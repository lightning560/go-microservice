package server

import (
	"context"
	v1 "redbook/api/domain/user/v1"
	basev1 "redbook/api/redbookpb/v1"
	"redbook/app/domain/user/model"
)

func (s *UserServer) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserResponse, error) {
	var (
		m = &model.User{
			Id:       req.User.Id,
			Uid:      req.User.Uid,
			Password: req.User.Password,
			Phone:    req.User.Phone,
			Email:    req.User.Email,
		}
	)
	err := s.userbiz.CreateUser(ctx, m)
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserResponse{}, nil
}

// func (s *UserServer) GetUser(ctx context.Context, req *v1.UidReq) (*v1.UserResponse, error) {
// 	rv, err := s.userbiz.GetUser(ctx, req.Uid)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &v1.UserResponse{
// 		User: &v1.User{
// 			Id:        rv.Id,
// 			Uid:       rv.Uid,
// 			Phone:     rv.Phone,
// 			Email:     rv.Email,
// 			CreatedAt: rv.CreatedAt,
// 		},
// 	}, nil
// }

func (s *UserServer) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*basev1.EmptyResponse, error) {
	var (
		m = &model.User{
			Id:    req.User.Id,
			Uid:   req.User.Uid,
			Phone: req.User.Phone,
			Email: req.User.Email,
		}
	)
	err := s.userbiz.UpdateUser(ctx, m)
	if err != nil {
		return nil, err
	}
	return &basev1.EmptyResponse{}, nil
}

func (s *UserServer) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordReq) (*basev1.EmptyResponse, error) {
	err := s.userbiz.UpdatePassword(ctx, req.Uid, req.Password)
	if err != nil {
		return nil, err
	}
	return &basev1.EmptyResponse{}, nil
}
