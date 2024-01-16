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
	"github.com/google/uuid"
)

// CoinCreate is the builder for creating a Coin entity.
type CoinCreate struct {
	config
	mutation *CoinMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cc *CoinCreate) SetCreatedAt(u uint32) *CoinCreate {
	cc.mutation.SetCreatedAt(u)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CoinCreate) SetNillableCreatedAt(u *uint32) *CoinCreate {
	if u != nil {
		cc.SetCreatedAt(*u)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CoinCreate) SetUpdatedAt(u uint32) *CoinCreate {
	cc.mutation.SetUpdatedAt(u)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CoinCreate) SetNillableUpdatedAt(u *uint32) *CoinCreate {
	if u != nil {
		cc.SetUpdatedAt(*u)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CoinCreate) SetDeletedAt(u uint32) *CoinCreate {
	cc.mutation.SetDeletedAt(u)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CoinCreate) SetNillableDeletedAt(u *uint32) *CoinCreate {
	if u != nil {
		cc.SetDeletedAt(*u)
	}
	return cc
}

// SetEntID sets the "ent_id" field.
func (cc *CoinCreate) SetEntID(u uuid.UUID) *CoinCreate {
	cc.mutation.SetEntID(u)
	return cc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cc *CoinCreate) SetNillableEntID(u *uuid.UUID) *CoinCreate {
	if u != nil {
		cc.SetEntID(*u)
	}
	return cc
}

// SetMiningpoolType sets the "miningpool_type" field.
func (cc *CoinCreate) SetMiningpoolType(s string) *CoinCreate {
	cc.mutation.SetMiningpoolType(s)
	return cc
}

// SetSite sets the "site" field.
func (cc *CoinCreate) SetSite(s string) *CoinCreate {
	cc.mutation.SetSite(s)
	return cc
}

// SetNillableSite sets the "site" field if the given value is not nil.
func (cc *CoinCreate) SetNillableSite(s *string) *CoinCreate {
	if s != nil {
		cc.SetSite(*s)
	}
	return cc
}

// SetCoinType sets the "coin_type" field.
func (cc *CoinCreate) SetCoinType(s string) *CoinCreate {
	cc.mutation.SetCoinType(s)
	return cc
}

// SetRevenueType sets the "revenue_type" field.
func (cc *CoinCreate) SetRevenueType(s string) *CoinCreate {
	cc.mutation.SetRevenueType(s)
	return cc
}

// SetFeeRate sets the "fee_rate" field.
func (cc *CoinCreate) SetFeeRate(f float32) *CoinCreate {
	cc.mutation.SetFeeRate(f)
	return cc
}

// SetNillableFeeRate sets the "fee_rate" field if the given value is not nil.
func (cc *CoinCreate) SetNillableFeeRate(f *float32) *CoinCreate {
	if f != nil {
		cc.SetFeeRate(*f)
	}
	return cc
}

// SetFixedRevenueAble sets the "fixed_revenue_able" field.
func (cc *CoinCreate) SetFixedRevenueAble(b bool) *CoinCreate {
	cc.mutation.SetFixedRevenueAble(b)
	return cc
}

// SetRemark sets the "remark" field.
func (cc *CoinCreate) SetRemark(s string) *CoinCreate {
	cc.mutation.SetRemark(s)
	return cc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (cc *CoinCreate) SetNillableRemark(s *string) *CoinCreate {
	if s != nil {
		cc.SetRemark(*s)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CoinCreate) SetID(u uint32) *CoinCreate {
	cc.mutation.SetID(u)
	return cc
}

// Mutation returns the CoinMutation object of the builder.
func (cc *CoinCreate) Mutation() *CoinMutation {
	return cc.mutation
}

// Save creates the Coin in the database.
func (cc *CoinCreate) Save(ctx context.Context) (*Coin, error) {
	var (
		err  error
		node *Coin
	)
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (cc *CoinCreate) SaveX(ctx context.Context) *Coin {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CoinCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CoinCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CoinCreate) defaults() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		if coin.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized coin.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := coin.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		if coin.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coin.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coin.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		if coin.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized coin.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := coin.DefaultDeletedAt()
		cc.mutation.SetDeletedAt(v)
	}
	if _, ok := cc.mutation.EntID(); !ok {
		if coin.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized coin.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := coin.DefaultEntID()
		cc.mutation.SetEntID(v)
	}
	if _, ok := cc.mutation.Site(); !ok {
		v := coin.DefaultSite
		cc.mutation.SetSite(v)
	}
	if _, ok := cc.mutation.FeeRate(); !ok {
		v := coin.DefaultFeeRate
		cc.mutation.SetFeeRate(v)
	}
	if _, ok := cc.mutation.Remark(); !ok {
		v := coin.DefaultRemark
		cc.mutation.SetRemark(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *CoinCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Coin.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Coin.updated_at"`)}
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Coin.deleted_at"`)}
	}
	if _, ok := cc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "Coin.ent_id"`)}
	}
	if _, ok := cc.mutation.MiningpoolType(); !ok {
		return &ValidationError{Name: "miningpool_type", err: errors.New(`ent: missing required field "Coin.miningpool_type"`)}
	}
	if _, ok := cc.mutation.CoinType(); !ok {
		return &ValidationError{Name: "coin_type", err: errors.New(`ent: missing required field "Coin.coin_type"`)}
	}
	if _, ok := cc.mutation.RevenueType(); !ok {
		return &ValidationError{Name: "revenue_type", err: errors.New(`ent: missing required field "Coin.revenue_type"`)}
	}
	if _, ok := cc.mutation.FixedRevenueAble(); !ok {
		return &ValidationError{Name: "fixed_revenue_able", err: errors.New(`ent: missing required field "Coin.fixed_revenue_able"`)}
	}
	return nil
}

func (cc *CoinCreate) sqlSave(ctx context.Context) (*Coin, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (cc *CoinCreate) createSpec() (*Coin, *sqlgraph.CreateSpec) {
	var (
		_node = &Coin{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: coin.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: coin.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coin.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coin.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coin.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coin.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := cc.mutation.MiningpoolType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coin.FieldMiningpoolType,
		})
		_node.MiningpoolType = value
	}
	if value, ok := cc.mutation.Site(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coin.FieldSite,
		})
		_node.Site = value
	}
	if value, ok := cc.mutation.CoinType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coin.FieldCoinType,
		})
		_node.CoinType = value
	}
	if value, ok := cc.mutation.RevenueType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coin.FieldRevenueType,
		})
		_node.RevenueType = value
	}
	if value, ok := cc.mutation.FeeRate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: coin.FieldFeeRate,
		})
		_node.FeeRate = value
	}
	if value, ok := cc.mutation.FixedRevenueAble(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coin.FieldFixedRevenueAble,
		})
		_node.FixedRevenueAble = value
	}
	if value, ok := cc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coin.FieldRemark,
		})
		_node.Remark = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Coin.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CoinUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (cc *CoinCreate) OnConflict(opts ...sql.ConflictOption) *CoinUpsertOne {
	cc.conflict = opts
	return &CoinUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Coin.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cc *CoinCreate) OnConflictColumns(columns ...string) *CoinUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CoinUpsertOne{
		create: cc,
	}
}

type (
	// CoinUpsertOne is the builder for "upsert"-ing
	//  one Coin node.
	CoinUpsertOne struct {
		create *CoinCreate
	}

	// CoinUpsert is the "OnConflict" setter.
	CoinUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *CoinUpsert) SetCreatedAt(v uint32) *CoinUpsert {
	u.Set(coin.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CoinUpsert) UpdateCreatedAt() *CoinUpsert {
	u.SetExcluded(coin.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CoinUpsert) AddCreatedAt(v uint32) *CoinUpsert {
	u.Add(coin.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CoinUpsert) SetUpdatedAt(v uint32) *CoinUpsert {
	u.Set(coin.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CoinUpsert) UpdateUpdatedAt() *CoinUpsert {
	u.SetExcluded(coin.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CoinUpsert) AddUpdatedAt(v uint32) *CoinUpsert {
	u.Add(coin.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CoinUpsert) SetDeletedAt(v uint32) *CoinUpsert {
	u.Set(coin.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CoinUpsert) UpdateDeletedAt() *CoinUpsert {
	u.SetExcluded(coin.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CoinUpsert) AddDeletedAt(v uint32) *CoinUpsert {
	u.Add(coin.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *CoinUpsert) SetEntID(v uuid.UUID) *CoinUpsert {
	u.Set(coin.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *CoinUpsert) UpdateEntID() *CoinUpsert {
	u.SetExcluded(coin.FieldEntID)
	return u
}

// SetMiningpoolType sets the "miningpool_type" field.
func (u *CoinUpsert) SetMiningpoolType(v string) *CoinUpsert {
	u.Set(coin.FieldMiningpoolType, v)
	return u
}

// UpdateMiningpoolType sets the "miningpool_type" field to the value that was provided on create.
func (u *CoinUpsert) UpdateMiningpoolType() *CoinUpsert {
	u.SetExcluded(coin.FieldMiningpoolType)
	return u
}

// SetSite sets the "site" field.
func (u *CoinUpsert) SetSite(v string) *CoinUpsert {
	u.Set(coin.FieldSite, v)
	return u
}

// UpdateSite sets the "site" field to the value that was provided on create.
func (u *CoinUpsert) UpdateSite() *CoinUpsert {
	u.SetExcluded(coin.FieldSite)
	return u
}

// ClearSite clears the value of the "site" field.
func (u *CoinUpsert) ClearSite() *CoinUpsert {
	u.SetNull(coin.FieldSite)
	return u
}

// SetCoinType sets the "coin_type" field.
func (u *CoinUpsert) SetCoinType(v string) *CoinUpsert {
	u.Set(coin.FieldCoinType, v)
	return u
}

// UpdateCoinType sets the "coin_type" field to the value that was provided on create.
func (u *CoinUpsert) UpdateCoinType() *CoinUpsert {
	u.SetExcluded(coin.FieldCoinType)
	return u
}

// SetRevenueType sets the "revenue_type" field.
func (u *CoinUpsert) SetRevenueType(v string) *CoinUpsert {
	u.Set(coin.FieldRevenueType, v)
	return u
}

// UpdateRevenueType sets the "revenue_type" field to the value that was provided on create.
func (u *CoinUpsert) UpdateRevenueType() *CoinUpsert {
	u.SetExcluded(coin.FieldRevenueType)
	return u
}

// SetFeeRate sets the "fee_rate" field.
func (u *CoinUpsert) SetFeeRate(v float32) *CoinUpsert {
	u.Set(coin.FieldFeeRate, v)
	return u
}

// UpdateFeeRate sets the "fee_rate" field to the value that was provided on create.
func (u *CoinUpsert) UpdateFeeRate() *CoinUpsert {
	u.SetExcluded(coin.FieldFeeRate)
	return u
}

// AddFeeRate adds v to the "fee_rate" field.
func (u *CoinUpsert) AddFeeRate(v float32) *CoinUpsert {
	u.Add(coin.FieldFeeRate, v)
	return u
}

// ClearFeeRate clears the value of the "fee_rate" field.
func (u *CoinUpsert) ClearFeeRate() *CoinUpsert {
	u.SetNull(coin.FieldFeeRate)
	return u
}

// SetFixedRevenueAble sets the "fixed_revenue_able" field.
func (u *CoinUpsert) SetFixedRevenueAble(v bool) *CoinUpsert {
	u.Set(coin.FieldFixedRevenueAble, v)
	return u
}

// UpdateFixedRevenueAble sets the "fixed_revenue_able" field to the value that was provided on create.
func (u *CoinUpsert) UpdateFixedRevenueAble() *CoinUpsert {
	u.SetExcluded(coin.FieldFixedRevenueAble)
	return u
}

// SetRemark sets the "remark" field.
func (u *CoinUpsert) SetRemark(v string) *CoinUpsert {
	u.Set(coin.FieldRemark, v)
	return u
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *CoinUpsert) UpdateRemark() *CoinUpsert {
	u.SetExcluded(coin.FieldRemark)
	return u
}

// ClearRemark clears the value of the "remark" field.
func (u *CoinUpsert) ClearRemark() *CoinUpsert {
	u.SetNull(coin.FieldRemark)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Coin.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(coin.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CoinUpsertOne) UpdateNewValues() *CoinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(coin.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Coin.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CoinUpsertOne) Ignore() *CoinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CoinUpsertOne) DoNothing() *CoinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CoinCreate.OnConflict
// documentation for more info.
func (u *CoinUpsertOne) Update(set func(*CoinUpsert)) *CoinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CoinUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CoinUpsertOne) SetCreatedAt(v uint32) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CoinUpsertOne) AddCreatedAt(v uint32) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CoinUpsertOne) UpdateCreatedAt() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CoinUpsertOne) SetUpdatedAt(v uint32) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CoinUpsertOne) AddUpdatedAt(v uint32) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CoinUpsertOne) UpdateUpdatedAt() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CoinUpsertOne) SetDeletedAt(v uint32) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CoinUpsertOne) AddDeletedAt(v uint32) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CoinUpsertOne) UpdateDeletedAt() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *CoinUpsertOne) SetEntID(v uuid.UUID) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *CoinUpsertOne) UpdateEntID() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateEntID()
	})
}

// SetMiningpoolType sets the "miningpool_type" field.
func (u *CoinUpsertOne) SetMiningpoolType(v string) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.SetMiningpoolType(v)
	})
}

// UpdateMiningpoolType sets the "miningpool_type" field to the value that was provided on create.
func (u *CoinUpsertOne) UpdateMiningpoolType() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateMiningpoolType()
	})
}

// SetSite sets the "site" field.
func (u *CoinUpsertOne) SetSite(v string) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.SetSite(v)
	})
}

// UpdateSite sets the "site" field to the value that was provided on create.
func (u *CoinUpsertOne) UpdateSite() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateSite()
	})
}

// ClearSite clears the value of the "site" field.
func (u *CoinUpsertOne) ClearSite() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.ClearSite()
	})
}

// SetCoinType sets the "coin_type" field.
func (u *CoinUpsertOne) SetCoinType(v string) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.SetCoinType(v)
	})
}

// UpdateCoinType sets the "coin_type" field to the value that was provided on create.
func (u *CoinUpsertOne) UpdateCoinType() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateCoinType()
	})
}

// SetRevenueType sets the "revenue_type" field.
func (u *CoinUpsertOne) SetRevenueType(v string) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.SetRevenueType(v)
	})
}

// UpdateRevenueType sets the "revenue_type" field to the value that was provided on create.
func (u *CoinUpsertOne) UpdateRevenueType() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateRevenueType()
	})
}

// SetFeeRate sets the "fee_rate" field.
func (u *CoinUpsertOne) SetFeeRate(v float32) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.SetFeeRate(v)
	})
}

// AddFeeRate adds v to the "fee_rate" field.
func (u *CoinUpsertOne) AddFeeRate(v float32) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.AddFeeRate(v)
	})
}

// UpdateFeeRate sets the "fee_rate" field to the value that was provided on create.
func (u *CoinUpsertOne) UpdateFeeRate() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateFeeRate()
	})
}

// ClearFeeRate clears the value of the "fee_rate" field.
func (u *CoinUpsertOne) ClearFeeRate() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.ClearFeeRate()
	})
}

// SetFixedRevenueAble sets the "fixed_revenue_able" field.
func (u *CoinUpsertOne) SetFixedRevenueAble(v bool) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.SetFixedRevenueAble(v)
	})
}

// UpdateFixedRevenueAble sets the "fixed_revenue_able" field to the value that was provided on create.
func (u *CoinUpsertOne) UpdateFixedRevenueAble() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateFixedRevenueAble()
	})
}

// SetRemark sets the "remark" field.
func (u *CoinUpsertOne) SetRemark(v string) *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *CoinUpsertOne) UpdateRemark() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *CoinUpsertOne) ClearRemark() *CoinUpsertOne {
	return u.Update(func(s *CoinUpsert) {
		s.ClearRemark()
	})
}

// Exec executes the query.
func (u *CoinUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CoinCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CoinUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CoinUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CoinUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CoinCreateBulk is the builder for creating many Coin entities in bulk.
type CoinCreateBulk struct {
	config
	builders []*CoinCreate
	conflict []sql.ConflictOption
}

// Save creates the Coin entities in the database.
func (ccb *CoinCreateBulk) Save(ctx context.Context) ([]*Coin, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Coin, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CoinMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CoinCreateBulk) SaveX(ctx context.Context) []*Coin {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CoinCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CoinCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Coin.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CoinUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ccb *CoinCreateBulk) OnConflict(opts ...sql.ConflictOption) *CoinUpsertBulk {
	ccb.conflict = opts
	return &CoinUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Coin.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccb *CoinCreateBulk) OnConflictColumns(columns ...string) *CoinUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CoinUpsertBulk{
		create: ccb,
	}
}

// CoinUpsertBulk is the builder for "upsert"-ing
// a bulk of Coin nodes.
type CoinUpsertBulk struct {
	create *CoinCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Coin.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(coin.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CoinUpsertBulk) UpdateNewValues() *CoinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(coin.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Coin.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CoinUpsertBulk) Ignore() *CoinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CoinUpsertBulk) DoNothing() *CoinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CoinCreateBulk.OnConflict
// documentation for more info.
func (u *CoinUpsertBulk) Update(set func(*CoinUpsert)) *CoinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CoinUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CoinUpsertBulk) SetCreatedAt(v uint32) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CoinUpsertBulk) AddCreatedAt(v uint32) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CoinUpsertBulk) UpdateCreatedAt() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CoinUpsertBulk) SetUpdatedAt(v uint32) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CoinUpsertBulk) AddUpdatedAt(v uint32) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CoinUpsertBulk) UpdateUpdatedAt() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CoinUpsertBulk) SetDeletedAt(v uint32) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CoinUpsertBulk) AddDeletedAt(v uint32) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CoinUpsertBulk) UpdateDeletedAt() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *CoinUpsertBulk) SetEntID(v uuid.UUID) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *CoinUpsertBulk) UpdateEntID() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateEntID()
	})
}

// SetMiningpoolType sets the "miningpool_type" field.
func (u *CoinUpsertBulk) SetMiningpoolType(v string) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.SetMiningpoolType(v)
	})
}

// UpdateMiningpoolType sets the "miningpool_type" field to the value that was provided on create.
func (u *CoinUpsertBulk) UpdateMiningpoolType() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateMiningpoolType()
	})
}

// SetSite sets the "site" field.
func (u *CoinUpsertBulk) SetSite(v string) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.SetSite(v)
	})
}

// UpdateSite sets the "site" field to the value that was provided on create.
func (u *CoinUpsertBulk) UpdateSite() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateSite()
	})
}

// ClearSite clears the value of the "site" field.
func (u *CoinUpsertBulk) ClearSite() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.ClearSite()
	})
}

// SetCoinType sets the "coin_type" field.
func (u *CoinUpsertBulk) SetCoinType(v string) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.SetCoinType(v)
	})
}

// UpdateCoinType sets the "coin_type" field to the value that was provided on create.
func (u *CoinUpsertBulk) UpdateCoinType() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateCoinType()
	})
}

// SetRevenueType sets the "revenue_type" field.
func (u *CoinUpsertBulk) SetRevenueType(v string) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.SetRevenueType(v)
	})
}

// UpdateRevenueType sets the "revenue_type" field to the value that was provided on create.
func (u *CoinUpsertBulk) UpdateRevenueType() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateRevenueType()
	})
}

// SetFeeRate sets the "fee_rate" field.
func (u *CoinUpsertBulk) SetFeeRate(v float32) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.SetFeeRate(v)
	})
}

// AddFeeRate adds v to the "fee_rate" field.
func (u *CoinUpsertBulk) AddFeeRate(v float32) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.AddFeeRate(v)
	})
}

// UpdateFeeRate sets the "fee_rate" field to the value that was provided on create.
func (u *CoinUpsertBulk) UpdateFeeRate() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateFeeRate()
	})
}

// ClearFeeRate clears the value of the "fee_rate" field.
func (u *CoinUpsertBulk) ClearFeeRate() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.ClearFeeRate()
	})
}

// SetFixedRevenueAble sets the "fixed_revenue_able" field.
func (u *CoinUpsertBulk) SetFixedRevenueAble(v bool) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.SetFixedRevenueAble(v)
	})
}

// UpdateFixedRevenueAble sets the "fixed_revenue_able" field to the value that was provided on create.
func (u *CoinUpsertBulk) UpdateFixedRevenueAble() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateFixedRevenueAble()
	})
}

// SetRemark sets the "remark" field.
func (u *CoinUpsertBulk) SetRemark(v string) *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *CoinUpsertBulk) UpdateRemark() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *CoinUpsertBulk) ClearRemark() *CoinUpsertBulk {
	return u.Update(func(s *CoinUpsert) {
		s.ClearRemark()
	})
}

// Exec executes the query.
func (u *CoinUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CoinCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CoinCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CoinUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
