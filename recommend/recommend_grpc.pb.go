// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.0
// source: recommend/recommend.proto

package recommendproto

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
	RecommendService_GetRecommendCategoryIDs_FullMethodName = "/recommend.RecommendService/GetRecommendCategoryIDs"
	RecommendService_GetRecommendProductIDs_FullMethodName  = "/recommend.RecommendService/GetRecommendProductIDs"
)

// RecommendServiceClient is the client API for RecommendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The greeting service definition.
type RecommendServiceClient interface {
	// get category recommend
	GetRecommendCategoryIDs(ctx context.Context, in *GetRecommendCategoryIDsRequest, opts ...grpc.CallOption) (*GetRecommendCategoryIDsResponse, error)
	// get product recommend
	GetRecommendProductIDs(ctx context.Context, in *GetRecommendProductIDsRequest, opts ...grpc.CallOption) (*GetRecommendProductIDsResponse, error)
}

type recommendServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecommendServiceClient(cc grpc.ClientConnInterface) RecommendServiceClient {
	return &recommendServiceClient{cc}
}

func (c *recommendServiceClient) GetRecommendCategoryIDs(ctx context.Context, in *GetRecommendCategoryIDsRequest, opts ...grpc.CallOption) (*GetRecommendCategoryIDsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRecommendCategoryIDsResponse)
	err := c.cc.Invoke(ctx, RecommendService_GetRecommendCategoryIDs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendServiceClient) GetRecommendProductIDs(ctx context.Context, in *GetRecommendProductIDsRequest, opts ...grpc.CallOption) (*GetRecommendProductIDsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRecommendProductIDsResponse)
	err := c.cc.Invoke(ctx, RecommendService_GetRecommendProductIDs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommendServiceServer is the server API for RecommendService service.
// All implementations must embed UnimplementedRecommendServiceServer
// for forward compatibility.
//
// The greeting service definition.
type RecommendServiceServer interface {
	// get category recommend
	GetRecommendCategoryIDs(context.Context, *GetRecommendCategoryIDsRequest) (*GetRecommendCategoryIDsResponse, error)
	// get product recommend
	GetRecommendProductIDs(context.Context, *GetRecommendProductIDsRequest) (*GetRecommendProductIDsResponse, error)
	mustEmbedUnimplementedRecommendServiceServer()
}

// UnimplementedRecommendServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRecommendServiceServer struct{}

func (UnimplementedRecommendServiceServer) GetRecommendCategoryIDs(context.Context, *GetRecommendCategoryIDsRequest) (*GetRecommendCategoryIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendCategoryIDs not implemented")
}
func (UnimplementedRecommendServiceServer) GetRecommendProductIDs(context.Context, *GetRecommendProductIDsRequest) (*GetRecommendProductIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendProductIDs not implemented")
}
func (UnimplementedRecommendServiceServer) mustEmbedUnimplementedRecommendServiceServer() {}
func (UnimplementedRecommendServiceServer) testEmbeddedByValue()                          {}

// UnsafeRecommendServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecommendServiceServer will
// result in compilation errors.
type UnsafeRecommendServiceServer interface {
	mustEmbedUnimplementedRecommendServiceServer()
}

func RegisterRecommendServiceServer(s grpc.ServiceRegistrar, srv RecommendServiceServer) {
	// If the following call pancis, it indicates UnimplementedRecommendServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RecommendService_ServiceDesc, srv)
}

func _RecommendService_GetRecommendCategoryIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecommendCategoryIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendServiceServer).GetRecommendCategoryIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecommendService_GetRecommendCategoryIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendServiceServer).GetRecommendCategoryIDs(ctx, req.(*GetRecommendCategoryIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendService_GetRecommendProductIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecommendProductIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendServiceServer).GetRecommendProductIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecommendService_GetRecommendProductIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendServiceServer).GetRecommendProductIDs(ctx, req.(*GetRecommendProductIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecommendService_ServiceDesc is the grpc.ServiceDesc for RecommendService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecommendService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "recommend.RecommendService",
	HandlerType: (*RecommendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecommendCategoryIDs",
			Handler:    _RecommendService_GetRecommendCategoryIDs_Handler,
		},
		{
			MethodName: "GetRecommendProductIDs",
			Handler:    _RecommendService_GetRecommendProductIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recommend/recommend.proto",
}
