// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package application

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// 应用管理
	CreateApplication(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	UpdateApplication(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	QueryApplication(ctx context.Context, in *QueryApplicationRequest, opts ...grpc.CallOption) (*ApplicationSet, error)
	DescribeApplication(ctx context.Context, in *DescribeApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	DeleteApplication(ctx context.Context, in *DeleteApplicationRequest, opts ...grpc.CallOption) (*Application, error)
	// 应用事件处理
	HandleApplicationEvent(ctx context.Context, in *ApplicationEvent, opts ...grpc.CallOption) (*Application, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreateApplication(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/infraboard.workflow.application.Service/CreateApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateApplication(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/infraboard.workflow.application.Service/UpdateApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryApplication(ctx context.Context, in *QueryApplicationRequest, opts ...grpc.CallOption) (*ApplicationSet, error) {
	out := new(ApplicationSet)
	err := c.cc.Invoke(ctx, "/infraboard.workflow.application.Service/QueryApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DescribeApplication(ctx context.Context, in *DescribeApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/infraboard.workflow.application.Service/DescribeApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteApplication(ctx context.Context, in *DeleteApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/infraboard.workflow.application.Service/DeleteApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) HandleApplicationEvent(ctx context.Context, in *ApplicationEvent, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/infraboard.workflow.application.Service/HandleApplicationEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// 应用管理
	CreateApplication(context.Context, *CreateApplicationRequest) (*Application, error)
	UpdateApplication(context.Context, *UpdateApplicationRequest) (*Application, error)
	QueryApplication(context.Context, *QueryApplicationRequest) (*ApplicationSet, error)
	DescribeApplication(context.Context, *DescribeApplicationRequest) (*Application, error)
	DeleteApplication(context.Context, *DeleteApplicationRequest) (*Application, error)
	// 应用事件处理
	HandleApplicationEvent(context.Context, *ApplicationEvent) (*Application, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) CreateApplication(context.Context, *CreateApplicationRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApplication not implemented")
}
func (UnimplementedServiceServer) UpdateApplication(context.Context, *UpdateApplicationRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApplication not implemented")
}
func (UnimplementedServiceServer) QueryApplication(context.Context, *QueryApplicationRequest) (*ApplicationSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryApplication not implemented")
}
func (UnimplementedServiceServer) DescribeApplication(context.Context, *DescribeApplicationRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeApplication not implemented")
}
func (UnimplementedServiceServer) DeleteApplication(context.Context, *DeleteApplicationRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApplication not implemented")
}
func (UnimplementedServiceServer) HandleApplicationEvent(context.Context, *ApplicationEvent) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleApplicationEvent not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_CreateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.workflow.application.Service/CreateApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateApplication(ctx, req.(*CreateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.workflow.application.Service/UpdateApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateApplication(ctx, req.(*UpdateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.workflow.application.Service/QueryApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryApplication(ctx, req.(*QueryApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DescribeApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribeApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.workflow.application.Service/DescribeApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribeApplication(ctx, req.(*DescribeApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.workflow.application.Service/DeleteApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteApplication(ctx, req.(*DeleteApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_HandleApplicationEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).HandleApplicationEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.workflow.application.Service/HandleApplicationEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).HandleApplicationEvent(ctx, req.(*ApplicationEvent))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.workflow.application.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApplication",
			Handler:    _Service_CreateApplication_Handler,
		},
		{
			MethodName: "UpdateApplication",
			Handler:    _Service_UpdateApplication_Handler,
		},
		{
			MethodName: "QueryApplication",
			Handler:    _Service_QueryApplication_Handler,
		},
		{
			MethodName: "DescribeApplication",
			Handler:    _Service_DescribeApplication_Handler,
		},
		{
			MethodName: "DeleteApplication",
			Handler:    _Service_DeleteApplication_Handler,
		},
		{
			MethodName: "HandleApplicationEvent",
			Handler:    _Service_HandleApplicationEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/app/application/pb/application.proto",
}
