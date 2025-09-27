import grpc
from google.protobuf import empty_pb2
from loguru import logger
from peewee import DoesNotExist
import json

from common.utils.page import make_page_response
from goods_srv.model.models import Goods, Category, Brands, GoodsCategoryBrand, Banner
from common.proto.pb import goods_pb2, goods_pb2_grpc,common_pb2


class GoodsServicer(goods_pb2_grpc.GoodsServicer):

    def convert_model_to_message(self, goods):
        info_rsp = goods_pb2.GoodsResponse()
        info_rsp.id = goods.id
        info_rsp.categoryId = goods.category_id
        info_rsp.name = goods.name
        info_rsp.goodsSn = goods.goods_sn
        info_rsp.clickNum = goods.click_num
        info_rsp.soldNum = goods.sold_num
        info_rsp.favNum = goods.fav_num
        info_rsp.marketPrice = goods.market_price
        info_rsp.shopPrice = goods.shop_price
        info_rsp.goodsBrief = goods.goods_brief
        info_rsp.goodsFrontImage = goods.goods_front_image
        info_rsp.isNew = goods.is_new
        info_rsp.isHot = goods.is_hot
        info_rsp.onSale = goods.on_sale
        info_rsp.descImages.extend(goods.desc_images)
        info_rsp.images.extend(goods.desc_images)
        info_rsp.category.id = goods.category.id
        info_rsp.category.name = goods.category.name

        info_rsp.brand.id = goods.brand.id
        info_rsp.brand.name = goods.brand.name
        info_rsp.brand.logo = goods.brand.logo
        return info_rsp

    @logger.catch
    def GoodsPageList(self, request: goods_pb2.GoodsFilterPageRequest, context):
        """商品列表页"""
        rsp = goods_pb2.GoodsListResponse()
        goods = Goods.select()
        if request.keyWords:
            goods = goods.filter(Goods.name.contains(request.keyWords))
        if request.isHot:
            goods = goods.filter(Goods.is_hot == True)
        if request.isNew:
            goods = goods.filter(Goods.is_new == True)
        if request.isTab:
            goods = goods.filter(Goods.is_hot == True)
        if request.priceMin:
            goods = goods.filter(Goods.shop_price >= request.priceMin)
        if request.priceMax:
            goods = goods.filter(Goods.shop_price <= request.priceMax)
        if request.brand:
            goods = goods.filter(Goods.brand_id == request.brand)

        if request.topCategory:
            # 通过category查询商品，这个category可能是一级、二级或者三级
            ids = []
            try:
                category = Category.get(Category.id == request.topCategory)
                level = category.level
                if level == 2:
                    categorys = Category.select().where(Category.parent_category == request.topCategory)
                    for category in categorys:
                        ids.append(category.id)
                elif level == 1:
                    c2 = Category.alias()
                    categorys = Category.select().where(Category.parent_category_id.in_(
                        c2.select(c2.id).where(c2.parent_category_id == request.topCategory)))
                    for category in categorys:
                        ids.append(category.id)
                elif level == 3:
                    ids.append(request.top_category)

                goods = goods.where(Goods.category_id.in_(ids))
            except Exception as e:
                pass
            goods = goods.filter(Goods.is_hot == True)
        page_num = 1
        page_size = 10
        page_request= request.pageRequest
        if page_request.pageNum:
            page_num = page_request.pageNum
        if page_request.pageSize:
            page_size = page_request.pageSize
        total = goods.count()
        result = goods.paginate(page_num, page_size)

        for good in result:
            rsp.list.append(self.convert_model_to_message(good))
        page_response = make_page_response(total, page_request)
        rsp.page.CopyFrom(page_response)

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
    def GetGoodsDetail(self, request: goods_pb2.GoodsDetailRequest, context):
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
    def CreateGoods(self, request: goods_pb2.CreateGoodsRequest, context):
        """新建商品"""
        try:
            category = Category.get(Category.id == request.categoryId)
        except DoesNotExist as e:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("商品分类不存在，创建商品失败")
            return goods_pb2.GoodsInfoResponse()

        try:
            brand = Brands.get(Brands.id == request.brandId)
        except DoesNotExist as e:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("品牌不存在，创建商品失败")
            return goods_pb2.GoodsInfoResponse()

        goods = Goods()
        goods.brand = brand
        goods.category = category
        goods.name = request.name
        goods.goods_sn = request.goodsSn
        goods.stocks = request.stocks
        goods.market_price = request.marketPrice
        goods.shop_price = request.shopPrice
        goods.goods_brief = request.goodsBrief
        goods.goods_desc = request.goodsDesc
        goods.ship_free = request.shipFree
        goods.images = list(request.images)
        goods.desc_images = list(request.descImages)
        goods.goods_front_image = request.goodsFrontImage
        goods.is_new = request.isNew
        goods.is_hot = request.isHot
        goods.on_sale = request.onSale
        goods.save()
        # TODO 此处完善库存的设置 - 分布式事务

        return self.convert_model_to_message(goods)

    @logger.catch
    def UpdateGoods(self, request: goods_pb2.UpdateGoodsRequest, context):
        """更新商品"""
        try:
            goods = Goods.get(Goods.id == request.id)
        except DoesNotExist as e:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("商品不存在，更新商品失败")
            return goods_pb2.GoodsInfoResponse()

        if request.categoryId:
            try:
                category = Category.get(Category.id == request.categoryId)
                goods.category = category
            except DoesNotExist as e:
                context.set_code(grpc.StatusCode.NOT_FOUND)
                context.set_details("商品分类不存在，创建商品失败")
                return goods_pb2.GoodsInfoResponse()
        if request.brandId:
            try:
                brand = Brands.get(Brands.id == request.brandId)
                goods.brand = brand
            except DoesNotExist as e:
                context.set_code(grpc.StatusCode.NOT_FOUND)
                context.set_details("品牌不存在，创建商品失败")
                return goods_pb2.GoodsInfoResponse()

        goods.name = request.name
        goods.goods_sn = request.goodsSn
        goods.stocks = request.stocks
        goods.market_price = request.marketPrice
        goods.shop_price = request.shopPrice
        goods.goods_brief = request.goodsBrief
        goods.goods_desc = request.goodsDesc
        goods.ship_free = request.shipFree
        goods.images = list(request.images)
        goods.desc_images = list(request.descImages)
        goods.goods_front_image = request.goodsFrontImage
        goods.is_new = request.isNew
        goods.is_hot = request.isHot
        goods.on_sale = request.onSale
        goods.save()
        # TODO 此处完善库存的设置 - 分布式事务
        return self.convert_model_to_message(goods)

    def category_model_to_dict(self, category: Category) -> dict:
        re = {}
        re["id"] = category.id
        re["name"] = category.name
        re["level"] = category.level
        re["parent"] = category.parent_category_id
        re['is_tab'] = category.is_tab
        return re

    @logger.catch
    def GetAllCategoryList(self, request: empty_pb2.Empty, context):
        # 商品分类
        """

        :param request:
        :param context:
        :return:  [{"name":"xxx","id":"xxx","sub_category":[{},{}]},{},{}]
        """
        level1 = []
        level2 = []
        level3 = []
        category_list_rsp = goods_pb2.CategoryListResponse()
        category_list_rsp.total = Category.select().count()
        for category in Category.select():
            category_rsp = goods_pb2.CategoryResponse()
            category_rsp.id = category.id
            category_rsp.name = category.name
            category_rsp.parentCategory = category.parent_category_id
            category_rsp.level = category.level
            category_rsp.isTab = category.is_tab
            category_list_rsp.data.append(category_rsp)
            if category.level == 1:
                level1.append(self.category_model_to_dict(category))
            elif category.level == 2:
                level2.append(self.category_model_to_dict(category))
            elif category.level == 3:
                level3.append(self.category_model_to_dict(category))

        for data3 in level3:
            for data2 in level2:
                if data3['parent'] == data2['id']:
                    if "sub_category" not in data2:
                        data2['sub_category'] = [data3]
                    else:
                        data2['sub_category'].append(data3)

        for data2 in level2:
            for data1 in level1:
                if data2['parent'] == data1['id']:
                    if "sub_category" not in data1:
                        data1['sub_category'] = [data2]
                    else:
                        data1['sub_category'].append(data2)
        category_list_rsp.jsonData = json.dumps(level1)

        return category_list_rsp

    @logger.catch
    def GetSubCategory(self, request: goods_pb2.CategoryListRequest, context):
        category_list_rsp = goods_pb2.SubCategoryListResponse()
        try:
            category_info = Category.get(Category.id == request.id)
            category_rsp = goods_pb2.CategoryInfoResponse()
            category_rsp.id = category_info.id
            category_rsp.name = category_info.name
            category_rsp.level = category_info.level
            category_rsp.is_tab = category_info.is_tab
            if category_info.parent_category:
                category_list_rsp.info.parent_category = category_info.parent_category_id
        except DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details('商品子分类记录不存在')
            return empty_pb2.Empty()

        categorys = Category.select().where(Category.parent_category == request.id)
        for category in categorys:
            category_rsp = goods_pb2.CategoryResponse()
            category_rsp.id = category.id
            category_rsp.name = category.name
            if category_info.parent_category:
                category_rsp.parent_category = category_info.parent_category_id
            category_rsp.level = category.level
            category_rsp.is_tab = category.is_tab

            category_list_rsp.sub_category_list.append(category_rsp)
        return category_list_rsp

    @logger.catch
    def CreateCategory(self, request:goods_pb2.CreateCategoryRequest, context):
        try:
            category = Category()
            category.name = request.name
            if request.level != 1:
                category.parent_category = request.parent_category
            category.level = request.level
            category.is_tab = request.is_tab
            category.save()
            category_rsp = goods_pb2.CategoryResponse()
            category_rsp.id = category.id
            category_rsp.name = category.name
            if category_rsp.parent_category:
                category_rsp.parent_category = category_rsp.parent_category.id
            category_rsp.level = category_rsp.level
            category_rsp.is_tab = category_rsp.is_tab

        except Exception as e:
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details("商品分类新增数据失败:" + str(e))
            return goods_pb2.CategoryResponse()
        return category_rsp

    @logger.catch
    def DeleteCategory(self, request:common_pb2.IdsRequest, context):
        try:
            category = Category.get(request.id)
            category.delete_instance()
            # TODO 删除响应的category下的商品
            return empty_pb2.Empty()
        except DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("删除商品分类时，记录不存在")
            return empty_pb2.Empty()

    @logger.catch
    def UpdateCategory(self, request, context):
        try:
            category = Category.get(request.id)
            if request.name:
                category.name = request.name
            if request.parent_category:
                category.parent_category = request.parent_category
            if request.level:
                category.level = request.level
            if request.is_tab:
                category.is_tab = request.is_tab

            category.save()

            return empty_pb2.Empty()
        except DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("更新商品分类时，记录不存在")
            return empty_pb2.Empty()

    @logger.catch
    def BrandPageList(self, request: goods_pb2.BrandFilterPageRequest, context):
        rsp = goods_pb2.BrandListResponse()
        brands = Brands.select()
        rsp.total = brands.count()
        result = brands.paginate(request.page, request.size)
        for brand in result:
            brand_rsp = goods_pb2.BrandInfoResponse()
            brand_rsp.id = brand.id
            brand_rsp.name = brand.name
            brand_rsp.logo = brand.logo
            rsp.data.append(brand_rsp)
        return rsp

    @logger.catch
    def CreateBrand(self, request: goods_pb2.CreateBrandRequest, context):
        brands = Brands.select().where(Brands.name == request.name)
        if brands:
            context.set_code(grpc.StatusCode.ALREADY_EXISTS)
            context.set_details('创建brand时，品牌已经存在')
            return goods_pb2.BrandInfoResponse()
        brand = Brands()
        brand.name = request.name
        brand.logo = request.logo
        brand.save()

        brand_rsp = goods_pb2.BrandInfoResponse()
        brand_rsp.id = brand.id
        brand_rsp.name = brand.name
        brand_rsp.logo = brand.logo
        return brand_rsp

    @logger.catch
    def DeleteBrand(self, request: common_pb2.IdsRequest, context):
        try:
            banner = Brands.get(request.id)
            banner.delete_innstance()
            return empty_pb2.Empty()
        except DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("删除brand时，记录不存在")
            return empty_pb2.Empty()

    @logger.catch
    def UpdateBrand(self, request: goods_pb2.UpdateBrandRequest, context):
        try:
            brand = Brands.get(request.id)
            if request.name:
                brand.name = request.name
            if request.logo:
                brand.logo = request.logo
            brand.save()
            return empty_pb2.Empty()
        except DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("更新brand时，记录不存在")
            return empty_pb2.Empty()

    @logger.catch
    def BannerPageList(self, request, context):
        rsp = goods_pb2.BannerListResponse()
        banners = Banner.select()
        rsp.total = banners.count()
        for banner in banners:
            banner_rsp = goods_pb2.BannerResponse()
            banner_rsp.id = banner.id
            banner_rsp.image = banner.image
            banner_rsp.index = banner.index
            banner_rsp.url = banner.url
            rsp.data.append(banner_rsp)
        return rsp

    @logger.catch
    def CreateBanner(self, request: goods_pb2.CreateBannerRequest, context):
        banner = Banner()
        banner.image = request.image
        banner.index = request.index
        banner.url = request.url
        banner.save()

        banner_rsp = goods_pb2.BannerResponse()
        banner_rsp.id = banner.id
        banner_rsp.image = banner.image
        banner_rsp.index = banner.index
        banner_rsp.url = banner.url

        return banner_rsp

    @logger.catch
    def DeleteBanner(self, request: common_pb2.IdsRequest, context):
        try:
            banner = Banner.get(request.id)
            banner.delete_innstance()
            return empty_pb2.Empty()
        except DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("删除banner时，记录不存在")
            return empty_pb2.Empty()

    @logger.catch
    def UpdateBanner(self, request: goods_pb2.UpdateBannerRequest, context):
        try:
            banner = Banner.get(request.id)
            if request.image:
                banner.image = request.image
            if request.index:
                banner.index = request.index
            if request.url:
                banner.url = request.url
            banner.save()
            return empty_pb2.Empty()
        except DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("更新banner时，记录不存在")
            return empty_pb2.Empty()

    @logger.catch
    def CategoryBrandPageList(self, request: goods_pb2.CategoryBrandFilterRequest, context):
        rsp = goods_pb2.CategoryBrandListResponse()
        category_brands = GoodsCategoryBrand.select()
        page = 1
        size = 10
        if request.page:
            page = request.page
        if request.size:
            size = request.size
        rsp.total = category_brands.count()
        category_brands = category_brands.paginate(page, size)
        for category_brand in category_brands:
            category_brand_rsp = goods_pb2.CategoryBrandResponse()
            category_brand_rsp.id = category_brand.id
            category_brand_rsp.brand.name = category_brand.brand.name
            category_brand_rsp.brand.logo = category_brand.brand.logo

            category_brand_rsp.category.id = category_brand.category.id
            category_brand_rsp.category.name = category_brand.category.name
            category_brand_rsp.category.parentCategory = category_brand.category.parent_category_id
            category_brand_rsp.category.level = category_brand.category.level
            category_brand_rsp.category.isTab = category_brand.category.is_tab
            rsp.data.append(category_brand_rsp)
        return rsp

    @logger.catch
    def GetCategoryBrandList(self, request: goods_pb2.CategoryBrandFilterRequest, context):
        # 获取某一个分类的所有品牌
        rsp = goods_pb2.BrandListResponse()
        try:
            category = Category.get(Category.id == request.id)
            category_brands = GoodsCategoryBrand.select().where(GoodsCategoryBrand.category == category)
            rsp.total = category_brands.count()
            for category_brand in category_brands:
                brand_rsp = goods_pb2.BrandInfoResponse()
                brand_rsp.id = category_brand.brand.id
                brand_rsp.name = category_brand.brand.name
                brand_rsp.logo = category_brand.brand.logo
                rsp.data.append(brand_rsp)
        except DoesNotExist:
            context.set_code(grpc.StatusCode.ALREADY_EXISTS)
            context.set_details('获取分类下的品牌，记录存在')
            return rsp
        return rsp

    @logger.catch
    def CreateCategoryBrand(self, request: goods_pb2.CreateCategoryRequest, context):
        category_brand = GoodsCategoryBrand()
        try:
            brand = Brands.get(request.brand_id)
            category_brand.brand = brand
            category = Category.get(request.categoryId)
            category_brand.category = category
            category_brand.save()
            rsp = goods_pb2.CategoryBrandResponse()
            rsp.id = category_brand.id
            return rsp
        except DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("商品品牌记录不存在")
            return goods_pb2.CategoryBrandResponse()

        except Exception as e:
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details("内部错误")
            return goods_pb2.CategoryBrandResponse()

    @logger.catch
    def DeleteCategoryBrand(self, request: common_pb2.IdsRequest, context):
        try:
            category_brand = GoodsCategoryBrand.get(request.id)
            category_brand.delete_instance()
            return empty_pb2.Empty()
        except DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("删除商品分类时，记录不存在")
            return empty_pb2.Empty()

    @logger.catch
    def UpdateCategoryBrand(self, request: goods_pb2.UpdateCategoryBrandRequest, context):
        try:
            category_brand = GoodsCategoryBrand.get(request.id)
            brand = Brands.get(request.brand_id)
            category_brand.brand = brand
            category = Category.get(request.categoryId)
            category_brand.category = category
            category_brand.save()

            return empty_pb2.Empty()
        except DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("更新商品分类时，记录不存在")
            return empty_pb2.Empty()
        except Exception as e:
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details("内部错误")
            return empty_pb2.Empty()
