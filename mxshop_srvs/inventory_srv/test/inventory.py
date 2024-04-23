import grpc
import consul
from inventory_srv.proto import inventory_pb2_grpc, inventory_pb2
from inventory_srv.settings import settings


class InventoryTest:
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
            raise Exception("Service not found")
        channel = grpc.insecure_channel(f"{ip}:{port}")
        self.stub = inventory_pb2_grpc.InventoryStub(channel)

    def set_inv(self):
        rsp = self.stub.SetInv(inventory_pb2.GoodsInvInfo(goods_id=1,num=100))
        print(rsp)

    def get_inv(self):
        rsp = self.stub.InvDetail(inventory_pb2.GoodsInvInfo(goods_id=2))
        print(rsp)

    def sell(self):
        goods_list = [(1,10),(2,11),(3,15),(4,119)]
        request = inventory_pb2.SellInfo()
        for goodsId, num in goods_list:
            request.goods_info.append(inventory_pb2.GoodsInvInfo(goods_id=goodsId, num=num))
        rsp = self.stub.Sell(request)

    def reback(self):
        goods_list = [(1,10),(2,11),(3,15),(4,119)]
        request = inventory_pb2.SellInfo()
        for goodsId, num in goods_list:
            request.goods_info.append(inventory_pb2.GoodsInvInfo(goods_id=goodsId, num=num))
        rsp = self.stub.Reback(request)

if __name__ == '__main__':
    inventory = InventoryTest()

    inventory.set_inv()
    inventory.get_inv()
    inventory.reback()
