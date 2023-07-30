package model

import (
	"redbook/app/domain/mall/internal/data/ent"
	"redbook/app/domain/mall/internal/data/ent/schema"
	"time"
)

type OrderItem struct {
	CollectionId int64
	ProductId    int64
	Quantity     int32
	Product      Product
}

type OrderItems []*OrderItem

type Order struct {
	Id            int64
	Uid           int64
	OrderItems    OrderItems
	State         int32
	TotalAmount   int64
	TotalQuantity int64
	CreatedAt     int64
	UpdatedAt     int64
}

func (m *OrderItem) ToEntity() *schema.OrderItem {
	return &schema.OrderItem{
		CollectionId: m.CollectionId,
		ProductId:    m.ProductId,
		Quantity:     m.Quantity,
	}
}
func (m *OrderItem) FromEntity(entity *schema.OrderItem) *OrderItem {
	return &OrderItem{
		CollectionId: entity.CollectionId,
		ProductId:    entity.ProductId,
		Quantity:     entity.Quantity,
	}
}

func (ms *OrderItems) ListToEntity() []schema.OrderItem {
	var rv []schema.OrderItem
	for _, m := range *ms {
		rv = append(rv, *m.ToEntity())
	}
	return rv
}

func (ms *OrderItems) ListFromEntity(entities []schema.OrderItem) {
	for _, entity := range entities {
		m := &OrderItem{}
		m.FromEntity(&entity)
		*ms = append(*ms, m)
	}
}

func (m *Order) ToEntity() *ent.Order {
	ca := time.Unix(m.CreatedAt, 0)
	ua := time.Unix(m.UpdatedAt, 0)
	return &ent.Order{
		ID:            m.Id,
		UID:           m.Uid,
		OrderItems:    m.OrderItems.ListToEntity(),
		State:         m.State,
		TotalAmount:   m.TotalAmount,
		TotalQuantity: m.TotalQuantity,
		CreatedAt:     ca,
		UpdatedAt:     ua,
	}
}

func (m *Order) FromEntity(entity *ent.Order) *Order {
	orderItems := &OrderItems{}
	orderItems.ListFromEntity(entity.OrderItems)
	return &Order{
		Id:            entity.ID,
		Uid:           entity.UID,
		OrderItems:    *orderItems,
		State:         entity.State,
		TotalAmount:   entity.TotalAmount,
		TotalQuantity: entity.TotalQuantity,
		CreatedAt:     entity.CreatedAt.Unix(),
		UpdatedAt:     entity.UpdatedAt.Unix(),
	}
}
