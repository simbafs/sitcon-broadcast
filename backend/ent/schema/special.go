package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Special holds the schema definition for the Special entity.
type Special struct {
	ent.Schema
}

// Fields of the Special.
func (Special) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Default(""),
		field.String("data").Default(""),
	}
}

// Edges of the Special.
func (Special) Edges() []ent.Edge {
	return nil
}
