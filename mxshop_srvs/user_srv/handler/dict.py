import peewee

import grpc
from loguru import logger
from peewee import IntegrityError

from common.proto.pb import dict_pb2_grpc,dict_pb2
from user_srv.model.models import DictType


class DictServicer(dict_pb2_grpc.DictServicer):
    def convert_dict_type_to_rsp(self, dict_type:DictType)->dict_pb2.DictTypeResponse:
        dict_type_rsp = dict_pb2.DictTypeResponse()
        dict_type_rsp.id = dict_type.id
        dict_type_rsp.dictName = dict_type.dict_name
        dict_type_rsp.dictCode = dict_type.dict_code
        # dict_type_rsp.status = dict_type.status
        return dict_type_rsp

    @logger.catch
    def CreateDictType(self, request:dict_pb2.CreateDictTypeRequest, context)->dict_pb2.DictTypeResponse:
        try:
            DictType.get(DictType.dict_code==request.dictCode)
            context.set_details(f'DictType with code {request.dictCode} already exists.')
            context.set_code(grpc.StatusCode.ALREADY_EXISTS)
            return dict_pb2.DictTypeResponse()
        except peewee.DoesNotExist as e:
            pass
        except Exception as e:
            # 其他异常
            context.set_details(f'Error checking DictType existence: {str(e)}')
            context.set_code(grpc.StatusCode.INTERNAL)
            return dict_pb2.DictTypeResponse()
        # 创建字典类型
        dict_type = DictType()
        dict_type.dict_code = request.dictCode
        dict_type.dict_name = request.dictName
        try:
            dict_type.save()
        except IntegrityError as e:
            context.set_details(f'Integrity error: {str(e)}')
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            return dict_pb2.DictTypeResponse()
        except Exception as e:
            context.set_details(f'Failed to save DictType: {str(e)}')
            context.set_code(grpc.StatusCode.INTERNAL)
            return dict_pb2.DictTypeResponse()

        return self.convert_dict_type_to_rsp(dict_type)