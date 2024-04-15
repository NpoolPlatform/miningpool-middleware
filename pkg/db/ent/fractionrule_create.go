// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fractionrule"
	"github.com/google/uuid"
)

// FractionRuleCreate is the builder for creating a FractionRule entity.
type FractionRuleCreate struct {
	config
	mutation *FractionRuleMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (frc *FractionRuleCreate) SetCreatedAt(u uint32) *FractionRuleCreate {
	frc.mutation.SetCreatedAt(u)
	return frc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (frc *FractionRuleCreate) SetNillableCreatedAt(u *uint32) *FractionRuleCreate {
	if u != nil {
		frc.SetCreatedAt(*u)
	}
	return frc
}

// SetUpdatedAt sets the "updated_at" field.
func (frc *FractionRuleCreate) SetUpdatedAt(u uint32) *FractionRuleCreate {
	frc.mutation.SetUpdatedAt(u)
	return frc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (frc *FractionRuleCreate) SetNillableUpdatedAt(u *uint32) *FractionRuleCreate {
	if u != nil {
		frc.SetUpdatedAt(*u)
	}
	return frc
}

// SetDeletedAt sets the "deleted_at" field.
func (frc *FractionRuleCreate) SetDeletedAt(u uint32) *FractionRuleCreate {
	frc.mutation.SetDeletedAt(u)
	return frc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (frc *FractionRuleCreate) SetNillableDeletedAt(u *uint32) *FractionRuleCreate {
	if u != nil {
		frc.SetDeletedAt(*u)
	}
	return frc
}

// SetEntID sets the "ent_id" field.
func (frc *FractionRuleCreate) SetEntID(u uuid.UUID) *FractionRuleCreate {
	frc.mutation.SetEntID(u)
	return frc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (frc *FractionRuleCreate) SetNillableEntID(u *uuid.UUID) *FractionRuleCreate {
	if u != nil {
		frc.SetEntID(*u)
	}
	return frc
}

// SetMiningpoolType sets the "miningpool_type" field.
func (frc *FractionRuleCreate) SetMiningpoolType(s string) *FractionRuleCreate {
	frc.mutation.SetMiningpoolType(s)
	return frc
}

// SetCoinType sets the "coin_type" field.
func (frc *FractionRuleCreate) SetCoinType(s string) *FractionRuleCreate {
	frc.mutation.SetCoinType(s)
	return frc
}

// SetWithdrawInterval sets the "withdraw_interval" field.
func (frc *FractionRuleCreate) SetWithdrawInterval(u uint32) *FractionRuleCreate {
	frc.mutation.SetWithdrawInterval(u)
	return frc
}

// SetMinAmount sets the "min_amount" field.
func (frc *FractionRuleCreate) SetMinAmount(f float32) *FractionRuleCreate {
	frc.mutation.SetMinAmount(f)
	return frc
}

// SetWithdrawRate sets the "withdraw_rate" field.
func (frc *FractionRuleCreate) SetWithdrawRate(f float32) *FractionRuleCreate {
	frc.mutation.SetWithdrawRate(f)
	return frc
}

// SetID sets the "id" field.
func (frc *FractionRuleCreate) SetID(u uint32) *FractionRuleCreate {
	frc.mutation.SetID(u)
	return frc
}

// Mutation returns the FractionRuleMutation object of the builder.
func (frc *FractionRuleCreate) Mutation() *FractionRuleMutation {
	return frc.mutation
}

// Save creates the FractionRule in the database.
func (frc *FractionRuleCreate) Save(ctx context.Context) (*FractionRule, error) {
	var (
		err  error
		node *FractionRule
	)
	if err := frc.defaults(); err != nil {
		return nil, err
	}
	if len(frc.hooks) == 0 {
		if err = frc.check(); err != nil {
			return nil, err
		}
		node, err = frc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FractionRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = frc.check(); err != nil {
				return nil, err
			}
			frc.mutation = mutation
			if node, err = frc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(frc.hooks) - 1; i >= 0; i-- {
			if frc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = frc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, frc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*FractionRule)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FractionRuleMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (frc *FractionRuleCreate) SaveX(ctx context.Context) *FractionRule {
	v, err := frc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (frc *FractionRuleCreate) Exec(ctx context.Context) error {
	_, err := frc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (frc *FractionRuleCreate) ExecX(ctx context.Context) {
	if err := frc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (frc *FractionRuleCreate) defaults() error {
	if _, ok := frc.mutation.CreatedAt(); !ok {
		if fractionrule.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized fractionrule.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := fractionrule.DefaultCreatedAt()
		frc.mutation.SetCreatedAt(v)
	}
	if _, ok := frc.mutation.UpdatedAt(); !ok {
		if fractionrule.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized fractionrule.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := fractionrule.DefaultUpdatedAt()
		frc.mutation.SetUpdatedAt(v)
	}
	if _, ok := frc.mutation.DeletedAt(); !ok {
		if fractionrule.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized fractionrule.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := fractionrule.DefaultDeletedAt()
		frc.mutation.SetDeletedAt(v)
	}
	if _, ok := frc.mutation.EntID(); !ok {
		if fractionrule.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized fractionrule.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := fractionrule.DefaultEntID()
		frc.mutation.SetEntID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (frc *FractionRuleCreate) check() error {
	if _, ok := frc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FractionRule.created_at"`)}
	}
	if _, ok := frc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FractionRule.updated_at"`)}
	}
	if _, ok := frc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "FractionRule.deleted_at"`)}
	}
	if _, ok := frc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "FractionRule.ent_id"`)}
	}
	if _, ok := frc.mutation.MiningpoolType(); !ok {
		return &ValidationError{Name: "miningpool_type", err: errors.New(`ent: missing required field "FractionRule.miningpool_type"`)}
	}
	if _, ok := frc.mutation.CoinType(); !ok {
		return &ValidationError{Name: "coin_type", err: errors.New(`ent: missing required field "FractionRule.coin_type"`)}
	}
	if _, ok := frc.mutation.WithdrawInterval(); !ok {
		return &ValidationError{Name: "withdraw_interval", err: errors.New(`ent: missing required field "FractionRule.withdraw_interval"`)}
	}
	if _, ok := frc.mutation.MinAmount(); !ok {
		return &ValidationError{Name: "min_amount", err: errors.New(`ent: missing required field "FractionRule.min_amount"`)}
	}
	if _, ok := frc.mutation.WithdrawRate(); !ok {
		return &ValidationError{Name: "withdraw_rate", err: errors.New(`ent: missing required field "FractionRule.withdraw_rate"`)}
	}
	return nil
}

func (frc *FractionRuleCreate) sqlSave(ctx context.Context) (*FractionRule, error) {
	_node, _spec := frc.createSpec()
	if err := sqlgraph.CreateNode(ctx, frc.driver, _spec); err != nil {
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

func (frc *FractionRuleCreate) createSpec() (*FractionRule, *sqlgraph.CreateSpec) {
	var (
		_node = &FractionRule{config: frc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: fractionrule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: fractionrule.FieldID,
			},
		}
	)
	_spec.OnConflict = frc.conflict
	if id, ok := frc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := frc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionrule.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := frc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionrule.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := frc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionrule.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := frc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fractionrule.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := frc.mutation.MiningpoolType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fractionrule.FieldMiningpoolType,
		})
		_node.MiningpoolType = value
	}
	if value, ok := frc.mutation.CoinType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fractionrule.FieldCoinType,
		})
		_node.CoinType = value
	}
	if value, ok := frc.mutation.WithdrawInterval(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionrule.FieldWithdrawInterval,
		})
		_node.WithdrawInterval = value
	}
	if value, ok := frc.mutation.MinAmount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: fractionrule.FieldMinAmount,
		})
		_node.MinAmount = value
	}
	if value, ok := frc.mutation.WithdrawRate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: fractionrule.FieldWithdrawRate,
		})
		_node.WithdrawRate = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FractionRule.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FractionRuleUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (frc *FractionRuleCreate) OnConflict(opts ...sql.ConflictOption) *FractionRuleUpsertOne {
	frc.conflict = opts
	return &FractionRuleUpsertOne{
		create: frc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FractionRule.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (frc *FractionRuleCreate) OnConflictColumns(columns ...string) *FractionRuleUpsertOne {
	frc.conflict = append(frc.conflict, sql.ConflictColumns(columns...))
	return &FractionRuleUpsertOne{
		create: frc,
	}
}

type (
	// FractionRuleUpsertOne is the builder for "upsert"-ing
	//  one FractionRule node.
	FractionRuleUpsertOne struct {
		create *FractionRuleCreate
	}

	// FractionRuleUpsert is the "OnConflict" setter.
	FractionRuleUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *FractionRuleUpsert) SetCreatedAt(v uint32) *FractionRuleUpsert {
	u.Set(fractionrule.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FractionRuleUpsert) UpdateCreatedAt() *FractionRuleUpsert {
	u.SetExcluded(fractionrule.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FractionRuleUpsert) AddCreatedAt(v uint32) *FractionRuleUpsert {
	u.Add(fractionrule.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FractionRuleUpsert) SetUpdatedAt(v uint32) *FractionRuleUpsert {
	u.Set(fractionrule.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FractionRuleUpsert) UpdateUpdatedAt() *FractionRuleUpsert {
	u.SetExcluded(fractionrule.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FractionRuleUpsert) AddUpdatedAt(v uint32) *FractionRuleUpsert {
	u.Add(fractionrule.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FractionRuleUpsert) SetDeletedAt(v uint32) *FractionRuleUpsert {
	u.Set(fractionrule.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FractionRuleUpsert) UpdateDeletedAt() *FractionRuleUpsert {
	u.SetExcluded(fractionrule.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FractionRuleUpsert) AddDeletedAt(v uint32) *FractionRuleUpsert {
	u.Add(fractionrule.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *FractionRuleUpsert) SetEntID(v uuid.UUID) *FractionRuleUpsert {
	u.Set(fractionrule.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *FractionRuleUpsert) UpdateEntID() *FractionRuleUpsert {
	u.SetExcluded(fractionrule.FieldEntID)
	return u
}

// SetMiningpoolType sets the "miningpool_type" field.
func (u *FractionRuleUpsert) SetMiningpoolType(v string) *FractionRuleUpsert {
	u.Set(fractionrule.FieldMiningpoolType, v)
	return u
}

// UpdateMiningpoolType sets the "miningpool_type" field to the value that was provided on create.
func (u *FractionRuleUpsert) UpdateMiningpoolType() *FractionRuleUpsert {
	u.SetExcluded(fractionrule.FieldMiningpoolType)
	return u
}

// SetCoinType sets the "coin_type" field.
func (u *FractionRuleUpsert) SetCoinType(v string) *FractionRuleUpsert {
	u.Set(fractionrule.FieldCoinType, v)
	return u
}

// UpdateCoinType sets the "coin_type" field to the value that was provided on create.
func (u *FractionRuleUpsert) UpdateCoinType() *FractionRuleUpsert {
	u.SetExcluded(fractionrule.FieldCoinType)
	return u
}

// SetWithdrawInterval sets the "withdraw_interval" field.
func (u *FractionRuleUpsert) SetWithdrawInterval(v uint32) *FractionRuleUpsert {
	u.Set(fractionrule.FieldWithdrawInterval, v)
	return u
}

// UpdateWithdrawInterval sets the "withdraw_interval" field to the value that was provided on create.
func (u *FractionRuleUpsert) UpdateWithdrawInterval() *FractionRuleUpsert {
	u.SetExcluded(fractionrule.FieldWithdrawInterval)
	return u
}

// AddWithdrawInterval adds v to the "withdraw_interval" field.
func (u *FractionRuleUpsert) AddWithdrawInterval(v uint32) *FractionRuleUpsert {
	u.Add(fractionrule.FieldWithdrawInterval, v)
	return u
}

// SetMinAmount sets the "min_amount" field.
func (u *FractionRuleUpsert) SetMinAmount(v float32) *FractionRuleUpsert {
	u.Set(fractionrule.FieldMinAmount, v)
	return u
}

// UpdateMinAmount sets the "min_amount" field to the value that was provided on create.
func (u *FractionRuleUpsert) UpdateMinAmount() *FractionRuleUpsert {
	u.SetExcluded(fractionrule.FieldMinAmount)
	return u
}

// AddMinAmount adds v to the "min_amount" field.
func (u *FractionRuleUpsert) AddMinAmount(v float32) *FractionRuleUpsert {
	u.Add(fractionrule.FieldMinAmount, v)
	return u
}

// SetWithdrawRate sets the "withdraw_rate" field.
func (u *FractionRuleUpsert) SetWithdrawRate(v float32) *FractionRuleUpsert {
	u.Set(fractionrule.FieldWithdrawRate, v)
	return u
}

// UpdateWithdrawRate sets the "withdraw_rate" field to the value that was provided on create.
func (u *FractionRuleUpsert) UpdateWithdrawRate() *FractionRuleUpsert {
	u.SetExcluded(fractionrule.FieldWithdrawRate)
	return u
}

// AddWithdrawRate adds v to the "withdraw_rate" field.
func (u *FractionRuleUpsert) AddWithdrawRate(v float32) *FractionRuleUpsert {
	u.Add(fractionrule.FieldWithdrawRate, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.FractionRule.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(fractionrule.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FractionRuleUpsertOne) UpdateNewValues() *FractionRuleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(fractionrule.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FractionRule.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FractionRuleUpsertOne) Ignore() *FractionRuleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FractionRuleUpsertOne) DoNothing() *FractionRuleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FractionRuleCreate.OnConflict
// documentation for more info.
func (u *FractionRuleUpsertOne) Update(set func(*FractionRuleUpsert)) *FractionRuleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FractionRuleUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FractionRuleUpsertOne) SetCreatedAt(v uint32) *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FractionRuleUpsertOne) AddCreatedAt(v uint32) *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FractionRuleUpsertOne) UpdateCreatedAt() *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FractionRuleUpsertOne) SetUpdatedAt(v uint32) *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FractionRuleUpsertOne) AddUpdatedAt(v uint32) *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FractionRuleUpsertOne) UpdateUpdatedAt() *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FractionRuleUpsertOne) SetDeletedAt(v uint32) *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FractionRuleUpsertOne) AddDeletedAt(v uint32) *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FractionRuleUpsertOne) UpdateDeletedAt() *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *FractionRuleUpsertOne) SetEntID(v uuid.UUID) *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *FractionRuleUpsertOne) UpdateEntID() *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateEntID()
	})
}

// SetMiningpoolType sets the "miningpool_type" field.
func (u *FractionRuleUpsertOne) SetMiningpoolType(v string) *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetMiningpoolType(v)
	})
}

// UpdateMiningpoolType sets the "miningpool_type" field to the value that was provided on create.
func (u *FractionRuleUpsertOne) UpdateMiningpoolType() *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateMiningpoolType()
	})
}

// SetCoinType sets the "coin_type" field.
func (u *FractionRuleUpsertOne) SetCoinType(v string) *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetCoinType(v)
	})
}

// UpdateCoinType sets the "coin_type" field to the value that was provided on create.
func (u *FractionRuleUpsertOne) UpdateCoinType() *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateCoinType()
	})
}

// SetWithdrawInterval sets the "withdraw_interval" field.
func (u *FractionRuleUpsertOne) SetWithdrawInterval(v uint32) *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetWithdrawInterval(v)
	})
}

// AddWithdrawInterval adds v to the "withdraw_interval" field.
func (u *FractionRuleUpsertOne) AddWithdrawInterval(v uint32) *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.AddWithdrawInterval(v)
	})
}

// UpdateWithdrawInterval sets the "withdraw_interval" field to the value that was provided on create.
func (u *FractionRuleUpsertOne) UpdateWithdrawInterval() *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateWithdrawInterval()
	})
}

// SetMinAmount sets the "min_amount" field.
func (u *FractionRuleUpsertOne) SetMinAmount(v float32) *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetMinAmount(v)
	})
}

// AddMinAmount adds v to the "min_amount" field.
func (u *FractionRuleUpsertOne) AddMinAmount(v float32) *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.AddMinAmount(v)
	})
}

// UpdateMinAmount sets the "min_amount" field to the value that was provided on create.
func (u *FractionRuleUpsertOne) UpdateMinAmount() *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateMinAmount()
	})
}

// SetWithdrawRate sets the "withdraw_rate" field.
func (u *FractionRuleUpsertOne) SetWithdrawRate(v float32) *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetWithdrawRate(v)
	})
}

// AddWithdrawRate adds v to the "withdraw_rate" field.
func (u *FractionRuleUpsertOne) AddWithdrawRate(v float32) *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.AddWithdrawRate(v)
	})
}

// UpdateWithdrawRate sets the "withdraw_rate" field to the value that was provided on create.
func (u *FractionRuleUpsertOne) UpdateWithdrawRate() *FractionRuleUpsertOne {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateWithdrawRate()
	})
}

// Exec executes the query.
func (u *FractionRuleUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FractionRuleCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FractionRuleUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FractionRuleUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FractionRuleUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FractionRuleCreateBulk is the builder for creating many FractionRule entities in bulk.
type FractionRuleCreateBulk struct {
	config
	builders []*FractionRuleCreate
	conflict []sql.ConflictOption
}

// Save creates the FractionRule entities in the database.
func (frcb *FractionRuleCreateBulk) Save(ctx context.Context) ([]*FractionRule, error) {
	specs := make([]*sqlgraph.CreateSpec, len(frcb.builders))
	nodes := make([]*FractionRule, len(frcb.builders))
	mutators := make([]Mutator, len(frcb.builders))
	for i := range frcb.builders {
		func(i int, root context.Context) {
			builder := frcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FractionRuleMutation)
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
					_, err = mutators[i+1].Mutate(root, frcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = frcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, frcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, frcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (frcb *FractionRuleCreateBulk) SaveX(ctx context.Context) []*FractionRule {
	v, err := frcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (frcb *FractionRuleCreateBulk) Exec(ctx context.Context) error {
	_, err := frcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (frcb *FractionRuleCreateBulk) ExecX(ctx context.Context) {
	if err := frcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FractionRule.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FractionRuleUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (frcb *FractionRuleCreateBulk) OnConflict(opts ...sql.ConflictOption) *FractionRuleUpsertBulk {
	frcb.conflict = opts
	return &FractionRuleUpsertBulk{
		create: frcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FractionRule.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (frcb *FractionRuleCreateBulk) OnConflictColumns(columns ...string) *FractionRuleUpsertBulk {
	frcb.conflict = append(frcb.conflict, sql.ConflictColumns(columns...))
	return &FractionRuleUpsertBulk{
		create: frcb,
	}
}

// FractionRuleUpsertBulk is the builder for "upsert"-ing
// a bulk of FractionRule nodes.
type FractionRuleUpsertBulk struct {
	create *FractionRuleCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.FractionRule.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(fractionrule.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FractionRuleUpsertBulk) UpdateNewValues() *FractionRuleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(fractionrule.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FractionRule.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FractionRuleUpsertBulk) Ignore() *FractionRuleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FractionRuleUpsertBulk) DoNothing() *FractionRuleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FractionRuleCreateBulk.OnConflict
// documentation for more info.
func (u *FractionRuleUpsertBulk) Update(set func(*FractionRuleUpsert)) *FractionRuleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FractionRuleUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FractionRuleUpsertBulk) SetCreatedAt(v uint32) *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FractionRuleUpsertBulk) AddCreatedAt(v uint32) *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FractionRuleUpsertBulk) UpdateCreatedAt() *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FractionRuleUpsertBulk) SetUpdatedAt(v uint32) *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FractionRuleUpsertBulk) AddUpdatedAt(v uint32) *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FractionRuleUpsertBulk) UpdateUpdatedAt() *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FractionRuleUpsertBulk) SetDeletedAt(v uint32) *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FractionRuleUpsertBulk) AddDeletedAt(v uint32) *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FractionRuleUpsertBulk) UpdateDeletedAt() *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *FractionRuleUpsertBulk) SetEntID(v uuid.UUID) *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *FractionRuleUpsertBulk) UpdateEntID() *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateEntID()
	})
}

// SetMiningpoolType sets the "miningpool_type" field.
func (u *FractionRuleUpsertBulk) SetMiningpoolType(v string) *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetMiningpoolType(v)
	})
}

// UpdateMiningpoolType sets the "miningpool_type" field to the value that was provided on create.
func (u *FractionRuleUpsertBulk) UpdateMiningpoolType() *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateMiningpoolType()
	})
}

// SetCoinType sets the "coin_type" field.
func (u *FractionRuleUpsertBulk) SetCoinType(v string) *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetCoinType(v)
	})
}

// UpdateCoinType sets the "coin_type" field to the value that was provided on create.
func (u *FractionRuleUpsertBulk) UpdateCoinType() *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateCoinType()
	})
}

// SetWithdrawInterval sets the "withdraw_interval" field.
func (u *FractionRuleUpsertBulk) SetWithdrawInterval(v uint32) *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetWithdrawInterval(v)
	})
}

// AddWithdrawInterval adds v to the "withdraw_interval" field.
func (u *FractionRuleUpsertBulk) AddWithdrawInterval(v uint32) *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.AddWithdrawInterval(v)
	})
}

// UpdateWithdrawInterval sets the "withdraw_interval" field to the value that was provided on create.
func (u *FractionRuleUpsertBulk) UpdateWithdrawInterval() *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateWithdrawInterval()
	})
}

// SetMinAmount sets the "min_amount" field.
func (u *FractionRuleUpsertBulk) SetMinAmount(v float32) *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetMinAmount(v)
	})
}

// AddMinAmount adds v to the "min_amount" field.
func (u *FractionRuleUpsertBulk) AddMinAmount(v float32) *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.AddMinAmount(v)
	})
}

// UpdateMinAmount sets the "min_amount" field to the value that was provided on create.
func (u *FractionRuleUpsertBulk) UpdateMinAmount() *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateMinAmount()
	})
}

// SetWithdrawRate sets the "withdraw_rate" field.
func (u *FractionRuleUpsertBulk) SetWithdrawRate(v float32) *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.SetWithdrawRate(v)
	})
}

// AddWithdrawRate adds v to the "withdraw_rate" field.
func (u *FractionRuleUpsertBulk) AddWithdrawRate(v float32) *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.AddWithdrawRate(v)
	})
}

// UpdateWithdrawRate sets the "withdraw_rate" field to the value that was provided on create.
func (u *FractionRuleUpsertBulk) UpdateWithdrawRate() *FractionRuleUpsertBulk {
	return u.Update(func(s *FractionRuleUpsert) {
		s.UpdateWithdrawRate()
	})
}

// Exec executes the query.
func (u *FractionRuleUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FractionRuleCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FractionRuleCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FractionRuleUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
