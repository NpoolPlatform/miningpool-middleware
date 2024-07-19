// Code generated by ent, DO NOT EDIT.

package coin

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the coin type in the database.
	Label = "coin"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldEntID holds the string denoting the ent_id field in the database.
	FieldEntID = "ent_id"
	// FieldPoolID holds the string denoting the pool_id field in the database.
	FieldPoolID = "pool_id"
	// FieldCoinTypeID holds the string denoting the coin_type_id field in the database.
	FieldCoinTypeID = "coin_type_id"
	// FieldCoinType holds the string denoting the coin_type field in the database.
	FieldCoinType = "coin_type"
	// FieldFeeRatio holds the string denoting the fee_ratio field in the database.
	FieldFeeRatio = "fee_ratio"
	// FieldFixedRevenueAble holds the string denoting the fixed_revenue_able field in the database.
	FieldFixedRevenueAble = "fixed_revenue_able"
	// FieldLeastTransferAmount holds the string denoting the least_transfer_amount field in the database.
	FieldLeastTransferAmount = "least_transfer_amount"
	// FieldBenefitIntervalSeconds holds the string denoting the benefit_interval_seconds field in the database.
	FieldBenefitIntervalSeconds = "benefit_interval_seconds"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// Table holds the table name of the coin in the database.
	Table = "coins"
)

// Columns holds all SQL columns for coin fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldPoolID,
	FieldCoinTypeID,
	FieldCoinType,
	FieldFeeRatio,
	FieldFixedRevenueAble,
	FieldLeastTransferAmount,
	FieldBenefitIntervalSeconds,
	FieldRemark,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/runtime"
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultEntID holds the default value on creation for the "ent_id" field.
	DefaultEntID func() uuid.UUID
	// DefaultPoolID holds the default value on creation for the "pool_id" field.
	DefaultPoolID func() uuid.UUID
	// DefaultCoinTypeID holds the default value on creation for the "coin_type_id" field.
	DefaultCoinTypeID func() uuid.UUID
	// DefaultCoinType holds the default value on creation for the "coin_type" field.
	DefaultCoinType string
	// DefaultFeeRatio holds the default value on creation for the "fee_ratio" field.
	DefaultFeeRatio decimal.Decimal
	// DefaultFixedRevenueAble holds the default value on creation for the "fixed_revenue_able" field.
	DefaultFixedRevenueAble bool
	// DefaultLeastTransferAmount holds the default value on creation for the "least_transfer_amount" field.
	DefaultLeastTransferAmount decimal.Decimal
	// DefaultBenefitIntervalSeconds holds the default value on creation for the "benefit_interval_seconds" field.
	DefaultBenefitIntervalSeconds uint32
	// DefaultRemark holds the default value on creation for the "remark" field.
	DefaultRemark string
)
