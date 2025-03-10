// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/special"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SpecialCreate is the builder for creating a Special entity.
type SpecialCreate struct {
	config
	mutation *SpecialMutation
	hooks    []Hook
}

// SetData sets the "data" field.
func (sc *SpecialCreate) SetData(s string) *SpecialCreate {
	sc.mutation.SetData(s)
	return sc
}

// SetNillableData sets the "data" field if the given value is not nil.
func (sc *SpecialCreate) SetNillableData(s *string) *SpecialCreate {
	if s != nil {
		sc.SetData(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SpecialCreate) SetID(s string) *SpecialCreate {
	sc.mutation.SetID(s)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SpecialCreate) SetNillableID(s *string) *SpecialCreate {
	if s != nil {
		sc.SetID(*s)
	}
	return sc
}

// Mutation returns the SpecialMutation object of the builder.
func (sc *SpecialCreate) Mutation() *SpecialMutation {
	return sc.mutation
}

// Save creates the Special in the database.
func (sc *SpecialCreate) Save(ctx context.Context) (*Special, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SpecialCreate) SaveX(ctx context.Context) *Special {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SpecialCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SpecialCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SpecialCreate) defaults() {
	if _, ok := sc.mutation.Data(); !ok {
		v := special.DefaultData
		sc.mutation.SetData(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := special.DefaultID
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SpecialCreate) check() error {
	if _, ok := sc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`ent: missing required field "Special.data"`)}
	}
	return nil
}

func (sc *SpecialCreate) sqlSave(ctx context.Context) (*Special, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Special.ID type: %T", _spec.ID.Value)
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SpecialCreate) createSpec() (*Special, *sqlgraph.CreateSpec) {
	var (
		_node = &Special{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(special.Table, sqlgraph.NewFieldSpec(special.FieldID, field.TypeString))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.Data(); ok {
		_spec.SetField(special.FieldData, field.TypeString, value)
		_node.Data = value
	}
	return _node, _spec
}

// SpecialCreateBulk is the builder for creating many Special entities in bulk.
type SpecialCreateBulk struct {
	config
	err      error
	builders []*SpecialCreate
}

// Save creates the Special entities in the database.
func (scb *SpecialCreateBulk) Save(ctx context.Context) ([]*Special, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Special, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SpecialMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SpecialCreateBulk) SaveX(ctx context.Context) []*Special {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SpecialCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SpecialCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
