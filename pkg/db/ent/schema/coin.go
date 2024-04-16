package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/mixin"
	"github.com/shopspring/decimal"
)

// Coin holds the schema definition for the Coin entity.
type Coin struct {
	ent.Schema
}

func (Coin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Coin.
func (Coin) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("miningpool_type").Optional().Default(""),
		field.
			String("coin_type").Optional().Default(""),
		field.
			JSON("revenue_types", []string{}).Optional().Default([]string{}),
		field.
			Other("fee_rate", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),

		field.
			Bool("fixed_revenue_able").Optional().Default(false),
		field.
			String("remark").Optional().Default(""),
		field.
			Other("threshold", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the Coin.
func (Coin) Edges() []ent.Edge {
	return nil
}
