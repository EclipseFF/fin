// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: admin.proto

package admin

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

// AdminPanelServiceClient is the client API for AdminPanelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminPanelServiceClient interface {
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	GetVideos(ctx context.Context, in *GetVideosRequest, opts ...grpc.CallOption) (*GetVideosResponse, error)
}

type adminPanelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminPanelServiceClient(cc grpc.ClientConnInterface) AdminPanelServiceClient {
	return &adminPanelServiceClient{cc}
}

func (c *adminPanelServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminPanelService/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminPanelServiceClient) GetVideos(ctx context.Context, in *GetVideosRequest, opts ...grpc.CallOption) (*GetVideosResponse, error) {
	out := new(GetVideosResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminPanelService/GetVideos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminPanelServiceServer is the server API for AdminPanelService service.
// All implementations must embed UnimplementedAdminPanelServiceServer
// for forward compatibility
type AdminPanelServiceServer interface {
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	GetVideos(context.Context, *GetVideosRequest) (*GetVideosResponse, error)
	mustEmbedUnimplementedAdminPanelServiceServer()
}

// UnimplementedAdminPanelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminPanelServiceServer struct {
}

func (UnimplementedAdminPanelServiceServer) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedAdminPanelServiceServer) GetVideos(context.Context, *GetVideosRequest) (*GetVideosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideos not implemented")
}
func (UnimplementedAdminPanelServiceServer) mustEmbedUnimplementedAdminPanelServiceServer() {}

// UnsafeAdminPanelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminPanelServiceServer will
// result in compilation errors.
type UnsafeAdminPanelServiceServer interface {
	mustEmbedUnimplementedAdminPanelServiceServer()
}

func RegisterAdminPanelServiceServer(s grpc.ServiceRegistrar, srv AdminPanelServiceServer) {
	s.RegisterService(&AdminPanelService_ServiceDesc, srv)
}

func _AdminPanelService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminPanelServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminPanelService/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminPanelServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminPanelService_GetVideos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminPanelServiceServer).GetVideos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminPanelService/GetVideos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminPanelServiceServer).GetVideos(ctx, req.(*GetVideosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminPanelService_ServiceDesc is the grpc.ServiceDesc for AdminPanelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminPanelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.AdminPanelService",
	HandlerType: (*AdminPanelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _AdminPanelService_CreateGroup_Handler,
		},
		{
			MethodName: "GetVideos",
			Handler:    _AdminPanelService_GetVideos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}
