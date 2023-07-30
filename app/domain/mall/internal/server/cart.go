package server

import (
	"context"
	"fmt"
	mallv1 "redbook/api/domain/mall/v1"
	"redbook/app/domain/mall/model"
)

func (s *MallServer) CreateCartItem(c context.Context, req *mallv1.CreateCartItemReq) (*mallv1.CreateCartItemResponse, error) {
	rv, err := s.cartBiz.CreateCartItem(c, req.Uid, req.ProductId, req.CollectionId, req.Quantity)
	if err != nil {
		return nil, err
	}
	return &mallv1.CreateCartItemResponse{
		Id: rv,
	}, nil
}

func (s *MallServer) MapCartItemByUid(c context.Context, req *mallv1.MapCartItemByUidReq) (*mallv1.MapCartItemByUidResponse, error) {
	rv, err := s.cartBiz.ListCartItemByUid(c, req.Uid)
	if err != nil {
		return nil, err
	}
	fmt.Println("svr#MapCartItemByUid rv: ", rv)
	pids := make([]int64, 0)
	for _, v := range rv {
		pids = append(pids, v.ProductId)
	}
	fmt.Println("svr#MapCartItemByUid pids: ", pids)
	productList, err := s.productBiz.ListProductByIds(c, pids)
	if err != nil {
		return nil, err
	}
	productMap := make(map[int64]*model.Product)
	for _, v := range productList {
		productMap[v.Id] = v
	}
	res := rv.ListToProto()
	for _, v := range res {
		v.ProductCard = productMap[v.ProductId].ToCard().ToProto()
	}
	mapCartItem := make(map[int64]*mallv1.CartItem)
	for _, v := range res {
		mapCartItem[v.Id] = v
	}
	return &mallv1.MapCartItemByUidResponse{
		CartItems: mapCartItem,
	}, nil
}

func (s *MallServer) UpdateCartItemQuantity(c context.Context, req *mallv1.UpdateCartItemQuantityReq) (*mallv1.UpdateCartItemQuantityResponse, error) {
	err := s.cartBiz.UpdateCartItemQuantity(c, req.Id, req.Quantity)
	if err != nil {
		return nil, err
	}
	return &mallv1.UpdateCartItemQuantityResponse{}, nil
}
