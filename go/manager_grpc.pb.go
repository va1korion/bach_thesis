// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: manager.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	// gets worker status
	GetStatus(ctx context.Context, in *Worker, opts ...grpc.CallOption) (*Status, error)
	// streams from client for nn purposes
	StreamTS(ctx context.Context, opts ...grpc.CallOption) (Manager_StreamTSClient, error)
	// streams from elsewhere, still gets video after nn
	// Source is either ip (of a camera) or dir
	GetVideo(ctx context.Context, in *Source, opts ...grpc.CallOption) (Manager_GetVideoClient, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) GetStatus(ctx context.Context, in *Worker, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/manager.Manager/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) StreamTS(ctx context.Context, opts ...grpc.CallOption) (Manager_StreamTSClient, error) {
	stream, err := c.cc.NewStream(ctx, &Manager_ServiceDesc.Streams[0], "/manager.Manager/StreamTS", opts...)
	if err != nil {
		return nil, err
	}
	x := &managerStreamTSClient{stream}
	return x, nil
}

type Manager_StreamTSClient interface {
	Send(*Video) error
	Recv() (*Result, error)
	grpc.ClientStream
}

type managerStreamTSClient struct {
	grpc.ClientStream
}

func (x *managerStreamTSClient) Send(m *Video) error {
	return x.ClientStream.SendMsg(m)
}

func (x *managerStreamTSClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *managerClient) GetVideo(ctx context.Context, in *Source, opts ...grpc.CallOption) (Manager_GetVideoClient, error) {
	stream, err := c.cc.NewStream(ctx, &Manager_ServiceDesc.Streams[1], "/manager.Manager/GetVideo", opts...)
	if err != nil {
		return nil, err
	}
	x := &managerGetVideoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Manager_GetVideoClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type managerGetVideoClient struct {
	grpc.ClientStream
}

func (x *managerGetVideoClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	// gets worker status
	GetStatus(context.Context, *Worker) (*Status, error)
	// streams from client for nn purposes
	StreamTS(Manager_StreamTSServer) error
	// streams from elsewhere, still gets video after nn
	// Source is either ip (of a camera) or dir
	GetVideo(*Source, Manager_GetVideoServer) error
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) GetStatus(context.Context, *Worker) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedManagerServer) StreamTS(Manager_StreamTSServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTS not implemented")
}
func (UnimplementedManagerServer) GetVideo(*Source, Manager_GetVideoServer) error {
	return status.Errorf(codes.Unimplemented, "method GetVideo not implemented")
}
func (UnimplementedManagerServer) mustEmbedUnimplementedManagerServer() {}

// UnsafeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServer will
// result in compilation errors.
type UnsafeManagerServer interface {
	mustEmbedUnimplementedManagerServer()
}

func RegisterManagerServer(s grpc.ServiceRegistrar, srv ManagerServer) {
	s.RegisterService(&Manager_ServiceDesc, srv)
}

func _Manager_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Worker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Manager/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetStatus(ctx, req.(*Worker))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_StreamTS_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ManagerServer).StreamTS(&managerStreamTSServer{stream})
}

type Manager_StreamTSServer interface {
	Send(*Result) error
	Recv() (*Video, error)
	grpc.ServerStream
}

type managerStreamTSServer struct {
	grpc.ServerStream
}

func (x *managerStreamTSServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func (x *managerStreamTSServer) Recv() (*Video, error) {
	m := new(Video)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Manager_GetVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Source)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManagerServer).GetVideo(m, &managerGetVideoServer{stream})
}

type Manager_GetVideoServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type managerGetVideoServer struct {
	grpc.ServerStream
}

func (x *managerGetVideoServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _Manager_GetStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTS",
			Handler:       _Manager_StreamTS_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetVideo",
			Handler:       _Manager_GetVideo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "manager.proto",
}
