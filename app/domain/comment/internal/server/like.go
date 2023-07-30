package server

import (
	"context"
	commentv1 "redbook/api/domain/comment/v1"
)

func (s *CommentServer) AddLike(c context.Context, req *commentv1.AddLikeReq) (*commentv1.AddLikeResponse, error) {
	err := s.likebiz.AddLike(c, req.ReplyId, req.Uid)
	if err != nil {
		return nil, err
	}
	return &commentv1.AddLikeResponse{}, nil
}

func (s *CommentServer) CancelLike(c context.Context, req *commentv1.CancelLikeReq) (*commentv1.CancelLikeResponse, error) {
	err := s.likebiz.CancelLike(c, req.ReplyId, req.Uid)
	if err != nil {
		return nil, err
	}
	return &commentv1.CancelLikeResponse{}, nil
}

func (s *CommentServer) IsLike(c context.Context, req *commentv1.IsLikeReq) (*commentv1.IsLikeResponse, error) {
	rv, err := s.likebiz.IsLike(c, req.ReplyId, req.Uid)
	if err != nil {
		return nil, err
	}
	return &commentv1.IsLikeResponse{
		Result: rv,
	}, nil
}
