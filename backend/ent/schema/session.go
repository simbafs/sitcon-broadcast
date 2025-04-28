package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	// TODO: broadcastTo and broadcastFrom
	return []ent.Field{
		field.Int8("idx").Comment("just for sorting"),
		field.Bool("finish").Default(false),
		field.Int64("start"),
		field.Int64("end"),
		field.String("room").Immutable(),
		field.String("session_id").Immutable(),

		field.String("title"),
		// extra data, such as title, speakers, etc.
		field.JSON("data", map[string]any{}),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("next", Session.Type),
	}
}
