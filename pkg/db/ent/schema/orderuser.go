package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// OrderUser holds the schema definition for the OrderUser entity.
type OrderUser struct {
	ent.Schema
}

func (OrderUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the OrderUser.
func (OrderUser) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("good_user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("name").Optional().Default(""),
		field.
			Other("proportion", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("revenue_address").Optional().Default(""),
		field.
			Text("read_page_link").Optional().Default(""),
		field.
			Bool("auto_pay").Optional().Default(false),
	}
}

// Edges of the OrderUser.
func (OrderUser) Edges() []ent.Edge {
	return nil
}

func (OrderUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id"),
		index.Fields("good_user_id"),
	}
}
