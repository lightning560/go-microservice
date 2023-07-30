// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package domainmallv1

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

// MallRpcClient is the client API for MallRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MallRpcClient interface {
	CreateShop(ctx context.Context, in *CreateShopReq, opts ...grpc.CallOption) (*CreateShopResponse, error)
	GetShopById(ctx context.Context, in *GetShopByIdReq, opts ...grpc.CallOption) (*GetShopByIdResponse, error)
	CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResponse, error)
	GetProductById(ctx context.Context, in *GetProductByIdReq, opts ...grpc.CallOption) (*GetProductByIdResponse, error)
	MapProductByIds(ctx context.Context, in *MapProductByIdsReq, opts ...grpc.CallOption) (*MapProductByIdsResponse, error)
	CreateCollection(ctx context.Context, in *CreateCollectionReq, opts ...grpc.CallOption) (*CreateCollectionResponse, error)
	GetCollectionById(ctx context.Context, in *GetCollectionByIdReq, opts ...grpc.CallOption) (*GetCollectionByIdResponse, error)
	GetCollectionCardById(ctx context.Context, in *GetCollectionCardByIdReq, opts ...grpc.CallOption) (*GetCollectionCardByIdResponse, error)
	MapCollectionCardByIds(ctx context.Context, in *MapCollectionCardByIdsReq, opts ...grpc.CallOption) (*MapCollectionCardByIdsResponse, error)
	MapCollectionCardByShopId(ctx context.Context, in *MapCollectionCardByShopIdReq, opts ...grpc.CallOption) (*MapCollectionCardByShopIdResponse, error)
	UpdateCollectionState(ctx context.Context, in *UpdateCollectionStateReq, opts ...grpc.CallOption) (*UpdateCollectionStateResponse, error)
	UpdateCollectionSku(ctx context.Context, in *UpdateCollectionSkuReq, opts ...grpc.CallOption) (*UpdateCollectionSkuResponse, error)
	CreateCartItem(ctx context.Context, in *CreateCartItemReq, opts ...grpc.CallOption) (*CreateCartItemResponse, error)
	// rpc GetCartItem(GetCartItemReq) returns (GetCartItemResponse) {}
	MapCartItemByUid(ctx context.Context, in *MapCartItemByUidReq, opts ...grpc.CallOption) (*MapCartItemByUidResponse, error)
	UpdateCartItemQuantity(ctx context.Context, in *UpdateCartItemQuantityReq, opts ...grpc.CallOption) (*UpdateCartItemQuantityResponse, error)
	DeleteCartItem(ctx context.Context, in *DeleteCartItemReq, opts ...grpc.CallOption) (*DeleteCartItemResponse, error)
	CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	GetOrderById(ctx context.Context, in *GetOrderByIdReq, opts ...grpc.CallOption) (*GetOrderByIdResponse, error)
	MapOrderByIds(ctx context.Context, in *MapOrderByIdsReq, opts ...grpc.CallOption) (*MapOrderByIdsResponse, error)
	MapOrderByUid(ctx context.Context, in *MapOrderByUidReq, opts ...grpc.CallOption) (*MapOrderByUidResponse, error)
	UpdateOrderState(ctx context.Context, in *UpdateOrderStateReq, opts ...grpc.CallOption) (*UpdateOrderStateResponse, error)
}

type mallRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewMallRpcClient(cc grpc.ClientConnInterface) MallRpcClient {
	return &mallRpcClient{cc}
}

func (c *mallRpcClient) CreateShop(ctx context.Context, in *CreateShopReq, opts ...grpc.CallOption) (*CreateShopResponse, error) {
	out := new(CreateShopResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/CreateShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) GetShopById(ctx context.Context, in *GetShopByIdReq, opts ...grpc.CallOption) (*GetShopByIdResponse, error) {
	out := new(GetShopByIdResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/GetShopById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) GetProductById(ctx context.Context, in *GetProductByIdReq, opts ...grpc.CallOption) (*GetProductByIdResponse, error) {
	out := new(GetProductByIdResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/GetProductById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) MapProductByIds(ctx context.Context, in *MapProductByIdsReq, opts ...grpc.CallOption) (*MapProductByIdsResponse, error) {
	out := new(MapProductByIdsResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/MapProductByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) CreateCollection(ctx context.Context, in *CreateCollectionReq, opts ...grpc.CallOption) (*CreateCollectionResponse, error) {
	out := new(CreateCollectionResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/CreateCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) GetCollectionById(ctx context.Context, in *GetCollectionByIdReq, opts ...grpc.CallOption) (*GetCollectionByIdResponse, error) {
	out := new(GetCollectionByIdResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/GetCollectionById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) GetCollectionCardById(ctx context.Context, in *GetCollectionCardByIdReq, opts ...grpc.CallOption) (*GetCollectionCardByIdResponse, error) {
	out := new(GetCollectionCardByIdResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/GetCollectionCardById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) MapCollectionCardByIds(ctx context.Context, in *MapCollectionCardByIdsReq, opts ...grpc.CallOption) (*MapCollectionCardByIdsResponse, error) {
	out := new(MapCollectionCardByIdsResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/MapCollectionCardByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) MapCollectionCardByShopId(ctx context.Context, in *MapCollectionCardByShopIdReq, opts ...grpc.CallOption) (*MapCollectionCardByShopIdResponse, error) {
	out := new(MapCollectionCardByShopIdResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/MapCollectionCardByShopId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) UpdateCollectionState(ctx context.Context, in *UpdateCollectionStateReq, opts ...grpc.CallOption) (*UpdateCollectionStateResponse, error) {
	out := new(UpdateCollectionStateResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/UpdateCollectionState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) UpdateCollectionSku(ctx context.Context, in *UpdateCollectionSkuReq, opts ...grpc.CallOption) (*UpdateCollectionSkuResponse, error) {
	out := new(UpdateCollectionSkuResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/UpdateCollectionSku", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) CreateCartItem(ctx context.Context, in *CreateCartItemReq, opts ...grpc.CallOption) (*CreateCartItemResponse, error) {
	out := new(CreateCartItemResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/CreateCartItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) MapCartItemByUid(ctx context.Context, in *MapCartItemByUidReq, opts ...grpc.CallOption) (*MapCartItemByUidResponse, error) {
	out := new(MapCartItemByUidResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/MapCartItemByUid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) UpdateCartItemQuantity(ctx context.Context, in *UpdateCartItemQuantityReq, opts ...grpc.CallOption) (*UpdateCartItemQuantityResponse, error) {
	out := new(UpdateCartItemQuantityResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/UpdateCartItemQuantity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) DeleteCartItem(ctx context.Context, in *DeleteCartItemReq, opts ...grpc.CallOption) (*DeleteCartItemResponse, error) {
	out := new(DeleteCartItemResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/DeleteCartItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) GetOrderById(ctx context.Context, in *GetOrderByIdReq, opts ...grpc.CallOption) (*GetOrderByIdResponse, error) {
	out := new(GetOrderByIdResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/GetOrderById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) MapOrderByIds(ctx context.Context, in *MapOrderByIdsReq, opts ...grpc.CallOption) (*MapOrderByIdsResponse, error) {
	out := new(MapOrderByIdsResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/MapOrderByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) MapOrderByUid(ctx context.Context, in *MapOrderByUidReq, opts ...grpc.CallOption) (*MapOrderByUidResponse, error) {
	out := new(MapOrderByUidResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/MapOrderByUid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallRpcClient) UpdateOrderState(ctx context.Context, in *UpdateOrderStateReq, opts ...grpc.CallOption) (*UpdateOrderStateResponse, error) {
	out := new(UpdateOrderStateResponse)
	err := c.cc.Invoke(ctx, "/domain.mall.v1.MallRpc/UpdateOrderState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MallRpcServer is the server API for MallRpc service.
// All implementations must embed UnimplementedMallRpcServer
// for forward compatibility
type MallRpcServer interface {
	CreateShop(context.Context, *CreateShopReq) (*CreateShopResponse, error)
	GetShopById(context.Context, *GetShopByIdReq) (*GetShopByIdResponse, error)
	CreateProduct(context.Context, *CreateProductReq) (*CreateProductResponse, error)
	GetProductById(context.Context, *GetProductByIdReq) (*GetProductByIdResponse, error)
	MapProductByIds(context.Context, *MapProductByIdsReq) (*MapProductByIdsResponse, error)
	CreateCollection(context.Context, *CreateCollectionReq) (*CreateCollectionResponse, error)
	GetCollectionById(context.Context, *GetCollectionByIdReq) (*GetCollectionByIdResponse, error)
	GetCollectionCardById(context.Context, *GetCollectionCardByIdReq) (*GetCollectionCardByIdResponse, error)
	MapCollectionCardByIds(context.Context, *MapCollectionCardByIdsReq) (*MapCollectionCardByIdsResponse, error)
	MapCollectionCardByShopId(context.Context, *MapCollectionCardByShopIdReq) (*MapCollectionCardByShopIdResponse, error)
	UpdateCollectionState(context.Context, *UpdateCollectionStateReq) (*UpdateCollectionStateResponse, error)
	UpdateCollectionSku(context.Context, *UpdateCollectionSkuReq) (*UpdateCollectionSkuResponse, error)
	CreateCartItem(context.Context, *CreateCartItemReq) (*CreateCartItemResponse, error)
	// rpc GetCartItem(GetCartItemReq) returns (GetCartItemResponse) {}
	MapCartItemByUid(context.Context, *MapCartItemByUidReq) (*MapCartItemByUidResponse, error)
	UpdateCartItemQuantity(context.Context, *UpdateCartItemQuantityReq) (*UpdateCartItemQuantityResponse, error)
	DeleteCartItem(context.Context, *DeleteCartItemReq) (*DeleteCartItemResponse, error)
	CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderResponse, error)
	GetOrderById(context.Context, *GetOrderByIdReq) (*GetOrderByIdResponse, error)
	MapOrderByIds(context.Context, *MapOrderByIdsReq) (*MapOrderByIdsResponse, error)
	MapOrderByUid(context.Context, *MapOrderByUidReq) (*MapOrderByUidResponse, error)
	UpdateOrderState(context.Context, *UpdateOrderStateReq) (*UpdateOrderStateResponse, error)
	mustEmbedUnimplementedMallRpcServer()
}

// UnimplementedMallRpcServer must be embedded to have forward compatible implementations.
type UnimplementedMallRpcServer struct {
}

func (UnimplementedMallRpcServer) CreateShop(context.Context, *CreateShopReq) (*CreateShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShop not implemented")
}
func (UnimplementedMallRpcServer) GetShopById(context.Context, *GetShopByIdReq) (*GetShopByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopById not implemented")
}
func (UnimplementedMallRpcServer) CreateProduct(context.Context, *CreateProductReq) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedMallRpcServer) GetProductById(context.Context, *GetProductByIdReq) (*GetProductByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductById not implemented")
}
func (UnimplementedMallRpcServer) MapProductByIds(context.Context, *MapProductByIdsReq) (*MapProductByIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapProductByIds not implemented")
}
func (UnimplementedMallRpcServer) CreateCollection(context.Context, *CreateCollectionReq) (*CreateCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollection not implemented")
}
func (UnimplementedMallRpcServer) GetCollectionById(context.Context, *GetCollectionByIdReq) (*GetCollectionByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollectionById not implemented")
}
func (UnimplementedMallRpcServer) GetCollectionCardById(context.Context, *GetCollectionCardByIdReq) (*GetCollectionCardByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollectionCardById not implemented")
}
func (UnimplementedMallRpcServer) MapCollectionCardByIds(context.Context, *MapCollectionCardByIdsReq) (*MapCollectionCardByIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapCollectionCardByIds not implemented")
}
func (UnimplementedMallRpcServer) MapCollectionCardByShopId(context.Context, *MapCollectionCardByShopIdReq) (*MapCollectionCardByShopIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapCollectionCardByShopId not implemented")
}
func (UnimplementedMallRpcServer) UpdateCollectionState(context.Context, *UpdateCollectionStateReq) (*UpdateCollectionStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCollectionState not implemented")
}
func (UnimplementedMallRpcServer) UpdateCollectionSku(context.Context, *UpdateCollectionSkuReq) (*UpdateCollectionSkuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCollectionSku not implemented")
}
func (UnimplementedMallRpcServer) CreateCartItem(context.Context, *CreateCartItemReq) (*CreateCartItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCartItem not implemented")
}
func (UnimplementedMallRpcServer) MapCartItemByUid(context.Context, *MapCartItemByUidReq) (*MapCartItemByUidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapCartItemByUid not implemented")
}
func (UnimplementedMallRpcServer) UpdateCartItemQuantity(context.Context, *UpdateCartItemQuantityReq) (*UpdateCartItemQuantityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCartItemQuantity not implemented")
}
func (UnimplementedMallRpcServer) DeleteCartItem(context.Context, *DeleteCartItemReq) (*DeleteCartItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCartItem not implemented")
}
func (UnimplementedMallRpcServer) CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedMallRpcServer) GetOrderById(context.Context, *GetOrderByIdReq) (*GetOrderByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderById not implemented")
}
func (UnimplementedMallRpcServer) MapOrderByIds(context.Context, *MapOrderByIdsReq) (*MapOrderByIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapOrderByIds not implemented")
}
func (UnimplementedMallRpcServer) MapOrderByUid(context.Context, *MapOrderByUidReq) (*MapOrderByUidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapOrderByUid not implemented")
}
func (UnimplementedMallRpcServer) UpdateOrderState(context.Context, *UpdateOrderStateReq) (*UpdateOrderStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderState not implemented")
}
func (UnimplementedMallRpcServer) mustEmbedUnimplementedMallRpcServer() {}

// UnsafeMallRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MallRpcServer will
// result in compilation errors.
type UnsafeMallRpcServer interface {
	mustEmbedUnimplementedMallRpcServer()
}

func RegisterMallRpcServer(s grpc.ServiceRegistrar, srv MallRpcServer) {
	s.RegisterService(&MallRpc_ServiceDesc, srv)
}

func _MallRpc_CreateShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).CreateShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/CreateShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).CreateShop(ctx, req.(*CreateShopReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_GetShopById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShopByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).GetShopById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/GetShopById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).GetShopById(ctx, req.(*GetShopByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).CreateProduct(ctx, req.(*CreateProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_GetProductById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).GetProductById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/GetProductById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).GetProductById(ctx, req.(*GetProductByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_MapProductByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapProductByIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).MapProductByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/MapProductByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).MapProductByIds(ctx, req.(*MapProductByIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_CreateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCollectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).CreateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/CreateCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).CreateCollection(ctx, req.(*CreateCollectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_GetCollectionById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollectionByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).GetCollectionById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/GetCollectionById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).GetCollectionById(ctx, req.(*GetCollectionByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_GetCollectionCardById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollectionCardByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).GetCollectionCardById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/GetCollectionCardById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).GetCollectionCardById(ctx, req.(*GetCollectionCardByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_MapCollectionCardByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapCollectionCardByIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).MapCollectionCardByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/MapCollectionCardByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).MapCollectionCardByIds(ctx, req.(*MapCollectionCardByIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_MapCollectionCardByShopId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapCollectionCardByShopIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).MapCollectionCardByShopId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/MapCollectionCardByShopId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).MapCollectionCardByShopId(ctx, req.(*MapCollectionCardByShopIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_UpdateCollectionState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCollectionStateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).UpdateCollectionState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/UpdateCollectionState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).UpdateCollectionState(ctx, req.(*UpdateCollectionStateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_UpdateCollectionSku_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCollectionSkuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).UpdateCollectionSku(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/UpdateCollectionSku",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).UpdateCollectionSku(ctx, req.(*UpdateCollectionSkuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_CreateCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCartItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).CreateCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/CreateCartItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).CreateCartItem(ctx, req.(*CreateCartItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_MapCartItemByUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapCartItemByUidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).MapCartItemByUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/MapCartItemByUid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).MapCartItemByUid(ctx, req.(*MapCartItemByUidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_UpdateCartItemQuantity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCartItemQuantityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).UpdateCartItemQuantity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/UpdateCartItemQuantity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).UpdateCartItemQuantity(ctx, req.(*UpdateCartItemQuantityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_DeleteCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCartItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).DeleteCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/DeleteCartItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).DeleteCartItem(ctx, req.(*DeleteCartItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).CreateOrder(ctx, req.(*CreateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_GetOrderById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).GetOrderById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/GetOrderById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).GetOrderById(ctx, req.(*GetOrderByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_MapOrderByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapOrderByIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).MapOrderByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/MapOrderByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).MapOrderByIds(ctx, req.(*MapOrderByIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_MapOrderByUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapOrderByUidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).MapOrderByUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/MapOrderByUid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).MapOrderByUid(ctx, req.(*MapOrderByUidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MallRpc_UpdateOrderState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderStateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallRpcServer).UpdateOrderState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/domain.mall.v1.MallRpc/UpdateOrderState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallRpcServer).UpdateOrderState(ctx, req.(*UpdateOrderStateReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MallRpc_ServiceDesc is the grpc.ServiceDesc for MallRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MallRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "domain.mall.v1.MallRpc",
	HandlerType: (*MallRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShop",
			Handler:    _MallRpc_CreateShop_Handler,
		},
		{
			MethodName: "GetShopById",
			Handler:    _MallRpc_GetShopById_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _MallRpc_CreateProduct_Handler,
		},
		{
			MethodName: "GetProductById",
			Handler:    _MallRpc_GetProductById_Handler,
		},
		{
			MethodName: "MapProductByIds",
			Handler:    _MallRpc_MapProductByIds_Handler,
		},
		{
			MethodName: "CreateCollection",
			Handler:    _MallRpc_CreateCollection_Handler,
		},
		{
			MethodName: "GetCollectionById",
			Handler:    _MallRpc_GetCollectionById_Handler,
		},
		{
			MethodName: "GetCollectionCardById",
			Handler:    _MallRpc_GetCollectionCardById_Handler,
		},
		{
			MethodName: "MapCollectionCardByIds",
			Handler:    _MallRpc_MapCollectionCardByIds_Handler,
		},
		{
			MethodName: "MapCollectionCardByShopId",
			Handler:    _MallRpc_MapCollectionCardByShopId_Handler,
		},
		{
			MethodName: "UpdateCollectionState",
			Handler:    _MallRpc_UpdateCollectionState_Handler,
		},
		{
			MethodName: "UpdateCollectionSku",
			Handler:    _MallRpc_UpdateCollectionSku_Handler,
		},
		{
			MethodName: "CreateCartItem",
			Handler:    _MallRpc_CreateCartItem_Handler,
		},
		{
			MethodName: "MapCartItemByUid",
			Handler:    _MallRpc_MapCartItemByUid_Handler,
		},
		{
			MethodName: "UpdateCartItemQuantity",
			Handler:    _MallRpc_UpdateCartItemQuantity_Handler,
		},
		{
			MethodName: "DeleteCartItem",
			Handler:    _MallRpc_DeleteCartItem_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _MallRpc_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrderById",
			Handler:    _MallRpc_GetOrderById_Handler,
		},
		{
			MethodName: "MapOrderByIds",
			Handler:    _MallRpc_MapOrderByIds_Handler,
		},
		{
			MethodName: "MapOrderByUid",
			Handler:    _MallRpc_MapOrderByUid_Handler,
		},
		{
			MethodName: "UpdateOrderState",
			Handler:    _MallRpc_UpdateOrderState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "domain/mall/v1/mall.proto",
}
