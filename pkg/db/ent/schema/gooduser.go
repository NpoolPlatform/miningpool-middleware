package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// GoodUser holds the schema definition for the GoodUser entity.
type GoodUser struct {
	ent.Schema
}

func (GoodUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the GoodUser.
func (GoodUser) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("root_user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("name").Optional().Default(""),
		field.
			UUID("coin_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Float32("hash_rate").Optional().Default(0),
		field.
			Text("read_page_link").Optional().Default(""),
	}
}

// Edges of the GoodUser.
func (GoodUser) Edges() []ent.Edge {
	return nil
}

func (GoodUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("root_user_id"),
	}
}
