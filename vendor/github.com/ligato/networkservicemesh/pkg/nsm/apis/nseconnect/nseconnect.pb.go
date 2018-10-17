// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nseconnect.proto

package nseconnect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/ligato/networkservicemesh/pkg/nsm/apis/common"
import netmesh "github.com/ligato/networkservicemesh/pkg/nsm/apis/netmesh"

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

// EndpointConnectionRequest is sent by a NSM to NSE to build a connection.
type EndpointConnectionRequest struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	NetworkServiceName   string   `protobuf:"bytes,2,opt,name=network_service_name,json=networkServiceName,proto3" json:"network_service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndpointConnectionRequest) Reset()         { *m = EndpointConnectionRequest{} }
func (m *EndpointConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*EndpointConnectionRequest) ProtoMessage()    {}
func (*EndpointConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_nseconnect_54fa9345bcf3911d, []int{0}
}
func (m *EndpointConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointConnectionRequest.Unmarshal(m, b)
}
func (m *EndpointConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointConnectionRequest.Marshal(b, m, deterministic)
}
func (dst *EndpointConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointConnectionRequest.Merge(dst, src)
}
func (m *EndpointConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_EndpointConnectionRequest.Size(m)
}
func (m *EndpointConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointConnectionRequest proto.InternalMessageInfo

func (m *EndpointConnectionRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *EndpointConnectionRequest) GetNetworkServiceName() string {
	if m != nil {
		return m.NetworkServiceName
	}
	return ""
}

// EndpointConnectionReply is sent back by NSE to NSM with information required for
// dataplane programming.
type EndpointConnectionReply struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	NetworkServiceName   string   `protobuf:"bytes,2,opt,name=network_service_name,json=networkServiceName,proto3" json:"network_service_name,omitempty"`
	LinuxNamespace       string   `protobuf:"bytes,3,opt,name=linux_namespace,json=linuxNamespace,proto3" json:"linux_namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndpointConnectionReply) Reset()         { *m = EndpointConnectionReply{} }
func (m *EndpointConnectionReply) String() string { return proto.CompactTextString(m) }
func (*EndpointConnectionReply) ProtoMessage()    {}
func (*EndpointConnectionReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_nseconnect_54fa9345bcf3911d, []int{1}
}
func (m *EndpointConnectionReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointConnectionReply.Unmarshal(m, b)
}
func (m *EndpointConnectionReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointConnectionReply.Marshal(b, m, deterministic)
}
func (dst *EndpointConnectionReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointConnectionReply.Merge(dst, src)
}
func (m *EndpointConnectionReply) XXX_Size() int {
	return xxx_messageInfo_EndpointConnectionReply.Size(m)
}
func (m *EndpointConnectionReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointConnectionReply.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointConnectionReply proto.InternalMessageInfo

func (m *EndpointConnectionReply) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *EndpointConnectionReply) GetNetworkServiceName() string {
	if m != nil {
		return m.NetworkServiceName
	}
	return ""
}

func (m *EndpointConnectionReply) GetLinuxNamespace() string {
	if m != nil {
		return m.LinuxNamespace
	}
	return ""
}

// EndpointConnectionInterface is sent by a NSM to NSE to inform NSE about
// the interface name which was created by the dataplane controller.
type EndpointConnectionMechanism struct {
	RequestId            string                 `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	LocalMechanism       *common.LocalMechanism `protobuf:"bytes,2,opt,name=local_mechanism,json=localMechanism,proto3" json:"local_mechanism,omitempty"`
	NetworkServiceName   string                 `protobuf:"bytes,3,opt,name=network_service_name,json=networkServiceName,proto3" json:"network_service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *EndpointConnectionMechanism) Reset()         { *m = EndpointConnectionMechanism{} }
func (m *EndpointConnectionMechanism) String() string { return proto.CompactTextString(m) }
func (*EndpointConnectionMechanism) ProtoMessage()    {}
func (*EndpointConnectionMechanism) Descriptor() ([]byte, []int) {
	return fileDescriptor_nseconnect_54fa9345bcf3911d, []int{2}
}
func (m *EndpointConnectionMechanism) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointConnectionMechanism.Unmarshal(m, b)
}
func (m *EndpointConnectionMechanism) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointConnectionMechanism.Marshal(b, m, deterministic)
}
func (dst *EndpointConnectionMechanism) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointConnectionMechanism.Merge(dst, src)
}
func (m *EndpointConnectionMechanism) XXX_Size() int {
	return xxx_messageInfo_EndpointConnectionMechanism.Size(m)
}
func (m *EndpointConnectionMechanism) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointConnectionMechanism.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointConnectionMechanism proto.InternalMessageInfo

func (m *EndpointConnectionMechanism) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *EndpointConnectionMechanism) GetLocalMechanism() *common.LocalMechanism {
	if m != nil {
		return m.LocalMechanism
	}
	return nil
}

func (m *EndpointConnectionMechanism) GetNetworkServiceName() string {
	if m != nil {
		return m.NetworkServiceName
	}
	return ""
}

// EndpointConnectionInterfaceReply is sent back by NSE to NSM to confirm the presence
// of the programmed interface.
type EndpointConnectionMechanismReply struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	MechanismFound       bool     `protobuf:"varint,2,opt,name=mechanism_found,json=mechanismFound,proto3" json:"mechanism_found,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndpointConnectionMechanismReply) Reset()         { *m = EndpointConnectionMechanismReply{} }
func (m *EndpointConnectionMechanismReply) String() string { return proto.CompactTextString(m) }
func (*EndpointConnectionMechanismReply) ProtoMessage()    {}
func (*EndpointConnectionMechanismReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_nseconnect_54fa9345bcf3911d, []int{3}
}
func (m *EndpointConnectionMechanismReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointConnectionMechanismReply.Unmarshal(m, b)
}
func (m *EndpointConnectionMechanismReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointConnectionMechanismReply.Marshal(b, m, deterministic)
}
func (dst *EndpointConnectionMechanismReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointConnectionMechanismReply.Merge(dst, src)
}
func (m *EndpointConnectionMechanismReply) XXX_Size() int {
	return xxx_messageInfo_EndpointConnectionMechanismReply.Size(m)
}
func (m *EndpointConnectionMechanismReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointConnectionMechanismReply.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointConnectionMechanismReply proto.InternalMessageInfo

func (m *EndpointConnectionMechanismReply) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *EndpointConnectionMechanismReply) GetMechanismFound() bool {
	if m != nil {
		return m.MechanismFound
	}
	return false
}

type EndpointAdvertiseRequest struct {
	RequestId            string                          `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	NetworkEndpoint      *netmesh.NetworkServiceEndpoint `protobuf:"bytes,2,opt,name=network_endpoint,json=networkEndpoint,proto3" json:"network_endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *EndpointAdvertiseRequest) Reset()         { *m = EndpointAdvertiseRequest{} }
func (m *EndpointAdvertiseRequest) String() string { return proto.CompactTextString(m) }
func (*EndpointAdvertiseRequest) ProtoMessage()    {}
func (*EndpointAdvertiseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_nseconnect_54fa9345bcf3911d, []int{4}
}
func (m *EndpointAdvertiseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointAdvertiseRequest.Unmarshal(m, b)
}
func (m *EndpointAdvertiseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointAdvertiseRequest.Marshal(b, m, deterministic)
}
func (dst *EndpointAdvertiseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointAdvertiseRequest.Merge(dst, src)
}
func (m *EndpointAdvertiseRequest) XXX_Size() int {
	return xxx_messageInfo_EndpointAdvertiseRequest.Size(m)
}
func (m *EndpointAdvertiseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointAdvertiseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointAdvertiseRequest proto.InternalMessageInfo

func (m *EndpointAdvertiseRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *EndpointAdvertiseRequest) GetNetworkEndpoint() *netmesh.NetworkServiceEndpoint {
	if m != nil {
		return m.NetworkEndpoint
	}
	return nil
}

type EndpointAdvertiseReply struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Accepted             bool     `protobuf:"varint,2,opt,name=accepted,proto3" json:"accepted,omitempty"`
	AdmissionError       string   `protobuf:"bytes,3,opt,name=admission_error,json=admissionError,proto3" json:"admission_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndpointAdvertiseReply) Reset()         { *m = EndpointAdvertiseReply{} }
func (m *EndpointAdvertiseReply) String() string { return proto.CompactTextString(m) }
func (*EndpointAdvertiseReply) ProtoMessage()    {}
func (*EndpointAdvertiseReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_nseconnect_54fa9345bcf3911d, []int{5}
}
func (m *EndpointAdvertiseReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointAdvertiseReply.Unmarshal(m, b)
}
func (m *EndpointAdvertiseReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointAdvertiseReply.Marshal(b, m, deterministic)
}
func (dst *EndpointAdvertiseReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointAdvertiseReply.Merge(dst, src)
}
func (m *EndpointAdvertiseReply) XXX_Size() int {
	return xxx_messageInfo_EndpointAdvertiseReply.Size(m)
}
func (m *EndpointAdvertiseReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointAdvertiseReply.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointAdvertiseReply proto.InternalMessageInfo

func (m *EndpointAdvertiseReply) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *EndpointAdvertiseReply) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *EndpointAdvertiseReply) GetAdmissionError() string {
	if m != nil {
		return m.AdmissionError
	}
	return ""
}

// Message sent by NSE to NSM informing of a removal previously advertised
// endpoint. NSM will attempt to locate Customer Resource and delete it.
//
type EndpointRemoveRequest struct {
	RequestId            string                          `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	NetworkEndpoint      *netmesh.NetworkServiceEndpoint `protobuf:"bytes,2,opt,name=network_endpoint,json=networkEndpoint,proto3" json:"network_endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *EndpointRemoveRequest) Reset()         { *m = EndpointRemoveRequest{} }
func (m *EndpointRemoveRequest) String() string { return proto.CompactTextString(m) }
func (*EndpointRemoveRequest) ProtoMessage()    {}
func (*EndpointRemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_nseconnect_54fa9345bcf3911d, []int{6}
}
func (m *EndpointRemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointRemoveRequest.Unmarshal(m, b)
}
func (m *EndpointRemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointRemoveRequest.Marshal(b, m, deterministic)
}
func (dst *EndpointRemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointRemoveRequest.Merge(dst, src)
}
func (m *EndpointRemoveRequest) XXX_Size() int {
	return xxx_messageInfo_EndpointRemoveRequest.Size(m)
}
func (m *EndpointRemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointRemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointRemoveRequest proto.InternalMessageInfo

func (m *EndpointRemoveRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *EndpointRemoveRequest) GetNetworkEndpoint() *netmesh.NetworkServiceEndpoint {
	if m != nil {
		return m.NetworkEndpoint
	}
	return nil
}

type EndpointRemoveReply struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Accepted             bool     `protobuf:"varint,2,opt,name=accepted,proto3" json:"accepted,omitempty"`
	AdmissionError       string   `protobuf:"bytes,3,opt,name=admission_error,json=admissionError,proto3" json:"admission_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndpointRemoveReply) Reset()         { *m = EndpointRemoveReply{} }
func (m *EndpointRemoveReply) String() string { return proto.CompactTextString(m) }
func (*EndpointRemoveReply) ProtoMessage()    {}
func (*EndpointRemoveReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_nseconnect_54fa9345bcf3911d, []int{7}
}
func (m *EndpointRemoveReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointRemoveReply.Unmarshal(m, b)
}
func (m *EndpointRemoveReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointRemoveReply.Marshal(b, m, deterministic)
}
func (dst *EndpointRemoveReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointRemoveReply.Merge(dst, src)
}
func (m *EndpointRemoveReply) XXX_Size() int {
	return xxx_messageInfo_EndpointRemoveReply.Size(m)
}
func (m *EndpointRemoveReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointRemoveReply.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointRemoveReply proto.InternalMessageInfo

func (m *EndpointRemoveReply) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *EndpointRemoveReply) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *EndpointRemoveReply) GetAdmissionError() string {
	if m != nil {
		return m.AdmissionError
	}
	return ""
}

func init() {
	proto.RegisterType((*EndpointConnectionRequest)(nil), "nseconnect.EndpointConnectionRequest")
	proto.RegisterType((*EndpointConnectionReply)(nil), "nseconnect.EndpointConnectionReply")
	proto.RegisterType((*EndpointConnectionMechanism)(nil), "nseconnect.EndpointConnectionMechanism")
	proto.RegisterType((*EndpointConnectionMechanismReply)(nil), "nseconnect.EndpointConnectionMechanismReply")
	proto.RegisterType((*EndpointAdvertiseRequest)(nil), "nseconnect.EndpointAdvertiseRequest")
	proto.RegisterType((*EndpointAdvertiseReply)(nil), "nseconnect.EndpointAdvertiseReply")
	proto.RegisterType((*EndpointRemoveRequest)(nil), "nseconnect.EndpointRemoveRequest")
	proto.RegisterType((*EndpointRemoveReply)(nil), "nseconnect.EndpointRemoveReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EndpointConnectionClient is the client API for EndpointConnection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndpointConnectionClient interface {
	RequestEndpointConnection(ctx context.Context, in *EndpointConnectionRequest, opts ...grpc.CallOption) (*EndpointConnectionReply, error)
	SendEndpointConnectionMechanism(ctx context.Context, in *EndpointConnectionMechanism, opts ...grpc.CallOption) (*EndpointConnectionMechanismReply, error)
}

type endpointConnectionClient struct {
	cc *grpc.ClientConn
}

func NewEndpointConnectionClient(cc *grpc.ClientConn) EndpointConnectionClient {
	return &endpointConnectionClient{cc}
}

func (c *endpointConnectionClient) RequestEndpointConnection(ctx context.Context, in *EndpointConnectionRequest, opts ...grpc.CallOption) (*EndpointConnectionReply, error) {
	out := new(EndpointConnectionReply)
	err := c.cc.Invoke(ctx, "/nseconnect.EndpointConnection/RequestEndpointConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointConnectionClient) SendEndpointConnectionMechanism(ctx context.Context, in *EndpointConnectionMechanism, opts ...grpc.CallOption) (*EndpointConnectionMechanismReply, error) {
	out := new(EndpointConnectionMechanismReply)
	err := c.cc.Invoke(ctx, "/nseconnect.EndpointConnection/SendEndpointConnectionMechanism", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointConnectionServer is the server API for EndpointConnection service.
type EndpointConnectionServer interface {
	RequestEndpointConnection(context.Context, *EndpointConnectionRequest) (*EndpointConnectionReply, error)
	SendEndpointConnectionMechanism(context.Context, *EndpointConnectionMechanism) (*EndpointConnectionMechanismReply, error)
}

func RegisterEndpointConnectionServer(s *grpc.Server, srv EndpointConnectionServer) {
	s.RegisterService(&_EndpointConnection_serviceDesc, srv)
}

func _EndpointConnection_RequestEndpointConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointConnectionServer).RequestEndpointConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nseconnect.EndpointConnection/RequestEndpointConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointConnectionServer).RequestEndpointConnection(ctx, req.(*EndpointConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndpointConnection_SendEndpointConnectionMechanism_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointConnectionMechanism)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointConnectionServer).SendEndpointConnectionMechanism(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nseconnect.EndpointConnection/SendEndpointConnectionMechanism",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointConnectionServer).SendEndpointConnectionMechanism(ctx, req.(*EndpointConnectionMechanism))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndpointConnection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nseconnect.EndpointConnection",
	HandlerType: (*EndpointConnectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestEndpointConnection",
			Handler:    _EndpointConnection_RequestEndpointConnection_Handler,
		},
		{
			MethodName: "SendEndpointConnectionMechanism",
			Handler:    _EndpointConnection_SendEndpointConnectionMechanism_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nseconnect.proto",
}

// EndpointOperationsClient is the client API for EndpointOperations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndpointOperationsClient interface {
	AdvertiseEndpoint(ctx context.Context, in *EndpointAdvertiseRequest, opts ...grpc.CallOption) (*EndpointAdvertiseReply, error)
	RemoveEndpoint(ctx context.Context, in *EndpointRemoveRequest, opts ...grpc.CallOption) (*EndpointRemoveReply, error)
}

type endpointOperationsClient struct {
	cc *grpc.ClientConn
}

func NewEndpointOperationsClient(cc *grpc.ClientConn) EndpointOperationsClient {
	return &endpointOperationsClient{cc}
}

func (c *endpointOperationsClient) AdvertiseEndpoint(ctx context.Context, in *EndpointAdvertiseRequest, opts ...grpc.CallOption) (*EndpointAdvertiseReply, error) {
	out := new(EndpointAdvertiseReply)
	err := c.cc.Invoke(ctx, "/nseconnect.EndpointOperations/AdvertiseEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointOperationsClient) RemoveEndpoint(ctx context.Context, in *EndpointRemoveRequest, opts ...grpc.CallOption) (*EndpointRemoveReply, error) {
	out := new(EndpointRemoveReply)
	err := c.cc.Invoke(ctx, "/nseconnect.EndpointOperations/RemoveEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointOperationsServer is the server API for EndpointOperations service.
type EndpointOperationsServer interface {
	AdvertiseEndpoint(context.Context, *EndpointAdvertiseRequest) (*EndpointAdvertiseReply, error)
	RemoveEndpoint(context.Context, *EndpointRemoveRequest) (*EndpointRemoveReply, error)
}

func RegisterEndpointOperationsServer(s *grpc.Server, srv EndpointOperationsServer) {
	s.RegisterService(&_EndpointOperations_serviceDesc, srv)
}

func _EndpointOperations_AdvertiseEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointAdvertiseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointOperationsServer).AdvertiseEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nseconnect.EndpointOperations/AdvertiseEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointOperationsServer).AdvertiseEndpoint(ctx, req.(*EndpointAdvertiseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndpointOperations_RemoveEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointOperationsServer).RemoveEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nseconnect.EndpointOperations/RemoveEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointOperationsServer).RemoveEndpoint(ctx, req.(*EndpointRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndpointOperations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nseconnect.EndpointOperations",
	HandlerType: (*EndpointOperationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdvertiseEndpoint",
			Handler:    _EndpointOperations_AdvertiseEndpoint_Handler,
		},
		{
			MethodName: "RemoveEndpoint",
			Handler:    _EndpointOperations_RemoveEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nseconnect.proto",
}

func init() { proto.RegisterFile("nseconnect.proto", fileDescriptor_nseconnect_54fa9345bcf3911d) }

var fileDescriptor_nseconnect_54fa9345bcf3911d = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x26, 0x2e, 0xc8, 0xee, 0x11, 0xda, 0x35, 0xea, 0xda, 0x8d, 0x48, 0xd7, 0xa8, 0xd4, 0x0b,
	0x69, 0xa4, 0x3e, 0x80, 0x88, 0x54, 0x51, 0xb4, 0x42, 0xd6, 0x5b, 0x09, 0xb3, 0x93, 0x63, 0x3b,
	0x6e, 0x66, 0x26, 0x66, 0xd2, 0xba, 0x05, 0xaf, 0x04, 0x9f, 0xc0, 0x07, 0xf1, 0x39, 0x7c, 0x22,
	0x6f, 0x25, 0x93, 0x99, 0x71, 0xbb, 0x9b, 0xfe, 0x08, 0xe2, 0x5e, 0xa5, 0xf9, 0xce, 0x37, 0xdf,
	0xf9, 0xbe, 0xd3, 0x93, 0x81, 0x5d, 0xa1, 0x90, 0x4a, 0x21, 0x90, 0x96, 0xfd, 0xbc, 0x90, 0xa5,
	0xf4, 0xe1, 0x0f, 0x12, 0xbc, 0x18, 0xb3, 0x72, 0x32, 0x3d, 0xea, 0x53, 0xc9, 0xa3, 0x8c, 0x8d,
	0x49, 0x29, 0x23, 0x81, 0xe5, 0x67, 0x59, 0x1c, 0x2b, 0x2c, 0x66, 0x8c, 0x22, 0x47, 0x35, 0x89,
	0xf2, 0xe3, 0x71, 0x24, 0x14, 0x8f, 0x48, 0xce, 0x54, 0x55, 0xd7, 0xa0, 0x79, 0xd6, 0xa2, 0xc1,
	0xf0, 0xef, 0x85, 0xa8, 0xe4, 0x5c, 0x0a, 0xf3, 0xa8, 0x65, 0xc2, 0x0c, 0xf6, 0x87, 0x22, 0xcd,
	0x25, 0x13, 0xe5, 0xb3, 0xda, 0x22, 0x93, 0x22, 0xc6, 0x4f, 0x53, 0x54, 0xa5, 0x7f, 0x1b, 0xa0,
	0xa8, 0x7f, 0x26, 0x2c, 0xed, 0x78, 0x07, 0xde, 0x83, 0x9d, 0x78, 0xc7, 0x20, 0x2f, 0x53, 0xff,
	0x11, 0x5c, 0x37, 0x2d, 0x13, 0xd3, 0x33, 0x11, 0x84, 0x63, 0xe7, 0x92, 0x26, 0xfa, 0xa6, 0x76,
	0x58, 0x97, 0x46, 0x84, 0x63, 0xf8, 0xdd, 0x83, 0x9b, 0x4d, 0xed, 0xf2, 0x6c, 0xfe, 0xcf, 0x9b,
	0xf9, 0x3d, 0x68, 0x67, 0x4c, 0x4c, 0x4f, 0x34, 0x4f, 0xe5, 0x84, 0x62, 0x67, 0x4b, 0x93, 0x5b,
	0x1a, 0x1e, 0x59, 0x34, 0xfc, 0xe1, 0xc1, 0xad, 0xf3, 0xae, 0xde, 0x20, 0x9d, 0x10, 0xc1, 0x14,
	0x5f, 0xe7, 0xec, 0x09, 0xb4, 0x33, 0x49, 0x49, 0x96, 0x70, 0x7b, 0x42, 0x9b, 0xba, 0x32, 0xd8,
	0xeb, 0x9b, 0x51, 0xbf, 0xae, 0xca, 0x4e, 0x2f, 0x6e, 0x65, 0x0b, 0xef, 0x4b, 0xa3, 0x6d, 0x2d,
	0x9d, 0xe3, 0x47, 0x38, 0x58, 0x61, 0x78, 0xa3, 0x79, 0xf6, 0xa0, 0xed, 0xfc, 0x26, 0x1f, 0xe4,
	0x54, 0xa4, 0xda, 0xf5, 0x76, 0xdc, 0x72, 0xf0, 0xf3, 0x0a, 0x0d, 0xbf, 0x79, 0xd0, 0xb1, 0xcd,
	0x9e, 0xa6, 0x33, 0x2c, 0x4a, 0xa6, 0x70, 0xc3, 0x0d, 0x79, 0x05, 0xbb, 0x36, 0x19, 0x1a, 0x09,
	0x33, 0x9b, 0x6e, 0xdf, 0xae, 0xf3, 0x68, 0x21, 0x9e, 0xed, 0x14, 0xb7, 0xcd, 0x41, 0x0b, 0x84,
	0x5f, 0x60, 0xaf, 0xc1, 0xc6, 0x06, 0x49, 0x03, 0xd8, 0x26, 0x94, 0x62, 0x5e, 0xa2, 0x8d, 0xe8,
	0xde, 0xab, 0x29, 0x90, 0x94, 0x33, 0xa5, 0x98, 0x14, 0x09, 0x16, 0x85, 0x2c, 0xec, 0x8e, 0x38,
	0x78, 0x58, 0xa1, 0xe1, 0x57, 0x0f, 0x6e, 0x38, 0x6f, 0xc8, 0xe5, 0xec, 0x22, 0x46, 0x30, 0x87,
	0x6b, 0x67, 0x3d, 0xfc, 0xa7, 0xfc, 0x83, 0x5f, 0x1e, 0xf8, 0xe7, 0x57, 0xce, 0x1f, 0xc3, 0xbe,
	0x99, 0x43, 0x43, 0xf1, 0x7e, 0xff, 0xd4, 0x55, 0xb8, 0xf4, 0x96, 0x09, 0xee, 0xae, 0xa3, 0x55,
	0x19, 0x4f, 0xa0, 0x7b, 0x88, 0x22, 0x5d, 0xf5, 0x99, 0xf6, 0x56, 0xeb, 0x38, 0x62, 0xf0, 0x70,
	0x43, 0xa2, 0xee, 0x3c, 0xf8, 0x79, 0x2a, 0xf9, 0xdb, 0x1c, 0x0b, 0x52, 0x71, 0x94, 0xff, 0x1e,
	0xae, 0xba, 0x35, 0xb4, 0x65, 0xff, 0x5e, 0x93, 0xf2, 0xd9, 0x8f, 0x26, 0x08, 0xd7, 0xb0, 0xaa,
	0xbc, 0xef, 0xa0, 0x55, 0xff, 0xc5, 0x4e, 0xfb, 0x4e, 0xd3, 0xa9, 0x85, 0x55, 0x0c, 0xba, 0xab,
	0x28, 0x79, 0x36, 0x3f, 0xba, 0xac, 0x2f, 0xfd, 0xc7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xce,
	0x65, 0x5c, 0xee, 0xa4, 0x06, 0x00, 0x00,
}