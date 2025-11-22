import grpc
import consul
from common.proto.pb import order_pb2, order_pb2_grpc
from order_srv.settings import settings
from loguru import logger


class OrderTest:
    def __init__(self):
        c = consul.Consul(host=settings.CONSUL_HOST, port=settings.CONSUL_PORT)
        services = c.agent.services()
        ip = ""
        port = ""
        for key, value in services.items():
            if value["Service"] == settings.SERVICE_NAME:
                ip = value["Address"]
                port = value["Port"]
                break

        if not ip:
            raise Exception("Order Service not found")
        channel = grpc.insecure_channel(f"{ip}:{port}")
        self.stub = order_pb2_grpc.OrderStub(channel)

    def create_cart_item(self):
        rsp = self.stub.CreateCartItem(order_pb2.CartItemRequest(userId=1, goodsId=1, nums=2))
        print(rsp)

    def create_order(self):
        rsp = self.stub.CreateOrder(order_pb2.OrderRequest(userId=1, address="北京市",
                                                           mobile="12345678901", post="请尽快发货"))
        print(rsp)

    def order_list(self):
        rsp = self.stub.OrderList(order_pb2.OrderFilterRequest(userId=1))
        logger.info(rsp)


if __name__ == '__main__':
    order = OrderTest()
    # order.create_order()
    order.order_list()
