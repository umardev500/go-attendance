package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),

		field.String("name").
			MaxLen(50).
			NotEmpty().
			Nillable().Unique(),

		field.String("location").
			MaxLen(100).
			NotEmpty().
			Nillable(),

		field.Time("installed_at").
			Default(time.Now).
			Nillable(),

		field.Bool("is_active").
			Default(true).
			Nillable(),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attendances", Attendance.Type),
	}
}
