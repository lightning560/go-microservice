package model

import (
	domainmallv1 "redbook/api/domain/mall/v1"
	interfacemallv1 "redbook/api/interface/mall/v1"
)

type CartItem struct {
	Id           int64       `json:"id,omitempty" bson:"id,omitempty"`
	Uid          int64       `json:"uid,omitempty" bson:"uid,omitempty"`
	CollectionId int64       `json:"collection_id,omitempty" bson:"collection_id,omitempty"`
	ProductId    int64       `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Quantity     int32       `json:"quantity,omitempty" bson:"quantity,omitempty"`
	CreatedAt    int64       `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    int64       `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	ProductCard  ProductCard `json:"product_card,omitempty" bson:"product_card,omitempty"`
}

type CartItems []*CartItem

func (m *CartItem) ToDomainProto() *domainmallv1.CartItem {
	return &domainmallv1.CartItem{
		Id:           m.Id,
		Uid:          m.Uid,
		CollectionId: m.CollectionId,
		ProductId:    m.ProductId,
		Quantity:     m.Quantity,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func (m *CartItem) FromDomainProto(item *domainmallv1.CartItem) {
	m.Id = item.Id
	m.Uid = item.Uid
	m.CollectionId = item.CollectionId
	m.ProductId = item.ProductId
	m.Quantity = item.Quantity
	m.CreatedAt = item.CreatedAt
	m.UpdatedAt = item.UpdatedAt
	m.ProductCard = ProductCard{}
	m.ProductCard.FromDomainProto(item.ProductCard)
}

func (ms *CartItems) ListToDomainProto() []*domainmallv1.CartItem {
	var res []*domainmallv1.CartItem
	for _, m := range *ms {
		res = append(res, m.ToDomainProto())
	}
	return res
}

func (ms *CartItems) ListFromDomainProto(items []*domainmallv1.CartItem) {
	for _, v := range items {
		m := &CartItem{}
		m.FromDomainProto(v)
		*ms = append(*ms, m)
	}
}

func (m *CartItem) ToInterfaceProto() *interfacemallv1.CartItem {
	return &interfacemallv1.CartItem{
		Id:           m.Id,
		Uid:          m.Uid,
		CollectionId: m.CollectionId,
		ProductId:    m.ProductId,
		Quantity:     m.Quantity,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
		ProductCard:  m.ProductCard.ToInterfaceProto(),
	}
}
func (m *CartItem) FromInterfaceProto(c *interfacemallv1.CartItem) {
	m.Id = c.Id
	m.Uid = c.Uid
	m.CollectionId = c.CollectionId
	m.ProductId = c.ProductId
	m.Quantity = c.Quantity
}

func (ms *CartItems) ListToInterfaceProto() []*interfacemallv1.CartItem {
	var res []*interfacemallv1.CartItem
	for _, m := range *ms {
		res = append(res, m.ToInterfaceProto())
	}
	return res
}
