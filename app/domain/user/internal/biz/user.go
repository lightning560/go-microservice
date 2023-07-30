package biz

import (
	"context"
	"fmt"
	"log"
	"redbook/app/domain/user/model"
)

type IUserRepo interface {
	CreateUser(ctx context.Context, m *model.User) error
	GetUser(ctx context.Context, uid int64) (*model.User, error)
	UpdateUser(ctx context.Context, m *model.User) error
	UpdatePassword(ctx context.Context, m *model.User) error
}

type UserBiz struct {
	repo IUserRepo
}

func NewUserBiz(repo IUserRepo) *UserBiz {
	return &UserBiz{repo: repo}
}

func (ubz *UserBiz) CreateUser(ctx context.Context, m *model.User) (err error) {
	fmt.Println("email:", m.Email, "+password_hash:", m.Password)
	log.Println(m)
	err = ubz.repo.CreateUser(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (ubz *UserBiz) GetUser(ctx context.Context, uid int64) (rv *model.User, err error) {
	rv, err = ubz.repo.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (ubz *UserBiz) UpdateUser(ctx context.Context, m *model.User) (err error) {
	err = ubz.repo.UpdateUser(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (ubz *UserBiz) UpdatePassword(ctx context.Context, uid int64, password string) (err error) {
	var (
		m = &model.User{Uid: uid, Password: password}
	)
	err = ubz.repo.UpdatePassword(ctx, m)
	if err != nil {
		return err
	}
	return nil
}
