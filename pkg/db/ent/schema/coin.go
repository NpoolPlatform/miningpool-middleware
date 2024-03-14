package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/mixin"
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
			String("miningpool_type"),
		field.
			String("site").Optional().Default(""),
		field.
			String("coin_type"),
		field.
			JSON("revenue_types", []string{}).Optional().Default([]string{}),
		field.
			Float32("fee_rate").Optional().Default(0),
		field.
			Bool("fixed_revenue_able").Optional().Default(false),
		field.
			String("remark").Optional().Default(""),
		field.
			Float32("threshold").Optional().Default(0),
	}
}

// Edges of the Coin.
func (Coin) Edges() []ent.Edge {
	return nil
}
