package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Card struct {
	ent.Schema
}

func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.String("card_uid").NotEmpty().Unique(),
		field.Time("issued_at").Immutable().Default(time.Now),
		field.Bool("is_active").Default(true),
	}
}

func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("card").Unique(),
		edge.To("scan_logs", ScanLog.Type),
	}
}
