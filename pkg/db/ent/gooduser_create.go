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
	"github.com/google/uuid"
)

// GoodUserCreate is the builder for creating a GoodUser entity.
type GoodUserCreate struct {
	config
	mutation *GoodUserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (guc *GoodUserCreate) SetCreatedAt(u uint32) *GoodUserCreate {
	guc.mutation.SetCreatedAt(u)
	return guc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (guc *GoodUserCreate) SetNillableCreatedAt(u *uint32) *GoodUserCreate {
	if u != nil {
		guc.SetCreatedAt(*u)
	}
	return guc
}

// SetUpdatedAt sets the "updated_at" field.
func (guc *GoodUserCreate) SetUpdatedAt(u uint32) *GoodUserCreate {
	guc.mutation.SetUpdatedAt(u)
	return guc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (guc *GoodUserCreate) SetNillableUpdatedAt(u *uint32) *GoodUserCreate {
	if u != nil {
		guc.SetUpdatedAt(*u)
	}
	return guc
}

// SetDeletedAt sets the "deleted_at" field.
func (guc *GoodUserCreate) SetDeletedAt(u uint32) *GoodUserCreate {
	guc.mutation.SetDeletedAt(u)
	return guc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (guc *GoodUserCreate) SetNillableDeletedAt(u *uint32) *GoodUserCreate {
	if u != nil {
		guc.SetDeletedAt(*u)
	}
	return guc
}

// SetEntID sets the "ent_id" field.
func (guc *GoodUserCreate) SetEntID(u uuid.UUID) *GoodUserCreate {
	guc.mutation.SetEntID(u)
	return guc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (guc *GoodUserCreate) SetNillableEntID(u *uuid.UUID) *GoodUserCreate {
	if u != nil {
		guc.SetEntID(*u)
	}
	return guc
}

// SetRootUserID sets the "root_user_id" field.
func (guc *GoodUserCreate) SetRootUserID(u uuid.UUID) *GoodUserCreate {
	guc.mutation.SetRootUserID(u)
	return guc
}

// SetNillableRootUserID sets the "root_user_id" field if the given value is not nil.
func (guc *GoodUserCreate) SetNillableRootUserID(u *uuid.UUID) *GoodUserCreate {
	if u != nil {
		guc.SetRootUserID(*u)
	}
	return guc
}

// SetName sets the "name" field.
func (guc *GoodUserCreate) SetName(s string) *GoodUserCreate {
	guc.mutation.SetName(s)
	return guc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (guc *GoodUserCreate) SetNillableName(s *string) *GoodUserCreate {
	if s != nil {
		guc.SetName(*s)
	}
	return guc
}

// SetPoolCoinTypeID sets the "pool_coin_type_id" field.
func (guc *GoodUserCreate) SetPoolCoinTypeID(u uuid.UUID) *GoodUserCreate {
	guc.mutation.SetPoolCoinTypeID(u)
	return guc
}

// SetNillablePoolCoinTypeID sets the "pool_coin_type_id" field if the given value is not nil.
func (guc *GoodUserCreate) SetNillablePoolCoinTypeID(u *uuid.UUID) *GoodUserCreate {
	if u != nil {
		guc.SetPoolCoinTypeID(*u)
	}
	return guc
}

// SetReadPageLink sets the "read_page_link" field.
func (guc *GoodUserCreate) SetReadPageLink(s string) *GoodUserCreate {
	guc.mutation.SetReadPageLink(s)
	return guc
}

// SetNillableReadPageLink sets the "read_page_link" field if the given value is not nil.
func (guc *GoodUserCreate) SetNillableReadPageLink(s *string) *GoodUserCreate {
	if s != nil {
		guc.SetReadPageLink(*s)
	}
	return guc
}

// SetID sets the "id" field.
func (guc *GoodUserCreate) SetID(u uint32) *GoodUserCreate {
	guc.mutation.SetID(u)
	return guc
}

// Mutation returns the GoodUserMutation object of the builder.
func (guc *GoodUserCreate) Mutation() *GoodUserMutation {
	return guc.mutation
}

// Save creates the GoodUser in the database.
func (guc *GoodUserCreate) Save(ctx context.Context) (*GoodUser, error) {
	var (
		err  error
		node *GoodUser
	)
	if err := guc.defaults(); err != nil {
		return nil, err
	}
	if len(guc.hooks) == 0 {
		if err = guc.check(); err != nil {
			return nil, err
		}
		node, err = guc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = guc.check(); err != nil {
				return nil, err
			}
			guc.mutation = mutation
			if node, err = guc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(guc.hooks) - 1; i >= 0; i-- {
			if guc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, guc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (guc *GoodUserCreate) SaveX(ctx context.Context) *GoodUser {
	v, err := guc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (guc *GoodUserCreate) Exec(ctx context.Context) error {
	_, err := guc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guc *GoodUserCreate) ExecX(ctx context.Context) {
	if err := guc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guc *GoodUserCreate) defaults() error {
	if _, ok := guc.mutation.CreatedAt(); !ok {
		if gooduser.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized gooduser.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := gooduser.DefaultCreatedAt()
		guc.mutation.SetCreatedAt(v)
	}
	if _, ok := guc.mutation.UpdatedAt(); !ok {
		if gooduser.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized gooduser.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := gooduser.DefaultUpdatedAt()
		guc.mutation.SetUpdatedAt(v)
	}
	if _, ok := guc.mutation.DeletedAt(); !ok {
		if gooduser.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized gooduser.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := gooduser.DefaultDeletedAt()
		guc.mutation.SetDeletedAt(v)
	}
	if _, ok := guc.mutation.EntID(); !ok {
		if gooduser.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized gooduser.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := gooduser.DefaultEntID()
		guc.mutation.SetEntID(v)
	}
	if _, ok := guc.mutation.RootUserID(); !ok {
		if gooduser.DefaultRootUserID == nil {
			return fmt.Errorf("ent: uninitialized gooduser.DefaultRootUserID (forgotten import ent/runtime?)")
		}
		v := gooduser.DefaultRootUserID()
		guc.mutation.SetRootUserID(v)
	}
	if _, ok := guc.mutation.Name(); !ok {
		v := gooduser.DefaultName
		guc.mutation.SetName(v)
	}
	if _, ok := guc.mutation.PoolCoinTypeID(); !ok {
		if gooduser.DefaultPoolCoinTypeID == nil {
			return fmt.Errorf("ent: uninitialized gooduser.DefaultPoolCoinTypeID (forgotten import ent/runtime?)")
		}
		v := gooduser.DefaultPoolCoinTypeID()
		guc.mutation.SetPoolCoinTypeID(v)
	}
	if _, ok := guc.mutation.ReadPageLink(); !ok {
		v := gooduser.DefaultReadPageLink
		guc.mutation.SetReadPageLink(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (guc *GoodUserCreate) check() error {
	if _, ok := guc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "GoodUser.created_at"`)}
	}
	if _, ok := guc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "GoodUser.updated_at"`)}
	}
	if _, ok := guc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "GoodUser.deleted_at"`)}
	}
	if _, ok := guc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "GoodUser.ent_id"`)}
	}
	return nil
}

func (guc *GoodUserCreate) sqlSave(ctx context.Context) (*GoodUser, error) {
	_node, _spec := guc.createSpec()
	if err := sqlgraph.CreateNode(ctx, guc.driver, _spec); err != nil {
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

func (guc *GoodUserCreate) createSpec() (*GoodUser, *sqlgraph.CreateSpec) {
	var (
		_node = &GoodUser{config: guc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: gooduser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: gooduser.FieldID,
			},
		}
	)
	_spec.OnConflict = guc.conflict
	if id, ok := guc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := guc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gooduser.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := guc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gooduser.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := guc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gooduser.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := guc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: gooduser.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := guc.mutation.RootUserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: gooduser.FieldRootUserID,
		})
		_node.RootUserID = value
	}
	if value, ok := guc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gooduser.FieldName,
		})
		_node.Name = value
	}
	if value, ok := guc.mutation.PoolCoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: gooduser.FieldPoolCoinTypeID,
		})
		_node.PoolCoinTypeID = value
	}
	if value, ok := guc.mutation.ReadPageLink(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gooduser.FieldReadPageLink,
		})
		_node.ReadPageLink = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GoodUser.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GoodUserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (guc *GoodUserCreate) OnConflict(opts ...sql.ConflictOption) *GoodUserUpsertOne {
	guc.conflict = opts
	return &GoodUserUpsertOne{
		create: guc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GoodUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (guc *GoodUserCreate) OnConflictColumns(columns ...string) *GoodUserUpsertOne {
	guc.conflict = append(guc.conflict, sql.ConflictColumns(columns...))
	return &GoodUserUpsertOne{
		create: guc,
	}
}

type (
	// GoodUserUpsertOne is the builder for "upsert"-ing
	//  one GoodUser node.
	GoodUserUpsertOne struct {
		create *GoodUserCreate
	}

	// GoodUserUpsert is the "OnConflict" setter.
	GoodUserUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *GoodUserUpsert) SetCreatedAt(v uint32) *GoodUserUpsert {
	u.Set(gooduser.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GoodUserUpsert) UpdateCreatedAt() *GoodUserUpsert {
	u.SetExcluded(gooduser.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *GoodUserUpsert) AddCreatedAt(v uint32) *GoodUserUpsert {
	u.Add(gooduser.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GoodUserUpsert) SetUpdatedAt(v uint32) *GoodUserUpsert {
	u.Set(gooduser.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GoodUserUpsert) UpdateUpdatedAt() *GoodUserUpsert {
	u.SetExcluded(gooduser.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *GoodUserUpsert) AddUpdatedAt(v uint32) *GoodUserUpsert {
	u.Add(gooduser.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *GoodUserUpsert) SetDeletedAt(v uint32) *GoodUserUpsert {
	u.Set(gooduser.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *GoodUserUpsert) UpdateDeletedAt() *GoodUserUpsert {
	u.SetExcluded(gooduser.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *GoodUserUpsert) AddDeletedAt(v uint32) *GoodUserUpsert {
	u.Add(gooduser.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *GoodUserUpsert) SetEntID(v uuid.UUID) *GoodUserUpsert {
	u.Set(gooduser.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *GoodUserUpsert) UpdateEntID() *GoodUserUpsert {
	u.SetExcluded(gooduser.FieldEntID)
	return u
}

// SetRootUserID sets the "root_user_id" field.
func (u *GoodUserUpsert) SetRootUserID(v uuid.UUID) *GoodUserUpsert {
	u.Set(gooduser.FieldRootUserID, v)
	return u
}

// UpdateRootUserID sets the "root_user_id" field to the value that was provided on create.
func (u *GoodUserUpsert) UpdateRootUserID() *GoodUserUpsert {
	u.SetExcluded(gooduser.FieldRootUserID)
	return u
}

// ClearRootUserID clears the value of the "root_user_id" field.
func (u *GoodUserUpsert) ClearRootUserID() *GoodUserUpsert {
	u.SetNull(gooduser.FieldRootUserID)
	return u
}

// SetName sets the "name" field.
func (u *GoodUserUpsert) SetName(v string) *GoodUserUpsert {
	u.Set(gooduser.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GoodUserUpsert) UpdateName() *GoodUserUpsert {
	u.SetExcluded(gooduser.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *GoodUserUpsert) ClearName() *GoodUserUpsert {
	u.SetNull(gooduser.FieldName)
	return u
}

// SetPoolCoinTypeID sets the "pool_coin_type_id" field.
func (u *GoodUserUpsert) SetPoolCoinTypeID(v uuid.UUID) *GoodUserUpsert {
	u.Set(gooduser.FieldPoolCoinTypeID, v)
	return u
}

// UpdatePoolCoinTypeID sets the "pool_coin_type_id" field to the value that was provided on create.
func (u *GoodUserUpsert) UpdatePoolCoinTypeID() *GoodUserUpsert {
	u.SetExcluded(gooduser.FieldPoolCoinTypeID)
	return u
}

// ClearPoolCoinTypeID clears the value of the "pool_coin_type_id" field.
func (u *GoodUserUpsert) ClearPoolCoinTypeID() *GoodUserUpsert {
	u.SetNull(gooduser.FieldPoolCoinTypeID)
	return u
}

// SetReadPageLink sets the "read_page_link" field.
func (u *GoodUserUpsert) SetReadPageLink(v string) *GoodUserUpsert {
	u.Set(gooduser.FieldReadPageLink, v)
	return u
}

// UpdateReadPageLink sets the "read_page_link" field to the value that was provided on create.
func (u *GoodUserUpsert) UpdateReadPageLink() *GoodUserUpsert {
	u.SetExcluded(gooduser.FieldReadPageLink)
	return u
}

// ClearReadPageLink clears the value of the "read_page_link" field.
func (u *GoodUserUpsert) ClearReadPageLink() *GoodUserUpsert {
	u.SetNull(gooduser.FieldReadPageLink)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.GoodUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(gooduser.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GoodUserUpsertOne) UpdateNewValues() *GoodUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(gooduser.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GoodUser.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GoodUserUpsertOne) Ignore() *GoodUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GoodUserUpsertOne) DoNothing() *GoodUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GoodUserCreate.OnConflict
// documentation for more info.
func (u *GoodUserUpsertOne) Update(set func(*GoodUserUpsert)) *GoodUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GoodUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *GoodUserUpsertOne) SetCreatedAt(v uint32) *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *GoodUserUpsertOne) AddCreatedAt(v uint32) *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GoodUserUpsertOne) UpdateCreatedAt() *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GoodUserUpsertOne) SetUpdatedAt(v uint32) *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *GoodUserUpsertOne) AddUpdatedAt(v uint32) *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GoodUserUpsertOne) UpdateUpdatedAt() *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *GoodUserUpsertOne) SetDeletedAt(v uint32) *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *GoodUserUpsertOne) AddDeletedAt(v uint32) *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *GoodUserUpsertOne) UpdateDeletedAt() *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *GoodUserUpsertOne) SetEntID(v uuid.UUID) *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *GoodUserUpsertOne) UpdateEntID() *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.UpdateEntID()
	})
}

// SetRootUserID sets the "root_user_id" field.
func (u *GoodUserUpsertOne) SetRootUserID(v uuid.UUID) *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.SetRootUserID(v)
	})
}

// UpdateRootUserID sets the "root_user_id" field to the value that was provided on create.
func (u *GoodUserUpsertOne) UpdateRootUserID() *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.UpdateRootUserID()
	})
}

// ClearRootUserID clears the value of the "root_user_id" field.
func (u *GoodUserUpsertOne) ClearRootUserID() *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.ClearRootUserID()
	})
}

// SetName sets the "name" field.
func (u *GoodUserUpsertOne) SetName(v string) *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GoodUserUpsertOne) UpdateName() *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *GoodUserUpsertOne) ClearName() *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.ClearName()
	})
}

// SetPoolCoinTypeID sets the "pool_coin_type_id" field.
func (u *GoodUserUpsertOne) SetPoolCoinTypeID(v uuid.UUID) *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.SetPoolCoinTypeID(v)
	})
}

// UpdatePoolCoinTypeID sets the "pool_coin_type_id" field to the value that was provided on create.
func (u *GoodUserUpsertOne) UpdatePoolCoinTypeID() *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.UpdatePoolCoinTypeID()
	})
}

// ClearPoolCoinTypeID clears the value of the "pool_coin_type_id" field.
func (u *GoodUserUpsertOne) ClearPoolCoinTypeID() *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.ClearPoolCoinTypeID()
	})
}

// SetReadPageLink sets the "read_page_link" field.
func (u *GoodUserUpsertOne) SetReadPageLink(v string) *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.SetReadPageLink(v)
	})
}

// UpdateReadPageLink sets the "read_page_link" field to the value that was provided on create.
func (u *GoodUserUpsertOne) UpdateReadPageLink() *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.UpdateReadPageLink()
	})
}

// ClearReadPageLink clears the value of the "read_page_link" field.
func (u *GoodUserUpsertOne) ClearReadPageLink() *GoodUserUpsertOne {
	return u.Update(func(s *GoodUserUpsert) {
		s.ClearReadPageLink()
	})
}

// Exec executes the query.
func (u *GoodUserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GoodUserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GoodUserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GoodUserUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GoodUserUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GoodUserCreateBulk is the builder for creating many GoodUser entities in bulk.
type GoodUserCreateBulk struct {
	config
	builders []*GoodUserCreate
	conflict []sql.ConflictOption
}

// Save creates the GoodUser entities in the database.
func (gucb *GoodUserCreateBulk) Save(ctx context.Context) ([]*GoodUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gucb.builders))
	nodes := make([]*GoodUser, len(gucb.builders))
	mutators := make([]Mutator, len(gucb.builders))
	for i := range gucb.builders {
		func(i int, root context.Context) {
			builder := gucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GoodUserMutation)
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
					_, err = mutators[i+1].Mutate(root, gucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gucb *GoodUserCreateBulk) SaveX(ctx context.Context) []*GoodUser {
	v, err := gucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gucb *GoodUserCreateBulk) Exec(ctx context.Context) error {
	_, err := gucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gucb *GoodUserCreateBulk) ExecX(ctx context.Context) {
	if err := gucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GoodUser.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GoodUserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (gucb *GoodUserCreateBulk) OnConflict(opts ...sql.ConflictOption) *GoodUserUpsertBulk {
	gucb.conflict = opts
	return &GoodUserUpsertBulk{
		create: gucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GoodUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gucb *GoodUserCreateBulk) OnConflictColumns(columns ...string) *GoodUserUpsertBulk {
	gucb.conflict = append(gucb.conflict, sql.ConflictColumns(columns...))
	return &GoodUserUpsertBulk{
		create: gucb,
	}
}

// GoodUserUpsertBulk is the builder for "upsert"-ing
// a bulk of GoodUser nodes.
type GoodUserUpsertBulk struct {
	create *GoodUserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GoodUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(gooduser.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GoodUserUpsertBulk) UpdateNewValues() *GoodUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(gooduser.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GoodUser.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GoodUserUpsertBulk) Ignore() *GoodUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GoodUserUpsertBulk) DoNothing() *GoodUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GoodUserCreateBulk.OnConflict
// documentation for more info.
func (u *GoodUserUpsertBulk) Update(set func(*GoodUserUpsert)) *GoodUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GoodUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *GoodUserUpsertBulk) SetCreatedAt(v uint32) *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *GoodUserUpsertBulk) AddCreatedAt(v uint32) *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GoodUserUpsertBulk) UpdateCreatedAt() *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GoodUserUpsertBulk) SetUpdatedAt(v uint32) *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *GoodUserUpsertBulk) AddUpdatedAt(v uint32) *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GoodUserUpsertBulk) UpdateUpdatedAt() *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *GoodUserUpsertBulk) SetDeletedAt(v uint32) *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *GoodUserUpsertBulk) AddDeletedAt(v uint32) *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *GoodUserUpsertBulk) UpdateDeletedAt() *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *GoodUserUpsertBulk) SetEntID(v uuid.UUID) *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *GoodUserUpsertBulk) UpdateEntID() *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.UpdateEntID()
	})
}

// SetRootUserID sets the "root_user_id" field.
func (u *GoodUserUpsertBulk) SetRootUserID(v uuid.UUID) *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.SetRootUserID(v)
	})
}

// UpdateRootUserID sets the "root_user_id" field to the value that was provided on create.
func (u *GoodUserUpsertBulk) UpdateRootUserID() *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.UpdateRootUserID()
	})
}

// ClearRootUserID clears the value of the "root_user_id" field.
func (u *GoodUserUpsertBulk) ClearRootUserID() *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.ClearRootUserID()
	})
}

// SetName sets the "name" field.
func (u *GoodUserUpsertBulk) SetName(v string) *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GoodUserUpsertBulk) UpdateName() *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *GoodUserUpsertBulk) ClearName() *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.ClearName()
	})
}

// SetPoolCoinTypeID sets the "pool_coin_type_id" field.
func (u *GoodUserUpsertBulk) SetPoolCoinTypeID(v uuid.UUID) *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.SetPoolCoinTypeID(v)
	})
}

// UpdatePoolCoinTypeID sets the "pool_coin_type_id" field to the value that was provided on create.
func (u *GoodUserUpsertBulk) UpdatePoolCoinTypeID() *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.UpdatePoolCoinTypeID()
	})
}

// ClearPoolCoinTypeID clears the value of the "pool_coin_type_id" field.
func (u *GoodUserUpsertBulk) ClearPoolCoinTypeID() *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.ClearPoolCoinTypeID()
	})
}

// SetReadPageLink sets the "read_page_link" field.
func (u *GoodUserUpsertBulk) SetReadPageLink(v string) *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.SetReadPageLink(v)
	})
}

// UpdateReadPageLink sets the "read_page_link" field to the value that was provided on create.
func (u *GoodUserUpsertBulk) UpdateReadPageLink() *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.UpdateReadPageLink()
	})
}

// ClearReadPageLink clears the value of the "read_page_link" field.
func (u *GoodUserUpsertBulk) ClearReadPageLink() *GoodUserUpsertBulk {
	return u.Update(func(s *GoodUserUpsert) {
		s.ClearReadPageLink()
	})
}

// Exec executes the query.
func (u *GoodUserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GoodUserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GoodUserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GoodUserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
