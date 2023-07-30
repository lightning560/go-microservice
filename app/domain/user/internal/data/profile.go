package data

import (
	"context"
	"fmt"
	"miopkg/client/eredis"
	"miopkg/errors"
	"miopkg/util/xlist"
	"redbook/app/domain/user/internal/biz"
	"redbook/app/domain/user/model"
	"redbook/common/basemodel"

	"golang.org/x/sync/singleflight"
)

type profileRepo struct {
	data *Data
}

func NewProfileRepo(data *Data) biz.IProfileRepo {
	return &profileRepo{data: data}
}

var cacheSingleFlights = [1]*singleflight.Group{{}}

func (r *profileRepo) UpdateProfile(ctx context.Context, m *model.Profile) error {
	err := r.RawUpdateProfile(ctx, m)
	if err != nil {
		return err
	}
	r.data.fanout.Do(ctx, func(ctx context.Context) {
		r.CacheUpdateProfile(ctx, m.Uid, m)
	})
	return nil
}

func (r *profileRepo) GetUserInfoByUid(ctx context.Context, uid int64) (*basemodel.UserInfo, error) {
	addCache := true
	profile, err := r.CacheGetUserProfileByUid(ctx, uid)
	if err != nil && err != eredis.Nil {
		fmt.Println("CacheGetUserProfileByUid err:", err)
		addCache = false
		err = nil
	}
	res := profile.ToInfo()
	if res != nil {
		return res, err
	}
	// res, err = r.RawGetUserProfileByUid(ctx, uid)
	rv, err, _ := cacheSingleFlights[0].Do(keyProfile(uid), func() (rv interface{}, e error) {
		rv, e = r.RawGetUserProfileByUid(ctx, uid)
		return
	})
	profile = rv.(*model.Profile)
	if err != nil {
		return nil, err
	}
	res = profile.ToInfo()
	miss := res
	if !addCache {
		return res, err
	}
	r.data.fanout.Do(ctx, func(ctx context.Context) {
		fmt.Println("fanout add profile to redis")
		r.CacheAddUserInfo(ctx, uid, miss)
	})
	return res, err
}

func (r *profileRepo) ListUserInfoByUids(ctx context.Context, uids []int64) ([]*model.Profile, error) {
	addCache := true
	cacheList, err := r.CacheListUserProfileByUids(ctx, uids)
	if err != nil {
		addCache = false
		err = nil
	}
	cacheUids := make([]int64, 0, len(cacheList))
	for _, v := range cacheList {
		if v == nil {
			continue
		}
		cacheUids = append(cacheUids, v.Uid)
	}
	residualUids := xlist.Int64Difference(uids, cacheUids)
	if len(residualUids) == 0 {
		return cacheList, err
	}
	rawList, err := r.RawListUserProfileByUids(ctx, residualUids)
	if err != nil {
		return cacheList, err
	}
	res := append(cacheList, rawList...)
	if !addCache {
		return res, err
	}
	r.data.fanout.Do(ctx, func(ctx context.Context) {
		r.CacheAddListUserProfile(ctx, rawList)
	})
	return res, err
}

func (r *profileRepo) GetUserProfileByUid(ctx context.Context, uid int64) (*model.Profile, error) {
	addCache := true
	profile, err := r.CacheGetUserProfileByUid(ctx, uid)
	if err != nil && err != eredis.Nil {
		fmt.Println("CacheGetUserProfileByUid err:", err)
		addCache = false
		err = nil
	}
	if profile != nil {
		return profile, err
	}
	// res, err = r.RawGetUserProfileByUid(ctx, uid)
	rv, err, _ := cacheSingleFlights[0].Do(keyProfile(uid), func() (rv interface{}, e error) {
		rv, e = r.RawGetUserProfileByUid(ctx, uid)
		return
	})
	profile = rv.(*model.Profile)
	if err != nil {
		return nil, err
	}

	if !addCache {
		return profile, err
	}
	miss := profile
	r.data.fanout.Do(ctx, func(ctx context.Context) {
		fmt.Println("fanout add profile to redis")
		r.CacheAddUserProfile(ctx, uid, miss)
	})
	return profile, err
}

func (r *profileRepo) ListUserProfileByUids(ctx context.Context, uids []int64) ([]*model.Profile, error) {
	addCache := true
	cacheList, err := r.CacheListUserProfileByUids(ctx, uids)
	if err != nil {
		addCache = false
		err = nil
	}
	fmt.Println("cacheList:", cacheList, "err:", err)
	cacheUids := make([]int64, 0, len(cacheList))
	for _, v := range cacheList {
		if v == nil {
			continue
		}
		cacheUids = append(cacheUids, v.Uid)
	}
	residualUids := xlist.Int64Difference(uids, cacheUids)
	if len(residualUids) == 0 {
		return cacheList, err
	}
	rawList, err := r.RawListUserProfileByUids(ctx, residualUids)
	if err != nil {
		return cacheList, err
	}
	fmt.Println("cacheList:", cacheList, "+len(cacheList)", len(cacheList))
	fmt.Println("rawList:", rawList, "+len(rawList)", len(rawList))
	res := make([]*model.Profile, 0, len(uids))
	if len(cacheList) == 0 && len(rawList) == 0 {
		return nil, errors.ErrNoDocuments
	} else if len(cacheList) == 0 {
		res = append(res, rawList...)
	} else if len(rawList) == 0 {
		res = append(res, cacheList...)
	} else {
		res = append(cacheList, rawList...)
	}
	fmt.Println("res:", res)
	if !addCache {
		return res, err
	}
	r.data.fanout.Do(ctx, func(ctx context.Context) {
		r.CacheAddListUserProfile(ctx, rawList)
	})
	fmt.Println("end of ListUserProfileByUids,res:", res)
	return res, err
}
