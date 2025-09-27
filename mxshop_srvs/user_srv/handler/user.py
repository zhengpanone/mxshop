import time

import grpc
import peewee
from google.protobuf import empty_pb2
from passlib.handlers.pbkdf2 import pbkdf2_sha256

from common.proto.pb import user_pb2, user_pb2_grpc
from common.utils.page import make_page_response
from user_srv.model.models import User

from loguru import logger
from math import ceil


class UserServicer(user_pb2_grpc.UserServicer):

    def convert_user_to_rsp(self, user):
        user_info_rsp = user_pb2.UserResponse()
        user_info_rsp.id = user.id
        user_info_rsp.password = user.password
        user_info_rsp.mobile = user.mobile
        user_info_rsp.role = user.role
        if user.nickname:
            user_info_rsp.nickname = user.nickname
        if user.gender:
            user_info_rsp.gender = user.gender
        if user.birthday:
            user_info_rsp.birthday = int(time.mktime(user.birthday.timetuple()))
        return user_info_rsp

    @logger.catch
    def GetUserPageList(self, request: user_pb2.UserFilterPageInfo, context):
        rsp = user_pb2.UserPageResponse()
        pageSize = 10
        pageNum = 1
        page_request = request.pageRequest
        if page_request.pageSize:
            pageSize = page_request.pageSize
        if page_request.pageNum:
            pageNum = page_request.pageNum

        users = User.select().paginate(pageNum, pageSize)
        total = users.count()

        page_response = make_page_response(total=total,page=page_request)
        rsp.page.CopyFrom(page_response)

        for user in users:
            user_info_rsp = self.convert_user_to_rsp(user)
            rsp.list.append(user_info_rsp)

        return rsp

    @logger.catch
    def GetUserById(self, request, context):
        try:
            user = User.get(User.id == request.id)
            return self.convert_user_to_rsp(user)
        except peewee.DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details('用户不存在')
            return user_pb2.UserResponse()

    @logger.catch
    def GetUserByMobile(self, request, context):
        try:
            user = User.get(User.mobile == request.mobile)
            return self.convert_user_to_rsp(user)
        except peewee.DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details('用户不存在')
            return user_pb2.UserResponse()

    @logger.catch
    def CheckPassword(self, request: user_pb2.PasswordCheckRequest, context):
        return user_pb2.PasswordCheckRequest(success=pbkdf2_sha256.verify(request.password, request.encryptedPassword))

    @logger.catch
    def CreateUser(self, request:user_pb2.CreateUserRequest, context):
        try:
            User.get(User.mobile == request.mobile)

            context.set_code(grpc.StatusCode.ALREADY_EXISTS)
            context.set_details("用户已存在")
            return user_pb2.UserResponse()
        except peewee.DoesNotExist as e:
            pass

        user = User()
        user.nick_name = request.nickname
        user.mobile = request.mobile
        user.role = '1'
        user.password = pbkdf2_sha256.hash(request.password)
        user.save()

        return self.convert_user_to_rsp(user)

    @logger.catch
    def DeleteUserByIds(self, request, context):
        try:
            User.delete().where(User.id.in_(request.ids)).execute()
            context.set_code(grpc.StatusCode.OK)
            return empty_pb2.Empty()
        except Exception as e:
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(f"用户删除失败:{str(e)}")
            return empty_pb2.Empty()
