// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package client_main

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

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	UnaryComms(ctx context.Context, in *UnaryNormalMessage, opts ...grpc.CallOption) (*UnaryServerMessage, error)
	ServerStreamComms(ctx context.Context, in *StreamNormalMessage, opts ...grpc.CallOption) (ChatService_ServerStreamCommsClient, error)
	ClientStreamComms(ctx context.Context, opts ...grpc.CallOption) (ChatService_ClientStreamCommsClient, error)
	BiDirectionStreamComms(ctx context.Context, opts ...grpc.CallOption) (ChatService_BiDirectionStreamCommsClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) UnaryComms(ctx context.Context, in *UnaryNormalMessage, opts ...grpc.CallOption) (*UnaryServerMessage, error) {
	out := new(UnaryServerMessage)
	err := c.cc.Invoke(ctx, "/main.ChatService/unaryComms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) ServerStreamComms(ctx context.Context, in *StreamNormalMessage, opts ...grpc.CallOption) (ChatService_ServerStreamCommsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/main.ChatService/serverStreamComms", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceServerStreamCommsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_ServerStreamCommsClient interface {
	Recv() (*StreamServerMessage, error)
	grpc.ClientStream
}

type chatServiceServerStreamCommsClient struct {
	grpc.ClientStream
}

func (x *chatServiceServerStreamCommsClient) Recv() (*StreamServerMessage, error) {
	m := new(StreamServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) ClientStreamComms(ctx context.Context, opts ...grpc.CallOption) (ChatService_ClientStreamCommsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[1], "/main.ChatService/clientStreamComms", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceClientStreamCommsClient{stream}
	return x, nil
}

type ChatService_ClientStreamCommsClient interface {
	Send(*StreamNormalMessage) error
	CloseAndRecv() (*StreamServerMessage, error)
	grpc.ClientStream
}

type chatServiceClientStreamCommsClient struct {
	grpc.ClientStream
}

func (x *chatServiceClientStreamCommsClient) Send(m *StreamNormalMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceClientStreamCommsClient) CloseAndRecv() (*StreamServerMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) BiDirectionStreamComms(ctx context.Context, opts ...grpc.CallOption) (ChatService_BiDirectionStreamCommsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[2], "/main.ChatService/biDirectionStreamComms", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceBiDirectionStreamCommsClient{stream}
	return x, nil
}

type ChatService_BiDirectionStreamCommsClient interface {
	Send(*StreamNormalMessage) error
	Recv() (*StreamServerMessage, error)
	grpc.ClientStream
}

type chatServiceBiDirectionStreamCommsClient struct {
	grpc.ClientStream
}

func (x *chatServiceBiDirectionStreamCommsClient) Send(m *StreamNormalMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceBiDirectionStreamCommsClient) Recv() (*StreamServerMessage, error) {
	m := new(StreamServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	UnaryComms(context.Context, *UnaryNormalMessage) (*UnaryServerMessage, error)
	ServerStreamComms(*StreamNormalMessage, ChatService_ServerStreamCommsServer) error
	ClientStreamComms(ChatService_ClientStreamCommsServer) error
	BiDirectionStreamComms(ChatService_BiDirectionStreamCommsServer) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) UnaryComms(context.Context, *UnaryNormalMessage) (*UnaryServerMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryComms not implemented")
}
func (UnimplementedChatServiceServer) ServerStreamComms(*StreamNormalMessage, ChatService_ServerStreamCommsServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamComms not implemented")
}
func (UnimplementedChatServiceServer) ClientStreamComms(ChatService_ClientStreamCommsServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamComms not implemented")
}
func (UnimplementedChatServiceServer) BiDirectionStreamComms(ChatService_BiDirectionStreamCommsServer) error {
	return status.Errorf(codes.Unimplemented, "method BiDirectionStreamComms not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_UnaryComms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryNormalMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UnaryComms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ChatService/unaryComms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UnaryComms(ctx, req.(*UnaryNormalMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_ServerStreamComms_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamNormalMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).ServerStreamComms(m, &chatServiceServerStreamCommsServer{stream})
}

type ChatService_ServerStreamCommsServer interface {
	Send(*StreamServerMessage) error
	grpc.ServerStream
}

type chatServiceServerStreamCommsServer struct {
	grpc.ServerStream
}

func (x *chatServiceServerStreamCommsServer) Send(m *StreamServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_ClientStreamComms_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).ClientStreamComms(&chatServiceClientStreamCommsServer{stream})
}

type ChatService_ClientStreamCommsServer interface {
	SendAndClose(*StreamServerMessage) error
	Recv() (*StreamNormalMessage, error)
	grpc.ServerStream
}

type chatServiceClientStreamCommsServer struct {
	grpc.ServerStream
}

func (x *chatServiceClientStreamCommsServer) SendAndClose(m *StreamServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceClientStreamCommsServer) Recv() (*StreamNormalMessage, error) {
	m := new(StreamNormalMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_BiDirectionStreamComms_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).BiDirectionStreamComms(&chatServiceBiDirectionStreamCommsServer{stream})
}

type ChatService_BiDirectionStreamCommsServer interface {
	Send(*StreamServerMessage) error
	Recv() (*StreamNormalMessage, error)
	grpc.ServerStream
}

type chatServiceBiDirectionStreamCommsServer struct {
	grpc.ServerStream
}

func (x *chatServiceBiDirectionStreamCommsServer) Send(m *StreamServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceBiDirectionStreamCommsServer) Recv() (*StreamNormalMessage, error) {
	m := new(StreamNormalMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "unaryComms",
			Handler:    _ChatService_UnaryComms_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "serverStreamComms",
			Handler:       _ChatService_ServerStreamComms_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "clientStreamComms",
			Handler:       _ChatService_ClientStreamComms_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "biDirectionStreamComms",
			Handler:       _ChatService_BiDirectionStreamComms_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}

// ChatServiceTwoClient is the client API for ChatServiceTwo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceTwoClient interface {
	UnaryComms2(ctx context.Context, in *UnaryNormalMessage, opts ...grpc.CallOption) (*UnaryServerMessage, error)
	ServerStreamComms2(ctx context.Context, in *StreamNormalMessage, opts ...grpc.CallOption) (ChatServiceTwo_ServerStreamComms2Client, error)
	ClientStreamComms2(ctx context.Context, opts ...grpc.CallOption) (ChatServiceTwo_ClientStreamComms2Client, error)
	BiDirectionStreamComms2(ctx context.Context, opts ...grpc.CallOption) (ChatServiceTwo_BiDirectionStreamComms2Client, error)
}

type chatServiceTwoClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceTwoClient(cc grpc.ClientConnInterface) ChatServiceTwoClient {
	return &chatServiceTwoClient{cc}
}

func (c *chatServiceTwoClient) UnaryComms2(ctx context.Context, in *UnaryNormalMessage, opts ...grpc.CallOption) (*UnaryServerMessage, error) {
	out := new(UnaryServerMessage)
	err := c.cc.Invoke(ctx, "/main.ChatServiceTwo/unaryComms2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceTwoClient) ServerStreamComms2(ctx context.Context, in *StreamNormalMessage, opts ...grpc.CallOption) (ChatServiceTwo_ServerStreamComms2Client, error) {
	stream, err := c.cc.NewStream(ctx, &ChatServiceTwo_ServiceDesc.Streams[0], "/main.ChatServiceTwo/serverStreamComms2", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceTwoServerStreamComms2Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatServiceTwo_ServerStreamComms2Client interface {
	Recv() (*StreamServerMessage, error)
	grpc.ClientStream
}

type chatServiceTwoServerStreamComms2Client struct {
	grpc.ClientStream
}

func (x *chatServiceTwoServerStreamComms2Client) Recv() (*StreamServerMessage, error) {
	m := new(StreamServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceTwoClient) ClientStreamComms2(ctx context.Context, opts ...grpc.CallOption) (ChatServiceTwo_ClientStreamComms2Client, error) {
	stream, err := c.cc.NewStream(ctx, &ChatServiceTwo_ServiceDesc.Streams[1], "/main.ChatServiceTwo/clientStreamComms2", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceTwoClientStreamComms2Client{stream}
	return x, nil
}

type ChatServiceTwo_ClientStreamComms2Client interface {
	Send(*StreamNormalMessage) error
	CloseAndRecv() (*StreamServerMessage, error)
	grpc.ClientStream
}

type chatServiceTwoClientStreamComms2Client struct {
	grpc.ClientStream
}

func (x *chatServiceTwoClientStreamComms2Client) Send(m *StreamNormalMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceTwoClientStreamComms2Client) CloseAndRecv() (*StreamServerMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceTwoClient) BiDirectionStreamComms2(ctx context.Context, opts ...grpc.CallOption) (ChatServiceTwo_BiDirectionStreamComms2Client, error) {
	stream, err := c.cc.NewStream(ctx, &ChatServiceTwo_ServiceDesc.Streams[2], "/main.ChatServiceTwo/biDirectionStreamComms2", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceTwoBiDirectionStreamComms2Client{stream}
	return x, nil
}

type ChatServiceTwo_BiDirectionStreamComms2Client interface {
	Send(*StreamNormalMessage) error
	Recv() (*StreamServerMessage, error)
	grpc.ClientStream
}

type chatServiceTwoBiDirectionStreamComms2Client struct {
	grpc.ClientStream
}

func (x *chatServiceTwoBiDirectionStreamComms2Client) Send(m *StreamNormalMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceTwoBiDirectionStreamComms2Client) Recv() (*StreamServerMessage, error) {
	m := new(StreamServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceTwoServer is the server API for ChatServiceTwo service.
// All implementations must embed UnimplementedChatServiceTwoServer
// for forward compatibility
type ChatServiceTwoServer interface {
	UnaryComms2(context.Context, *UnaryNormalMessage) (*UnaryServerMessage, error)
	ServerStreamComms2(*StreamNormalMessage, ChatServiceTwo_ServerStreamComms2Server) error
	ClientStreamComms2(ChatServiceTwo_ClientStreamComms2Server) error
	BiDirectionStreamComms2(ChatServiceTwo_BiDirectionStreamComms2Server) error
	mustEmbedUnimplementedChatServiceTwoServer()
}

// UnimplementedChatServiceTwoServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceTwoServer struct {
}

func (UnimplementedChatServiceTwoServer) UnaryComms2(context.Context, *UnaryNormalMessage) (*UnaryServerMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryComms2 not implemented")
}
func (UnimplementedChatServiceTwoServer) ServerStreamComms2(*StreamNormalMessage, ChatServiceTwo_ServerStreamComms2Server) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamComms2 not implemented")
}
func (UnimplementedChatServiceTwoServer) ClientStreamComms2(ChatServiceTwo_ClientStreamComms2Server) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamComms2 not implemented")
}
func (UnimplementedChatServiceTwoServer) BiDirectionStreamComms2(ChatServiceTwo_BiDirectionStreamComms2Server) error {
	return status.Errorf(codes.Unimplemented, "method BiDirectionStreamComms2 not implemented")
}
func (UnimplementedChatServiceTwoServer) mustEmbedUnimplementedChatServiceTwoServer() {}

// UnsafeChatServiceTwoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceTwoServer will
// result in compilation errors.
type UnsafeChatServiceTwoServer interface {
	mustEmbedUnimplementedChatServiceTwoServer()
}

func RegisterChatServiceTwoServer(s grpc.ServiceRegistrar, srv ChatServiceTwoServer) {
	s.RegisterService(&ChatServiceTwo_ServiceDesc, srv)
}

func _ChatServiceTwo_UnaryComms2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryNormalMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceTwoServer).UnaryComms2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ChatServiceTwo/unaryComms2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceTwoServer).UnaryComms2(ctx, req.(*UnaryNormalMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatServiceTwo_ServerStreamComms2_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamNormalMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceTwoServer).ServerStreamComms2(m, &chatServiceTwoServerStreamComms2Server{stream})
}

type ChatServiceTwo_ServerStreamComms2Server interface {
	Send(*StreamServerMessage) error
	grpc.ServerStream
}

type chatServiceTwoServerStreamComms2Server struct {
	grpc.ServerStream
}

func (x *chatServiceTwoServerStreamComms2Server) Send(m *StreamServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatServiceTwo_ClientStreamComms2_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceTwoServer).ClientStreamComms2(&chatServiceTwoClientStreamComms2Server{stream})
}

type ChatServiceTwo_ClientStreamComms2Server interface {
	SendAndClose(*StreamServerMessage) error
	Recv() (*StreamNormalMessage, error)
	grpc.ServerStream
}

type chatServiceTwoClientStreamComms2Server struct {
	grpc.ServerStream
}

func (x *chatServiceTwoClientStreamComms2Server) SendAndClose(m *StreamServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceTwoClientStreamComms2Server) Recv() (*StreamNormalMessage, error) {
	m := new(StreamNormalMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatServiceTwo_BiDirectionStreamComms2_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceTwoServer).BiDirectionStreamComms2(&chatServiceTwoBiDirectionStreamComms2Server{stream})
}

type ChatServiceTwo_BiDirectionStreamComms2Server interface {
	Send(*StreamServerMessage) error
	Recv() (*StreamNormalMessage, error)
	grpc.ServerStream
}

type chatServiceTwoBiDirectionStreamComms2Server struct {
	grpc.ServerStream
}

func (x *chatServiceTwoBiDirectionStreamComms2Server) Send(m *StreamServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceTwoBiDirectionStreamComms2Server) Recv() (*StreamNormalMessage, error) {
	m := new(StreamNormalMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceTwo_ServiceDesc is the grpc.ServiceDesc for ChatServiceTwo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatServiceTwo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.ChatServiceTwo",
	HandlerType: (*ChatServiceTwoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "unaryComms2",
			Handler:    _ChatServiceTwo_UnaryComms2_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "serverStreamComms2",
			Handler:       _ChatServiceTwo_ServerStreamComms2_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "clientStreamComms2",
			Handler:       _ChatServiceTwo_ClientStreamComms2_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "biDirectionStreamComms2",
			Handler:       _ChatServiceTwo_BiDirectionStreamComms2_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}
