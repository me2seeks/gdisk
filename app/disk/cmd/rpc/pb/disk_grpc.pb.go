// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: disk.proto

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

// DiskClient is the client API for Disk service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiskClient interface {
	FileUploadPrepare(ctx context.Context, in *FileUploadPrepareRep, opts ...grpc.CallOption) (*FileUploadPrepareResp, error)
	UpdateFile(ctx context.Context, in *UpdateFileReq, opts ...grpc.CallOption) (*UpdateFileResp, error)
	//  rpc Statistics(StatisticsReq) returns(StatisticsResp);
	ListFile(ctx context.Context, in *ListFileReq, opts ...grpc.CallOption) (*ListFileResp, error)
	UnscopedFile(ctx context.Context, in *UnscopedFileReq, opts ...grpc.CallOption) (*UnscopedFileResp, error)
}

type diskClient struct {
	cc grpc.ClientConnInterface
}

func NewDiskClient(cc grpc.ClientConnInterface) DiskClient {
	return &diskClient{cc}
}

func (c *diskClient) FileUploadPrepare(ctx context.Context, in *FileUploadPrepareRep, opts ...grpc.CallOption) (*FileUploadPrepareResp, error) {
	out := new(FileUploadPrepareResp)
	err := c.cc.Invoke(ctx, "/pb.disk/FileUploadPrepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskClient) UpdateFile(ctx context.Context, in *UpdateFileReq, opts ...grpc.CallOption) (*UpdateFileResp, error) {
	out := new(UpdateFileResp)
	err := c.cc.Invoke(ctx, "/pb.disk/UpdateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskClient) ListFile(ctx context.Context, in *ListFileReq, opts ...grpc.CallOption) (*ListFileResp, error) {
	out := new(ListFileResp)
	err := c.cc.Invoke(ctx, "/pb.disk/ListFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskClient) UnscopedFile(ctx context.Context, in *UnscopedFileReq, opts ...grpc.CallOption) (*UnscopedFileResp, error) {
	out := new(UnscopedFileResp)
	err := c.cc.Invoke(ctx, "/pb.disk/UnscopedFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiskServer is the server API for Disk service.
// All implementations must embed UnimplementedDiskServer
// for forward compatibility
type DiskServer interface {
	FileUploadPrepare(context.Context, *FileUploadPrepareRep) (*FileUploadPrepareResp, error)
	UpdateFile(context.Context, *UpdateFileReq) (*UpdateFileResp, error)
	//  rpc Statistics(StatisticsReq) returns(StatisticsResp);
	ListFile(context.Context, *ListFileReq) (*ListFileResp, error)
	UnscopedFile(context.Context, *UnscopedFileReq) (*UnscopedFileResp, error)
	mustEmbedUnimplementedDiskServer()
}

// UnimplementedDiskServer must be embedded to have forward compatible implementations.
type UnimplementedDiskServer struct {
}

func (UnimplementedDiskServer) FileUploadPrepare(context.Context, *FileUploadPrepareRep) (*FileUploadPrepareResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileUploadPrepare not implemented")
}
func (UnimplementedDiskServer) UpdateFile(context.Context, *UpdateFileReq) (*UpdateFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFile not implemented")
}
func (UnimplementedDiskServer) ListFile(context.Context, *ListFileReq) (*ListFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFile not implemented")
}
func (UnimplementedDiskServer) UnscopedFile(context.Context, *UnscopedFileReq) (*UnscopedFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnscopedFile not implemented")
}
func (UnimplementedDiskServer) mustEmbedUnimplementedDiskServer() {}

// UnsafeDiskServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiskServer will
// result in compilation errors.
type UnsafeDiskServer interface {
	mustEmbedUnimplementedDiskServer()
}

func RegisterDiskServer(s grpc.ServiceRegistrar, srv DiskServer) {
	s.RegisterService(&Disk_ServiceDesc, srv)
}

func _Disk_FileUploadPrepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileUploadPrepareRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskServer).FileUploadPrepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.disk/FileUploadPrepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskServer).FileUploadPrepare(ctx, req.(*FileUploadPrepareRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _Disk_UpdateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskServer).UpdateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.disk/UpdateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskServer).UpdateFile(ctx, req.(*UpdateFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Disk_ListFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskServer).ListFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.disk/ListFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskServer).ListFile(ctx, req.(*ListFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Disk_UnscopedFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnscopedFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskServer).UnscopedFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.disk/UnscopedFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskServer).UnscopedFile(ctx, req.(*UnscopedFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Disk_ServiceDesc is the grpc.ServiceDesc for Disk service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Disk_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.disk",
	HandlerType: (*DiskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FileUploadPrepare",
			Handler:    _Disk_FileUploadPrepare_Handler,
		},
		{
			MethodName: "UpdateFile",
			Handler:    _Disk_UpdateFile_Handler,
		},
		{
			MethodName: "ListFile",
			Handler:    _Disk_ListFile_Handler,
		},
		{
			MethodName: "UnscopedFile",
			Handler:    _Disk_UnscopedFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "disk.proto",
}
