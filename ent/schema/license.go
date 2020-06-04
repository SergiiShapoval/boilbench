package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

type License struct {
	ent.Schema
}

func (License) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pilot", Pilot.Type).
			Unique().Required(),
	}
}
