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
	ProductService_GetRecommendCategoryIDs_FullMethodName = "/recommend.ProductService/GetRecommendCategoryIDs"
	ProductService_GetRecommendProductIDs_FullMethodName  = "/recommend.ProductService/GetRecommendProductIDs"
)

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The greeting service definition.
type ProductServiceClient interface {
	// get category recommend
	GetRecommendCategoryIDs(ctx context.Context, in *GetRecommendCategoryIDsRequest, opts ...grpc.CallOption) (*GetRecommendCategoryIDsResponse, error)
	// get product recommend
	GetRecommendProductIDs(ctx context.Context, in *GetRecommendCategoryIDsRequest, opts ...grpc.CallOption) (*GetRecommendProductIDsResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetRecommendCategoryIDs(ctx context.Context, in *GetRecommendCategoryIDsRequest, opts ...grpc.CallOption) (*GetRecommendCategoryIDsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRecommendCategoryIDsResponse)
	err := c.cc.Invoke(ctx, ProductService_GetRecommendCategoryIDs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetRecommendProductIDs(ctx context.Context, in *GetRecommendCategoryIDsRequest, opts ...grpc.CallOption) (*GetRecommendProductIDsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRecommendProductIDsResponse)
	err := c.cc.Invoke(ctx, ProductService_GetRecommendProductIDs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility.
//
// The greeting service definition.
type ProductServiceServer interface {
	// get category recommend
	GetRecommendCategoryIDs(context.Context, *GetRecommendCategoryIDsRequest) (*GetRecommendCategoryIDsResponse, error)
	// get product recommend
	GetRecommendProductIDs(context.Context, *GetRecommendCategoryIDsRequest) (*GetRecommendProductIDsResponse, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductServiceServer struct{}

func (UnimplementedProductServiceServer) GetRecommendCategoryIDs(context.Context, *GetRecommendCategoryIDsRequest) (*GetRecommendCategoryIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendCategoryIDs not implemented")
}
func (UnimplementedProductServiceServer) GetRecommendProductIDs(context.Context, *GetRecommendCategoryIDsRequest) (*GetRecommendProductIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendProductIDs not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}
func (UnimplementedProductServiceServer) testEmbeddedByValue()                        {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	// If the following call pancis, it indicates UnimplementedProductServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_GetRecommendCategoryIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecommendCategoryIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetRecommendCategoryIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetRecommendCategoryIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetRecommendCategoryIDs(ctx, req.(*GetRecommendCategoryIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetRecommendProductIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecommendCategoryIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetRecommendProductIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetRecommendProductIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetRecommendProductIDs(ctx, req.(*GetRecommendCategoryIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "recommend.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecommendCategoryIDs",
			Handler:    _ProductService_GetRecommendCategoryIDs_Handler,
		},
		{
			MethodName: "GetRecommendProductIDs",
			Handler:    _ProductService_GetRecommendProductIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recommend/recommend.proto",
}
