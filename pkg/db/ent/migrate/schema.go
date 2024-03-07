// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CoinsColumns holds the columns for the "coins" table.
	CoinsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "miningpool_type", Type: field.TypeString},
		{Name: "site", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "coin_type", Type: field.TypeString},
		{Name: "revenue_type", Type: field.TypeString},
		{Name: "fee_rate", Type: field.TypeFloat32, Nullable: true, Default: 0},
		{Name: "fixed_revenue_able", Type: field.TypeBool},
		{Name: "remark", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// CoinsTable holds the schema information for the "coins" table.
	CoinsTable = &schema.Table{
		Name:       "coins",
		Columns:    CoinsColumns,
		PrimaryKey: []*schema.Column{CoinsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "coin_ent_id",
				Unique:  true,
				Columns: []*schema.Column{CoinsColumns[4]},
			},
		},
	}
	// FractionsColumns holds the columns for the "fractions" table.
	FractionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "order_mininguser_id", Type: field.TypeUUID},
		{Name: "withdraw_state", Type: field.TypeString},
		{Name: "withdraw_time", Type: field.TypeString},
		{Name: "pay_time", Type: field.TypeUint32, Nullable: true},
	}
	// FractionsTable holds the schema information for the "fractions" table.
	FractionsTable = &schema.Table{
		Name:       "fractions",
		Columns:    FractionsColumns,
		PrimaryKey: []*schema.Column{FractionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "fraction_ent_id",
				Unique:  true,
				Columns: []*schema.Column{FractionsColumns[4]},
			},
		},
	}
	// FractionRulesColumns holds the columns for the "fraction_rules" table.
	FractionRulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "miningpool_type", Type: field.TypeString},
		{Name: "coin_type", Type: field.TypeString},
		{Name: "withdraw_interval", Type: field.TypeUint32},
		{Name: "min_amount", Type: field.TypeFloat32},
		{Name: "withdraw_rate", Type: field.TypeFloat32},
	}
	// FractionRulesTable holds the schema information for the "fraction_rules" table.
	FractionRulesTable = &schema.Table{
		Name:       "fraction_rules",
		Columns:    FractionRulesColumns,
		PrimaryKey: []*schema.Column{FractionRulesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "fractionrule_ent_id",
				Unique:  true,
				Columns: []*schema.Column{FractionRulesColumns[4]},
			},
		},
	}
	// GoodUsersColumns holds the columns for the "good_users" table.
	GoodUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "auth_token", Type: field.TypeString},
		{Name: "authed", Type: field.TypeBool},
		{Name: "start", Type: field.TypeUint32},
		{Name: "end", Type: field.TypeUint32},
		{Name: "hash_rate", Type: field.TypeFloat64, Nullable: true, Default: 0},
		{Name: "read_page_link", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// GoodUsersTable holds the schema information for the "good_users" table.
	GoodUsersTable = &schema.Table{
		Name:       "good_users",
		Columns:    GoodUsersColumns,
		PrimaryKey: []*schema.Column{GoodUsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "gooduser_ent_id",
				Unique:  true,
				Columns: []*schema.Column{GoodUsersColumns[4]},
			},
		},
	}
	// OrderUsersColumns holds the columns for the "order_users" table.
	OrderUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "good_user_id", Type: field.TypeString},
		{Name: "coin_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "proportion", Type: field.TypeFloat32},
		{Name: "start", Type: field.TypeUint32},
		{Name: "end", Type: field.TypeUint32},
		{Name: "compensation_time", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "revenue_address", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "threshold", Type: field.TypeFloat32},
		{Name: "read_page_link", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "auto_pay", Type: field.TypeBool},
	}
	// OrderUsersTable holds the schema information for the "order_users" table.
	OrderUsersTable = &schema.Table{
		Name:       "order_users",
		Columns:    OrderUsersColumns,
		PrimaryKey: []*schema.Column{OrderUsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "orderuser_ent_id",
				Unique:  true,
				Columns: []*schema.Column{OrderUsersColumns[4]},
			},
		},
	}
	// RootUsersColumns holds the columns for the "root_users" table.
	RootUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "miningpool_type", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "auth_token", Type: field.TypeString},
		{Name: "authed", Type: field.TypeBool},
		{Name: "remark", Type: field.TypeString},
	}
	// RootUsersTable holds the schema information for the "root_users" table.
	RootUsersTable = &schema.Table{
		Name:       "root_users",
		Columns:    RootUsersColumns,
		PrimaryKey: []*schema.Column{RootUsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "rootuser_ent_id",
				Unique:  true,
				Columns: []*schema.Column{RootUsersColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CoinsTable,
		FractionsTable,
		FractionRulesTable,
		GoodUsersTable,
		OrderUsersTable,
		RootUsersTable,
	}
)

func init() {
}
