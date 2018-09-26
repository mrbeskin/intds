// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

package dataArray

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetIntDataArrayStreamRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIntDataArrayStreamRequest) Reset()         { *m = GetIntDataArrayStreamRequest{} }
func (m *GetIntDataArrayStreamRequest) String() string { return proto.CompactTextString(m) }
func (*GetIntDataArrayStreamRequest) ProtoMessage()    {}
func (*GetIntDataArrayStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

func (m *GetIntDataArrayStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIntDataArrayStreamRequest.Unmarshal(m, b)
}
func (m *GetIntDataArrayStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIntDataArrayStreamRequest.Marshal(b, m, deterministic)
}
func (m *GetIntDataArrayStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIntDataArrayStreamRequest.Merge(m, src)
}
func (m *GetIntDataArrayStreamRequest) XXX_Size() int {
	return xxx_messageInfo_GetIntDataArrayStreamRequest.Size(m)
}
func (m *GetIntDataArrayStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIntDataArrayStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetIntDataArrayStreamRequest proto.InternalMessageInfo

type ServerSendResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerSendResponse) Reset()         { *m = ServerSendResponse{} }
func (m *ServerSendResponse) String() string { return proto.CompactTextString(m) }
func (*ServerSendResponse) ProtoMessage()    {}
func (*ServerSendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{1}
}

func (m *ServerSendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerSendResponse.Unmarshal(m, b)
}
func (m *ServerSendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerSendResponse.Marshal(b, m, deterministic)
}
func (m *ServerSendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerSendResponse.Merge(m, src)
}
func (m *ServerSendResponse) XXX_Size() int {
	return xxx_messageInfo_ServerSendResponse.Size(m)
}
func (m *ServerSendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerSendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerSendResponse proto.InternalMessageInfo

type IntDataArray struct {
	Data                 []int32  `protobuf:"varint,1,rep,packed,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntDataArray) Reset()         { *m = IntDataArray{} }
func (m *IntDataArray) String() string { return proto.CompactTextString(m) }
func (*IntDataArray) ProtoMessage()    {}
func (*IntDataArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{2}
}

func (m *IntDataArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntDataArray.Unmarshal(m, b)
}
func (m *IntDataArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntDataArray.Marshal(b, m, deterministic)
}
func (m *IntDataArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntDataArray.Merge(m, src)
}
func (m *IntDataArray) XXX_Size() int {
	return xxx_messageInfo_IntDataArray.Size(m)
}
func (m *IntDataArray) XXX_DiscardUnknown() {
	xxx_messageInfo_IntDataArray.DiscardUnknown(m)
}

var xxx_messageInfo_IntDataArray proto.InternalMessageInfo

func (m *IntDataArray) GetData() []int32 {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetIntDataArrayStreamRequest)(nil), "dataArray.GetIntDataArrayStreamRequest")
	proto.RegisterType((*ServerSendResponse)(nil), "dataArray.ServerSendResponse")
	proto.RegisterType((*IntDataArray)(nil), "dataArray.IntDataArray")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataServerClient is the client API for DataServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataServerClient interface {
	WriteIntDataArray(ctx context.Context, in *IntDataArray, opts ...grpc.CallOption) (*ServerSendResponse, error)
	// A simple service for reading a stream of IntDataArrays.
	GetIntDataArrays(ctx context.Context, in *GetIntDataArrayStreamRequest, opts ...grpc.CallOption) (DataServer_GetIntDataArraysClient, error)
}

type dataServerClient struct {
	cc *grpc.ClientConn
}

func NewDataServerClient(cc *grpc.ClientConn) DataServerClient {
	return &dataServerClient{cc}
}

func (c *dataServerClient) WriteIntDataArray(ctx context.Context, in *IntDataArray, opts ...grpc.CallOption) (*ServerSendResponse, error) {
	out := new(ServerSendResponse)
	err := c.cc.Invoke(ctx, "/dataArray.DataServer/WriteIntDataArray", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServerClient) GetIntDataArrays(ctx context.Context, in *GetIntDataArrayStreamRequest, opts ...grpc.CallOption) (DataServer_GetIntDataArraysClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataServer_serviceDesc.Streams[0], "/dataArray.DataServer/GetIntDataArrays", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServerGetIntDataArraysClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataServer_GetIntDataArraysClient interface {
	Recv() (*IntDataArray, error)
	grpc.ClientStream
}

type dataServerGetIntDataArraysClient struct {
	grpc.ClientStream
}

func (x *dataServerGetIntDataArraysClient) Recv() (*IntDataArray, error) {
	m := new(IntDataArray)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataServerServer is the server API for DataServer service.
type DataServerServer interface {
	WriteIntDataArray(context.Context, *IntDataArray) (*ServerSendResponse, error)
	// A simple service for reading a stream of IntDataArrays.
	GetIntDataArrays(*GetIntDataArrayStreamRequest, DataServer_GetIntDataArraysServer) error
}

func RegisterDataServerServer(s *grpc.Server, srv DataServerServer) {
	s.RegisterService(&_DataServer_serviceDesc, srv)
}

func _DataServer_WriteIntDataArray_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntDataArray)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServerServer).WriteIntDataArray(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataArray.DataServer/WriteIntDataArray",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServerServer).WriteIntDataArray(ctx, req.(*IntDataArray))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataServer_GetIntDataArrays_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetIntDataArrayStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataServerServer).GetIntDataArrays(m, &dataServerGetIntDataArraysServer{stream})
}

type DataServer_GetIntDataArraysServer interface {
	Send(*IntDataArray) error
	grpc.ServerStream
}

type dataServerGetIntDataArraysServer struct {
	grpc.ServerStream
}

func (x *dataServerGetIntDataArraysServer) Send(m *IntDataArray) error {
	return x.ServerStream.SendMsg(m)
}

var _DataServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataArray.DataServer",
	HandlerType: (*DataServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteIntDataArray",
			Handler:    _DataServer_WriteIntDataArray_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetIntDataArrays",
			Handler:       _DataServer_GetIntDataArrays_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "data.proto",
}

// DataClientClient is the client API for DataClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataClientClient interface {
	WriteIntDataArray(ctx context.Context, in *IntDataArray, opts ...grpc.CallOption) (*IntDataArray, error)
}

type dataClientClient struct {
	cc *grpc.ClientConn
}

func NewDataClientClient(cc *grpc.ClientConn) DataClientClient {
	return &dataClientClient{cc}
}

func (c *dataClientClient) WriteIntDataArray(ctx context.Context, in *IntDataArray, opts ...grpc.CallOption) (*IntDataArray, error) {
	out := new(IntDataArray)
	err := c.cc.Invoke(ctx, "/dataArray.DataClient/WriteIntDataArray", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataClientServer is the server API for DataClient service.
type DataClientServer interface {
	WriteIntDataArray(context.Context, *IntDataArray) (*IntDataArray, error)
}

func RegisterDataClientServer(s *grpc.Server, srv DataClientServer) {
	s.RegisterService(&_DataClient_serviceDesc, srv)
}

func _DataClient_WriteIntDataArray_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntDataArray)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataClientServer).WriteIntDataArray(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataArray.DataClient/WriteIntDataArray",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataClientServer).WriteIntDataArray(ctx, req.(*IntDataArray))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataClient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataArray.DataClient",
	HandlerType: (*DataClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteIntDataArray",
			Handler:    _DataClient_WriteIntDataArray_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}

func init() { proto.RegisterFile("data.proto", fileDescriptor_871986018790d2fd) }

var fileDescriptor_871986018790d2fd = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0x2c, 0x49,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0xb1, 0x1d, 0x8b, 0x8a, 0x12, 0x2b, 0x95,
	0xe4, 0xb8, 0x64, 0xdc, 0x53, 0x4b, 0x3c, 0xf3, 0x4a, 0x5c, 0x60, 0x42, 0xc1, 0x25, 0x45, 0xa9,
	0x89, 0xb9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x4a, 0x22, 0x5c, 0x42, 0xc1, 0xa9, 0x45,
	0x65, 0xa9, 0x45, 0xc1, 0xa9, 0x79, 0x29, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x4a,
	0x4a, 0x5c, 0x3c, 0xc8, 0x5a, 0x84, 0x84, 0xb8, 0x58, 0x40, 0x46, 0x4a, 0x30, 0x2a, 0x30, 0x6b,
	0xb0, 0x06, 0x81, 0xd9, 0x46, 0x5b, 0x19, 0xb9, 0xb8, 0x40, 0x2a, 0x20, 0xda, 0x85, 0x7c, 0xb9,
	0x04, 0xc3, 0x8b, 0x32, 0x4b, 0x52, 0x51, 0xf4, 0x89, 0xeb, 0xc1, 0x5d, 0xa2, 0x87, 0x2c, 0x21,
	0x25, 0x8b, 0x24, 0x81, 0xc5, 0x7e, 0x06, 0xa1, 0x08, 0x2e, 0x01, 0x34, 0x77, 0x17, 0x0b, 0xa9,
	0x23, 0x69, 0xc2, 0xe7, 0x29, 0x29, 0x5c, 0xd6, 0x2a, 0x31, 0x18, 0x30, 0x1a, 0x85, 0x42, 0x9c,
	0xed, 0x9c, 0x93, 0x99, 0x9a, 0x57, 0x22, 0xe4, 0x4e, 0x92, 0xb3, 0x71, 0x1b, 0x9c, 0xc4, 0x06,
	0x0e, 0x7a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x88, 0x8f, 0x90, 0x88, 0x01, 0x00,
	0x00,
}
