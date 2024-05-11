// Code generated by ent, DO NOT EDIT.

package orderuser

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the orderuser type in the database.
	Label = "order_user"
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
	// FieldGoodUserID holds the string denoting the good_user_id field in the database.
	FieldGoodUserID = "good_user_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldProportion holds the string denoting the proportion field in the database.
	FieldProportion = "proportion"
	// FieldRevenueAddress holds the string denoting the revenue_address field in the database.
	FieldRevenueAddress = "revenue_address"
	// FieldReadPageLink holds the string denoting the read_page_link field in the database.
	FieldReadPageLink = "read_page_link"
	// FieldAutoPay holds the string denoting the auto_pay field in the database.
	FieldAutoPay = "auto_pay"
	// Table holds the table name of the orderuser in the database.
	Table = "order_users"
)

// Columns holds all SQL columns for orderuser fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldGoodUserID,
	FieldUserID,
	FieldAppID,
	FieldName,
	FieldProportion,
	FieldRevenueAddress,
	FieldReadPageLink,
	FieldAutoPay,
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
//
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
	// DefaultGoodUserID holds the default value on creation for the "good_user_id" field.
	DefaultGoodUserID func() uuid.UUID
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID func() uuid.UUID
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultProportion holds the default value on creation for the "proportion" field.
	DefaultProportion decimal.Decimal
	// DefaultRevenueAddress holds the default value on creation for the "revenue_address" field.
	DefaultRevenueAddress string
	// DefaultReadPageLink holds the default value on creation for the "read_page_link" field.
	DefaultReadPageLink string
	// DefaultAutoPay holds the default value on creation for the "auto_pay" field.
	DefaultAutoPay bool
)
