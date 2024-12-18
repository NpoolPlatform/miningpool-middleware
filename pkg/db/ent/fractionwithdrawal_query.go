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
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fractionwithdrawal"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/predicate"
)

// FractionWithdrawalQuery is the builder for querying FractionWithdrawal entities.
type FractionWithdrawalQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FractionWithdrawal
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FractionWithdrawalQuery builder.
func (fwq *FractionWithdrawalQuery) Where(ps ...predicate.FractionWithdrawal) *FractionWithdrawalQuery {
	fwq.predicates = append(fwq.predicates, ps...)
	return fwq
}

// Limit adds a limit step to the query.
func (fwq *FractionWithdrawalQuery) Limit(limit int) *FractionWithdrawalQuery {
	fwq.limit = &limit
	return fwq
}

// Offset adds an offset step to the query.
func (fwq *FractionWithdrawalQuery) Offset(offset int) *FractionWithdrawalQuery {
	fwq.offset = &offset
	return fwq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fwq *FractionWithdrawalQuery) Unique(unique bool) *FractionWithdrawalQuery {
	fwq.unique = &unique
	return fwq
}

// Order adds an order step to the query.
func (fwq *FractionWithdrawalQuery) Order(o ...OrderFunc) *FractionWithdrawalQuery {
	fwq.order = append(fwq.order, o...)
	return fwq
}

// First returns the first FractionWithdrawal entity from the query.
// Returns a *NotFoundError when no FractionWithdrawal was found.
func (fwq *FractionWithdrawalQuery) First(ctx context.Context) (*FractionWithdrawal, error) {
	nodes, err := fwq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{fractionwithdrawal.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fwq *FractionWithdrawalQuery) FirstX(ctx context.Context) *FractionWithdrawal {
	node, err := fwq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FractionWithdrawal ID from the query.
// Returns a *NotFoundError when no FractionWithdrawal ID was found.
func (fwq *FractionWithdrawalQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = fwq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{fractionwithdrawal.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fwq *FractionWithdrawalQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := fwq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FractionWithdrawal entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FractionWithdrawal entity is found.
// Returns a *NotFoundError when no FractionWithdrawal entities are found.
func (fwq *FractionWithdrawalQuery) Only(ctx context.Context) (*FractionWithdrawal, error) {
	nodes, err := fwq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{fractionwithdrawal.Label}
	default:
		return nil, &NotSingularError{fractionwithdrawal.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fwq *FractionWithdrawalQuery) OnlyX(ctx context.Context) *FractionWithdrawal {
	node, err := fwq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FractionWithdrawal ID in the query.
// Returns a *NotSingularError when more than one FractionWithdrawal ID is found.
// Returns a *NotFoundError when no entities are found.
func (fwq *FractionWithdrawalQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = fwq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{fractionwithdrawal.Label}
	default:
		err = &NotSingularError{fractionwithdrawal.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fwq *FractionWithdrawalQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := fwq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FractionWithdrawals.
func (fwq *FractionWithdrawalQuery) All(ctx context.Context) ([]*FractionWithdrawal, error) {
	if err := fwq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fwq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fwq *FractionWithdrawalQuery) AllX(ctx context.Context) []*FractionWithdrawal {
	nodes, err := fwq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FractionWithdrawal IDs.
func (fwq *FractionWithdrawalQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := fwq.Select(fractionwithdrawal.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fwq *FractionWithdrawalQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := fwq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fwq *FractionWithdrawalQuery) Count(ctx context.Context) (int, error) {
	if err := fwq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fwq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fwq *FractionWithdrawalQuery) CountX(ctx context.Context) int {
	count, err := fwq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fwq *FractionWithdrawalQuery) Exist(ctx context.Context) (bool, error) {
	if err := fwq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fwq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fwq *FractionWithdrawalQuery) ExistX(ctx context.Context) bool {
	exist, err := fwq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FractionWithdrawalQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fwq *FractionWithdrawalQuery) Clone() *FractionWithdrawalQuery {
	if fwq == nil {
		return nil
	}
	return &FractionWithdrawalQuery{
		config:     fwq.config,
		limit:      fwq.limit,
		offset:     fwq.offset,
		order:      append([]OrderFunc{}, fwq.order...),
		predicates: append([]predicate.FractionWithdrawal{}, fwq.predicates...),
		// clone intermediate query.
		sql:    fwq.sql.Clone(),
		path:   fwq.path,
		unique: fwq.unique,
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
//	client.FractionWithdrawal.Query().
//		GroupBy(fractionwithdrawal.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fwq *FractionWithdrawalQuery) GroupBy(field string, fields ...string) *FractionWithdrawalGroupBy {
	grbuild := &FractionWithdrawalGroupBy{config: fwq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fwq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fwq.sqlQuery(ctx), nil
	}
	grbuild.label = fractionwithdrawal.Label
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
//	client.FractionWithdrawal.Query().
//		Select(fractionwithdrawal.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (fwq *FractionWithdrawalQuery) Select(fields ...string) *FractionWithdrawalSelect {
	fwq.fields = append(fwq.fields, fields...)
	selbuild := &FractionWithdrawalSelect{FractionWithdrawalQuery: fwq}
	selbuild.label = fractionwithdrawal.Label
	selbuild.flds, selbuild.scan = &fwq.fields, selbuild.Scan
	return selbuild
}

func (fwq *FractionWithdrawalQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fwq.fields {
		if !fractionwithdrawal.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fwq.path != nil {
		prev, err := fwq.path(ctx)
		if err != nil {
			return err
		}
		fwq.sql = prev
	}
	if fractionwithdrawal.Policy == nil {
		return errors.New("ent: uninitialized fractionwithdrawal.Policy (forgotten import ent/runtime?)")
	}
	if err := fractionwithdrawal.Policy.EvalQuery(ctx, fwq); err != nil {
		return err
	}
	return nil
}

func (fwq *FractionWithdrawalQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FractionWithdrawal, error) {
	var (
		nodes = []*FractionWithdrawal{}
		_spec = fwq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FractionWithdrawal).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &FractionWithdrawal{config: fwq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(fwq.modifiers) > 0 {
		_spec.Modifiers = fwq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fwq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (fwq *FractionWithdrawalQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fwq.querySpec()
	if len(fwq.modifiers) > 0 {
		_spec.Modifiers = fwq.modifiers
	}
	_spec.Node.Columns = fwq.fields
	if len(fwq.fields) > 0 {
		_spec.Unique = fwq.unique != nil && *fwq.unique
	}
	return sqlgraph.CountNodes(ctx, fwq.driver, _spec)
}

func (fwq *FractionWithdrawalQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fwq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fwq *FractionWithdrawalQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fractionwithdrawal.Table,
			Columns: fractionwithdrawal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: fractionwithdrawal.FieldID,
			},
		},
		From:   fwq.sql,
		Unique: true,
	}
	if unique := fwq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fwq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fractionwithdrawal.FieldID)
		for i := range fields {
			if fields[i] != fractionwithdrawal.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fwq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fwq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fwq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fwq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fwq *FractionWithdrawalQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fwq.driver.Dialect())
	t1 := builder.Table(fractionwithdrawal.Table)
	columns := fwq.fields
	if len(columns) == 0 {
		columns = fractionwithdrawal.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fwq.sql != nil {
		selector = fwq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fwq.unique != nil && *fwq.unique {
		selector.Distinct()
	}
	for _, m := range fwq.modifiers {
		m(selector)
	}
	for _, p := range fwq.predicates {
		p(selector)
	}
	for _, p := range fwq.order {
		p(selector)
	}
	if offset := fwq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fwq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (fwq *FractionWithdrawalQuery) ForUpdate(opts ...sql.LockOption) *FractionWithdrawalQuery {
	if fwq.driver.Dialect() == dialect.Postgres {
		fwq.Unique(false)
	}
	fwq.modifiers = append(fwq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return fwq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (fwq *FractionWithdrawalQuery) ForShare(opts ...sql.LockOption) *FractionWithdrawalQuery {
	if fwq.driver.Dialect() == dialect.Postgres {
		fwq.Unique(false)
	}
	fwq.modifiers = append(fwq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return fwq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fwq *FractionWithdrawalQuery) Modify(modifiers ...func(s *sql.Selector)) *FractionWithdrawalSelect {
	fwq.modifiers = append(fwq.modifiers, modifiers...)
	return fwq.Select()
}

// FractionWithdrawalGroupBy is the group-by builder for FractionWithdrawal entities.
type FractionWithdrawalGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fwgb *FractionWithdrawalGroupBy) Aggregate(fns ...AggregateFunc) *FractionWithdrawalGroupBy {
	fwgb.fns = append(fwgb.fns, fns...)
	return fwgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fwgb *FractionWithdrawalGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fwgb.path(ctx)
	if err != nil {
		return err
	}
	fwgb.sql = query
	return fwgb.sqlScan(ctx, v)
}

func (fwgb *FractionWithdrawalGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fwgb.fields {
		if !fractionwithdrawal.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fwgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fwgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fwgb *FractionWithdrawalGroupBy) sqlQuery() *sql.Selector {
	selector := fwgb.sql.Select()
	aggregation := make([]string, 0, len(fwgb.fns))
	for _, fn := range fwgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fwgb.fields)+len(fwgb.fns))
		for _, f := range fwgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fwgb.fields...)...)
}

// FractionWithdrawalSelect is the builder for selecting fields of FractionWithdrawal entities.
type FractionWithdrawalSelect struct {
	*FractionWithdrawalQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fws *FractionWithdrawalSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fws.prepareQuery(ctx); err != nil {
		return err
	}
	fws.sql = fws.FractionWithdrawalQuery.sqlQuery(ctx)
	return fws.sqlScan(ctx, v)
}

func (fws *FractionWithdrawalSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fws.sql.Query()
	if err := fws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fws *FractionWithdrawalSelect) Modify(modifiers ...func(s *sql.Selector)) *FractionWithdrawalSelect {
	fws.modifiers = append(fws.modifiers, modifiers...)
	return fws
}
