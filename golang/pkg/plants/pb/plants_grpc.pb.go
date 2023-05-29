// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/plants.proto

package pb

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

const (
	PlantsService_List_FullMethodName = "/plants.PlantsService/List"
)

// PlantsServiceClient is the client API for PlantsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlantsServiceClient interface {
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ObjectList, error)
}

type plantsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlantsServiceClient(cc grpc.ClientConnInterface) PlantsServiceClient {
	return &plantsServiceClient{cc}
}

func (c *plantsServiceClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ObjectList, error) {
	out := new(ObjectList)
	err := c.cc.Invoke(ctx, PlantsService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlantsServiceServer is the server API for PlantsService service.
// All implementations must embed UnimplementedPlantsServiceServer
// for forward compatibility
type PlantsServiceServer interface {
	List(context.Context, *Empty) (*ObjectList, error)
	mustEmbedUnimplementedPlantsServiceServer()
}

// UnimplementedPlantsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPlantsServiceServer struct {
}

func (UnimplementedPlantsServiceServer) List(context.Context, *Empty) (*ObjectList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPlantsServiceServer) mustEmbedUnimplementedPlantsServiceServer() {}

// UnsafePlantsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlantsServiceServer will
// result in compilation errors.
type UnsafePlantsServiceServer interface {
	mustEmbedUnimplementedPlantsServiceServer()
}

func RegisterPlantsServiceServer(s grpc.ServiceRegistrar, srv PlantsServiceServer) {
	s.RegisterService(&PlantsService_ServiceDesc, srv)
}

func _PlantsService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlantsServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlantsService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlantsServiceServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// PlantsService_ServiceDesc is the grpc.ServiceDesc for PlantsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlantsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "plants.PlantsService",
	HandlerType: (*PlantsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _PlantsService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/plants.proto",
}
