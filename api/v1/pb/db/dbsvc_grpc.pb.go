// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package db

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

// DatabaseClient is the client API for Database service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error)
	ServiceStatus(ctx context.Context, in *ServiceStatusRequest, opts ...grpc.CallOption) (*ServiceStatusReply, error)
}

type databaseClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseClient(cc grpc.ClientConnInterface) DatabaseClient {
	return &databaseClient{cc}
}

func (c *databaseClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := c.cc.Invoke(ctx, "/pb.database/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/pb.database/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error) {
	out := new(UpdateReply)
	err := c.cc.Invoke(ctx, "/pb.database/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error) {
	out := new(RemoveReply)
	err := c.cc.Invoke(ctx, "/pb.database/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) ServiceStatus(ctx context.Context, in *ServiceStatusRequest, opts ...grpc.CallOption) (*ServiceStatusReply, error) {
	out := new(ServiceStatusReply)
	err := c.cc.Invoke(ctx, "/pb.database/ServiceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServer is the server API for Database service.
// All implementations must embed UnimplementedDatabaseServer
// for forward compatibility
type DatabaseServer interface {
	Add(context.Context, *AddRequest) (*AddReply, error)
	Get(context.Context, *GetRequest) (*GetReply, error)
	Update(context.Context, *UpdateRequest) (*UpdateReply, error)
	Remove(context.Context, *RemoveRequest) (*RemoveReply, error)
	ServiceStatus(context.Context, *ServiceStatusRequest) (*ServiceStatusReply, error)
	mustEmbedUnimplementedDatabaseServer()
}

// UnimplementedDatabaseServer must be embedded to have forward compatible implementations.
type UnimplementedDatabaseServer struct {
}

func (UnimplementedDatabaseServer) Add(context.Context, *AddRequest) (*AddReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedDatabaseServer) Get(context.Context, *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDatabaseServer) Update(context.Context, *UpdateRequest) (*UpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDatabaseServer) Remove(context.Context, *RemoveRequest) (*RemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedDatabaseServer) ServiceStatus(context.Context, *ServiceStatusRequest) (*ServiceStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceStatus not implemented")
}
func (UnimplementedDatabaseServer) mustEmbedUnimplementedDatabaseServer() {}

// UnsafeDatabaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseServer will
// result in compilation errors.
type UnsafeDatabaseServer interface {
	mustEmbedUnimplementedDatabaseServer()
}

func RegisterDatabaseServer(s grpc.ServiceRegistrar, srv DatabaseServer) {
	s.RegisterService(&Database_ServiceDesc, srv)
}

func _Database_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.database/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.database/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.database/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.database/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_ServiceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).ServiceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.database/ServiceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).ServiceStatus(ctx, req.(*ServiceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Database_ServiceDesc is the grpc.ServiceDesc for Database service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Database_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.database",
	HandlerType: (*DatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Database_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Database_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Database_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Database_Remove_Handler,
		},
		{
			MethodName: "ServiceStatus",
			Handler:    _Database_ServiceStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/pb/db/dbsvc.proto",
}
