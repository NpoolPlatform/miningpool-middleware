// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fraction"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// FractionUpdate is the builder for updating Fraction entities.
type FractionUpdate struct {
	config
	hooks     []Hook
	mutation  *FractionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the FractionUpdate builder.
func (fu *FractionUpdate) Where(ps ...predicate.Fraction) *FractionUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetCreatedAt sets the "created_at" field.
func (fu *FractionUpdate) SetCreatedAt(u uint32) *FractionUpdate {
	fu.mutation.ResetCreatedAt()
	fu.mutation.SetCreatedAt(u)
	return fu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fu *FractionUpdate) SetNillableCreatedAt(u *uint32) *FractionUpdate {
	if u != nil {
		fu.SetCreatedAt(*u)
	}
	return fu
}

// AddCreatedAt adds u to the "created_at" field.
func (fu *FractionUpdate) AddCreatedAt(u int32) *FractionUpdate {
	fu.mutation.AddCreatedAt(u)
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FractionUpdate) SetUpdatedAt(u uint32) *FractionUpdate {
	fu.mutation.ResetUpdatedAt()
	fu.mutation.SetUpdatedAt(u)
	return fu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (fu *FractionUpdate) AddUpdatedAt(u int32) *FractionUpdate {
	fu.mutation.AddUpdatedAt(u)
	return fu
}

// SetDeletedAt sets the "deleted_at" field.
func (fu *FractionUpdate) SetDeletedAt(u uint32) *FractionUpdate {
	fu.mutation.ResetDeletedAt()
	fu.mutation.SetDeletedAt(u)
	return fu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fu *FractionUpdate) SetNillableDeletedAt(u *uint32) *FractionUpdate {
	if u != nil {
		fu.SetDeletedAt(*u)
	}
	return fu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (fu *FractionUpdate) AddDeletedAt(u int32) *FractionUpdate {
	fu.mutation.AddDeletedAt(u)
	return fu
}

// SetEntID sets the "ent_id" field.
func (fu *FractionUpdate) SetEntID(u uuid.UUID) *FractionUpdate {
	fu.mutation.SetEntID(u)
	return fu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (fu *FractionUpdate) SetNillableEntID(u *uuid.UUID) *FractionUpdate {
	if u != nil {
		fu.SetEntID(*u)
	}
	return fu
}

// SetAppID sets the "app_id" field.
func (fu *FractionUpdate) SetAppID(u uuid.UUID) *FractionUpdate {
	fu.mutation.SetAppID(u)
	return fu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (fu *FractionUpdate) SetNillableAppID(u *uuid.UUID) *FractionUpdate {
	if u != nil {
		fu.SetAppID(*u)
	}
	return fu
}

// ClearAppID clears the value of the "app_id" field.
func (fu *FractionUpdate) ClearAppID() *FractionUpdate {
	fu.mutation.ClearAppID()
	return fu
}

// SetUserID sets the "user_id" field.
func (fu *FractionUpdate) SetUserID(u uuid.UUID) *FractionUpdate {
	fu.mutation.SetUserID(u)
	return fu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (fu *FractionUpdate) SetNillableUserID(u *uuid.UUID) *FractionUpdate {
	if u != nil {
		fu.SetUserID(*u)
	}
	return fu
}

// ClearUserID clears the value of the "user_id" field.
func (fu *FractionUpdate) ClearUserID() *FractionUpdate {
	fu.mutation.ClearUserID()
	return fu
}

// SetOrderUserID sets the "order_user_id" field.
func (fu *FractionUpdate) SetOrderUserID(u uuid.UUID) *FractionUpdate {
	fu.mutation.SetOrderUserID(u)
	return fu
}

// SetNillableOrderUserID sets the "order_user_id" field if the given value is not nil.
func (fu *FractionUpdate) SetNillableOrderUserID(u *uuid.UUID) *FractionUpdate {
	if u != nil {
		fu.SetOrderUserID(*u)
	}
	return fu
}

// ClearOrderUserID clears the value of the "order_user_id" field.
func (fu *FractionUpdate) ClearOrderUserID() *FractionUpdate {
	fu.mutation.ClearOrderUserID()
	return fu
}

// SetWithdrawState sets the "withdraw_state" field.
func (fu *FractionUpdate) SetWithdrawState(s string) *FractionUpdate {
	fu.mutation.SetWithdrawState(s)
	return fu
}

// SetNillableWithdrawState sets the "withdraw_state" field if the given value is not nil.
func (fu *FractionUpdate) SetNillableWithdrawState(s *string) *FractionUpdate {
	if s != nil {
		fu.SetWithdrawState(*s)
	}
	return fu
}

// ClearWithdrawState clears the value of the "withdraw_state" field.
func (fu *FractionUpdate) ClearWithdrawState() *FractionUpdate {
	fu.mutation.ClearWithdrawState()
	return fu
}

// SetWithdrawAt sets the "withdraw_at" field.
func (fu *FractionUpdate) SetWithdrawAt(u uint32) *FractionUpdate {
	fu.mutation.ResetWithdrawAt()
	fu.mutation.SetWithdrawAt(u)
	return fu
}

// SetNillableWithdrawAt sets the "withdraw_at" field if the given value is not nil.
func (fu *FractionUpdate) SetNillableWithdrawAt(u *uint32) *FractionUpdate {
	if u != nil {
		fu.SetWithdrawAt(*u)
	}
	return fu
}

// AddWithdrawAt adds u to the "withdraw_at" field.
func (fu *FractionUpdate) AddWithdrawAt(u int32) *FractionUpdate {
	fu.mutation.AddWithdrawAt(u)
	return fu
}

// ClearWithdrawAt clears the value of the "withdraw_at" field.
func (fu *FractionUpdate) ClearWithdrawAt() *FractionUpdate {
	fu.mutation.ClearWithdrawAt()
	return fu
}

// SetPromisePayAt sets the "promise_pay_at" field.
func (fu *FractionUpdate) SetPromisePayAt(u uint32) *FractionUpdate {
	fu.mutation.ResetPromisePayAt()
	fu.mutation.SetPromisePayAt(u)
	return fu
}

// SetNillablePromisePayAt sets the "promise_pay_at" field if the given value is not nil.
func (fu *FractionUpdate) SetNillablePromisePayAt(u *uint32) *FractionUpdate {
	if u != nil {
		fu.SetPromisePayAt(*u)
	}
	return fu
}

// AddPromisePayAt adds u to the "promise_pay_at" field.
func (fu *FractionUpdate) AddPromisePayAt(u int32) *FractionUpdate {
	fu.mutation.AddPromisePayAt(u)
	return fu
}

// ClearPromisePayAt clears the value of the "promise_pay_at" field.
func (fu *FractionUpdate) ClearPromisePayAt() *FractionUpdate {
	fu.mutation.ClearPromisePayAt()
	return fu
}

// SetMsg sets the "msg" field.
func (fu *FractionUpdate) SetMsg(s string) *FractionUpdate {
	fu.mutation.SetMsg(s)
	return fu
}

// SetNillableMsg sets the "msg" field if the given value is not nil.
func (fu *FractionUpdate) SetNillableMsg(s *string) *FractionUpdate {
	if s != nil {
		fu.SetMsg(*s)
	}
	return fu
}

// ClearMsg clears the value of the "msg" field.
func (fu *FractionUpdate) ClearMsg() *FractionUpdate {
	fu.mutation.ClearMsg()
	return fu
}

// Mutation returns the FractionMutation object of the builder.
func (fu *FractionUpdate) Mutation() *FractionMutation {
	return fu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FractionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := fu.defaults(); err != nil {
		return 0, err
	}
	if len(fu.hooks) == 0 {
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FractionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			if fu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FractionUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FractionUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FractionUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *FractionUpdate) defaults() error {
	if _, ok := fu.mutation.UpdatedAt(); !ok {
		if fraction.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized fraction.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := fraction.UpdateDefaultUpdatedAt()
		fu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fu *FractionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FractionUpdate {
	fu.modifiers = append(fu.modifiers, modifiers...)
	return fu
}

func (fu *FractionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fraction.Table,
			Columns: fraction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: fraction.FieldID,
			},
		},
	}
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldCreatedAt,
		})
	}
	if value, ok := fu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldCreatedAt,
		})
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldUpdatedAt,
		})
	}
	if value, ok := fu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldUpdatedAt,
		})
	}
	if value, ok := fu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldDeletedAt,
		})
	}
	if value, ok := fu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldDeletedAt,
		})
	}
	if value, ok := fu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fraction.FieldEntID,
		})
	}
	if value, ok := fu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fraction.FieldAppID,
		})
	}
	if fu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fraction.FieldAppID,
		})
	}
	if value, ok := fu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fraction.FieldUserID,
		})
	}
	if fu.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fraction.FieldUserID,
		})
	}
	if value, ok := fu.mutation.OrderUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fraction.FieldOrderUserID,
		})
	}
	if fu.mutation.OrderUserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fraction.FieldOrderUserID,
		})
	}
	if value, ok := fu.mutation.WithdrawState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fraction.FieldWithdrawState,
		})
	}
	if fu.mutation.WithdrawStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: fraction.FieldWithdrawState,
		})
	}
	if value, ok := fu.mutation.WithdrawAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldWithdrawAt,
		})
	}
	if value, ok := fu.mutation.AddedWithdrawAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldWithdrawAt,
		})
	}
	if fu.mutation.WithdrawAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: fraction.FieldWithdrawAt,
		})
	}
	if value, ok := fu.mutation.PromisePayAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldPromisePayAt,
		})
	}
	if value, ok := fu.mutation.AddedPromisePayAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldPromisePayAt,
		})
	}
	if fu.mutation.PromisePayAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: fraction.FieldPromisePayAt,
		})
	}
	if value, ok := fu.mutation.Msg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fraction.FieldMsg,
		})
	}
	if fu.mutation.MsgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: fraction.FieldMsg,
		})
	}
	_spec.Modifiers = fu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fraction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// FractionUpdateOne is the builder for updating a single Fraction entity.
type FractionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *FractionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (fuo *FractionUpdateOne) SetCreatedAt(u uint32) *FractionUpdateOne {
	fuo.mutation.ResetCreatedAt()
	fuo.mutation.SetCreatedAt(u)
	return fuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fuo *FractionUpdateOne) SetNillableCreatedAt(u *uint32) *FractionUpdateOne {
	if u != nil {
		fuo.SetCreatedAt(*u)
	}
	return fuo
}

// AddCreatedAt adds u to the "created_at" field.
func (fuo *FractionUpdateOne) AddCreatedAt(u int32) *FractionUpdateOne {
	fuo.mutation.AddCreatedAt(u)
	return fuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FractionUpdateOne) SetUpdatedAt(u uint32) *FractionUpdateOne {
	fuo.mutation.ResetUpdatedAt()
	fuo.mutation.SetUpdatedAt(u)
	return fuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (fuo *FractionUpdateOne) AddUpdatedAt(u int32) *FractionUpdateOne {
	fuo.mutation.AddUpdatedAt(u)
	return fuo
}

// SetDeletedAt sets the "deleted_at" field.
func (fuo *FractionUpdateOne) SetDeletedAt(u uint32) *FractionUpdateOne {
	fuo.mutation.ResetDeletedAt()
	fuo.mutation.SetDeletedAt(u)
	return fuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fuo *FractionUpdateOne) SetNillableDeletedAt(u *uint32) *FractionUpdateOne {
	if u != nil {
		fuo.SetDeletedAt(*u)
	}
	return fuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (fuo *FractionUpdateOne) AddDeletedAt(u int32) *FractionUpdateOne {
	fuo.mutation.AddDeletedAt(u)
	return fuo
}

// SetEntID sets the "ent_id" field.
func (fuo *FractionUpdateOne) SetEntID(u uuid.UUID) *FractionUpdateOne {
	fuo.mutation.SetEntID(u)
	return fuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (fuo *FractionUpdateOne) SetNillableEntID(u *uuid.UUID) *FractionUpdateOne {
	if u != nil {
		fuo.SetEntID(*u)
	}
	return fuo
}

// SetAppID sets the "app_id" field.
func (fuo *FractionUpdateOne) SetAppID(u uuid.UUID) *FractionUpdateOne {
	fuo.mutation.SetAppID(u)
	return fuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (fuo *FractionUpdateOne) SetNillableAppID(u *uuid.UUID) *FractionUpdateOne {
	if u != nil {
		fuo.SetAppID(*u)
	}
	return fuo
}

// ClearAppID clears the value of the "app_id" field.
func (fuo *FractionUpdateOne) ClearAppID() *FractionUpdateOne {
	fuo.mutation.ClearAppID()
	return fuo
}

// SetUserID sets the "user_id" field.
func (fuo *FractionUpdateOne) SetUserID(u uuid.UUID) *FractionUpdateOne {
	fuo.mutation.SetUserID(u)
	return fuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (fuo *FractionUpdateOne) SetNillableUserID(u *uuid.UUID) *FractionUpdateOne {
	if u != nil {
		fuo.SetUserID(*u)
	}
	return fuo
}

// ClearUserID clears the value of the "user_id" field.
func (fuo *FractionUpdateOne) ClearUserID() *FractionUpdateOne {
	fuo.mutation.ClearUserID()
	return fuo
}

// SetOrderUserID sets the "order_user_id" field.
func (fuo *FractionUpdateOne) SetOrderUserID(u uuid.UUID) *FractionUpdateOne {
	fuo.mutation.SetOrderUserID(u)
	return fuo
}

// SetNillableOrderUserID sets the "order_user_id" field if the given value is not nil.
func (fuo *FractionUpdateOne) SetNillableOrderUserID(u *uuid.UUID) *FractionUpdateOne {
	if u != nil {
		fuo.SetOrderUserID(*u)
	}
	return fuo
}

// ClearOrderUserID clears the value of the "order_user_id" field.
func (fuo *FractionUpdateOne) ClearOrderUserID() *FractionUpdateOne {
	fuo.mutation.ClearOrderUserID()
	return fuo
}

// SetWithdrawState sets the "withdraw_state" field.
func (fuo *FractionUpdateOne) SetWithdrawState(s string) *FractionUpdateOne {
	fuo.mutation.SetWithdrawState(s)
	return fuo
}

// SetNillableWithdrawState sets the "withdraw_state" field if the given value is not nil.
func (fuo *FractionUpdateOne) SetNillableWithdrawState(s *string) *FractionUpdateOne {
	if s != nil {
		fuo.SetWithdrawState(*s)
	}
	return fuo
}

// ClearWithdrawState clears the value of the "withdraw_state" field.
func (fuo *FractionUpdateOne) ClearWithdrawState() *FractionUpdateOne {
	fuo.mutation.ClearWithdrawState()
	return fuo
}

// SetWithdrawAt sets the "withdraw_at" field.
func (fuo *FractionUpdateOne) SetWithdrawAt(u uint32) *FractionUpdateOne {
	fuo.mutation.ResetWithdrawAt()
	fuo.mutation.SetWithdrawAt(u)
	return fuo
}

// SetNillableWithdrawAt sets the "withdraw_at" field if the given value is not nil.
func (fuo *FractionUpdateOne) SetNillableWithdrawAt(u *uint32) *FractionUpdateOne {
	if u != nil {
		fuo.SetWithdrawAt(*u)
	}
	return fuo
}

// AddWithdrawAt adds u to the "withdraw_at" field.
func (fuo *FractionUpdateOne) AddWithdrawAt(u int32) *FractionUpdateOne {
	fuo.mutation.AddWithdrawAt(u)
	return fuo
}

// ClearWithdrawAt clears the value of the "withdraw_at" field.
func (fuo *FractionUpdateOne) ClearWithdrawAt() *FractionUpdateOne {
	fuo.mutation.ClearWithdrawAt()
	return fuo
}

// SetPromisePayAt sets the "promise_pay_at" field.
func (fuo *FractionUpdateOne) SetPromisePayAt(u uint32) *FractionUpdateOne {
	fuo.mutation.ResetPromisePayAt()
	fuo.mutation.SetPromisePayAt(u)
	return fuo
}

// SetNillablePromisePayAt sets the "promise_pay_at" field if the given value is not nil.
func (fuo *FractionUpdateOne) SetNillablePromisePayAt(u *uint32) *FractionUpdateOne {
	if u != nil {
		fuo.SetPromisePayAt(*u)
	}
	return fuo
}

// AddPromisePayAt adds u to the "promise_pay_at" field.
func (fuo *FractionUpdateOne) AddPromisePayAt(u int32) *FractionUpdateOne {
	fuo.mutation.AddPromisePayAt(u)
	return fuo
}

// ClearPromisePayAt clears the value of the "promise_pay_at" field.
func (fuo *FractionUpdateOne) ClearPromisePayAt() *FractionUpdateOne {
	fuo.mutation.ClearPromisePayAt()
	return fuo
}

// SetMsg sets the "msg" field.
func (fuo *FractionUpdateOne) SetMsg(s string) *FractionUpdateOne {
	fuo.mutation.SetMsg(s)
	return fuo
}

// SetNillableMsg sets the "msg" field if the given value is not nil.
func (fuo *FractionUpdateOne) SetNillableMsg(s *string) *FractionUpdateOne {
	if s != nil {
		fuo.SetMsg(*s)
	}
	return fuo
}

// ClearMsg clears the value of the "msg" field.
func (fuo *FractionUpdateOne) ClearMsg() *FractionUpdateOne {
	fuo.mutation.ClearMsg()
	return fuo
}

// Mutation returns the FractionMutation object of the builder.
func (fuo *FractionUpdateOne) Mutation() *FractionMutation {
	return fuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FractionUpdateOne) Select(field string, fields ...string) *FractionUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Fraction entity.
func (fuo *FractionUpdateOne) Save(ctx context.Context) (*Fraction, error) {
	var (
		err  error
		node *Fraction
	)
	if err := fuo.defaults(); err != nil {
		return nil, err
	}
	if len(fuo.hooks) == 0 {
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FractionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			if fuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Fraction)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FractionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FractionUpdateOne) SaveX(ctx context.Context) *Fraction {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FractionUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FractionUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *FractionUpdateOne) defaults() error {
	if _, ok := fuo.mutation.UpdatedAt(); !ok {
		if fraction.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized fraction.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := fraction.UpdateDefaultUpdatedAt()
		fuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fuo *FractionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FractionUpdateOne {
	fuo.modifiers = append(fuo.modifiers, modifiers...)
	return fuo
}

func (fuo *FractionUpdateOne) sqlSave(ctx context.Context) (_node *Fraction, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fraction.Table,
			Columns: fraction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: fraction.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Fraction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fraction.FieldID)
		for _, f := range fields {
			if !fraction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != fraction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldCreatedAt,
		})
	}
	if value, ok := fuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldCreatedAt,
		})
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldUpdatedAt,
		})
	}
	if value, ok := fuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldUpdatedAt,
		})
	}
	if value, ok := fuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldDeletedAt,
		})
	}
	if value, ok := fuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldDeletedAt,
		})
	}
	if value, ok := fuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fraction.FieldEntID,
		})
	}
	if value, ok := fuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fraction.FieldAppID,
		})
	}
	if fuo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fraction.FieldAppID,
		})
	}
	if value, ok := fuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fraction.FieldUserID,
		})
	}
	if fuo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fraction.FieldUserID,
		})
	}
	if value, ok := fuo.mutation.OrderUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fraction.FieldOrderUserID,
		})
	}
	if fuo.mutation.OrderUserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fraction.FieldOrderUserID,
		})
	}
	if value, ok := fuo.mutation.WithdrawState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fraction.FieldWithdrawState,
		})
	}
	if fuo.mutation.WithdrawStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: fraction.FieldWithdrawState,
		})
	}
	if value, ok := fuo.mutation.WithdrawAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldWithdrawAt,
		})
	}
	if value, ok := fuo.mutation.AddedWithdrawAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldWithdrawAt,
		})
	}
	if fuo.mutation.WithdrawAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: fraction.FieldWithdrawAt,
		})
	}
	if value, ok := fuo.mutation.PromisePayAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldPromisePayAt,
		})
	}
	if value, ok := fuo.mutation.AddedPromisePayAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fraction.FieldPromisePayAt,
		})
	}
	if fuo.mutation.PromisePayAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: fraction.FieldPromisePayAt,
		})
	}
	if value, ok := fuo.mutation.Msg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fraction.FieldMsg,
		})
	}
	if fuo.mutation.MsgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: fraction.FieldMsg,
		})
	}
	_spec.Modifiers = fuo.modifiers
	_node = &Fraction{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fraction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
