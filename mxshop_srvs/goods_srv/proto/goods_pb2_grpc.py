# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from goods_srv.proto import goods_pb2 as goods__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


class GoodsStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GoodsList = channel.unary_unary(
                '/Goods/GoodsList',
                request_serializer=goods__pb2.GoodsFilterRequest.SerializeToString,
                response_deserializer=goods__pb2.GoodsListResponse.FromString,
                )
        self.BatchGetGoods = channel.unary_unary(
                '/Goods/BatchGetGoods',
                request_serializer=goods__pb2.BatchGoodsIdInfo.SerializeToString,
                response_deserializer=goods__pb2.GoodsListResponse.FromString,
                )
        self.CreateGoods = channel.unary_unary(
                '/Goods/CreateGoods',
                request_serializer=goods__pb2.CreateGoodsInfo.SerializeToString,
                response_deserializer=goods__pb2.GoodsInfoResponse.FromString,
                )
        self.DeleteGoods = channel.unary_unary(
                '/Goods/DeleteGoods',
                request_serializer=goods__pb2.DeleteGoodsInfo.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.UpdateGoods = channel.unary_unary(
                '/Goods/UpdateGoods',
                request_serializer=goods__pb2.CreateGoodsInfo.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.GetGoodsDetail = channel.unary_unary(
                '/Goods/GetGoodsDetail',
                request_serializer=goods__pb2.GoodInfoRequest.SerializeToString,
                response_deserializer=goods__pb2.GoodsInfoResponse.FromString,
                )


class GoodsServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GoodsList(self, request, context):
        """商品接口
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BatchGetGoods(self, request, context):
        """用户提交订单有多个商品，批量查询商品信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateGoods(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteGoods(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateGoods(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetGoodsDetail(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_GoodsServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GoodsList': grpc.unary_unary_rpc_method_handler(
                    servicer.GoodsList,
                    request_deserializer=goods__pb2.GoodsFilterRequest.FromString,
                    response_serializer=goods__pb2.GoodsListResponse.SerializeToString,
            ),
            'BatchGetGoods': grpc.unary_unary_rpc_method_handler(
                    servicer.BatchGetGoods,
                    request_deserializer=goods__pb2.BatchGoodsIdInfo.FromString,
                    response_serializer=goods__pb2.GoodsListResponse.SerializeToString,
            ),
            'CreateGoods': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateGoods,
                    request_deserializer=goods__pb2.CreateGoodsInfo.FromString,
                    response_serializer=goods__pb2.GoodsInfoResponse.SerializeToString,
            ),
            'DeleteGoods': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteGoods,
                    request_deserializer=goods__pb2.DeleteGoodsInfo.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'UpdateGoods': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateGoods,
                    request_deserializer=goods__pb2.CreateGoodsInfo.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'GetGoodsDetail': grpc.unary_unary_rpc_method_handler(
                    servicer.GetGoodsDetail,
                    request_deserializer=goods__pb2.GoodInfoRequest.FromString,
                    response_serializer=goods__pb2.GoodsInfoResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'Goods', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Goods(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GoodsList(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Goods/GoodsList',
            goods__pb2.GoodsFilterRequest.SerializeToString,
            goods__pb2.GoodsListResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def BatchGetGoods(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Goods/BatchGetGoods',
            goods__pb2.BatchGoodsIdInfo.SerializeToString,
            goods__pb2.GoodsListResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateGoods(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Goods/CreateGoods',
            goods__pb2.CreateGoodsInfo.SerializeToString,
            goods__pb2.GoodsInfoResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteGoods(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Goods/DeleteGoods',
            goods__pb2.DeleteGoodsInfo.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateGoods(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Goods/UpdateGoods',
            goods__pb2.CreateGoodsInfo.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetGoodsDetail(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Goods/GetGoodsDetail',
            goods__pb2.GoodInfoRequest.SerializeToString,
            goods__pb2.GoodsInfoResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
