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
	return []ent.Field{
		field.String("id"),
		field.String("title"),
		field.String("type"),
		field.Strings("speakers"),
		field.String("room"),
		field.Strings("broadcast"),
		field.Time("start"),
		field.Time("end"),
		field.String("slido"), // json: qa
		field.String("slide"),
		field.String("hackmd"), // json: co_write
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return nil
}
