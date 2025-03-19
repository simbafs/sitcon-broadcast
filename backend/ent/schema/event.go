package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("url").Comment("url to session.json"),
		field.String("script").Comment("js script to process session.json to ent.Session"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return nil
}
