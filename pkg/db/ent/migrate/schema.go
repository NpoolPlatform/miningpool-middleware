// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppPoolsColumns holds the columns for the "app_pools" table.
	AppPoolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "pool_id", Type: field.TypeUUID, Nullable: true},
	}
	// AppPoolsTable holds the schema information for the "app_pools" table.
	AppPoolsTable = &schema.Table{
		Name:       "app_pools",
		Columns:    AppPoolsColumns,
		PrimaryKey: []*schema.Column{AppPoolsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "apppool_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppPoolsColumns[4]},
			},
		},
	}
	// CoinsColumns holds the columns for the "coins" table.
	CoinsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "pool_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "fee_ratio", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "fixed_revenue_able", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "least_transfer_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "benefit_interval_seconds", Type: field.TypeUint32, Nullable: true, Default: 0},
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
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "order_user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "withdraw_state", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "withdraw_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "promise_pay_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "msg", Type: field.TypeString, Nullable: true, Default: ""},
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
			{
				Name:    "fraction_app_id_user_id",
				Unique:  false,
				Columns: []*schema.Column{FractionsColumns[5], FractionsColumns[6]},
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
		{Name: "pool_coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "withdraw_interval", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "min_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "withdraw_rate", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
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
		{Name: "root_user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "pool_coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "read_page_link", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
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
			{
				Name:    "gooduser_root_user_id",
				Unique:  false,
				Columns: []*schema.Column{GoodUsersColumns[5]},
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
		{Name: "good_user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "proportion", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "revenue_address", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "read_page_link", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "auto_pay", Type: field.TypeBool, Nullable: true, Default: false},
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
			{
				Name:    "orderuser_app_id",
				Unique:  false,
				Columns: []*schema.Column{OrderUsersColumns[7]},
			},
			{
				Name:    "orderuser_good_user_id",
				Unique:  false,
				Columns: []*schema.Column{OrderUsersColumns[5]},
			},
		},
	}
	// PoolsColumns holds the columns for the "pools" table.
	PoolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "miningpool_type", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "site", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "logo", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// PoolsTable holds the schema information for the "pools" table.
	PoolsTable = &schema.Table{
		Name:       "pools",
		Columns:    PoolsColumns,
		PrimaryKey: []*schema.Column{PoolsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "pool_ent_id",
				Unique:  true,
				Columns: []*schema.Column{PoolsColumns[4]},
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
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "pool_id", Type: field.TypeUUID, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "auth_token", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "auth_token_salt", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "authed", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "remark", Type: field.TypeString, Nullable: true, Default: ""},
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
		AppPoolsTable,
		CoinsTable,
		FractionsTable,
		FractionRulesTable,
		GoodUsersTable,
		OrderUsersTable,
		PoolsTable,
		RootUsersTable,
	}
)

func init() {
}
