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

// OrderMininguserID applies equality check predicate on the "order_mininguser_id" field. It's identical to OrderMininguserIDEQ.
func OrderMininguserID(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderMininguserID), v))
	})
}

// WithdrawState applies equality check predicate on the "withdraw_state" field. It's identical to WithdrawStateEQ.
func WithdrawState(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWithdrawState), v))
	})
}

// WithdrawTime applies equality check predicate on the "withdraw_time" field. It's identical to WithdrawTimeEQ.
func WithdrawTime(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWithdrawTime), v))
	})
}

// PayTime applies equality check predicate on the "pay_time" field. It's identical to PayTimeEQ.
func PayTime(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPayTime), v))
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

// OrderMininguserIDEQ applies the EQ predicate on the "order_mininguser_id" field.
func OrderMininguserIDEQ(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderMininguserID), v))
	})
}

// OrderMininguserIDNEQ applies the NEQ predicate on the "order_mininguser_id" field.
func OrderMininguserIDNEQ(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderMininguserID), v))
	})
}

// OrderMininguserIDIn applies the In predicate on the "order_mininguser_id" field.
func OrderMininguserIDIn(vs ...uuid.UUID) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderMininguserID), v...))
	})
}

// OrderMininguserIDNotIn applies the NotIn predicate on the "order_mininguser_id" field.
func OrderMininguserIDNotIn(vs ...uuid.UUID) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderMininguserID), v...))
	})
}

// OrderMininguserIDGT applies the GT predicate on the "order_mininguser_id" field.
func OrderMininguserIDGT(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderMininguserID), v))
	})
}

// OrderMininguserIDGTE applies the GTE predicate on the "order_mininguser_id" field.
func OrderMininguserIDGTE(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderMininguserID), v))
	})
}

// OrderMininguserIDLT applies the LT predicate on the "order_mininguser_id" field.
func OrderMininguserIDLT(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderMininguserID), v))
	})
}

// OrderMininguserIDLTE applies the LTE predicate on the "order_mininguser_id" field.
func OrderMininguserIDLTE(v uuid.UUID) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderMininguserID), v))
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

// WithdrawTimeEQ applies the EQ predicate on the "withdraw_time" field.
func WithdrawTimeEQ(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWithdrawTime), v))
	})
}

// WithdrawTimeNEQ applies the NEQ predicate on the "withdraw_time" field.
func WithdrawTimeNEQ(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWithdrawTime), v))
	})
}

// WithdrawTimeIn applies the In predicate on the "withdraw_time" field.
func WithdrawTimeIn(vs ...string) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldWithdrawTime), v...))
	})
}

// WithdrawTimeNotIn applies the NotIn predicate on the "withdraw_time" field.
func WithdrawTimeNotIn(vs ...string) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldWithdrawTime), v...))
	})
}

// WithdrawTimeGT applies the GT predicate on the "withdraw_time" field.
func WithdrawTimeGT(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWithdrawTime), v))
	})
}

// WithdrawTimeGTE applies the GTE predicate on the "withdraw_time" field.
func WithdrawTimeGTE(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWithdrawTime), v))
	})
}

// WithdrawTimeLT applies the LT predicate on the "withdraw_time" field.
func WithdrawTimeLT(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWithdrawTime), v))
	})
}

// WithdrawTimeLTE applies the LTE predicate on the "withdraw_time" field.
func WithdrawTimeLTE(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWithdrawTime), v))
	})
}

// WithdrawTimeContains applies the Contains predicate on the "withdraw_time" field.
func WithdrawTimeContains(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldWithdrawTime), v))
	})
}

// WithdrawTimeHasPrefix applies the HasPrefix predicate on the "withdraw_time" field.
func WithdrawTimeHasPrefix(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldWithdrawTime), v))
	})
}

// WithdrawTimeHasSuffix applies the HasSuffix predicate on the "withdraw_time" field.
func WithdrawTimeHasSuffix(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldWithdrawTime), v))
	})
}

// WithdrawTimeEqualFold applies the EqualFold predicate on the "withdraw_time" field.
func WithdrawTimeEqualFold(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldWithdrawTime), v))
	})
}

// WithdrawTimeContainsFold applies the ContainsFold predicate on the "withdraw_time" field.
func WithdrawTimeContainsFold(v string) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldWithdrawTime), v))
	})
}

// PayTimeEQ applies the EQ predicate on the "pay_time" field.
func PayTimeEQ(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPayTime), v))
	})
}

// PayTimeNEQ applies the NEQ predicate on the "pay_time" field.
func PayTimeNEQ(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPayTime), v))
	})
}

// PayTimeIn applies the In predicate on the "pay_time" field.
func PayTimeIn(vs ...uint32) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPayTime), v...))
	})
}

// PayTimeNotIn applies the NotIn predicate on the "pay_time" field.
func PayTimeNotIn(vs ...uint32) predicate.Fraction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPayTime), v...))
	})
}

// PayTimeGT applies the GT predicate on the "pay_time" field.
func PayTimeGT(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPayTime), v))
	})
}

// PayTimeGTE applies the GTE predicate on the "pay_time" field.
func PayTimeGTE(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPayTime), v))
	})
}

// PayTimeLT applies the LT predicate on the "pay_time" field.
func PayTimeLT(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPayTime), v))
	})
}

// PayTimeLTE applies the LTE predicate on the "pay_time" field.
func PayTimeLTE(v uint32) predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPayTime), v))
	})
}

// PayTimeIsNil applies the IsNil predicate on the "pay_time" field.
func PayTimeIsNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPayTime)))
	})
}

// PayTimeNotNil applies the NotNil predicate on the "pay_time" field.
func PayTimeNotNil() predicate.Fraction {
	return predicate.Fraction(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPayTime)))
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
