// Code generated by ent, DO NOT EDIT.

package gooduser

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// RootUserID applies equality check predicate on the "root_user_id" field. It's identical to RootUserIDEQ.
func RootUserID(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRootUserID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// MiningpoolType applies equality check predicate on the "miningpool_type" field. It's identical to MiningpoolTypeEQ.
func MiningpoolType(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMiningpoolType), v))
	})
}

// CoinType applies equality check predicate on the "coin_type" field. It's identical to CoinTypeEQ.
func CoinType(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinType), v))
	})
}

// HashRate applies equality check predicate on the "hash_rate" field. It's identical to HashRateEQ.
func HashRate(v float32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHashRate), v))
	})
}

// ReadPageLink applies equality check predicate on the "read_page_link" field. It's identical to ReadPageLinkEQ.
func ReadPageLink(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReadPageLink), v))
	})
}

// RevenueType applies equality check predicate on the "revenue_type" field. It's identical to RevenueTypeEQ.
func RevenueType(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRevenueType), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// RootUserIDEQ applies the EQ predicate on the "root_user_id" field.
func RootUserIDEQ(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRootUserID), v))
	})
}

// RootUserIDNEQ applies the NEQ predicate on the "root_user_id" field.
func RootUserIDNEQ(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRootUserID), v))
	})
}

// RootUserIDIn applies the In predicate on the "root_user_id" field.
func RootUserIDIn(vs ...uuid.UUID) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRootUserID), v...))
	})
}

// RootUserIDNotIn applies the NotIn predicate on the "root_user_id" field.
func RootUserIDNotIn(vs ...uuid.UUID) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRootUserID), v...))
	})
}

// RootUserIDGT applies the GT predicate on the "root_user_id" field.
func RootUserIDGT(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRootUserID), v))
	})
}

// RootUserIDGTE applies the GTE predicate on the "root_user_id" field.
func RootUserIDGTE(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRootUserID), v))
	})
}

// RootUserIDLT applies the LT predicate on the "root_user_id" field.
func RootUserIDLT(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRootUserID), v))
	})
}

// RootUserIDLTE applies the LTE predicate on the "root_user_id" field.
func RootUserIDLTE(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRootUserID), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// MiningpoolTypeEQ applies the EQ predicate on the "miningpool_type" field.
func MiningpoolTypeEQ(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeNEQ applies the NEQ predicate on the "miningpool_type" field.
func MiningpoolTypeNEQ(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeIn applies the In predicate on the "miningpool_type" field.
func MiningpoolTypeIn(vs ...string) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMiningpoolType), v...))
	})
}

// MiningpoolTypeNotIn applies the NotIn predicate on the "miningpool_type" field.
func MiningpoolTypeNotIn(vs ...string) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMiningpoolType), v...))
	})
}

// MiningpoolTypeGT applies the GT predicate on the "miningpool_type" field.
func MiningpoolTypeGT(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeGTE applies the GTE predicate on the "miningpool_type" field.
func MiningpoolTypeGTE(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeLT applies the LT predicate on the "miningpool_type" field.
func MiningpoolTypeLT(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeLTE applies the LTE predicate on the "miningpool_type" field.
func MiningpoolTypeLTE(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeContains applies the Contains predicate on the "miningpool_type" field.
func MiningpoolTypeContains(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeHasPrefix applies the HasPrefix predicate on the "miningpool_type" field.
func MiningpoolTypeHasPrefix(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeHasSuffix applies the HasSuffix predicate on the "miningpool_type" field.
func MiningpoolTypeHasSuffix(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeEqualFold applies the EqualFold predicate on the "miningpool_type" field.
func MiningpoolTypeEqualFold(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeContainsFold applies the ContainsFold predicate on the "miningpool_type" field.
func MiningpoolTypeContainsFold(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMiningpoolType), v))
	})
}

// CoinTypeEQ applies the EQ predicate on the "coin_type" field.
func CoinTypeEQ(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinType), v))
	})
}

// CoinTypeNEQ applies the NEQ predicate on the "coin_type" field.
func CoinTypeNEQ(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinType), v))
	})
}

// CoinTypeIn applies the In predicate on the "coin_type" field.
func CoinTypeIn(vs ...string) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCoinType), v...))
	})
}

// CoinTypeNotIn applies the NotIn predicate on the "coin_type" field.
func CoinTypeNotIn(vs ...string) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCoinType), v...))
	})
}

// CoinTypeGT applies the GT predicate on the "coin_type" field.
func CoinTypeGT(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinType), v))
	})
}

// CoinTypeGTE applies the GTE predicate on the "coin_type" field.
func CoinTypeGTE(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinType), v))
	})
}

// CoinTypeLT applies the LT predicate on the "coin_type" field.
func CoinTypeLT(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinType), v))
	})
}

// CoinTypeLTE applies the LTE predicate on the "coin_type" field.
func CoinTypeLTE(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinType), v))
	})
}

// CoinTypeContains applies the Contains predicate on the "coin_type" field.
func CoinTypeContains(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCoinType), v))
	})
}

// CoinTypeHasPrefix applies the HasPrefix predicate on the "coin_type" field.
func CoinTypeHasPrefix(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCoinType), v))
	})
}

// CoinTypeHasSuffix applies the HasSuffix predicate on the "coin_type" field.
func CoinTypeHasSuffix(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCoinType), v))
	})
}

// CoinTypeEqualFold applies the EqualFold predicate on the "coin_type" field.
func CoinTypeEqualFold(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCoinType), v))
	})
}

// CoinTypeContainsFold applies the ContainsFold predicate on the "coin_type" field.
func CoinTypeContainsFold(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCoinType), v))
	})
}

// HashRateEQ applies the EQ predicate on the "hash_rate" field.
func HashRateEQ(v float32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHashRate), v))
	})
}

// HashRateNEQ applies the NEQ predicate on the "hash_rate" field.
func HashRateNEQ(v float32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHashRate), v))
	})
}

// HashRateIn applies the In predicate on the "hash_rate" field.
func HashRateIn(vs ...float32) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldHashRate), v...))
	})
}

// HashRateNotIn applies the NotIn predicate on the "hash_rate" field.
func HashRateNotIn(vs ...float32) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldHashRate), v...))
	})
}

// HashRateGT applies the GT predicate on the "hash_rate" field.
func HashRateGT(v float32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHashRate), v))
	})
}

// HashRateGTE applies the GTE predicate on the "hash_rate" field.
func HashRateGTE(v float32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHashRate), v))
	})
}

// HashRateLT applies the LT predicate on the "hash_rate" field.
func HashRateLT(v float32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHashRate), v))
	})
}

// HashRateLTE applies the LTE predicate on the "hash_rate" field.
func HashRateLTE(v float32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHashRate), v))
	})
}

// HashRateIsNil applies the IsNil predicate on the "hash_rate" field.
func HashRateIsNil() predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHashRate)))
	})
}

// HashRateNotNil applies the NotNil predicate on the "hash_rate" field.
func HashRateNotNil() predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHashRate)))
	})
}

// ReadPageLinkEQ applies the EQ predicate on the "read_page_link" field.
func ReadPageLinkEQ(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReadPageLink), v))
	})
}

// ReadPageLinkNEQ applies the NEQ predicate on the "read_page_link" field.
func ReadPageLinkNEQ(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReadPageLink), v))
	})
}

// ReadPageLinkIn applies the In predicate on the "read_page_link" field.
func ReadPageLinkIn(vs ...string) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldReadPageLink), v...))
	})
}

// ReadPageLinkNotIn applies the NotIn predicate on the "read_page_link" field.
func ReadPageLinkNotIn(vs ...string) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldReadPageLink), v...))
	})
}

// ReadPageLinkGT applies the GT predicate on the "read_page_link" field.
func ReadPageLinkGT(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReadPageLink), v))
	})
}

// ReadPageLinkGTE applies the GTE predicate on the "read_page_link" field.
func ReadPageLinkGTE(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReadPageLink), v))
	})
}

// ReadPageLinkLT applies the LT predicate on the "read_page_link" field.
func ReadPageLinkLT(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReadPageLink), v))
	})
}

// ReadPageLinkLTE applies the LTE predicate on the "read_page_link" field.
func ReadPageLinkLTE(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReadPageLink), v))
	})
}

// ReadPageLinkContains applies the Contains predicate on the "read_page_link" field.
func ReadPageLinkContains(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldReadPageLink), v))
	})
}

// ReadPageLinkHasPrefix applies the HasPrefix predicate on the "read_page_link" field.
func ReadPageLinkHasPrefix(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldReadPageLink), v))
	})
}

// ReadPageLinkHasSuffix applies the HasSuffix predicate on the "read_page_link" field.
func ReadPageLinkHasSuffix(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldReadPageLink), v))
	})
}

// ReadPageLinkEqualFold applies the EqualFold predicate on the "read_page_link" field.
func ReadPageLinkEqualFold(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldReadPageLink), v))
	})
}

// ReadPageLinkContainsFold applies the ContainsFold predicate on the "read_page_link" field.
func ReadPageLinkContainsFold(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldReadPageLink), v))
	})
}

// RevenueTypeEQ applies the EQ predicate on the "revenue_type" field.
func RevenueTypeEQ(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRevenueType), v))
	})
}

// RevenueTypeNEQ applies the NEQ predicate on the "revenue_type" field.
func RevenueTypeNEQ(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRevenueType), v))
	})
}

// RevenueTypeIn applies the In predicate on the "revenue_type" field.
func RevenueTypeIn(vs ...string) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRevenueType), v...))
	})
}

// RevenueTypeNotIn applies the NotIn predicate on the "revenue_type" field.
func RevenueTypeNotIn(vs ...string) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRevenueType), v...))
	})
}

// RevenueTypeGT applies the GT predicate on the "revenue_type" field.
func RevenueTypeGT(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRevenueType), v))
	})
}

// RevenueTypeGTE applies the GTE predicate on the "revenue_type" field.
func RevenueTypeGTE(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRevenueType), v))
	})
}

// RevenueTypeLT applies the LT predicate on the "revenue_type" field.
func RevenueTypeLT(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRevenueType), v))
	})
}

// RevenueTypeLTE applies the LTE predicate on the "revenue_type" field.
func RevenueTypeLTE(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRevenueType), v))
	})
}

// RevenueTypeContains applies the Contains predicate on the "revenue_type" field.
func RevenueTypeContains(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRevenueType), v))
	})
}

// RevenueTypeHasPrefix applies the HasPrefix predicate on the "revenue_type" field.
func RevenueTypeHasPrefix(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRevenueType), v))
	})
}

// RevenueTypeHasSuffix applies the HasSuffix predicate on the "revenue_type" field.
func RevenueTypeHasSuffix(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRevenueType), v))
	})
}

// RevenueTypeEqualFold applies the EqualFold predicate on the "revenue_type" field.
func RevenueTypeEqualFold(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRevenueType), v))
	})
}

// RevenueTypeContainsFold applies the ContainsFold predicate on the "revenue_type" field.
func RevenueTypeContainsFold(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRevenueType), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GoodUser) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GoodUser) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GoodUser) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		p(s.Not())
	})
}
