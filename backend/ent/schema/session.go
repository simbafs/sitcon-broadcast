package schema

import (
	"entgo.io/ent"
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
		field.Int64("start"),
		field.Int64("end"),
		field.String("session_id").Immutable(),
		field.String("room").Immutable(),
		field.String("next").Immutable(),
		field.String("title"),
		field.String("speaker"),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return nil
}
