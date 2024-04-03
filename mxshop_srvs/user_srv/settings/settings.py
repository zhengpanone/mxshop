from playhouse.pool import PooledMySQLDatabase
from playhouse.shortcuts import ReconnectMixin
import nacos
import json
from loguru import logger


class ReconnectMysqlDatabase(ReconnectMixin, PooledMySQLDatabase):
    # python的mro
    pass


def get_server_ip():
    """
    获取ip
    :return: ip地址
    """
    import socket
    hostname = socket.gethostname()
    ip_address = socket.gethostbyname(hostname)
    return ip_address


NACOS = {
    "host": "127.0.0.1",
    "port": 8848,
    "namespace": "ac2997b0-1569-47d9-a792-efc10375341b",
    "dataId": "user-srv.json",
    "groupId": "dev",
    "user": "nacos",
    "password": "nacos"
}

client = nacos.NacosClient(f"{NACOS['host']}:{NACOS['port']}", namespace=NACOS['namespace'], username=NACOS["user"],
                           password=NACOS["password"])
data = json.loads(client.get_config(NACOS["dataId"], NACOS["groupId"]))
logger.info(data)


def update_config(args):
    logger.info("配置产生变化")
    print(json.loads(data))


# 数据库配置
MYSQL_DB = data['mysql']['db']
MYSQL_HOST = data['mysql']['host']
MYSQL_PORT = data['mysql']['port']
MYSQL_USER = data['mysql']['user']
MYSQL_PASSWORD = data['mysql']['password']
DB = ReconnectMysqlDatabase(MYSQL_DB, host=MYSQL_HOST, port=MYSQL_PORT, user=MYSQL_USER, password=MYSQL_PASSWORD)

# consul配置
CONSUL_HOST = data["consul"]["host"]
CONSUL_PORT = data["consul"]["port"]

# 服务相关的配置
SERVICE_NAME = data["name"]
SERVICE_HOST = get_server_ip()
SERVICE_TAGS = data["tags"]
