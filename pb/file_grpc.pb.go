// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// FilesClient is the client API for Files service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilesClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (Files_UploadClient, error)
}

type filesClient struct {
	cc grpc.ClientConnInterface
}

func NewFilesClient(cc grpc.ClientConnInterface) FilesClient {
	return &filesClient{cc}
}

func (c *filesClient) Upload(ctx context.Context, opts ...grpc.CallOption) (Files_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Files_serviceDesc.Streams[0], "/file.Files/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &filesUploadClient{stream}
	return x, nil
}

type Files_UploadClient interface {
	Send(*FileChunk) error
	CloseAndRecv() (*File, error)
	grpc.ClientStream
}

type filesUploadClient struct {
	grpc.ClientStream
}

func (x *filesUploadClient) Send(m *FileChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *filesUploadClient) CloseAndRecv() (*File, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(File)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FilesServer is the server API for Files service.
// All implementations must embed UnimplementedFilesServer
// for forward compatibility
type FilesServer interface {
	Upload(Files_UploadServer) error
	mustEmbedUnimplementedFilesServer()
}

// UnimplementedFilesServer must be embedded to have forward compatible implementations.
type UnimplementedFilesServer struct {
}

func (UnimplementedFilesServer) Upload(Files_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedFilesServer) mustEmbedUnimplementedFilesServer() {}

// UnsafeFilesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilesServer will
// result in compilation errors.
type UnsafeFilesServer interface {
	mustEmbedUnimplementedFilesServer()
}

func RegisterFilesServer(s grpc.ServiceRegistrar, srv FilesServer) {
	s.RegisterService(&_Files_serviceDesc, srv)
}

func _Files_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FilesServer).Upload(&filesUploadServer{stream})
}

type Files_UploadServer interface {
	SendAndClose(*File) error
	Recv() (*FileChunk, error)
	grpc.ServerStream
}

type filesUploadServer struct {
	grpc.ServerStream
}

func (x *filesUploadServer) SendAndClose(m *File) error {
	return x.ServerStream.SendMsg(m)
}

func (x *filesUploadServer) Recv() (*FileChunk, error) {
	m := new(FileChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Files_serviceDesc = grpc.ServiceDesc{
	ServiceName: "file.Files",
	HandlerType: (*FilesServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Files_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "file.proto",
}