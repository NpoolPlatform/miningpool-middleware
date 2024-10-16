// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fractionwithdrawal"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// FractionWithdrawalUpdate is the builder for updating FractionWithdrawal entities.
type FractionWithdrawalUpdate struct {
	config
	hooks     []Hook
	mutation  *FractionWithdrawalMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the FractionWithdrawalUpdate builder.
func (fwu *FractionWithdrawalUpdate) Where(ps ...predicate.FractionWithdrawal) *FractionWithdrawalUpdate {
	fwu.mutation.Where(ps...)
	return fwu
}

// SetCreatedAt sets the "created_at" field.
func (fwu *FractionWithdrawalUpdate) SetCreatedAt(u uint32) *FractionWithdrawalUpdate {
	fwu.mutation.ResetCreatedAt()
	fwu.mutation.SetCreatedAt(u)
	return fwu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fwu *FractionWithdrawalUpdate) SetNillableCreatedAt(u *uint32) *FractionWithdrawalUpdate {
	if u != nil {
		fwu.SetCreatedAt(*u)
	}
	return fwu
}

// AddCreatedAt adds u to the "created_at" field.
func (fwu *FractionWithdrawalUpdate) AddCreatedAt(u int32) *FractionWithdrawalUpdate {
	fwu.mutation.AddCreatedAt(u)
	return fwu
}

// SetUpdatedAt sets the "updated_at" field.
func (fwu *FractionWithdrawalUpdate) SetUpdatedAt(u uint32) *FractionWithdrawalUpdate {
	fwu.mutation.ResetUpdatedAt()
	fwu.mutation.SetUpdatedAt(u)
	return fwu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (fwu *FractionWithdrawalUpdate) AddUpdatedAt(u int32) *FractionWithdrawalUpdate {
	fwu.mutation.AddUpdatedAt(u)
	return fwu
}

// SetDeletedAt sets the "deleted_at" field.
func (fwu *FractionWithdrawalUpdate) SetDeletedAt(u uint32) *FractionWithdrawalUpdate {
	fwu.mutation.ResetDeletedAt()
	fwu.mutation.SetDeletedAt(u)
	return fwu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fwu *FractionWithdrawalUpdate) SetNillableDeletedAt(u *uint32) *FractionWithdrawalUpdate {
	if u != nil {
		fwu.SetDeletedAt(*u)
	}
	return fwu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (fwu *FractionWithdrawalUpdate) AddDeletedAt(u int32) *FractionWithdrawalUpdate {
	fwu.mutation.AddDeletedAt(u)
	return fwu
}

// SetEntID sets the "ent_id" field.
func (fwu *FractionWithdrawalUpdate) SetEntID(u uuid.UUID) *FractionWithdrawalUpdate {
	fwu.mutation.SetEntID(u)
	return fwu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (fwu *FractionWithdrawalUpdate) SetNillableEntID(u *uuid.UUID) *FractionWithdrawalUpdate {
	if u != nil {
		fwu.SetEntID(*u)
	}
	return fwu
}

// SetAppID sets the "app_id" field.
func (fwu *FractionWithdrawalUpdate) SetAppID(u uuid.UUID) *FractionWithdrawalUpdate {
	fwu.mutation.SetAppID(u)
	return fwu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (fwu *FractionWithdrawalUpdate) SetNillableAppID(u *uuid.UUID) *FractionWithdrawalUpdate {
	if u != nil {
		fwu.SetAppID(*u)
	}
	return fwu
}

// ClearAppID clears the value of the "app_id" field.
func (fwu *FractionWithdrawalUpdate) ClearAppID() *FractionWithdrawalUpdate {
	fwu.mutation.ClearAppID()
	return fwu
}

// SetUserID sets the "user_id" field.
func (fwu *FractionWithdrawalUpdate) SetUserID(u uuid.UUID) *FractionWithdrawalUpdate {
	fwu.mutation.SetUserID(u)
	return fwu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (fwu *FractionWithdrawalUpdate) SetNillableUserID(u *uuid.UUID) *FractionWithdrawalUpdate {
	if u != nil {
		fwu.SetUserID(*u)
	}
	return fwu
}

// ClearUserID clears the value of the "user_id" field.
func (fwu *FractionWithdrawalUpdate) ClearUserID() *FractionWithdrawalUpdate {
	fwu.mutation.ClearUserID()
	return fwu
}

// SetOrderUserID sets the "order_user_id" field.
func (fwu *FractionWithdrawalUpdate) SetOrderUserID(u uuid.UUID) *FractionWithdrawalUpdate {
	fwu.mutation.SetOrderUserID(u)
	return fwu
}

// SetNillableOrderUserID sets the "order_user_id" field if the given value is not nil.
func (fwu *FractionWithdrawalUpdate) SetNillableOrderUserID(u *uuid.UUID) *FractionWithdrawalUpdate {
	if u != nil {
		fwu.SetOrderUserID(*u)
	}
	return fwu
}

// ClearOrderUserID clears the value of the "order_user_id" field.
func (fwu *FractionWithdrawalUpdate) ClearOrderUserID() *FractionWithdrawalUpdate {
	fwu.mutation.ClearOrderUserID()
	return fwu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (fwu *FractionWithdrawalUpdate) SetCoinTypeID(u uuid.UUID) *FractionWithdrawalUpdate {
	fwu.mutation.SetCoinTypeID(u)
	return fwu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (fwu *FractionWithdrawalUpdate) SetNillableCoinTypeID(u *uuid.UUID) *FractionWithdrawalUpdate {
	if u != nil {
		fwu.SetCoinTypeID(*u)
	}
	return fwu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (fwu *FractionWithdrawalUpdate) ClearCoinTypeID() *FractionWithdrawalUpdate {
	fwu.mutation.ClearCoinTypeID()
	return fwu
}

// SetFractionWithdrawState sets the "fraction_withdraw_state" field.
func (fwu *FractionWithdrawalUpdate) SetFractionWithdrawState(s string) *FractionWithdrawalUpdate {
	fwu.mutation.SetFractionWithdrawState(s)
	return fwu
}

// SetNillableFractionWithdrawState sets the "fraction_withdraw_state" field if the given value is not nil.
func (fwu *FractionWithdrawalUpdate) SetNillableFractionWithdrawState(s *string) *FractionWithdrawalUpdate {
	if s != nil {
		fwu.SetFractionWithdrawState(*s)
	}
	return fwu
}

// ClearFractionWithdrawState clears the value of the "fraction_withdraw_state" field.
func (fwu *FractionWithdrawalUpdate) ClearFractionWithdrawState() *FractionWithdrawalUpdate {
	fwu.mutation.ClearFractionWithdrawState()
	return fwu
}

// SetWithdrawAt sets the "withdraw_at" field.
func (fwu *FractionWithdrawalUpdate) SetWithdrawAt(u uint32) *FractionWithdrawalUpdate {
	fwu.mutation.ResetWithdrawAt()
	fwu.mutation.SetWithdrawAt(u)
	return fwu
}

// SetNillableWithdrawAt sets the "withdraw_at" field if the given value is not nil.
func (fwu *FractionWithdrawalUpdate) SetNillableWithdrawAt(u *uint32) *FractionWithdrawalUpdate {
	if u != nil {
		fwu.SetWithdrawAt(*u)
	}
	return fwu
}

// AddWithdrawAt adds u to the "withdraw_at" field.
func (fwu *FractionWithdrawalUpdate) AddWithdrawAt(u int32) *FractionWithdrawalUpdate {
	fwu.mutation.AddWithdrawAt(u)
	return fwu
}

// ClearWithdrawAt clears the value of the "withdraw_at" field.
func (fwu *FractionWithdrawalUpdate) ClearWithdrawAt() *FractionWithdrawalUpdate {
	fwu.mutation.ClearWithdrawAt()
	return fwu
}

// SetPromisePayAt sets the "promise_pay_at" field.
func (fwu *FractionWithdrawalUpdate) SetPromisePayAt(u uint32) *FractionWithdrawalUpdate {
	fwu.mutation.ResetPromisePayAt()
	fwu.mutation.SetPromisePayAt(u)
	return fwu
}

// SetNillablePromisePayAt sets the "promise_pay_at" field if the given value is not nil.
func (fwu *FractionWithdrawalUpdate) SetNillablePromisePayAt(u *uint32) *FractionWithdrawalUpdate {
	if u != nil {
		fwu.SetPromisePayAt(*u)
	}
	return fwu
}

// AddPromisePayAt adds u to the "promise_pay_at" field.
func (fwu *FractionWithdrawalUpdate) AddPromisePayAt(u int32) *FractionWithdrawalUpdate {
	fwu.mutation.AddPromisePayAt(u)
	return fwu
}

// ClearPromisePayAt clears the value of the "promise_pay_at" field.
func (fwu *FractionWithdrawalUpdate) ClearPromisePayAt() *FractionWithdrawalUpdate {
	fwu.mutation.ClearPromisePayAt()
	return fwu
}

// SetMsg sets the "msg" field.
func (fwu *FractionWithdrawalUpdate) SetMsg(s string) *FractionWithdrawalUpdate {
	fwu.mutation.SetMsg(s)
	return fwu
}

// SetNillableMsg sets the "msg" field if the given value is not nil.
func (fwu *FractionWithdrawalUpdate) SetNillableMsg(s *string) *FractionWithdrawalUpdate {
	if s != nil {
		fwu.SetMsg(*s)
	}
	return fwu
}

// ClearMsg clears the value of the "msg" field.
func (fwu *FractionWithdrawalUpdate) ClearMsg() *FractionWithdrawalUpdate {
	fwu.mutation.ClearMsg()
	return fwu
}

// Mutation returns the FractionWithdrawalMutation object of the builder.
func (fwu *FractionWithdrawalUpdate) Mutation() *FractionWithdrawalMutation {
	return fwu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fwu *FractionWithdrawalUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := fwu.defaults(); err != nil {
		return 0, err
	}
	if len(fwu.hooks) == 0 {
		affected, err = fwu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FractionWithdrawalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fwu.mutation = mutation
			affected, err = fwu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fwu.hooks) - 1; i >= 0; i-- {
			if fwu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fwu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fwu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fwu *FractionWithdrawalUpdate) SaveX(ctx context.Context) int {
	affected, err := fwu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fwu *FractionWithdrawalUpdate) Exec(ctx context.Context) error {
	_, err := fwu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fwu *FractionWithdrawalUpdate) ExecX(ctx context.Context) {
	if err := fwu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fwu *FractionWithdrawalUpdate) defaults() error {
	if _, ok := fwu.mutation.UpdatedAt(); !ok {
		if fractionwithdrawal.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized fractionwithdrawal.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := fractionwithdrawal.UpdateDefaultUpdatedAt()
		fwu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fwu *FractionWithdrawalUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FractionWithdrawalUpdate {
	fwu.modifiers = append(fwu.modifiers, modifiers...)
	return fwu
}

func (fwu *FractionWithdrawalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fractionwithdrawal.Table,
			Columns: fractionwithdrawal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: fractionwithdrawal.FieldID,
			},
		},
	}
	if ps := fwu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fwu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldCreatedAt,
		})
	}
	if value, ok := fwu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldCreatedAt,
		})
	}
	if value, ok := fwu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldUpdatedAt,
		})
	}
	if value, ok := fwu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldUpdatedAt,
		})
	}
	if value, ok := fwu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldDeletedAt,
		})
	}
	if value, ok := fwu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldDeletedAt,
		})
	}
	if value, ok := fwu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fractionwithdrawal.FieldEntID,
		})
	}
	if value, ok := fwu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fractionwithdrawal.FieldAppID,
		})
	}
	if fwu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fractionwithdrawal.FieldAppID,
		})
	}
	if value, ok := fwu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fractionwithdrawal.FieldUserID,
		})
	}
	if fwu.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fractionwithdrawal.FieldUserID,
		})
	}
	if value, ok := fwu.mutation.OrderUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fractionwithdrawal.FieldOrderUserID,
		})
	}
	if fwu.mutation.OrderUserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fractionwithdrawal.FieldOrderUserID,
		})
	}
	if value, ok := fwu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fractionwithdrawal.FieldCoinTypeID,
		})
	}
	if fwu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fractionwithdrawal.FieldCoinTypeID,
		})
	}
	if value, ok := fwu.mutation.FractionWithdrawState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fractionwithdrawal.FieldFractionWithdrawState,
		})
	}
	if fwu.mutation.FractionWithdrawStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: fractionwithdrawal.FieldFractionWithdrawState,
		})
	}
	if value, ok := fwu.mutation.WithdrawAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldWithdrawAt,
		})
	}
	if value, ok := fwu.mutation.AddedWithdrawAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldWithdrawAt,
		})
	}
	if fwu.mutation.WithdrawAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: fractionwithdrawal.FieldWithdrawAt,
		})
	}
	if value, ok := fwu.mutation.PromisePayAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldPromisePayAt,
		})
	}
	if value, ok := fwu.mutation.AddedPromisePayAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldPromisePayAt,
		})
	}
	if fwu.mutation.PromisePayAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: fractionwithdrawal.FieldPromisePayAt,
		})
	}
	if value, ok := fwu.mutation.Msg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fractionwithdrawal.FieldMsg,
		})
	}
	if fwu.mutation.MsgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: fractionwithdrawal.FieldMsg,
		})
	}
	_spec.Modifiers = fwu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, fwu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fractionwithdrawal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// FractionWithdrawalUpdateOne is the builder for updating a single FractionWithdrawal entity.
type FractionWithdrawalUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *FractionWithdrawalMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (fwuo *FractionWithdrawalUpdateOne) SetCreatedAt(u uint32) *FractionWithdrawalUpdateOne {
	fwuo.mutation.ResetCreatedAt()
	fwuo.mutation.SetCreatedAt(u)
	return fwuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fwuo *FractionWithdrawalUpdateOne) SetNillableCreatedAt(u *uint32) *FractionWithdrawalUpdateOne {
	if u != nil {
		fwuo.SetCreatedAt(*u)
	}
	return fwuo
}

// AddCreatedAt adds u to the "created_at" field.
func (fwuo *FractionWithdrawalUpdateOne) AddCreatedAt(u int32) *FractionWithdrawalUpdateOne {
	fwuo.mutation.AddCreatedAt(u)
	return fwuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fwuo *FractionWithdrawalUpdateOne) SetUpdatedAt(u uint32) *FractionWithdrawalUpdateOne {
	fwuo.mutation.ResetUpdatedAt()
	fwuo.mutation.SetUpdatedAt(u)
	return fwuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (fwuo *FractionWithdrawalUpdateOne) AddUpdatedAt(u int32) *FractionWithdrawalUpdateOne {
	fwuo.mutation.AddUpdatedAt(u)
	return fwuo
}

// SetDeletedAt sets the "deleted_at" field.
func (fwuo *FractionWithdrawalUpdateOne) SetDeletedAt(u uint32) *FractionWithdrawalUpdateOne {
	fwuo.mutation.ResetDeletedAt()
	fwuo.mutation.SetDeletedAt(u)
	return fwuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fwuo *FractionWithdrawalUpdateOne) SetNillableDeletedAt(u *uint32) *FractionWithdrawalUpdateOne {
	if u != nil {
		fwuo.SetDeletedAt(*u)
	}
	return fwuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (fwuo *FractionWithdrawalUpdateOne) AddDeletedAt(u int32) *FractionWithdrawalUpdateOne {
	fwuo.mutation.AddDeletedAt(u)
	return fwuo
}

// SetEntID sets the "ent_id" field.
func (fwuo *FractionWithdrawalUpdateOne) SetEntID(u uuid.UUID) *FractionWithdrawalUpdateOne {
	fwuo.mutation.SetEntID(u)
	return fwuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (fwuo *FractionWithdrawalUpdateOne) SetNillableEntID(u *uuid.UUID) *FractionWithdrawalUpdateOne {
	if u != nil {
		fwuo.SetEntID(*u)
	}
	return fwuo
}

// SetAppID sets the "app_id" field.
func (fwuo *FractionWithdrawalUpdateOne) SetAppID(u uuid.UUID) *FractionWithdrawalUpdateOne {
	fwuo.mutation.SetAppID(u)
	return fwuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (fwuo *FractionWithdrawalUpdateOne) SetNillableAppID(u *uuid.UUID) *FractionWithdrawalUpdateOne {
	if u != nil {
		fwuo.SetAppID(*u)
	}
	return fwuo
}

// ClearAppID clears the value of the "app_id" field.
func (fwuo *FractionWithdrawalUpdateOne) ClearAppID() *FractionWithdrawalUpdateOne {
	fwuo.mutation.ClearAppID()
	return fwuo
}

// SetUserID sets the "user_id" field.
func (fwuo *FractionWithdrawalUpdateOne) SetUserID(u uuid.UUID) *FractionWithdrawalUpdateOne {
	fwuo.mutation.SetUserID(u)
	return fwuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (fwuo *FractionWithdrawalUpdateOne) SetNillableUserID(u *uuid.UUID) *FractionWithdrawalUpdateOne {
	if u != nil {
		fwuo.SetUserID(*u)
	}
	return fwuo
}

// ClearUserID clears the value of the "user_id" field.
func (fwuo *FractionWithdrawalUpdateOne) ClearUserID() *FractionWithdrawalUpdateOne {
	fwuo.mutation.ClearUserID()
	return fwuo
}

// SetOrderUserID sets the "order_user_id" field.
func (fwuo *FractionWithdrawalUpdateOne) SetOrderUserID(u uuid.UUID) *FractionWithdrawalUpdateOne {
	fwuo.mutation.SetOrderUserID(u)
	return fwuo
}

// SetNillableOrderUserID sets the "order_user_id" field if the given value is not nil.
func (fwuo *FractionWithdrawalUpdateOne) SetNillableOrderUserID(u *uuid.UUID) *FractionWithdrawalUpdateOne {
	if u != nil {
		fwuo.SetOrderUserID(*u)
	}
	return fwuo
}

// ClearOrderUserID clears the value of the "order_user_id" field.
func (fwuo *FractionWithdrawalUpdateOne) ClearOrderUserID() *FractionWithdrawalUpdateOne {
	fwuo.mutation.ClearOrderUserID()
	return fwuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (fwuo *FractionWithdrawalUpdateOne) SetCoinTypeID(u uuid.UUID) *FractionWithdrawalUpdateOne {
	fwuo.mutation.SetCoinTypeID(u)
	return fwuo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (fwuo *FractionWithdrawalUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *FractionWithdrawalUpdateOne {
	if u != nil {
		fwuo.SetCoinTypeID(*u)
	}
	return fwuo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (fwuo *FractionWithdrawalUpdateOne) ClearCoinTypeID() *FractionWithdrawalUpdateOne {
	fwuo.mutation.ClearCoinTypeID()
	return fwuo
}

// SetFractionWithdrawState sets the "fraction_withdraw_state" field.
func (fwuo *FractionWithdrawalUpdateOne) SetFractionWithdrawState(s string) *FractionWithdrawalUpdateOne {
	fwuo.mutation.SetFractionWithdrawState(s)
	return fwuo
}

// SetNillableFractionWithdrawState sets the "fraction_withdraw_state" field if the given value is not nil.
func (fwuo *FractionWithdrawalUpdateOne) SetNillableFractionWithdrawState(s *string) *FractionWithdrawalUpdateOne {
	if s != nil {
		fwuo.SetFractionWithdrawState(*s)
	}
	return fwuo
}

// ClearFractionWithdrawState clears the value of the "fraction_withdraw_state" field.
func (fwuo *FractionWithdrawalUpdateOne) ClearFractionWithdrawState() *FractionWithdrawalUpdateOne {
	fwuo.mutation.ClearFractionWithdrawState()
	return fwuo
}

// SetWithdrawAt sets the "withdraw_at" field.
func (fwuo *FractionWithdrawalUpdateOne) SetWithdrawAt(u uint32) *FractionWithdrawalUpdateOne {
	fwuo.mutation.ResetWithdrawAt()
	fwuo.mutation.SetWithdrawAt(u)
	return fwuo
}

// SetNillableWithdrawAt sets the "withdraw_at" field if the given value is not nil.
func (fwuo *FractionWithdrawalUpdateOne) SetNillableWithdrawAt(u *uint32) *FractionWithdrawalUpdateOne {
	if u != nil {
		fwuo.SetWithdrawAt(*u)
	}
	return fwuo
}

// AddWithdrawAt adds u to the "withdraw_at" field.
func (fwuo *FractionWithdrawalUpdateOne) AddWithdrawAt(u int32) *FractionWithdrawalUpdateOne {
	fwuo.mutation.AddWithdrawAt(u)
	return fwuo
}

// ClearWithdrawAt clears the value of the "withdraw_at" field.
func (fwuo *FractionWithdrawalUpdateOne) ClearWithdrawAt() *FractionWithdrawalUpdateOne {
	fwuo.mutation.ClearWithdrawAt()
	return fwuo
}

// SetPromisePayAt sets the "promise_pay_at" field.
func (fwuo *FractionWithdrawalUpdateOne) SetPromisePayAt(u uint32) *FractionWithdrawalUpdateOne {
	fwuo.mutation.ResetPromisePayAt()
	fwuo.mutation.SetPromisePayAt(u)
	return fwuo
}

// SetNillablePromisePayAt sets the "promise_pay_at" field if the given value is not nil.
func (fwuo *FractionWithdrawalUpdateOne) SetNillablePromisePayAt(u *uint32) *FractionWithdrawalUpdateOne {
	if u != nil {
		fwuo.SetPromisePayAt(*u)
	}
	return fwuo
}

// AddPromisePayAt adds u to the "promise_pay_at" field.
func (fwuo *FractionWithdrawalUpdateOne) AddPromisePayAt(u int32) *FractionWithdrawalUpdateOne {
	fwuo.mutation.AddPromisePayAt(u)
	return fwuo
}

// ClearPromisePayAt clears the value of the "promise_pay_at" field.
func (fwuo *FractionWithdrawalUpdateOne) ClearPromisePayAt() *FractionWithdrawalUpdateOne {
	fwuo.mutation.ClearPromisePayAt()
	return fwuo
}

// SetMsg sets the "msg" field.
func (fwuo *FractionWithdrawalUpdateOne) SetMsg(s string) *FractionWithdrawalUpdateOne {
	fwuo.mutation.SetMsg(s)
	return fwuo
}

// SetNillableMsg sets the "msg" field if the given value is not nil.
func (fwuo *FractionWithdrawalUpdateOne) SetNillableMsg(s *string) *FractionWithdrawalUpdateOne {
	if s != nil {
		fwuo.SetMsg(*s)
	}
	return fwuo
}

// ClearMsg clears the value of the "msg" field.
func (fwuo *FractionWithdrawalUpdateOne) ClearMsg() *FractionWithdrawalUpdateOne {
	fwuo.mutation.ClearMsg()
	return fwuo
}

// Mutation returns the FractionWithdrawalMutation object of the builder.
func (fwuo *FractionWithdrawalUpdateOne) Mutation() *FractionWithdrawalMutation {
	return fwuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fwuo *FractionWithdrawalUpdateOne) Select(field string, fields ...string) *FractionWithdrawalUpdateOne {
	fwuo.fields = append([]string{field}, fields...)
	return fwuo
}

// Save executes the query and returns the updated FractionWithdrawal entity.
func (fwuo *FractionWithdrawalUpdateOne) Save(ctx context.Context) (*FractionWithdrawal, error) {
	var (
		err  error
		node *FractionWithdrawal
	)
	if err := fwuo.defaults(); err != nil {
		return nil, err
	}
	if len(fwuo.hooks) == 0 {
		node, err = fwuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FractionWithdrawalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fwuo.mutation = mutation
			node, err = fwuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fwuo.hooks) - 1; i >= 0; i-- {
			if fwuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fwuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fwuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*FractionWithdrawal)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FractionWithdrawalMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fwuo *FractionWithdrawalUpdateOne) SaveX(ctx context.Context) *FractionWithdrawal {
	node, err := fwuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fwuo *FractionWithdrawalUpdateOne) Exec(ctx context.Context) error {
	_, err := fwuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fwuo *FractionWithdrawalUpdateOne) ExecX(ctx context.Context) {
	if err := fwuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fwuo *FractionWithdrawalUpdateOne) defaults() error {
	if _, ok := fwuo.mutation.UpdatedAt(); !ok {
		if fractionwithdrawal.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized fractionwithdrawal.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := fractionwithdrawal.UpdateDefaultUpdatedAt()
		fwuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fwuo *FractionWithdrawalUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FractionWithdrawalUpdateOne {
	fwuo.modifiers = append(fwuo.modifiers, modifiers...)
	return fwuo
}

func (fwuo *FractionWithdrawalUpdateOne) sqlSave(ctx context.Context) (_node *FractionWithdrawal, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fractionwithdrawal.Table,
			Columns: fractionwithdrawal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: fractionwithdrawal.FieldID,
			},
		},
	}
	id, ok := fwuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FractionWithdrawal.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fwuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fractionwithdrawal.FieldID)
		for _, f := range fields {
			if !fractionwithdrawal.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != fractionwithdrawal.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fwuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fwuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldCreatedAt,
		})
	}
	if value, ok := fwuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldCreatedAt,
		})
	}
	if value, ok := fwuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldUpdatedAt,
		})
	}
	if value, ok := fwuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldUpdatedAt,
		})
	}
	if value, ok := fwuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldDeletedAt,
		})
	}
	if value, ok := fwuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldDeletedAt,
		})
	}
	if value, ok := fwuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fractionwithdrawal.FieldEntID,
		})
	}
	if value, ok := fwuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fractionwithdrawal.FieldAppID,
		})
	}
	if fwuo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fractionwithdrawal.FieldAppID,
		})
	}
	if value, ok := fwuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fractionwithdrawal.FieldUserID,
		})
	}
	if fwuo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fractionwithdrawal.FieldUserID,
		})
	}
	if value, ok := fwuo.mutation.OrderUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fractionwithdrawal.FieldOrderUserID,
		})
	}
	if fwuo.mutation.OrderUserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fractionwithdrawal.FieldOrderUserID,
		})
	}
	if value, ok := fwuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fractionwithdrawal.FieldCoinTypeID,
		})
	}
	if fwuo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fractionwithdrawal.FieldCoinTypeID,
		})
	}
	if value, ok := fwuo.mutation.FractionWithdrawState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fractionwithdrawal.FieldFractionWithdrawState,
		})
	}
	if fwuo.mutation.FractionWithdrawStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: fractionwithdrawal.FieldFractionWithdrawState,
		})
	}
	if value, ok := fwuo.mutation.WithdrawAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldWithdrawAt,
		})
	}
	if value, ok := fwuo.mutation.AddedWithdrawAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldWithdrawAt,
		})
	}
	if fwuo.mutation.WithdrawAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: fractionwithdrawal.FieldWithdrawAt,
		})
	}
	if value, ok := fwuo.mutation.PromisePayAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldPromisePayAt,
		})
	}
	if value, ok := fwuo.mutation.AddedPromisePayAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fractionwithdrawal.FieldPromisePayAt,
		})
	}
	if fwuo.mutation.PromisePayAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: fractionwithdrawal.FieldPromisePayAt,
		})
	}
	if value, ok := fwuo.mutation.Msg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fractionwithdrawal.FieldMsg,
		})
	}
	if fwuo.mutation.MsgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: fractionwithdrawal.FieldMsg,
		})
	}
	_spec.Modifiers = fwuo.modifiers
	_node = &FractionWithdrawal{config: fwuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fwuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fractionwithdrawal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}