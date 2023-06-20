package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Node{},
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name"),

		field.String("last_name"),

		field.String("email_address").
			Unique(),

		field.String("account_address").
			Unique(),

		field.String("alias").
			Unique(),

		field.String("bio").
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("following", User.Type).
			Annotations(
				entgql.RelayConnection(),
			).
			From("followers").
			Annotations(
				entgql.RelayConnection(),
			),

		edge.To("posts", Post.Type).
			Annotations(
				entgql.RelayConnection(),
			),

		edge.From("groups", Group.Type).
			Ref("users").
			Annotations(
				entgql.RelayConnection(),
			),
	}
}
