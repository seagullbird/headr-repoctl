// Code generated by protoc-gen-go. DO NOT EDIT.
// source: repoctlsvc.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	repoctlsvc.proto

It has these top-level messages:
	NewSiteRequest
	NewSiteReply
*/
package pb

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

// The sum request contains two parameters.
type NewSiteRequest struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Sitename string `protobuf:"bytes,2,opt,name=sitename" json:"sitename,omitempty"`
}

func (m *NewSiteRequest) Reset()                    { *m = NewSiteRequest{} }
func (m *NewSiteRequest) String() string            { return proto.CompactTextString(m) }
func (*NewSiteRequest) ProtoMessage()               {}
func (*NewSiteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NewSiteRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *NewSiteRequest) GetSitename() string {
	if m != nil {
		return m.Sitename
	}
	return ""
}

// The sum response contains the result of the calculation.
type NewSiteReply struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *NewSiteReply) Reset()                    { *m = NewSiteReply{} }
func (m *NewSiteReply) String() string            { return proto.CompactTextString(m) }
func (*NewSiteReply) ProtoMessage()               {}
func (*NewSiteReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NewSiteReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*NewSiteRequest)(nil), "pb.NewSiteRequest")
	proto.RegisterType((*NewSiteReply)(nil), "pb.NewSiteReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Repoctl service

type RepoctlClient interface {
	// Sums two integers.
	NewSite(ctx context.Context, in *NewSiteRequest, opts ...grpc.CallOption) (*NewSiteReply, error)
}

type repoctlClient struct {
	cc *grpc.ClientConn
}

func NewRepoctlClient(cc *grpc.ClientConn) RepoctlClient {
	return &repoctlClient{cc}
}

func (c *repoctlClient) NewSite(ctx context.Context, in *NewSiteRequest, opts ...grpc.CallOption) (*NewSiteReply, error) {
	out := new(NewSiteReply)
	err := grpc.Invoke(ctx, "/pb.Repoctl/NewSite", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Repoctl service

type RepoctlServer interface {
	// Sums two integers.
	NewSite(context.Context, *NewSiteRequest) (*NewSiteReply, error)
}

func RegisterRepoctlServer(s *grpc.Server, srv RepoctlServer) {
	s.RegisterService(&_Repoctl_serviceDesc, srv)
}

func _Repoctl_NewSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewSiteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoctlServer).NewSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Repoctl/NewSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoctlServer).NewSite(ctx, req.(*NewSiteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Repoctl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Repoctl",
	HandlerType: (*RepoctlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewSite",
			Handler:    _Repoctl_NewSite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "repoctlsvc.proto",
}

func init() { proto.RegisterFile("repoctlsvc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x4a, 0x2d, 0xc8,
	0x4f, 0x2e, 0xc9, 0x29, 0x2e, 0x4b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x52, 0x72, 0xe2, 0xe2, 0xf3, 0x4b, 0x2d, 0x0f, 0xce, 0x2c, 0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x4d, 0xcd, 0x4d, 0xcc, 0xcc, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x82, 0x70, 0x84, 0xa4, 0xb8, 0x38, 0x8a, 0x33, 0x4b, 0x52, 0xf3, 0x12, 0x73, 0x53,
	0x25, 0x98, 0xc0, 0x12, 0x70, 0xbe, 0x92, 0x02, 0x17, 0x0f, 0xdc, 0x8c, 0x82, 0x9c, 0x4a, 0x21,
	0x01, 0x2e, 0xe6, 0xd4, 0xa2, 0x22, 0xa8, 0x7e, 0x10, 0xd3, 0xc8, 0x86, 0x8b, 0x3d, 0x08, 0x62,
	0xbb, 0x90, 0x21, 0x17, 0x3b, 0x54, 0xb1, 0x90, 0x90, 0x5e, 0x41, 0x92, 0x1e, 0xaa, 0xed, 0x52,
	0x02, 0x28, 0x62, 0x05, 0x39, 0x95, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xe7, 0x1a, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x88, 0x2f, 0x83, 0x23, 0xc2, 0x00, 0x00, 0x00,
}
