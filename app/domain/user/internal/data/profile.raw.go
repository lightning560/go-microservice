package data

import (
	"context"
	"fmt"
	"miopkg/log"
	"redbook/app/domain/user/internal/data/ent/profile"
	"redbook/app/domain/user/model"
	"redbook/common/basemodel"

	"miopkg/errors"
)

func (r *profileRepo) RawGetUserInfoByUid(ctx context.Context, uid int64) (*model.Profile, error) {
	rv, err := r.data.userWDB.Debug().Profile.Query().Where(profile.UID(uid)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Profile{
		Id:   rv.ID,
		Uid:  rv.UID,
		Name: rv.Name,
		Avatar: &basemodel.Image{
			Url: rv.Avatar,
		},
		Sign:  rv.Sign,
		Level: rv.Level,
		Sex:   rv.Sex,
	}, nil
}

func (r *profileRepo) RawListUserInfoByUids(ctx context.Context, uids []int64) ([]*model.Profile, error) {
	rv, err := r.data.userWDB.Profile.Query().Where(profile.UIDIn(uids...)).All(ctx)
	if err != nil {
		return nil, err
	}
	var res []*model.Profile
	for _, v := range rv {
		res = append(res, &model.Profile{
			Id:   v.ID,
			Uid:  v.UID,
			Name: v.Name,
			Avatar: &basemodel.Image{
				Url: v.Avatar,
			},
			Level: v.Level,
		})
	}
	return res, nil
}

func (r *profileRepo) RawGetUserCardByUid(ctx context.Context, uid int64) (*model.Profile, error) {
	rv, err := r.data.userWDB.Profile.Query().Where(profile.UID(uid)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Profile{
		Id:   rv.ID,
		Uid:  rv.UID,
		Name: rv.Name,
		Avatar: &basemodel.Image{
			Url: rv.Avatar,
		},
		Sign:  rv.Sign,
		Level: rv.Level,
		State: rv.State,
	}, nil
}

func (r *profileRepo) RawGetUserProfileByUid(ctx context.Context, uid int64) (*model.Profile, error) {
	rv, err := r.data.userWDB.Debug().Profile.Query().Where(profile.UID(uid)).Only(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "RawGetUserProfileByUid")
	}
	return &model.Profile{
		Id:   rv.ID,
		Uid:  rv.UID,
		Name: rv.Name,
		Avatar: &basemodel.Image{
			Url: rv.Avatar,
		},
		Sign:  rv.Sign,
		Level: rv.Level,
		State: rv.State,
	}, nil
}

func (r *profileRepo) RawListUserProfileByUids(ctx context.Context, uids []int64) ([]*model.Profile, error) {
	rv, err := r.data.userWDB.Profile.Query().Where(profile.UIDIn(uids...)).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(rv) == 0 {
		return nil, errors.ErrNoDocuments
	}
	var res []*model.Profile
	for _, v := range rv {
		res = append(res, &model.Profile{
			Id:   v.ID,
			Uid:  v.UID,
			Name: v.Name,
			Avatar: &basemodel.Image{
				Url: v.Avatar,
			},
			Level: v.Level,
		})
	}
	return res, nil
}
func (r *profileRepo) RawUpdateProfile(ctx context.Context, m *model.Profile) error {
	res, err := r.data.userWDB.Profile.UpdateOneID(m.Id).SetUID(m.Uid).SetName(m.Name).SetAvatar(m.Avatar.Url).SetSign(m.Sign).SetLevel(m.Level).SetSex(m.Sex).Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating Profile: %w", err)
	}
	log.Info(fmt.Sprintf("Profile was created:%p", res), log.FieldMod("user"))
	return nil
}
