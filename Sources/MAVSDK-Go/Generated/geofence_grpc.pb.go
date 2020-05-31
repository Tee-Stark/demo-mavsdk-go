// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mavsdkgo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GeofenceServiceClient is the client API for GeofenceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeofenceServiceClient interface {
	//
	// Upload a geofence.
	//
	// Polygons are uploaded to a drone. Once uploaded, the geofence will remain
	// on the drone even if a connection is lost.
	UploadGeofence(ctx context.Context, in *UploadGeofenceRequest, opts ...grpc.CallOption) (*UploadGeofenceResponse, error)
}

type geofenceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGeofenceServiceClient(cc grpc.ClientConnInterface) GeofenceServiceClient {
	return &geofenceServiceClient{cc}
}

func (c *geofenceServiceClient) UploadGeofence(ctx context.Context, in *UploadGeofenceRequest, opts ...grpc.CallOption) (*UploadGeofenceResponse, error) {
	out := new(UploadGeofenceResponse)
	err := c.cc.Invoke(ctx, "/mavsdk.rpc.geofence.GeofenceService/UploadGeofence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeofenceServiceServer is the server API for GeofenceService service.
type GeofenceServiceServer interface {
	//
	// Upload a geofence.
	//
	// Polygons are uploaded to a drone. Once uploaded, the geofence will remain
	// on the drone even if a connection is lost.
	UploadGeofence(context.Context, *UploadGeofenceRequest) (*UploadGeofenceResponse, error)
}

// UnimplementedGeofenceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGeofenceServiceServer struct {
}

func (*UnimplementedGeofenceServiceServer) UploadGeofence(context.Context, *UploadGeofenceRequest) (*UploadGeofenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadGeofence not implemented")
}

func RegisterGeofenceServiceServer(s *grpc.Server, srv GeofenceServiceServer) {
	s.RegisterService(&_GeofenceService_serviceDesc, srv)
}

func _GeofenceService_UploadGeofence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadGeofenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeofenceServiceServer).UploadGeofence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mavsdk.rpc.geofence.GeofenceService/UploadGeofence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeofenceServiceServer).UploadGeofence(ctx, req.(*UploadGeofenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GeofenceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mavsdk.rpc.geofence.GeofenceService",
	HandlerType: (*GeofenceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadGeofence",
			Handler:    _GeofenceService_UploadGeofence_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "geofence.proto",
}