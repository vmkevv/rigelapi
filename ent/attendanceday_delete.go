// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/attendanceday"
	"github.com/vmkevv/rigelapi/ent/predicate"
)

// AttendanceDayDelete is the builder for deleting a AttendanceDay entity.
type AttendanceDayDelete struct {
	config
	hooks    []Hook
	mutation *AttendanceDayMutation
}

// Where appends a list predicates to the AttendanceDayDelete builder.
func (add *AttendanceDayDelete) Where(ps ...predicate.AttendanceDay) *AttendanceDayDelete {
	add.mutation.Where(ps...)
	return add
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (add *AttendanceDayDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(add.hooks) == 0 {
		affected, err = add.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AttendanceDayMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			add.mutation = mutation
			affected, err = add.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(add.hooks) - 1; i >= 0; i-- {
			if add.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = add.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, add.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (add *AttendanceDayDelete) ExecX(ctx context.Context) int {
	n, err := add.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (add *AttendanceDayDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: attendanceday.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: attendanceday.FieldID,
			},
		},
	}
	if ps := add.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, add.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AttendanceDayDeleteOne is the builder for deleting a single AttendanceDay entity.
type AttendanceDayDeleteOne struct {
	add *AttendanceDayDelete
}

// Exec executes the deletion query.
func (addo *AttendanceDayDeleteOne) Exec(ctx context.Context) error {
	n, err := addo.add.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{attendanceday.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (addo *AttendanceDayDeleteOne) ExecX(ctx context.Context) {
	addo.add.ExecX(ctx)
}
