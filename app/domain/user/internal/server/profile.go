package server

import (
	"context"
	v1 "redbook/api/domain/user/v1"
	basev1 "redbook/api/redbookpb/v1"
	"redbook/app/domain/user/model"
	"redbook/common/basemodel"
)

func (s *UserServer) UpdateProfile(ctx context.Context, req *v1.UpdateProfileReq) (*basev1.EmptyResponse, error) {
	var (
		m = &model.Profile{
			Uid:  req.Profile.Uid,
			Name: req.Profile.Name,
			Avatar: &basemodel.Image{
				Url:    req.Profile.Avatar.Url,
				Width:  req.Profile.Avatar.Width,
				Height: req.Profile.Avatar.Height,
				SizeKb: req.Profile.Avatar.SizeKb,
				Name:   req.Profile.Avatar.Name,
				Hash:   req.Profile.Avatar.Hash,
			},
			Sign: req.Profile.Sign,
		}
	)
	err := s.profilebiz.UpdateProfile(ctx, m)
	if err != nil {
		return nil, err
	}
	return &basev1.EmptyResponse{}, nil
}

func (s *UserServer) GetUserInfoByUid(ctx context.Context, req *v1.GetUserInfoByUidReq) (*v1.GetUserInfoByUidResponse, error) {
	rv, err := s.profilebiz.GetUserInfoByUid(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserInfoByUidResponse{
		Info: &basev1.UserInfo{
			Id:     rv.Id,
			Uid:    rv.Uid,
			Name:   rv.Name,
			Sex:    rv.Sex,
			Avatar: rv.Avatar.ToProto(),
			Level:  rv.Level,
			Sign:   rv.Sign,
			State:  rv.State,
		},
	}, nil
}

func (s *UserServer) MapUserInfoByUids(ctx context.Context, req *v1.MapUserInfoByUidsReq) (*v1.MapUserInfoByUidsResponse, error) {
	rv, err := s.profilebiz.ListUserInfoByUids(ctx, req.Uids)
	if err != nil {
		return nil, err
	}
	res := make(map[int64]*basev1.UserInfo, len(rv))
	for _, v := range rv {
		res[v.Uid] = v.ToProto()
	}
	return &v1.MapUserInfoByUidsResponse{
		Infos: res,
		Total: int32(len(res)),
	}, nil

}

func (s *UserServer) GetUserCardByUid(c context.Context, req *v1.GetUserCardByUidReq) (*v1.GetUserCardByUidResponse, error) {
	rv, err := s.profilebiz.GetUserCardByUid(c, req.Uid)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserCardByUidResponse{
		Card: &v1.Card{
			Uid:    rv.Uid,
			Name:   rv.Name,
			Avatar: rv.Avatar.ToProto(),
			Level:  int32(rv.Level),
		},
	}, nil
}

func (s *UserServer) GetUserProfileByUid(c context.Context, req *v1.GetUserProfileByUidReq) (*v1.GetUserProfileByUidResponse, error) {
	rv, err := s.profilebiz.GetUserProfileByUid(c, req.Uid)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserProfileByUidResponse{
		Profile: &v1.Profile{
			Uid:  rv.Uid,
			Name: rv.Name,
			Avatar: &basev1.Image{
				Url:   rv.Avatar.Url,
				Width: rv.Avatar.Width,
			},
			Level: rv.Level,
			Sign:  rv.Sign,
		},
	}, nil
}
