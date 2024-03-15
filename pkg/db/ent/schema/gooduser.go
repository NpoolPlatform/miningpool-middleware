package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
			UUID("good_id", uuid.UUID{}),
		field.
			UUID("root_user_id", uuid.UUID{}),
		field.
			String("name"),
		field.
			String("miningpool_type"),
		field.
			String("coin_type"),
		field.
			Float32("hash_rate"),
		field.
			String("read_page_link"),
		field.
			String("revenue_type"),
	}
}

// Edges of the GoodUser.
func (GoodUser) Edges() []ent.Edge {
	return nil
}
