import logging
import os
import signal
import socket
import sys
from loguru import logger
import argparse

BASE_DIR = os.path.dirname(os.path.abspath(os.path.dirname(__file__)))
sys.path.insert(0, BASE_DIR)

import grpc
from concurrent import futures
from userop_srv.proto import userfav_pb2_grpc, message_pb2_grpc, address_pb2_grpc
from userop_srv.handler.message import MessageServicer
from userop_srv.handler.address import AddressServicer
from userop_srv.handler.user_fav import UserFavServicer
from common.register import consul
from common.grpc_health.v1 import health_pb2_grpc, health
from userop_srv.settings import settings
from functools import partial


def on_exit(signum, frame, service_id):
    register = consul.ConsulRegister(settings.CONSUL_HOST, settings.CONSUL_PORT)
    logger.info(f"注销{service_id}服务")
    result = register.deregister(service_id)
    if result:
        logger.info(f"注销用户服务：{service_id} 成功")
    else:
        logger.error(f"注销用户服务：{service_id} 失败")

    sys.exit(0)


def get_free_tcp_port():
    tcp = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    tcp.bind(("", 0))
    addr, port = tcp.getsockname()
    tcp.close()

    return port


def test_db(args):
    print("配置文件产生变化")
    print(args)


def server():
    parser = argparse.ArgumentParser()
    parser.add_argument('--host', nargs="?", type=str, default='0.0.0.0', help='server host')
    parser.add_argument('--port', nargs="?", type=int, default=0, help='server port')
    args = parser.parse_args()
    if args.port == 0:
        port = get_free_tcp_port()
    else:
        port = args.port

    logger.add("logs/userop_srv_{time}.log")
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

    # 注册用户收藏服务
    userfav_pb2_grpc.add_UserFavServicer_to_server(UserFavServicer(), server)
    # 收件地址服务
    address_pb2_grpc.add_AddressServicer_to_server(AddressServicer(), server)
    # 消息服务
    message_pb2_grpc.add_MessageServicer_to_server(MessageServicer(), server)

    # 注册健康检查
    health_pb2_grpc.add_HealthServicer_to_server(health.HealthServicer(), server)

    import uuid
    service_id = str(uuid.uuid1())
    # server.add_insecure_port('[::]:50051')
    server.add_insecure_port(f'{args.host}:{port}')
    # 主进程退出信号监听
    """
        windows下支撑的信号是有限的
        SIGINT ctrl+c 终端
        SIGTERM kill 发出的软件终止
        
    """
    signal.signal(signal.SIGINT, partial(on_exit, service_id=service_id))
    signal.signal(signal.SIGTERM, partial(on_exit, service_id=service_id))
    logger.info(f'Starting server http://{args.host}:{port}')
    server.start()

    logger.info(f"用户操作服务注册到注册中心")
    register = consul.ConsulRegister(settings.CONSUL_HOST, settings.CONSUL_PORT)
    if not register.register(settings.SERVICE_NAME, service_id, settings.SERVICE_HOST, port,
                             settings.SERVICE_TAGS, None):
        logger.info(f"用户操作服务注册失败")
        sys.exit(0)

    server.wait_for_termination()


if __name__ == '__main__':
    # print(get_free_tcp_port())
    logging.basicConfig()
    settings.client.add_config_watcher(settings.NACOS["dataId"], settings.NACOS["groupId"],
                                       settings.update_config)  # 这个逻辑必须放在__name__ == '__main__'中
    server()
