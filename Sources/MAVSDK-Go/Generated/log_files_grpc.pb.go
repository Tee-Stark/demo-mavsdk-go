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

// LogFilesServiceClient is the client API for LogFilesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogFilesServiceClient interface {
	// Get List of log files.
	GetEntries(ctx context.Context, in *GetEntriesRequest, opts ...grpc.CallOption) (*GetEntriesResponse, error)
	// Download log file.
	SubscribeDownloadLogFile(ctx context.Context, in *SubscribeDownloadLogFileRequest, opts ...grpc.CallOption) (LogFilesService_SubscribeDownloadLogFileClient, error)
}

type logFilesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogFilesServiceClient(cc grpc.ClientConnInterface) LogFilesServiceClient {
	return &logFilesServiceClient{cc}
}

func (c *logFilesServiceClient) GetEntries(ctx context.Context, in *GetEntriesRequest, opts ...grpc.CallOption) (*GetEntriesResponse, error) {
	out := new(GetEntriesResponse)
	err := c.cc.Invoke(ctx, "/mavsdk.rpc.log_files.LogFilesService/GetEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logFilesServiceClient) SubscribeDownloadLogFile(ctx context.Context, in *SubscribeDownloadLogFileRequest, opts ...grpc.CallOption) (LogFilesService_SubscribeDownloadLogFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogFilesService_serviceDesc.Streams[0], "/mavsdk.rpc.log_files.LogFilesService/SubscribeDownloadLogFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &logFilesServiceSubscribeDownloadLogFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogFilesService_SubscribeDownloadLogFileClient interface {
	Recv() (*DownloadLogFileResponse, error)
	grpc.ClientStream
}

type logFilesServiceSubscribeDownloadLogFileClient struct {
	grpc.ClientStream
}

func (x *logFilesServiceSubscribeDownloadLogFileClient) Recv() (*DownloadLogFileResponse, error) {
	m := new(DownloadLogFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogFilesServiceServer is the server API for LogFilesService service.
type LogFilesServiceServer interface {
	// Get List of log files.
	GetEntries(context.Context, *GetEntriesRequest) (*GetEntriesResponse, error)
	// Download log file.
	SubscribeDownloadLogFile(*SubscribeDownloadLogFileRequest, LogFilesService_SubscribeDownloadLogFileServer) error
}

// UnimplementedLogFilesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLogFilesServiceServer struct {
}

func (*UnimplementedLogFilesServiceServer) GetEntries(context.Context, *GetEntriesRequest) (*GetEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntries not implemented")
}
func (*UnimplementedLogFilesServiceServer) SubscribeDownloadLogFile(*SubscribeDownloadLogFileRequest, LogFilesService_SubscribeDownloadLogFileServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeDownloadLogFile not implemented")
}

func RegisterLogFilesServiceServer(s *grpc.Server, srv LogFilesServiceServer) {
	s.RegisterService(&_LogFilesService_serviceDesc, srv)
}

func _LogFilesService_GetEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogFilesServiceServer).GetEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mavsdk.rpc.log_files.LogFilesService/GetEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogFilesServiceServer).GetEntries(ctx, req.(*GetEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogFilesService_SubscribeDownloadLogFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeDownloadLogFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogFilesServiceServer).SubscribeDownloadLogFile(m, &logFilesServiceSubscribeDownloadLogFileServer{stream})
}

type LogFilesService_SubscribeDownloadLogFileServer interface {
	Send(*DownloadLogFileResponse) error
	grpc.ServerStream
}

type logFilesServiceSubscribeDownloadLogFileServer struct {
	grpc.ServerStream
}

func (x *logFilesServiceSubscribeDownloadLogFileServer) Send(m *DownloadLogFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _LogFilesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mavsdk.rpc.log_files.LogFilesService",
	HandlerType: (*LogFilesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntries",
			Handler:    _LogFilesService_GetEntries_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeDownloadLogFile",
			Handler:       _LogFilesService_SubscribeDownloadLogFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "log_files.proto",
}
