package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Province holds the schema definition for Province entity.
type Province struct {
	ent.Schema
}

// Fields of the Province.
func (Province) Fields() []ent.Field {
	return []ent.Field{
		
		field.String("Province_Name").NotEmpty(),
	}
}

// Edges of the Province.
func (Province) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Province_ID", User.Type),
	}
}
