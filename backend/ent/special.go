// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/special"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Special is the model entity for the Special schema.
type Special struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Data holds the value of the "data" field.
	Data         string `json:"data,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Special) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case special.FieldID, special.FieldData:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Special fields.
func (s *Special) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case special.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				s.ID = value.String
			}
		case special.FieldData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value.Valid {
				s.Data = value.String
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Special.
// This includes values selected through modifiers, order, etc.
func (s *Special) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Special.
// Note that you need to call Special.Unwrap() before calling this method if this Special
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Special) Update() *SpecialUpdateOne {
	return NewSpecialClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Special entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Special) Unwrap() *Special {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Special is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Special) String() string {
	var builder strings.Builder
	builder.WriteString("Special(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("data=")
	builder.WriteString(s.Data)
	builder.WriteByte(')')
	return builder.String()
}

// Specials is a parsable slice of Special.
type Specials []*Special
