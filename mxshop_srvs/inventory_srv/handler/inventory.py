import grpc
from google.protobuf import empty_pb2
from loguru import logger
from peewee import DoesNotExist
import json
from inventory_srv.model.models import Inventory
from inventory_srv.proto import inventory_pb2, inventory_pb2_grpc


class InventoryServicer(inventory_pb2_grpc.InventoryServicer):

    pass
