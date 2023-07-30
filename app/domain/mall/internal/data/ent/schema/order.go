package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type OrderItem struct {
	CollectionId int64 `json:"collection_id,omitempty" bson:"collection_id,omitempty"`
	ProductId    int64 `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Quantity     int32 `json:"quantity,omitempty" bson:"quantity,omitempty"`
}

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Positive(),
		field.Int64("uid").
			Unique().Positive(),
		field.JSON("order_items", []OrderItem{}).
			Optional(),
		field.Int32("state"),
		field.Int64("total_amount").
			Optional(),
		field.Int64("total_quantity").
			Optional(),
		field.Time("created_at").
			Immutable().
			Default(time.Now).SchemaType(map[string]string{
			"mysql": "datetime",
		}),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			"mysql": "datetime",
		}).
			UpdateDefault(time.Now),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}
