// Code generated by protoc-gen-go. DO NOT EDIT.
// source: forwarder.proto

package forwarder

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	connection "github.com/networkservicemesh/networkservicemesh/controlplane/api/connection"
	crossconnect "github.com/networkservicemesh/networkservicemesh/controlplane/api/crossconnect"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Message sent by forwarder module informing NSM of any changes in its
// operations parameters or constraints
type MechanismUpdate struct {
	RemoteMechanisms     []*connection.Mechanism `protobuf:"bytes,1,rep,name=remote_mechanisms,json=remoteMechanisms,proto3" json:"remote_mechanisms,omitempty"`
	LocalMechanisms      []*connection.Mechanism `protobuf:"bytes,2,rep,name=local_mechanisms,json=localMechanisms,proto3" json:"local_mechanisms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MechanismUpdate) Reset()         { *m = MechanismUpdate{} }
func (m *MechanismUpdate) String() string { return proto.CompactTextString(m) }
func (*MechanismUpdate) ProtoMessage()    {}
func (*MechanismUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_19bff53f4d11db23, []int{0}
}

func (m *MechanismUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MechanismUpdate.Unmarshal(m, b)
}
func (m *MechanismUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MechanismUpdate.Marshal(b, m, deterministic)
}
func (m *MechanismUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MechanismUpdate.Merge(m, src)
}
func (m *MechanismUpdate) XXX_Size() int {
	return xxx_messageInfo_MechanismUpdate.Size(m)
}
func (m *MechanismUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_MechanismUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_MechanismUpdate proto.InternalMessageInfo

func (m *MechanismUpdate) GetRemoteMechanisms() []*connection.Mechanism {
	if m != nil {
		return m.RemoteMechanisms
	}
	return nil
}

func (m *MechanismUpdate) GetLocalMechanisms() []*connection.Mechanism {
	if m != nil {
		return m.LocalMechanisms
	}
	return nil
}

func init() {
	proto.RegisterType((*MechanismUpdate)(nil), "forwarder.MechanismUpdate")
}

func init() { proto.RegisterFile("forwarder.proto", fileDescriptor_19bff53f4d11db23) }

var fileDescriptor_19bff53f4d11db23 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x4e, 0x02, 0x31,
	0x10, 0xc7, 0xb3, 0x1a, 0x35, 0xd4, 0x03, 0xd0, 0x44, 0x43, 0x7a, 0x32, 0x9e, 0x3c, 0x15, 0x83,
	0x47, 0x2f, 0x2a, 0xd1, 0xc4, 0x03, 0x17, 0x12, 0x8f, 0x6a, 0x4a, 0x19, 0xa0, 0xb1, 0xed, 0xd4,
	0xb6, 0x48, 0x78, 0x06, 0x1f, 0xc0, 0xd7, 0x35, 0x4b, 0x97, 0xdd, 0xc5, 0xe8, 0x9e, 0xbc, 0x34,
	0x33, 0xff, 0x99, 0xfe, 0xe6, 0xa3, 0x25, 0xed, 0x19, 0xfa, 0x95, 0xf0, 0x53, 0xf0, 0xdc, 0x79,
	0x8c, 0x48, 0x5b, 0xa5, 0xc0, 0x9e, 0xe7, 0x2a, 0x2e, 0x96, 0x13, 0x2e, 0xd1, 0xf4, 0x2d, 0xc4,
	0x15, 0xfa, 0xb7, 0x00, 0xfe, 0x43, 0x49, 0x30, 0x10, 0x16, 0xbf, 0x49, 0x12, 0x6d, 0xf4, 0xa8,
	0x9d, 0x16, 0x16, 0xfa, 0xc2, 0xa9, 0x5c, 0xb0, 0x20, 0xa3, 0x42, 0x5b, 0x33, 0x53, 0x25, 0x26,
	0xfe, 0x01, 0xef, 0x31, 0x84, 0x02, 0xbc, 0xe3, 0x14, 0x25, 0x7a, 0x2e, 0xae, 0x1d, 0x84, 0x3e,
	0x18, 0x17, 0xd7, 0xe9, 0x4c, 0x91, 0xf3, 0xaf, 0x8c, 0xb4, 0x47, 0x20, 0x17, 0xc2, 0xaa, 0x60,
	0x9e, 0xdc, 0x54, 0x44, 0xa0, 0x77, 0xa4, 0xeb, 0xc1, 0x60, 0x84, 0x57, 0xb3, 0x8d, 0x84, 0x5e,
	0x76, 0xb6, 0x7f, 0x71, 0x3c, 0x38, 0xe1, 0xb5, 0xf6, 0xcb, 0x7b, 0xe3, 0x4e, 0xca, 0x2f, 0x85,
	0x40, 0x6f, 0x48, 0x47, 0xa3, 0x14, 0xba, 0x8e, 0xd8, 0x6b, 0x42, 0xb4, 0x37, 0xe9, 0x15, 0x61,
	0xf0, 0x99, 0x91, 0xd6, 0xc3, 0xf6, 0x0d, 0xe8, 0x2d, 0x39, 0x1a, 0xc3, 0xfb, 0x12, 0x42, 0xa4,
	0x8c, 0xef, 0x4c, 0x38, 0xcc, 0x9d, 0x61, 0x72, 0x58, 0x43, 0x8c, 0x5e, 0x93, 0x83, 0xa1, 0xc6,
	0x00, 0x8d, 0x80, 0x53, 0x3e, 0x47, 0x9c, 0x6b, 0x48, 0xeb, 0x99, 0x2c, 0x67, 0xfc, 0x3e, 0xdf,
	0xd6, 0xe0, 0x85, 0x74, 0xab, 0xde, 0x46, 0x68, 0x55, 0x44, 0x4f, 0x1f, 0x49, 0xb7, 0x30, 0x6b,
	0x93, 0xff, 0x41, 0x60, 0x8c, 0x57, 0x5f, 0xec, 0xc7, 0xc6, 0x2f, 0xb3, 0xc9, 0xe1, 0x26, 0xfb,
	0xea, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x88, 0x27, 0xe6, 0x88, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ForwarderClient is the client API for Forwarder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ForwarderClient interface {
	Request(ctx context.Context, in *crossconnect.CrossConnect, opts ...grpc.CallOption) (*crossconnect.CrossConnect, error)
	Close(ctx context.Context, in *crossconnect.CrossConnect, opts ...grpc.CallOption) (*empty.Empty, error)
}

type forwarderClient struct {
	cc grpc.ClientConnInterface
}

func NewForwarderClient(cc grpc.ClientConnInterface) ForwarderClient {
	return &forwarderClient{cc}
}

func (c *forwarderClient) Request(ctx context.Context, in *crossconnect.CrossConnect, opts ...grpc.CallOption) (*crossconnect.CrossConnect, error) {
	out := new(crossconnect.CrossConnect)
	err := c.cc.Invoke(ctx, "/forwarder.Forwarder/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forwarderClient) Close(ctx context.Context, in *crossconnect.CrossConnect, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/forwarder.Forwarder/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ForwarderServer is the server API for Forwarder service.
type ForwarderServer interface {
	Request(context.Context, *crossconnect.CrossConnect) (*crossconnect.CrossConnect, error)
	Close(context.Context, *crossconnect.CrossConnect) (*empty.Empty, error)
}

// UnimplementedForwarderServer can be embedded to have forward compatible implementations.
type UnimplementedForwarderServer struct {
}

func (*UnimplementedForwarderServer) Request(ctx context.Context, req *crossconnect.CrossConnect) (*crossconnect.CrossConnect, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (*UnimplementedForwarderServer) Close(ctx context.Context, req *crossconnect.CrossConnect) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}

func RegisterForwarderServer(s *grpc.Server, srv ForwarderServer) {
	s.RegisterService(&_Forwarder_serviceDesc, srv)
}

func _Forwarder_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(crossconnect.CrossConnect)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForwarderServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/forwarder.Forwarder/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForwarderServer).Request(ctx, req.(*crossconnect.CrossConnect))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forwarder_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(crossconnect.CrossConnect)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForwarderServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/forwarder.Forwarder/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForwarderServer).Close(ctx, req.(*crossconnect.CrossConnect))
	}
	return interceptor(ctx, in, info, handler)
}

var _Forwarder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "forwarder.Forwarder",
	HandlerType: (*ForwarderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _Forwarder_Request_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Forwarder_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "forwarder.proto",
}

// MechanismsMonitorClient is the client API for MechanismsMonitor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MechanismsMonitorClient interface {
	MonitorMechanisms(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (MechanismsMonitor_MonitorMechanismsClient, error)
}

type mechanismsMonitorClient struct {
	cc grpc.ClientConnInterface
}

func NewMechanismsMonitorClient(cc grpc.ClientConnInterface) MechanismsMonitorClient {
	return &mechanismsMonitorClient{cc}
}

func (c *mechanismsMonitorClient) MonitorMechanisms(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (MechanismsMonitor_MonitorMechanismsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MechanismsMonitor_serviceDesc.Streams[0], "/forwarder.MechanismsMonitor/MonitorMechanisms", opts...)
	if err != nil {
		return nil, err
	}
	x := &mechanismsMonitorMonitorMechanismsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MechanismsMonitor_MonitorMechanismsClient interface {
	Recv() (*MechanismUpdate, error)
	grpc.ClientStream
}

type mechanismsMonitorMonitorMechanismsClient struct {
	grpc.ClientStream
}

func (x *mechanismsMonitorMonitorMechanismsClient) Recv() (*MechanismUpdate, error) {
	m := new(MechanismUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MechanismsMonitorServer is the server API for MechanismsMonitor service.
type MechanismsMonitorServer interface {
	MonitorMechanisms(*empty.Empty, MechanismsMonitor_MonitorMechanismsServer) error
}

// UnimplementedMechanismsMonitorServer can be embedded to have forward compatible implementations.
type UnimplementedMechanismsMonitorServer struct {
}

func (*UnimplementedMechanismsMonitorServer) MonitorMechanisms(req *empty.Empty, srv MechanismsMonitor_MonitorMechanismsServer) error {
	return status.Errorf(codes.Unimplemented, "method MonitorMechanisms not implemented")
}

func RegisterMechanismsMonitorServer(s *grpc.Server, srv MechanismsMonitorServer) {
	s.RegisterService(&_MechanismsMonitor_serviceDesc, srv)
}

func _MechanismsMonitor_MonitorMechanisms_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MechanismsMonitorServer).MonitorMechanisms(m, &mechanismsMonitorMonitorMechanismsServer{stream})
}

type MechanismsMonitor_MonitorMechanismsServer interface {
	Send(*MechanismUpdate) error
	grpc.ServerStream
}

type mechanismsMonitorMonitorMechanismsServer struct {
	grpc.ServerStream
}

func (x *mechanismsMonitorMonitorMechanismsServer) Send(m *MechanismUpdate) error {
	return x.ServerStream.SendMsg(m)
}

var _MechanismsMonitor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "forwarder.MechanismsMonitor",
	HandlerType: (*MechanismsMonitorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorMechanisms",
			Handler:       _MechanismsMonitor_MonitorMechanisms_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "forwarder.proto",
}
