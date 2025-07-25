// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/temp_sensor.proto

package proto

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
	TempSensorService_CreateTempSensor_FullMethodName = "/proto.TempSensorService/CreateTempSensor"
	TempSensorService_GetAllTempSensor_FullMethodName = "/proto.TempSensorService/GetAllTempSensor"
)

// TempSensorServiceClient is the client API for TempSensorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TempSensorServiceClient interface {
	CreateTempSensor(ctx context.Context, in *CreateTempSensorReq, opts ...grpc.CallOption) (*CreateTempSensorRes, error)
	GetAllTempSensor(ctx context.Context, in *GetAllTempSensorReq, opts ...grpc.CallOption) (*GetAllTempSensorRes, error)
}

type tempSensorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTempSensorServiceClient(cc grpc.ClientConnInterface) TempSensorServiceClient {
	return &tempSensorServiceClient{cc}
}

func (c *tempSensorServiceClient) CreateTempSensor(ctx context.Context, in *CreateTempSensorReq, opts ...grpc.CallOption) (*CreateTempSensorRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTempSensorRes)
	err := c.cc.Invoke(ctx, TempSensorService_CreateTempSensor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tempSensorServiceClient) GetAllTempSensor(ctx context.Context, in *GetAllTempSensorReq, opts ...grpc.CallOption) (*GetAllTempSensorRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllTempSensorRes)
	err := c.cc.Invoke(ctx, TempSensorService_GetAllTempSensor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TempSensorServiceServer is the server API for TempSensorService service.
// All implementations must embed UnimplementedTempSensorServiceServer
// for forward compatibility.
type TempSensorServiceServer interface {
	CreateTempSensor(context.Context, *CreateTempSensorReq) (*CreateTempSensorRes, error)
	GetAllTempSensor(context.Context, *GetAllTempSensorReq) (*GetAllTempSensorRes, error)
	mustEmbedUnimplementedTempSensorServiceServer()
}

// UnimplementedTempSensorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTempSensorServiceServer struct{}

func (UnimplementedTempSensorServiceServer) CreateTempSensor(context.Context, *CreateTempSensorReq) (*CreateTempSensorRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTempSensor not implemented")
}
func (UnimplementedTempSensorServiceServer) GetAllTempSensor(context.Context, *GetAllTempSensorReq) (*GetAllTempSensorRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTempSensor not implemented")
}
func (UnimplementedTempSensorServiceServer) mustEmbedUnimplementedTempSensorServiceServer() {}
func (UnimplementedTempSensorServiceServer) testEmbeddedByValue()                           {}

// UnsafeTempSensorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TempSensorServiceServer will
// result in compilation errors.
type UnsafeTempSensorServiceServer interface {
	mustEmbedUnimplementedTempSensorServiceServer()
}

func RegisterTempSensorServiceServer(s grpc.ServiceRegistrar, srv TempSensorServiceServer) {
	// If the following call pancis, it indicates UnimplementedTempSensorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TempSensorService_ServiceDesc, srv)
}

func _TempSensorService_CreateTempSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTempSensorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TempSensorServiceServer).CreateTempSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TempSensorService_CreateTempSensor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TempSensorServiceServer).CreateTempSensor(ctx, req.(*CreateTempSensorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TempSensorService_GetAllTempSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTempSensorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TempSensorServiceServer).GetAllTempSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TempSensorService_GetAllTempSensor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TempSensorServiceServer).GetAllTempSensor(ctx, req.(*GetAllTempSensorReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TempSensorService_ServiceDesc is the grpc.ServiceDesc for TempSensorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TempSensorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TempSensorService",
	HandlerType: (*TempSensorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTempSensor",
			Handler:    _TempSensorService_CreateTempSensor_Handler,
		},
		{
			MethodName: "GetAllTempSensor",
			Handler:    _TempSensorService_GetAllTempSensor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/temp_sensor.proto",
}
