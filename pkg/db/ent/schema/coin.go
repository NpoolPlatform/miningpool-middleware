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
			String("revenue_type"),
		field.
			Float32("fee_rate").Optional().Default(0),
		field.
			Bool("fixed_revenue_able"),
		field.
			String("remark").Optional().Default(""),
	}
}

// Edges of the Coin.
func (Coin) Edges() []ent.Edge {
	return nil
}
