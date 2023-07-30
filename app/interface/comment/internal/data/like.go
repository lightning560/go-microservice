package data

import (
	"context"
	commentv1 "redbook/api/domain/comment/v1"
	"redbook/app/interface/comment/internal/biz"
)

type likeRepo struct {
	data *Data
}

func NewLikeRepo(data *Data) biz.ILikeRepo {
	return &likeRepo{data: data}
}

// func (r *likeRepo) AddLike(c context.Context, replyId, uid int64) (err error) {
// 	if _, err = r.data.commentRpcClient.AddLike(c, &commentv1.AddLikeReq{
// 		ReplyId: replyId,
// 		Uid:     uid,
// 	}); err != nil {
// 		return err
// 	}
// 	return nil
// }

func (r *likeRepo) CancelLike(c context.Context, replyId, uid int64) error {
	_, err := r.data.commentRpcClient.CancelLike(c, &commentv1.CancelLikeReq{
		ReplyId: replyId,
		Uid:     uid,
	})
	if err != nil {
		return err
	}
	return err
}

func (r *likeRepo) IsLike(c context.Context, replyId, uid int64) (bool, error) {
	rv, err := r.data.commentRpcClient.IsLike(c, &commentv1.IsLikeReq{
		ReplyId: replyId,
		Uid:     uid,
	})
	if err != nil {
		return false, err
	}
	return rv.Result, err
}
