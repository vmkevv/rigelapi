// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/attendancesync"
	"github.com/vmkevv/rigelapi/ent/predicate"
	"github.com/vmkevv/rigelapi/ent/teacher"
)

// AttendanceSyncQuery is the builder for querying AttendanceSync entities.
type AttendanceSyncQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	predicates  []predicate.AttendanceSync
	withTeacher *TeacherQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AttendanceSyncQuery builder.
func (asq *AttendanceSyncQuery) Where(ps ...predicate.AttendanceSync) *AttendanceSyncQuery {
	asq.predicates = append(asq.predicates, ps...)
	return asq
}

// Limit adds a limit step to the query.
func (asq *AttendanceSyncQuery) Limit(limit int) *AttendanceSyncQuery {
	asq.limit = &limit
	return asq
}

// Offset adds an offset step to the query.
func (asq *AttendanceSyncQuery) Offset(offset int) *AttendanceSyncQuery {
	asq.offset = &offset
	return asq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (asq *AttendanceSyncQuery) Unique(unique bool) *AttendanceSyncQuery {
	asq.unique = &unique
	return asq
}

// Order adds an order step to the query.
func (asq *AttendanceSyncQuery) Order(o ...OrderFunc) *AttendanceSyncQuery {
	asq.order = append(asq.order, o...)
	return asq
}

// QueryTeacher chains the current query on the "teacher" edge.
func (asq *AttendanceSyncQuery) QueryTeacher() *TeacherQuery {
	query := &TeacherQuery{config: asq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := asq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := asq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attendancesync.Table, attendancesync.FieldID, selector),
			sqlgraph.To(teacher.Table, teacher.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attendancesync.TeacherTable, attendancesync.TeacherColumn),
		)
		fromU = sqlgraph.SetNeighbors(asq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AttendanceSync entity from the query.
// Returns a *NotFoundError when no AttendanceSync was found.
func (asq *AttendanceSyncQuery) First(ctx context.Context) (*AttendanceSync, error) {
	nodes, err := asq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{attendancesync.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (asq *AttendanceSyncQuery) FirstX(ctx context.Context) *AttendanceSync {
	node, err := asq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AttendanceSync ID from the query.
// Returns a *NotFoundError when no AttendanceSync ID was found.
func (asq *AttendanceSyncQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = asq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{attendancesync.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (asq *AttendanceSyncQuery) FirstIDX(ctx context.Context) string {
	id, err := asq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AttendanceSync entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AttendanceSync entity is found.
// Returns a *NotFoundError when no AttendanceSync entities are found.
func (asq *AttendanceSyncQuery) Only(ctx context.Context) (*AttendanceSync, error) {
	nodes, err := asq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{attendancesync.Label}
	default:
		return nil, &NotSingularError{attendancesync.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (asq *AttendanceSyncQuery) OnlyX(ctx context.Context) *AttendanceSync {
	node, err := asq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AttendanceSync ID in the query.
// Returns a *NotSingularError when more than one AttendanceSync ID is found.
// Returns a *NotFoundError when no entities are found.
func (asq *AttendanceSyncQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = asq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{attendancesync.Label}
	default:
		err = &NotSingularError{attendancesync.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (asq *AttendanceSyncQuery) OnlyIDX(ctx context.Context) string {
	id, err := asq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AttendanceSyncs.
func (asq *AttendanceSyncQuery) All(ctx context.Context) ([]*AttendanceSync, error) {
	if err := asq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return asq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (asq *AttendanceSyncQuery) AllX(ctx context.Context) []*AttendanceSync {
	nodes, err := asq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AttendanceSync IDs.
func (asq *AttendanceSyncQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := asq.Select(attendancesync.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (asq *AttendanceSyncQuery) IDsX(ctx context.Context) []string {
	ids, err := asq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (asq *AttendanceSyncQuery) Count(ctx context.Context) (int, error) {
	if err := asq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return asq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (asq *AttendanceSyncQuery) CountX(ctx context.Context) int {
	count, err := asq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (asq *AttendanceSyncQuery) Exist(ctx context.Context) (bool, error) {
	if err := asq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return asq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (asq *AttendanceSyncQuery) ExistX(ctx context.Context) bool {
	exist, err := asq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AttendanceSyncQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (asq *AttendanceSyncQuery) Clone() *AttendanceSyncQuery {
	if asq == nil {
		return nil
	}
	return &AttendanceSyncQuery{
		config:      asq.config,
		limit:       asq.limit,
		offset:      asq.offset,
		order:       append([]OrderFunc{}, asq.order...),
		predicates:  append([]predicate.AttendanceSync{}, asq.predicates...),
		withTeacher: asq.withTeacher.Clone(),
		// clone intermediate query.
		sql:    asq.sql.Clone(),
		path:   asq.path,
		unique: asq.unique,
	}
}

// WithTeacher tells the query-builder to eager-load the nodes that are connected to
// the "teacher" edge. The optional arguments are used to configure the query builder of the edge.
func (asq *AttendanceSyncQuery) WithTeacher(opts ...func(*TeacherQuery)) *AttendanceSyncQuery {
	query := &TeacherQuery{config: asq.config}
	for _, opt := range opts {
		opt(query)
	}
	asq.withTeacher = query
	return asq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		LastSyncID string `json:"last_sync_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AttendanceSync.Query().
//		GroupBy(attendancesync.FieldLastSyncID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (asq *AttendanceSyncQuery) GroupBy(field string, fields ...string) *AttendanceSyncGroupBy {
	grbuild := &AttendanceSyncGroupBy{config: asq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := asq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return asq.sqlQuery(ctx), nil
	}
	grbuild.label = attendancesync.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		LastSyncID string `json:"last_sync_id,omitempty"`
//	}
//
//	client.AttendanceSync.Query().
//		Select(attendancesync.FieldLastSyncID).
//		Scan(ctx, &v)
func (asq *AttendanceSyncQuery) Select(fields ...string) *AttendanceSyncSelect {
	asq.fields = append(asq.fields, fields...)
	selbuild := &AttendanceSyncSelect{AttendanceSyncQuery: asq}
	selbuild.label = attendancesync.Label
	selbuild.flds, selbuild.scan = &asq.fields, selbuild.Scan
	return selbuild
}

func (asq *AttendanceSyncQuery) prepareQuery(ctx context.Context) error {
	for _, f := range asq.fields {
		if !attendancesync.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if asq.path != nil {
		prev, err := asq.path(ctx)
		if err != nil {
			return err
		}
		asq.sql = prev
	}
	return nil
}

func (asq *AttendanceSyncQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AttendanceSync, error) {
	var (
		nodes       = []*AttendanceSync{}
		withFKs     = asq.withFKs
		_spec       = asq.querySpec()
		loadedTypes = [1]bool{
			asq.withTeacher != nil,
		}
	)
	if asq.withTeacher != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, attendancesync.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AttendanceSync).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AttendanceSync{config: asq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, asq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := asq.withTeacher; query != nil {
		if err := asq.loadTeacher(ctx, query, nodes, nil,
			func(n *AttendanceSync, e *Teacher) { n.Edges.Teacher = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (asq *AttendanceSyncQuery) loadTeacher(ctx context.Context, query *TeacherQuery, nodes []*AttendanceSync, init func(*AttendanceSync), assign func(*AttendanceSync, *Teacher)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*AttendanceSync)
	for i := range nodes {
		if nodes[i].teacher_attendance_syncs == nil {
			continue
		}
		fk := *nodes[i].teacher_attendance_syncs
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(teacher.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "teacher_attendance_syncs" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (asq *AttendanceSyncQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := asq.querySpec()
	_spec.Node.Columns = asq.fields
	if len(asq.fields) > 0 {
		_spec.Unique = asq.unique != nil && *asq.unique
	}
	return sqlgraph.CountNodes(ctx, asq.driver, _spec)
}

func (asq *AttendanceSyncQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := asq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (asq *AttendanceSyncQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   attendancesync.Table,
			Columns: attendancesync.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: attendancesync.FieldID,
			},
		},
		From:   asq.sql,
		Unique: true,
	}
	if unique := asq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := asq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attendancesync.FieldID)
		for i := range fields {
			if fields[i] != attendancesync.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := asq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := asq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := asq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := asq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (asq *AttendanceSyncQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(asq.driver.Dialect())
	t1 := builder.Table(attendancesync.Table)
	columns := asq.fields
	if len(columns) == 0 {
		columns = attendancesync.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if asq.sql != nil {
		selector = asq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if asq.unique != nil && *asq.unique {
		selector.Distinct()
	}
	for _, p := range asq.predicates {
		p(selector)
	}
	for _, p := range asq.order {
		p(selector)
	}
	if offset := asq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := asq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AttendanceSyncGroupBy is the group-by builder for AttendanceSync entities.
type AttendanceSyncGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (asgb *AttendanceSyncGroupBy) Aggregate(fns ...AggregateFunc) *AttendanceSyncGroupBy {
	asgb.fns = append(asgb.fns, fns...)
	return asgb
}

// Scan applies the group-by query and scans the result into the given value.
func (asgb *AttendanceSyncGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := asgb.path(ctx)
	if err != nil {
		return err
	}
	asgb.sql = query
	return asgb.sqlScan(ctx, v)
}

func (asgb *AttendanceSyncGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range asgb.fields {
		if !attendancesync.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := asgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := asgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (asgb *AttendanceSyncGroupBy) sqlQuery() *sql.Selector {
	selector := asgb.sql.Select()
	aggregation := make([]string, 0, len(asgb.fns))
	for _, fn := range asgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(asgb.fields)+len(asgb.fns))
		for _, f := range asgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(asgb.fields...)...)
}

// AttendanceSyncSelect is the builder for selecting fields of AttendanceSync entities.
type AttendanceSyncSelect struct {
	*AttendanceSyncQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ass *AttendanceSyncSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ass.prepareQuery(ctx); err != nil {
		return err
	}
	ass.sql = ass.AttendanceSyncQuery.sqlQuery(ctx)
	return ass.sqlScan(ctx, v)
}

func (ass *AttendanceSyncSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ass.sql.Query()
	if err := ass.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
