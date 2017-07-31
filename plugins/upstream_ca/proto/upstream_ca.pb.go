// Code generated by protoc-gen-go. DO NOT EDIT.
// source: upstream_ca.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	upstream_ca.proto

It has these top-level messages:
	SubmitCSRRequest
	SubmitCSRResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto2 "github.com/spiffe/control-plane/plugins/common/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// ConfigureRequest from public import github.com/spiffe/control-plane/plugins/common/proto/common.proto
type ConfigureRequest proto2.ConfigureRequest

func (m *ConfigureRequest) Reset()         { (*proto2.ConfigureRequest)(m).Reset() }
func (m *ConfigureRequest) String() string { return (*proto2.ConfigureRequest)(m).String() }
func (*ConfigureRequest) ProtoMessage()    {}
func (m *ConfigureRequest) GetConfiguration() string {
	return (*proto2.ConfigureRequest)(m).GetConfiguration()
}

// ConfigureResponse from public import github.com/spiffe/control-plane/plugins/common/proto/common.proto
type ConfigureResponse proto2.ConfigureResponse

func (m *ConfigureResponse) Reset()         { (*proto2.ConfigureResponse)(m).Reset() }
func (m *ConfigureResponse) String() string { return (*proto2.ConfigureResponse)(m).String() }
func (*ConfigureResponse) ProtoMessage()    {}
func (m *ConfigureResponse) GetErrorList() []string {
	return (*proto2.ConfigureResponse)(m).GetErrorList()
}

// GetPluginInfoRequest from public import github.com/spiffe/control-plane/plugins/common/proto/common.proto
type GetPluginInfoRequest proto2.GetPluginInfoRequest

func (m *GetPluginInfoRequest) Reset()         { (*proto2.GetPluginInfoRequest)(m).Reset() }
func (m *GetPluginInfoRequest) String() string { return (*proto2.GetPluginInfoRequest)(m).String() }
func (*GetPluginInfoRequest) ProtoMessage()    {}

// GetPluginInfoResponse from public import github.com/spiffe/control-plane/plugins/common/proto/common.proto
type GetPluginInfoResponse proto2.GetPluginInfoResponse

func (m *GetPluginInfoResponse) Reset()         { (*proto2.GetPluginInfoResponse)(m).Reset() }
func (m *GetPluginInfoResponse) String() string { return (*proto2.GetPluginInfoResponse)(m).String() }
func (*GetPluginInfoResponse) ProtoMessage()    {}
func (m *GetPluginInfoResponse) GetPluginName() string {
	return (*proto2.GetPluginInfoResponse)(m).GetPluginName()
}
func (m *GetPluginInfoResponse) GetDescription() string {
	return (*proto2.GetPluginInfoResponse)(m).GetDescription()
}
func (m *GetPluginInfoResponse) GetDateCreated() string {
	return (*proto2.GetPluginInfoResponse)(m).GetDateCreated()
}
func (m *GetPluginInfoResponse) GetLocation() string {
	return (*proto2.GetPluginInfoResponse)(m).GetLocation()
}
func (m *GetPluginInfoResponse) GetVersion() string {
	return (*proto2.GetPluginInfoResponse)(m).GetVersion()
}
func (m *GetPluginInfoResponse) GetAuthor() string {
	return (*proto2.GetPluginInfoResponse)(m).GetAuthor()
}
func (m *GetPluginInfoResponse) GetCompany() string {
	return (*proto2.GetPluginInfoResponse)(m).GetCompany()
}

type SubmitCSRRequest struct {
	Csr []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
}

func (m *SubmitCSRRequest) Reset()                    { *m = SubmitCSRRequest{} }
func (m *SubmitCSRRequest) String() string            { return proto1.CompactTextString(m) }
func (*SubmitCSRRequest) ProtoMessage()               {}
func (*SubmitCSRRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SubmitCSRRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

type SubmitCSRResponse struct {
	Cert                []byte `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
	UpstreamTrustBundle []byte `protobuf:"bytes,2,opt,name=upstreamTrustBundle,proto3" json:"upstreamTrustBundle,omitempty"`
}

func (m *SubmitCSRResponse) Reset()                    { *m = SubmitCSRResponse{} }
func (m *SubmitCSRResponse) String() string            { return proto1.CompactTextString(m) }
func (*SubmitCSRResponse) ProtoMessage()               {}
func (*SubmitCSRResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SubmitCSRResponse) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

func (m *SubmitCSRResponse) GetUpstreamTrustBundle() []byte {
	if m != nil {
		return m.UpstreamTrustBundle
	}
	return nil
}

func init() {
	proto1.RegisterType((*SubmitCSRRequest)(nil), "proto.SubmitCSRRequest")
	proto1.RegisterType((*SubmitCSRResponse)(nil), "proto.SubmitCSRResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Node service

type NodeClient interface {
	Configure(ctx context.Context, in *proto2.ConfigureRequest, opts ...grpc.CallOption) (*proto2.ConfigureResponse, error)
	GetPluginInfo(ctx context.Context, in *proto2.GetPluginInfoRequest, opts ...grpc.CallOption) (*proto2.GetPluginInfoResponse, error)
	SubmitCSR(ctx context.Context, in *SubmitCSRRequest, opts ...grpc.CallOption) (*SubmitCSRResponse, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Configure(ctx context.Context, in *proto2.ConfigureRequest, opts ...grpc.CallOption) (*proto2.ConfigureResponse, error) {
	out := new(proto2.ConfigureResponse)
	err := grpc.Invoke(ctx, "/proto.node/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) GetPluginInfo(ctx context.Context, in *proto2.GetPluginInfoRequest, opts ...grpc.CallOption) (*proto2.GetPluginInfoResponse, error) {
	out := new(proto2.GetPluginInfoResponse)
	err := grpc.Invoke(ctx, "/proto.node/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) SubmitCSR(ctx context.Context, in *SubmitCSRRequest, opts ...grpc.CallOption) (*SubmitCSRResponse, error) {
	out := new(SubmitCSRResponse)
	err := grpc.Invoke(ctx, "/proto.node/SubmitCSR", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeServer interface {
	Configure(context.Context, *proto2.ConfigureRequest) (*proto2.ConfigureResponse, error)
	GetPluginInfo(context.Context, *proto2.GetPluginInfoRequest) (*proto2.GetPluginInfoResponse, error)
	SubmitCSR(context.Context, *SubmitCSRRequest) (*SubmitCSRResponse, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto2.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.node/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Configure(ctx, req.(*proto2.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto2.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.node/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).GetPluginInfo(ctx, req.(*proto2.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_SubmitCSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitCSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).SubmitCSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.node/SubmitCSR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).SubmitCSR(ctx, req.(*SubmitCSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _Node_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _Node_GetPluginInfo_Handler,
		},
		{
			MethodName: "SubmitCSR",
			Handler:    _Node_SubmitCSR_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upstream_ca.proto",
}

func init() { proto1.RegisterFile("upstream_ca.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x09, 0x14, 0xa4, 0x9e, 0x40, 0x6a, 0xcd, 0x40, 0x54, 0x18, 0x50, 0xc5, 0xc0, 0x42,
	0x8c, 0x60, 0xaf, 0x04, 0x1d, 0x10, 0x4c, 0x55, 0xca, 0xc2, 0x84, 0x12, 0xf7, 0x12, 0x2c, 0xc5,
	0x3e, 0xe3, 0x3f, 0x5f, 0x94, 0x4f, 0x84, 0x70, 0xdd, 0x08, 0x68, 0x27, 0x9f, 0xef, 0xbd, 0x7b,
	0xfa, 0xdd, 0xc1, 0x38, 0x18, 0xe7, 0x2d, 0x56, 0xea, 0x5d, 0x54, 0x85, 0xb1, 0xe4, 0x89, 0x1d,
	0xc6, 0x67, 0xf2, 0xd0, 0x4a, 0xff, 0x11, 0xea, 0x42, 0x90, 0xe2, 0xce, 0xc8, 0xa6, 0x41, 0x2e,
	0x48, 0x7b, 0x4b, 0xdd, 0x8d, 0xe9, 0x2a, 0x8d, 0xdc, 0x74, 0xa1, 0x95, 0xda, 0x71, 0x41, 0x4a,
	0x91, 0xe6, 0x71, 0x2a, 0x7d, 0xd6, 0x49, 0xd3, 0x2b, 0x18, 0x2d, 0x43, 0xad, 0xa4, 0x9f, 0x2f,
	0xcb, 0x12, 0x3f, 0x03, 0x3a, 0xcf, 0x46, 0x70, 0x20, 0x9c, 0xcd, 0xb3, 0xcb, 0xec, 0xfa, 0xb8,
	0xfc, 0x29, 0xa7, 0x6f, 0x30, 0xfe, 0xe5, 0x72, 0x86, 0xb4, 0x43, 0xc6, 0x60, 0x20, 0xd0, 0xfa,
	0xe4, 0x8b, 0x35, 0xbb, 0x85, 0xd3, 0x0d, 0xed, 0xab, 0x0d, 0xce, 0x3f, 0x06, 0xbd, 0xea, 0x30,
	0xdf, 0x8f, 0x96, 0x5d, 0xd2, 0xdd, 0x57, 0x06, 0x03, 0x4d, 0x2b, 0x64, 0x33, 0x18, 0xce, 0x49,
	0x37, 0xb2, 0x0d, 0x16, 0xd9, 0xd9, 0x1a, 0xaf, 0xe8, 0x3b, 0x89, 0x6d, 0x92, 0x6f, 0x0b, 0x09,
	0xe7, 0x05, 0x4e, 0x9e, 0xd0, 0x2f, 0xe2, 0xc6, 0xcf, 0xba, 0x21, 0x76, 0x9e, 0xac, 0x7f, 0xba,
	0x9b, 0x9c, 0x8b, 0xdd, 0x62, 0xca, 0x9a, 0xc1, 0xb0, 0xdf, 0xb7, 0x67, 0xf9, 0x7f, 0xa7, 0x9e,
	0x65, 0xeb, 0x34, 0x8b, 0xbd, 0xfa, 0x28, 0x4a, 0xf7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xba,
	0x99, 0xd9, 0xa9, 0xbd, 0x01, 0x00, 0x00,
}
