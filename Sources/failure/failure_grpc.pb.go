// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.2
// source: failure.proto

package failure

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

// FailureServiceClient is the client API for FailureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FailureServiceClient interface {
	// Injects a failure.
	Inject(ctx context.Context, in *InjectRequest, opts ...grpc.CallOption) (*InjectResponse, error)
}

type failureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFailureServiceClient(cc grpc.ClientConnInterface) FailureServiceClient {
	return &failureServiceClient{cc}
}

func (c *failureServiceClient) Inject(ctx context.Context, in *InjectRequest, opts ...grpc.CallOption) (*InjectResponse, error) {
	out := new(InjectResponse)
	err := c.cc.Invoke(ctx, "/mavsdk.rpc.failure.FailureService/Inject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FailureServiceServer is the server API for FailureService service.
// All implementations must embed UnimplementedFailureServiceServer
// for forward compatibility
type FailureServiceServer interface {
	// Injects a failure.
	Inject(context.Context, *InjectRequest) (*InjectResponse, error)
	mustEmbedUnimplementedFailureServiceServer()
}

// UnimplementedFailureServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFailureServiceServer struct {
}

func (UnimplementedFailureServiceServer) Inject(context.Context, *InjectRequest) (*InjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inject not implemented")
}
func (UnimplementedFailureServiceServer) mustEmbedUnimplementedFailureServiceServer() {}

// UnsafeFailureServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FailureServiceServer will
// result in compilation errors.
type UnsafeFailureServiceServer interface {
	mustEmbedUnimplementedFailureServiceServer()
}

func RegisterFailureServiceServer(s grpc.ServiceRegistrar, srv FailureServiceServer) {
	s.RegisterService(&FailureService_ServiceDesc, srv)
}

func _FailureService_Inject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FailureServiceServer).Inject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mavsdk.rpc.failure.FailureService/Inject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FailureServiceServer).Inject(ctx, req.(*InjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FailureService_ServiceDesc is the grpc.ServiceDesc for FailureService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FailureService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mavsdk.rpc.failure.FailureService",
	HandlerType: (*FailureServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Inject",
			Handler:    _FailureService_Inject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "failure.proto",
}