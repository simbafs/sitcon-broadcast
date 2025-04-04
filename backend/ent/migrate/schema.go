// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "url", Type: field.TypeString},
		{Name: "script", Type: field.TypeString},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "idx", Type: field.TypeInt8},
		{Name: "finish", Type: field.TypeBool, Default: false},
		{Name: "start", Type: field.TypeInt64},
		{Name: "end", Type: field.TypeInt64},
		{Name: "room", Type: field.TypeString},
		{Name: "session_id", Type: field.TypeString},
		{Name: "next", Type: field.TypeString},
		{Name: "title", Type: field.TypeString},
		{Name: "data", Type: field.TypeJSON},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EventsTable,
		SessionsTable,
	}
)

func init() {
}
