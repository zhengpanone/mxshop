from nacos import NacosClient

from common.register import base


class NacosRegister(base.Register):

    def __init__(self, client: NacosClient):
        self.client = client

    def register(self, name, id, address, port, tags, check):
        # 服务注册：注册到Nacos，实现服务发现的自动化。heartbeat_interval可以调整后台心跳间隔时间，默认为5秒。
        self.client.add_naming_instance(service_name=name, address=address, port=port,heartbeat_interval=5)

    def deregister(self, service_id):
        pass

    def get_all_services(self):
        pass

    def filter_service(self, filter):
        pass
