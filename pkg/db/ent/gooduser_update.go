// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/gooduser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GoodUserUpdate is the builder for updating GoodUser entities.
type GoodUserUpdate struct {
	config
	hooks     []Hook
	mutation  *GoodUserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the GoodUserUpdate builder.
func (guu *GoodUserUpdate) Where(ps ...predicate.GoodUser) *GoodUserUpdate {
	guu.mutation.Where(ps...)
	return guu
}

// SetCreatedAt sets the "created_at" field.
func (guu *GoodUserUpdate) SetCreatedAt(u uint32) *GoodUserUpdate {
	guu.mutation.ResetCreatedAt()
	guu.mutation.SetCreatedAt(u)
	return guu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (guu *GoodUserUpdate) SetNillableCreatedAt(u *uint32) *GoodUserUpdate {
	if u != nil {
		guu.SetCreatedAt(*u)
	}
	return guu
}

// AddCreatedAt adds u to the "created_at" field.
func (guu *GoodUserUpdate) AddCreatedAt(u int32) *GoodUserUpdate {
	guu.mutation.AddCreatedAt(u)
	return guu
}

// SetUpdatedAt sets the "updated_at" field.
func (guu *GoodUserUpdate) SetUpdatedAt(u uint32) *GoodUserUpdate {
	guu.mutation.ResetUpdatedAt()
	guu.mutation.SetUpdatedAt(u)
	return guu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (guu *GoodUserUpdate) AddUpdatedAt(u int32) *GoodUserUpdate {
	guu.mutation.AddUpdatedAt(u)
	return guu
}

// SetDeletedAt sets the "deleted_at" field.
func (guu *GoodUserUpdate) SetDeletedAt(u uint32) *GoodUserUpdate {
	guu.mutation.ResetDeletedAt()
	guu.mutation.SetDeletedAt(u)
	return guu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (guu *GoodUserUpdate) SetNillableDeletedAt(u *uint32) *GoodUserUpdate {
	if u != nil {
		guu.SetDeletedAt(*u)
	}
	return guu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (guu *GoodUserUpdate) AddDeletedAt(u int32) *GoodUserUpdate {
	guu.mutation.AddDeletedAt(u)
	return guu
}

// SetEntID sets the "ent_id" field.
func (guu *GoodUserUpdate) SetEntID(u uuid.UUID) *GoodUserUpdate {
	guu.mutation.SetEntID(u)
	return guu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (guu *GoodUserUpdate) SetNillableEntID(u *uuid.UUID) *GoodUserUpdate {
	if u != nil {
		guu.SetEntID(*u)
	}
	return guu
}

// SetRootUserID sets the "root_user_id" field.
func (guu *GoodUserUpdate) SetRootUserID(u uuid.UUID) *GoodUserUpdate {
	guu.mutation.SetRootUserID(u)
	return guu
}

// SetNillableRootUserID sets the "root_user_id" field if the given value is not nil.
func (guu *GoodUserUpdate) SetNillableRootUserID(u *uuid.UUID) *GoodUserUpdate {
	if u != nil {
		guu.SetRootUserID(*u)
	}
	return guu
}

// ClearRootUserID clears the value of the "root_user_id" field.
func (guu *GoodUserUpdate) ClearRootUserID() *GoodUserUpdate {
	guu.mutation.ClearRootUserID()
	return guu
}

// SetName sets the "name" field.
func (guu *GoodUserUpdate) SetName(s string) *GoodUserUpdate {
	guu.mutation.SetName(s)
	return guu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (guu *GoodUserUpdate) SetNillableName(s *string) *GoodUserUpdate {
	if s != nil {
		guu.SetName(*s)
	}
	return guu
}

// ClearName clears the value of the "name" field.
func (guu *GoodUserUpdate) ClearName() *GoodUserUpdate {
	guu.mutation.ClearName()
	return guu
}

// SetCoinID sets the "coin_id" field.
func (guu *GoodUserUpdate) SetCoinID(u uuid.UUID) *GoodUserUpdate {
	guu.mutation.SetCoinID(u)
	return guu
}

// SetNillableCoinID sets the "coin_id" field if the given value is not nil.
func (guu *GoodUserUpdate) SetNillableCoinID(u *uuid.UUID) *GoodUserUpdate {
	if u != nil {
		guu.SetCoinID(*u)
	}
	return guu
}

// ClearCoinID clears the value of the "coin_id" field.
func (guu *GoodUserUpdate) ClearCoinID() *GoodUserUpdate {
	guu.mutation.ClearCoinID()
	return guu
}

// SetRevenueID sets the "revenue_id" field.
func (guu *GoodUserUpdate) SetRevenueID(u uuid.UUID) *GoodUserUpdate {
	guu.mutation.SetRevenueID(u)
	return guu
}

// SetNillableRevenueID sets the "revenue_id" field if the given value is not nil.
func (guu *GoodUserUpdate) SetNillableRevenueID(u *uuid.UUID) *GoodUserUpdate {
	if u != nil {
		guu.SetRevenueID(*u)
	}
	return guu
}

// ClearRevenueID clears the value of the "revenue_id" field.
func (guu *GoodUserUpdate) ClearRevenueID() *GoodUserUpdate {
	guu.mutation.ClearRevenueID()
	return guu
}

// SetHashRate sets the "hash_rate" field.
func (guu *GoodUserUpdate) SetHashRate(f float32) *GoodUserUpdate {
	guu.mutation.ResetHashRate()
	guu.mutation.SetHashRate(f)
	return guu
}

// SetNillableHashRate sets the "hash_rate" field if the given value is not nil.
func (guu *GoodUserUpdate) SetNillableHashRate(f *float32) *GoodUserUpdate {
	if f != nil {
		guu.SetHashRate(*f)
	}
	return guu
}

// AddHashRate adds f to the "hash_rate" field.
func (guu *GoodUserUpdate) AddHashRate(f float32) *GoodUserUpdate {
	guu.mutation.AddHashRate(f)
	return guu
}

// ClearHashRate clears the value of the "hash_rate" field.
func (guu *GoodUserUpdate) ClearHashRate() *GoodUserUpdate {
	guu.mutation.ClearHashRate()
	return guu
}

// SetReadPageLink sets the "read_page_link" field.
func (guu *GoodUserUpdate) SetReadPageLink(s string) *GoodUserUpdate {
	guu.mutation.SetReadPageLink(s)
	return guu
}

// SetNillableReadPageLink sets the "read_page_link" field if the given value is not nil.
func (guu *GoodUserUpdate) SetNillableReadPageLink(s *string) *GoodUserUpdate {
	if s != nil {
		guu.SetReadPageLink(*s)
	}
	return guu
}

// ClearReadPageLink clears the value of the "read_page_link" field.
func (guu *GoodUserUpdate) ClearReadPageLink() *GoodUserUpdate {
	guu.mutation.ClearReadPageLink()
	return guu
}

// Mutation returns the GoodUserMutation object of the builder.
func (guu *GoodUserUpdate) Mutation() *GoodUserMutation {
	return guu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (guu *GoodUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := guu.defaults(); err != nil {
		return 0, err
	}
	if len(guu.hooks) == 0 {
		affected, err = guu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			guu.mutation = mutation
			affected, err = guu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(guu.hooks) - 1; i >= 0; i-- {
			if guu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (guu *GoodUserUpdate) SaveX(ctx context.Context) int {
	affected, err := guu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (guu *GoodUserUpdate) Exec(ctx context.Context) error {
	_, err := guu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guu *GoodUserUpdate) ExecX(ctx context.Context) {
	if err := guu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guu *GoodUserUpdate) defaults() error {
	if _, ok := guu.mutation.UpdatedAt(); !ok {
		if gooduser.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized gooduser.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := gooduser.UpdateDefaultUpdatedAt()
		guu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (guu *GoodUserUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GoodUserUpdate {
	guu.modifiers = append(guu.modifiers, modifiers...)
	return guu
}

func (guu *GoodUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gooduser.Table,
			Columns: gooduser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: gooduser.FieldID,
			},
		},
	}
	if ps := guu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gooduser.FieldCreatedAt,
		})
	}
	if value, ok := guu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gooduser.FieldCreatedAt,
		})
	}
	if value, ok := guu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gooduser.FieldUpdatedAt,
		})
	}
	if value, ok := guu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gooduser.FieldUpdatedAt,
		})
	}
	if value, ok := guu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gooduser.FieldDeletedAt,
		})
	}
	if value, ok := guu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gooduser.FieldDeletedAt,
		})
	}
	if value, ok := guu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: gooduser.FieldEntID,
		})
	}
	if value, ok := guu.mutation.RootUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: gooduser.FieldRootUserID,
		})
	}
	if guu.mutation.RootUserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: gooduser.FieldRootUserID,
		})
	}
	if value, ok := guu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gooduser.FieldName,
		})
	}
	if guu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: gooduser.FieldName,
		})
	}
	if value, ok := guu.mutation.CoinID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: gooduser.FieldCoinID,
		})
	}
	if guu.mutation.CoinIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: gooduser.FieldCoinID,
		})
	}
	if value, ok := guu.mutation.RevenueID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: gooduser.FieldRevenueID,
		})
	}
	if guu.mutation.RevenueIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: gooduser.FieldRevenueID,
		})
	}
	if value, ok := guu.mutation.HashRate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: gooduser.FieldHashRate,
		})
	}
	if value, ok := guu.mutation.AddedHashRate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: gooduser.FieldHashRate,
		})
	}
	if guu.mutation.HashRateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Column: gooduser.FieldHashRate,
		})
	}
	if value, ok := guu.mutation.ReadPageLink(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gooduser.FieldReadPageLink,
		})
	}
	if guu.mutation.ReadPageLinkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: gooduser.FieldReadPageLink,
		})
	}
	_spec.Modifiers = guu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, guu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gooduser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GoodUserUpdateOne is the builder for updating a single GoodUser entity.
type GoodUserUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *GoodUserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (guuo *GoodUserUpdateOne) SetCreatedAt(u uint32) *GoodUserUpdateOne {
	guuo.mutation.ResetCreatedAt()
	guuo.mutation.SetCreatedAt(u)
	return guuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (guuo *GoodUserUpdateOne) SetNillableCreatedAt(u *uint32) *GoodUserUpdateOne {
	if u != nil {
		guuo.SetCreatedAt(*u)
	}
	return guuo
}

// AddCreatedAt adds u to the "created_at" field.
func (guuo *GoodUserUpdateOne) AddCreatedAt(u int32) *GoodUserUpdateOne {
	guuo.mutation.AddCreatedAt(u)
	return guuo
}

// SetUpdatedAt sets the "updated_at" field.
func (guuo *GoodUserUpdateOne) SetUpdatedAt(u uint32) *GoodUserUpdateOne {
	guuo.mutation.ResetUpdatedAt()
	guuo.mutation.SetUpdatedAt(u)
	return guuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (guuo *GoodUserUpdateOne) AddUpdatedAt(u int32) *GoodUserUpdateOne {
	guuo.mutation.AddUpdatedAt(u)
	return guuo
}

// SetDeletedAt sets the "deleted_at" field.
func (guuo *GoodUserUpdateOne) SetDeletedAt(u uint32) *GoodUserUpdateOne {
	guuo.mutation.ResetDeletedAt()
	guuo.mutation.SetDeletedAt(u)
	return guuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (guuo *GoodUserUpdateOne) SetNillableDeletedAt(u *uint32) *GoodUserUpdateOne {
	if u != nil {
		guuo.SetDeletedAt(*u)
	}
	return guuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (guuo *GoodUserUpdateOne) AddDeletedAt(u int32) *GoodUserUpdateOne {
	guuo.mutation.AddDeletedAt(u)
	return guuo
}

// SetEntID sets the "ent_id" field.
func (guuo *GoodUserUpdateOne) SetEntID(u uuid.UUID) *GoodUserUpdateOne {
	guuo.mutation.SetEntID(u)
	return guuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (guuo *GoodUserUpdateOne) SetNillableEntID(u *uuid.UUID) *GoodUserUpdateOne {
	if u != nil {
		guuo.SetEntID(*u)
	}
	return guuo
}

// SetRootUserID sets the "root_user_id" field.
func (guuo *GoodUserUpdateOne) SetRootUserID(u uuid.UUID) *GoodUserUpdateOne {
	guuo.mutation.SetRootUserID(u)
	return guuo
}

// SetNillableRootUserID sets the "root_user_id" field if the given value is not nil.
func (guuo *GoodUserUpdateOne) SetNillableRootUserID(u *uuid.UUID) *GoodUserUpdateOne {
	if u != nil {
		guuo.SetRootUserID(*u)
	}
	return guuo
}

// ClearRootUserID clears the value of the "root_user_id" field.
func (guuo *GoodUserUpdateOne) ClearRootUserID() *GoodUserUpdateOne {
	guuo.mutation.ClearRootUserID()
	return guuo
}

// SetName sets the "name" field.
func (guuo *GoodUserUpdateOne) SetName(s string) *GoodUserUpdateOne {
	guuo.mutation.SetName(s)
	return guuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (guuo *GoodUserUpdateOne) SetNillableName(s *string) *GoodUserUpdateOne {
	if s != nil {
		guuo.SetName(*s)
	}
	return guuo
}

// ClearName clears the value of the "name" field.
func (guuo *GoodUserUpdateOne) ClearName() *GoodUserUpdateOne {
	guuo.mutation.ClearName()
	return guuo
}

// SetCoinID sets the "coin_id" field.
func (guuo *GoodUserUpdateOne) SetCoinID(u uuid.UUID) *GoodUserUpdateOne {
	guuo.mutation.SetCoinID(u)
	return guuo
}

// SetNillableCoinID sets the "coin_id" field if the given value is not nil.
func (guuo *GoodUserUpdateOne) SetNillableCoinID(u *uuid.UUID) *GoodUserUpdateOne {
	if u != nil {
		guuo.SetCoinID(*u)
	}
	return guuo
}

// ClearCoinID clears the value of the "coin_id" field.
func (guuo *GoodUserUpdateOne) ClearCoinID() *GoodUserUpdateOne {
	guuo.mutation.ClearCoinID()
	return guuo
}

// SetRevenueID sets the "revenue_id" field.
func (guuo *GoodUserUpdateOne) SetRevenueID(u uuid.UUID) *GoodUserUpdateOne {
	guuo.mutation.SetRevenueID(u)
	return guuo
}

// SetNillableRevenueID sets the "revenue_id" field if the given value is not nil.
func (guuo *GoodUserUpdateOne) SetNillableRevenueID(u *uuid.UUID) *GoodUserUpdateOne {
	if u != nil {
		guuo.SetRevenueID(*u)
	}
	return guuo
}

// ClearRevenueID clears the value of the "revenue_id" field.
func (guuo *GoodUserUpdateOne) ClearRevenueID() *GoodUserUpdateOne {
	guuo.mutation.ClearRevenueID()
	return guuo
}

// SetHashRate sets the "hash_rate" field.
func (guuo *GoodUserUpdateOne) SetHashRate(f float32) *GoodUserUpdateOne {
	guuo.mutation.ResetHashRate()
	guuo.mutation.SetHashRate(f)
	return guuo
}

// SetNillableHashRate sets the "hash_rate" field if the given value is not nil.
func (guuo *GoodUserUpdateOne) SetNillableHashRate(f *float32) *GoodUserUpdateOne {
	if f != nil {
		guuo.SetHashRate(*f)
	}
	return guuo
}

// AddHashRate adds f to the "hash_rate" field.
func (guuo *GoodUserUpdateOne) AddHashRate(f float32) *GoodUserUpdateOne {
	guuo.mutation.AddHashRate(f)
	return guuo
}

// ClearHashRate clears the value of the "hash_rate" field.
func (guuo *GoodUserUpdateOne) ClearHashRate() *GoodUserUpdateOne {
	guuo.mutation.ClearHashRate()
	return guuo
}

// SetReadPageLink sets the "read_page_link" field.
func (guuo *GoodUserUpdateOne) SetReadPageLink(s string) *GoodUserUpdateOne {
	guuo.mutation.SetReadPageLink(s)
	return guuo
}

// SetNillableReadPageLink sets the "read_page_link" field if the given value is not nil.
func (guuo *GoodUserUpdateOne) SetNillableReadPageLink(s *string) *GoodUserUpdateOne {
	if s != nil {
		guuo.SetReadPageLink(*s)
	}
	return guuo
}

// ClearReadPageLink clears the value of the "read_page_link" field.
func (guuo *GoodUserUpdateOne) ClearReadPageLink() *GoodUserUpdateOne {
	guuo.mutation.ClearReadPageLink()
	return guuo
}

// Mutation returns the GoodUserMutation object of the builder.
func (guuo *GoodUserUpdateOne) Mutation() *GoodUserMutation {
	return guuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guuo *GoodUserUpdateOne) Select(field string, fields ...string) *GoodUserUpdateOne {
	guuo.fields = append([]string{field}, fields...)
	return guuo
}

// Save executes the query and returns the updated GoodUser entity.
func (guuo *GoodUserUpdateOne) Save(ctx context.Context) (*GoodUser, error) {
	var (
		err  error
		node *GoodUser
	)
	if err := guuo.defaults(); err != nil {
		return nil, err
	}
	if len(guuo.hooks) == 0 {
		node, err = guuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			guuo.mutation = mutation
			node, err = guuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guuo.hooks) - 1; i >= 0; i-- {
			if guuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, guuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GoodUser)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GoodUserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guuo *GoodUserUpdateOne) SaveX(ctx context.Context) *GoodUser {
	node, err := guuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guuo *GoodUserUpdateOne) Exec(ctx context.Context) error {
	_, err := guuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guuo *GoodUserUpdateOne) ExecX(ctx context.Context) {
	if err := guuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guuo *GoodUserUpdateOne) defaults() error {
	if _, ok := guuo.mutation.UpdatedAt(); !ok {
		if gooduser.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized gooduser.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := gooduser.UpdateDefaultUpdatedAt()
		guuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (guuo *GoodUserUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GoodUserUpdateOne {
	guuo.modifiers = append(guuo.modifiers, modifiers...)
	return guuo
}

func (guuo *GoodUserUpdateOne) sqlSave(ctx context.Context) (_node *GoodUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gooduser.Table,
			Columns: gooduser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: gooduser.FieldID,
			},
		},
	}
	id, ok := guuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GoodUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gooduser.FieldID)
		for _, f := range fields {
			if !gooduser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != gooduser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gooduser.FieldCreatedAt,
		})
	}
	if value, ok := guuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gooduser.FieldCreatedAt,
		})
	}
	if value, ok := guuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gooduser.FieldUpdatedAt,
		})
	}
	if value, ok := guuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gooduser.FieldUpdatedAt,
		})
	}
	if value, ok := guuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gooduser.FieldDeletedAt,
		})
	}
	if value, ok := guuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gooduser.FieldDeletedAt,
		})
	}
	if value, ok := guuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: gooduser.FieldEntID,
		})
	}
	if value, ok := guuo.mutation.RootUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: gooduser.FieldRootUserID,
		})
	}
	if guuo.mutation.RootUserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: gooduser.FieldRootUserID,
		})
	}
	if value, ok := guuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gooduser.FieldName,
		})
	}
	if guuo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: gooduser.FieldName,
		})
	}
	if value, ok := guuo.mutation.CoinID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: gooduser.FieldCoinID,
		})
	}
	if guuo.mutation.CoinIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: gooduser.FieldCoinID,
		})
	}
	if value, ok := guuo.mutation.RevenueID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: gooduser.FieldRevenueID,
		})
	}
	if guuo.mutation.RevenueIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: gooduser.FieldRevenueID,
		})
	}
	if value, ok := guuo.mutation.HashRate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: gooduser.FieldHashRate,
		})
	}
	if value, ok := guuo.mutation.AddedHashRate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: gooduser.FieldHashRate,
		})
	}
	if guuo.mutation.HashRateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Column: gooduser.FieldHashRate,
		})
	}
	if value, ok := guuo.mutation.ReadPageLink(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gooduser.FieldReadPageLink,
		})
	}
	if guuo.mutation.ReadPageLinkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: gooduser.FieldReadPageLink,
		})
	}
	_spec.Modifiers = guuo.modifiers
	_node = &GoodUser{config: guuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gooduser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
