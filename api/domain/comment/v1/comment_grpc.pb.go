// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package commentv1

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

// CommentRpcClient is the client API for CommentRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentRpcClient interface {
	CreateSubject(ctx context.Context, in *CreateSubjectReq, opts ...grpc.CallOption) (*CreateSubjectResponse, error)
	GetSubjectById(ctx context.Context, in *GetSubjectByIdReq, opts ...grpc.CallOption) (*GetSubjectByIdResponse, error)
	GetSubjectByBelong(ctx context.Context, in *GetSubjectByBelongReq, opts ...grpc.CallOption) (*GetSubjectByBelongResponse, error)
	CreateReply(ctx context.Context, in *CreateReplyReq, opts ...grpc.CallOption) (*CreateReplyResponse, error)
	// rpc GetReplyById(GetReplyByIdReq) returns (GetReplyByIdReponse);
	ListReplyByIds(ctx context.Context, in *ListReplyByIdsReq, opts ...grpc.CallOption) (*ListReplyByIdsResponse, error)
	ListFloorBySubjectSortReply(ctx context.Context, in *ListFloorBySubjectSortReplyReq, opts ...grpc.CallOption) (*ListFloorBySubjectSortReplyResponse, error)
	ListFloorBySubjectSortTime(ctx context.Context, in *ListFloorBySubjectSortTimeReq, opts ...grpc.CallOption) (*ListFloorBySubjectSortTimeResponse, error)
	ListReplyByFloorSortLike(ctx context.Context, in *ListReplyByFloorSortLikeReq, opts ...grpc.CallOption) (*ListReplyByFloorSortLikeResponse, error)
	ListReplyByFloorSortTime(ctx context.Context, in *ListReplyByFloorSortTimeReq, opts ...grpc.CallOption) (*ListReplyByFloorSortTimeResponse, error)
	AddLike(ctx context.Context, in *AddLikeReq, opts ...grpc.CallOption) (*AddLikeResponse, error)
	CancelLike(ctx context.Context, in *CancelLikeReq, opts ...grpc.CallOption) (*CancelLikeResponse, error)
	IsLike(ctx context.Context, in *IsLikeReq, opts ...grpc.CallOption) (*IsLikeResponse, error)
}

type commentRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentRpcClient(cc grpc.ClientConnInterface) CommentRpcClient {
	return &commentRpcClient{cc}
}

func (c *commentRpcClient) CreateSubject(ctx context.Context, in *CreateSubjectReq, opts ...grpc.CallOption) (*CreateSubjectResponse, error) {
	out := new(CreateSubjectResponse)
	err := c.cc.Invoke(ctx, "/domain.comment.v1.CommentRpc/CreateSubject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) GetSubjectById(ctx context.Context, in *GetSubjectByIdReq, opts ...grpc.CallOption) (*GetSubjectByIdResponse, error) {
	out := new(GetSubjectByIdResponse)
	err := c.cc.Invoke(ctx, "/domain.comment.v1.CommentRpc/GetSubjectById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) GetSubjectByBelong(ctx context.Context, in *GetSubjectByBelongReq, opts ...grpc.CallOption) (*GetSubjectByBelongResponse, error) {
	out := new(GetSubjectByBelongResponse)
	err := c.cc.Invoke(ctx, "/domain.comment.v1.CommentRpc/GetSubjectByBelong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) CreateReply(ctx context.Context, in *CreateReplyReq, opts ...grpc.CallOption) (*CreateReplyResponse, error) {
	out := new(CreateReplyResponse)
	err := c.cc.Invoke(ctx, "/domain.comment.v1.CommentRpc/CreateReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) ListReplyByIds(ctx context.Context, in *ListReplyByIdsReq, opts ...grpc.CallOption) (*ListReplyByIdsResponse, error) {
	out := new(ListReplyByIdsResponse)
	err := c.cc.Invoke(ctx, "/domain.comment.v1.CommentRpc/ListReplyByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) ListFloorBySubjectSortReply(ctx context.Context, in *ListFloorBySubjectSortReplyReq, opts ...grpc.CallOption) (*ListFloorBySubjectSortReplyResponse, error) {
	out := new(ListFloorBySubjectSortReplyResponse)
	err := c.cc.Invoke(ctx, "/domain.comment.v1.CommentRpc/ListFloorBySubjectSortReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) ListFloorBySubjectSortTime(ctx context.Context, in *ListFloorBySubjectSortTimeReq, opts ...grpc.CallOption) (*ListFloorBySubjectSortTimeResponse, error) {
	out := new(ListFloorBySubjectSortTimeResponse)
	err := c.cc.Invoke(ctx, "/domain.comment.v1.CommentRpc/ListFloorBySubjectSortTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) ListReplyByFloorSortLike(ctx context.Context, in *ListReplyByFloorSortLikeReq, opts ...grpc.CallOption) (*ListReplyByFloorSortLikeResponse, error) {
	out := new(ListReplyByFloorSortLikeResponse)
	err := c.cc.Invoke(ctx, "/domain.comment.v1.CommentRpc/ListReplyByFloorSortLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) ListReplyByFloorSortTime(ctx context.Context, in *ListReplyByFloorSortTimeReq, opts ...grpc.CallOption) (*ListReplyByFloorSortTimeResponse, error) {
	out := new(ListReplyByFloorSortTimeResponse)
	err := c.cc.Invoke(ctx, "/domain.comment.v1.CommentRpc/ListReplyByFloorSortTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) AddLike(ctx context.Context, in *AddLikeReq, opts ...grpc.CallOption) (*AddLikeResponse, error) {
	out := new(AddLikeResponse)
	err := c.cc.Invoke(ctx, "/domain.comment.v1.CommentRpc/AddLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) CancelLike(ctx context.Context, in *CancelLikeReq, opts ...grpc.CallOption) (*CancelLikeResponse, error) {
	out := new(CancelLikeResponse)
	err := c.cc.Invoke(ctx, "/domain.comment.v1.CommentRpc/CancelLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) IsLike(ctx context.Context, in *IsLikeReq, opts ...grpc.CallOption) (*IsLikeResponse, error) {
	out := new(IsLikeResponse)
	err := c.cc.Invoke(ctx, "/domain.comment.v1.CommentRpc/IsLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentRpcServer is the server API for CommentRpc service.
// All implementations must embed UnimplementedCommentRpcServer
// for forward compatibility
type CommentRpcServer interface {
	CreateSubject(context.Context, *CreateSubjectReq) (*CreateSubjectResponse, error)
	GetSubjectById(context.Context, *GetSubjectByIdReq) (*GetSubjectByIdResponse, error)
	GetSubjectByBelong(context.Context, *GetSubjectByBelongReq) (*GetSubjectByBelongResponse, error)
	CreateReply(context.Context, *CreateReplyReq) (*CreateReplyResponse, error)
	// rpc GetReplyById(GetReplyByIdReq) returns (GetReplyByIdReponse);
	ListReplyByIds(context.Context, *ListReplyByIdsReq) (*ListReplyByIdsResponse, error)
	ListFloorBySubjectSortReply(context.Context, *ListFloorBySubjectSortReplyReq) (*ListFloorBySubjectSortReplyResponse, error)
	ListFloorBySubjectSortTime(context.Context, *ListFloorBySubjectSortTimeReq) (*ListFloorBySubjectSortTimeResponse, error)
	ListReplyByFloorSortLike(context.Context, *ListReplyByFloorSortLikeReq) (*ListReplyByFloorSortLikeResponse, error)
	ListReplyByFloorSortTime(context.Context, *ListReplyByFloorSortTimeReq) (*ListReplyByFloorSortTimeResponse, error)
	AddLike(context.Context, *AddLikeReq) (*AddLikeResponse, error)
	CancelLike(context.Context, *CancelLikeReq) (*CancelLikeResponse, error)
	IsLike(context.Context, *IsLikeReq) (*IsLikeResponse, error)
	mustEmbedUnimplementedCommentRpcServer()
}

// UnimplementedCommentRpcServer must be embedded to have forward compatible implementations.
type UnimplementedCommentRpcServer struct {
}

func (UnimplementedCommentRpcServer) CreateSubject(context.Context, *CreateSubjectReq) (*CreateSubjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubject not implemented")
}
func (UnimplementedCommentRpcServer) GetSubjectById(context.Context, *GetSubjectByIdReq) (*GetSubjectByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubjectById not implemented")
}
func (UnimplementedCommentRpcServer) GetSubjectByBelong(context.Context, *GetSubjectByBelongReq) (*GetSubjectByBelongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubjectByBelong not implemented")
}
func (UnimplementedCommentRpcServer) CreateReply(context.Context, *CreateReplyReq) (*CreateReplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReply not implemented")
}
func (UnimplementedCommentRpcServer) ListReplyByIds(context.Context, *ListReplyByIdsReq) (*ListReplyByIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReplyByIds not implemented")
}
func (UnimplementedCommentRpcServer) ListFloorBySubjectSortReply(context.Context, *ListFloorBySubjectSortReplyReq) (*ListFloorBySubjectSortReplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFloorBySubjectSortReply not implemented")
}
func (UnimplementedCommentRpcServer) ListFloorBySubjectSortTime(context.Context, *ListFloorBySubjectSortTimeReq) (*ListFloorBySubjectSortTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFloorBySubjectSortTime not implemented")
}
func (UnimplementedCommentRpcServer) ListReplyByFloorSortLike(context.Context, *ListReplyByFloorSortLikeReq) (*ListReplyByFloorSortLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReplyByFloorSortLike not implemented")
}
func (UnimplementedCommentRpcServer) ListReplyByFloorSortTime(context.Context, *ListReplyByFloorSortTimeReq) (*ListReplyByFloorSortTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReplyByFloorSortTime not implemented")
}
func (UnimplementedCommentRpcServer) AddLike(context.Context, *AddLikeReq) (*AddLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLike not implemented")
}
func (UnimplementedCommentRpcServer) CancelLike(context.Context, *CancelLikeReq) (*CancelLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelLike not implemented")
}
func (UnimplementedCommentRpcServer) IsLike(context.Context, *IsLikeReq) (*IsLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsLike not implemented")
}
func (UnimplementedCommentRpcServer) mustEmbedUnimplementedCommentRpcServer() {}

// UnsafeCommentRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentRpcServer will
// result in compilation errors.
type UnsafeCommentRpcServer interface {
	mustEmbedUnimplementedCommentRpcServer()
}

func RegisterCommentRpcServer(s grpc.ServiceRegistrar, srv CommentRpcServer) {
	s.RegisterService(&CommentRpc_ServiceDesc, srv)
}

func _CommentRpc_CreateSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).CreateSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.comment.v1.CommentRpc/CreateSubject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).CreateSubject(ctx, req.(*CreateSubjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_GetSubjectById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubjectByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).GetSubjectById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.comment.v1.CommentRpc/GetSubjectById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).GetSubjectById(ctx, req.(*GetSubjectByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_GetSubjectByBelong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubjectByBelongReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).GetSubjectByBelong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.comment.v1.CommentRpc/GetSubjectByBelong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).GetSubjectByBelong(ctx, req.(*GetSubjectByBelongReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_CreateReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).CreateReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.comment.v1.CommentRpc/CreateReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).CreateReply(ctx, req.(*CreateReplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_ListReplyByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReplyByIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).ListReplyByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.comment.v1.CommentRpc/ListReplyByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).ListReplyByIds(ctx, req.(*ListReplyByIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_ListFloorBySubjectSortReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFloorBySubjectSortReplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).ListFloorBySubjectSortReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.comment.v1.CommentRpc/ListFloorBySubjectSortReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).ListFloorBySubjectSortReply(ctx, req.(*ListFloorBySubjectSortReplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_ListFloorBySubjectSortTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFloorBySubjectSortTimeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).ListFloorBySubjectSortTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.comment.v1.CommentRpc/ListFloorBySubjectSortTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).ListFloorBySubjectSortTime(ctx, req.(*ListFloorBySubjectSortTimeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_ListReplyByFloorSortLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReplyByFloorSortLikeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).ListReplyByFloorSortLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.comment.v1.CommentRpc/ListReplyByFloorSortLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).ListReplyByFloorSortLike(ctx, req.(*ListReplyByFloorSortLikeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_ListReplyByFloorSortTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReplyByFloorSortTimeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).ListReplyByFloorSortTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.comment.v1.CommentRpc/ListReplyByFloorSortTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).ListReplyByFloorSortTime(ctx, req.(*ListReplyByFloorSortTimeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_AddLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLikeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).AddLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.comment.v1.CommentRpc/AddLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).AddLike(ctx, req.(*AddLikeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_CancelLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelLikeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).CancelLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.comment.v1.CommentRpc/CancelLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).CancelLike(ctx, req.(*CancelLikeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_IsLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsLikeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).IsLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.comment.v1.CommentRpc/IsLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).IsLike(ctx, req.(*IsLikeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentRpc_ServiceDesc is the grpc.ServiceDesc for CommentRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "domain.comment.v1.CommentRpc",
	HandlerType: (*CommentRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSubject",
			Handler:    _CommentRpc_CreateSubject_Handler,
		},
		{
			MethodName: "GetSubjectById",
			Handler:    _CommentRpc_GetSubjectById_Handler,
		},
		{
			MethodName: "GetSubjectByBelong",
			Handler:    _CommentRpc_GetSubjectByBelong_Handler,
		},
		{
			MethodName: "CreateReply",
			Handler:    _CommentRpc_CreateReply_Handler,
		},
		{
			MethodName: "ListReplyByIds",
			Handler:    _CommentRpc_ListReplyByIds_Handler,
		},
		{
			MethodName: "ListFloorBySubjectSortReply",
			Handler:    _CommentRpc_ListFloorBySubjectSortReply_Handler,
		},
		{
			MethodName: "ListFloorBySubjectSortTime",
			Handler:    _CommentRpc_ListFloorBySubjectSortTime_Handler,
		},
		{
			MethodName: "ListReplyByFloorSortLike",
			Handler:    _CommentRpc_ListReplyByFloorSortLike_Handler,
		},
		{
			MethodName: "ListReplyByFloorSortTime",
			Handler:    _CommentRpc_ListReplyByFloorSortTime_Handler,
		},
		{
			MethodName: "AddLike",
			Handler:    _CommentRpc_AddLike_Handler,
		},
		{
			MethodName: "CancelLike",
			Handler:    _CommentRpc_CancelLike_Handler,
		},
		{
			MethodName: "IsLike",
			Handler:    _CommentRpc_IsLike_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "domain/comment/v1/comment.proto",
}
