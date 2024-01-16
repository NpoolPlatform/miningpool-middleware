package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/mixin"
	"github.com/google/uuid"
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
			UUID("order_id", uuid.UUID{}),
		field.
			String("good_user_id"),
		field.
			String("coin_id"),
		field.
			String("name"),
		field.
			Float32("proportion"),
		field.
			Uint32("start"),
		field.
			Uint32("end"),
		field.
			Uint32("compensation_time").Optional().Default(0),
		field.
			String("revenue_address").Optional().Default(""),
		field.
			Float32("threshold"),
		field.
			String("read_page_link").Optional().Default(""),
		field.
			Bool("auto_pay"),
	}
}

// Edges of the OrderUser.
func (OrderUser) Edges() []ent.Edge {
	return nil
}
