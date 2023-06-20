package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Mixin of the Notification.
func (Notification) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Node{},
	}
}

func (Notification) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return nil
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return nil
}
