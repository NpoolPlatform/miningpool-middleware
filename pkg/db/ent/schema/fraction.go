package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
			UUID("app_id", uuid.UUID{}).Optional().
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("user_id", uuid.UUID{}).Optional().
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("order_user_id", uuid.UUID{}).Optional().
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("coin_type_id", uuid.UUID{}).Optional().
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("withdraw_state").Optional().Default(""),
		field.
			Uint32("withdraw_at").Optional().Default(0),
		field.
			Uint32("promise_pay_at").Optional().Default(0),
		field.
			String("msg").Optional().Default(""),
	}
}

// Edges of the Fraction.
func (Fraction) Edges() []ent.Edge {
	return nil
}

func (Fraction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "user_id"),
	}
}
