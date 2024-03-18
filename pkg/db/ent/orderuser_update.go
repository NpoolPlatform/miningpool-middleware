// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// OrderUserUpdate is the builder for updating OrderUser entities.
type OrderUserUpdate struct {
	config
	hooks     []Hook
	mutation  *OrderUserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrderUserUpdate builder.
func (ouu *OrderUserUpdate) Where(ps ...predicate.OrderUser) *OrderUserUpdate {
	ouu.mutation.Where(ps...)
	return ouu
}

// SetCreatedAt sets the "created_at" field.
func (ouu *OrderUserUpdate) SetCreatedAt(u uint32) *OrderUserUpdate {
	ouu.mutation.ResetCreatedAt()
	ouu.mutation.SetCreatedAt(u)
	return ouu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ouu *OrderUserUpdate) SetNillableCreatedAt(u *uint32) *OrderUserUpdate {
	if u != nil {
		ouu.SetCreatedAt(*u)
	}
	return ouu
}

// AddCreatedAt adds u to the "created_at" field.
func (ouu *OrderUserUpdate) AddCreatedAt(u int32) *OrderUserUpdate {
	ouu.mutation.AddCreatedAt(u)
	return ouu
}

// SetUpdatedAt sets the "updated_at" field.
func (ouu *OrderUserUpdate) SetUpdatedAt(u uint32) *OrderUserUpdate {
	ouu.mutation.ResetUpdatedAt()
	ouu.mutation.SetUpdatedAt(u)
	return ouu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ouu *OrderUserUpdate) AddUpdatedAt(u int32) *OrderUserUpdate {
	ouu.mutation.AddUpdatedAt(u)
	return ouu
}

// SetDeletedAt sets the "deleted_at" field.
func (ouu *OrderUserUpdate) SetDeletedAt(u uint32) *OrderUserUpdate {
	ouu.mutation.ResetDeletedAt()
	ouu.mutation.SetDeletedAt(u)
	return ouu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ouu *OrderUserUpdate) SetNillableDeletedAt(u *uint32) *OrderUserUpdate {
	if u != nil {
		ouu.SetDeletedAt(*u)
	}
	return ouu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ouu *OrderUserUpdate) AddDeletedAt(u int32) *OrderUserUpdate {
	ouu.mutation.AddDeletedAt(u)
	return ouu
}

// SetEntID sets the "ent_id" field.
func (ouu *OrderUserUpdate) SetEntID(u uuid.UUID) *OrderUserUpdate {
	ouu.mutation.SetEntID(u)
	return ouu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ouu *OrderUserUpdate) SetNillableEntID(u *uuid.UUID) *OrderUserUpdate {
	if u != nil {
		ouu.SetEntID(*u)
	}
	return ouu
}

// SetRootUserID sets the "root_user_id" field.
func (ouu *OrderUserUpdate) SetRootUserID(u uuid.UUID) *OrderUserUpdate {
	ouu.mutation.SetRootUserID(u)
	return ouu
}

// SetGoodUserID sets the "good_user_id" field.
func (ouu *OrderUserUpdate) SetGoodUserID(u uuid.UUID) *OrderUserUpdate {
	ouu.mutation.SetGoodUserID(u)
	return ouu
}

// SetOrderID sets the "order_id" field.
func (ouu *OrderUserUpdate) SetOrderID(u uuid.UUID) *OrderUserUpdate {
	ouu.mutation.SetOrderID(u)
	return ouu
}

// SetName sets the "name" field.
func (ouu *OrderUserUpdate) SetName(s string) *OrderUserUpdate {
	ouu.mutation.SetName(s)
	return ouu
}

// SetMiningpoolType sets the "miningpool_type" field.
func (ouu *OrderUserUpdate) SetMiningpoolType(s string) *OrderUserUpdate {
	ouu.mutation.SetMiningpoolType(s)
	return ouu
}

// SetCoinType sets the "coin_type" field.
func (ouu *OrderUserUpdate) SetCoinType(s string) *OrderUserUpdate {
	ouu.mutation.SetCoinType(s)
	return ouu
}

// SetProportion sets the "proportion" field.
func (ouu *OrderUserUpdate) SetProportion(f float32) *OrderUserUpdate {
	ouu.mutation.ResetProportion()
	ouu.mutation.SetProportion(f)
	return ouu
}

// SetNillableProportion sets the "proportion" field if the given value is not nil.
func (ouu *OrderUserUpdate) SetNillableProportion(f *float32) *OrderUserUpdate {
	if f != nil {
		ouu.SetProportion(*f)
	}
	return ouu
}

// AddProportion adds f to the "proportion" field.
func (ouu *OrderUserUpdate) AddProportion(f float32) *OrderUserUpdate {
	ouu.mutation.AddProportion(f)
	return ouu
}

// ClearProportion clears the value of the "proportion" field.
func (ouu *OrderUserUpdate) ClearProportion() *OrderUserUpdate {
	ouu.mutation.ClearProportion()
	return ouu
}

// SetRevenueAddress sets the "revenue_address" field.
func (ouu *OrderUserUpdate) SetRevenueAddress(s string) *OrderUserUpdate {
	ouu.mutation.SetRevenueAddress(s)
	return ouu
}

// SetNillableRevenueAddress sets the "revenue_address" field if the given value is not nil.
func (ouu *OrderUserUpdate) SetNillableRevenueAddress(s *string) *OrderUserUpdate {
	if s != nil {
		ouu.SetRevenueAddress(*s)
	}
	return ouu
}

// ClearRevenueAddress clears the value of the "revenue_address" field.
func (ouu *OrderUserUpdate) ClearRevenueAddress() *OrderUserUpdate {
	ouu.mutation.ClearRevenueAddress()
	return ouu
}

// SetReadPageLink sets the "read_page_link" field.
func (ouu *OrderUserUpdate) SetReadPageLink(s string) *OrderUserUpdate {
	ouu.mutation.SetReadPageLink(s)
	return ouu
}

// SetAutoPay sets the "auto_pay" field.
func (ouu *OrderUserUpdate) SetAutoPay(b bool) *OrderUserUpdate {
	ouu.mutation.SetAutoPay(b)
	return ouu
}

// SetNillableAutoPay sets the "auto_pay" field if the given value is not nil.
func (ouu *OrderUserUpdate) SetNillableAutoPay(b *bool) *OrderUserUpdate {
	if b != nil {
		ouu.SetAutoPay(*b)
	}
	return ouu
}

// ClearAutoPay clears the value of the "auto_pay" field.
func (ouu *OrderUserUpdate) ClearAutoPay() *OrderUserUpdate {
	ouu.mutation.ClearAutoPay()
	return ouu
}

// Mutation returns the OrderUserMutation object of the builder.
func (ouu *OrderUserUpdate) Mutation() *OrderUserMutation {
	return ouu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ouu *OrderUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ouu.defaults(); err != nil {
		return 0, err
	}
	if len(ouu.hooks) == 0 {
		affected, err = ouu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ouu.mutation = mutation
			affected, err = ouu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ouu.hooks) - 1; i >= 0; i-- {
			if ouu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouu *OrderUserUpdate) SaveX(ctx context.Context) int {
	affected, err := ouu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ouu *OrderUserUpdate) Exec(ctx context.Context) error {
	_, err := ouu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouu *OrderUserUpdate) ExecX(ctx context.Context) {
	if err := ouu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouu *OrderUserUpdate) defaults() error {
	if _, ok := ouu.mutation.UpdatedAt(); !ok {
		if orderuser.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderuser.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderuser.UpdateDefaultUpdatedAt()
		ouu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ouu *OrderUserUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderUserUpdate {
	ouu.modifiers = append(ouu.modifiers, modifiers...)
	return ouu
}

func (ouu *OrderUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderuser.Table,
			Columns: orderuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderuser.FieldID,
			},
		},
	}
	if ps := ouu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderuser.FieldCreatedAt,
		})
	}
	if value, ok := ouu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderuser.FieldCreatedAt,
		})
	}
	if value, ok := ouu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderuser.FieldUpdatedAt,
		})
	}
	if value, ok := ouu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderuser.FieldUpdatedAt,
		})
	}
	if value, ok := ouu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderuser.FieldDeletedAt,
		})
	}
	if value, ok := ouu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderuser.FieldDeletedAt,
		})
	}
	if value, ok := ouu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderuser.FieldEntID,
		})
	}
	if value, ok := ouu.mutation.RootUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderuser.FieldRootUserID,
		})
	}
	if value, ok := ouu.mutation.GoodUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderuser.FieldGoodUserID,
		})
	}
	if value, ok := ouu.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderuser.FieldOrderID,
		})
	}
	if value, ok := ouu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderuser.FieldName,
		})
	}
	if value, ok := ouu.mutation.MiningpoolType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderuser.FieldMiningpoolType,
		})
	}
	if value, ok := ouu.mutation.CoinType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderuser.FieldCoinType,
		})
	}
	if value, ok := ouu.mutation.Proportion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: orderuser.FieldProportion,
		})
	}
	if value, ok := ouu.mutation.AddedProportion(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: orderuser.FieldProportion,
		})
	}
	if ouu.mutation.ProportionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Column: orderuser.FieldProportion,
		})
	}
	if value, ok := ouu.mutation.RevenueAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderuser.FieldRevenueAddress,
		})
	}
	if ouu.mutation.RevenueAddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderuser.FieldRevenueAddress,
		})
	}
	if value, ok := ouu.mutation.ReadPageLink(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderuser.FieldReadPageLink,
		})
	}
	if value, ok := ouu.mutation.AutoPay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: orderuser.FieldAutoPay,
		})
	}
	if ouu.mutation.AutoPayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: orderuser.FieldAutoPay,
		})
	}
	_spec.Modifiers = ouu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ouu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrderUserUpdateOne is the builder for updating a single OrderUser entity.
type OrderUserUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrderUserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (ouuo *OrderUserUpdateOne) SetCreatedAt(u uint32) *OrderUserUpdateOne {
	ouuo.mutation.ResetCreatedAt()
	ouuo.mutation.SetCreatedAt(u)
	return ouuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ouuo *OrderUserUpdateOne) SetNillableCreatedAt(u *uint32) *OrderUserUpdateOne {
	if u != nil {
		ouuo.SetCreatedAt(*u)
	}
	return ouuo
}

// AddCreatedAt adds u to the "created_at" field.
func (ouuo *OrderUserUpdateOne) AddCreatedAt(u int32) *OrderUserUpdateOne {
	ouuo.mutation.AddCreatedAt(u)
	return ouuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ouuo *OrderUserUpdateOne) SetUpdatedAt(u uint32) *OrderUserUpdateOne {
	ouuo.mutation.ResetUpdatedAt()
	ouuo.mutation.SetUpdatedAt(u)
	return ouuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ouuo *OrderUserUpdateOne) AddUpdatedAt(u int32) *OrderUserUpdateOne {
	ouuo.mutation.AddUpdatedAt(u)
	return ouuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ouuo *OrderUserUpdateOne) SetDeletedAt(u uint32) *OrderUserUpdateOne {
	ouuo.mutation.ResetDeletedAt()
	ouuo.mutation.SetDeletedAt(u)
	return ouuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ouuo *OrderUserUpdateOne) SetNillableDeletedAt(u *uint32) *OrderUserUpdateOne {
	if u != nil {
		ouuo.SetDeletedAt(*u)
	}
	return ouuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ouuo *OrderUserUpdateOne) AddDeletedAt(u int32) *OrderUserUpdateOne {
	ouuo.mutation.AddDeletedAt(u)
	return ouuo
}

// SetEntID sets the "ent_id" field.
func (ouuo *OrderUserUpdateOne) SetEntID(u uuid.UUID) *OrderUserUpdateOne {
	ouuo.mutation.SetEntID(u)
	return ouuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ouuo *OrderUserUpdateOne) SetNillableEntID(u *uuid.UUID) *OrderUserUpdateOne {
	if u != nil {
		ouuo.SetEntID(*u)
	}
	return ouuo
}

// SetRootUserID sets the "root_user_id" field.
func (ouuo *OrderUserUpdateOne) SetRootUserID(u uuid.UUID) *OrderUserUpdateOne {
	ouuo.mutation.SetRootUserID(u)
	return ouuo
}

// SetGoodUserID sets the "good_user_id" field.
func (ouuo *OrderUserUpdateOne) SetGoodUserID(u uuid.UUID) *OrderUserUpdateOne {
	ouuo.mutation.SetGoodUserID(u)
	return ouuo
}

// SetOrderID sets the "order_id" field.
func (ouuo *OrderUserUpdateOne) SetOrderID(u uuid.UUID) *OrderUserUpdateOne {
	ouuo.mutation.SetOrderID(u)
	return ouuo
}

// SetName sets the "name" field.
func (ouuo *OrderUserUpdateOne) SetName(s string) *OrderUserUpdateOne {
	ouuo.mutation.SetName(s)
	return ouuo
}

// SetMiningpoolType sets the "miningpool_type" field.
func (ouuo *OrderUserUpdateOne) SetMiningpoolType(s string) *OrderUserUpdateOne {
	ouuo.mutation.SetMiningpoolType(s)
	return ouuo
}

// SetCoinType sets the "coin_type" field.
func (ouuo *OrderUserUpdateOne) SetCoinType(s string) *OrderUserUpdateOne {
	ouuo.mutation.SetCoinType(s)
	return ouuo
}

// SetProportion sets the "proportion" field.
func (ouuo *OrderUserUpdateOne) SetProportion(f float32) *OrderUserUpdateOne {
	ouuo.mutation.ResetProportion()
	ouuo.mutation.SetProportion(f)
	return ouuo
}

// SetNillableProportion sets the "proportion" field if the given value is not nil.
func (ouuo *OrderUserUpdateOne) SetNillableProportion(f *float32) *OrderUserUpdateOne {
	if f != nil {
		ouuo.SetProportion(*f)
	}
	return ouuo
}

// AddProportion adds f to the "proportion" field.
func (ouuo *OrderUserUpdateOne) AddProportion(f float32) *OrderUserUpdateOne {
	ouuo.mutation.AddProportion(f)
	return ouuo
}

// ClearProportion clears the value of the "proportion" field.
func (ouuo *OrderUserUpdateOne) ClearProportion() *OrderUserUpdateOne {
	ouuo.mutation.ClearProportion()
	return ouuo
}

// SetRevenueAddress sets the "revenue_address" field.
func (ouuo *OrderUserUpdateOne) SetRevenueAddress(s string) *OrderUserUpdateOne {
	ouuo.mutation.SetRevenueAddress(s)
	return ouuo
}

// SetNillableRevenueAddress sets the "revenue_address" field if the given value is not nil.
func (ouuo *OrderUserUpdateOne) SetNillableRevenueAddress(s *string) *OrderUserUpdateOne {
	if s != nil {
		ouuo.SetRevenueAddress(*s)
	}
	return ouuo
}

// ClearRevenueAddress clears the value of the "revenue_address" field.
func (ouuo *OrderUserUpdateOne) ClearRevenueAddress() *OrderUserUpdateOne {
	ouuo.mutation.ClearRevenueAddress()
	return ouuo
}

// SetReadPageLink sets the "read_page_link" field.
func (ouuo *OrderUserUpdateOne) SetReadPageLink(s string) *OrderUserUpdateOne {
	ouuo.mutation.SetReadPageLink(s)
	return ouuo
}

// SetAutoPay sets the "auto_pay" field.
func (ouuo *OrderUserUpdateOne) SetAutoPay(b bool) *OrderUserUpdateOne {
	ouuo.mutation.SetAutoPay(b)
	return ouuo
}

// SetNillableAutoPay sets the "auto_pay" field if the given value is not nil.
func (ouuo *OrderUserUpdateOne) SetNillableAutoPay(b *bool) *OrderUserUpdateOne {
	if b != nil {
		ouuo.SetAutoPay(*b)
	}
	return ouuo
}

// ClearAutoPay clears the value of the "auto_pay" field.
func (ouuo *OrderUserUpdateOne) ClearAutoPay() *OrderUserUpdateOne {
	ouuo.mutation.ClearAutoPay()
	return ouuo
}

// Mutation returns the OrderUserMutation object of the builder.
func (ouuo *OrderUserUpdateOne) Mutation() *OrderUserMutation {
	return ouuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouuo *OrderUserUpdateOne) Select(field string, fields ...string) *OrderUserUpdateOne {
	ouuo.fields = append([]string{field}, fields...)
	return ouuo
}

// Save executes the query and returns the updated OrderUser entity.
func (ouuo *OrderUserUpdateOne) Save(ctx context.Context) (*OrderUser, error) {
	var (
		err  error
		node *OrderUser
	)
	if err := ouuo.defaults(); err != nil {
		return nil, err
	}
	if len(ouuo.hooks) == 0 {
		node, err = ouuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ouuo.mutation = mutation
			node, err = ouuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouuo.hooks) - 1; i >= 0; i-- {
			if ouuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ouuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrderUser)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderUserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouuo *OrderUserUpdateOne) SaveX(ctx context.Context) *OrderUser {
	node, err := ouuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouuo *OrderUserUpdateOne) Exec(ctx context.Context) error {
	_, err := ouuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouuo *OrderUserUpdateOne) ExecX(ctx context.Context) {
	if err := ouuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouuo *OrderUserUpdateOne) defaults() error {
	if _, ok := ouuo.mutation.UpdatedAt(); !ok {
		if orderuser.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderuser.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderuser.UpdateDefaultUpdatedAt()
		ouuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ouuo *OrderUserUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderUserUpdateOne {
	ouuo.modifiers = append(ouuo.modifiers, modifiers...)
	return ouuo
}

func (ouuo *OrderUserUpdateOne) sqlSave(ctx context.Context) (_node *OrderUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderuser.Table,
			Columns: orderuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderuser.FieldID,
			},
		},
	}
	id, ok := ouuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderuser.FieldID)
		for _, f := range fields {
			if !orderuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderuser.FieldCreatedAt,
		})
	}
	if value, ok := ouuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderuser.FieldCreatedAt,
		})
	}
	if value, ok := ouuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderuser.FieldUpdatedAt,
		})
	}
	if value, ok := ouuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderuser.FieldUpdatedAt,
		})
	}
	if value, ok := ouuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderuser.FieldDeletedAt,
		})
	}
	if value, ok := ouuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderuser.FieldDeletedAt,
		})
	}
	if value, ok := ouuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderuser.FieldEntID,
		})
	}
	if value, ok := ouuo.mutation.RootUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderuser.FieldRootUserID,
		})
	}
	if value, ok := ouuo.mutation.GoodUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderuser.FieldGoodUserID,
		})
	}
	if value, ok := ouuo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderuser.FieldOrderID,
		})
	}
	if value, ok := ouuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderuser.FieldName,
		})
	}
	if value, ok := ouuo.mutation.MiningpoolType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderuser.FieldMiningpoolType,
		})
	}
	if value, ok := ouuo.mutation.CoinType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderuser.FieldCoinType,
		})
	}
	if value, ok := ouuo.mutation.Proportion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: orderuser.FieldProportion,
		})
	}
	if value, ok := ouuo.mutation.AddedProportion(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: orderuser.FieldProportion,
		})
	}
	if ouuo.mutation.ProportionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Column: orderuser.FieldProportion,
		})
	}
	if value, ok := ouuo.mutation.RevenueAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderuser.FieldRevenueAddress,
		})
	}
	if ouuo.mutation.RevenueAddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderuser.FieldRevenueAddress,
		})
	}
	if value, ok := ouuo.mutation.ReadPageLink(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderuser.FieldReadPageLink,
		})
	}
	if value, ok := ouuo.mutation.AutoPay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: orderuser.FieldAutoPay,
		})
	}
	if ouuo.mutation.AutoPayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: orderuser.FieldAutoPay,
		})
	}
	_spec.Modifiers = ouuo.modifiers
	_node = &OrderUser{config: ouuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
