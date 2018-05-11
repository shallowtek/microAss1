// Code generated by protoc-gen-go. DO NOT EDIT.
// source: redisGateway.proto

package redisgateway

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
	Key                  string   `protobuf:"bytes,1,opt,name=Key,json=key" json:"Key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=Value,json=value" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyRequest) Reset()         { *m = KeyRequest{} }
func (m *KeyRequest) String() string { return proto.CompactTextString(m) }
func (*KeyRequest) ProtoMessage()    {}
func (*KeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_redisGateway_2de937eafae8f24a, []int{0}
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

func init() {
	proto.RegisterType((*KeyRequest)(nil), "redisgateway.KeyRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RedisGateway service

type RedisGatewayClient interface {
	GetData(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*KeyRequest, error)
	SetData(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*KeyRequest, error)
}

type redisGatewayClient struct {
	cc *grpc.ClientConn
}

func NewRedisGatewayClient(cc *grpc.ClientConn) RedisGatewayClient {
	return &redisGatewayClient{cc}
}

func (c *redisGatewayClient) GetData(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*KeyRequest, error) {
	out := new(KeyRequest)
	err := grpc.Invoke(ctx, "/RedisGateway/getData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisGatewayClient) SetData(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*KeyRequest, error) {
	out := new(KeyRequest)
	err := grpc.Invoke(ctx, "/RedisGateway/setData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RedisGateway service

type RedisGatewayServer interface {
	GetData(context.Context, *KeyRequest) (*KeyRequest, error)
	SetData(context.Context, *KeyRequest) (*KeyRequest, error)
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
		FullMethod: "/RedisGateway/GetData",
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
		FullMethod: "/RedisGateway/SetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisGatewayServer).SetData(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RedisGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RedisGateway",
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

func init() { proto.RegisterFile("redisGateway.proto", fileDescriptor_redisGateway_2de937eafae8f24a) }

var fileDescriptor_redisGateway_2de937eafae8f24a = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4a, 0x4d, 0xc9,
	0x2c, 0x76, 0x4f, 0x2c, 0x49, 0x2d, 0x4f, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x01, 0x8b, 0xa5, 0x43, 0xc4, 0x94, 0x4c, 0xb8, 0xb8, 0xbc, 0x53, 0x2b, 0x83, 0x52, 0x0b, 0x4b,
	0x53, 0x8b, 0x4b, 0x84, 0x04, 0xb8, 0x98, 0xbd, 0x53, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0x98, 0xb3, 0x53, 0x2b, 0x85, 0x44, 0xb8, 0x58, 0xc3, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98,
	0xc0, 0x62, 0xac, 0x65, 0x20, 0x8e, 0xd1, 0x04, 0x46, 0x2e, 0x9e, 0x20, 0x24, 0xa3, 0x85, 0xec,
	0xb9, 0xd8, 0xd3, 0x53, 0x4b, 0x5c, 0x12, 0x4b, 0x12, 0x85, 0x24, 0xf4, 0x90, 0x2d, 0xd0, 0x43,
	0x98, 0x2e, 0x85, 0x53, 0x46, 0x89, 0x01, 0x64, 0x40, 0x31, 0x25, 0x06, 0x38, 0x19, 0x73, 0xc9,
	0xe4, 0x66, 0x26, 0x17, 0xe5, 0x3b, 0x16, 0x17, 0x1b, 0xea, 0x21, 0xbb, 0x4d, 0x1f, 0xec, 0x6d,
	0x27, 0x41, 0x64, 0xb1, 0x00, 0x90, 0x50, 0x00, 0x63, 0x12, 0x1b, 0x58, 0xce, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x0e, 0x53, 0x13, 0xc6, 0x28, 0x01, 0x00, 0x00,
}
