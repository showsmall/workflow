// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pipeline

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
	CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*Pipeline, error)
	QueryPipeline(ctx context.Context, in *QueryPipelineRequest, opts ...grpc.CallOption) (*PipelineSet, error)
	DescribePipeline(ctx context.Context, in *DescribePipelineRequest, opts ...grpc.CallOption) (*Pipeline, error)
	CreateStep(ctx context.Context, in *CreateStepRequest, opts ...grpc.CallOption) (*Step, error)
	QueryStep(ctx context.Context, in *QueryStepRequest, opts ...grpc.CallOption) (*StepSet, error)
	DescribeStep(ctx context.Context, in *DescribeStepRequest, opts ...grpc.CallOption) (*Step, error)
	DeleteStep(ctx context.Context, in *DeleteStepRequest, opts ...grpc.CallOption) (*Step, error)
	CancelStep(ctx context.Context, in *CancelStepRequest, opts ...grpc.CallOption) (*Step, error)
	AuditStep(ctx context.Context, in *AuditStepRequest, opts ...grpc.CallOption) (*Step, error)
	DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*Pipeline, error)
	CreateAction(ctx context.Context, in *CreateActionRequest, opts ...grpc.CallOption) (*Action, error)
	QueryAction(ctx context.Context, in *QueryActionRequest, opts ...grpc.CallOption) (*ActionSet, error)
	DescribeAction(ctx context.Context, in *DescribeActionRequest, opts ...grpc.CallOption) (*Action, error)
	DeleteAction(ctx context.Context, in *DeleteActionRequest, opts ...grpc.CallOption) (*Action, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*Pipeline, error) {
	out := new(Pipeline)
	err := c.cc.Invoke(ctx, "/workflow.pipeline.Service/CreatePipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryPipeline(ctx context.Context, in *QueryPipelineRequest, opts ...grpc.CallOption) (*PipelineSet, error) {
	out := new(PipelineSet)
	err := c.cc.Invoke(ctx, "/workflow.pipeline.Service/QueryPipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DescribePipeline(ctx context.Context, in *DescribePipelineRequest, opts ...grpc.CallOption) (*Pipeline, error) {
	out := new(Pipeline)
	err := c.cc.Invoke(ctx, "/workflow.pipeline.Service/DescribePipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateStep(ctx context.Context, in *CreateStepRequest, opts ...grpc.CallOption) (*Step, error) {
	out := new(Step)
	err := c.cc.Invoke(ctx, "/workflow.pipeline.Service/CreateStep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryStep(ctx context.Context, in *QueryStepRequest, opts ...grpc.CallOption) (*StepSet, error) {
	out := new(StepSet)
	err := c.cc.Invoke(ctx, "/workflow.pipeline.Service/QueryStep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DescribeStep(ctx context.Context, in *DescribeStepRequest, opts ...grpc.CallOption) (*Step, error) {
	out := new(Step)
	err := c.cc.Invoke(ctx, "/workflow.pipeline.Service/DescribeStep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteStep(ctx context.Context, in *DeleteStepRequest, opts ...grpc.CallOption) (*Step, error) {
	out := new(Step)
	err := c.cc.Invoke(ctx, "/workflow.pipeline.Service/DeleteStep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CancelStep(ctx context.Context, in *CancelStepRequest, opts ...grpc.CallOption) (*Step, error) {
	out := new(Step)
	err := c.cc.Invoke(ctx, "/workflow.pipeline.Service/CancelStep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AuditStep(ctx context.Context, in *AuditStepRequest, opts ...grpc.CallOption) (*Step, error) {
	out := new(Step)
	err := c.cc.Invoke(ctx, "/workflow.pipeline.Service/AuditStep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*Pipeline, error) {
	out := new(Pipeline)
	err := c.cc.Invoke(ctx, "/workflow.pipeline.Service/DeletePipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateAction(ctx context.Context, in *CreateActionRequest, opts ...grpc.CallOption) (*Action, error) {
	out := new(Action)
	err := c.cc.Invoke(ctx, "/workflow.pipeline.Service/CreateAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryAction(ctx context.Context, in *QueryActionRequest, opts ...grpc.CallOption) (*ActionSet, error) {
	out := new(ActionSet)
	err := c.cc.Invoke(ctx, "/workflow.pipeline.Service/QueryAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DescribeAction(ctx context.Context, in *DescribeActionRequest, opts ...grpc.CallOption) (*Action, error) {
	out := new(Action)
	err := c.cc.Invoke(ctx, "/workflow.pipeline.Service/DescribeAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteAction(ctx context.Context, in *DeleteActionRequest, opts ...grpc.CallOption) (*Action, error) {
	out := new(Action)
	err := c.cc.Invoke(ctx, "/workflow.pipeline.Service/DeleteAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	CreatePipeline(context.Context, *CreatePipelineRequest) (*Pipeline, error)
	QueryPipeline(context.Context, *QueryPipelineRequest) (*PipelineSet, error)
	DescribePipeline(context.Context, *DescribePipelineRequest) (*Pipeline, error)
	CreateStep(context.Context, *CreateStepRequest) (*Step, error)
	QueryStep(context.Context, *QueryStepRequest) (*StepSet, error)
	DescribeStep(context.Context, *DescribeStepRequest) (*Step, error)
	DeleteStep(context.Context, *DeleteStepRequest) (*Step, error)
	CancelStep(context.Context, *CancelStepRequest) (*Step, error)
	AuditStep(context.Context, *AuditStepRequest) (*Step, error)
	DeletePipeline(context.Context, *DeletePipelineRequest) (*Pipeline, error)
	CreateAction(context.Context, *CreateActionRequest) (*Action, error)
	QueryAction(context.Context, *QueryActionRequest) (*ActionSet, error)
	DescribeAction(context.Context, *DescribeActionRequest) (*Action, error)
	DeleteAction(context.Context, *DeleteActionRequest) (*Action, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) CreatePipeline(context.Context, *CreatePipelineRequest) (*Pipeline, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePipeline not implemented")
}
func (UnimplementedServiceServer) QueryPipeline(context.Context, *QueryPipelineRequest) (*PipelineSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPipeline not implemented")
}
func (UnimplementedServiceServer) DescribePipeline(context.Context, *DescribePipelineRequest) (*Pipeline, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePipeline not implemented")
}
func (UnimplementedServiceServer) CreateStep(context.Context, *CreateStepRequest) (*Step, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStep not implemented")
}
func (UnimplementedServiceServer) QueryStep(context.Context, *QueryStepRequest) (*StepSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryStep not implemented")
}
func (UnimplementedServiceServer) DescribeStep(context.Context, *DescribeStepRequest) (*Step, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeStep not implemented")
}
func (UnimplementedServiceServer) DeleteStep(context.Context, *DeleteStepRequest) (*Step, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStep not implemented")
}
func (UnimplementedServiceServer) CancelStep(context.Context, *CancelStepRequest) (*Step, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelStep not implemented")
}
func (UnimplementedServiceServer) AuditStep(context.Context, *AuditStepRequest) (*Step, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuditStep not implemented")
}
func (UnimplementedServiceServer) DeletePipeline(context.Context, *DeletePipelineRequest) (*Pipeline, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePipeline not implemented")
}
func (UnimplementedServiceServer) CreateAction(context.Context, *CreateActionRequest) (*Action, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAction not implemented")
}
func (UnimplementedServiceServer) QueryAction(context.Context, *QueryActionRequest) (*ActionSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAction not implemented")
}
func (UnimplementedServiceServer) DescribeAction(context.Context, *DescribeActionRequest) (*Action, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeAction not implemented")
}
func (UnimplementedServiceServer) DeleteAction(context.Context, *DeleteActionRequest) (*Action, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAction not implemented")
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

func _Service_CreatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.pipeline.Service/CreatePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreatePipeline(ctx, req.(*CreatePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.pipeline.Service/QueryPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryPipeline(ctx, req.(*QueryPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DescribePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.pipeline.Service/DescribePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribePipeline(ctx, req.(*DescribePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateStep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateStep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.pipeline.Service/CreateStep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateStep(ctx, req.(*CreateStepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryStep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryStep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.pipeline.Service/QueryStep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryStep(ctx, req.(*QueryStepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DescribeStep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeStepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribeStep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.pipeline.Service/DescribeStep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribeStep(ctx, req.(*DescribeStepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteStep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteStep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.pipeline.Service/DeleteStep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteStep(ctx, req.(*DeleteStepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CancelStep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelStepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CancelStep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.pipeline.Service/CancelStep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CancelStep(ctx, req.(*CancelStepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_AuditStep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditStepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AuditStep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.pipeline.Service/AuditStep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AuditStep(ctx, req.(*AuditStepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeletePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeletePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.pipeline.Service/DeletePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeletePipeline(ctx, req.(*DeletePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.pipeline.Service/CreateAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateAction(ctx, req.(*CreateActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.pipeline.Service/QueryAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryAction(ctx, req.(*QueryActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DescribeAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribeAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.pipeline.Service/DescribeAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribeAction(ctx, req.(*DescribeActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.pipeline.Service/DeleteAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteAction(ctx, req.(*DeleteActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "workflow.pipeline.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePipeline",
			Handler:    _Service_CreatePipeline_Handler,
		},
		{
			MethodName: "QueryPipeline",
			Handler:    _Service_QueryPipeline_Handler,
		},
		{
			MethodName: "DescribePipeline",
			Handler:    _Service_DescribePipeline_Handler,
		},
		{
			MethodName: "CreateStep",
			Handler:    _Service_CreateStep_Handler,
		},
		{
			MethodName: "QueryStep",
			Handler:    _Service_QueryStep_Handler,
		},
		{
			MethodName: "DescribeStep",
			Handler:    _Service_DescribeStep_Handler,
		},
		{
			MethodName: "DeleteStep",
			Handler:    _Service_DeleteStep_Handler,
		},
		{
			MethodName: "CancelStep",
			Handler:    _Service_CancelStep_Handler,
		},
		{
			MethodName: "AuditStep",
			Handler:    _Service_AuditStep_Handler,
		},
		{
			MethodName: "DeletePipeline",
			Handler:    _Service_DeletePipeline_Handler,
		},
		{
			MethodName: "CreateAction",
			Handler:    _Service_CreateAction_Handler,
		},
		{
			MethodName: "QueryAction",
			Handler:    _Service_QueryAction_Handler,
		},
		{
			MethodName: "DescribeAction",
			Handler:    _Service_DescribeAction_Handler,
		},
		{
			MethodName: "DeleteAction",
			Handler:    _Service_DeleteAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/pkg/pipeline/pb/pipeline.proto",
}
