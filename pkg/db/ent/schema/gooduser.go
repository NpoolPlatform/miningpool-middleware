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
			String("name"),
		field.
			String("auth_token"),
		field.
			Bool("authed"),
		field.
			Uint32("start"),
		field.
			Uint32("end"),
		field.
			Float("hash_rate").Optional().Default(0),
		field.
			String("read_page_link").Optional().Default(""),
	}
}

// Edges of the GoodUser.
func (GoodUser) Edges() []ent.Edge {
	return nil
}
