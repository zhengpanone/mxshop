import grpc
from peewee import DoesNotExist

from userop_srv.proto import address_pb2, address_pb2_grpc
from userop_srv.model.models import Address
from google.protobuf import empty_pb2
from loguru import logger


class AddressServicer(address_pb2_grpc.AddressServicer):
    @logger.catch
    def GetAddressList(self, request: address_pb2.AddressRequest, context):
        rsp = address_pb2.AddressListResponse()
        address = Address.select()
        if request.userId:
            address = address.where(Address.user == request.userId)
        rsp.total = address.count()
        for add in address:
            add_rsp = address_pb2.AddressResponse()

            add_rsp.id = add.id
            add_rsp.userId = add.user
            add_rsp.province = add.province
            add_rsp.city = add.city
            add_rsp.district = add.district
            add_rsp.address = add.address
            add_rsp.singerName = add.singer_name
            add_rsp.singerMobile = add.singer_mobile
            rsp.data.append(add_rsp)

        return rsp

    @logger.catch
    def CreateAddress(self, request: address_pb2.AddressRequest, context):
        addr = Address()
        addr.user = request.userId
        addr.province = request.province
        addr.city = request.city
        addr.district = request.district
        addr.address = request.address
        addr.singer_name = request.singerName
        addr.singer_mobile = request.singerName
        addr.save()

        rsp = address_pb2.AddressResponse(id=addr.id)
        return rsp

    @logger.catch
    def DeleteAddress(self, request: address_pb2.AddressRequest, context):
        try:
            address = Address.get(request.id)
            address.delete_instance()
            return empty_pb2.Empty()
        except DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details('记录不存在')
            return empty_pb2.Empty()

    @logger.catch
    def UpdateAddress(self, request, context):
        try:
            address = Address.get(request.id)
            if request.province:
                address.province = request.province
            if request.city:
                address.city = request.city
            if request.district:
                address.district = request.district
            if request.singerName:
                address.singer_name = request.singerName
            if request.singerMobile:
                address.singer_mobile = request.singerMobile
            address.save()
            return empty_pb2.Empty()
        except DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details('记录不存在')
            return empty_pb2.Empty()


