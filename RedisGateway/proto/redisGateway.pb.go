// Code generated by protoc-gen-go. DO NOT EDIT.
// source: redisGateway.proto

package redisGateway

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// The request message containing the user's name.
type KeyRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=Value,json=value" json:"Value,omitempty"`
}

func (m *KeyRequest) Reset()         { *m = KeyRequest{} }
func (m *KeyRequest) String() string { return proto.CompactTextString(m) }
func (*KeyRequest) ProtoMessage()    {}
func (*KeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_redisGateway_8ad2f783009b06f4, []int{0}
}
func (m *KeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyRequest.Unmarshal(m, b)
}
func (m *KeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyRequest.Marshal(b, m, deterministic)
}
func (dst *KeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRequest.Merge(dst, src)
}
func (m *KeyRequest) XXX_Size() int {
	return xxx_messageInfo_KeyRequest.Size(m)
}
func (m *KeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRequest proto.InternalMessageInfo

func (m *KeyRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Empty struct {

}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_redisGateway_8ad2f783009b06f4, []int{1}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// The response message containing the greetings
type DataReply struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`

}

func (m *DataReply) Reset()         { *m = DataReply{} }
func (m *DataReply) String() string { return proto.CompactTextString(m) }
func (*DataReply) ProtoMessage()    {}
func (*DataReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_redisGateway_8ad2f783009b06f4, []int{2}
}
func (m *DataReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataReply.Unmarshal(m, b)
}
func (m *DataReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataReply.Marshal(b, m, deterministic)
}
func (dst *DataReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataReply.Merge(dst, src)
}
func (m *DataReply) XXX_Size() int {
	return xxx_messageInfo_DataReply.Size(m)
}
func (m *DataReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DataReply.DiscardUnknown(m)
}

var xxx_messageInfo_DataReply proto.InternalMessageInfo

func (m *DataReply) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DataReply) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*KeyRequest)(nil), "redisGateway.KeyRequest")
	proto.RegisterType((*Empty)(nil), "redisGateway.Empty")
	proto.RegisterType((*DataReply)(nil), "redisGateway.DataReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RedisGateway service

type RedisGatewayClient interface {
	GetData(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*DataReply, error)
	SetData(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*Empty, error)
}

type redisGatewayClient struct {
	cc *grpc.ClientConn
}

func NewRedisGatewayClient(cc *grpc.ClientConn) RedisGatewayClient {
	return &redisGatewayClient{cc}
}

func (c *redisGatewayClient) GetData(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*DataReply, error) {
	out := new(DataReply)
	err := grpc.Invoke(ctx, "/redisGateway.redisGateway/getData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisGatewayClient) SetData(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/redisGateway.redisGateway/setData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RedisGateway service

type RedisGatewayServer interface {
	GetData(context.Context, *KeyRequest) (*DataReply, error)
	SetData(context.Context, *KeyRequest) (*Empty, error)
}

func RegisterRedisGatewayServer(s *grpc.Server, srv RedisGatewayServer) {
	s.RegisterService(&_RedisGateway_serviceDesc, srv)
}

func _RedisGateway_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisGatewayServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/redisGateway.redisGateway/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisGatewayServer).GetData(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisGateway_SetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisGatewayServer).SetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/redisGateway.redisGateway/SetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisGatewayServer).SetData(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RedisGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "redisGateway.redisGateway",
	HandlerType: (*RedisGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getData",
			Handler:    _RedisGateway_GetData_Handler,
		},
		{
			MethodName: "setData",
			Handler:    _RedisGateway_SetData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "redisGateway.proto",
}

func init() { proto.RegisterFile("redisGateway.proto", fileDescriptor_redisGateway_8ad2f783009b06f4) }

var fileDescriptor_redisGateway_8ad2f783009b06f4 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4a, 0x4d, 0xc9,
	0x2c, 0x76, 0x4f, 0x2c, 0x49, 0x2d, 0x4f, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x41, 0x16, 0x53, 0x32, 0xe1, 0xe2, 0xf2, 0x4e, 0xad, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x11, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31,
	0x85, 0x44, 0xb8, 0x58, 0xc3, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0xc0, 0x62, 0xac, 0x65, 0x20,
	0x8e, 0x12, 0x3b, 0x17, 0xab, 0x6b, 0x6e, 0x41, 0x49, 0xa5, 0x92, 0x31, 0x17, 0xa7, 0x4b, 0x62,
	0x49, 0x62, 0x50, 0x6a, 0x41, 0x4e, 0x25, 0x76, 0xdd, 0x65, 0x18, 0xba, 0x8d, 0xba, 0x18, 0xb9,
	0x50, 0x1c, 0x21, 0x64, 0xc7, 0xc5, 0x9e, 0x9e, 0x5a, 0x02, 0x32, 0x48, 0x48, 0x42, 0x0f, 0xc5,
	0xc9, 0x08, 0xb7, 0x49, 0x89, 0xa3, 0xca, 0xc0, 0xad, 0x55, 0x62, 0x10, 0xb2, 0xe2, 0x62, 0x2f,
	0x26, 0xa8, 0x5f, 0x18, 0x55, 0x06, 0xe2, 0x7e, 0x06, 0x27, 0x5b, 0x2e, 0xb5, 0xe4, 0xfc, 0x5c,
	0x3d, 0xdf, 0xc4, 0x92, 0x92, 0xe0, 0x8c, 0xc4, 0x9c, 0x9c, 0xfc, 0x72, 0xbd, 0xdc, 0xcc, 0xe4,
	0xa2, 0x7c, 0xc7, 0xe2, 0x62, 0x43, 0x14, 0x1d, 0x4e, 0x82, 0x41, 0x48, 0xbc, 0x00, 0x50, 0x58,
	0x06, 0x30, 0x26, 0xb1, 0x81, 0x03, 0xd5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x32, 0x77, 0x1c,
	0x31, 0x6a, 0x01, 0x00, 0x00,
}
