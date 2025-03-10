// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/predicate"
	"backend/ent/session"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks    []Hook
	mutation *SessionMutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetTitle sets the "title" field.
func (su *SessionUpdate) SetTitle(s string) *SessionUpdate {
	su.mutation.SetTitle(s)
	return su
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (su *SessionUpdate) SetNillableTitle(s *string) *SessionUpdate {
	if s != nil {
		su.SetTitle(*s)
	}
	return su
}

// SetType sets the "type" field.
func (su *SessionUpdate) SetType(s string) *SessionUpdate {
	su.mutation.SetType(s)
	return su
}

// SetNillableType sets the "type" field if the given value is not nil.
func (su *SessionUpdate) SetNillableType(s *string) *SessionUpdate {
	if s != nil {
		su.SetType(*s)
	}
	return su
}

// SetSpeakers sets the "speakers" field.
func (su *SessionUpdate) SetSpeakers(s []string) *SessionUpdate {
	su.mutation.SetSpeakers(s)
	return su
}

// AppendSpeakers appends s to the "speakers" field.
func (su *SessionUpdate) AppendSpeakers(s []string) *SessionUpdate {
	su.mutation.AppendSpeakers(s)
	return su
}

// SetRoom sets the "room" field.
func (su *SessionUpdate) SetRoom(s string) *SessionUpdate {
	su.mutation.SetRoom(s)
	return su
}

// SetNillableRoom sets the "room" field if the given value is not nil.
func (su *SessionUpdate) SetNillableRoom(s *string) *SessionUpdate {
	if s != nil {
		su.SetRoom(*s)
	}
	return su
}

// SetBroadcast sets the "broadcast" field.
func (su *SessionUpdate) SetBroadcast(s []string) *SessionUpdate {
	su.mutation.SetBroadcast(s)
	return su
}

// AppendBroadcast appends s to the "broadcast" field.
func (su *SessionUpdate) AppendBroadcast(s []string) *SessionUpdate {
	su.mutation.AppendBroadcast(s)
	return su
}

// SetStart sets the "start" field.
func (su *SessionUpdate) SetStart(i int64) *SessionUpdate {
	su.mutation.ResetStart()
	su.mutation.SetStart(i)
	return su
}

// SetNillableStart sets the "start" field if the given value is not nil.
func (su *SessionUpdate) SetNillableStart(i *int64) *SessionUpdate {
	if i != nil {
		su.SetStart(*i)
	}
	return su
}

// AddStart adds i to the "start" field.
func (su *SessionUpdate) AddStart(i int64) *SessionUpdate {
	su.mutation.AddStart(i)
	return su
}

// SetEnd sets the "end" field.
func (su *SessionUpdate) SetEnd(i int64) *SessionUpdate {
	su.mutation.ResetEnd()
	su.mutation.SetEnd(i)
	return su
}

// SetNillableEnd sets the "end" field if the given value is not nil.
func (su *SessionUpdate) SetNillableEnd(i *int64) *SessionUpdate {
	if i != nil {
		su.SetEnd(*i)
	}
	return su
}

// AddEnd adds i to the "end" field.
func (su *SessionUpdate) AddEnd(i int64) *SessionUpdate {
	su.mutation.AddEnd(i)
	return su
}

// SetSlido sets the "slido" field.
func (su *SessionUpdate) SetSlido(s string) *SessionUpdate {
	su.mutation.SetSlido(s)
	return su
}

// SetNillableSlido sets the "slido" field if the given value is not nil.
func (su *SessionUpdate) SetNillableSlido(s *string) *SessionUpdate {
	if s != nil {
		su.SetSlido(*s)
	}
	return su
}

// SetSlide sets the "slide" field.
func (su *SessionUpdate) SetSlide(s string) *SessionUpdate {
	su.mutation.SetSlide(s)
	return su
}

// SetNillableSlide sets the "slide" field if the given value is not nil.
func (su *SessionUpdate) SetNillableSlide(s *string) *SessionUpdate {
	if s != nil {
		su.SetSlide(*s)
	}
	return su
}

// SetHackmd sets the "hackmd" field.
func (su *SessionUpdate) SetHackmd(s string) *SessionUpdate {
	su.mutation.SetHackmd(s)
	return su
}

// SetNillableHackmd sets the "hackmd" field if the given value is not nil.
func (su *SessionUpdate) SetNillableHackmd(s *string) *SessionUpdate {
	if s != nil {
		su.SetHackmd(*s)
	}
	return su
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Title(); ok {
		_spec.SetField(session.FieldTitle, field.TypeString, value)
	}
	if value, ok := su.mutation.GetType(); ok {
		_spec.SetField(session.FieldType, field.TypeString, value)
	}
	if value, ok := su.mutation.Speakers(); ok {
		_spec.SetField(session.FieldSpeakers, field.TypeJSON, value)
	}
	if value, ok := su.mutation.AppendedSpeakers(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, session.FieldSpeakers, value)
		})
	}
	if value, ok := su.mutation.Room(); ok {
		_spec.SetField(session.FieldRoom, field.TypeString, value)
	}
	if value, ok := su.mutation.Broadcast(); ok {
		_spec.SetField(session.FieldBroadcast, field.TypeJSON, value)
	}
	if value, ok := su.mutation.AppendedBroadcast(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, session.FieldBroadcast, value)
		})
	}
	if value, ok := su.mutation.Start(); ok {
		_spec.SetField(session.FieldStart, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedStart(); ok {
		_spec.AddField(session.FieldStart, field.TypeInt64, value)
	}
	if value, ok := su.mutation.End(); ok {
		_spec.SetField(session.FieldEnd, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedEnd(); ok {
		_spec.AddField(session.FieldEnd, field.TypeInt64, value)
	}
	if value, ok := su.mutation.Slido(); ok {
		_spec.SetField(session.FieldSlido, field.TypeString, value)
	}
	if value, ok := su.mutation.Slide(); ok {
		_spec.SetField(session.FieldSlide, field.TypeString, value)
	}
	if value, ok := su.mutation.Hackmd(); ok {
		_spec.SetField(session.FieldHackmd, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SessionMutation
}

// SetTitle sets the "title" field.
func (suo *SessionUpdateOne) SetTitle(s string) *SessionUpdateOne {
	suo.mutation.SetTitle(s)
	return suo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableTitle(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetTitle(*s)
	}
	return suo
}

// SetType sets the "type" field.
func (suo *SessionUpdateOne) SetType(s string) *SessionUpdateOne {
	suo.mutation.SetType(s)
	return suo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableType(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetType(*s)
	}
	return suo
}

// SetSpeakers sets the "speakers" field.
func (suo *SessionUpdateOne) SetSpeakers(s []string) *SessionUpdateOne {
	suo.mutation.SetSpeakers(s)
	return suo
}

// AppendSpeakers appends s to the "speakers" field.
func (suo *SessionUpdateOne) AppendSpeakers(s []string) *SessionUpdateOne {
	suo.mutation.AppendSpeakers(s)
	return suo
}

// SetRoom sets the "room" field.
func (suo *SessionUpdateOne) SetRoom(s string) *SessionUpdateOne {
	suo.mutation.SetRoom(s)
	return suo
}

// SetNillableRoom sets the "room" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableRoom(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetRoom(*s)
	}
	return suo
}

// SetBroadcast sets the "broadcast" field.
func (suo *SessionUpdateOne) SetBroadcast(s []string) *SessionUpdateOne {
	suo.mutation.SetBroadcast(s)
	return suo
}

// AppendBroadcast appends s to the "broadcast" field.
func (suo *SessionUpdateOne) AppendBroadcast(s []string) *SessionUpdateOne {
	suo.mutation.AppendBroadcast(s)
	return suo
}

// SetStart sets the "start" field.
func (suo *SessionUpdateOne) SetStart(i int64) *SessionUpdateOne {
	suo.mutation.ResetStart()
	suo.mutation.SetStart(i)
	return suo
}

// SetNillableStart sets the "start" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableStart(i *int64) *SessionUpdateOne {
	if i != nil {
		suo.SetStart(*i)
	}
	return suo
}

// AddStart adds i to the "start" field.
func (suo *SessionUpdateOne) AddStart(i int64) *SessionUpdateOne {
	suo.mutation.AddStart(i)
	return suo
}

// SetEnd sets the "end" field.
func (suo *SessionUpdateOne) SetEnd(i int64) *SessionUpdateOne {
	suo.mutation.ResetEnd()
	suo.mutation.SetEnd(i)
	return suo
}

// SetNillableEnd sets the "end" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableEnd(i *int64) *SessionUpdateOne {
	if i != nil {
		suo.SetEnd(*i)
	}
	return suo
}

// AddEnd adds i to the "end" field.
func (suo *SessionUpdateOne) AddEnd(i int64) *SessionUpdateOne {
	suo.mutation.AddEnd(i)
	return suo
}

// SetSlido sets the "slido" field.
func (suo *SessionUpdateOne) SetSlido(s string) *SessionUpdateOne {
	suo.mutation.SetSlido(s)
	return suo
}

// SetNillableSlido sets the "slido" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableSlido(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetSlido(*s)
	}
	return suo
}

// SetSlide sets the "slide" field.
func (suo *SessionUpdateOne) SetSlide(s string) *SessionUpdateOne {
	suo.mutation.SetSlide(s)
	return suo
}

// SetNillableSlide sets the "slide" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableSlide(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetSlide(*s)
	}
	return suo
}

// SetHackmd sets the "hackmd" field.
func (suo *SessionUpdateOne) SetHackmd(s string) *SessionUpdateOne {
	suo.mutation.SetHackmd(s)
	return suo
}

// SetNillableHackmd sets the "hackmd" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableHackmd(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetHackmd(*s)
	}
	return suo
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (suo *SessionUpdateOne) Where(ps ...predicate.Session) *SessionUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SessionUpdateOne) Select(field string, fields ...string) *SessionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Session entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (_node *Session, err error) {
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Session.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for _, f := range fields {
			if !session.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != session.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Title(); ok {
		_spec.SetField(session.FieldTitle, field.TypeString, value)
	}
	if value, ok := suo.mutation.GetType(); ok {
		_spec.SetField(session.FieldType, field.TypeString, value)
	}
	if value, ok := suo.mutation.Speakers(); ok {
		_spec.SetField(session.FieldSpeakers, field.TypeJSON, value)
	}
	if value, ok := suo.mutation.AppendedSpeakers(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, session.FieldSpeakers, value)
		})
	}
	if value, ok := suo.mutation.Room(); ok {
		_spec.SetField(session.FieldRoom, field.TypeString, value)
	}
	if value, ok := suo.mutation.Broadcast(); ok {
		_spec.SetField(session.FieldBroadcast, field.TypeJSON, value)
	}
	if value, ok := suo.mutation.AppendedBroadcast(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, session.FieldBroadcast, value)
		})
	}
	if value, ok := suo.mutation.Start(); ok {
		_spec.SetField(session.FieldStart, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedStart(); ok {
		_spec.AddField(session.FieldStart, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.End(); ok {
		_spec.SetField(session.FieldEnd, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedEnd(); ok {
		_spec.AddField(session.FieldEnd, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.Slido(); ok {
		_spec.SetField(session.FieldSlido, field.TypeString, value)
	}
	if value, ok := suo.mutation.Slide(); ok {
		_spec.SetField(session.FieldSlide, field.TypeString, value)
	}
	if value, ok := suo.mutation.Hackmd(); ok {
		_spec.SetField(session.FieldHackmd, field.TypeString, value)
	}
	_node = &Session{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
