package data

import (
	"context"
	"fmt"
	"miopkg/log"
	"miopkg/util/randgen"
	"miopkg/util/snowflake"
	"redbook/app/domain/user/internal/biz"
	"redbook/app/domain/user/internal/data/ent"
	"redbook/app/domain/user/internal/data/ent/user"
	"redbook/app/domain/user/model"
)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.IUserRepo {
	//从wire改为编码方式
	// data, _, err := NewData()
	// if err != nil {
	// 	log.Println("NewUserRepoApi error:%w", err)
	// 	return nil
	// }
	return &userRepo{data: data}
}

func (r *userRepo) CreateUser(ctx context.Context, u *model.User) error {
	id := snowflake.GenInt64()
	tx, err := r.data.userWDB.Debug().Tx(ctx)
	if err != nil {
		return fmt.Errorf("CreateUser starting a transaction: %w", err)
	}
	// uidstring := randword.RandomWord(8, "num")
	// uid, err := strconv.ParseInt(uidstring, 10, 64)
	// if err != nil {
	// 	log.Errorf("CreateUser uid strconv.ParseInt error:%w", err)
	// 	return err
	// }
	uid := randgen.RandomInt64(8)
	ures, err := tx.User.
		Create().SetID(id).SetUID(uid).SetPhone(u.Phone).SetEmail(u.Email).SetPassword(u.Password).Save(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("CreateUser failed creating user: %w", err))
	}
	log.Info(fmt.Sprintf("user was created:%p", ures), log.FieldMod("user"))
	pres, err := tx.Profile.Create().SetID(u.Id).SetUID(u.Uid).Save(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("CreateUser failed creating profile: %w", err))
	}
	log.Info(fmt.Sprintf("profile was created:%p", pres), log.FieldMod("user"))
	return tx.Commit()
	// return &model.User{
	// 	UserName: res.Username,
	// 	Pwd:      []byte(res.Password),
	// }, nil
}

func (r *userRepo) GetUser(ctx context.Context, uid int64) (*model.User, error) {
	urv, err := r.data.userRDB.Debug().User.Query().Where(user.UID(uid)).Only(ctx)
	if err != nil {
		log.Error("GetUser can't query user by uid: %w", log.FieldErr(err), log.FieldMod("user"))
		return nil, err
	}
	return &model.User{
		Id:        urv.ID,
		Uid:       urv.UID,
		Phone:     urv.Phone,
		Email:     urv.Email,
		CreatedAt: urv.CreatedAt.Unix(),
	}, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, u *model.User) error {
	err := r.data.userWDB.Debug().User.UpdateOneID(u.Id).SetPhone(u.Phone).SetEmail(u.Email).Exec(ctx)
	switch {
	// If the entity does not meet a specific condition,
	// the operation will return an "ent.NotFoundError".
	case ent.IsNotFound(err):
		fmt.Println("user item was not found")
		log.Errorf("UpdateUser can't find user by id: %w", log.FieldErr(err), log.FieldMod("user"))
	// Any other error.
	case err != nil:
		fmt.Println("update error:", err)
		log.Errorf("UpdateUser can't update user: %w", log.FieldErr(err), log.FieldMod("user"))
	}
	return err
	// urv, err := r.data.userWDB.Debug().User.Query().Where(user.UID(u.Uid)).Only(ctx)
	// if err != nil {
	// 	log.Error("UpdateUser can't query user by uid: %w", log.FieldErr(err), log.FieldMod("user"))
	// 	return err
	// }
	// _, err = urv.Update().SetPhone(u.Phone).SetEmail(u.Email).Save(ctx)
	// if err != nil {
	// 	log.Error("can't update user: %w", log.FieldErr(err), log.FieldMod("user"))
	// 	return err
	// }
	// return err
}

func (r *userRepo) UpdatePassword(ctx context.Context, u *model.User) error {
	urv, err := r.data.userWDB.Debug().User.Query().Where(user.UID(u.Uid)).Only(ctx)
	if err != nil {
		log.Error("UpdatePassword can't query user by uid: %w", log.FieldErr(err), log.FieldMod("user"))
		return err
	}
	_, err = urv.Update().SetPassword(u.Password).Save(ctx)
	if err != nil {
		log.Error("can't update password: %w", log.FieldErr(err), log.FieldMod("user"))
		return err
	}
	return err
}
