import time

import grpc
import peewee
from passlib.handlers.pbkdf2 import pbkdf2_sha256

from user_srv.proto import user_pb2, user_pb2_grpc
from user_srv.model.models import User

from loguru import logger


class UserServicer(user_pb2_grpc.UserServicer):

    def convert_user_to_rsp(self, user):
        user_info_rsp = user_pb2.UserInfoResponse()
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
    def GetUserList(self, request: user_pb2.PageInfo, context):
        rsp = user_pb2.UserListResponse()
        size = 10
        page = 1
        if request.size:
            size = request.size
        if request.page:
            page = request.page

        users = User.select().paginate(page, size)
        rsp.total = users.count()

        for user in users:
            user_info_rsp = self.convert_user_to_rsp(user)
            rsp.data.append(user_info_rsp)

        return rsp

    @logger.catch
    def GetUserById(self, request, context):
        try:
            user = User.get(User.id == request.id)
            return self.convert_user_to_rsp(user)
        except peewee.DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details('用户不存在')
            return user_pb2.UserInfoResponse()

    @logger.catch
    def GetUserByMobile(self, request, context):
        try:
            user = User.get(User.mobile == request.mobile)
            return self.convert_user_to_rsp(user)
        except peewee.DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details('用户不存在')
            return user_pb2.UserInfoResponse()

    @logger.catch
    def CheckPassword(self, request: user_pb2.PasswordCheckInfo, context):
        return user_pb2.CheckPasswordResponse(success=pbkdf2_sha256.verify(request.password, request.encryptedPassword))

    @logger.catch
    def CreateUser(self, request:user_pb2.CreateUserInfo, context):
        try:
            User.get(User.mobile == request.mobile)

            context.set_code(grpc.StatusCode.ALREADY_EXISTS)
            context.set_details("用户已存在")
            return user_pb2.UserInfoResponse()
        except peewee.DoesNotExist as e:
            pass

        user = User()
        user.nick_name = request.nickname
        user.mobile = request.mobile
        user.password = pbkdf2_sha256.hash(request.password)
        user.save()

        return self.convert_user_to_rsp(user)
