package data

import (
	"context"
	"miopkg/util/snowflake"
	"redbook/app/domain/mall/internal/biz"
	"redbook/app/domain/mall/internal/data/ent"
	"redbook/app/domain/mall/internal/data/ent/order"
	"redbook/app/domain/mall/model"
	"time"
)

type orderRepo struct {
	data *Data
}

func NewOrderRepo(data *Data) biz.IOrderRepo {
	return &orderRepo{data: data}
}
func (r *orderRepo) CreateOrder(ctx context.Context, order *model.Order) (int64, error) {
	id := snowflake.GenInt64()
	rv, err := r.data.orderWDB.Order.Create().
		SetID(id).
		SetUID(order.Uid).
		SetState(1).
		SetOrderItems(order.OrderItems.ListToEntity()).
		SetTotalAmount(order.TotalAmount).
		SetTotalQuantity(order.TotalQuantity).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return -1, err
	}
	return rv.ID, nil
}
func (r *orderRepo) GetOrderById(ctx context.Context, id int64) (*ent.Order, error) {
	rv, err := r.data.orderWDB.Order.Query().Where(order.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return rv, nil
}
func (r *orderRepo) ListOrderByIds(ctx context.Context, ids []int64) ([]*ent.Order, error) {
	rv, err := r.data.orderWDB.Order.
		Query().
		Where(order.IDIn(ids...)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return rv, nil
}
func (r *orderRepo) ListOrderByUid(ctx context.Context, uid int64, offset, limit int32, sortBy, sortOrder string) ([]*ent.Order, error) {

	query := r.data.orderWDB.Order.
		Query().
		Where(order.UID(uid)).
		Offset(int(offset)).
		Limit(int(limit))
	if sortBy != "" {
		if sortOrder == "asc" {
			query = query.Order(ent.Asc(sortBy))
		} else {
			query = query.Order(ent.Desc(sortBy))
		}
	}
	rv, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return rv, nil
}
func (r *orderRepo) UpdateOrderState(ctx context.Context, id int64, state int32) error {
	_, err := r.data.orderWDB.Order.
		UpdateOneID(id).
		SetState(state).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
