package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// FractionRule holds the schema definition for the FractionRule entity.
type FractionRule struct {
	ent.Schema
}

func (FractionRule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the FractionRule.
func (FractionRule) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("pool_coin_type_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Uint32("withdraw_interval").Optional().Default(0),
		field.
			Other("min_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("withdraw_rate", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the FractionRule.
func (FractionRule) Edges() []ent.Edge {
	return nil
}
