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
		field.String("id").Default(""),
		field.String("title").Default(""),
		field.String("type").Default(""),
		field.Strings("speakers").Default([]string{}),
		field.String("room").Default(""),
		field.Strings("broadcast").Default([]string{}),
		field.Int64("start").Default(0),
		field.Int64("end").Default(0),
		field.String("slido").Default(""), // json: qa
		field.String("slide").Default(""),
		field.String("hackmd").Default(""), // json: co_write
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return nil
}
