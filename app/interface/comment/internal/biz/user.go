package biz

import (
	"context"
	"miopkg/errors"
	basev1 "redbook/api/redbookpb/v1"
	"redbook/common/basemodel"
)

type IUserRepo interface {
	GetUserInfoByUid(ctx context.Context, uid int64) (*basev1.UserInfo, error)
	MapUserInfoByUids(ctx context.Context, uids []int64) (map[int64]*basemodel.UserInfo, error)
}

type UserBiz struct {
	repo IUserRepo
}

func NewUserBiz(repo IUserRepo) *UserBiz {
	return &UserBiz{repo: repo}
}
func (bz *UserBiz) GetUserInfoByUid(ctx context.Context, uid int64) (*basemodel.UserInfo, error) {
	rv, err := bz.repo.GetUserInfoByUid(ctx, uid)
	if err != nil {
		return nil, errors.Wrap(err, "GetUserInfoByUid")
	}
	res := &basemodel.UserInfo{}
	res.FromProto(rv)
	return res, nil
}

func (u *UserBiz) MapUserInfoByUids(ctx context.Context, uids []int64) (map[int64]*basemodel.UserInfo, error) {
	return u.repo.MapUserInfoByUids(ctx, uids)
}
