package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Jet struct {
	ent.Schema
}

func (Jet) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name").Optional(),
		field.String("color").Nillable().Optional(),
		field.String("uuid").Optional(),
		field.String("identifier").Optional(),
		field.Bytes("cargo").Optional(),
		field.Bytes("manifest").Optional(),
	}
}

func (Jet) Edges() []ent.Edge {

	return []ent.Edge{
		edge.To("pilot", Pilot.Type).Unique(),
		edge.To("airport", Airport.Type).Unique(),
	}
}

func (Jet) Config() ent.Config {
	return ent.Config{}
}
