# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: order.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0border.proto\x1a\x1bgoogle/protobuf/empty.proto\"\x16\n\x08UserInfo\x12\n\n\x02id\x18\x01 \x01(\x05\"b\n\x14ShopCartInfoResponse\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0e\n\x06userId\x18\x02 \x01(\x05\x12\x0f\n\x07goodsId\x18\x03 \x01(\x05\x12\x0c\n\x04nums\x18\x04 \x01(\x05\x12\x0f\n\x07\x63hecked\x18\x05 \x01(\x08\"J\n\x14\x43\x61rtItemListResponse\x12\r\n\x05total\x18\x01 \x01(\x05\x12#\n\x04\x64\x61ta\x18\x02 \x03(\x0b\x32\x15.ShopCartInfoResponse\"Q\n\x0f\x43\x61rtItemRequest\x12\x0e\n\x06userId\x18\x01 \x01(\x05\x12\x0f\n\x07goodsId\x18\x02 \x01(\x05\x12\x0c\n\x04nums\x18\x03 \x01(\x05\x12\x0f\n\x07\x63hecked\x18\x04 \x01(\x08\"g\n\x0cOrderRequest\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0e\n\x06userId\x18\x02 \x01(\x05\x12\x0f\n\x07\x61\x64\x64ress\x18\x03 \x01(\t\x12\x0e\n\x06mobile\x18\x04 \x01(\t\x12\x0c\n\x04name\x18\x05 \x01(\t\x12\x0c\n\x04post\x18\x06 \x01(\t\"\xad\x01\n\x11OrderInfoResponse\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0e\n\x06userId\x18\x02 \x01(\x05\x12\x0f\n\x07orderSn\x18\x03 \x01(\t\x12\x0f\n\x07payType\x18\x04 \x01(\t\x12\x0e\n\x06status\x18\x05 \x01(\t\x12\x0c\n\x04post\x18\x06 \x01(\t\x12\r\n\x05total\x18\x07 \x01(\x02\x12\x0f\n\x07\x61\x64\x64ress\x18\x08 \x01(\t\x12\x0c\n\x04name\x18\t \x01(\t\x12\x0e\n\x06mobile\x18\n \x01(\t\"D\n\x11OrderListResponse\x12\r\n\x05total\x18\x01 \x01(\x05\x12 \n\x04\x64\x61ta\x18\x02 \x03(\x0b\x32\x12.OrderInfoResponse\"@\n\x12OrderFilterRequest\x12\x0e\n\x06userId\x18\x01 \x01(\x05\x12\x0c\n\x04page\x18\x02 \x01(\x05\x12\x0c\n\x04size\x18\x03 \x01(\x05\"\x8a\x01\n\x11OrderItemResponse\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0f\n\x07orderId\x18\x02 \x01(\x05\x12\x0f\n\x07goodsId\x18\x03 \x01(\x05\x12\x11\n\tgoodsName\x18\x04 \x01(\t\x12\x12\n\ngoodsImage\x18\x05 \x01(\t\x12\x12\n\ngoodsPrice\x18\x06 \x01(\t\x12\x0c\n\x04nums\x18\x07 \x01(\x05\"c\n\x17OrderInfoDetailResponse\x12%\n\torderInfo\x18\x01 \x01(\x0b\x32\x12.OrderInfoResponse\x12!\n\x05goods\x18\x02 \x03(\x0b\x32\x12.OrderItemResponse\".\n\x0bOrderStatus\x12\x0f\n\x07orderSn\x18\x01 \x01(\t\x12\x0e\n\x06status\x18\x02 \x01(\t2\xc6\x03\n\x05Order\x12/\n\x0b\x43\x61rItemList\x12\t.UserInfo\x1a\x15.CartItemListResponse\x12\x39\n\x0e\x43reateCartItem\x12\x10.CartItemRequest\x1a\x15.ShopCartInfoResponse\x12:\n\x0eUpdateCartItem\x12\x10.CartItemRequest\x1a\x16.google.protobuf.Empty\x12:\n\x0e\x44\x65leteCartItem\x12\x10.CartItemRequest\x1a\x16.google.protobuf.Empty\x12\x30\n\x0b\x43reateOrder\x12\r.OrderRequest\x1a\x12.OrderInfoResponse\x12\x34\n\tOrderList\x12\x13.OrderFilterRequest\x1a\x12.OrderListResponse\x12\x36\n\x0bOrderDetail\x12\r.OrderRequest\x1a\x18.OrderInfoDetailResponse\x12\x39\n\x11UpdateOrderStatus\x12\x0c.OrderStatus\x1a\x16.google.protobuf.EmptyB\tZ\x07.;protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'order_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\007.;proto'
  _globals['_USERINFO']._serialized_start=44
  _globals['_USERINFO']._serialized_end=66
  _globals['_SHOPCARTINFORESPONSE']._serialized_start=68
  _globals['_SHOPCARTINFORESPONSE']._serialized_end=166
  _globals['_CARTITEMLISTRESPONSE']._serialized_start=168
  _globals['_CARTITEMLISTRESPONSE']._serialized_end=242
  _globals['_CARTITEMREQUEST']._serialized_start=244
  _globals['_CARTITEMREQUEST']._serialized_end=325
  _globals['_ORDERREQUEST']._serialized_start=327
  _globals['_ORDERREQUEST']._serialized_end=430
  _globals['_ORDERINFORESPONSE']._serialized_start=433
  _globals['_ORDERINFORESPONSE']._serialized_end=606
  _globals['_ORDERLISTRESPONSE']._serialized_start=608
  _globals['_ORDERLISTRESPONSE']._serialized_end=676
  _globals['_ORDERFILTERREQUEST']._serialized_start=678
  _globals['_ORDERFILTERREQUEST']._serialized_end=742
  _globals['_ORDERITEMRESPONSE']._serialized_start=745
  _globals['_ORDERITEMRESPONSE']._serialized_end=883
  _globals['_ORDERINFODETAILRESPONSE']._serialized_start=885
  _globals['_ORDERINFODETAILRESPONSE']._serialized_end=984
  _globals['_ORDERSTATUS']._serialized_start=986
  _globals['_ORDERSTATUS']._serialized_end=1032
  _globals['_ORDER']._serialized_start=1035
  _globals['_ORDER']._serialized_end=1489
# @@protoc_insertion_point(module_scope)
