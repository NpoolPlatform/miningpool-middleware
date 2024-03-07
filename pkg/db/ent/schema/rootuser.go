package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/mixin"
)

// RootUser holds the schema definition for the RootUser entity.
type RootUser struct {
	ent.Schema
}

func (RootUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the RootUser.
func (RootUser) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name"),
		field.
			String("miningpool_type"),
		field.
			String("email"),
		field.
			String("auth_token"),
		field.
			Bool("authed"),
		field.
			String("remark"),
	}
}

// Edges of the RootUser.
func (RootUser) Edges() []ent.Edge {
	return nil
}
