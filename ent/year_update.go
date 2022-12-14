// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/area"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/period"
	"github.com/vmkevv/rigelapi/ent/predicate"
	"github.com/vmkevv/rigelapi/ent/year"
)

// YearUpdate is the builder for updating Year entities.
type YearUpdate struct {
	config
	hooks    []Hook
	mutation *YearMutation
}

// Where appends a list predicates to the YearUpdate builder.
func (yu *YearUpdate) Where(ps ...predicate.Year) *YearUpdate {
	yu.mutation.Where(ps...)
	return yu
}

// SetValue sets the "value" field.
func (yu *YearUpdate) SetValue(i int) *YearUpdate {
	yu.mutation.ResetValue()
	yu.mutation.SetValue(i)
	return yu
}

// AddValue adds i to the "value" field.
func (yu *YearUpdate) AddValue(i int) *YearUpdate {
	yu.mutation.AddValue(i)
	return yu
}

// AddClassIDs adds the "classes" edge to the Class entity by IDs.
func (yu *YearUpdate) AddClassIDs(ids ...string) *YearUpdate {
	yu.mutation.AddClassIDs(ids...)
	return yu
}

// AddClasses adds the "classes" edges to the Class entity.
func (yu *YearUpdate) AddClasses(c ...*Class) *YearUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return yu.AddClassIDs(ids...)
}

// AddPeriodIDs adds the "periods" edge to the Period entity by IDs.
func (yu *YearUpdate) AddPeriodIDs(ids ...string) *YearUpdate {
	yu.mutation.AddPeriodIDs(ids...)
	return yu
}

// AddPeriods adds the "periods" edges to the Period entity.
func (yu *YearUpdate) AddPeriods(p ...*Period) *YearUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return yu.AddPeriodIDs(ids...)
}

// AddAreaIDs adds the "areas" edge to the Area entity by IDs.
func (yu *YearUpdate) AddAreaIDs(ids ...string) *YearUpdate {
	yu.mutation.AddAreaIDs(ids...)
	return yu
}

// AddAreas adds the "areas" edges to the Area entity.
func (yu *YearUpdate) AddAreas(a ...*Area) *YearUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return yu.AddAreaIDs(ids...)
}

// Mutation returns the YearMutation object of the builder.
func (yu *YearUpdate) Mutation() *YearMutation {
	return yu.mutation
}

// ClearClasses clears all "classes" edges to the Class entity.
func (yu *YearUpdate) ClearClasses() *YearUpdate {
	yu.mutation.ClearClasses()
	return yu
}

// RemoveClassIDs removes the "classes" edge to Class entities by IDs.
func (yu *YearUpdate) RemoveClassIDs(ids ...string) *YearUpdate {
	yu.mutation.RemoveClassIDs(ids...)
	return yu
}

// RemoveClasses removes "classes" edges to Class entities.
func (yu *YearUpdate) RemoveClasses(c ...*Class) *YearUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return yu.RemoveClassIDs(ids...)
}

// ClearPeriods clears all "periods" edges to the Period entity.
func (yu *YearUpdate) ClearPeriods() *YearUpdate {
	yu.mutation.ClearPeriods()
	return yu
}

// RemovePeriodIDs removes the "periods" edge to Period entities by IDs.
func (yu *YearUpdate) RemovePeriodIDs(ids ...string) *YearUpdate {
	yu.mutation.RemovePeriodIDs(ids...)
	return yu
}

// RemovePeriods removes "periods" edges to Period entities.
func (yu *YearUpdate) RemovePeriods(p ...*Period) *YearUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return yu.RemovePeriodIDs(ids...)
}

// ClearAreas clears all "areas" edges to the Area entity.
func (yu *YearUpdate) ClearAreas() *YearUpdate {
	yu.mutation.ClearAreas()
	return yu
}

// RemoveAreaIDs removes the "areas" edge to Area entities by IDs.
func (yu *YearUpdate) RemoveAreaIDs(ids ...string) *YearUpdate {
	yu.mutation.RemoveAreaIDs(ids...)
	return yu
}

// RemoveAreas removes "areas" edges to Area entities.
func (yu *YearUpdate) RemoveAreas(a ...*Area) *YearUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return yu.RemoveAreaIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (yu *YearUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(yu.hooks) == 0 {
		affected, err = yu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*YearMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			yu.mutation = mutation
			affected, err = yu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(yu.hooks) - 1; i >= 0; i-- {
			if yu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = yu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, yu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (yu *YearUpdate) SaveX(ctx context.Context) int {
	affected, err := yu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (yu *YearUpdate) Exec(ctx context.Context) error {
	_, err := yu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (yu *YearUpdate) ExecX(ctx context.Context) {
	if err := yu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (yu *YearUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   year.Table,
			Columns: year.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: year.FieldID,
			},
		},
	}
	if ps := yu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := yu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: year.FieldValue,
		})
	}
	if value, ok := yu.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: year.FieldValue,
		})
	}
	if yu.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.ClassesTable,
			Columns: []string{year.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: class.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := yu.mutation.RemovedClassesIDs(); len(nodes) > 0 && !yu.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.ClassesTable,
			Columns: []string{year.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := yu.mutation.ClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.ClassesTable,
			Columns: []string{year.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if yu.mutation.PeriodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.PeriodsTable,
			Columns: []string{year.PeriodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: period.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := yu.mutation.RemovedPeriodsIDs(); len(nodes) > 0 && !yu.mutation.PeriodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.PeriodsTable,
			Columns: []string{year.PeriodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: period.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := yu.mutation.PeriodsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.PeriodsTable,
			Columns: []string{year.PeriodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: period.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if yu.mutation.AreasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.AreasTable,
			Columns: []string{year.AreasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: area.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := yu.mutation.RemovedAreasIDs(); len(nodes) > 0 && !yu.mutation.AreasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.AreasTable,
			Columns: []string{year.AreasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: area.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := yu.mutation.AreasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.AreasTable,
			Columns: []string{year.AreasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: area.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, yu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{year.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// YearUpdateOne is the builder for updating a single Year entity.
type YearUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *YearMutation
}

// SetValue sets the "value" field.
func (yuo *YearUpdateOne) SetValue(i int) *YearUpdateOne {
	yuo.mutation.ResetValue()
	yuo.mutation.SetValue(i)
	return yuo
}

// AddValue adds i to the "value" field.
func (yuo *YearUpdateOne) AddValue(i int) *YearUpdateOne {
	yuo.mutation.AddValue(i)
	return yuo
}

// AddClassIDs adds the "classes" edge to the Class entity by IDs.
func (yuo *YearUpdateOne) AddClassIDs(ids ...string) *YearUpdateOne {
	yuo.mutation.AddClassIDs(ids...)
	return yuo
}

// AddClasses adds the "classes" edges to the Class entity.
func (yuo *YearUpdateOne) AddClasses(c ...*Class) *YearUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return yuo.AddClassIDs(ids...)
}

// AddPeriodIDs adds the "periods" edge to the Period entity by IDs.
func (yuo *YearUpdateOne) AddPeriodIDs(ids ...string) *YearUpdateOne {
	yuo.mutation.AddPeriodIDs(ids...)
	return yuo
}

// AddPeriods adds the "periods" edges to the Period entity.
func (yuo *YearUpdateOne) AddPeriods(p ...*Period) *YearUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return yuo.AddPeriodIDs(ids...)
}

// AddAreaIDs adds the "areas" edge to the Area entity by IDs.
func (yuo *YearUpdateOne) AddAreaIDs(ids ...string) *YearUpdateOne {
	yuo.mutation.AddAreaIDs(ids...)
	return yuo
}

// AddAreas adds the "areas" edges to the Area entity.
func (yuo *YearUpdateOne) AddAreas(a ...*Area) *YearUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return yuo.AddAreaIDs(ids...)
}

// Mutation returns the YearMutation object of the builder.
func (yuo *YearUpdateOne) Mutation() *YearMutation {
	return yuo.mutation
}

// ClearClasses clears all "classes" edges to the Class entity.
func (yuo *YearUpdateOne) ClearClasses() *YearUpdateOne {
	yuo.mutation.ClearClasses()
	return yuo
}

// RemoveClassIDs removes the "classes" edge to Class entities by IDs.
func (yuo *YearUpdateOne) RemoveClassIDs(ids ...string) *YearUpdateOne {
	yuo.mutation.RemoveClassIDs(ids...)
	return yuo
}

// RemoveClasses removes "classes" edges to Class entities.
func (yuo *YearUpdateOne) RemoveClasses(c ...*Class) *YearUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return yuo.RemoveClassIDs(ids...)
}

// ClearPeriods clears all "periods" edges to the Period entity.
func (yuo *YearUpdateOne) ClearPeriods() *YearUpdateOne {
	yuo.mutation.ClearPeriods()
	return yuo
}

// RemovePeriodIDs removes the "periods" edge to Period entities by IDs.
func (yuo *YearUpdateOne) RemovePeriodIDs(ids ...string) *YearUpdateOne {
	yuo.mutation.RemovePeriodIDs(ids...)
	return yuo
}

// RemovePeriods removes "periods" edges to Period entities.
func (yuo *YearUpdateOne) RemovePeriods(p ...*Period) *YearUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return yuo.RemovePeriodIDs(ids...)
}

// ClearAreas clears all "areas" edges to the Area entity.
func (yuo *YearUpdateOne) ClearAreas() *YearUpdateOne {
	yuo.mutation.ClearAreas()
	return yuo
}

// RemoveAreaIDs removes the "areas" edge to Area entities by IDs.
func (yuo *YearUpdateOne) RemoveAreaIDs(ids ...string) *YearUpdateOne {
	yuo.mutation.RemoveAreaIDs(ids...)
	return yuo
}

// RemoveAreas removes "areas" edges to Area entities.
func (yuo *YearUpdateOne) RemoveAreas(a ...*Area) *YearUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return yuo.RemoveAreaIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (yuo *YearUpdateOne) Select(field string, fields ...string) *YearUpdateOne {
	yuo.fields = append([]string{field}, fields...)
	return yuo
}

// Save executes the query and returns the updated Year entity.
func (yuo *YearUpdateOne) Save(ctx context.Context) (*Year, error) {
	var (
		err  error
		node *Year
	)
	if len(yuo.hooks) == 0 {
		node, err = yuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*YearMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			yuo.mutation = mutation
			node, err = yuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(yuo.hooks) - 1; i >= 0; i-- {
			if yuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = yuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, yuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Year)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from YearMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (yuo *YearUpdateOne) SaveX(ctx context.Context) *Year {
	node, err := yuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (yuo *YearUpdateOne) Exec(ctx context.Context) error {
	_, err := yuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (yuo *YearUpdateOne) ExecX(ctx context.Context) {
	if err := yuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (yuo *YearUpdateOne) sqlSave(ctx context.Context) (_node *Year, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   year.Table,
			Columns: year.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: year.FieldID,
			},
		},
	}
	id, ok := yuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Year.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := yuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, year.FieldID)
		for _, f := range fields {
			if !year.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != year.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := yuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := yuo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: year.FieldValue,
		})
	}
	if value, ok := yuo.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: year.FieldValue,
		})
	}
	if yuo.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.ClassesTable,
			Columns: []string{year.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: class.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := yuo.mutation.RemovedClassesIDs(); len(nodes) > 0 && !yuo.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.ClassesTable,
			Columns: []string{year.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := yuo.mutation.ClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.ClassesTable,
			Columns: []string{year.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if yuo.mutation.PeriodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.PeriodsTable,
			Columns: []string{year.PeriodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: period.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := yuo.mutation.RemovedPeriodsIDs(); len(nodes) > 0 && !yuo.mutation.PeriodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.PeriodsTable,
			Columns: []string{year.PeriodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: period.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := yuo.mutation.PeriodsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.PeriodsTable,
			Columns: []string{year.PeriodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: period.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if yuo.mutation.AreasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.AreasTable,
			Columns: []string{year.AreasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: area.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := yuo.mutation.RemovedAreasIDs(); len(nodes) > 0 && !yuo.mutation.AreasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.AreasTable,
			Columns: []string{year.AreasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: area.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := yuo.mutation.AreasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.AreasTable,
			Columns: []string{year.AreasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: area.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Year{config: yuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, yuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{year.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
