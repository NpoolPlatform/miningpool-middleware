// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// PoolUpdate is the builder for updating Pool entities.
type PoolUpdate struct {
	config
	hooks     []Hook
	mutation  *PoolMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PoolUpdate builder.
func (pu *PoolUpdate) Where(ps ...predicate.Pool) *PoolUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PoolUpdate) SetCreatedAt(u uint32) *PoolUpdate {
	pu.mutation.ResetCreatedAt()
	pu.mutation.SetCreatedAt(u)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PoolUpdate) SetNillableCreatedAt(u *uint32) *PoolUpdate {
	if u != nil {
		pu.SetCreatedAt(*u)
	}
	return pu
}

// AddCreatedAt adds u to the "created_at" field.
func (pu *PoolUpdate) AddCreatedAt(u int32) *PoolUpdate {
	pu.mutation.AddCreatedAt(u)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PoolUpdate) SetUpdatedAt(u uint32) *PoolUpdate {
	pu.mutation.ResetUpdatedAt()
	pu.mutation.SetUpdatedAt(u)
	return pu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (pu *PoolUpdate) AddUpdatedAt(u int32) *PoolUpdate {
	pu.mutation.AddUpdatedAt(u)
	return pu
}

// SetDeletedAt sets the "deleted_at" field.
func (pu *PoolUpdate) SetDeletedAt(u uint32) *PoolUpdate {
	pu.mutation.ResetDeletedAt()
	pu.mutation.SetDeletedAt(u)
	return pu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pu *PoolUpdate) SetNillableDeletedAt(u *uint32) *PoolUpdate {
	if u != nil {
		pu.SetDeletedAt(*u)
	}
	return pu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (pu *PoolUpdate) AddDeletedAt(u int32) *PoolUpdate {
	pu.mutation.AddDeletedAt(u)
	return pu
}

// SetEntID sets the "ent_id" field.
func (pu *PoolUpdate) SetEntID(u uuid.UUID) *PoolUpdate {
	pu.mutation.SetEntID(u)
	return pu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (pu *PoolUpdate) SetNillableEntID(u *uuid.UUID) *PoolUpdate {
	if u != nil {
		pu.SetEntID(*u)
	}
	return pu
}

// SetMiningpoolType sets the "miningpool_type" field.
func (pu *PoolUpdate) SetMiningpoolType(s string) *PoolUpdate {
	pu.mutation.SetMiningpoolType(s)
	return pu
}

// SetNillableMiningpoolType sets the "miningpool_type" field if the given value is not nil.
func (pu *PoolUpdate) SetNillableMiningpoolType(s *string) *PoolUpdate {
	if s != nil {
		pu.SetMiningpoolType(*s)
	}
	return pu
}

// ClearMiningpoolType clears the value of the "miningpool_type" field.
func (pu *PoolUpdate) ClearMiningpoolType() *PoolUpdate {
	pu.mutation.ClearMiningpoolType()
	return pu
}

// SetName sets the "name" field.
func (pu *PoolUpdate) SetName(s string) *PoolUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PoolUpdate) SetNillableName(s *string) *PoolUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// ClearName clears the value of the "name" field.
func (pu *PoolUpdate) ClearName() *PoolUpdate {
	pu.mutation.ClearName()
	return pu
}

// SetSite sets the "site" field.
func (pu *PoolUpdate) SetSite(s string) *PoolUpdate {
	pu.mutation.SetSite(s)
	return pu
}

// SetNillableSite sets the "site" field if the given value is not nil.
func (pu *PoolUpdate) SetNillableSite(s *string) *PoolUpdate {
	if s != nil {
		pu.SetSite(*s)
	}
	return pu
}

// ClearSite clears the value of the "site" field.
func (pu *PoolUpdate) ClearSite() *PoolUpdate {
	pu.mutation.ClearSite()
	return pu
}

// SetLogo sets the "logo" field.
func (pu *PoolUpdate) SetLogo(s string) *PoolUpdate {
	pu.mutation.SetLogo(s)
	return pu
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (pu *PoolUpdate) SetNillableLogo(s *string) *PoolUpdate {
	if s != nil {
		pu.SetLogo(*s)
	}
	return pu
}

// ClearLogo clears the value of the "logo" field.
func (pu *PoolUpdate) ClearLogo() *PoolUpdate {
	pu.mutation.ClearLogo()
	return pu
}

// SetDescription sets the "description" field.
func (pu *PoolUpdate) SetDescription(s string) *PoolUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PoolUpdate) SetNillableDescription(s *string) *PoolUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *PoolUpdate) ClearDescription() *PoolUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// Mutation returns the PoolMutation object of the builder.
func (pu *PoolUpdate) Mutation() *PoolMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PoolUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := pu.defaults(); err != nil {
		return 0, err
	}
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PoolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PoolUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PoolUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PoolUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PoolUpdate) defaults() error {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		if pool.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized pool.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := pool.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *PoolUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PoolUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *PoolUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pool.Table,
			Columns: pool.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: pool.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pool.FieldCreatedAt,
		})
	}
	if value, ok := pu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pool.FieldCreatedAt,
		})
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pool.FieldUpdatedAt,
		})
	}
	if value, ok := pu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pool.FieldUpdatedAt,
		})
	}
	if value, ok := pu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pool.FieldDeletedAt,
		})
	}
	if value, ok := pu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pool.FieldDeletedAt,
		})
	}
	if value, ok := pu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: pool.FieldEntID,
		})
	}
	if value, ok := pu.mutation.MiningpoolType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pool.FieldMiningpoolType,
		})
	}
	if pu.mutation.MiningpoolTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pool.FieldMiningpoolType,
		})
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pool.FieldName,
		})
	}
	if pu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pool.FieldName,
		})
	}
	if value, ok := pu.mutation.Site(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pool.FieldSite,
		})
	}
	if pu.mutation.SiteCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pool.FieldSite,
		})
	}
	if value, ok := pu.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pool.FieldLogo,
		})
	}
	if pu.mutation.LogoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pool.FieldLogo,
		})
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pool.FieldDescription,
		})
	}
	if pu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pool.FieldDescription,
		})
	}
	_spec.Modifiers = pu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PoolUpdateOne is the builder for updating a single Pool entity.
type PoolUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PoolMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (puo *PoolUpdateOne) SetCreatedAt(u uint32) *PoolUpdateOne {
	puo.mutation.ResetCreatedAt()
	puo.mutation.SetCreatedAt(u)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PoolUpdateOne) SetNillableCreatedAt(u *uint32) *PoolUpdateOne {
	if u != nil {
		puo.SetCreatedAt(*u)
	}
	return puo
}

// AddCreatedAt adds u to the "created_at" field.
func (puo *PoolUpdateOne) AddCreatedAt(u int32) *PoolUpdateOne {
	puo.mutation.AddCreatedAt(u)
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PoolUpdateOne) SetUpdatedAt(u uint32) *PoolUpdateOne {
	puo.mutation.ResetUpdatedAt()
	puo.mutation.SetUpdatedAt(u)
	return puo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (puo *PoolUpdateOne) AddUpdatedAt(u int32) *PoolUpdateOne {
	puo.mutation.AddUpdatedAt(u)
	return puo
}

// SetDeletedAt sets the "deleted_at" field.
func (puo *PoolUpdateOne) SetDeletedAt(u uint32) *PoolUpdateOne {
	puo.mutation.ResetDeletedAt()
	puo.mutation.SetDeletedAt(u)
	return puo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (puo *PoolUpdateOne) SetNillableDeletedAt(u *uint32) *PoolUpdateOne {
	if u != nil {
		puo.SetDeletedAt(*u)
	}
	return puo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (puo *PoolUpdateOne) AddDeletedAt(u int32) *PoolUpdateOne {
	puo.mutation.AddDeletedAt(u)
	return puo
}

// SetEntID sets the "ent_id" field.
func (puo *PoolUpdateOne) SetEntID(u uuid.UUID) *PoolUpdateOne {
	puo.mutation.SetEntID(u)
	return puo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (puo *PoolUpdateOne) SetNillableEntID(u *uuid.UUID) *PoolUpdateOne {
	if u != nil {
		puo.SetEntID(*u)
	}
	return puo
}

// SetMiningpoolType sets the "miningpool_type" field.
func (puo *PoolUpdateOne) SetMiningpoolType(s string) *PoolUpdateOne {
	puo.mutation.SetMiningpoolType(s)
	return puo
}

// SetNillableMiningpoolType sets the "miningpool_type" field if the given value is not nil.
func (puo *PoolUpdateOne) SetNillableMiningpoolType(s *string) *PoolUpdateOne {
	if s != nil {
		puo.SetMiningpoolType(*s)
	}
	return puo
}

// ClearMiningpoolType clears the value of the "miningpool_type" field.
func (puo *PoolUpdateOne) ClearMiningpoolType() *PoolUpdateOne {
	puo.mutation.ClearMiningpoolType()
	return puo
}

// SetName sets the "name" field.
func (puo *PoolUpdateOne) SetName(s string) *PoolUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PoolUpdateOne) SetNillableName(s *string) *PoolUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// ClearName clears the value of the "name" field.
func (puo *PoolUpdateOne) ClearName() *PoolUpdateOne {
	puo.mutation.ClearName()
	return puo
}

// SetSite sets the "site" field.
func (puo *PoolUpdateOne) SetSite(s string) *PoolUpdateOne {
	puo.mutation.SetSite(s)
	return puo
}

// SetNillableSite sets the "site" field if the given value is not nil.
func (puo *PoolUpdateOne) SetNillableSite(s *string) *PoolUpdateOne {
	if s != nil {
		puo.SetSite(*s)
	}
	return puo
}

// ClearSite clears the value of the "site" field.
func (puo *PoolUpdateOne) ClearSite() *PoolUpdateOne {
	puo.mutation.ClearSite()
	return puo
}

// SetLogo sets the "logo" field.
func (puo *PoolUpdateOne) SetLogo(s string) *PoolUpdateOne {
	puo.mutation.SetLogo(s)
	return puo
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (puo *PoolUpdateOne) SetNillableLogo(s *string) *PoolUpdateOne {
	if s != nil {
		puo.SetLogo(*s)
	}
	return puo
}

// ClearLogo clears the value of the "logo" field.
func (puo *PoolUpdateOne) ClearLogo() *PoolUpdateOne {
	puo.mutation.ClearLogo()
	return puo
}

// SetDescription sets the "description" field.
func (puo *PoolUpdateOne) SetDescription(s string) *PoolUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PoolUpdateOne) SetNillableDescription(s *string) *PoolUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *PoolUpdateOne) ClearDescription() *PoolUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// Mutation returns the PoolMutation object of the builder.
func (puo *PoolUpdateOne) Mutation() *PoolMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PoolUpdateOne) Select(field string, fields ...string) *PoolUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pool entity.
func (puo *PoolUpdateOne) Save(ctx context.Context) (*Pool, error) {
	var (
		err  error
		node *Pool
	)
	if err := puo.defaults(); err != nil {
		return nil, err
	}
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PoolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Pool)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PoolMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PoolUpdateOne) SaveX(ctx context.Context) *Pool {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PoolUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PoolUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PoolUpdateOne) defaults() error {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		if pool.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized pool.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := pool.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *PoolUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PoolUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *PoolUpdateOne) sqlSave(ctx context.Context) (_node *Pool, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pool.Table,
			Columns: pool.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: pool.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Pool.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pool.FieldID)
		for _, f := range fields {
			if !pool.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pool.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pool.FieldCreatedAt,
		})
	}
	if value, ok := puo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pool.FieldCreatedAt,
		})
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pool.FieldUpdatedAt,
		})
	}
	if value, ok := puo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pool.FieldUpdatedAt,
		})
	}
	if value, ok := puo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pool.FieldDeletedAt,
		})
	}
	if value, ok := puo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pool.FieldDeletedAt,
		})
	}
	if value, ok := puo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: pool.FieldEntID,
		})
	}
	if value, ok := puo.mutation.MiningpoolType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pool.FieldMiningpoolType,
		})
	}
	if puo.mutation.MiningpoolTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pool.FieldMiningpoolType,
		})
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pool.FieldName,
		})
	}
	if puo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pool.FieldName,
		})
	}
	if value, ok := puo.mutation.Site(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pool.FieldSite,
		})
	}
	if puo.mutation.SiteCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pool.FieldSite,
		})
	}
	if value, ok := puo.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pool.FieldLogo,
		})
	}
	if puo.mutation.LogoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pool.FieldLogo,
		})
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pool.FieldDescription,
		})
	}
	if puo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pool.FieldDescription,
		})
	}
	_spec.Modifiers = puo.modifiers
	_node = &Pool{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
