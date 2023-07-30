package server

import (
	"context"
	mallv1 "redbook/api/domain/mall/v1"
	"redbook/common/basemodel"

	"google.golang.org/grpc/status"
)

func (s *MallServer) CreateShop(ctx context.Context, req *mallv1.CreateShopReq) (*mallv1.CreateShopResponse, error) {
	logo := &basemodel.Image{}
	logo.FromProto(req.Logo)
	id, err := s.shopBiz.CreateShop(ctx, req.Uid, req.Name, req.Sign, logo)
	if err != nil {
		st := status.Errorf(100, "status error test: %v", err)
		return nil, status.Convert(st).Err()
	}
	return &mallv1.CreateShopResponse{Id: id}, nil
}

func (s *MallServer) GetShopById(ctx context.Context, req *mallv1.GetShopByIdReq) (*mallv1.GetShopByIdResponse, error) {
	m, err := s.shopBiz.GetShopById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &mallv1.GetShopByIdResponse{Shop: m.ToProto()}, nil
}
