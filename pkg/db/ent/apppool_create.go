// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/apppool"
	"github.com/google/uuid"
)

// AppPoolCreate is the builder for creating a AppPool entity.
type AppPoolCreate struct {
	config
	mutation *AppPoolMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (apc *AppPoolCreate) SetCreatedAt(u uint32) *AppPoolCreate {
	apc.mutation.SetCreatedAt(u)
	return apc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (apc *AppPoolCreate) SetNillableCreatedAt(u *uint32) *AppPoolCreate {
	if u != nil {
		apc.SetCreatedAt(*u)
	}
	return apc
}

// SetUpdatedAt sets the "updated_at" field.
func (apc *AppPoolCreate) SetUpdatedAt(u uint32) *AppPoolCreate {
	apc.mutation.SetUpdatedAt(u)
	return apc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (apc *AppPoolCreate) SetNillableUpdatedAt(u *uint32) *AppPoolCreate {
	if u != nil {
		apc.SetUpdatedAt(*u)
	}
	return apc
}

// SetDeletedAt sets the "deleted_at" field.
func (apc *AppPoolCreate) SetDeletedAt(u uint32) *AppPoolCreate {
	apc.mutation.SetDeletedAt(u)
	return apc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (apc *AppPoolCreate) SetNillableDeletedAt(u *uint32) *AppPoolCreate {
	if u != nil {
		apc.SetDeletedAt(*u)
	}
	return apc
}

// SetEntID sets the "ent_id" field.
func (apc *AppPoolCreate) SetEntID(u uuid.UUID) *AppPoolCreate {
	apc.mutation.SetEntID(u)
	return apc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (apc *AppPoolCreate) SetNillableEntID(u *uuid.UUID) *AppPoolCreate {
	if u != nil {
		apc.SetEntID(*u)
	}
	return apc
}

// SetAppID sets the "app_id" field.
func (apc *AppPoolCreate) SetAppID(u uuid.UUID) *AppPoolCreate {
	apc.mutation.SetAppID(u)
	return apc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (apc *AppPoolCreate) SetNillableAppID(u *uuid.UUID) *AppPoolCreate {
	if u != nil {
		apc.SetAppID(*u)
	}
	return apc
}

// SetPoolID sets the "pool_id" field.
func (apc *AppPoolCreate) SetPoolID(u uuid.UUID) *AppPoolCreate {
	apc.mutation.SetPoolID(u)
	return apc
}

// SetNillablePoolID sets the "pool_id" field if the given value is not nil.
func (apc *AppPoolCreate) SetNillablePoolID(u *uuid.UUID) *AppPoolCreate {
	if u != nil {
		apc.SetPoolID(*u)
	}
	return apc
}

// SetID sets the "id" field.
func (apc *AppPoolCreate) SetID(u uint32) *AppPoolCreate {
	apc.mutation.SetID(u)
	return apc
}

// Mutation returns the AppPoolMutation object of the builder.
func (apc *AppPoolCreate) Mutation() *AppPoolMutation {
	return apc.mutation
}

// Save creates the AppPool in the database.
func (apc *AppPoolCreate) Save(ctx context.Context) (*AppPool, error) {
	var (
		err  error
		node *AppPool
	)
	if err := apc.defaults(); err != nil {
		return nil, err
	}
	if len(apc.hooks) == 0 {
		if err = apc.check(); err != nil {
			return nil, err
		}
		node, err = apc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppPoolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = apc.check(); err != nil {
				return nil, err
			}
			apc.mutation = mutation
			if node, err = apc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(apc.hooks) - 1; i >= 0; i-- {
			if apc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = apc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, apc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppPool)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppPoolMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (apc *AppPoolCreate) SaveX(ctx context.Context) *AppPool {
	v, err := apc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (apc *AppPoolCreate) Exec(ctx context.Context) error {
	_, err := apc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apc *AppPoolCreate) ExecX(ctx context.Context) {
	if err := apc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (apc *AppPoolCreate) defaults() error {
	if _, ok := apc.mutation.CreatedAt(); !ok {
		if apppool.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized apppool.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := apppool.DefaultCreatedAt()
		apc.mutation.SetCreatedAt(v)
	}
	if _, ok := apc.mutation.UpdatedAt(); !ok {
		if apppool.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized apppool.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := apppool.DefaultUpdatedAt()
		apc.mutation.SetUpdatedAt(v)
	}
	if _, ok := apc.mutation.DeletedAt(); !ok {
		if apppool.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized apppool.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := apppool.DefaultDeletedAt()
		apc.mutation.SetDeletedAt(v)
	}
	if _, ok := apc.mutation.EntID(); !ok {
		if apppool.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized apppool.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := apppool.DefaultEntID()
		apc.mutation.SetEntID(v)
	}
	if _, ok := apc.mutation.AppID(); !ok {
		if apppool.DefaultAppID == nil {
			return fmt.Errorf("ent: uninitialized apppool.DefaultAppID (forgotten import ent/runtime?)")
		}
		v := apppool.DefaultAppID()
		apc.mutation.SetAppID(v)
	}
	if _, ok := apc.mutation.PoolID(); !ok {
		if apppool.DefaultPoolID == nil {
			return fmt.Errorf("ent: uninitialized apppool.DefaultPoolID (forgotten import ent/runtime?)")
		}
		v := apppool.DefaultPoolID()
		apc.mutation.SetPoolID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (apc *AppPoolCreate) check() error {
	if _, ok := apc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppPool.created_at"`)}
	}
	if _, ok := apc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppPool.updated_at"`)}
	}
	if _, ok := apc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "AppPool.deleted_at"`)}
	}
	if _, ok := apc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "AppPool.ent_id"`)}
	}
	return nil
}

func (apc *AppPoolCreate) sqlSave(ctx context.Context) (*AppPool, error) {
	_node, _spec := apc.createSpec()
	if err := sqlgraph.CreateNode(ctx, apc.driver, _spec); err != nil {
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

func (apc *AppPoolCreate) createSpec() (*AppPool, *sqlgraph.CreateSpec) {
	var (
		_node = &AppPool{config: apc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: apppool.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: apppool.FieldID,
			},
		}
	)
	_spec.OnConflict = apc.conflict
	if id, ok := apc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := apc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: apppool.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := apc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: apppool.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := apc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: apppool.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := apc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: apppool.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := apc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: apppool.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := apc.mutation.PoolID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: apppool.FieldPoolID,
		})
		_node.PoolID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppPool.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppPoolUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (apc *AppPoolCreate) OnConflict(opts ...sql.ConflictOption) *AppPoolUpsertOne {
	apc.conflict = opts
	return &AppPoolUpsertOne{
		create: apc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppPool.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (apc *AppPoolCreate) OnConflictColumns(columns ...string) *AppPoolUpsertOne {
	apc.conflict = append(apc.conflict, sql.ConflictColumns(columns...))
	return &AppPoolUpsertOne{
		create: apc,
	}
}

type (
	// AppPoolUpsertOne is the builder for "upsert"-ing
	//  one AppPool node.
	AppPoolUpsertOne struct {
		create *AppPoolCreate
	}

	// AppPoolUpsert is the "OnConflict" setter.
	AppPoolUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AppPoolUpsert) SetCreatedAt(v uint32) *AppPoolUpsert {
	u.Set(apppool.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppPoolUpsert) UpdateCreatedAt() *AppPoolUpsert {
	u.SetExcluded(apppool.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppPoolUpsert) AddCreatedAt(v uint32) *AppPoolUpsert {
	u.Add(apppool.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppPoolUpsert) SetUpdatedAt(v uint32) *AppPoolUpsert {
	u.Set(apppool.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppPoolUpsert) UpdateUpdatedAt() *AppPoolUpsert {
	u.SetExcluded(apppool.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppPoolUpsert) AddUpdatedAt(v uint32) *AppPoolUpsert {
	u.Add(apppool.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppPoolUpsert) SetDeletedAt(v uint32) *AppPoolUpsert {
	u.Set(apppool.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppPoolUpsert) UpdateDeletedAt() *AppPoolUpsert {
	u.SetExcluded(apppool.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppPoolUpsert) AddDeletedAt(v uint32) *AppPoolUpsert {
	u.Add(apppool.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *AppPoolUpsert) SetEntID(v uuid.UUID) *AppPoolUpsert {
	u.Set(apppool.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppPoolUpsert) UpdateEntID() *AppPoolUpsert {
	u.SetExcluded(apppool.FieldEntID)
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppPoolUpsert) SetAppID(v uuid.UUID) *AppPoolUpsert {
	u.Set(apppool.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppPoolUpsert) UpdateAppID() *AppPoolUpsert {
	u.SetExcluded(apppool.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppPoolUpsert) ClearAppID() *AppPoolUpsert {
	u.SetNull(apppool.FieldAppID)
	return u
}

// SetPoolID sets the "pool_id" field.
func (u *AppPoolUpsert) SetPoolID(v uuid.UUID) *AppPoolUpsert {
	u.Set(apppool.FieldPoolID, v)
	return u
}

// UpdatePoolID sets the "pool_id" field to the value that was provided on create.
func (u *AppPoolUpsert) UpdatePoolID() *AppPoolUpsert {
	u.SetExcluded(apppool.FieldPoolID)
	return u
}

// ClearPoolID clears the value of the "pool_id" field.
func (u *AppPoolUpsert) ClearPoolID() *AppPoolUpsert {
	u.SetNull(apppool.FieldPoolID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppPool.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(apppool.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppPoolUpsertOne) UpdateNewValues() *AppPoolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(apppool.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppPool.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppPoolUpsertOne) Ignore() *AppPoolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppPoolUpsertOne) DoNothing() *AppPoolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppPoolCreate.OnConflict
// documentation for more info.
func (u *AppPoolUpsertOne) Update(set func(*AppPoolUpsert)) *AppPoolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppPoolUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppPoolUpsertOne) SetCreatedAt(v uint32) *AppPoolUpsertOne {
	return u.Update(func(s *AppPoolUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppPoolUpsertOne) AddCreatedAt(v uint32) *AppPoolUpsertOne {
	return u.Update(func(s *AppPoolUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppPoolUpsertOne) UpdateCreatedAt() *AppPoolUpsertOne {
	return u.Update(func(s *AppPoolUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppPoolUpsertOne) SetUpdatedAt(v uint32) *AppPoolUpsertOne {
	return u.Update(func(s *AppPoolUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppPoolUpsertOne) AddUpdatedAt(v uint32) *AppPoolUpsertOne {
	return u.Update(func(s *AppPoolUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppPoolUpsertOne) UpdateUpdatedAt() *AppPoolUpsertOne {
	return u.Update(func(s *AppPoolUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppPoolUpsertOne) SetDeletedAt(v uint32) *AppPoolUpsertOne {
	return u.Update(func(s *AppPoolUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppPoolUpsertOne) AddDeletedAt(v uint32) *AppPoolUpsertOne {
	return u.Update(func(s *AppPoolUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppPoolUpsertOne) UpdateDeletedAt() *AppPoolUpsertOne {
	return u.Update(func(s *AppPoolUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AppPoolUpsertOne) SetEntID(v uuid.UUID) *AppPoolUpsertOne {
	return u.Update(func(s *AppPoolUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppPoolUpsertOne) UpdateEntID() *AppPoolUpsertOne {
	return u.Update(func(s *AppPoolUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppPoolUpsertOne) SetAppID(v uuid.UUID) *AppPoolUpsertOne {
	return u.Update(func(s *AppPoolUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppPoolUpsertOne) UpdateAppID() *AppPoolUpsertOne {
	return u.Update(func(s *AppPoolUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppPoolUpsertOne) ClearAppID() *AppPoolUpsertOne {
	return u.Update(func(s *AppPoolUpsert) {
		s.ClearAppID()
	})
}

// SetPoolID sets the "pool_id" field.
func (u *AppPoolUpsertOne) SetPoolID(v uuid.UUID) *AppPoolUpsertOne {
	return u.Update(func(s *AppPoolUpsert) {
		s.SetPoolID(v)
	})
}

// UpdatePoolID sets the "pool_id" field to the value that was provided on create.
func (u *AppPoolUpsertOne) UpdatePoolID() *AppPoolUpsertOne {
	return u.Update(func(s *AppPoolUpsert) {
		s.UpdatePoolID()
	})
}

// ClearPoolID clears the value of the "pool_id" field.
func (u *AppPoolUpsertOne) ClearPoolID() *AppPoolUpsertOne {
	return u.Update(func(s *AppPoolUpsert) {
		s.ClearPoolID()
	})
}

// Exec executes the query.
func (u *AppPoolUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppPoolCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppPoolUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppPoolUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppPoolUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppPoolCreateBulk is the builder for creating many AppPool entities in bulk.
type AppPoolCreateBulk struct {
	config
	builders []*AppPoolCreate
	conflict []sql.ConflictOption
}

// Save creates the AppPool entities in the database.
func (apcb *AppPoolCreateBulk) Save(ctx context.Context) ([]*AppPool, error) {
	specs := make([]*sqlgraph.CreateSpec, len(apcb.builders))
	nodes := make([]*AppPool, len(apcb.builders))
	mutators := make([]Mutator, len(apcb.builders))
	for i := range apcb.builders {
		func(i int, root context.Context) {
			builder := apcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppPoolMutation)
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
					_, err = mutators[i+1].Mutate(root, apcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = apcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, apcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, apcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (apcb *AppPoolCreateBulk) SaveX(ctx context.Context) []*AppPool {
	v, err := apcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (apcb *AppPoolCreateBulk) Exec(ctx context.Context) error {
	_, err := apcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apcb *AppPoolCreateBulk) ExecX(ctx context.Context) {
	if err := apcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppPool.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppPoolUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (apcb *AppPoolCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppPoolUpsertBulk {
	apcb.conflict = opts
	return &AppPoolUpsertBulk{
		create: apcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppPool.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (apcb *AppPoolCreateBulk) OnConflictColumns(columns ...string) *AppPoolUpsertBulk {
	apcb.conflict = append(apcb.conflict, sql.ConflictColumns(columns...))
	return &AppPoolUpsertBulk{
		create: apcb,
	}
}

// AppPoolUpsertBulk is the builder for "upsert"-ing
// a bulk of AppPool nodes.
type AppPoolUpsertBulk struct {
	create *AppPoolCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppPool.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(apppool.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppPoolUpsertBulk) UpdateNewValues() *AppPoolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(apppool.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppPool.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppPoolUpsertBulk) Ignore() *AppPoolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppPoolUpsertBulk) DoNothing() *AppPoolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppPoolCreateBulk.OnConflict
// documentation for more info.
func (u *AppPoolUpsertBulk) Update(set func(*AppPoolUpsert)) *AppPoolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppPoolUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppPoolUpsertBulk) SetCreatedAt(v uint32) *AppPoolUpsertBulk {
	return u.Update(func(s *AppPoolUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppPoolUpsertBulk) AddCreatedAt(v uint32) *AppPoolUpsertBulk {
	return u.Update(func(s *AppPoolUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppPoolUpsertBulk) UpdateCreatedAt() *AppPoolUpsertBulk {
	return u.Update(func(s *AppPoolUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppPoolUpsertBulk) SetUpdatedAt(v uint32) *AppPoolUpsertBulk {
	return u.Update(func(s *AppPoolUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppPoolUpsertBulk) AddUpdatedAt(v uint32) *AppPoolUpsertBulk {
	return u.Update(func(s *AppPoolUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppPoolUpsertBulk) UpdateUpdatedAt() *AppPoolUpsertBulk {
	return u.Update(func(s *AppPoolUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppPoolUpsertBulk) SetDeletedAt(v uint32) *AppPoolUpsertBulk {
	return u.Update(func(s *AppPoolUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppPoolUpsertBulk) AddDeletedAt(v uint32) *AppPoolUpsertBulk {
	return u.Update(func(s *AppPoolUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppPoolUpsertBulk) UpdateDeletedAt() *AppPoolUpsertBulk {
	return u.Update(func(s *AppPoolUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AppPoolUpsertBulk) SetEntID(v uuid.UUID) *AppPoolUpsertBulk {
	return u.Update(func(s *AppPoolUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppPoolUpsertBulk) UpdateEntID() *AppPoolUpsertBulk {
	return u.Update(func(s *AppPoolUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppPoolUpsertBulk) SetAppID(v uuid.UUID) *AppPoolUpsertBulk {
	return u.Update(func(s *AppPoolUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppPoolUpsertBulk) UpdateAppID() *AppPoolUpsertBulk {
	return u.Update(func(s *AppPoolUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppPoolUpsertBulk) ClearAppID() *AppPoolUpsertBulk {
	return u.Update(func(s *AppPoolUpsert) {
		s.ClearAppID()
	})
}

// SetPoolID sets the "pool_id" field.
func (u *AppPoolUpsertBulk) SetPoolID(v uuid.UUID) *AppPoolUpsertBulk {
	return u.Update(func(s *AppPoolUpsert) {
		s.SetPoolID(v)
	})
}

// UpdatePoolID sets the "pool_id" field to the value that was provided on create.
func (u *AppPoolUpsertBulk) UpdatePoolID() *AppPoolUpsertBulk {
	return u.Update(func(s *AppPoolUpsert) {
		s.UpdatePoolID()
	})
}

// ClearPoolID clears the value of the "pool_id" field.
func (u *AppPoolUpsertBulk) ClearPoolID() *AppPoolUpsertBulk {
	return u.Update(func(s *AppPoolUpsert) {
		s.ClearPoolID()
	})
}

// Exec executes the query.
func (u *AppPoolUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppPoolCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppPoolCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppPoolUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
