// Code generated by ent, DO NOT EDIT.

package rootuser

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the rootuser type in the database.
	Label = "root_user"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPoolID holds the string denoting the pool_id field in the database.
	FieldPoolID = "pool_id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldAuthToken holds the string denoting the auth_token field in the database.
	FieldAuthToken = "auth_token"
	// FieldAuthTokenSalt holds the string denoting the auth_token_salt field in the database.
	FieldAuthTokenSalt = "auth_token_salt"
	// FieldAuthed holds the string denoting the authed field in the database.
	FieldAuthed = "authed"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// Table holds the table name of the rootuser in the database.
	Table = "root_users"
)

// Columns holds all SQL columns for rootuser fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldName,
	FieldPoolID,
	FieldEmail,
	FieldAuthToken,
	FieldAuthTokenSalt,
	FieldAuthed,
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
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultPoolID holds the default value on creation for the "pool_id" field.
	DefaultPoolID func() uuid.UUID
	// DefaultEmail holds the default value on creation for the "email" field.
	DefaultEmail string
	// DefaultAuthToken holds the default value on creation for the "auth_token" field.
	DefaultAuthToken string
	// DefaultAuthTokenSalt holds the default value on creation for the "auth_token_salt" field.
	DefaultAuthTokenSalt string
	// DefaultAuthed holds the default value on creation for the "authed" field.
	DefaultAuthed bool
	// DefaultRemark holds the default value on creation for the "remark" field.
	DefaultRemark string
)
