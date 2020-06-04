package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

type Hangar struct {
	ent.Schema
}

func (Hangar) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}
