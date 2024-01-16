// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CoinUpdate is the builder for updating Coin entities.
type CoinUpdate struct {
	config
	hooks     []Hook
	mutation  *CoinMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CoinUpdate builder.
func (cu *CoinUpdate) Where(ps ...predicate.Coin) *CoinUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CoinUpdate) SetCreatedAt(u uint32) *CoinUpdate {
	cu.mutation.ResetCreatedAt()
	cu.mutation.SetCreatedAt(u)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CoinUpdate) SetNillableCreatedAt(u *uint32) *CoinUpdate {
	if u != nil {
		cu.SetCreatedAt(*u)
	}
	return cu
}

// AddCreatedAt adds u to the "created_at" field.
func (cu *CoinUpdate) AddCreatedAt(u int32) *CoinUpdate {
	cu.mutation.AddCreatedAt(u)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CoinUpdate) SetUpdatedAt(u uint32) *CoinUpdate {
	cu.mutation.ResetUpdatedAt()
	cu.mutation.SetUpdatedAt(u)
	return cu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cu *CoinUpdate) AddUpdatedAt(u int32) *CoinUpdate {
	cu.mutation.AddUpdatedAt(u)
	return cu
}

// SetDeletedAt sets the "deleted_at" field.
func (cu *CoinUpdate) SetDeletedAt(u uint32) *CoinUpdate {
	cu.mutation.ResetDeletedAt()
	cu.mutation.SetDeletedAt(u)
	return cu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cu *CoinUpdate) SetNillableDeletedAt(u *uint32) *CoinUpdate {
	if u != nil {
		cu.SetDeletedAt(*u)
	}
	return cu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cu *CoinUpdate) AddDeletedAt(u int32) *CoinUpdate {
	cu.mutation.AddDeletedAt(u)
	return cu
}

// SetEntID sets the "ent_id" field.
func (cu *CoinUpdate) SetEntID(u uuid.UUID) *CoinUpdate {
	cu.mutation.SetEntID(u)
	return cu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cu *CoinUpdate) SetNillableEntID(u *uuid.UUID) *CoinUpdate {
	if u != nil {
		cu.SetEntID(*u)
	}
	return cu
}

// SetMiningpoolType sets the "miningpool_type" field.
func (cu *CoinUpdate) SetMiningpoolType(s string) *CoinUpdate {
	cu.mutation.SetMiningpoolType(s)
	return cu
}

// SetSite sets the "site" field.
func (cu *CoinUpdate) SetSite(s string) *CoinUpdate {
	cu.mutation.SetSite(s)
	return cu
}

// SetNillableSite sets the "site" field if the given value is not nil.
func (cu *CoinUpdate) SetNillableSite(s *string) *CoinUpdate {
	if s != nil {
		cu.SetSite(*s)
	}
	return cu
}

// ClearSite clears the value of the "site" field.
func (cu *CoinUpdate) ClearSite() *CoinUpdate {
	cu.mutation.ClearSite()
	return cu
}

// SetCoinType sets the "coin_type" field.
func (cu *CoinUpdate) SetCoinType(s string) *CoinUpdate {
	cu.mutation.SetCoinType(s)
	return cu
}

// SetRevenueType sets the "revenue_type" field.
func (cu *CoinUpdate) SetRevenueType(s string) *CoinUpdate {
	cu.mutation.SetRevenueType(s)
	return cu
}

// SetFeeRate sets the "fee_rate" field.
func (cu *CoinUpdate) SetFeeRate(f float32) *CoinUpdate {
	cu.mutation.ResetFeeRate()
	cu.mutation.SetFeeRate(f)
	return cu
}

// SetNillableFeeRate sets the "fee_rate" field if the given value is not nil.
func (cu *CoinUpdate) SetNillableFeeRate(f *float32) *CoinUpdate {
	if f != nil {
		cu.SetFeeRate(*f)
	}
	return cu
}

// AddFeeRate adds f to the "fee_rate" field.
func (cu *CoinUpdate) AddFeeRate(f float32) *CoinUpdate {
	cu.mutation.AddFeeRate(f)
	return cu
}

// ClearFeeRate clears the value of the "fee_rate" field.
func (cu *CoinUpdate) ClearFeeRate() *CoinUpdate {
	cu.mutation.ClearFeeRate()
	return cu
}

// SetFixedRevenueAble sets the "fixed_revenue_able" field.
func (cu *CoinUpdate) SetFixedRevenueAble(b bool) *CoinUpdate {
	cu.mutation.SetFixedRevenueAble(b)
	return cu
}

// SetRemark sets the "remark" field.
func (cu *CoinUpdate) SetRemark(s string) *CoinUpdate {
	cu.mutation.SetRemark(s)
	return cu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (cu *CoinUpdate) SetNillableRemark(s *string) *CoinUpdate {
	if s != nil {
		cu.SetRemark(*s)
	}
	return cu
}

// ClearRemark clears the value of the "remark" field.
func (cu *CoinUpdate) ClearRemark() *CoinUpdate {
	cu.mutation.ClearRemark()
	return cu
}

// Mutation returns the CoinMutation object of the builder.
func (cu *CoinUpdate) Mutation() *CoinMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CoinUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := cu.defaults(); err != nil {
		return 0, err
	}
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CoinUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CoinUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CoinUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CoinUpdate) defaults() error {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		if coin.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coin.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coin.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CoinUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CoinUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CoinUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coin.Table,
			Columns: coin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: coin.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coin.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coin.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coin.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coin.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coin.FieldDeletedAt,
		})
	}
	if value, ok := cu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coin.FieldDeletedAt,
		})
	}
	if value, ok := cu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coin.FieldEntID,
		})
	}
	if value, ok := cu.mutation.MiningpoolType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coin.FieldMiningpoolType,
		})
	}
	if value, ok := cu.mutation.Site(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coin.FieldSite,
		})
	}
	if cu.mutation.SiteCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coin.FieldSite,
		})
	}
	if value, ok := cu.mutation.CoinType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coin.FieldCoinType,
		})
	}
	if value, ok := cu.mutation.RevenueType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coin.FieldRevenueType,
		})
	}
	if value, ok := cu.mutation.FeeRate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: coin.FieldFeeRate,
		})
	}
	if value, ok := cu.mutation.AddedFeeRate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: coin.FieldFeeRate,
		})
	}
	if cu.mutation.FeeRateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Column: coin.FieldFeeRate,
		})
	}
	if value, ok := cu.mutation.FixedRevenueAble(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coin.FieldFixedRevenueAble,
		})
	}
	if value, ok := cu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coin.FieldRemark,
		})
	}
	if cu.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coin.FieldRemark,
		})
	}
	_spec.Modifiers = cu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CoinUpdateOne is the builder for updating a single Coin entity.
type CoinUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CoinMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CoinUpdateOne) SetCreatedAt(u uint32) *CoinUpdateOne {
	cuo.mutation.ResetCreatedAt()
	cuo.mutation.SetCreatedAt(u)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CoinUpdateOne) SetNillableCreatedAt(u *uint32) *CoinUpdateOne {
	if u != nil {
		cuo.SetCreatedAt(*u)
	}
	return cuo
}

// AddCreatedAt adds u to the "created_at" field.
func (cuo *CoinUpdateOne) AddCreatedAt(u int32) *CoinUpdateOne {
	cuo.mutation.AddCreatedAt(u)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CoinUpdateOne) SetUpdatedAt(u uint32) *CoinUpdateOne {
	cuo.mutation.ResetUpdatedAt()
	cuo.mutation.SetUpdatedAt(u)
	return cuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cuo *CoinUpdateOne) AddUpdatedAt(u int32) *CoinUpdateOne {
	cuo.mutation.AddUpdatedAt(u)
	return cuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cuo *CoinUpdateOne) SetDeletedAt(u uint32) *CoinUpdateOne {
	cuo.mutation.ResetDeletedAt()
	cuo.mutation.SetDeletedAt(u)
	return cuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cuo *CoinUpdateOne) SetNillableDeletedAt(u *uint32) *CoinUpdateOne {
	if u != nil {
		cuo.SetDeletedAt(*u)
	}
	return cuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cuo *CoinUpdateOne) AddDeletedAt(u int32) *CoinUpdateOne {
	cuo.mutation.AddDeletedAt(u)
	return cuo
}

// SetEntID sets the "ent_id" field.
func (cuo *CoinUpdateOne) SetEntID(u uuid.UUID) *CoinUpdateOne {
	cuo.mutation.SetEntID(u)
	return cuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cuo *CoinUpdateOne) SetNillableEntID(u *uuid.UUID) *CoinUpdateOne {
	if u != nil {
		cuo.SetEntID(*u)
	}
	return cuo
}

// SetMiningpoolType sets the "miningpool_type" field.
func (cuo *CoinUpdateOne) SetMiningpoolType(s string) *CoinUpdateOne {
	cuo.mutation.SetMiningpoolType(s)
	return cuo
}

// SetSite sets the "site" field.
func (cuo *CoinUpdateOne) SetSite(s string) *CoinUpdateOne {
	cuo.mutation.SetSite(s)
	return cuo
}

// SetNillableSite sets the "site" field if the given value is not nil.
func (cuo *CoinUpdateOne) SetNillableSite(s *string) *CoinUpdateOne {
	if s != nil {
		cuo.SetSite(*s)
	}
	return cuo
}

// ClearSite clears the value of the "site" field.
func (cuo *CoinUpdateOne) ClearSite() *CoinUpdateOne {
	cuo.mutation.ClearSite()
	return cuo
}

// SetCoinType sets the "coin_type" field.
func (cuo *CoinUpdateOne) SetCoinType(s string) *CoinUpdateOne {
	cuo.mutation.SetCoinType(s)
	return cuo
}

// SetRevenueType sets the "revenue_type" field.
func (cuo *CoinUpdateOne) SetRevenueType(s string) *CoinUpdateOne {
	cuo.mutation.SetRevenueType(s)
	return cuo
}

// SetFeeRate sets the "fee_rate" field.
func (cuo *CoinUpdateOne) SetFeeRate(f float32) *CoinUpdateOne {
	cuo.mutation.ResetFeeRate()
	cuo.mutation.SetFeeRate(f)
	return cuo
}

// SetNillableFeeRate sets the "fee_rate" field if the given value is not nil.
func (cuo *CoinUpdateOne) SetNillableFeeRate(f *float32) *CoinUpdateOne {
	if f != nil {
		cuo.SetFeeRate(*f)
	}
	return cuo
}

// AddFeeRate adds f to the "fee_rate" field.
func (cuo *CoinUpdateOne) AddFeeRate(f float32) *CoinUpdateOne {
	cuo.mutation.AddFeeRate(f)
	return cuo
}

// ClearFeeRate clears the value of the "fee_rate" field.
func (cuo *CoinUpdateOne) ClearFeeRate() *CoinUpdateOne {
	cuo.mutation.ClearFeeRate()
	return cuo
}

// SetFixedRevenueAble sets the "fixed_revenue_able" field.
func (cuo *CoinUpdateOne) SetFixedRevenueAble(b bool) *CoinUpdateOne {
	cuo.mutation.SetFixedRevenueAble(b)
	return cuo
}

// SetRemark sets the "remark" field.
func (cuo *CoinUpdateOne) SetRemark(s string) *CoinUpdateOne {
	cuo.mutation.SetRemark(s)
	return cuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (cuo *CoinUpdateOne) SetNillableRemark(s *string) *CoinUpdateOne {
	if s != nil {
		cuo.SetRemark(*s)
	}
	return cuo
}

// ClearRemark clears the value of the "remark" field.
func (cuo *CoinUpdateOne) ClearRemark() *CoinUpdateOne {
	cuo.mutation.ClearRemark()
	return cuo
}

// Mutation returns the CoinMutation object of the builder.
func (cuo *CoinUpdateOne) Mutation() *CoinMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CoinUpdateOne) Select(field string, fields ...string) *CoinUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Coin entity.
func (cuo *CoinUpdateOne) Save(ctx context.Context) (*Coin, error) {
	var (
		err  error
		node *Coin
	)
	if err := cuo.defaults(); err != nil {
		return nil, err
	}
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Coin)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CoinMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CoinUpdateOne) SaveX(ctx context.Context) *Coin {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CoinUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CoinUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CoinUpdateOne) defaults() error {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		if coin.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coin.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coin.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CoinUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CoinUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CoinUpdateOne) sqlSave(ctx context.Context) (_node *Coin, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coin.Table,
			Columns: coin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: coin.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Coin.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coin.FieldID)
		for _, f := range fields {
			if !coin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coin.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coin.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coin.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coin.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coin.FieldDeletedAt,
		})
	}
	if value, ok := cuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coin.FieldDeletedAt,
		})
	}
	if value, ok := cuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coin.FieldEntID,
		})
	}
	if value, ok := cuo.mutation.MiningpoolType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coin.FieldMiningpoolType,
		})
	}
	if value, ok := cuo.mutation.Site(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coin.FieldSite,
		})
	}
	if cuo.mutation.SiteCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coin.FieldSite,
		})
	}
	if value, ok := cuo.mutation.CoinType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coin.FieldCoinType,
		})
	}
	if value, ok := cuo.mutation.RevenueType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coin.FieldRevenueType,
		})
	}
	if value, ok := cuo.mutation.FeeRate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: coin.FieldFeeRate,
		})
	}
	if value, ok := cuo.mutation.AddedFeeRate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: coin.FieldFeeRate,
		})
	}
	if cuo.mutation.FeeRateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Column: coin.FieldFeeRate,
		})
	}
	if value, ok := cuo.mutation.FixedRevenueAble(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coin.FieldFixedRevenueAble,
		})
	}
	if value, ok := cuo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coin.FieldRemark,
		})
	}
	if cuo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coin.FieldRemark,
		})
	}
	_spec.Modifiers = cuo.modifiers
	_node = &Coin{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
