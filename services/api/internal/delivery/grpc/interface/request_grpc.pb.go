// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.2
// source: request.proto

package contact

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

// RequestServiceClient is the client API for RequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RequestServiceClient interface {
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	AddVideo(ctx context.Context, in *AddVideoRequest, opts ...grpc.CallOption) (*AddVideoResponse, error)
	ShowGroups(ctx context.Context, in *GroupsQuery, opts ...grpc.CallOption) (RequestService_ShowGroupsClient, error)
	RemoveVideo(ctx context.Context, in *RemoveVideoRequest, opts ...grpc.CallOption) (*RemoveVideoResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*CreateLoginResponse, error)
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error)
}

type requestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRequestServiceClient(cc grpc.ClientConnInterface) RequestServiceClient {
	return &requestServiceClient{cc}
}

func (c *requestServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, "/request.RequestService/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) AddVideo(ctx context.Context, in *AddVideoRequest, opts ...grpc.CallOption) (*AddVideoResponse, error) {
	out := new(AddVideoResponse)
	err := c.cc.Invoke(ctx, "/request.RequestService/AddVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) ShowGroups(ctx context.Context, in *GroupsQuery, opts ...grpc.CallOption) (RequestService_ShowGroupsClient, error) {
	stream, err := c.cc.NewStream(ctx, &RequestService_ServiceDesc.Streams[0], "/request.RequestService/ShowGroups", opts...)
	if err != nil {
		return nil, err
	}
	x := &requestServiceShowGroupsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RequestService_ShowGroupsClient interface {
	Recv() (*GroupsResponse, error)
	grpc.ClientStream
}

type requestServiceShowGroupsClient struct {
	grpc.ClientStream
}

func (x *requestServiceShowGroupsClient) Recv() (*GroupsResponse, error) {
	m := new(GroupsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *requestServiceClient) RemoveVideo(ctx context.Context, in *RemoveVideoRequest, opts ...grpc.CallOption) (*RemoveVideoResponse, error) {
	out := new(RemoveVideoResponse)
	err := c.cc.Invoke(ctx, "/request.RequestService/RemoveVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*CreateLoginResponse, error) {
	out := new(CreateLoginResponse)
	err := c.cc.Invoke(ctx, "/request.RequestService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error) {
	out := new(DeleteGroupResponse)
	err := c.cc.Invoke(ctx, "/request.RequestService/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RequestServiceServer is the server API for RequestService service.
// All implementations must embed UnimplementedRequestServiceServer
// for forward compatibility
type RequestServiceServer interface {
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	AddVideo(context.Context, *AddVideoRequest) (*AddVideoResponse, error)
	ShowGroups(*GroupsQuery, RequestService_ShowGroupsServer) error
	RemoveVideo(context.Context, *RemoveVideoRequest) (*RemoveVideoResponse, error)
	Login(context.Context, *LoginRequest) (*CreateLoginResponse, error)
	DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error)
	mustEmbedUnimplementedRequestServiceServer()
}

// UnimplementedRequestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRequestServiceServer struct {
}

func (UnimplementedRequestServiceServer) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedRequestServiceServer) AddVideo(context.Context, *AddVideoRequest) (*AddVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVideo not implemented")
}
func (UnimplementedRequestServiceServer) ShowGroups(*GroupsQuery, RequestService_ShowGroupsServer) error {
	return status.Errorf(codes.Unimplemented, "method ShowGroups not implemented")
}
func (UnimplementedRequestServiceServer) RemoveVideo(context.Context, *RemoveVideoRequest) (*RemoveVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveVideo not implemented")
}
func (UnimplementedRequestServiceServer) Login(context.Context, *LoginRequest) (*CreateLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedRequestServiceServer) DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedRequestServiceServer) mustEmbedUnimplementedRequestServiceServer() {}

// UnsafeRequestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RequestServiceServer will
// result in compilation errors.
type UnsafeRequestServiceServer interface {
	mustEmbedUnimplementedRequestServiceServer()
}

func RegisterRequestServiceServer(s grpc.ServiceRegistrar, srv RequestServiceServer) {
	s.RegisterService(&RequestService_ServiceDesc, srv)
}

func _RequestService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/request.RequestService/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_AddVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).AddVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/request.RequestService/AddVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).AddVideo(ctx, req.(*AddVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_ShowGroups_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GroupsQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RequestServiceServer).ShowGroups(m, &requestServiceShowGroupsServer{stream})
}

type RequestService_ShowGroupsServer interface {
	Send(*GroupsResponse) error
	grpc.ServerStream
}

type requestServiceShowGroupsServer struct {
	grpc.ServerStream
}

func (x *requestServiceShowGroupsServer) Send(m *GroupsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RequestService_RemoveVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).RemoveVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/request.RequestService/RemoveVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).RemoveVideo(ctx, req.(*RemoveVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/request.RequestService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/request.RequestService/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RequestService_ServiceDesc is the grpc.ServiceDesc for RequestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RequestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "request.RequestService",
	HandlerType: (*RequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _RequestService_CreateGroup_Handler,
		},
		{
			MethodName: "AddVideo",
			Handler:    _RequestService_AddVideo_Handler,
		},
		{
			MethodName: "RemoveVideo",
			Handler:    _RequestService_RemoveVideo_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _RequestService_Login_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _RequestService_DeleteGroup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ShowGroups",
			Handler:       _RequestService_ShowGroups_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "request.proto",
}