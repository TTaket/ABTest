// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: layer.proto

package pb_layer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LayerService_CreateLayer_FullMethodName = "/layer.LayerService/CreateLayer"
	LayerService_GetLayer_FullMethodName    = "/layer.LayerService/GetLayer"
	LayerService_UpdateLayer_FullMethodName = "/layer.LayerService/UpdateLayer"
	LayerService_DeleteLayer_FullMethodName = "/layer.LayerService/DeleteLayer"
	LayerService_ListLayers_FullMethodName  = "/layer.LayerService/ListLayers"
)

// LayerServiceClient is the client API for LayerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 层域服务
type LayerServiceClient interface {
	// 创建层域
	CreateLayer(ctx context.Context, in *CreateLayerRequest, opts ...grpc.CallOption) (*CreateLayerResponse, error)
	// 获取层域
	GetLayer(ctx context.Context, in *GetLayerRequest, opts ...grpc.CallOption) (*GetLayerResponse, error)
	// 更新层域
	UpdateLayer(ctx context.Context, in *UpdateLayerRequest, opts ...grpc.CallOption) (*UpdateLayerResponse, error)
	// 删除层域
	DeleteLayer(ctx context.Context, in *DeleteLayerRequest, opts ...grpc.CallOption) (*DeleteLayerResponse, error)
	// 列出层域
	ListLayers(ctx context.Context, in *ListLayersRequest, opts ...grpc.CallOption) (*ListLayersResponse, error)
}

type layerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLayerServiceClient(cc grpc.ClientConnInterface) LayerServiceClient {
	return &layerServiceClient{cc}
}

func (c *layerServiceClient) CreateLayer(ctx context.Context, in *CreateLayerRequest, opts ...grpc.CallOption) (*CreateLayerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateLayerResponse)
	err := c.cc.Invoke(ctx, LayerService_CreateLayer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *layerServiceClient) GetLayer(ctx context.Context, in *GetLayerRequest, opts ...grpc.CallOption) (*GetLayerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLayerResponse)
	err := c.cc.Invoke(ctx, LayerService_GetLayer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *layerServiceClient) UpdateLayer(ctx context.Context, in *UpdateLayerRequest, opts ...grpc.CallOption) (*UpdateLayerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateLayerResponse)
	err := c.cc.Invoke(ctx, LayerService_UpdateLayer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *layerServiceClient) DeleteLayer(ctx context.Context, in *DeleteLayerRequest, opts ...grpc.CallOption) (*DeleteLayerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteLayerResponse)
	err := c.cc.Invoke(ctx, LayerService_DeleteLayer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *layerServiceClient) ListLayers(ctx context.Context, in *ListLayersRequest, opts ...grpc.CallOption) (*ListLayersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLayersResponse)
	err := c.cc.Invoke(ctx, LayerService_ListLayers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LayerServiceServer is the server API for LayerService service.
// All implementations must embed UnimplementedLayerServiceServer
// for forward compatibility.
//
// 层域服务
type LayerServiceServer interface {
	// 创建层域
	CreateLayer(context.Context, *CreateLayerRequest) (*CreateLayerResponse, error)
	// 获取层域
	GetLayer(context.Context, *GetLayerRequest) (*GetLayerResponse, error)
	// 更新层域
	UpdateLayer(context.Context, *UpdateLayerRequest) (*UpdateLayerResponse, error)
	// 删除层域
	DeleteLayer(context.Context, *DeleteLayerRequest) (*DeleteLayerResponse, error)
	// 列出层域
	ListLayers(context.Context, *ListLayersRequest) (*ListLayersResponse, error)
	mustEmbedUnimplementedLayerServiceServer()
}

// UnimplementedLayerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLayerServiceServer struct{}

func (UnimplementedLayerServiceServer) CreateLayer(context.Context, *CreateLayerRequest) (*CreateLayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLayer not implemented")
}
func (UnimplementedLayerServiceServer) GetLayer(context.Context, *GetLayerRequest) (*GetLayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLayer not implemented")
}
func (UnimplementedLayerServiceServer) UpdateLayer(context.Context, *UpdateLayerRequest) (*UpdateLayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLayer not implemented")
}
func (UnimplementedLayerServiceServer) DeleteLayer(context.Context, *DeleteLayerRequest) (*DeleteLayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLayer not implemented")
}
func (UnimplementedLayerServiceServer) ListLayers(context.Context, *ListLayersRequest) (*ListLayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLayers not implemented")
}
func (UnimplementedLayerServiceServer) mustEmbedUnimplementedLayerServiceServer() {}
func (UnimplementedLayerServiceServer) testEmbeddedByValue()                      {}

// UnsafeLayerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LayerServiceServer will
// result in compilation errors.
type UnsafeLayerServiceServer interface {
	mustEmbedUnimplementedLayerServiceServer()
}

func RegisterLayerServiceServer(s grpc.ServiceRegistrar, srv LayerServiceServer) {
	// If the following call pancis, it indicates UnimplementedLayerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LayerService_ServiceDesc, srv)
}

func _LayerService_CreateLayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LayerServiceServer).CreateLayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LayerService_CreateLayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LayerServiceServer).CreateLayer(ctx, req.(*CreateLayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LayerService_GetLayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LayerServiceServer).GetLayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LayerService_GetLayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LayerServiceServer).GetLayer(ctx, req.(*GetLayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LayerService_UpdateLayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LayerServiceServer).UpdateLayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LayerService_UpdateLayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LayerServiceServer).UpdateLayer(ctx, req.(*UpdateLayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LayerService_DeleteLayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LayerServiceServer).DeleteLayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LayerService_DeleteLayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LayerServiceServer).DeleteLayer(ctx, req.(*DeleteLayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LayerService_ListLayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LayerServiceServer).ListLayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LayerService_ListLayers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LayerServiceServer).ListLayers(ctx, req.(*ListLayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LayerService_ServiceDesc is the grpc.ServiceDesc for LayerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LayerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "layer.LayerService",
	HandlerType: (*LayerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLayer",
			Handler:    _LayerService_CreateLayer_Handler,
		},
		{
			MethodName: "GetLayer",
			Handler:    _LayerService_GetLayer_Handler,
		},
		{
			MethodName: "UpdateLayer",
			Handler:    _LayerService_UpdateLayer_Handler,
		},
		{
			MethodName: "DeleteLayer",
			Handler:    _LayerService_DeleteLayer_Handler,
		},
		{
			MethodName: "ListLayers",
			Handler:    _LayerService_ListLayers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "layer.proto",
}
