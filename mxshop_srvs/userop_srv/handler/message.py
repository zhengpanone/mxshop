from common.proto.pb import message_pb2_grpc, message_pb2
from userop_srv.model.models import LeavingMessages

from loguru import logger


class MessageServicer(message_pb2_grpc.MessageServicer):
    @logger.catch
    def MessageList(self, request: message_pb2.MessageRequest, context):
        rsp = message_pb2.MessageListResponse()
        messages = LeavingMessages.select()
        if request.userId:
            messages = messages.where(LeavingMessages.user == request.userId)

        rsp.total = messages.count()
        for message in messages:
            brand_rsp = message_pb2.MessageResponse()
            brand_rsp.id = message.id
            brand_rsp.userId = message.user
            brand_rsp.subject = message.subject
            brand_rsp.message = message.message
            brand_rsp.file = message.file
            rsp.data.append(brand_rsp)

        return rsp

    def CreateMessage(self, request: message_pb2.MessageRequest, context):
        message = LeavingMessages()
        message.user = request.userId
        message.message_type = request.messageType
        message.subject = request.subject
        message.message = request.message
        message.file = request.file

        message.save()

        rsp = message_pb2.MessageResponse()
        rsp.id = message.id
        rsp.userId = message.user
        rsp.subject = message.subject
        rsp.message = message.message
        rsp.file = message.file
        rsp.messageType = message.message_type
