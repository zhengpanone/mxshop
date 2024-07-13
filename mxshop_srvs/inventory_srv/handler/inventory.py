import grpc
from google.protobuf import empty_pb2
from loguru import logger
from peewee import DoesNotExist
import json
from inventory_srv.model.models import Inventory
from inventory_srv.proto import inventory_pb2, inventory_pb2_grpc
from inventory_srv.settings import settings


class InventoryServicer(inventory_pb2_grpc.InventoryServicer):
    @logger.catch
    def Sell(self, request:inventory_pb2.SellInfo, context):
        # 扣减库存
        with settings.DB.atomic() as txn:
            for item in request.goodsInfo:
                # 查询库存
                try:
                    goods_inv = Inventory.get(Inventory.goods == item.goodsId)
                except DoesNotExist as e:
                    txn.rollback() # 事务回滚
                    context.set_code(grpc.StatusCode.NOT_FOUND)
                    context.set_details('商品未找到')
                    return empty_pb2.Empty()
                if goods_inv.stocks < item.num:
                    # 库存不足
                    context.set_code(grpc.StatusCode.RESOURCE_EXHAUSTED)
                    context.set_details("库存不足")
                    txn.rollback()  # 事务回滚
                    return empty_pb2.Empty()
                else:
                    # TODO 这里可能会引起数据不一致 --分布式锁
                    goods_inv.stocks -= item.num
                    goods_inv.save()
        return empty_pb2.Empty()

    @logger.catch
    def SetInv(self, request: inventory_pb2.GoodsInvInfo, context):
        # 设置库存，后面修改库存可以可以使用这个接口
        force_insert = False
        invs = Inventory.select().where(Inventory.goods == request.goods_id)
        if not invs:
            inv = Inventory()
            inv.goods = request.goods_id
            force_insert = True
        else:
            inv = invs[0]
        inv.stocks = request.num
        inv.save(force_insert=force_insert)
        return empty_pb2.Empty()

    @logger.catch
    def InvDetail(self, request: inventory_pb2.GoodsInvInfo, context):
        # 获取某个商品的库存详情
        try:
            inv = Inventory.get(Inventory.goods == request.goods_id)
            return inventory_pb2.GoodsInvInfo(goods_id=inv.goods, num=inv.stocks)
        except DoesNotExist as e:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("没有找到该商品的库存记录")
            return inventory_pb2.GoodsInvInfo()

    @logger.catch
    def Reback(self, request, context):
        # 库存归还，有两张情况：1、订单自动归还，2订单创建失败 3 手动归还
        with settings.DB.atomic() as txn:
            for item in request.goodsInfo:
                # 查询库存
                try:
                    goods_inv = Inventory.get(Inventory.goods==item.goods_id)
                except DoesNotExist as e:
                    txn.rollback() # 事务回滚
                    context.set_code(grpc.StatusCode.NOT_FOUND)
                    context.set_details('商品未找到')
                    return empty_pb2.Empty()
                # TODO 这里可能会引起数据不一致 -- 分布式锁
                goods_inv.stocks += item.num
                goods_inv.save()
            return empty_pb2.Empty()
