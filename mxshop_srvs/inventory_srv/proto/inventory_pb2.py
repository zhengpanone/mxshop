# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: inventory.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0finventory.proto\x1a\x1bgoogle/protobuf/empty.proto\"-\n\x0cGoodsInvInfo\x12\x10\n\x08goods_id\x18\x01 \x01(\x05\x12\x0b\n\x03num\x18\x02 \x01(\x05\"-\n\x08SellInfo\x12!\n\ngoods_info\x18\x01 \x03(\x0b\x32\r.GoodsInvInfo2\xbf\x01\n\tInventory\x12/\n\x06SetInv\x12\r.GoodsInvInfo\x1a\x16.google.protobuf.Empty\x12)\n\tInvDetail\x12\r.GoodsInvInfo\x1a\r.GoodsInvInfo\x12)\n\x04Sell\x12\t.SellInfo\x1a\x16.google.protobuf.Empty\x12+\n\x06Reback\x12\t.SellInfo\x1a\x16.google.protobuf.EmptyB\tZ\x07.;protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'inventory_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\007.;proto'
  _globals['_GOODSINVINFO']._serialized_start=48
  _globals['_GOODSINVINFO']._serialized_end=93
  _globals['_SELLINFO']._serialized_start=95
  _globals['_SELLINFO']._serialized_end=140
  _globals['_INVENTORY']._serialized_start=143
  _globals['_INVENTORY']._serialized_end=334
# @@protoc_insertion_point(module_scope)
