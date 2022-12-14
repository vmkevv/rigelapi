// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/municipio"
	"github.com/vmkevv/rigelapi/ent/predicate"
	"github.com/vmkevv/rigelapi/ent/provincia"
	"github.com/vmkevv/rigelapi/ent/school"
)

// MunicipioUpdate is the builder for updating Municipio entities.
type MunicipioUpdate struct {
	config
	hooks    []Hook
	mutation *MunicipioMutation
}

// Where appends a list predicates to the MunicipioUpdate builder.
func (mu *MunicipioUpdate) Where(ps ...predicate.Municipio) *MunicipioUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetName sets the "name" field.
func (mu *MunicipioUpdate) SetName(s string) *MunicipioUpdate {
	mu.mutation.SetName(s)
	return mu
}

// AddSchoolIDs adds the "schools" edge to the School entity by IDs.
func (mu *MunicipioUpdate) AddSchoolIDs(ids ...string) *MunicipioUpdate {
	mu.mutation.AddSchoolIDs(ids...)
	return mu
}

// AddSchools adds the "schools" edges to the School entity.
func (mu *MunicipioUpdate) AddSchools(s ...*School) *MunicipioUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mu.AddSchoolIDs(ids...)
}

// SetProvinciaID sets the "provincia" edge to the Provincia entity by ID.
func (mu *MunicipioUpdate) SetProvinciaID(id string) *MunicipioUpdate {
	mu.mutation.SetProvinciaID(id)
	return mu
}

// SetNillableProvinciaID sets the "provincia" edge to the Provincia entity by ID if the given value is not nil.
func (mu *MunicipioUpdate) SetNillableProvinciaID(id *string) *MunicipioUpdate {
	if id != nil {
		mu = mu.SetProvinciaID(*id)
	}
	return mu
}

// SetProvincia sets the "provincia" edge to the Provincia entity.
func (mu *MunicipioUpdate) SetProvincia(p *Provincia) *MunicipioUpdate {
	return mu.SetProvinciaID(p.ID)
}

// Mutation returns the MunicipioMutation object of the builder.
func (mu *MunicipioUpdate) Mutation() *MunicipioMutation {
	return mu.mutation
}

// ClearSchools clears all "schools" edges to the School entity.
func (mu *MunicipioUpdate) ClearSchools() *MunicipioUpdate {
	mu.mutation.ClearSchools()
	return mu
}

// RemoveSchoolIDs removes the "schools" edge to School entities by IDs.
func (mu *MunicipioUpdate) RemoveSchoolIDs(ids ...string) *MunicipioUpdate {
	mu.mutation.RemoveSchoolIDs(ids...)
	return mu
}

// RemoveSchools removes "schools" edges to School entities.
func (mu *MunicipioUpdate) RemoveSchools(s ...*School) *MunicipioUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mu.RemoveSchoolIDs(ids...)
}

// ClearProvincia clears the "provincia" edge to the Provincia entity.
func (mu *MunicipioUpdate) ClearProvincia() *MunicipioUpdate {
	mu.mutation.ClearProvincia()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MunicipioUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MunicipioMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MunicipioUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MunicipioUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MunicipioUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MunicipioUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   municipio.Table,
			Columns: municipio.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: municipio.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: municipio.FieldName,
		})
	}
	if mu.mutation.SchoolsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   municipio.SchoolsTable,
			Columns: []string{municipio.SchoolsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: school.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedSchoolsIDs(); len(nodes) > 0 && !mu.mutation.SchoolsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   municipio.SchoolsTable,
			Columns: []string{municipio.SchoolsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: school.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.SchoolsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   municipio.SchoolsTable,
			Columns: []string{municipio.SchoolsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: school.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ProvinciaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   municipio.ProvinciaTable,
			Columns: []string{municipio.ProvinciaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: provincia.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ProvinciaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   municipio.ProvinciaTable,
			Columns: []string{municipio.ProvinciaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: provincia.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{municipio.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MunicipioUpdateOne is the builder for updating a single Municipio entity.
type MunicipioUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MunicipioMutation
}

// SetName sets the "name" field.
func (muo *MunicipioUpdateOne) SetName(s string) *MunicipioUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// AddSchoolIDs adds the "schools" edge to the School entity by IDs.
func (muo *MunicipioUpdateOne) AddSchoolIDs(ids ...string) *MunicipioUpdateOne {
	muo.mutation.AddSchoolIDs(ids...)
	return muo
}

// AddSchools adds the "schools" edges to the School entity.
func (muo *MunicipioUpdateOne) AddSchools(s ...*School) *MunicipioUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return muo.AddSchoolIDs(ids...)
}

// SetProvinciaID sets the "provincia" edge to the Provincia entity by ID.
func (muo *MunicipioUpdateOne) SetProvinciaID(id string) *MunicipioUpdateOne {
	muo.mutation.SetProvinciaID(id)
	return muo
}

// SetNillableProvinciaID sets the "provincia" edge to the Provincia entity by ID if the given value is not nil.
func (muo *MunicipioUpdateOne) SetNillableProvinciaID(id *string) *MunicipioUpdateOne {
	if id != nil {
		muo = muo.SetProvinciaID(*id)
	}
	return muo
}

// SetProvincia sets the "provincia" edge to the Provincia entity.
func (muo *MunicipioUpdateOne) SetProvincia(p *Provincia) *MunicipioUpdateOne {
	return muo.SetProvinciaID(p.ID)
}

// Mutation returns the MunicipioMutation object of the builder.
func (muo *MunicipioUpdateOne) Mutation() *MunicipioMutation {
	return muo.mutation
}

// ClearSchools clears all "schools" edges to the School entity.
func (muo *MunicipioUpdateOne) ClearSchools() *MunicipioUpdateOne {
	muo.mutation.ClearSchools()
	return muo
}

// RemoveSchoolIDs removes the "schools" edge to School entities by IDs.
func (muo *MunicipioUpdateOne) RemoveSchoolIDs(ids ...string) *MunicipioUpdateOne {
	muo.mutation.RemoveSchoolIDs(ids...)
	return muo
}

// RemoveSchools removes "schools" edges to School entities.
func (muo *MunicipioUpdateOne) RemoveSchools(s ...*School) *MunicipioUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return muo.RemoveSchoolIDs(ids...)
}

// ClearProvincia clears the "provincia" edge to the Provincia entity.
func (muo *MunicipioUpdateOne) ClearProvincia() *MunicipioUpdateOne {
	muo.mutation.ClearProvincia()
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MunicipioUpdateOne) Select(field string, fields ...string) *MunicipioUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Municipio entity.
func (muo *MunicipioUpdateOne) Save(ctx context.Context) (*Municipio, error) {
	var (
		err  error
		node *Municipio
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MunicipioMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, muo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Municipio)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MunicipioMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MunicipioUpdateOne) SaveX(ctx context.Context) *Municipio {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MunicipioUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MunicipioUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MunicipioUpdateOne) sqlSave(ctx context.Context) (_node *Municipio, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   municipio.Table,
			Columns: municipio.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: municipio.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Municipio.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, municipio.FieldID)
		for _, f := range fields {
			if !municipio.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != municipio.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: municipio.FieldName,
		})
	}
	if muo.mutation.SchoolsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   municipio.SchoolsTable,
			Columns: []string{municipio.SchoolsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: school.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedSchoolsIDs(); len(nodes) > 0 && !muo.mutation.SchoolsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   municipio.SchoolsTable,
			Columns: []string{municipio.SchoolsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: school.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.SchoolsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   municipio.SchoolsTable,
			Columns: []string{municipio.SchoolsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: school.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ProvinciaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   municipio.ProvinciaTable,
			Columns: []string{municipio.ProvinciaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: provincia.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ProvinciaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   municipio.ProvinciaTable,
			Columns: []string{municipio.ProvinciaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: provincia.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Municipio{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{municipio.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
