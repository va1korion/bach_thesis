# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: worker.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0cworker.proto\x12\x07manager\"\x19\n\nMsgRequest\x12\x0b\n\x03img\x18\x01 \x01(\x0c\"\x19\n\x08MsgReply\x12\r\n\x05reply\x18\x01 \x01(\x0c\x32\x44\n\tImageTest\x12\x37\n\x07\x41nalyse\x12\x13.manager.MsgRequest\x1a\x11.manager.MsgReply\"\x00(\x01\x30\x01\x42\x15P\x01Z\x0bworker_api/\xa2\x02\x03HLWb\x06proto3')



_MSGREQUEST = DESCRIPTOR.message_types_by_name['MsgRequest']
_MSGREPLY = DESCRIPTOR.message_types_by_name['MsgReply']
MsgRequest = _reflection.GeneratedProtocolMessageType('MsgRequest', (_message.Message,), {
  'DESCRIPTOR' : _MSGREQUEST,
  '__module__' : 'worker_pb2'
  # @@protoc_insertion_point(class_scope:manager.MsgRequest)
  })
_sym_db.RegisterMessage(MsgRequest)

MsgReply = _reflection.GeneratedProtocolMessageType('MsgReply', (_message.Message,), {
  'DESCRIPTOR' : _MSGREPLY,
  '__module__' : 'worker_pb2'
  # @@protoc_insertion_point(class_scope:manager.MsgReply)
  })
_sym_db.RegisterMessage(MsgReply)

_IMAGETEST = DESCRIPTOR.services_by_name['ImageTest']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'P\001Z\013worker_api/\242\002\003HLW'
  _MSGREQUEST._serialized_start=25
  _MSGREQUEST._serialized_end=50
  _MSGREPLY._serialized_start=52
  _MSGREPLY._serialized_end=77
  _IMAGETEST._serialized_start=79
  _IMAGETEST._serialized_end=147
# @@protoc_insertion_point(module_scope)