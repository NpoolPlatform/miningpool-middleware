// Code generated by ent, DO NOT EDIT.

package coin

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// MiningpoolType applies equality check predicate on the "miningpool_type" field. It's identical to MiningpoolTypeEQ.
func MiningpoolType(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMiningpoolType), v))
	})
}

// CoinType applies equality check predicate on the "coin_type" field. It's identical to CoinTypeEQ.
func CoinType(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinType), v))
	})
}

// FeeRate applies equality check predicate on the "fee_rate" field. It's identical to FeeRateEQ.
func FeeRate(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeeRate), v))
	})
}

// FixedRevenueAble applies equality check predicate on the "fixed_revenue_able" field. It's identical to FixedRevenueAbleEQ.
func FixedRevenueAble(v bool) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFixedRevenueAble), v))
	})
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemark), v))
	})
}

// Threshold applies equality check predicate on the "threshold" field. It's identical to ThresholdEQ.
func Threshold(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThreshold), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// MiningpoolTypeEQ applies the EQ predicate on the "miningpool_type" field.
func MiningpoolTypeEQ(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeNEQ applies the NEQ predicate on the "miningpool_type" field.
func MiningpoolTypeNEQ(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeIn applies the In predicate on the "miningpool_type" field.
func MiningpoolTypeIn(vs ...string) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMiningpoolType), v...))
	})
}

// MiningpoolTypeNotIn applies the NotIn predicate on the "miningpool_type" field.
func MiningpoolTypeNotIn(vs ...string) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMiningpoolType), v...))
	})
}

// MiningpoolTypeGT applies the GT predicate on the "miningpool_type" field.
func MiningpoolTypeGT(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeGTE applies the GTE predicate on the "miningpool_type" field.
func MiningpoolTypeGTE(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeLT applies the LT predicate on the "miningpool_type" field.
func MiningpoolTypeLT(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeLTE applies the LTE predicate on the "miningpool_type" field.
func MiningpoolTypeLTE(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeContains applies the Contains predicate on the "miningpool_type" field.
func MiningpoolTypeContains(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeHasPrefix applies the HasPrefix predicate on the "miningpool_type" field.
func MiningpoolTypeHasPrefix(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeHasSuffix applies the HasSuffix predicate on the "miningpool_type" field.
func MiningpoolTypeHasSuffix(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeIsNil applies the IsNil predicate on the "miningpool_type" field.
func MiningpoolTypeIsNil() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMiningpoolType)))
	})
}

// MiningpoolTypeNotNil applies the NotNil predicate on the "miningpool_type" field.
func MiningpoolTypeNotNil() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMiningpoolType)))
	})
}

// MiningpoolTypeEqualFold applies the EqualFold predicate on the "miningpool_type" field.
func MiningpoolTypeEqualFold(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMiningpoolType), v))
	})
}

// MiningpoolTypeContainsFold applies the ContainsFold predicate on the "miningpool_type" field.
func MiningpoolTypeContainsFold(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMiningpoolType), v))
	})
}

// CoinTypeEQ applies the EQ predicate on the "coin_type" field.
func CoinTypeEQ(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinType), v))
	})
}

// CoinTypeNEQ applies the NEQ predicate on the "coin_type" field.
func CoinTypeNEQ(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinType), v))
	})
}

// CoinTypeIn applies the In predicate on the "coin_type" field.
func CoinTypeIn(vs ...string) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCoinType), v...))
	})
}

// CoinTypeNotIn applies the NotIn predicate on the "coin_type" field.
func CoinTypeNotIn(vs ...string) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCoinType), v...))
	})
}

// CoinTypeGT applies the GT predicate on the "coin_type" field.
func CoinTypeGT(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinType), v))
	})
}

// CoinTypeGTE applies the GTE predicate on the "coin_type" field.
func CoinTypeGTE(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinType), v))
	})
}

// CoinTypeLT applies the LT predicate on the "coin_type" field.
func CoinTypeLT(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinType), v))
	})
}

// CoinTypeLTE applies the LTE predicate on the "coin_type" field.
func CoinTypeLTE(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinType), v))
	})
}

// CoinTypeContains applies the Contains predicate on the "coin_type" field.
func CoinTypeContains(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCoinType), v))
	})
}

// CoinTypeHasPrefix applies the HasPrefix predicate on the "coin_type" field.
func CoinTypeHasPrefix(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCoinType), v))
	})
}

// CoinTypeHasSuffix applies the HasSuffix predicate on the "coin_type" field.
func CoinTypeHasSuffix(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCoinType), v))
	})
}

// CoinTypeIsNil applies the IsNil predicate on the "coin_type" field.
func CoinTypeIsNil() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCoinType)))
	})
}

// CoinTypeNotNil applies the NotNil predicate on the "coin_type" field.
func CoinTypeNotNil() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCoinType)))
	})
}

// CoinTypeEqualFold applies the EqualFold predicate on the "coin_type" field.
func CoinTypeEqualFold(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCoinType), v))
	})
}

// CoinTypeContainsFold applies the ContainsFold predicate on the "coin_type" field.
func CoinTypeContainsFold(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCoinType), v))
	})
}

// RevenueTypesIsNil applies the IsNil predicate on the "revenue_types" field.
func RevenueTypesIsNil() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRevenueTypes)))
	})
}

// RevenueTypesNotNil applies the NotNil predicate on the "revenue_types" field.
func RevenueTypesNotNil() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRevenueTypes)))
	})
}

// FeeRateEQ applies the EQ predicate on the "fee_rate" field.
func FeeRateEQ(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeeRate), v))
	})
}

// FeeRateNEQ applies the NEQ predicate on the "fee_rate" field.
func FeeRateNEQ(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFeeRate), v))
	})
}

// FeeRateIn applies the In predicate on the "fee_rate" field.
func FeeRateIn(vs ...decimal.Decimal) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFeeRate), v...))
	})
}

// FeeRateNotIn applies the NotIn predicate on the "fee_rate" field.
func FeeRateNotIn(vs ...decimal.Decimal) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFeeRate), v...))
	})
}

// FeeRateGT applies the GT predicate on the "fee_rate" field.
func FeeRateGT(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFeeRate), v))
	})
}

// FeeRateGTE applies the GTE predicate on the "fee_rate" field.
func FeeRateGTE(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFeeRate), v))
	})
}

// FeeRateLT applies the LT predicate on the "fee_rate" field.
func FeeRateLT(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFeeRate), v))
	})
}

// FeeRateLTE applies the LTE predicate on the "fee_rate" field.
func FeeRateLTE(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFeeRate), v))
	})
}

// FeeRateIsNil applies the IsNil predicate on the "fee_rate" field.
func FeeRateIsNil() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFeeRate)))
	})
}

// FeeRateNotNil applies the NotNil predicate on the "fee_rate" field.
func FeeRateNotNil() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFeeRate)))
	})
}

// FixedRevenueAbleEQ applies the EQ predicate on the "fixed_revenue_able" field.
func FixedRevenueAbleEQ(v bool) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFixedRevenueAble), v))
	})
}

// FixedRevenueAbleNEQ applies the NEQ predicate on the "fixed_revenue_able" field.
func FixedRevenueAbleNEQ(v bool) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFixedRevenueAble), v))
	})
}

// FixedRevenueAbleIsNil applies the IsNil predicate on the "fixed_revenue_able" field.
func FixedRevenueAbleIsNil() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFixedRevenueAble)))
	})
}

// FixedRevenueAbleNotNil applies the NotNil predicate on the "fixed_revenue_able" field.
func FixedRevenueAbleNotNil() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFixedRevenueAble)))
	})
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemark), v))
	})
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRemark), v))
	})
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRemark), v...))
	})
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRemark), v...))
	})
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRemark), v))
	})
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRemark), v))
	})
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRemark), v))
	})
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRemark), v))
	})
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRemark), v))
	})
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRemark), v))
	})
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRemark), v))
	})
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRemark)))
	})
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRemark)))
	})
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRemark), v))
	})
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRemark), v))
	})
}

// ThresholdEQ applies the EQ predicate on the "threshold" field.
func ThresholdEQ(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThreshold), v))
	})
}

// ThresholdNEQ applies the NEQ predicate on the "threshold" field.
func ThresholdNEQ(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThreshold), v))
	})
}

// ThresholdIn applies the In predicate on the "threshold" field.
func ThresholdIn(vs ...decimal.Decimal) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldThreshold), v...))
	})
}

// ThresholdNotIn applies the NotIn predicate on the "threshold" field.
func ThresholdNotIn(vs ...decimal.Decimal) predicate.Coin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldThreshold), v...))
	})
}

// ThresholdGT applies the GT predicate on the "threshold" field.
func ThresholdGT(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThreshold), v))
	})
}

// ThresholdGTE applies the GTE predicate on the "threshold" field.
func ThresholdGTE(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThreshold), v))
	})
}

// ThresholdLT applies the LT predicate on the "threshold" field.
func ThresholdLT(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThreshold), v))
	})
}

// ThresholdLTE applies the LTE predicate on the "threshold" field.
func ThresholdLTE(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThreshold), v))
	})
}

// ThresholdIsNil applies the IsNil predicate on the "threshold" field.
func ThresholdIsNil() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldThreshold)))
	})
}

// ThresholdNotNil applies the NotNil predicate on the "threshold" field.
func ThresholdNotNil() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldThreshold)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Coin) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Coin) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
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
func Not(p predicate.Coin) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		p(s.Not())
	})
}
