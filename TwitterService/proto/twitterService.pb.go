// Code generated by protoc-gen-go. DO NOT EDIT.
// source: twitterService.proto

/*
Package twitterservice is a generated protocol buffer package.

It is generated from these files:
	twitterService.proto

It has these top-level messages:
	TweetsRequest
	TweetsReply
*/
package twitterservice

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
type TweetsRequest struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Minutes string `protobuf:"bytes,2,opt,name=minutes" json:"minutes,omitempty"`
}

func (m *TweetsRequest) Reset()                    { *m = TweetsRequest{} }
func (m *TweetsRequest) String() string            { return proto.CompactTextString(m) }
func (*TweetsRequest) ProtoMessage()               {}
func (*TweetsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TweetsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TweetsRequest) GetMinutes() string {
	if m != nil {
		return m.Minutes
	}
	return ""
}

// The response message containing the greetings
type TweetsReply struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *TweetsReply) Reset()                    { *m = TweetsReply{} }
func (m *TweetsReply) String() string            { return proto.CompactTextString(m) }
func (*TweetsReply) ProtoMessage()               {}
func (*TweetsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TweetsReply) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*TweetsRequest)(nil), "twitterservice.TweetsRequest")
	proto.RegisterType((*TweetsReply)(nil), "twitterservice.TweetsReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TwitterService service

type TwitterServiceClient interface {
	// Sends a greeting
	GetTweets(ctx context.Context, in *TweetsRequest, opts ...grpc.CallOption) (TwitterService_GetTweetsClient, error)
}

type twitterServiceClient struct {
	cc *grpc.ClientConn
}

func NewTwitterServiceClient(cc *grpc.ClientConn) TwitterServiceClient {
	return &twitterServiceClient{cc}
}

func (c *twitterServiceClient) GetTweets(ctx context.Context, in *TweetsRequest, opts ...grpc.CallOption) (TwitterService_GetTweetsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TwitterService_serviceDesc.Streams[0], c.cc, "/twitterservice.TwitterService/getTweets", opts...)
	if err != nil {
		return nil, err
	}
	x := &twitterServiceGetTweetsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TwitterService_GetTweetsClient interface {
	Recv() (*TweetsReply, error)
	grpc.ClientStream
}

type twitterServiceGetTweetsClient struct {
	grpc.ClientStream
}

func (x *twitterServiceGetTweetsClient) Recv() (*TweetsReply, error) {
	m := new(TweetsReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TwitterService service

type TwitterServiceServer interface {
	// Sends a greeting
	GetTweets(*TweetsRequest, TwitterService_GetTweetsServer) error
}

func RegisterTwitterServiceServer(s *grpc.Server, srv TwitterServiceServer) {
	s.RegisterService(&_TwitterService_serviceDesc, srv)
}

func _TwitterService_GetTweets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TweetsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TwitterServiceServer).GetTweets(m, &twitterServiceGetTweetsServer{stream})
}

type TwitterService_GetTweetsServer interface {
	Send(*TweetsReply) error
	grpc.ServerStream
}

type twitterServiceGetTweetsServer struct {
	grpc.ServerStream
}

func (x *twitterServiceGetTweetsServer) Send(m *TweetsReply) error {
	return x.ServerStream.SendMsg(m)
}

var _TwitterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "twitterservice.TwitterService",
	HandlerType: (*TwitterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getTweets",
			Handler:       _TwitterService_GetTweets_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "twitterService.proto",
}

func init() { proto.RegisterFile("twitterService.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x29, 0xcf, 0x2c,
	0x29, 0x49, 0x2d, 0x0a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x83, 0x8a, 0x16, 0x43, 0x44, 0x95, 0x6c, 0xb9, 0x78, 0x43, 0xca, 0x53, 0x53, 0x4b,
	0x8a, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xcc, 0xbc,
	0xd2, 0x92, 0xd4, 0x62, 0x09, 0x26, 0xb0, 0x30, 0x8c, 0xab, 0xa4, 0xc8, 0xc5, 0x0d, 0xd3, 0x5e,
	0x90, 0x53, 0x09, 0xd2, 0x5c, 0x92, 0x5a, 0x51, 0x02, 0xd3, 0x0c, 0x62, 0x1b, 0xc5, 0x72, 0xf1,
	0x85, 0xa0, 0xb8, 0x44, 0xc8, 0x9b, 0x8b, 0x33, 0x3d, 0xb5, 0x04, 0xa2, 0x4f, 0x48, 0x56, 0x0f,
	0xd5, 0x45, 0x7a, 0x28, 0xce, 0x91, 0x92, 0xc6, 0x25, 0x5d, 0x90, 0x53, 0xa9, 0xc4, 0x60, 0xc0,
	0xe8, 0x64, 0xce, 0x25, 0x97, 0x9b, 0x99, 0x5c, 0x94, 0xef, 0x58, 0x5c, 0x6c, 0xa8, 0x87, 0x6a,
	0x91, 0x3e, 0xd8, 0xcb, 0x4e, 0xc2, 0xa8, 0xa2, 0x01, 0x20, 0xc1, 0x00, 0xc6, 0x24, 0x36, 0xb0,
	0xac, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x84, 0xe5, 0xd2, 0x82, 0x28, 0x01, 0x00, 0x00,
}
