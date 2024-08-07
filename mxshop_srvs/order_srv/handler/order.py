import grpc
from google.protobuf import empty_pb2
from loguru import logger
from peewee import DoesNotExist
from order_srv.model.models import OrderGoods, ShoppingCart, OrderInfo
from order_srv.proto import order_pb2, order_pb2_grpc
from order_srv.proto import goods_pb2, goods_pb2_grpc
from order_srv.proto import inventory_pb2, inventory_pb2_grpc
from common.register import consul
from order_srv.settings import settings


def generate_order_sn(user_id):
    # 当前时间+user_id+随机数
    from random import Random
    import time
    random_ins = Random()
    return f"{time.strftime("%Y%m%d%H%M%S")}{user_id}{random_ins.randint(10, 99)}"


class OrderServicer(order_pb2_grpc.OrderServicer):

    @logger.catch
    def CarItemList(self, request: order_pb2.UserInfo, context):
        # 获取用户的购物信息
        items = ShoppingCart.select().where(ShoppingCart.user == request.id)
        rsp = order_pb2.CartItemListResponse(total=items.count())
        for item in items:
            item_rsp = order_pb2.ShopCartInfoResponse()
            item_rsp.id = item.id
            item_rsp.userId = item.user
            item_rsp.goodsId = item.goods
            item_rsp.nums = item.nums
            item_rsp.checked = item.checked
            rsp.data.append(item_rsp)
        return rsp

    @logger.catch
    def CreateCartItem(self, request: order_pb2.CartItemRequest, context):
        # 添加商品到购物车
        existed_items = ShoppingCart.select().where(ShoppingCart.goods == request.goodsId,
                                                    ShoppingCart.user == request.userId);
        if existed_items:
            item = existed_items[0]
            item.nums += request.nums
        else:
            item = ShoppingCart()
            item.user = request.userId
            item.goods = request.goodsId
            item.nums = request.nums
        item.save()
        return order_pb2.ShopCartInfoResponse(id=item.id)

    @logger.catch
    def UpdateCartItem(self, request: order_pb2.CartItemRequest, context):
        """更新购物车条目和选中状态"""
        try:
            item = ShoppingCart.get(ShoppingCart.user == request.userId, ShoppingCart.goods == request.goodsId)
            if request.nums:
                item.nums = request.nums
            item.checked = request.checked
            item.save()
            return empty_pb2.Empty()
        except DoesNotExist as e:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("购物车记录不存在")
            return empty_pb2.Empty()

    @logger.catch
    def DeleteCartItem(self, request: order_pb2.CartItemRequest, context):
        """删除购物车条目"""
        try:
            item = ShoppingCart.get(ShoppingCart.user == request.userId, ShoppingCart.goods == request.goodsId)
            item.delete_instance()
            return empty_pb2.Empty()
        except DoesNotExist as e:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("购物车记录不存在")
            return empty_pb2.Empty()

    @logger.catch
    def OrderList(self, request, context):
        """订单列表"""
        rsp = order_pb2.OrderListResponse()
        orders = OrderInfo.select()
        if request.userId:
            orders = orders.where(OrderInfo.user == request.userId)
        rsp.total = orders.count()
        # 分页
        page = request.page if request.page else 1
        size = request.size if request.size else 10
        offset = (page - 1) * size
        orders = orders.limit(size).offset(offset)
        for order in orders:
            tmp_rsp = order_pb2.OrderInfoResponse()
            tmp_rsp.id = order.id
            tmp_rsp.userId = order.user
            tmp_rsp.orderSn = order.order_sn
            tmp_rsp.payType = order.pay_type
            tmp_rsp.status = order.status
            tmp_rsp.post = order.post
            tmp_rsp.total = order.order_amount
            tmp_rsp.address = order.address
            tmp_rsp.name = order.signer_name
            tmp_rsp.mobile = order.signer_mobile
            tmp_rsp.addTime = order.add_time.strftime('%Y-%m-%d %H:%M:%S')
            rsp.data.append(tmp_rsp)
        return rsp

    @logger.catch
    def OrderDetail(self, request: order_pb2.OrderRequest, context):
        """订单详情"""
        rsp = order_pb2.OrderInfoDetailResponse()
        try:
            # 订单信息
            if request.userId:
                order = OrderInfo.get(OrderInfo.id == request.id, OrderInfo.user == request.userId)
            else:
                order = OrderInfo.get(OrderInfo.id == request.id)
            rsp.orderInfo.id = order.id
            rsp.orderInfo.userId = order.user
            rsp.orderInfo.orderSn = order.order_sn
            rsp.orderInfo.payType = order.pay_type
            rsp.orderInfo.status = order.status
            rsp.orderInfo.post = order.post
            rsp.orderInfo.total = order.order_amount
            rsp.orderInfo.address = order.address
            rsp.orderInfo.name = order.signer_name
            rsp.orderInfo.mobile = order.signer_mobile
            # 商品信息
            order_goods = OrderGoods.select().where(OrderGoods.order == order.id)
            for order_good in order_goods:
                order_goods_rsp = order_pb2.OrderItemResponse()
                order_goods_rsp.id = order_good.id
                order_goods_rsp.orderId = order.id
                order_goods_rsp.goodsId = order_good.goods
                order_goods_rsp.goodsName = order_good.goods_name
                order_goods_rsp.goodsImage = order_good.goods_image
                order_goods_rsp.goodsPrice = float(order_good.goods_price)
                order_goods_rsp.nums = order_good.nums
                rsp.goods.append(order_goods_rsp)
            return rsp
        except DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("订单记录不存在")
            return empty_pb2.Empty()

    @logger.catch
    def UpdateOrderStatus(self, request, context):
        """更新订单支付状态"""
        OrderInfo.update(status=request.status).where(OrderInfo.order_sn == request.orderSn).execute()
        return empty_pb2.Empty()

    @logger.catch
    def CreateOrder(self, request, context):
        """创建订单
        1、价格 -- 访问商品服务
        2、库存扣减 -- 访问库存服务
        3、订单基本信息表 -- 订单商品信息表
        4、从购物车中获取选到的商品
        5、从购物车中删除已购买的商品
        """
        with settings.DB.atomic() as txn:
            goods_ids = []
            goods_nums = {}
            order_goods_list = []
            order_amount = 0
            goods_sell_info = []
            # 获取购物车记录
            for cart_item in ShoppingCart.select().where(ShoppingCart.user == request.userId,
                                                         ShoppingCart.checked == True):
                goods_ids.append(cart_item.id)
                goods_nums[cart_item.goods] = cart_item.nums

            if not goods_ids:
                context.set_code(grpc.StatusCode.NOT_FOUND)
                context.set_details("没有选中结算的商品")
                return order_pb2.OrderInfoResponse()

            # 查询商品信息
            register = consul.ConsulRegister(settings.CONSUL_HOST, settings.CONSUL_PORT)
            goods_srv_host, goods_srv_port = register.get_host_port(f'Service == "{settings.GOODS_SRV_NAME}"')
            if not goods_srv_host:
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details("商品服务不可用")
                return order_pb2.OrderInfoResponse()

            goods_channel = grpc.insecure_channel(f"{goods_srv_host}:{goods_srv_port}")
            goods_stub = goods_pb2_grpc.GoodsStub(goods_channel)
            # 批量获取商品信息
            try:
                goods_info_rsp = goods_stub.BatchGetGoods(goods_pb2.BatchGoodsIdInfo(id=goods_ids))
            except grpc.RpcError as e:
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(f"商品服务不可用:{str(e)}")
                return order_pb2.OrderInfoResponse()
            for goods_info in goods_info_rsp.data:
                count = goods_nums[goods_info.id]
                order_amount += goods_info.shopPrice * count
                order_goods = OrderGoods(goods=goods_info.id, goods_name=goods_info.name,
                                         goods_image=goods_info.goodsFrontImage,
                                         goods_price=goods_info.shopPrice, nums=count)
                order_goods_list.append(order_goods)
                goods_sell_info.append(inventory_pb2.GoodsInvInfo(goodsId=goods_info.id, num=count))

            # 扣减库存
            # TODO 负载均衡
            inv_srv_host, inv_srv_port = register.get_host_port(f'Service == "{settings.INVENTORY_SRV_NAME}"')
            if not inv_srv_host:
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details("库存服务不可用")
                return order_pb2.OrderInfoResponse()

            inv_channel = grpc.insecure_channel(f"{inv_srv_host}:{inv_srv_port}")
            inv_stub = inventory_pb2_grpc.InventoryStub(inv_channel)
            try:
                inv_stub.Sell(inventory_pb2.SellInfo(goodsInfo=goods_sell_info))
            except grpc.RpcError as e:
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(f"扣减库存失败:{str(e)}")
                return order_pb2.OrderInfoResponse()
            try:
                # 创建订单
                order = OrderInfo()
                order.order_sn = generate_order_sn(request.userId)
                order.order_amount = order_amount
                order.address = request.address
                order.signer_name = request.name
                order.signer_mobile = request.mobile
                order.user = request.userId
                order.post = request.post
                order.save()

                # 批量插入订单商品表
                for order_good in order_goods_list:
                    order_good.order = order.id
                OrderGoods.bulk_create(order_goods_list)

                # 删除购物车的记录
                ShoppingCart.delete().where(ShoppingCart.user == request.userId, ShoppingCart.checked == True).execute()
            except Exception as e:
                txn.rollback()
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(f"订单创建失败:{str(e)}")
                return order_pb2.OrderInfoResponse()
        return order_pb2.OrderInfoResponse(id=order.id, orderSn=order.order_sn, total=order_amount)
