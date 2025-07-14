package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ScanLog struct {
	ent.Schema
}

func (ScanLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Immutable().
			Positive(),
		field.Time("scanned_at").Default(time.Now),
		field.Enum("status").
			Values("success", "rejected", "unknown_card").Default("success"),
		field.String("message").Nillable(),
	}
}

func (ScanLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("card", Card.Type).Ref("scan_logs").Unique().Required(),
	}
}
