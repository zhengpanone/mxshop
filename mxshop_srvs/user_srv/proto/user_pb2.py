# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nuser.proto\x1a\x1bgoogle/protobuf/empty.proto\"@\n\x11PasswordCheckInfo\x12\x10\n\x08password\x18\x01 \x01(\t\x12\x19\n\x11\x65ncryptedPassword\x18\x02 \x01(\t\"(\n\x15\x43heckPasswordResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\"&\n\x08PageInfo\x12\x0c\n\x04page\x18\x01 \x01(\r\x12\x0c\n\x04size\x18\x02 \x01(\r\"\x1f\n\rMobileRequest\x12\x0e\n\x06mobile\x18\x01 \x01(\t\"\x17\n\tIdRequest\x12\n\n\x02id\x18\x01 \x01(\x05\"D\n\x0e\x43reateUserInfo\x12\x10\n\x08nickname\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\x12\x0e\n\x06mobile\x18\x03 \x01(\t\"P\n\x0eUpdateUserInfo\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x10\n\x08nickname\x18\x02 \x01(\t\x12\x0e\n\x06gender\x18\x03 \x01(\t\x12\x10\n\x08\x62irthday\x18\x04 \x01(\x04\"B\n\x10UserListResponse\x12\r\n\x05total\x18\x01 \x01(\x05\x12\x1f\n\x04\x64\x61ta\x18\x02 \x03(\x0b\x32\x11.UserInfoResponse\"\x82\x01\n\x10UserInfoResponse\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x10\n\x08nickname\x18\x02 \x01(\t\x12\x10\n\x08password\x18\x03 \x01(\t\x12\x0e\n\x06mobile\x18\x04 \x01(\t\x12\x10\n\x08\x62irthday\x18\x05 \x01(\x04\x12\x0e\n\x06gender\x18\x06 \x01(\t\x12\x0c\n\x04role\x18\x07 \x01(\t2\xbd\x02\n\x04User\x12+\n\x0bGetUserList\x12\t.PageInfo\x1a\x11.UserListResponse\x12\x34\n\x0fGetUserByMobile\x12\x0e.MobileRequest\x1a\x11.UserInfoResponse\x12,\n\x0bGetUserById\x12\n.IdRequest\x1a\x11.UserInfoResponse\x12\x30\n\nCreateUser\x12\x0f.CreateUserInfo\x1a\x11.UserInfoResponse\x12\x35\n\nUpdateUser\x12\x0f.UpdateUserInfo\x1a\x16.google.protobuf.Empty\x12;\n\rCheckPassword\x12\x12.PasswordCheckInfo\x1a\x16.CheckPasswordResponseB\tZ\x07.;protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'user_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\007.;proto'
  _globals['_PASSWORDCHECKINFO']._serialized_start=43
  _globals['_PASSWORDCHECKINFO']._serialized_end=107
  _globals['_CHECKPASSWORDRESPONSE']._serialized_start=109
  _globals['_CHECKPASSWORDRESPONSE']._serialized_end=149
  _globals['_PAGEINFO']._serialized_start=151
  _globals['_PAGEINFO']._serialized_end=189
  _globals['_MOBILEREQUEST']._serialized_start=191
  _globals['_MOBILEREQUEST']._serialized_end=222
  _globals['_IDREQUEST']._serialized_start=224
  _globals['_IDREQUEST']._serialized_end=247
  _globals['_CREATEUSERINFO']._serialized_start=249
  _globals['_CREATEUSERINFO']._serialized_end=317
  _globals['_UPDATEUSERINFO']._serialized_start=319
  _globals['_UPDATEUSERINFO']._serialized_end=399
  _globals['_USERLISTRESPONSE']._serialized_start=401
  _globals['_USERLISTRESPONSE']._serialized_end=467
  _globals['_USERINFORESPONSE']._serialized_start=470
  _globals['_USERINFORESPONSE']._serialized_end=600
  _globals['_USER']._serialized_start=603
  _globals['_USER']._serialized_end=920
# @@protoc_insertion_point(module_scope)
