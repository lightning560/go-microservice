package data

import (
	"context"
	"miopkg/log"
	mallv1 "redbook/api/domain/mall/v1"
	"redbook/app/interface/mall/internal/biz"
)

type cartRepo struct {
	data *Data
}

func NewCartRepo(data *Data) biz.ICartRepo {
	return &cartRepo{data: data}
}

func (r *cartRepo) CreateCartItem(ctx context.Context, uid, collectionId, productId int64, quantity int32) (int64, error) {
	rv, err := r.data.mallRpc.CreateCartItem(ctx, &mallv1.CreateCartItemReq{
		Uid:          uid,
		CollectionId: collectionId,
		ProductId:    productId,
		Quantity:     quantity,
	})
	if err != nil {
		log.Errorf("CreateCartItem error: %v", err)
		return -1, err
	}
	return rv.Id, nil
}

// func (r *cartRepo) GetCartItem(ctx context.Context, id int64) (mallv1.CartItem, error) {
// 	resp, err := r.data.mallRpc.GetCartItem(ctx, &mallv1.GetCartItemReq{
// 		Id: id,
// 	})
// 	if err != nil {
// 		log.Errorf("GetCartItem error: %v", err)
// 		return mallv1.CartItem{}, err
// 	}
// 	return *resp.CartItem, nil
// }

func (r *cartRepo) ListCartItemByUid(ctx context.Context, uid int64) ([]*mallv1.CartItem, error) {
	rv, err := r.data.mallRpc.MapCartItemByUid(ctx, &mallv1.MapCartItemByUidReq{
		Uid: uid,
	})
	if err != nil {
		log.Errorf("ListCartItemByUid error: %v", err)
		return nil, err
	}
	res := make([]*mallv1.CartItem, 0, len(rv.CartItems))
	for _, v := range rv.CartItems {
		res = append(res, v)
	}
	return res, nil
}

func (r *cartRepo) UpdateCartItemQuantity(ctx context.Context, id int64, quantity int32) error {
	_, err := r.data.mallRpc.UpdateCartItemQuantity(ctx, &mallv1.UpdateCartItemQuantityReq{
		Id:       id,
		Quantity: quantity,
	})
	if err != nil {
		log.Errorf("UpdateCartItemQuantity error: %v", err)
		return err
	}
	return nil
}
