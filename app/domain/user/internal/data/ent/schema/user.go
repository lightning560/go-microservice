package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Unique().
			Positive(),
		field.Int64("uid").
			Unique().
			Positive(),
		field.String("phone").
			Optional(),
		field.String("email").
			Unique(),
		field.String("password").
			Sensitive(),
		field.Int("identification").
			Optional(),
		field.Bool("deleted").
			Default(false),
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

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profiles", Profile.Type),
	}
}
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "uid").
			Unique(),
	}
}
