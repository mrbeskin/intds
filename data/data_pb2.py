# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: data/data.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='data/data.proto',
  package='dataArray',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x0f\x64\x61ta/data.proto\x12\tdataArray\"\x1e\n\x1cGetIntDataArrayStreamRequest\"\x14\n\x12ServerSendResponse\"\x1c\n\x0cIntDataArray\x12\x0c\n\x04\x64\x61ta\x18\x01 \x03(\x05\x32\xb5\x01\n\nDataServer\x12M\n\x11WriteIntDataArray\x12\x17.dataArray.IntDataArray\x1a\x1d.dataArray.ServerSendResponse\"\x00\x12X\n\x10GetIntDataArrays\x12\'.dataArray.GetIntDataArrayStreamRequest\x1a\x17.dataArray.IntDataArray\"\x00\x30\x01\x32U\n\nDataClient\x12G\n\x11WriteIntDataArray\x12\x17.dataArray.IntDataArray\x1a\x17.dataArray.IntDataArray\"\x00\x62\x06proto3')
)




_GETINTDATAARRAYSTREAMREQUEST = _descriptor.Descriptor(
  name='GetIntDataArrayStreamRequest',
  full_name='dataArray.GetIntDataArrayStreamRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=30,
  serialized_end=60,
)


_SERVERSENDRESPONSE = _descriptor.Descriptor(
  name='ServerSendResponse',
  full_name='dataArray.ServerSendResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=62,
  serialized_end=82,
)


_INTDATAARRAY = _descriptor.Descriptor(
  name='IntDataArray',
  full_name='dataArray.IntDataArray',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='data', full_name='dataArray.IntDataArray.data', index=0,
      number=1, type=5, cpp_type=1, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=84,
  serialized_end=112,
)

DESCRIPTOR.message_types_by_name['GetIntDataArrayStreamRequest'] = _GETINTDATAARRAYSTREAMREQUEST
DESCRIPTOR.message_types_by_name['ServerSendResponse'] = _SERVERSENDRESPONSE
DESCRIPTOR.message_types_by_name['IntDataArray'] = _INTDATAARRAY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

GetIntDataArrayStreamRequest = _reflection.GeneratedProtocolMessageType('GetIntDataArrayStreamRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETINTDATAARRAYSTREAMREQUEST,
  __module__ = 'data.data_pb2'
  # @@protoc_insertion_point(class_scope:dataArray.GetIntDataArrayStreamRequest)
  ))
_sym_db.RegisterMessage(GetIntDataArrayStreamRequest)

ServerSendResponse = _reflection.GeneratedProtocolMessageType('ServerSendResponse', (_message.Message,), dict(
  DESCRIPTOR = _SERVERSENDRESPONSE,
  __module__ = 'data.data_pb2'
  # @@protoc_insertion_point(class_scope:dataArray.ServerSendResponse)
  ))
_sym_db.RegisterMessage(ServerSendResponse)

IntDataArray = _reflection.GeneratedProtocolMessageType('IntDataArray', (_message.Message,), dict(
  DESCRIPTOR = _INTDATAARRAY,
  __module__ = 'data.data_pb2'
  # @@protoc_insertion_point(class_scope:dataArray.IntDataArray)
  ))
_sym_db.RegisterMessage(IntDataArray)



_DATASERVER = _descriptor.ServiceDescriptor(
  name='DataServer',
  full_name='dataArray.DataServer',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=115,
  serialized_end=296,
  methods=[
  _descriptor.MethodDescriptor(
    name='WriteIntDataArray',
    full_name='dataArray.DataServer.WriteIntDataArray',
    index=0,
    containing_service=None,
    input_type=_INTDATAARRAY,
    output_type=_SERVERSENDRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetIntDataArrays',
    full_name='dataArray.DataServer.GetIntDataArrays',
    index=1,
    containing_service=None,
    input_type=_GETINTDATAARRAYSTREAMREQUEST,
    output_type=_INTDATAARRAY,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_DATASERVER)

DESCRIPTOR.services_by_name['DataServer'] = _DATASERVER


_DATACLIENT = _descriptor.ServiceDescriptor(
  name='DataClient',
  full_name='dataArray.DataClient',
  file=DESCRIPTOR,
  index=1,
  serialized_options=None,
  serialized_start=298,
  serialized_end=383,
  methods=[
  _descriptor.MethodDescriptor(
    name='WriteIntDataArray',
    full_name='dataArray.DataClient.WriteIntDataArray',
    index=0,
    containing_service=None,
    input_type=_INTDATAARRAY,
    output_type=_INTDATAARRAY,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_DATACLIENT)

DESCRIPTOR.services_by_name['DataClient'] = _DATACLIENT

# @@protoc_insertion_point(module_scope)