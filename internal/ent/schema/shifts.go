package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Shift struct {
	ent.Schema
}

func (Shift) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.String("name").NotEmpty().Unique(),
		field.Time("start_time").Default(time.Now),
		field.Time("end_time").Default(time.Now),
	}
}

func (Shift) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("shifts"),
	}
}
