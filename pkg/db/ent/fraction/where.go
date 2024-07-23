// Code generated by ent, DO NOT EDIT.

package fraction

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// OrderUserID applies equality check predicate on the "order_user_id" field. It's identical to OrderUserIDEQ.
func OrderUserID(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderUserID), v))
	})
}

// CoinTypeID applies equality check predicate on the "coin_type_id" field. It's identical to CoinTypeIDEQ.
func CoinTypeID(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// WithdrawState applies equality check predicate on the "withdraw_state" field. It's identical to WithdrawStateEQ.
func WithdrawState(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWithdrawState), v))
	})
}

// WithdrawAt applies equality check predicate on the "withdraw_at" field. It's identical to WithdrawAtEQ.
func WithdrawAt(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWithdrawAt), v))
	})
}

// PromisePayAt applies equality check predicate on the "promise_pay_at" field. It's identical to PromisePayAtEQ.
func PromisePayAt(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPromisePayAt), v))
	})
}

// Msg applies equality check predicate on the "msg" field. It's identical to MsgEQ.
func Msg(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMsg), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppID)))
	})
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppID)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// OrderUserIDEQ applies the EQ predicate on the "order_user_id" field.
func OrderUserIDEQ(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderUserID), v))
	})
}

// OrderUserIDNEQ applies the NEQ predicate on the "order_user_id" field.
func OrderUserIDNEQ(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderUserID), v))
	})
}

// OrderUserIDIn applies the In predicate on the "order_user_id" field.
func OrderUserIDIn(vs ...uuid.UUID) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderUserID), v...))
	})
}

// OrderUserIDNotIn applies the NotIn predicate on the "order_user_id" field.
func OrderUserIDNotIn(vs ...uuid.UUID) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderUserID), v...))
	})
}

// OrderUserIDGT applies the GT predicate on the "order_user_id" field.
func OrderUserIDGT(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderUserID), v))
	})
}

// OrderUserIDGTE applies the GTE predicate on the "order_user_id" field.
func OrderUserIDGTE(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderUserID), v))
	})
}

// OrderUserIDLT applies the LT predicate on the "order_user_id" field.
func OrderUserIDLT(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderUserID), v))
	})
}

// OrderUserIDLTE applies the LTE predicate on the "order_user_id" field.
func OrderUserIDLTE(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderUserID), v))
	})
}

// OrderUserIDIsNil applies the IsNil predicate on the "order_user_id" field.
func OrderUserIDIsNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrderUserID)))
	})
}

// OrderUserIDNotNil applies the NotNil predicate on the "order_user_id" field.
func OrderUserIDNotNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrderUserID)))
	})
}

// CoinTypeIDEQ applies the EQ predicate on the "coin_type_id" field.
func CoinTypeIDEQ(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDNEQ applies the NEQ predicate on the "coin_type_id" field.
func CoinTypeIDNEQ(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIn applies the In predicate on the "coin_type_id" field.
func CoinTypeIDIn(vs ...uuid.UUID) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDNotIn applies the NotIn predicate on the "coin_type_id" field.
func CoinTypeIDNotIn(vs ...uuid.UUID) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDGT applies the GT predicate on the "coin_type_id" field.
func CoinTypeIDGT(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDGTE applies the GTE predicate on the "coin_type_id" field.
func CoinTypeIDGTE(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLT applies the LT predicate on the "coin_type_id" field.
func CoinTypeIDLT(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLTE applies the LTE predicate on the "coin_type_id" field.
func CoinTypeIDLTE(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIsNil applies the IsNil predicate on the "coin_type_id" field.
func CoinTypeIDIsNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCoinTypeID)))
	})
}

// CoinTypeIDNotNil applies the NotNil predicate on the "coin_type_id" field.
func CoinTypeIDNotNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCoinTypeID)))
	})
}

// WithdrawStateEQ applies the EQ predicate on the "withdraw_state" field.
func WithdrawStateEQ(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWithdrawState), v))
	})
}

// WithdrawStateNEQ applies the NEQ predicate on the "withdraw_state" field.
func WithdrawStateNEQ(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWithdrawState), v))
	})
}

// WithdrawStateIn applies the In predicate on the "withdraw_state" field.
func WithdrawStateIn(vs ...string) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldWithdrawState), v...))
	})
}

// WithdrawStateNotIn applies the NotIn predicate on the "withdraw_state" field.
func WithdrawStateNotIn(vs ...string) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldWithdrawState), v...))
	})
}

// WithdrawStateGT applies the GT predicate on the "withdraw_state" field.
func WithdrawStateGT(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWithdrawState), v))
	})
}

// WithdrawStateGTE applies the GTE predicate on the "withdraw_state" field.
func WithdrawStateGTE(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWithdrawState), v))
	})
}

// WithdrawStateLT applies the LT predicate on the "withdraw_state" field.
func WithdrawStateLT(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWithdrawState), v))
	})
}

// WithdrawStateLTE applies the LTE predicate on the "withdraw_state" field.
func WithdrawStateLTE(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWithdrawState), v))
	})
}

// WithdrawStateContains applies the Contains predicate on the "withdraw_state" field.
func WithdrawStateContains(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldWithdrawState), v))
	})
}

// WithdrawStateHasPrefix applies the HasPrefix predicate on the "withdraw_state" field.
func WithdrawStateHasPrefix(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldWithdrawState), v))
	})
}

// WithdrawStateHasSuffix applies the HasSuffix predicate on the "withdraw_state" field.
func WithdrawStateHasSuffix(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldWithdrawState), v))
	})
}

// WithdrawStateIsNil applies the IsNil predicate on the "withdraw_state" field.
func WithdrawStateIsNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldWithdrawState)))
	})
}

// WithdrawStateNotNil applies the NotNil predicate on the "withdraw_state" field.
func WithdrawStateNotNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldWithdrawState)))
	})
}

// WithdrawStateEqualFold applies the EqualFold predicate on the "withdraw_state" field.
func WithdrawStateEqualFold(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldWithdrawState), v))
	})
}

// WithdrawStateContainsFold applies the ContainsFold predicate on the "withdraw_state" field.
func WithdrawStateContainsFold(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldWithdrawState), v))
	})
}

// WithdrawAtEQ applies the EQ predicate on the "withdraw_at" field.
func WithdrawAtEQ(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWithdrawAt), v))
	})
}

// WithdrawAtNEQ applies the NEQ predicate on the "withdraw_at" field.
func WithdrawAtNEQ(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWithdrawAt), v))
	})
}

// WithdrawAtIn applies the In predicate on the "withdraw_at" field.
func WithdrawAtIn(vs ...uint32) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldWithdrawAt), v...))
	})
}

// WithdrawAtNotIn applies the NotIn predicate on the "withdraw_at" field.
func WithdrawAtNotIn(vs ...uint32) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldWithdrawAt), v...))
	})
}

// WithdrawAtGT applies the GT predicate on the "withdraw_at" field.
func WithdrawAtGT(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWithdrawAt), v))
	})
}

// WithdrawAtGTE applies the GTE predicate on the "withdraw_at" field.
func WithdrawAtGTE(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWithdrawAt), v))
	})
}

// WithdrawAtLT applies the LT predicate on the "withdraw_at" field.
func WithdrawAtLT(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWithdrawAt), v))
	})
}

// WithdrawAtLTE applies the LTE predicate on the "withdraw_at" field.
func WithdrawAtLTE(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWithdrawAt), v))
	})
}

// WithdrawAtIsNil applies the IsNil predicate on the "withdraw_at" field.
func WithdrawAtIsNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldWithdrawAt)))
	})
}

// WithdrawAtNotNil applies the NotNil predicate on the "withdraw_at" field.
func WithdrawAtNotNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldWithdrawAt)))
	})
}

// PromisePayAtEQ applies the EQ predicate on the "promise_pay_at" field.
func PromisePayAtEQ(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPromisePayAt), v))
	})
}

// PromisePayAtNEQ applies the NEQ predicate on the "promise_pay_at" field.
func PromisePayAtNEQ(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPromisePayAt), v))
	})
}

// PromisePayAtIn applies the In predicate on the "promise_pay_at" field.
func PromisePayAtIn(vs ...uint32) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPromisePayAt), v...))
	})
}

// PromisePayAtNotIn applies the NotIn predicate on the "promise_pay_at" field.
func PromisePayAtNotIn(vs ...uint32) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPromisePayAt), v...))
	})
}

// PromisePayAtGT applies the GT predicate on the "promise_pay_at" field.
func PromisePayAtGT(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPromisePayAt), v))
	})
}

// PromisePayAtGTE applies the GTE predicate on the "promise_pay_at" field.
func PromisePayAtGTE(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPromisePayAt), v))
	})
}

// PromisePayAtLT applies the LT predicate on the "promise_pay_at" field.
func PromisePayAtLT(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPromisePayAt), v))
	})
}

// PromisePayAtLTE applies the LTE predicate on the "promise_pay_at" field.
func PromisePayAtLTE(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPromisePayAt), v))
	})
}

// PromisePayAtIsNil applies the IsNil predicate on the "promise_pay_at" field.
func PromisePayAtIsNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPromisePayAt)))
	})
}

// PromisePayAtNotNil applies the NotNil predicate on the "promise_pay_at" field.
func PromisePayAtNotNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPromisePayAt)))
	})
}

// MsgEQ applies the EQ predicate on the "msg" field.
func MsgEQ(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMsg), v))
	})
}

// MsgNEQ applies the NEQ predicate on the "msg" field.
func MsgNEQ(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMsg), v))
	})
}

// MsgIn applies the In predicate on the "msg" field.
func MsgIn(vs ...string) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMsg), v...))
	})
}

// MsgNotIn applies the NotIn predicate on the "msg" field.
func MsgNotIn(vs ...string) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMsg), v...))
	})
}

// MsgGT applies the GT predicate on the "msg" field.
func MsgGT(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMsg), v))
	})
}

// MsgGTE applies the GTE predicate on the "msg" field.
func MsgGTE(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMsg), v))
	})
}

// MsgLT applies the LT predicate on the "msg" field.
func MsgLT(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMsg), v))
	})
}

// MsgLTE applies the LTE predicate on the "msg" field.
func MsgLTE(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMsg), v))
	})
}

// MsgContains applies the Contains predicate on the "msg" field.
func MsgContains(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMsg), v))
	})
}

// MsgHasPrefix applies the HasPrefix predicate on the "msg" field.
func MsgHasPrefix(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMsg), v))
	})
}

// MsgHasSuffix applies the HasSuffix predicate on the "msg" field.
func MsgHasSuffix(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMsg), v))
	})
}

// MsgIsNil applies the IsNil predicate on the "msg" field.
func MsgIsNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMsg)))
	})
}

// MsgNotNil applies the NotNil predicate on the "msg" field.
func MsgNotNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMsg)))
	})
}

// MsgEqualFold applies the EqualFold predicate on the "msg" field.
func MsgEqualFold(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMsg), v))
	})
}

// MsgContainsFold applies the ContainsFold predicate on the "msg" field.
func MsgContainsFold(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMsg), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Fraction) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Fraction) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
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
func Not(p predicate.Fraction) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		p(s.Not())
	})
}
