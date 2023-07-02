package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Task struct {
	ent.Schema
}

//+ ent:config:./ent/user

// Fields of the User.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("age").
			Positive(),
		field.String("address"),
	}
}

// Edges of the User.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		// Task много к одному с User
		edge.From("owner", User.Type).
			Ref("tasks").
			Unique(),
	}
}
