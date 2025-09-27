import grpc
import consul
from common.proto.pb import goods_pb2_grpc, goods_pb2
from goods_srv.settings import settings


class GoodsTest:
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
        self.stub = goods_pb2_grpc.GoodsStub(channel)

    def create_goods(self):
        rsp = self.stub.CreateGoods(goods_pb2.CreateGoodsInfo(name="香蕉", goods_sn="xxxxx",market_price=12.11))
    def goods_list(self):
        rsp: goods_pb2.GoodsListResponse = self.stub.GoodsList(goods_pb2.GoodsFilterRequest(price_min=50))
        print(rsp.total)
        for item in rsp.data:
            print(item.name)

    def batch_get(self):
        ids = [1, 2, 3]
        rsp: goods_pb2.GoodsListResponse = self.stub.BatchGetGoods(
            goods_pb2.BatchGoodsIdInfo(id=ids)
        )
        for item in rsp.data:
            print(item.name)

    def delete_goods(self):
        rsp = self.stub.DeleteGoods(goods_pb2.DeleteGoodsInfo(id=1))
        print(rsp)

    def get_goods_detail(self, id):
        rsp = self.stub.GetGoodsDetail(goods_pb2.GoodInfoRequest(id=id))
        print(rsp)


if __name__ == '__main__':
    goods = GoodsTest()
    goods.create_goods()
    # goods.goods_list()
    # goods.batch_get()
    # goods.delete_goods()
    # goods.get_goods_detail(1)
