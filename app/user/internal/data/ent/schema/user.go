package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field{
	return []ent.Field{
		field.Int64("id"),
		field.String("name"),
		field.String("secret"),
		field.Uint64("age"),
		field.String("phone"),
	}
}

func (User) Edges() []ent.Edge{
	return []ent.Edge{
	}
}