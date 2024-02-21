// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: task/task.proto

package task

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

// TaskerClient is the client API for Tasker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskerClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type taskerClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskerClient(cc grpc.ClientConnInterface) TaskerClient {
	return &taskerClient{cc}
}

func (c *taskerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/taskgrpc.Tasker/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskerClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/taskgrpc.Tasker/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskerServer is the server API for Tasker service.
// All implementations must embed UnimplementedTaskerServer
// for forward compatibility
type TaskerServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	mustEmbedUnimplementedTaskerServer()
}

// UnimplementedTaskerServer must be embedded to have forward compatible implementations.
type UnimplementedTaskerServer struct {
}

func (UnimplementedTaskerServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTaskerServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTaskerServer) mustEmbedUnimplementedTaskerServer() {}

// UnsafeTaskerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskerServer will
// result in compilation errors.
type UnsafeTaskerServer interface {
	mustEmbedUnimplementedTaskerServer()
}

func RegisterTaskerServer(s grpc.ServiceRegistrar, srv TaskerServer) {
	s.RegisterService(&Tasker_ServiceDesc, srv)
}

func _Tasker_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taskgrpc.Tasker/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasker_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taskgrpc.Tasker/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskerServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tasker_ServiceDesc is the grpc.ServiceDesc for Tasker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tasker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "taskgrpc.Tasker",
	HandlerType: (*TaskerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Tasker_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Tasker_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task/task.proto",
}
