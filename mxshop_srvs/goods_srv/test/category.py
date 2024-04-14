import grpc
import consul
from google.protobuf import empty_pb2
from goods_srv.proto import category_pb2_grpc, category_pb2
from goods_srv.settings import settings


class CategoryTest:
    def __init__(self):
        c = consul.Consul(host="127.0.0.1", port=8500)
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
        self.stub = category_pb2_grpc.CategoryStub(channel)

    def category_all_list(self):
        rsp: category_pb2.CategoryListResponse = self.stub.GetAllCategoryList(empty_pb2.Empty())
        print(rsp.total)
        for item in rsp.data:
            print(item.name)


if __name__ == '__main__':
    category = CategoryTest()
    category.category_all_list()
