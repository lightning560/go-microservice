package server

import (
	"context"
	feedv1 "redbook/api/domain/feed/v1"
	basev1 "redbook/api/redbookpb/v1"
)

func (s *FeedServer) AddFollow(ctx context.Context, req *feedv1.AddFollowReq) (*basev1.EmptyResponse, error) {
	err := s.relationbiz.AddFollow(ctx, req.FollowerUid, req.FolloweeUid)
	if err != nil {
		return nil, err
	}
	return &basev1.EmptyResponse{}, nil
}

func (s *FeedServer) CancelFollow(ctx context.Context, req *feedv1.CancelFollowReq) (*basev1.EmptyResponse, error) {
	err := s.relationbiz.CancelFollow(ctx, req.FollowerUid, req.FolloweeUid)
	if err != nil {
		return nil, err
	}
	return &basev1.EmptyResponse{}, nil
}

func (s *FeedServer) IsFollow(ctx context.Context, req *feedv1.IsFollowReq) (*basev1.BoolResponse, error) {
	rv, err := s.relationbiz.IsFollow(ctx, req.FollowerUid, req.FolloweeUid)
	if err != nil {
		return nil, err
	}
	return &basev1.BoolResponse{Result: rv}, nil
}

func (s *FeedServer) ListFollower(ctx context.Context, req *feedv1.ListFollowerReq) (*feedv1.ListFollowerResponse, error) {
	rv, err := s.relationbiz.ListFollower(ctx, req.FolloweeUid)
	if err != nil {
		return nil, err
	}
	res := &feedv1.ListFollowerResponse{
		Relations: make([]*feedv1.Relation, 0, len(*rv)),
		Total:     int32(len(*rv)),
	}
	for _, v := range *rv {
		res.Relations = append(res.Relations, &feedv1.Relation{
			FollowerUid: v.FollowerUid,
			FolloweeUid: v.FolloweeUid,
		})
	}
	return res, nil
}

func (s *FeedServer) ListFollowee(ctx context.Context, req *feedv1.ListFolloweeReq) (*feedv1.ListFolloweeResponse, error) {
	rv, err := s.relationbiz.ListFollowee(ctx, req.FollowerUid)
	if err != nil {
		return nil, err
	}
	res := &feedv1.ListFolloweeResponse{
		Relations: make([]*feedv1.Relation, 0, len(*rv)),
		Total:     int32(len(*rv)),
	}
	for _, v := range *rv {
		res.Relations = append(res.Relations, &feedv1.Relation{
			FollowerUid: v.FollowerUid,
			FolloweeUid: v.FolloweeUid,
		})
	}
	return res, nil
}
