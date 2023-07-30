package model

import (
	mallv1 "redbook/api/domain/mall/v1"
	"redbook/app/domain/mall/model/entity"
)

type CartItem struct {
	Id_          string `json:"_id,omitempty" bson:"_id,omitempty"`
	Id           int64  `json:"id,omitempty" bson:"id,omitempty"`
	Uid          int64  `json:"uid,omitempty" bson:"uid,omitempty"`
	CollectionId int64  `json:"collection_id,omitempty" bson:"collection_id,omitempty"`
	ProductId    int64  `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Quantity     int32  `json:"quantity,omitempty" bson:"quantity,omitempty"`
	CreatedAt    int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    int64  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type CartItems []*CartItem

func (m *CartItem) FromEntity(entity *entity.CartItem) {
	m.Id_ = entity.Id_.Hex()
	m.Id = entity.Id
	m.Uid = entity.Uid
	m.CollectionId = entity.CollectionId
	m.ProductId = entity.ProductId
	m.Quantity = entity.Quantity
	m.CreatedAt = entity.CreatedAt
	m.UpdatedAt = entity.UpdatedAt
}

func (m *CartItem) ToProto() *mallv1.CartItem {
	return &mallv1.CartItem{
		Id:           m.Id,
		Uid:          m.Uid,
		CollectionId: m.CollectionId,
		ProductId:    m.ProductId,
		Quantity:     m.Quantity,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func (ms *CartItems) ListFromEntity(entity []*entity.CartItem) {
	for _, v := range entity {
		m := &CartItem{}
		m.FromEntity(v)
		*ms = append(*ms, m)
	}
}

func (ms *CartItems) ListToProto() []*mallv1.CartItem {
	res := []*mallv1.CartItem{}
	for _, v := range *ms {
		res = append(res, v.ToProto())
	}
	return res
}
