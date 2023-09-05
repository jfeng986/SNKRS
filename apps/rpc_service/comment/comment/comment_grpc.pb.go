// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: comment.proto

package comment

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
	CommentService_Comments_FullMethodName = "/comment.CommentService/Comments"
)

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentServiceClient interface {
	Comments(ctx context.Context, in *CommentsRequest, opts ...grpc.CallOption) (*CommentsResponse, error)
}

type commentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentServiceClient(cc grpc.ClientConnInterface) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) Comments(ctx context.Context, in *CommentsRequest, opts ...grpc.CallOption) (*CommentsResponse, error) {
	out := new(CommentsResponse)
	err := c.cc.Invoke(ctx, CommentService_Comments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
// All implementations must embed UnimplementedCommentServiceServer
// for forward compatibility
type CommentServiceServer interface {
	Comments(context.Context, *CommentsRequest) (*CommentsResponse, error)
	mustEmbedUnimplementedCommentServiceServer()
}

// UnimplementedCommentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServiceServer struct {
}

func (UnimplementedCommentServiceServer) Comments(context.Context, *CommentsRequest) (*CommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comments not implemented")
}
func (UnimplementedCommentServiceServer) mustEmbedUnimplementedCommentServiceServer() {}

// UnsafeCommentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServiceServer will
// result in compilation errors.
type UnsafeCommentServiceServer interface {
	mustEmbedUnimplementedCommentServiceServer()
}

func RegisterCommentServiceServer(s grpc.ServiceRegistrar, srv CommentServiceServer) {
	s.RegisterService(&CommentService_ServiceDesc, srv)
}

func _CommentService_Comments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).Comments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_Comments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).Comments(ctx, req.(*CommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentService_ServiceDesc is the grpc.ServiceDesc for CommentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Comments",
			Handler:    _CommentService_Comments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
