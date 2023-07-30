package data

import (
	"context"
	"fmt"
	"miopkg/errors"
	userv1 "redbook/api/domain/user/v1"
	basev1 "redbook/api/redbookpb/v1"
	"redbook/app/interface/comment/internal/biz"
	"redbook/common/basemodel"
)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.IUserRepo {
	return &userRepo{data: data}
}

func (r *userRepo) GetUserInfoByUid(ctx context.Context, uid int64) (*basev1.UserInfo, error) {
	rv, err := r.data.userRpcClient.GetUserInfoByUid(ctx, &userv1.GetUserInfoByUidReq{Uid: uid})
	if err != nil {
		return nil, errors.Wrap(err, "GetUserInfoByUid")
	}
	return rv.Info, nil
}

func (r *userRepo) MapUserInfoByUids(ctx context.Context, uids []int64) (map[int64]*basemodel.UserInfo, error) {
	rv, err := r.data.userRpcClient.MapUserInfoByUids(ctx, &userv1.MapUserInfoByUidsReq{Uids: uids})
	if err != nil {
		return nil, err
	}
	var res = make(map[int64]*basemodel.UserInfo, rv.Total)
	fmt.Println("len(rv.Infos)", len(rv.Infos))
	for k, v := range rv.Infos {
		m := &basemodel.UserInfo{}
		m.FromProto(v)
		res[k] = m
	}
	fmt.Println("len(res)", len(res))
	return res, nil
}
