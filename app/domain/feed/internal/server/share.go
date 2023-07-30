package server

import (
	"context"
	feedv1 "redbook/api/domain/feed/v1"
	basev1 "redbook/api/redbookpb/v1"
)

func (s *FeedServer) AddShare(ctx context.Context, req *feedv1.AddShareReq) (*basev1.EmptyResponse, error) {
	err := s.sharebiz.AddShare(ctx, req.Uid, req.PostId)
	if err != nil {
		return nil, err
	}
	return &basev1.EmptyResponse{}, nil
}

func (s *FeedServer) CancelShare(ctx context.Context, req *feedv1.CancelShareReq) (*basev1.EmptyResponse, error) {
	err := s.sharebiz.CancelShare(ctx, req.Uid, req.PostId)
	if err != nil {
		return nil, err
	}
	return &basev1.EmptyResponse{}, nil
}

func (s *FeedServer) IsShare(ctx context.Context, req *feedv1.IsShareReq) (*basev1.BoolResponse, error) {
	rv, err := s.sharebiz.IsShare(ctx, req.Uid, req.PostId)
	if err != nil {
		return nil, err
	}
	return &basev1.BoolResponse{Result: rv}, nil
}

func (s *FeedServer) ListShareByPost(ctx context.Context, req *feedv1.ListShareByPostReq) (*feedv1.ListShareByPostResponse, error) {
	rv, err := s.sharebiz.ListShareByPost(ctx, req.PostId)
	if err != nil {
		return nil, err
	}
	res := &feedv1.ListShareByPostResponse{
		Shares: make([]*feedv1.Share, 0, len(*rv)),
		Total:  int32(len(*rv)),
	}
	for _, v := range *rv {

		res.Shares = append(res.Shares, &feedv1.Share{
			Uid:       v.Uid,
			PostId:    v.PostId,
			CreatedAt: v.CreatedAt,
		})
	}
	return res, nil
}

func (s *FeedServer) ListShareByUid(ctx context.Context, req *feedv1.ListShareByUidReq) (*feedv1.ListShareByUidResponse, error) {
	rv, err := s.sharebiz.ListShareByUid(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	res := &feedv1.ListShareByUidResponse{
		Shares: make([]*feedv1.Share, 0, len(*rv)),
		Total:  int32(len(*rv)),
	}
	for _, v := range *rv {
		res.Shares = append(res.Shares, &feedv1.Share{
			Uid:       v.Uid,
			PostId:    v.PostId,
			CreatedAt: v.CreatedAt,
		})
	}
	return res, nil
}

func (s *FeedServer) GetShareCountByPost(ctx context.Context, req *feedv1.GetShareCountByPostReq) (*feedv1.GetShareCountByPostResponse, error) {
	rv, err := s.sharebiz.GetShareCountByPost(ctx, req.PostId)
	if err != nil {
		return nil, err
	}
	return &feedv1.GetShareCountByPostResponse{Count: rv}, nil
}
