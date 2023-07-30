package server

import (
	"context"
	feedv1 "redbook/api/domain/feed/v1"
	basev1 "redbook/api/redbookpb/v1"
)

func (s *FeedServer) AddLike(ctx context.Context, req *feedv1.AddLikeReq) (*basev1.EmptyResponse, error) {
	err := s.likebiz.AddLike(ctx, req.Uid, req.PostId)
	if err != nil {
		return nil, err
	}
	return &basev1.EmptyResponse{}, nil
}

func (s *FeedServer) CancelLike(ctx context.Context, req *feedv1.CancelLikeReq) (*basev1.EmptyResponse, error) {
	err := s.likebiz.CancelLike(ctx, req.Uid, req.PostId)
	if err != nil {
		return nil, err
	}
	return &basev1.EmptyResponse{}, nil
}

func (s *FeedServer) IsLike(ctx context.Context, req *feedv1.IsLikeReq) (*basev1.BoolResponse, error) {
	rv, err := s.likebiz.IsLike(ctx, req.Uid, req.PostId)
	if err != nil {
		return nil, err
	}
	return &basev1.BoolResponse{Result: rv}, nil
}
func (s *FeedServer) ListLikeByPost(ctx context.Context, req *feedv1.ListLikeByPostReq) (*feedv1.ListLikeByPostResponse, error) {
	rv, err := s.likebiz.ListLikeByPost(ctx, req.PostId)
	if err != nil {
		return nil, err
	}
	res := &feedv1.ListLikeByPostResponse{
		Likes: make([]*feedv1.Like, 0, len(*rv)),
		Total: int32(len(*rv)),
	}
	for _, v := range *rv {
		res.Likes = append(res.Likes, &feedv1.Like{
			Uid:       v.Uid,
			PostId:    v.PostId,
			CreatedAt: v.CreatedAt,
		})
	}
	return res, nil
}

func (s *FeedServer) ListLikeByUid(ctx context.Context, req *feedv1.ListLikeByUidReq) (*feedv1.ListLikeByUidResponse, error) {
	rv, err := s.likebiz.ListLikeByUid(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	res := &feedv1.ListLikeByUidResponse{
		Likes: make([]*feedv1.Like, 0, len(*rv)),
		Total: int32(len(*rv)),
	}
	for _, v := range *rv {
		res.Likes = append(res.Likes, &feedv1.Like{
			Uid:       v.Uid,
			PostId:    v.PostId,
			CreatedAt: v.CreatedAt,
		})
	}
	return res, nil
}

func (s *FeedServer) GetLikeCountByPost(ctx context.Context, req *feedv1.GetLikeCountByPostReq) (*feedv1.GetLikeCountByPostResponse, error) {
	rv, err := s.likebiz.GetLikeCountByPost(ctx, req.PostId)
	if err != nil {
		return nil, err
	}
	return &feedv1.GetLikeCountByPostResponse{Count: rv}, nil
}
