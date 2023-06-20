package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
)

// Setting holds the schema definition for the Setting entity.
type Setting struct {
	ent.Schema
}

// Mixin of the Setting.
func (Setting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Node{},
	}
}

func (Setting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}

// Fields of the Setting.
func (Setting) Fields() []ent.Field {
	return nil
}

// Edges of the Setting.
func (Setting) Edges() []ent.Edge {
	return nil
}
