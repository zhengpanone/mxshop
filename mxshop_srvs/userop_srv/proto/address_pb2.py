# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: address.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\raddress.proto\x1a\x1bgoogle/protobuf/empty.proto\"\x99\x01\n\x0e\x41\x64\x64ressRequest\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0e\n\x06userId\x18\x02 \x01(\x05\x12\x10\n\x08province\x18\x03 \x01(\t\x12\x0c\n\x04\x63ity\x18\x04 \x01(\t\x12\x10\n\x08\x64istrict\x18\x05 \x01(\t\x12\x0f\n\x07\x61\x64\x64ress\x18\x06 \x01(\t\x12\x12\n\nsingerName\x18\x07 \x01(\t\x12\x14\n\x0csingerMobile\x18\x08 \x01(\t\"\x9a\x01\n\x0f\x41\x64\x64ressResponse\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0e\n\x06userId\x18\x02 \x01(\x05\x12\x10\n\x08province\x18\x03 \x01(\t\x12\x0c\n\x04\x63ity\x18\x04 \x01(\t\x12\x10\n\x08\x64istrict\x18\x05 \x01(\t\x12\x0f\n\x07\x61\x64\x64ress\x18\x06 \x01(\t\x12\x12\n\nsingerName\x18\x07 \x01(\t\x12\x14\n\x0csingerMobile\x18\x08 \x01(\t\"D\n\x13\x41\x64\x64ressListResponse\x12\r\n\x05total\x18\x01 \x01(\x05\x12\x1e\n\x04\x64\x61ta\x18\x02 \x03(\x0b\x32\x10.AddressResponse2\xea\x01\n\x07\x41\x64\x64ress\x12\x37\n\x0eGetAddressList\x12\x0f.AddressRequest\x1a\x14.AddressListResponse\x12\x32\n\rCreateAddress\x12\x0f.AddressRequest\x1a\x10.AddressResponse\x12\x38\n\rDeleteAddress\x12\x0f.AddressRequest\x1a\x16.google.protobuf.Empty\x12\x38\n\rUpdateAddress\x12\x0f.AddressRequest\x1a\x16.google.protobuf.EmptyB\tZ\x07.;protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'address_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\007.;proto'
  _globals['_ADDRESSREQUEST']._serialized_start=47
  _globals['_ADDRESSREQUEST']._serialized_end=200
  _globals['_ADDRESSRESPONSE']._serialized_start=203
  _globals['_ADDRESSRESPONSE']._serialized_end=357
  _globals['_ADDRESSLISTRESPONSE']._serialized_start=359
  _globals['_ADDRESSLISTRESPONSE']._serialized_end=427
  _globals['_ADDRESS']._serialized_start=430
  _globals['_ADDRESS']._serialized_end=664
# @@protoc_insertion_point(module_scope)