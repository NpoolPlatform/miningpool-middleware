// Code generated by ent, DO NOT EDIT.

package fraction

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the fraction type in the database.
	Label = "fraction"
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
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldOrderUserID holds the string denoting the order_user_id field in the database.
	FieldOrderUserID = "order_user_id"
	// FieldWithdrawState holds the string denoting the withdraw_state field in the database.
	FieldWithdrawState = "withdraw_state"
	// FieldWithdrawTime holds the string denoting the withdraw_time field in the database.
	FieldWithdrawTime = "withdraw_time"
	// FieldPayTime holds the string denoting the pay_time field in the database.
	FieldPayTime = "pay_time"
	// FieldMsg holds the string denoting the msg field in the database.
	FieldMsg = "msg"
	// Table holds the table name of the fraction in the database.
	Table = "fractions"
)

// Columns holds all SQL columns for fraction fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldAppID,
	FieldUserID,
	FieldOrderUserID,
	FieldWithdrawState,
	FieldWithdrawTime,
	FieldPayTime,
	FieldMsg,
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
	// DefaultPayTime holds the default value on creation for the "pay_time" field.
	DefaultPayTime uint32
	// DefaultMsg holds the default value on creation for the "msg" field.
	DefaultMsg string
)
