// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/schema"
	"backend/ent/session"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescFinish is the schema descriptor for finish field.
	sessionDescFinish := sessionFields[1].Descriptor()
	// session.DefaultFinish holds the default value on creation for the finish field.
	session.DefaultFinish = sessionDescFinish.Default.(bool)
}
