// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/apperror"
)

// AppErrorCreate is the builder for creating a AppError entity.
type AppErrorCreate struct {
	config
	mutation *AppErrorMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (aec *AppErrorCreate) SetUserID(s string) *AppErrorCreate {
	aec.mutation.SetUserID(s)
	return aec
}

// SetCause sets the "cause" field.
func (aec *AppErrorCreate) SetCause(s string) *AppErrorCreate {
	aec.mutation.SetCause(s)
	return aec
}

// SetErrorMsg sets the "error_msg" field.
func (aec *AppErrorCreate) SetErrorMsg(s string) *AppErrorCreate {
	aec.mutation.SetErrorMsg(s)
	return aec
}

// SetErrorStack sets the "error_stack" field.
func (aec *AppErrorCreate) SetErrorStack(s string) *AppErrorCreate {
	aec.mutation.SetErrorStack(s)
	return aec
}

// SetID sets the "id" field.
func (aec *AppErrorCreate) SetID(s string) *AppErrorCreate {
	aec.mutation.SetID(s)
	return aec
}

// Mutation returns the AppErrorMutation object of the builder.
func (aec *AppErrorCreate) Mutation() *AppErrorMutation {
	return aec.mutation
}

// Save creates the AppError in the database.
func (aec *AppErrorCreate) Save(ctx context.Context) (*AppError, error) {
	var (
		err  error
		node *AppError
	)
	if len(aec.hooks) == 0 {
		if err = aec.check(); err != nil {
			return nil, err
		}
		node, err = aec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppErrorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aec.check(); err != nil {
				return nil, err
			}
			aec.mutation = mutation
			if node, err = aec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(aec.hooks) - 1; i >= 0; i-- {
			if aec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, aec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppError)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppErrorMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (aec *AppErrorCreate) SaveX(ctx context.Context) *AppError {
	v, err := aec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aec *AppErrorCreate) Exec(ctx context.Context) error {
	_, err := aec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aec *AppErrorCreate) ExecX(ctx context.Context) {
	if err := aec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aec *AppErrorCreate) check() error {
	if _, ok := aec.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "AppError.user_id"`)}
	}
	if _, ok := aec.mutation.Cause(); !ok {
		return &ValidationError{Name: "cause", err: errors.New(`ent: missing required field "AppError.cause"`)}
	}
	if _, ok := aec.mutation.ErrorMsg(); !ok {
		return &ValidationError{Name: "error_msg", err: errors.New(`ent: missing required field "AppError.error_msg"`)}
	}
	if _, ok := aec.mutation.ErrorStack(); !ok {
		return &ValidationError{Name: "error_stack", err: errors.New(`ent: missing required field "AppError.error_stack"`)}
	}
	return nil
}

func (aec *AppErrorCreate) sqlSave(ctx context.Context) (*AppError, error) {
	_node, _spec := aec.createSpec()
	if err := sqlgraph.CreateNode(ctx, aec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected AppError.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (aec *AppErrorCreate) createSpec() (*AppError, *sqlgraph.CreateSpec) {
	var (
		_node = &AppError{config: aec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: apperror.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: apperror.FieldID,
			},
		}
	)
	if id, ok := aec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := aec.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apperror.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := aec.mutation.Cause(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apperror.FieldCause,
		})
		_node.Cause = value
	}
	if value, ok := aec.mutation.ErrorMsg(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apperror.FieldErrorMsg,
		})
		_node.ErrorMsg = value
	}
	if value, ok := aec.mutation.ErrorStack(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apperror.FieldErrorStack,
		})
		_node.ErrorStack = value
	}
	return _node, _spec
}

// AppErrorCreateBulk is the builder for creating many AppError entities in bulk.
type AppErrorCreateBulk struct {
	config
	builders []*AppErrorCreate
}

// Save creates the AppError entities in the database.
func (aecb *AppErrorCreateBulk) Save(ctx context.Context) ([]*AppError, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aecb.builders))
	nodes := make([]*AppError, len(aecb.builders))
	mutators := make([]Mutator, len(aecb.builders))
	for i := range aecb.builders {
		func(i int, root context.Context) {
			builder := aecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppErrorMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, aecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, aecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aecb *AppErrorCreateBulk) SaveX(ctx context.Context) []*AppError {
	v, err := aecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aecb *AppErrorCreateBulk) Exec(ctx context.Context) error {
	_, err := aecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aecb *AppErrorCreateBulk) ExecX(ctx context.Context) {
	if err := aecb.Exec(ctx); err != nil {
		panic(err)
	}
}