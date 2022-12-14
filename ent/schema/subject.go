package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Subject holds the schema definition for the Subject entity.
type Subject struct {
	ent.Schema
}

// Fields of the Subject.
func (Subject) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
	}
}

// Edges of the Subject.
func (Subject) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("classes", Class.Type),
	}
}
