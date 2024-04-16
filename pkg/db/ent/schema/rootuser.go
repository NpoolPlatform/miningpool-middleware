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
			String("name").Optional().Default(""),
		field.
			String("miningpool_type").Optional().Default(""),
		field.
			String("email").Optional().Default(""),
		field.
			Text("auth_token").Optional().Default(""),
		field.
			Text("auth_token_salt").Optional().Default(""),
		field.
			Bool("authed").Optional().Default(false),
		field.
			String("remark").Optional().Default(""),
	}
}

// Edges of the RootUser.
func (RootUser) Edges() []ent.Edge {
	return nil
}
