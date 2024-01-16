package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// Fraction holds the schema definition for the Fraction entity.
type Fraction struct {
	ent.Schema
}

func (Fraction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Fraction.
func (Fraction) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("order_mininguser_id", uuid.UUID{}),
		field.
			String("withdraw_state"),
		field.
			String("withdraw_time"),
		field.
			Uint32("pay_time").Optional(),
	}
}

// Edges of the Fraction.
func (Fraction) Edges() []ent.Edge {
	return nil
}
