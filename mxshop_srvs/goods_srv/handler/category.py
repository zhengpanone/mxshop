import json

import grpc
from loguru import logger
from peewee import DoesNotExist

from goods_srv.model.models import Category
from goods_srv.proto import category_pb2, category_pb2_grpc
from google.protobuf import empty_pb2


class CategoryServicer(category_pb2_grpc.CategoryServicer):

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
        category_list_rsp = category_pb2.CategoryListResponse()
        category_list_rsp.total = Category.select().count()
        for category in Category.select():
            category_rsp = category_pb2.CategoryInfoResponse()
            category_rsp.id = category.id
            category_rsp.name = category.name
            category_rsp.parent_category = category.parent_category_id
            category_rsp.level = category.level
            category_rsp.is_tab = category.is_tab
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
        category_list_rsp.json_data = json.dumps(level1)

        return category_list_rsp

    @logger.catch
    def GetSubCategory(self, request: category_pb2.CategoryListRequest, context):
        category_list_rsp = category_pb2.SubCategoryListResponse()
        try:
            category_info = Category.get(Category.id == request.id)
            category_rsp = category_pb2.CategoryInfoResponse()
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
            category_rsp = category_pb2.CategoryInfoResponse()
            category_rsp.id = category.id
            category_rsp.name = category.name
            if category_info.parent_category:
                category_rsp.parent_category = category_info.parent_category_id
            category_rsp.level = category.level
            category_rsp.is_tab = category.is_tab

            category_list_rsp.sub_category_list.append(category_rsp)
        return category_list_rsp

    @logger.catch
    def CreateCategory(self, request, context):
        try:
            category = Category()
            category.name = request.name
            if request.level != 1:
                category.parent_category = request.parent_category
            category.level = request.level
            category.is_tab = request.is_tab
            category.save()
            category_rsp = category_pb2.CategoryInfoResponse()
            category_rsp.id = category.id
            category_rsp.name = category.name
            if category_rsp.parent_category:
                category_rsp.parent_category = category_rsp.parent_category.id
            category_rsp.level = category_rsp.level
            category_rsp.is_tab = category_rsp.is_tab

        except Exception as e:
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details("商品分类新增数据失败:" + str(e))
            return category_pb2.CategoryInfoResponse()
        return category_rsp

    @logger.catch
    def DeleteCategory(self, request, context):
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
