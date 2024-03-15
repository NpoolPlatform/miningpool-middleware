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
			UUID("root_user_id", uuid.UUID{}),
		field.
			UUID("good_user_id", uuid.UUID{}),
		field.
			UUID("order_id", uuid.UUID{}),
		field.
			String("name"),
		field.
			String("miningpool_type"),
		field.
			String("Coin_Type"),
		field.
			Float32("proportion"),
		field.
			String("revenue_address"),
		field.
			String("read_page_link"),
		field.
			Bool("auto_pay").Optional().Default(false),
	}
}

// Edges of the OrderUser.
func (OrderUser) Edges() []ent.Edge {
	return nil
}
