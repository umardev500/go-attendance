package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Immutable().
			Positive(),

		field.String("name").
			MaxLen(50).
			NotEmpty(),

		field.String("location").
			MaxLen(100).
			NotEmpty(),

		field.Time("installed_at").
			Default(time.Now),

		field.Bool("is_active").
			Default(true),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attendances", Attendance.Type),
	}
}
