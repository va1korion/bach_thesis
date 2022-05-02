// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/worker.proto

package worker_api

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

// ImageTestClient is the client API for ImageTest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageTestClient interface {
	// Sends a greeting
	Analyse(ctx context.Context, opts ...grpc.CallOption) (ImageTest_AnalyseClient, error)
}

type imageTestClient struct {
	cc grpc.ClientConnInterface
}

func NewImageTestClient(cc grpc.ClientConnInterface) ImageTestClient {
	return &imageTestClient{cc}
}

func (c *imageTestClient) Analyse(ctx context.Context, opts ...grpc.CallOption) (ImageTest_AnalyseClient, error) {
	stream, err := c.cc.NewStream(ctx, &ImageTest_ServiceDesc.Streams[0], "/manager.ImageTest/Analyse", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageTestAnalyseClient{stream}
	return x, nil
}

type ImageTest_AnalyseClient interface {
	Send(*MsgRequest) error
	Recv() (*MsgReply, error)
	grpc.ClientStream
}

type imageTestAnalyseClient struct {
	grpc.ClientStream
}

func (x *imageTestAnalyseClient) Send(m *MsgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imageTestAnalyseClient) Recv() (*MsgReply, error) {
	m := new(MsgReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImageTestServer is the server API for ImageTest service.
// All implementations must embed UnimplementedImageTestServer
// for forward compatibility
type ImageTestServer interface {
	// Sends a greeting
	Analyse(ImageTest_AnalyseServer) error
	mustEmbedUnimplementedImageTestServer()
}

// UnimplementedImageTestServer must be embedded to have forward compatible implementations.
type UnimplementedImageTestServer struct {
}

func (UnimplementedImageTestServer) Analyse(ImageTest_AnalyseServer) error {
	return status.Errorf(codes.Unimplemented, "method Analyse not implemented")
}
func (UnimplementedImageTestServer) mustEmbedUnimplementedImageTestServer() {}

// UnsafeImageTestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageTestServer will
// result in compilation errors.
type UnsafeImageTestServer interface {
	mustEmbedUnimplementedImageTestServer()
}

func RegisterImageTestServer(s grpc.ServiceRegistrar, srv ImageTestServer) {
	s.RegisterService(&ImageTest_ServiceDesc, srv)
}

func _ImageTest_Analyse_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImageTestServer).Analyse(&imageTestAnalyseServer{stream})
}

type ImageTest_AnalyseServer interface {
	Send(*MsgReply) error
	Recv() (*MsgRequest, error)
	grpc.ServerStream
}

type imageTestAnalyseServer struct {
	grpc.ServerStream
}

func (x *imageTestAnalyseServer) Send(m *MsgReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imageTestAnalyseServer) Recv() (*MsgRequest, error) {
	m := new(MsgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImageTest_ServiceDesc is the grpc.ServiceDesc for ImageTest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageTest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager.ImageTest",
	HandlerType: (*ImageTestServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Analyse",
			Handler:       _ImageTest_Analyse_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/worker.proto",
}