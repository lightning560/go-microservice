package server

import (
	"context"
	likev1 "redbook/api/domain/like/v1"
	"redbook/app/domain/like/internal/data/model"
)

func (s *LikeServer) Like(c context.Context, req *likev1.LikeReq) (*likev1.LikeReply, error) {
	m := &model.Like{
		Mid: req.Mid,
		Oid: req.Oid,
		Sid: req.Sid,
		Bid: req.Bid,
	}
	rv, err := s.likebz.Like(c, m)
	if err != nil {
		return nil, err
	}
	return &likev1.LikeReply{
		Mid:  rv.Mid,
		Oid:  rv.Oid,
		Like: rv.Like,
	}, nil
}
