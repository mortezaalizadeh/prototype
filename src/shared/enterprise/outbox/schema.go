package outbox

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Outbox holds the schema definition for the Outbox entity.
type Outbox struct {
	ent.Schema
}

// Fields of the Outbox.
func (Outbox) Fields() []ent.Field {
	return []ent.Field{
		field.Time("timestamp"),
		field.String("topic"),
		field.String("key"),
		field.Bytes("payload"),
		field.JSON("headers", map[string]string{}),
		field.Int("retry_count"),
		field.Enum("status").
			NamedValues(
				"Pending", "PENDING",
				"Succeeded", "SUCCEEDED",
				"Failed", "FAILED"),
		field.Time("last_retry").Optional(),
		field.Strings("processing_errors").Optional(),
	}
}

// Edges of the Outbox.
func (Outbox) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Outbox) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("last_retry"),
		index.Fields("status"),
	}
}
