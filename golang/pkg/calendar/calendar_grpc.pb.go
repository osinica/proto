// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/calendar.proto

package calendar

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
	CalendarService_CreateEvent_FullMethodName = "/calendar.CalendarService/CreateEvent"
	CalendarService_FindOne_FullMethodName     = "/calendar.CalendarService/FindOne"
)

// CalendarServiceClient is the client API for CalendarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalendarServiceClient interface {
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error)
	FindOne(ctx context.Context, in *FindOneRequest, opts ...grpc.CallOption) (*FindOneResponse, error)
}

type calendarServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarServiceClient(cc grpc.ClientConnInterface) CalendarServiceClient {
	return &calendarServiceClient{cc}
}

func (c *calendarServiceClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error) {
	out := new(CreateEventResponse)
	err := c.cc.Invoke(ctx, CalendarService_CreateEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) FindOne(ctx context.Context, in *FindOneRequest, opts ...grpc.CallOption) (*FindOneResponse, error) {
	out := new(FindOneResponse)
	err := c.cc.Invoke(ctx, CalendarService_FindOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServiceServer is the server API for CalendarService service.
// All implementations must embed UnimplementedCalendarServiceServer
// for forward compatibility
type CalendarServiceServer interface {
	CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error)
	FindOne(context.Context, *FindOneRequest) (*FindOneResponse, error)
	mustEmbedUnimplementedCalendarServiceServer()
}

// UnimplementedCalendarServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalendarServiceServer struct {
}

func (UnimplementedCalendarServiceServer) CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedCalendarServiceServer) FindOne(context.Context, *FindOneRequest) (*FindOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (UnimplementedCalendarServiceServer) mustEmbedUnimplementedCalendarServiceServer() {}

// UnsafeCalendarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalendarServiceServer will
// result in compilation errors.
type UnsafeCalendarServiceServer interface {
	mustEmbedUnimplementedCalendarServiceServer()
}

func RegisterCalendarServiceServer(s grpc.ServiceRegistrar, srv CalendarServiceServer) {
	s.RegisterService(&CalendarService_ServiceDesc, srv)
}

func _CalendarService_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarService_CreateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarService_FindOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).FindOne(ctx, req.(*FindOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalendarService_ServiceDesc is the grpc.ServiceDesc for CalendarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalendarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calendar.CalendarService",
	HandlerType: (*CalendarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _CalendarService_CreateEvent_Handler,
		},
		{
			MethodName: "FindOne",
			Handler:    _CalendarService_FindOne_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/calendar.proto",
}
