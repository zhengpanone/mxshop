import logging
import os
import signal
import sys
from loguru import logger
import argparse

BASE_DIR = os.path.dirname(os.path.abspath(os.path.dirname(__file__)))
sys.path.insert(0, BASE_DIR)

import grpc
from concurrent import futures
from user_srv.proto import user_pb2_grpc
from user_srv.handler.user import UserServicer


def on_exit(signum, frame):
    logger.info("程序进程中断")
    sys.exit(0)


def server():
    parser = argparse.ArgumentParser()
    parser.add_argument('--port', nargs="?", type=int, default=5001, help='server port')
    parser.add_argument('--host', nargs="?", type=str, default='0.0.0.0', help='server host')
    args = parser.parse_args()

    logger.add("logs/user_srv_{time}.log")
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    user_pb2_grpc.add_UserServicer_to_server(UserServicer(), server)
    # server.add_insecure_port('[::]:50051')
    server.add_insecure_port(f'{args.host}:{args.port}')
    # 主进程退出信号监听
    """
        windows下支撑的信号是有限的
        SIGINT ctrl+c 终端
        SIGTERM kill 发出的软件终止
        
    """
    signal.signal(signal.SIGINT, on_exit)
    signal.signal(signal.SIGTERM, on_exit)
    logger.info(f'Starting server http://{args.host}:{args.port}')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    server()
