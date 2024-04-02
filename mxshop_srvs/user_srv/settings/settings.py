from playhouse.pool import PooledMySQLDatabase
from playhouse.shortcuts import ReconnectMixin


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


# 数据库配置
MYSQL_DB = "mxshop_user_srv"
MYSQL_HOST = "127.0.0.1"
MYSQL_PORT = 3306
MYSQL_USER = "root"
MYSQL_PASSWORD = "root"
DB = ReconnectMysqlDatabase(MYSQL_DB, host=MYSQL_HOST, port=MYSQL_PORT, user=MYSQL_USER, password=MYSQL_PASSWORD)

# consul配置
CONSUL_HOST = "127.0.0.1"
CONSUL_PORT = 8500

# 服务相关的配置
SERVICE_NAME = "user-srv"
SERVICE_HOST = get_server_ip()
SERVICE_TAGS = ['imooc', 'zhengpanone', 'python', 'srv']
