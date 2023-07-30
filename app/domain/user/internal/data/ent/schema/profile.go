package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Address holds the schema definition for the Address entity.
type Profile struct {
	ent.Schema
}

// Fields of the Address.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			// Default(snowflake.GenInt64()).
			Unique(),
		field.Int64("uid").Unique(),
		field.String("name"),
		field.Int32("sex").
			Default(0),
		field.String("avatar").
			Default("https://dummyimage.com/100"),
		field.String("sign").Optional(),
		field.Time("birthday").Optional(),
		field.Int32("level").
			Default(0),
		field.Int8("verify_type").
			Default(0),
		field.Int32("attr").
			Default(0),
		field.Int32("state").
			Default(0),
		field.Bool("deleted"),
		field.Time("created_at").
			Immutable().
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).
			UpdateDefault(time.Now),
	}
}

// Edges of the Address.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("profiles").
			Required().
			Unique(),
	}
}
func (Profile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "uid").
			Unique(),
	}
}
