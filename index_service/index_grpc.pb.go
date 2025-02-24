// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.4
// source: index.proto

package indexservice

import (
	context "context"
	types "github.com/zlican/engine/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	IndexService_DeleteDoc_FullMethodName = "/indexservice.IndexService/DeleteDoc"
	IndexService_AddDoc_FullMethodName    = "/indexservice.IndexService/AddDoc"
	IndexService_Search_FullMethodName    = "/indexservice.IndexService/Search"
)

// IndexServiceClient is the client API for IndexService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 定义 gRPC 服务
type IndexServiceClient interface {
	DeleteDoc(ctx context.Context, in *DocID, opts ...grpc.CallOption) (*AffectedCount, error)
	AddDoc(ctx context.Context, in *types.Document, opts ...grpc.CallOption) (*AffectedCount, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResult, error)
}

type indexServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIndexServiceClient(cc grpc.ClientConnInterface) IndexServiceClient {
	return &indexServiceClient{cc}
}

func (c *indexServiceClient) DeleteDoc(ctx context.Context, in *DocID, opts ...grpc.CallOption) (*AffectedCount, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AffectedCount)
	err := c.cc.Invoke(ctx, IndexService_DeleteDoc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) AddDoc(ctx context.Context, in *types.Document, opts ...grpc.CallOption) (*AffectedCount, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AffectedCount)
	err := c.cc.Invoke(ctx, IndexService_AddDoc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchResult)
	err := c.cc.Invoke(ctx, IndexService_Search_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexServiceServer is the server API for IndexService service.
// All implementations must embed UnimplementedIndexServiceServer
// for forward compatibility.
//
// 定义 gRPC 服务
type IndexServiceServer interface {
	DeleteDoc(context.Context, *DocID) (*AffectedCount, error)
	AddDoc(context.Context, *types.Document) (*AffectedCount, error)
	Search(context.Context, *SearchRequest) (*SearchResult, error)
	mustEmbedUnimplementedIndexServiceServer()
}

// UnimplementedIndexServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedIndexServiceServer struct{}

func (UnimplementedIndexServiceServer) DeleteDoc(context.Context, *DocID) (*AffectedCount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDoc not implemented")
}
func (UnimplementedIndexServiceServer) AddDoc(context.Context, *types.Document) (*AffectedCount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDoc not implemented")
}
func (UnimplementedIndexServiceServer) Search(context.Context, *SearchRequest) (*SearchResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedIndexServiceServer) mustEmbedUnimplementedIndexServiceServer() {}
func (UnimplementedIndexServiceServer) testEmbeddedByValue()                      {}

// UnsafeIndexServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndexServiceServer will
// result in compilation errors.
type UnsafeIndexServiceServer interface {
	mustEmbedUnimplementedIndexServiceServer()
}

func RegisterIndexServiceServer(s grpc.ServiceRegistrar, srv IndexServiceServer) {
	// If the following call pancis, it indicates UnimplementedIndexServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&IndexService_ServiceDesc, srv)
}

func _IndexService_DeleteDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).DeleteDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IndexService_DeleteDoc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).DeleteDoc(ctx, req.(*DocID))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_AddDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Document)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).AddDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IndexService_AddDoc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).AddDoc(ctx, req.(*types.Document))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IndexService_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IndexService_ServiceDesc is the grpc.ServiceDesc for IndexService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IndexService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "indexservice.IndexService",
	HandlerType: (*IndexServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteDoc",
			Handler:    _IndexService_DeleteDoc_Handler,
		},
		{
			MethodName: "AddDoc",
			Handler:    _IndexService_AddDoc_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _IndexService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "index.proto",
}
