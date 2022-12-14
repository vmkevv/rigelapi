// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/attendancedaysyncs"
	"github.com/vmkevv/rigelapi/ent/predicate"
	"github.com/vmkevv/rigelapi/ent/teacher"
)

// AttendanceDaySyncsQuery is the builder for querying AttendanceDaySyncs entities.
type AttendanceDaySyncsQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	predicates  []predicate.AttendanceDaySyncs
	withTeacher *TeacherQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AttendanceDaySyncsQuery builder.
func (adsq *AttendanceDaySyncsQuery) Where(ps ...predicate.AttendanceDaySyncs) *AttendanceDaySyncsQuery {
	adsq.predicates = append(adsq.predicates, ps...)
	return adsq
}

// Limit adds a limit step to the query.
func (adsq *AttendanceDaySyncsQuery) Limit(limit int) *AttendanceDaySyncsQuery {
	adsq.limit = &limit
	return adsq
}

// Offset adds an offset step to the query.
func (adsq *AttendanceDaySyncsQuery) Offset(offset int) *AttendanceDaySyncsQuery {
	adsq.offset = &offset
	return adsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (adsq *AttendanceDaySyncsQuery) Unique(unique bool) *AttendanceDaySyncsQuery {
	adsq.unique = &unique
	return adsq
}

// Order adds an order step to the query.
func (adsq *AttendanceDaySyncsQuery) Order(o ...OrderFunc) *AttendanceDaySyncsQuery {
	adsq.order = append(adsq.order, o...)
	return adsq
}

// QueryTeacher chains the current query on the "teacher" edge.
func (adsq *AttendanceDaySyncsQuery) QueryTeacher() *TeacherQuery {
	query := &TeacherQuery{config: adsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := adsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := adsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attendancedaysyncs.Table, attendancedaysyncs.FieldID, selector),
			sqlgraph.To(teacher.Table, teacher.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attendancedaysyncs.TeacherTable, attendancedaysyncs.TeacherColumn),
		)
		fromU = sqlgraph.SetNeighbors(adsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AttendanceDaySyncs entity from the query.
// Returns a *NotFoundError when no AttendanceDaySyncs was found.
func (adsq *AttendanceDaySyncsQuery) First(ctx context.Context) (*AttendanceDaySyncs, error) {
	nodes, err := adsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{attendancedaysyncs.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (adsq *AttendanceDaySyncsQuery) FirstX(ctx context.Context) *AttendanceDaySyncs {
	node, err := adsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AttendanceDaySyncs ID from the query.
// Returns a *NotFoundError when no AttendanceDaySyncs ID was found.
func (adsq *AttendanceDaySyncsQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = adsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{attendancedaysyncs.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (adsq *AttendanceDaySyncsQuery) FirstIDX(ctx context.Context) string {
	id, err := adsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AttendanceDaySyncs entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AttendanceDaySyncs entity is found.
// Returns a *NotFoundError when no AttendanceDaySyncs entities are found.
func (adsq *AttendanceDaySyncsQuery) Only(ctx context.Context) (*AttendanceDaySyncs, error) {
	nodes, err := adsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{attendancedaysyncs.Label}
	default:
		return nil, &NotSingularError{attendancedaysyncs.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (adsq *AttendanceDaySyncsQuery) OnlyX(ctx context.Context) *AttendanceDaySyncs {
	node, err := adsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AttendanceDaySyncs ID in the query.
// Returns a *NotSingularError when more than one AttendanceDaySyncs ID is found.
// Returns a *NotFoundError when no entities are found.
func (adsq *AttendanceDaySyncsQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = adsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{attendancedaysyncs.Label}
	default:
		err = &NotSingularError{attendancedaysyncs.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (adsq *AttendanceDaySyncsQuery) OnlyIDX(ctx context.Context) string {
	id, err := adsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AttendanceDaySyncsSlice.
func (adsq *AttendanceDaySyncsQuery) All(ctx context.Context) ([]*AttendanceDaySyncs, error) {
	if err := adsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return adsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (adsq *AttendanceDaySyncsQuery) AllX(ctx context.Context) []*AttendanceDaySyncs {
	nodes, err := adsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AttendanceDaySyncs IDs.
func (adsq *AttendanceDaySyncsQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := adsq.Select(attendancedaysyncs.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (adsq *AttendanceDaySyncsQuery) IDsX(ctx context.Context) []string {
	ids, err := adsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (adsq *AttendanceDaySyncsQuery) Count(ctx context.Context) (int, error) {
	if err := adsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return adsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (adsq *AttendanceDaySyncsQuery) CountX(ctx context.Context) int {
	count, err := adsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (adsq *AttendanceDaySyncsQuery) Exist(ctx context.Context) (bool, error) {
	if err := adsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return adsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (adsq *AttendanceDaySyncsQuery) ExistX(ctx context.Context) bool {
	exist, err := adsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AttendanceDaySyncsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (adsq *AttendanceDaySyncsQuery) Clone() *AttendanceDaySyncsQuery {
	if adsq == nil {
		return nil
	}
	return &AttendanceDaySyncsQuery{
		config:      adsq.config,
		limit:       adsq.limit,
		offset:      adsq.offset,
		order:       append([]OrderFunc{}, adsq.order...),
		predicates:  append([]predicate.AttendanceDaySyncs{}, adsq.predicates...),
		withTeacher: adsq.withTeacher.Clone(),
		// clone intermediate query.
		sql:    adsq.sql.Clone(),
		path:   adsq.path,
		unique: adsq.unique,
	}
}

// WithTeacher tells the query-builder to eager-load the nodes that are connected to
// the "teacher" edge. The optional arguments are used to configure the query builder of the edge.
func (adsq *AttendanceDaySyncsQuery) WithTeacher(opts ...func(*TeacherQuery)) *AttendanceDaySyncsQuery {
	query := &TeacherQuery{config: adsq.config}
	for _, opt := range opts {
		opt(query)
	}
	adsq.withTeacher = query
	return adsq
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
//	client.AttendanceDaySyncs.Query().
//		GroupBy(attendancedaysyncs.FieldLastSyncID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (adsq *AttendanceDaySyncsQuery) GroupBy(field string, fields ...string) *AttendanceDaySyncsGroupBy {
	grbuild := &AttendanceDaySyncsGroupBy{config: adsq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := adsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return adsq.sqlQuery(ctx), nil
	}
	grbuild.label = attendancedaysyncs.Label
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
//	client.AttendanceDaySyncs.Query().
//		Select(attendancedaysyncs.FieldLastSyncID).
//		Scan(ctx, &v)
func (adsq *AttendanceDaySyncsQuery) Select(fields ...string) *AttendanceDaySyncsSelect {
	adsq.fields = append(adsq.fields, fields...)
	selbuild := &AttendanceDaySyncsSelect{AttendanceDaySyncsQuery: adsq}
	selbuild.label = attendancedaysyncs.Label
	selbuild.flds, selbuild.scan = &adsq.fields, selbuild.Scan
	return selbuild
}

func (adsq *AttendanceDaySyncsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range adsq.fields {
		if !attendancedaysyncs.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if adsq.path != nil {
		prev, err := adsq.path(ctx)
		if err != nil {
			return err
		}
		adsq.sql = prev
	}
	return nil
}

func (adsq *AttendanceDaySyncsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AttendanceDaySyncs, error) {
	var (
		nodes       = []*AttendanceDaySyncs{}
		withFKs     = adsq.withFKs
		_spec       = adsq.querySpec()
		loadedTypes = [1]bool{
			adsq.withTeacher != nil,
		}
	)
	if adsq.withTeacher != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, attendancedaysyncs.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AttendanceDaySyncs).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AttendanceDaySyncs{config: adsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, adsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := adsq.withTeacher; query != nil {
		if err := adsq.loadTeacher(ctx, query, nodes, nil,
			func(n *AttendanceDaySyncs, e *Teacher) { n.Edges.Teacher = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (adsq *AttendanceDaySyncsQuery) loadTeacher(ctx context.Context, query *TeacherQuery, nodes []*AttendanceDaySyncs, init func(*AttendanceDaySyncs), assign func(*AttendanceDaySyncs, *Teacher)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*AttendanceDaySyncs)
	for i := range nodes {
		if nodes[i].teacher_attendance_day_syncs == nil {
			continue
		}
		fk := *nodes[i].teacher_attendance_day_syncs
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
			return fmt.Errorf(`unexpected foreign-key "teacher_attendance_day_syncs" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (adsq *AttendanceDaySyncsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := adsq.querySpec()
	_spec.Node.Columns = adsq.fields
	if len(adsq.fields) > 0 {
		_spec.Unique = adsq.unique != nil && *adsq.unique
	}
	return sqlgraph.CountNodes(ctx, adsq.driver, _spec)
}

func (adsq *AttendanceDaySyncsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := adsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (adsq *AttendanceDaySyncsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   attendancedaysyncs.Table,
			Columns: attendancedaysyncs.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: attendancedaysyncs.FieldID,
			},
		},
		From:   adsq.sql,
		Unique: true,
	}
	if unique := adsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := adsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attendancedaysyncs.FieldID)
		for i := range fields {
			if fields[i] != attendancedaysyncs.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := adsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := adsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := adsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := adsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (adsq *AttendanceDaySyncsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(adsq.driver.Dialect())
	t1 := builder.Table(attendancedaysyncs.Table)
	columns := adsq.fields
	if len(columns) == 0 {
		columns = attendancedaysyncs.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if adsq.sql != nil {
		selector = adsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if adsq.unique != nil && *adsq.unique {
		selector.Distinct()
	}
	for _, p := range adsq.predicates {
		p(selector)
	}
	for _, p := range adsq.order {
		p(selector)
	}
	if offset := adsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := adsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AttendanceDaySyncsGroupBy is the group-by builder for AttendanceDaySyncs entities.
type AttendanceDaySyncsGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (adsgb *AttendanceDaySyncsGroupBy) Aggregate(fns ...AggregateFunc) *AttendanceDaySyncsGroupBy {
	adsgb.fns = append(adsgb.fns, fns...)
	return adsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (adsgb *AttendanceDaySyncsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := adsgb.path(ctx)
	if err != nil {
		return err
	}
	adsgb.sql = query
	return adsgb.sqlScan(ctx, v)
}

func (adsgb *AttendanceDaySyncsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range adsgb.fields {
		if !attendancedaysyncs.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := adsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := adsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (adsgb *AttendanceDaySyncsGroupBy) sqlQuery() *sql.Selector {
	selector := adsgb.sql.Select()
	aggregation := make([]string, 0, len(adsgb.fns))
	for _, fn := range adsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(adsgb.fields)+len(adsgb.fns))
		for _, f := range adsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(adsgb.fields...)...)
}

// AttendanceDaySyncsSelect is the builder for selecting fields of AttendanceDaySyncs entities.
type AttendanceDaySyncsSelect struct {
	*AttendanceDaySyncsQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (adss *AttendanceDaySyncsSelect) Scan(ctx context.Context, v interface{}) error {
	if err := adss.prepareQuery(ctx); err != nil {
		return err
	}
	adss.sql = adss.AttendanceDaySyncsQuery.sqlQuery(ctx)
	return adss.sqlScan(ctx, v)
}

func (adss *AttendanceDaySyncsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := adss.sql.Query()
	if err := adss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
