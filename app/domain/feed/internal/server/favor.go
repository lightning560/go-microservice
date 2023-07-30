package server

import (
	"context"
	feedv1 "redbook/api/domain/feed/v1"
	basev1 "redbook/api/redbookpb/v1"
)

func (s *FeedServer) AddFavor(ctx context.Context, req *feedv1.AddFavorReq) (*basev1.EmptyResponse, error) {
	err := s.favorbiz.AddFavor(ctx, req.Uid, req.PostId)
	if err != nil {
		return nil, err
	}
	return &basev1.EmptyResponse{}, nil
}

func (s *FeedServer) CancelFavor(ctx context.Context, req *feedv1.CancelFavorReq) (*basev1.EmptyResponse, error) {
	err := s.favorbiz.CancelFavor(ctx, req.Uid, req.PostId)
	if err != nil {
		return nil, err
	}
	return &basev1.EmptyResponse{}, nil
}

func (s *FeedServer) IsFavor(ctx context.Context, req *feedv1.IsFavorReq) (*basev1.BoolResponse, error) {
	rv, err := s.favorbiz.IsFavor(ctx, req.Uid, req.PostId)
	if err != nil {
		return nil, err
	}
	return &basev1.BoolResponse{Result: rv}, nil
}

func (s *FeedServer) ListFavorByPost(ctx context.Context, req *feedv1.ListFavorByPostReq) (*feedv1.ListFavorByPostResponse, error) {
	rv, err := s.favorbiz.ListFavorByPost(ctx, req.PostId)
	if err != nil {
		return nil, err
	}
	res := &feedv1.ListFavorByPostResponse{
		Favors: make([]*feedv1.Favor, 0, len(*rv)),
		Total:  int32(len(*rv)),
	}
	for _, v := range *rv {
		res.Favors = append(res.Favors, &feedv1.Favor{
			Uid:       v.Uid,
			PostId:    v.PostId,
			CreatedAt: v.CreatedAt,
		})
	}
	return res, nil
}

func (s *FeedServer) ListFavorByUid(ctx context.Context, req *feedv1.ListFavorByUidReq) (*feedv1.ListFavorByUidResponse, error) {
	rv, err := s.favorbiz.ListFavorByUid(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	res := &feedv1.ListFavorByUidResponse{
		Favors: make([]*feedv1.Favor, 0, len(*rv)),
		Total:  int32(len(*rv)),
	}
	for _, v := range *rv {
		res.Favors = append(res.Favors, &feedv1.Favor{
			Uid:       v.Uid,
			PostId:    v.PostId,
			CreatedAt: v.CreatedAt,
		})
	}
	return res, nil
}

func (s *FeedServer) GetFavorCountByPost(ctx context.Context, req *feedv1.GetFavorCountByPostReq) (*feedv1.GetFavorCountByPostResponse, error) {
	rv, err := s.favorbiz.GetFavorCountByPost(ctx, req.PostId)
	if err != nil {
		return nil, err
	}
	return &feedv1.GetFavorCountByPostResponse{Count: rv}, nil
}
