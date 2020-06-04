package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

type Language struct {
	ent.Schema
}

func (Language) Fields() []ent.Field {
	return []ent.Field{
		field.String("language"),
	}
}
