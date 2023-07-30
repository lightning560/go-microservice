package data

import (
	"context"
	"fmt"
	"miopkg/log"
	"miopkg/util/xtime"
	"redbook/app/domain/user/model"
	"redbook/common/basemodel"
	"strconv"

	"miopkg/errors"

	jsoniter "github.com/json-iterator/go"
)

const (
	_prefixInfo    = "i_"
	_prefixCard    = "c_"
	_prefixProfile = "p_"
)

func keyInfo(uid int64) string {
	return _prefixInfo + strconv.FormatInt(uid, 10)
}

func keyCard(uid int64) string {
	return _prefixCard + strconv.FormatInt(uid, 10)
}

func keyProfile(uid int64) string {
	return _prefixProfile + strconv.FormatInt(uid, 10)
}

func (r *profileRepo) CacheGetUserProfileByUid(ctx context.Context, uid int64) (p *model.Profile, err error) {
	key := keyProfile(uid)
	rv, err := r.data.rc.Get(ctx, key)
	// if err is redis.Nil, p is nil
	if err != nil {
		err = errors.Wrap(err, "CacheGetUserProfileByUid")
		return
	}
	fmt.Println("get profile from redis;qry:", rv, ";err:", err)
	err = jsoniter.UnmarshalFromString(rv, &p)
	if err != nil {
		err = errors.Wrap(err, "data profile cache scan")
		log.Errorf("failed UnmarshalFromString: %v", err)
		return
	}
	return
}

func (r *profileRepo) CacheListUserProfileByUids(ctx context.Context, uids []int64) (res []*model.Profile, err error) {
	var keys []string
	for _, uid := range uids {
		keys = append(keys, keyProfile(uid))
	}
	rv, err := r.data.rc.MGet(ctx, keys)
	if err != nil {
		err = errors.Wrap(err, "data profile cache")
		return nil, err
	}
	for _, v := range rv {
		var p *model.Profile
		if v == nil {
			// res = append(res, nil)
			continue
		}
		err = jsoniter.UnmarshalFromString(v.(string), &p)
		if err != nil {
			err = errors.Wrap(err, "profile cache Unmarshal From String")
			return nil, err
		}
		res = append(res, p)
	}
	return res, nil
}

func (r *profileRepo) CacheAddUserInfo(ctx context.Context, uid int64, p *basemodel.UserInfo) error {
	key := keyInfo(uid)
	val, err := jsoniter.Marshal(p)
	if err != nil {
		err = errors.Wrap(err, "data info cache marshal")
		return err
	}
	err = r.data.rc.SetNX(context.Background(), key, val, xtime.Duration("120s"))
	if err != nil {
		err = errors.Wrap(err, "data:info cache setex")
		return err
	}
	return nil
}

func (r *profileRepo) CacheAddUserProfile(ctx context.Context, uid int64, p *model.Profile) error {
	key := keyProfile(uid)
	fmt.Println("CacheAddProfile:", key)
	val, err := jsoniter.Marshal(p)
	if err != nil {
		err = errors.Wrap(err, "data profile cache marshal")
		return err
	}
	err = r.data.rc.SetNX(context.Background(), key, val, xtime.Duration("120s"))
	if err != nil {
		err = errors.Wrap(err, "data profile cache setex")
		return err
	}
	return nil
}

func (r *profileRepo) CacheAddListUserProfile(ctx context.Context, lst []*model.Profile) error {
	for _, v := range lst {
		key := keyProfile(v.Uid)
		val, err := jsoniter.Marshal(v)
		if err != nil {
			err = errors.Wrap(err, "data profile cache marshal")
			return err
		}
		err = r.data.rc.SetNX(context.Background(), key, val, xtime.Duration("60s"))
		if err != nil {
			err = errors.Wrap(err, "data profile cache setex")
			return err
		}
	}
	return nil
}

func (r *profileRepo) CacheUpdateProfile(ctx context.Context, uid int64, p *model.Profile) error {
	key := keyProfile(uid)
	val, err := jsoniter.Marshal(p)
	if err != nil {
		err = errors.Wrap(err, "data profile cache marshal")
		return err
	}

	err = r.data.rc.SetEX(ctx, key, val, xtime.Duration("60s"))

	if err != nil {
		err = errors.Wrap(err, "data profile cache setex")
		return err
	}
	return nil
}
