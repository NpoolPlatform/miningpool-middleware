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
			String("miningpool_type"),
		field.
			String("coin_type"),
		field.
			Uint32("withdraw_interval"),
		field.
			Float32("min_amount"),
		field.
			Float32("withdraw_rate"),
	}
}

// Edges of the FractionRule.
func (FractionRule) Edges() []ent.Edge {
	return nil
}
