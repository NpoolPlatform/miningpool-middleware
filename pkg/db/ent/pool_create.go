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
	"github.com/google/uuid"
)

// PoolCreate is the builder for creating a Pool entity.
type PoolCreate struct {
	config
	mutation *PoolMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (pc *PoolCreate) SetCreatedAt(u uint32) *PoolCreate {
	pc.mutation.SetCreatedAt(u)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PoolCreate) SetNillableCreatedAt(u *uint32) *PoolCreate {
	if u != nil {
		pc.SetCreatedAt(*u)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PoolCreate) SetUpdatedAt(u uint32) *PoolCreate {
	pc.mutation.SetUpdatedAt(u)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PoolCreate) SetNillableUpdatedAt(u *uint32) *PoolCreate {
	if u != nil {
		pc.SetUpdatedAt(*u)
	}
	return pc
}

// SetDeletedAt sets the "deleted_at" field.
func (pc *PoolCreate) SetDeletedAt(u uint32) *PoolCreate {
	pc.mutation.SetDeletedAt(u)
	return pc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pc *PoolCreate) SetNillableDeletedAt(u *uint32) *PoolCreate {
	if u != nil {
		pc.SetDeletedAt(*u)
	}
	return pc
}

// SetEntID sets the "ent_id" field.
func (pc *PoolCreate) SetEntID(u uuid.UUID) *PoolCreate {
	pc.mutation.SetEntID(u)
	return pc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (pc *PoolCreate) SetNillableEntID(u *uuid.UUID) *PoolCreate {
	if u != nil {
		pc.SetEntID(*u)
	}
	return pc
}

// SetMiningpoolType sets the "miningpool_type" field.
func (pc *PoolCreate) SetMiningpoolType(s string) *PoolCreate {
	pc.mutation.SetMiningpoolType(s)
	return pc
}

// SetNillableMiningpoolType sets the "miningpool_type" field if the given value is not nil.
func (pc *PoolCreate) SetNillableMiningpoolType(s *string) *PoolCreate {
	if s != nil {
		pc.SetMiningpoolType(*s)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PoolCreate) SetName(s string) *PoolCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pc *PoolCreate) SetNillableName(s *string) *PoolCreate {
	if s != nil {
		pc.SetName(*s)
	}
	return pc
}

// SetSite sets the "site" field.
func (pc *PoolCreate) SetSite(s string) *PoolCreate {
	pc.mutation.SetSite(s)
	return pc
}

// SetNillableSite sets the "site" field if the given value is not nil.
func (pc *PoolCreate) SetNillableSite(s *string) *PoolCreate {
	if s != nil {
		pc.SetSite(*s)
	}
	return pc
}

// SetLogo sets the "logo" field.
func (pc *PoolCreate) SetLogo(s string) *PoolCreate {
	pc.mutation.SetLogo(s)
	return pc
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (pc *PoolCreate) SetNillableLogo(s *string) *PoolCreate {
	if s != nil {
		pc.SetLogo(*s)
	}
	return pc
}

// SetDescription sets the "description" field.
func (pc *PoolCreate) SetDescription(s string) *PoolCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *PoolCreate) SetNillableDescription(s *string) *PoolCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PoolCreate) SetID(u uint32) *PoolCreate {
	pc.mutation.SetID(u)
	return pc
}

// Mutation returns the PoolMutation object of the builder.
func (pc *PoolCreate) Mutation() *PoolMutation {
	return pc.mutation
}

// Save creates the Pool in the database.
func (pc *PoolCreate) Save(ctx context.Context) (*Pool, error) {
	var (
		err  error
		node *Pool
	)
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PoolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (pc *PoolCreate) SaveX(ctx context.Context) *Pool {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PoolCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PoolCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PoolCreate) defaults() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		if pool.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized pool.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := pool.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		if pool.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized pool.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := pool.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.DeletedAt(); !ok {
		if pool.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized pool.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := pool.DefaultDeletedAt()
		pc.mutation.SetDeletedAt(v)
	}
	if _, ok := pc.mutation.EntID(); !ok {
		if pool.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized pool.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := pool.DefaultEntID()
		pc.mutation.SetEntID(v)
	}
	if _, ok := pc.mutation.MiningpoolType(); !ok {
		v := pool.DefaultMiningpoolType
		pc.mutation.SetMiningpoolType(v)
	}
	if _, ok := pc.mutation.Name(); !ok {
		v := pool.DefaultName
		pc.mutation.SetName(v)
	}
	if _, ok := pc.mutation.Site(); !ok {
		v := pool.DefaultSite
		pc.mutation.SetSite(v)
	}
	if _, ok := pc.mutation.Logo(); !ok {
		v := pool.DefaultLogo
		pc.mutation.SetLogo(v)
	}
	if _, ok := pc.mutation.Description(); !ok {
		v := pool.DefaultDescription
		pc.mutation.SetDescription(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *PoolCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Pool.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Pool.updated_at"`)}
	}
	if _, ok := pc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Pool.deleted_at"`)}
	}
	if _, ok := pc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "Pool.ent_id"`)}
	}
	return nil
}

func (pc *PoolCreate) sqlSave(ctx context.Context) (*Pool, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
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

func (pc *PoolCreate) createSpec() (*Pool, *sqlgraph.CreateSpec) {
	var (
		_node = &Pool{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: pool.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: pool.FieldID,
			},
		}
	)
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pool.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pool.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pool.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := pc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: pool.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := pc.mutation.MiningpoolType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pool.FieldMiningpoolType,
		})
		_node.MiningpoolType = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pool.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.Site(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pool.FieldSite,
		})
		_node.Site = value
	}
	if value, ok := pc.mutation.Logo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pool.FieldLogo,
		})
		_node.Logo = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pool.FieldDescription,
		})
		_node.Description = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Pool.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PoolUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (pc *PoolCreate) OnConflict(opts ...sql.ConflictOption) *PoolUpsertOne {
	pc.conflict = opts
	return &PoolUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Pool.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PoolCreate) OnConflictColumns(columns ...string) *PoolUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PoolUpsertOne{
		create: pc,
	}
}

type (
	// PoolUpsertOne is the builder for "upsert"-ing
	//  one Pool node.
	PoolUpsertOne struct {
		create *PoolCreate
	}

	// PoolUpsert is the "OnConflict" setter.
	PoolUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *PoolUpsert) SetCreatedAt(v uint32) *PoolUpsert {
	u.Set(pool.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PoolUpsert) UpdateCreatedAt() *PoolUpsert {
	u.SetExcluded(pool.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *PoolUpsert) AddCreatedAt(v uint32) *PoolUpsert {
	u.Add(pool.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PoolUpsert) SetUpdatedAt(v uint32) *PoolUpsert {
	u.Set(pool.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PoolUpsert) UpdateUpdatedAt() *PoolUpsert {
	u.SetExcluded(pool.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *PoolUpsert) AddUpdatedAt(v uint32) *PoolUpsert {
	u.Add(pool.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PoolUpsert) SetDeletedAt(v uint32) *PoolUpsert {
	u.Set(pool.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PoolUpsert) UpdateDeletedAt() *PoolUpsert {
	u.SetExcluded(pool.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *PoolUpsert) AddDeletedAt(v uint32) *PoolUpsert {
	u.Add(pool.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *PoolUpsert) SetEntID(v uuid.UUID) *PoolUpsert {
	u.Set(pool.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *PoolUpsert) UpdateEntID() *PoolUpsert {
	u.SetExcluded(pool.FieldEntID)
	return u
}

// SetMiningpoolType sets the "miningpool_type" field.
func (u *PoolUpsert) SetMiningpoolType(v string) *PoolUpsert {
	u.Set(pool.FieldMiningpoolType, v)
	return u
}

// UpdateMiningpoolType sets the "miningpool_type" field to the value that was provided on create.
func (u *PoolUpsert) UpdateMiningpoolType() *PoolUpsert {
	u.SetExcluded(pool.FieldMiningpoolType)
	return u
}

// ClearMiningpoolType clears the value of the "miningpool_type" field.
func (u *PoolUpsert) ClearMiningpoolType() *PoolUpsert {
	u.SetNull(pool.FieldMiningpoolType)
	return u
}

// SetName sets the "name" field.
func (u *PoolUpsert) SetName(v string) *PoolUpsert {
	u.Set(pool.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PoolUpsert) UpdateName() *PoolUpsert {
	u.SetExcluded(pool.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *PoolUpsert) ClearName() *PoolUpsert {
	u.SetNull(pool.FieldName)
	return u
}

// SetSite sets the "site" field.
func (u *PoolUpsert) SetSite(v string) *PoolUpsert {
	u.Set(pool.FieldSite, v)
	return u
}

// UpdateSite sets the "site" field to the value that was provided on create.
func (u *PoolUpsert) UpdateSite() *PoolUpsert {
	u.SetExcluded(pool.FieldSite)
	return u
}

// ClearSite clears the value of the "site" field.
func (u *PoolUpsert) ClearSite() *PoolUpsert {
	u.SetNull(pool.FieldSite)
	return u
}

// SetLogo sets the "logo" field.
func (u *PoolUpsert) SetLogo(v string) *PoolUpsert {
	u.Set(pool.FieldLogo, v)
	return u
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *PoolUpsert) UpdateLogo() *PoolUpsert {
	u.SetExcluded(pool.FieldLogo)
	return u
}

// ClearLogo clears the value of the "logo" field.
func (u *PoolUpsert) ClearLogo() *PoolUpsert {
	u.SetNull(pool.FieldLogo)
	return u
}

// SetDescription sets the "description" field.
func (u *PoolUpsert) SetDescription(v string) *PoolUpsert {
	u.Set(pool.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PoolUpsert) UpdateDescription() *PoolUpsert {
	u.SetExcluded(pool.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *PoolUpsert) ClearDescription() *PoolUpsert {
	u.SetNull(pool.FieldDescription)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Pool.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(pool.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PoolUpsertOne) UpdateNewValues() *PoolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(pool.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Pool.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PoolUpsertOne) Ignore() *PoolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PoolUpsertOne) DoNothing() *PoolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PoolCreate.OnConflict
// documentation for more info.
func (u *PoolUpsertOne) Update(set func(*PoolUpsert)) *PoolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PoolUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *PoolUpsertOne) SetCreatedAt(v uint32) *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *PoolUpsertOne) AddCreatedAt(v uint32) *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PoolUpsertOne) UpdateCreatedAt() *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PoolUpsertOne) SetUpdatedAt(v uint32) *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *PoolUpsertOne) AddUpdatedAt(v uint32) *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PoolUpsertOne) UpdateUpdatedAt() *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PoolUpsertOne) SetDeletedAt(v uint32) *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *PoolUpsertOne) AddDeletedAt(v uint32) *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PoolUpsertOne) UpdateDeletedAt() *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *PoolUpsertOne) SetEntID(v uuid.UUID) *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *PoolUpsertOne) UpdateEntID() *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateEntID()
	})
}

// SetMiningpoolType sets the "miningpool_type" field.
func (u *PoolUpsertOne) SetMiningpoolType(v string) *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.SetMiningpoolType(v)
	})
}

// UpdateMiningpoolType sets the "miningpool_type" field to the value that was provided on create.
func (u *PoolUpsertOne) UpdateMiningpoolType() *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateMiningpoolType()
	})
}

// ClearMiningpoolType clears the value of the "miningpool_type" field.
func (u *PoolUpsertOne) ClearMiningpoolType() *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.ClearMiningpoolType()
	})
}

// SetName sets the "name" field.
func (u *PoolUpsertOne) SetName(v string) *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PoolUpsertOne) UpdateName() *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *PoolUpsertOne) ClearName() *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.ClearName()
	})
}

// SetSite sets the "site" field.
func (u *PoolUpsertOne) SetSite(v string) *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.SetSite(v)
	})
}

// UpdateSite sets the "site" field to the value that was provided on create.
func (u *PoolUpsertOne) UpdateSite() *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateSite()
	})
}

// ClearSite clears the value of the "site" field.
func (u *PoolUpsertOne) ClearSite() *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.ClearSite()
	})
}

// SetLogo sets the "logo" field.
func (u *PoolUpsertOne) SetLogo(v string) *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.SetLogo(v)
	})
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *PoolUpsertOne) UpdateLogo() *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateLogo()
	})
}

// ClearLogo clears the value of the "logo" field.
func (u *PoolUpsertOne) ClearLogo() *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.ClearLogo()
	})
}

// SetDescription sets the "description" field.
func (u *PoolUpsertOne) SetDescription(v string) *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PoolUpsertOne) UpdateDescription() *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *PoolUpsertOne) ClearDescription() *PoolUpsertOne {
	return u.Update(func(s *PoolUpsert) {
		s.ClearDescription()
	})
}

// Exec executes the query.
func (u *PoolUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PoolCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PoolUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PoolUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PoolUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PoolCreateBulk is the builder for creating many Pool entities in bulk.
type PoolCreateBulk struct {
	config
	builders []*PoolCreate
	conflict []sql.ConflictOption
}

// Save creates the Pool entities in the database.
func (pcb *PoolCreateBulk) Save(ctx context.Context) ([]*Pool, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Pool, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PoolMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PoolCreateBulk) SaveX(ctx context.Context) []*Pool {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PoolCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PoolCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Pool.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PoolUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (pcb *PoolCreateBulk) OnConflict(opts ...sql.ConflictOption) *PoolUpsertBulk {
	pcb.conflict = opts
	return &PoolUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Pool.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PoolCreateBulk) OnConflictColumns(columns ...string) *PoolUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PoolUpsertBulk{
		create: pcb,
	}
}

// PoolUpsertBulk is the builder for "upsert"-ing
// a bulk of Pool nodes.
type PoolUpsertBulk struct {
	create *PoolCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Pool.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(pool.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PoolUpsertBulk) UpdateNewValues() *PoolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(pool.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Pool.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PoolUpsertBulk) Ignore() *PoolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PoolUpsertBulk) DoNothing() *PoolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PoolCreateBulk.OnConflict
// documentation for more info.
func (u *PoolUpsertBulk) Update(set func(*PoolUpsert)) *PoolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PoolUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *PoolUpsertBulk) SetCreatedAt(v uint32) *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *PoolUpsertBulk) AddCreatedAt(v uint32) *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PoolUpsertBulk) UpdateCreatedAt() *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PoolUpsertBulk) SetUpdatedAt(v uint32) *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *PoolUpsertBulk) AddUpdatedAt(v uint32) *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PoolUpsertBulk) UpdateUpdatedAt() *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PoolUpsertBulk) SetDeletedAt(v uint32) *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *PoolUpsertBulk) AddDeletedAt(v uint32) *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PoolUpsertBulk) UpdateDeletedAt() *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *PoolUpsertBulk) SetEntID(v uuid.UUID) *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *PoolUpsertBulk) UpdateEntID() *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateEntID()
	})
}

// SetMiningpoolType sets the "miningpool_type" field.
func (u *PoolUpsertBulk) SetMiningpoolType(v string) *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.SetMiningpoolType(v)
	})
}

// UpdateMiningpoolType sets the "miningpool_type" field to the value that was provided on create.
func (u *PoolUpsertBulk) UpdateMiningpoolType() *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateMiningpoolType()
	})
}

// ClearMiningpoolType clears the value of the "miningpool_type" field.
func (u *PoolUpsertBulk) ClearMiningpoolType() *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.ClearMiningpoolType()
	})
}

// SetName sets the "name" field.
func (u *PoolUpsertBulk) SetName(v string) *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PoolUpsertBulk) UpdateName() *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *PoolUpsertBulk) ClearName() *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.ClearName()
	})
}

// SetSite sets the "site" field.
func (u *PoolUpsertBulk) SetSite(v string) *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.SetSite(v)
	})
}

// UpdateSite sets the "site" field to the value that was provided on create.
func (u *PoolUpsertBulk) UpdateSite() *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateSite()
	})
}

// ClearSite clears the value of the "site" field.
func (u *PoolUpsertBulk) ClearSite() *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.ClearSite()
	})
}

// SetLogo sets the "logo" field.
func (u *PoolUpsertBulk) SetLogo(v string) *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.SetLogo(v)
	})
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *PoolUpsertBulk) UpdateLogo() *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateLogo()
	})
}

// ClearLogo clears the value of the "logo" field.
func (u *PoolUpsertBulk) ClearLogo() *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.ClearLogo()
	})
}

// SetDescription sets the "description" field.
func (u *PoolUpsertBulk) SetDescription(v string) *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PoolUpsertBulk) UpdateDescription() *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *PoolUpsertBulk) ClearDescription() *PoolUpsertBulk {
	return u.Update(func(s *PoolUpsert) {
		s.ClearDescription()
	})
}

// Exec executes the query.
func (u *PoolUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PoolCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PoolCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PoolUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
