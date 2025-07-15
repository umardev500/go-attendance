package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Attendance struct {
	ent.Schema
}

func (Attendance) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.Time("check_in").Default(time.Now),
		field.Time("check_out").Optional().Nillable(),
		field.Time("date").Default(time.Now),
	}
}

func (Attendance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("attendances").Unique(),
		edge.From("devices", Device.Type).Ref("attendances").Unique(),
	}
}
