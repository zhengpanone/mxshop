import random

import requests
from common.register import base
import consul


class ConsulRegister(base.Register):

    def __init__(self, host, port):
        self.host = host
        self.port = port
        self.consul = consul.Consul(host=host, port=port)

    def register(self, name: str, id: str, address: str, port: int, tags, check) -> bool:
        if check is None:
            check = {
                "GRPC": f"{address}:{port}",
                "GRPCUseTLS": False,
                "Timeout": "5s",
                "Interval": "5s",
                "DeregisterCriticalServiceAfter": "15s"
            }
        else:
            check = check

        success = self.consul.agent.service.register(name=name, service_id=id, address=address,
                                                     port=port, tags=tags, check=check)
        if success:
            return True
        else:
            return False

    def deregister(self, service_id) -> bool:
        return self.consul.agent.service.deregister(service_id =service_id)

    def get_all_services(self):
        return self.consul.agent.service()

    def filter_service(self, filter):
        url = f"http://{self.host}:{self.port}/v1/agent/services"
        params = {
            "filter": filter
        }
        return requests.get(url, params=params).json()

    def get_host_port(self, filter):
        data = self.filter_service(filter)
        if data:
            service_info = random.choice(list(data.values()))
            return service_info["Address"], service_info["Port"]

        return None, None
