package server

import (
	"context"
	mallv1 "redbook/api/domain/mall/v1"
)

func (s *MallServer) CreateOrder(ctx context.Context, req *mallv1.CreateOrderReq) (*mallv1.CreateOrderResponse, error) {
	return &mallv1.CreateOrderResponse{}, nil
}
