package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/mixin"
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
			String("miningpool_type").Optional().Default(""),
		field.
			String("coin_type").Optional().Default(""),
		field.
			Uint32("withdraw_interval").Optional().Default(0),
		field.
			Float32("min_amount").Optional().Default(0),
		field.
			Float32("withdraw_rate").Optional().Default(0),
	}
}

// Edges of the FractionRule.
func (FractionRule) Edges() []ent.Edge {
	return nil
}
