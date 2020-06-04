package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

type Airport struct {
	ent.Schema
}

func (Airport) Fields() []ent.Field {
	return []ent.Field{
		field.Int("size").Nillable().Optional(),
	}
}
