// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/attendancedaysyncs"
	"github.com/vmkevv/rigelapi/ent/teacher"
)

// AttendanceDaySyncsCreate is the builder for creating a AttendanceDaySyncs entity.
type AttendanceDaySyncsCreate struct {
	config
	mutation *AttendanceDaySyncsMutation
	hooks    []Hook
}

// SetLastSyncID sets the "last_sync_id" field.
func (adsc *AttendanceDaySyncsCreate) SetLastSyncID(s string) *AttendanceDaySyncsCreate {
	adsc.mutation.SetLastSyncID(s)
	return adsc
}

// SetID sets the "id" field.
func (adsc *AttendanceDaySyncsCreate) SetID(s string) *AttendanceDaySyncsCreate {
	adsc.mutation.SetID(s)
	return adsc
}

// SetTeacherID sets the "teacher" edge to the Teacher entity by ID.
func (adsc *AttendanceDaySyncsCreate) SetTeacherID(id string) *AttendanceDaySyncsCreate {
	adsc.mutation.SetTeacherID(id)
	return adsc
}

// SetNillableTeacherID sets the "teacher" edge to the Teacher entity by ID if the given value is not nil.
func (adsc *AttendanceDaySyncsCreate) SetNillableTeacherID(id *string) *AttendanceDaySyncsCreate {
	if id != nil {
		adsc = adsc.SetTeacherID(*id)
	}
	return adsc
}

// SetTeacher sets the "teacher" edge to the Teacher entity.
func (adsc *AttendanceDaySyncsCreate) SetTeacher(t *Teacher) *AttendanceDaySyncsCreate {
	return adsc.SetTeacherID(t.ID)
}

// Mutation returns the AttendanceDaySyncsMutation object of the builder.
func (adsc *AttendanceDaySyncsCreate) Mutation() *AttendanceDaySyncsMutation {
	return adsc.mutation
}

// Save creates the AttendanceDaySyncs in the database.
func (adsc *AttendanceDaySyncsCreate) Save(ctx context.Context) (*AttendanceDaySyncs, error) {
	var (
		err  error
		node *AttendanceDaySyncs
	)
	if len(adsc.hooks) == 0 {
		if err = adsc.check(); err != nil {
			return nil, err
		}
		node, err = adsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AttendanceDaySyncsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = adsc.check(); err != nil {
				return nil, err
			}
			adsc.mutation = mutation
			if node, err = adsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(adsc.hooks) - 1; i >= 0; i-- {
			if adsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = adsc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, adsc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AttendanceDaySyncs)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AttendanceDaySyncsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (adsc *AttendanceDaySyncsCreate) SaveX(ctx context.Context) *AttendanceDaySyncs {
	v, err := adsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (adsc *AttendanceDaySyncsCreate) Exec(ctx context.Context) error {
	_, err := adsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (adsc *AttendanceDaySyncsCreate) ExecX(ctx context.Context) {
	if err := adsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (adsc *AttendanceDaySyncsCreate) check() error {
	if _, ok := adsc.mutation.LastSyncID(); !ok {
		return &ValidationError{Name: "last_sync_id", err: errors.New(`ent: missing required field "AttendanceDaySyncs.last_sync_id"`)}
	}
	return nil
}

func (adsc *AttendanceDaySyncsCreate) sqlSave(ctx context.Context) (*AttendanceDaySyncs, error) {
	_node, _spec := adsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, adsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected AttendanceDaySyncs.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (adsc *AttendanceDaySyncsCreate) createSpec() (*AttendanceDaySyncs, *sqlgraph.CreateSpec) {
	var (
		_node = &AttendanceDaySyncs{config: adsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: attendancedaysyncs.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: attendancedaysyncs.FieldID,
			},
		}
	)
	if id, ok := adsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := adsc.mutation.LastSyncID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: attendancedaysyncs.FieldLastSyncID,
		})
		_node.LastSyncID = value
	}
	if nodes := adsc.mutation.TeacherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendancedaysyncs.TeacherTable,
			Columns: []string{attendancedaysyncs.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teacher.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.teacher_attendance_day_syncs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AttendanceDaySyncsCreateBulk is the builder for creating many AttendanceDaySyncs entities in bulk.
type AttendanceDaySyncsCreateBulk struct {
	config
	builders []*AttendanceDaySyncsCreate
}

// Save creates the AttendanceDaySyncs entities in the database.
func (adscb *AttendanceDaySyncsCreateBulk) Save(ctx context.Context) ([]*AttendanceDaySyncs, error) {
	specs := make([]*sqlgraph.CreateSpec, len(adscb.builders))
	nodes := make([]*AttendanceDaySyncs, len(adscb.builders))
	mutators := make([]Mutator, len(adscb.builders))
	for i := range adscb.builders {
		func(i int, root context.Context) {
			builder := adscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttendanceDaySyncsMutation)
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
					_, err = mutators[i+1].Mutate(root, adscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, adscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, adscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (adscb *AttendanceDaySyncsCreateBulk) SaveX(ctx context.Context) []*AttendanceDaySyncs {
	v, err := adscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (adscb *AttendanceDaySyncsCreateBulk) Exec(ctx context.Context) error {
	_, err := adscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (adscb *AttendanceDaySyncsCreateBulk) ExecX(ctx context.Context) {
	if err := adscb.Exec(ctx); err != nil {
		panic(err)
	}
}
