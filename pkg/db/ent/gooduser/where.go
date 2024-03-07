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

// GoodID applies equality check predicate on the "good_id" field. It's identical to GoodIDEQ.
func GoodID(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// AuthToken applies equality check predicate on the "auth_token" field. It's identical to AuthTokenEQ.
func AuthToken(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthToken), v))
	})
}

// Authed applies equality check predicate on the "authed" field. It's identical to AuthedEQ.
func Authed(v bool) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthed), v))
	})
}

// Start applies equality check predicate on the "start" field. It's identical to StartEQ.
func Start(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStart), v))
	})
}

// End applies equality check predicate on the "end" field. It's identical to EndEQ.
func End(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnd), v))
	})
}

// HashRate applies equality check predicate on the "hash_rate" field. It's identical to HashRateEQ.
func HashRate(v float64) predicate.GoodUser {
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

// GoodIDEQ applies the EQ predicate on the "good_id" field.
func GoodIDEQ(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// GoodIDNEQ applies the NEQ predicate on the "good_id" field.
func GoodIDNEQ(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoodID), v))
	})
}

// GoodIDIn applies the In predicate on the "good_id" field.
func GoodIDIn(vs ...uuid.UUID) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGoodID), v...))
	})
}

// GoodIDNotIn applies the NotIn predicate on the "good_id" field.
func GoodIDNotIn(vs ...uuid.UUID) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGoodID), v...))
	})
}

// GoodIDGT applies the GT predicate on the "good_id" field.
func GoodIDGT(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoodID), v))
	})
}

// GoodIDGTE applies the GTE predicate on the "good_id" field.
func GoodIDGTE(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoodID), v))
	})
}

// GoodIDLT applies the LT predicate on the "good_id" field.
func GoodIDLT(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoodID), v))
	})
}

// GoodIDLTE applies the LTE predicate on the "good_id" field.
func GoodIDLTE(v uuid.UUID) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoodID), v))
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

// AuthTokenEQ applies the EQ predicate on the "auth_token" field.
func AuthTokenEQ(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthToken), v))
	})
}

// AuthTokenNEQ applies the NEQ predicate on the "auth_token" field.
func AuthTokenNEQ(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAuthToken), v))
	})
}

// AuthTokenIn applies the In predicate on the "auth_token" field.
func AuthTokenIn(vs ...string) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAuthToken), v...))
	})
}

// AuthTokenNotIn applies the NotIn predicate on the "auth_token" field.
func AuthTokenNotIn(vs ...string) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAuthToken), v...))
	})
}

// AuthTokenGT applies the GT predicate on the "auth_token" field.
func AuthTokenGT(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAuthToken), v))
	})
}

// AuthTokenGTE applies the GTE predicate on the "auth_token" field.
func AuthTokenGTE(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAuthToken), v))
	})
}

// AuthTokenLT applies the LT predicate on the "auth_token" field.
func AuthTokenLT(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAuthToken), v))
	})
}

// AuthTokenLTE applies the LTE predicate on the "auth_token" field.
func AuthTokenLTE(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAuthToken), v))
	})
}

// AuthTokenContains applies the Contains predicate on the "auth_token" field.
func AuthTokenContains(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAuthToken), v))
	})
}

// AuthTokenHasPrefix applies the HasPrefix predicate on the "auth_token" field.
func AuthTokenHasPrefix(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAuthToken), v))
	})
}

// AuthTokenHasSuffix applies the HasSuffix predicate on the "auth_token" field.
func AuthTokenHasSuffix(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAuthToken), v))
	})
}

// AuthTokenEqualFold applies the EqualFold predicate on the "auth_token" field.
func AuthTokenEqualFold(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAuthToken), v))
	})
}

// AuthTokenContainsFold applies the ContainsFold predicate on the "auth_token" field.
func AuthTokenContainsFold(v string) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAuthToken), v))
	})
}

// AuthedEQ applies the EQ predicate on the "authed" field.
func AuthedEQ(v bool) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthed), v))
	})
}

// AuthedNEQ applies the NEQ predicate on the "authed" field.
func AuthedNEQ(v bool) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAuthed), v))
	})
}

// StartEQ applies the EQ predicate on the "start" field.
func StartEQ(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStart), v))
	})
}

// StartNEQ applies the NEQ predicate on the "start" field.
func StartNEQ(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStart), v))
	})
}

// StartIn applies the In predicate on the "start" field.
func StartIn(vs ...uint32) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStart), v...))
	})
}

// StartNotIn applies the NotIn predicate on the "start" field.
func StartNotIn(vs ...uint32) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStart), v...))
	})
}

// StartGT applies the GT predicate on the "start" field.
func StartGT(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStart), v))
	})
}

// StartGTE applies the GTE predicate on the "start" field.
func StartGTE(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStart), v))
	})
}

// StartLT applies the LT predicate on the "start" field.
func StartLT(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStart), v))
	})
}

// StartLTE applies the LTE predicate on the "start" field.
func StartLTE(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStart), v))
	})
}

// EndEQ applies the EQ predicate on the "end" field.
func EndEQ(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnd), v))
	})
}

// EndNEQ applies the NEQ predicate on the "end" field.
func EndNEQ(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnd), v))
	})
}

// EndIn applies the In predicate on the "end" field.
func EndIn(vs ...uint32) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEnd), v...))
	})
}

// EndNotIn applies the NotIn predicate on the "end" field.
func EndNotIn(vs ...uint32) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEnd), v...))
	})
}

// EndGT applies the GT predicate on the "end" field.
func EndGT(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEnd), v))
	})
}

// EndGTE applies the GTE predicate on the "end" field.
func EndGTE(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEnd), v))
	})
}

// EndLT applies the LT predicate on the "end" field.
func EndLT(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEnd), v))
	})
}

// EndLTE applies the LTE predicate on the "end" field.
func EndLTE(v uint32) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEnd), v))
	})
}

// HashRateEQ applies the EQ predicate on the "hash_rate" field.
func HashRateEQ(v float64) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHashRate), v))
	})
}

// HashRateNEQ applies the NEQ predicate on the "hash_rate" field.
func HashRateNEQ(v float64) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHashRate), v))
	})
}

// HashRateIn applies the In predicate on the "hash_rate" field.
func HashRateIn(vs ...float64) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldHashRate), v...))
	})
}

// HashRateNotIn applies the NotIn predicate on the "hash_rate" field.
func HashRateNotIn(vs ...float64) predicate.GoodUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldHashRate), v...))
	})
}

// HashRateGT applies the GT predicate on the "hash_rate" field.
func HashRateGT(v float64) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHashRate), v))
	})
}

// HashRateGTE applies the GTE predicate on the "hash_rate" field.
func HashRateGTE(v float64) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHashRate), v))
	})
}

// HashRateLT applies the LT predicate on the "hash_rate" field.
func HashRateLT(v float64) predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHashRate), v))
	})
}

// HashRateLTE applies the LTE predicate on the "hash_rate" field.
func HashRateLTE(v float64) predicate.GoodUser {
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

// ReadPageLinkIsNil applies the IsNil predicate on the "read_page_link" field.
func ReadPageLinkIsNil() predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldReadPageLink)))
	})
}

// ReadPageLinkNotNil applies the NotNil predicate on the "read_page_link" field.
func ReadPageLinkNotNil() predicate.GoodUser {
	return predicate.GoodUser(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldReadPageLink)))
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
