import grpc
from google.protobuf import empty_pb2
from loguru import logger
from peewee import DoesNotExist

from goods_srv.model.models import Goods, Category, Brands
from goods_srv.proto import goods_pb2, goods_pb2_grpc


class GoodsServicer(goods_pb2_grpc.GoodsServicer):

    def convert_model_to_message(self, goods):
        info_rsp = goods_pb2.GoodsInfoResponse()
        info_rsp.id = goods.id
        info_rsp.category_id = goods.category_id
        info_rsp.name = goods.name
        info_rsp.goods_sn = goods.goods_sn
        info_rsp.click_num = goods.click_num
        info_rsp.sold_num = goods.sold_num
        info_rsp.fav_num = goods.fav_num
        info_rsp.market_price = goods.market_price
        info_rsp.shop_price = goods.shop_price
        info_rsp.goods_brief = goods.goods_brief
        info_rsp.goods_front_image = goods.goods_front_image
        info_rsp.is_new = goods.is_new
        info_rsp.is_hot = goods.is_hot
        info_rsp.on_sale = goods.on_sale
        info_rsp.desc_images.extend(goods.desc_images)
        info_rsp.images.extend(goods.desc_images)
        info_rsp.category.id = goods.category.id
        info_rsp.category.name = goods.category.name

        info_rsp.brand.id = goods.brand.id
        info_rsp.brand.name = goods.brand.name
        info_rsp.brand.logo = goods.brand.logo

    @logger.catch
    def GoodsList(self, request: goods_pb2.GoodsFilterRequest, context):
        """商品列表页"""
        rsp = goods_pb2.GoodsListResponse()
        goods = Goods.select()
        if request.key_words:
            goods = goods.filter(Goods.name.constraints(request.key_words))
        if request.is_hot:
            goods = goods.filter(Goods.is_hot == True)
        if request.is_new:
            goods = goods.filter(Goods.is_new == True)
        if request.is_tab:
            goods = goods.filter(Goods.is_hot == True)
        if request.price_min:
            goods = goods.filter(Goods.shop_price >= request.price_min)
        if request.price_max:
            goods = goods.filter(Goods.shop_price <= request.price_max)
        if request.brand:
            goods = goods.filter(Goods.brand_id == request.brand)

        if request.top_category:
            # 通过category查询商品，这个category可能是一级、二级或者三级
            ids = []
            try:
                category = Category.get(Category.id == request.top_category)
                level = category.level
                if level == 2:
                    categorys = Category.select().where(Category.parent_category == request.top_category)
                    for category in categorys:
                        ids.append(category.id)
                elif level == 1:
                    c2 = Category.alias()
                    categorys = Category.select().where(Category.parent_category_id.in_(
                        c2.select(c2.id).where(c2.parent_category_id == request.top_category)))
                    for category in categorys:
                        ids.append(category.id)
                elif level == 3:
                    ids.append(request.top_category)

                goods = goods.where(Goods.category_id.in_(ids))
            except Exception as e:
                pass
            goods = goods.filter(Goods.is_hot == True)
        size = 10
        page = 1
        if request.page:
            page = request.page
        if request.size:
            size = request.size
        result = goods.paginate(page, size)
        rsp.total = result.count()
        for good in result:
            rsp.data.append(self.convert_model_to_message(good))

        return rsp

    @logger.catch
    def BatchGetGoods(self, request, context):
        # 批量获取商品信息
        rsp = goods_pb2.GoodsListResponse()
        goods = Goods.select().where(Goods.id.in_(list(request.id)))
        rsp.total = goods.count()
        for good in goods:
            rsp.data.append(self.convert_model_to_message(good))
        return rsp

    @logger.catch
    def DeleteGoods(self, request, context):
        """删除商品"""
        # goods = Goods.delete().where(Goods.id==request.id)
        try:
            Goods.get(Goods.id == request.id)
        except DoesNotExist as e:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("记录不存在")
            return empty_pb2.Empty()

    @logger.catch
    def GetGoodsDetail(self, request: goods_pb2.GoodInfoRequest, context):
        """获取商品详情"""
        try:
            goods = Goods.get(Goods.id == request.id)
            # 每次请求增加click_num
            goods.click_num += 1
            goods.save()
            return self.convert_model_to_message(goods)
        except DoesNotExist as e:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("商品不存在")
            return goods_pb2.GoodsInfoResponse()

    @logger.catch
    def CreateGoods(self, request: goods_pb2.CreateGoodsInfo, context):
        """新建商品"""
        try:
            category = Category.get(Category.id == request.category_id)
        except DoesNotExist as e:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("商品分类不存在，创建商品失败")
            return goods_pb2.GoodsInfoResponse()

        try:
            brand = Brands.get(Brands.id == request.brand_id)
        except DoesNotExist as e:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("品牌不存在，创建商品失败")
            return goods_pb2.GoodsInfoResponse()

        goods = Goods()
        goods.brand = brand
        goods.category = category
        goods.name = request.name
        goods.goods_sn = request.goods_sn
        goods.stocks = request.stocks
        goods.market_price = request.market_price
        goods.shop_price = request.shop_price
        goods.goods_brief = request.goods_brief
        goods.goods_desc = request.goods_desc
        goods.ship_free = request.ship_free
        goods.images = list(request.images)
        goods.desc_images = list(request.desc_images)
        goods.goods_front_image = request.goods_front_image
        goods.is_new = request.is_new
        goods.is_hot = request.is_hot
        goods.on_sale = request.on_sale
        goods.save()
        # TODO 此处完善库存的设置 - 分布式事务

        return self.convert_model_to_message(goods)

    @logger.catch
    def UpdateGoods(self, request: goods_pb2.CreateGoodsInfo, context):
        """更新商品"""
        try:
            category = Category.get(Category.id == request.category_id)
        except DoesNotExist as e:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("商品分类不存在，创建商品失败")
            return goods_pb2.GoodsInfoResponse()

        try:
            brand = Brands.get(Brands.id == request.brand_id)
        except DoesNotExist as e:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("品牌不存在，创建商品失败")
            return goods_pb2.GoodsInfoResponse()

        try:
            goods = Goods.get(Goods.id == request.id)
        except DoesNotExist as e:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("商品不存在，更新商品失败")
            return goods_pb2.GoodsInfoResponse()
        goods.brand = brand
        goods.category = category
        goods.name = request.name
        goods.goods_sn = request.goods_sn
        goods.stocks = request.stocks
        goods.market_price = request.market_price
        goods.shop_price = request.shop_price
        goods.goods_brief = request.goods_brief
        goods.goods_desc = request.goods_desc
        goods.ship_free = request.ship_free
        goods.images = list(request.images)
        goods.desc_images = list(request.desc_images)
        goods.goods_front_image = request.goods_front_image
        goods.is_new = request.is_new
        goods.is_hot = request.is_hot
        goods.on_sale = request.on_sale
        goods.save()
        # TODO 此处完善库存的设置 - 分布式事务
        return self.convert_model_to_message(goods)
