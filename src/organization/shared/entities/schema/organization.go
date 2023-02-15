package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().Annotations(entgql.OrderField("name")),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Organization) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}
