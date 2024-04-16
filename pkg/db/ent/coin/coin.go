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
	// FieldMiningpoolType holds the string denoting the miningpool_type field in the database.
	FieldMiningpoolType = "miningpool_type"
	// FieldCoinType holds the string denoting the coin_type field in the database.
	FieldCoinType = "coin_type"
	// FieldRevenueTypes holds the string denoting the revenue_types field in the database.
	FieldRevenueTypes = "revenue_types"
	// FieldFeeRate holds the string denoting the fee_rate field in the database.
	FieldFeeRate = "fee_rate"
	// FieldFixedRevenueAble holds the string denoting the fixed_revenue_able field in the database.
	FieldFixedRevenueAble = "fixed_revenue_able"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldThreshold holds the string denoting the threshold field in the database.
	FieldThreshold = "threshold"
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
	FieldMiningpoolType,
	FieldCoinType,
	FieldRevenueTypes,
	FieldFeeRate,
	FieldFixedRevenueAble,
	FieldRemark,
	FieldThreshold,
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
	// DefaultMiningpoolType holds the default value on creation for the "miningpool_type" field.
	DefaultMiningpoolType string
	// DefaultCoinType holds the default value on creation for the "coin_type" field.
	DefaultCoinType string
	// DefaultRevenueTypes holds the default value on creation for the "revenue_types" field.
	DefaultRevenueTypes []string
	// DefaultFeeRate holds the default value on creation for the "fee_rate" field.
	DefaultFeeRate decimal.Decimal
	// DefaultFixedRevenueAble holds the default value on creation for the "fixed_revenue_able" field.
	DefaultFixedRevenueAble bool
	// DefaultRemark holds the default value on creation for the "remark" field.
	DefaultRemark string
	// DefaultThreshold holds the default value on creation for the "threshold" field.
	DefaultThreshold decimal.Decimal
)
