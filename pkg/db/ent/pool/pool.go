// Code generated by ent, DO NOT EDIT.

package pool

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the pool type in the database.
	Label = "pool"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSite holds the string denoting the site field in the database.
	FieldSite = "site"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// Table holds the table name of the pool in the database.
	Table = "pools"
)

// Columns holds all SQL columns for pool fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldMiningpoolType,
	FieldName,
	FieldSite,
	FieldDescription,
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
	// DefaultSite holds the default value on creation for the "site" field.
	DefaultSite string
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
)
