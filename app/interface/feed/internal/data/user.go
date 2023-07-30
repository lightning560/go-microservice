package data

import (
	"context"
	userv1 "redbook/api/domain/user/v1"
	basev1 "redbook/api/redbookpb/v1"
	"redbook/app/interface/feed/internal/biz"
	"redbook/app/interface/feed/internal/biz/model"
	"redbook/common/basemodel"

	"github.com/pkg/errors"
)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.IUserRepo {
	return &userRepo{data: data}
}

func (r *userRepo) GetUserInfoByUid(ctx context.Context, uid int64) (*basev1.UserInfo, error) {
	rv, err := r.data.userRpc.GetUserInfoByUid(ctx, &userv1.GetUserInfoByUidReq{Uid: uid})
	if err != nil {
		return nil, errors.Wrap(err, "GetUserInfoByUid")
	}
	return rv.Info, nil
}

func (r *userRepo) GetUserCardByUid(ctx context.Context, uid int64) (*model.UserProflie, error) {
	rv, err := r.data.userRpc.GetUserByUidCard(ctx, &userv1.GetUserCardByUidReq{Uid: uid})
	if err != nil {
		return nil, errors.Wrap(err, "GetUserCardByUid")
	}
	return &model.UserProflie{
		Uid:  rv.Card.Uid,
		Name: rv.Card.Name,
	}, nil
}

func (r *userRepo) GetUserProfileByUid(ctx context.Context, uid int64) (*model.UserProflie, error) {
	rv, err := r.data.userRpc.GetUserProfileByUid(ctx, &userv1.GetUserProfileByUidReq{Uid: uid})
	if err != nil {
		return nil, err
	}
	return &model.UserProflie{
		Uid:  rv.Profile.Uid,
		Name: rv.Profile.Name,
		Avatar: &basemodel.Image{
			Url:    rv.Profile.Avatar.Url,
			Width:  rv.Profile.Avatar.Width,
			Height: rv.Profile.Avatar.Height,
			SizeKb: rv.Profile.Avatar.SizeKb,
			Name:   rv.Profile.Avatar.Name,
			Hash:   rv.Profile.Avatar.Hash,
		},
	}, nil
}

func (r *userRepo) MapUserInfoByUids(ctx context.Context, uids []int64) (map[int64]*basemodel.UserInfo, error) {
	rv, err := r.data.userRpc.MapUserInfoByUids(ctx, &userv1.MapUserInfoByUidsReq{Uids: uids})
	if err != nil {
		return nil, err
	}
	var res = make(map[int64]*basemodel.UserInfo, rv.Total)
	for k, v := range rv.Infos {
		m := &basemodel.UserInfo{}
		m.FromProto(v)
		res[k] = m
	}
	return res, nil
}
