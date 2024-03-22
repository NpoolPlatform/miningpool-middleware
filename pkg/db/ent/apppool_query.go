// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/apppool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/predicate"
)

// AppPoolQuery is the builder for querying AppPool entities.
type AppPoolQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppPool
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppPoolQuery builder.
func (apq *AppPoolQuery) Where(ps ...predicate.AppPool) *AppPoolQuery {
	apq.predicates = append(apq.predicates, ps...)
	return apq
}

// Limit adds a limit step to the query.
func (apq *AppPoolQuery) Limit(limit int) *AppPoolQuery {
	apq.limit = &limit
	return apq
}

// Offset adds an offset step to the query.
func (apq *AppPoolQuery) Offset(offset int) *AppPoolQuery {
	apq.offset = &offset
	return apq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (apq *AppPoolQuery) Unique(unique bool) *AppPoolQuery {
	apq.unique = &unique
	return apq
}

// Order adds an order step to the query.
func (apq *AppPoolQuery) Order(o ...OrderFunc) *AppPoolQuery {
	apq.order = append(apq.order, o...)
	return apq
}

// First returns the first AppPool entity from the query.
// Returns a *NotFoundError when no AppPool was found.
func (apq *AppPoolQuery) First(ctx context.Context) (*AppPool, error) {
	nodes, err := apq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{apppool.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (apq *AppPoolQuery) FirstX(ctx context.Context) *AppPool {
	node, err := apq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppPool ID from the query.
// Returns a *NotFoundError when no AppPool ID was found.
func (apq *AppPoolQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = apq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{apppool.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (apq *AppPoolQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := apq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppPool entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppPool entity is found.
// Returns a *NotFoundError when no AppPool entities are found.
func (apq *AppPoolQuery) Only(ctx context.Context) (*AppPool, error) {
	nodes, err := apq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{apppool.Label}
	default:
		return nil, &NotSingularError{apppool.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (apq *AppPoolQuery) OnlyX(ctx context.Context) *AppPool {
	node, err := apq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppPool ID in the query.
// Returns a *NotSingularError when more than one AppPool ID is found.
// Returns a *NotFoundError when no entities are found.
func (apq *AppPoolQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = apq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{apppool.Label}
	default:
		err = &NotSingularError{apppool.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (apq *AppPoolQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := apq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppPools.
func (apq *AppPoolQuery) All(ctx context.Context) ([]*AppPool, error) {
	if err := apq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return apq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (apq *AppPoolQuery) AllX(ctx context.Context) []*AppPool {
	nodes, err := apq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppPool IDs.
func (apq *AppPoolQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := apq.Select(apppool.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (apq *AppPoolQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := apq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (apq *AppPoolQuery) Count(ctx context.Context) (int, error) {
	if err := apq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return apq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (apq *AppPoolQuery) CountX(ctx context.Context) int {
	count, err := apq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (apq *AppPoolQuery) Exist(ctx context.Context) (bool, error) {
	if err := apq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return apq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (apq *AppPoolQuery) ExistX(ctx context.Context) bool {
	exist, err := apq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppPoolQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (apq *AppPoolQuery) Clone() *AppPoolQuery {
	if apq == nil {
		return nil
	}
	return &AppPoolQuery{
		config:     apq.config,
		limit:      apq.limit,
		offset:     apq.offset,
		order:      append([]OrderFunc{}, apq.order...),
		predicates: append([]predicate.AppPool{}, apq.predicates...),
		// clone intermediate query.
		sql:    apq.sql.Clone(),
		path:   apq.path,
		unique: apq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppPool.Query().
//		GroupBy(apppool.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (apq *AppPoolQuery) GroupBy(field string, fields ...string) *AppPoolGroupBy {
	grbuild := &AppPoolGroupBy{config: apq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := apq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return apq.sqlQuery(ctx), nil
	}
	grbuild.label = apppool.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.AppPool.Query().
//		Select(apppool.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (apq *AppPoolQuery) Select(fields ...string) *AppPoolSelect {
	apq.fields = append(apq.fields, fields...)
	selbuild := &AppPoolSelect{AppPoolQuery: apq}
	selbuild.label = apppool.Label
	selbuild.flds, selbuild.scan = &apq.fields, selbuild.Scan
	return selbuild
}

func (apq *AppPoolQuery) prepareQuery(ctx context.Context) error {
	for _, f := range apq.fields {
		if !apppool.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if apq.path != nil {
		prev, err := apq.path(ctx)
		if err != nil {
			return err
		}
		apq.sql = prev
	}
	if apppool.Policy == nil {
		return errors.New("ent: uninitialized apppool.Policy (forgotten import ent/runtime?)")
	}
	if err := apppool.Policy.EvalQuery(ctx, apq); err != nil {
		return err
	}
	return nil
}

func (apq *AppPoolQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppPool, error) {
	var (
		nodes = []*AppPool{}
		_spec = apq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AppPool).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AppPool{config: apq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(apq.modifiers) > 0 {
		_spec.Modifiers = apq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, apq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (apq *AppPoolQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := apq.querySpec()
	if len(apq.modifiers) > 0 {
		_spec.Modifiers = apq.modifiers
	}
	_spec.Node.Columns = apq.fields
	if len(apq.fields) > 0 {
		_spec.Unique = apq.unique != nil && *apq.unique
	}
	return sqlgraph.CountNodes(ctx, apq.driver, _spec)
}

func (apq *AppPoolQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := apq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (apq *AppPoolQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   apppool.Table,
			Columns: apppool.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: apppool.FieldID,
			},
		},
		From:   apq.sql,
		Unique: true,
	}
	if unique := apq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := apq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apppool.FieldID)
		for i := range fields {
			if fields[i] != apppool.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := apq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := apq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := apq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := apq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (apq *AppPoolQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(apq.driver.Dialect())
	t1 := builder.Table(apppool.Table)
	columns := apq.fields
	if len(columns) == 0 {
		columns = apppool.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if apq.sql != nil {
		selector = apq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if apq.unique != nil && *apq.unique {
		selector.Distinct()
	}
	for _, m := range apq.modifiers {
		m(selector)
	}
	for _, p := range apq.predicates {
		p(selector)
	}
	for _, p := range apq.order {
		p(selector)
	}
	if offset := apq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := apq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (apq *AppPoolQuery) ForUpdate(opts ...sql.LockOption) *AppPoolQuery {
	if apq.driver.Dialect() == dialect.Postgres {
		apq.Unique(false)
	}
	apq.modifiers = append(apq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return apq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (apq *AppPoolQuery) ForShare(opts ...sql.LockOption) *AppPoolQuery {
	if apq.driver.Dialect() == dialect.Postgres {
		apq.Unique(false)
	}
	apq.modifiers = append(apq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return apq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (apq *AppPoolQuery) Modify(modifiers ...func(s *sql.Selector)) *AppPoolSelect {
	apq.modifiers = append(apq.modifiers, modifiers...)
	return apq.Select()
}

// AppPoolGroupBy is the group-by builder for AppPool entities.
type AppPoolGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (apgb *AppPoolGroupBy) Aggregate(fns ...AggregateFunc) *AppPoolGroupBy {
	apgb.fns = append(apgb.fns, fns...)
	return apgb
}

// Scan applies the group-by query and scans the result into the given value.
func (apgb *AppPoolGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := apgb.path(ctx)
	if err != nil {
		return err
	}
	apgb.sql = query
	return apgb.sqlScan(ctx, v)
}

func (apgb *AppPoolGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range apgb.fields {
		if !apppool.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := apgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := apgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (apgb *AppPoolGroupBy) sqlQuery() *sql.Selector {
	selector := apgb.sql.Select()
	aggregation := make([]string, 0, len(apgb.fns))
	for _, fn := range apgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(apgb.fields)+len(apgb.fns))
		for _, f := range apgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(apgb.fields...)...)
}

// AppPoolSelect is the builder for selecting fields of AppPool entities.
type AppPoolSelect struct {
	*AppPoolQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (aps *AppPoolSelect) Scan(ctx context.Context, v interface{}) error {
	if err := aps.prepareQuery(ctx); err != nil {
		return err
	}
	aps.sql = aps.AppPoolQuery.sqlQuery(ctx)
	return aps.sqlScan(ctx, v)
}

func (aps *AppPoolSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := aps.sql.Query()
	if err := aps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (aps *AppPoolSelect) Modify(modifiers ...func(s *sql.Selector)) *AppPoolSelect {
	aps.modifiers = append(aps.modifiers, modifiers...)
	return aps
}
