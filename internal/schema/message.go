package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Mixin of the Message.
func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Node{},
	}
}

func (Message) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("body").
			NotEmpty(),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("to", Message.Type).
			Annotations(
				entgql.RelayConnection(),
			).
			From("from").
			Annotations(
				entgql.RelayConnection(),
			),
	}
}
