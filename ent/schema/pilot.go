package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

type Pilot struct {
	ent.Schema
	Name string
}

func (Pilot) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

func (Pilot) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("languages", Language.Type),
	}
}

func (Pilot) Index() []ent.Index {
	return []ent.Index{
		index.Fields("languages").StorageKey("idx_pilot_languages"),
	}
}
