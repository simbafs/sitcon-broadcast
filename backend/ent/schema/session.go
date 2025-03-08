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
		field.String("title"),
		field.String("id").Immutable(),
		field.String("room").Immutable(),
		field.Strings("broadcastTo").Immutable(),
		field.String("broadcastFrom").Immutable(),
		field.Int64("start"),
		field.Int64("end"),
		field.String("speaker"),
		field.String("qa"),
		field.String("slidoID"),
		field.String("slido_admin_link"),
		field.String("co_write"),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return nil
}
