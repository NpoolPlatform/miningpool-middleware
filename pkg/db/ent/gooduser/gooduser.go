// Code generated by ent, DO NOT EDIT.

package gooduser

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the gooduser type in the database.
	Label = "good_user"
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
	// FieldRootUserID holds the string denoting the root_user_id field in the database.
	FieldRootUserID = "root_user_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCoinID holds the string denoting the coin_id field in the database.
	FieldCoinID = "coin_id"
	// FieldHashRate holds the string denoting the hash_rate field in the database.
	FieldHashRate = "hash_rate"
	// FieldReadPageLink holds the string denoting the read_page_link field in the database.
	FieldReadPageLink = "read_page_link"
	// Table holds the table name of the gooduser in the database.
	Table = "good_users"
)

// Columns holds all SQL columns for gooduser fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldRootUserID,
	FieldName,
	FieldCoinID,
	FieldHashRate,
	FieldReadPageLink,
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
	// DefaultRootUserID holds the default value on creation for the "root_user_id" field.
	DefaultRootUserID func() uuid.UUID
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultCoinID holds the default value on creation for the "coin_id" field.
	DefaultCoinID func() uuid.UUID
	// DefaultHashRate holds the default value on creation for the "hash_rate" field.
	DefaultHashRate float32
	// DefaultReadPageLink holds the default value on creation for the "read_page_link" field.
	DefaultReadPageLink string
)
