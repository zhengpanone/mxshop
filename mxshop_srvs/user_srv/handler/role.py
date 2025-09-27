import grpc
import peewee

from user_srv.model.models import Role
from common.proto.pb import role_pb2_grpc,role_pb2
from loguru import logger


class RoleServicer(role_pb2_grpc.RoleServicer):

    def convert_role_to_rsp(self, role):
        role_info_rsp = role_pb2.RoleInfoResponse()
        role_info_rsp.id = role.id
        role_info_rsp.name = role.name
        role_info_rsp.remark = role.description
        role_info_rsp.status = role.status
        return role_info_rsp


    @logger.catch
    def CreateRole(self, request:role_pb2.CreateRoleInfo, context):
        try:
            Role.get(Role.name == request.name)
            context.set_code(grpc.StatusCode.ALREADY_EXISTS)
            context.set_details('Role already exists')
            return role_pb2.RoleInfoResponse()
        except peewee.DoesNotExist as e:
            pass
        role = Role()
        role.name=request.name
        role.description=request.remark
        role.status=request.status
        role.save()
        return self.convert_role_to_rsp(role)
        pass