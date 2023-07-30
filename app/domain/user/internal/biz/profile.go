package biz

import (
	"context"
	"miopkg/errors"
	"redbook/app/domain/user/model"
	"redbook/common/basemodel"
)

type IProfileRepo interface {
	UpdateProfile(context.Context, *model.Profile) error
	// GetUserInfoByUid(context.Context, int64) (*model.Profile, error)
	// ListUserInfoByUids(context.Context, []int64) ([]*model.Profile, error)
	// GetUserCardByUid(context.Context, int64) (*model.Profile, error)
	GetUserProfileByUid(context.Context, int64) (*model.Profile, error)
	ListUserProfileByUids(context.Context, []int64) ([]*model.Profile, error)
}

type ProfileBiz struct {
	repo IProfileRepo
}

func NewProfileBiz(repo IProfileRepo) *ProfileBiz {
	return &ProfileBiz{repo: repo}
}
func (bz *ProfileBiz) UpdateProfile(ctx context.Context, p *model.Profile) error {
	return bz.repo.UpdateProfile(ctx, p)
}

func (bz *ProfileBiz) GetUserInfoByUid(ctx context.Context, uid int64) (*basemodel.UserInfo, error) {
	rv, err := bz.repo.GetUserProfileByUid(ctx, uid)
	if err != nil {
		return nil, err
	}
	return rv.ToInfo(), nil
}
func (bz *ProfileBiz) ListUserInfoByUids(ctx context.Context, uids []int64) ([]*basemodel.UserInfo, error) {
	rv, err := bz.repo.ListUserProfileByUids(ctx, uids)
	if err != nil {
		return nil, errors.Wrap(err, "biz.ListUserInfoByUids")
	}
	var res []*basemodel.UserInfo
	for _, v := range rv {
		res = append(res, v.ToInfo())
	}
	return res, nil
}

func (bz *ProfileBiz) GetUserCardByUid(ctx context.Context, uid int64) (*model.Profile, error) {
	return bz.repo.GetUserProfileByUid(ctx, uid)
}
func (bz *ProfileBiz) GetUserProfileByUid(ctx context.Context, uid int64) (*model.Profile, error) {
	return bz.repo.GetUserProfileByUid(ctx, uid)
}
