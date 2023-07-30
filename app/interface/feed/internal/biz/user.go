package biz

import (
	"context"
	basev1 "redbook/api/redbookpb/v1"
	"redbook/app/interface/feed/internal/biz/model"
	"redbook/common/basemodel"

	"miopkg/errors"
)

type IUserRepo interface {
	GetUserInfoByUid(ctx context.Context, uid int64) (*basev1.UserInfo, error)
	GetUserCardByUid(ctx context.Context, uid int64) (*model.UserProflie, error)
	GetUserProfileByUid(ctx context.Context, uid int64) (*model.UserProflie, error)
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

func (u *UserBiz) GetUserCardByUid(ctx context.Context, uid int64) (*model.UserProflie, error) {
	return u.repo.GetUserCardByUid(ctx, uid)
}

func (u *UserBiz) GetUserProfileByUid(ctx context.Context, uid int64) (*model.UserProflie, error) {
	return u.repo.GetUserProfileByUid(ctx, uid)
}
func (u *UserBiz) MapUserInfoByUids(ctx context.Context, uids []int64) (map[int64]*basemodel.UserInfo, error) {
	return u.repo.MapUserInfoByUids(ctx, uids)
}
