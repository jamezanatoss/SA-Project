package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Gender holds the schema definition for Gender entity.
type Gender struct {
	ent.Schema
}

// Fields of the Gender.
func (Gender) Fields() []ent.Field {
	return []ent.Field{
		
		field.String("Gender_Name").NotEmpty(),
	}
}

// Edges of the Gender.
func (Gender) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Gender_ID", User.Type),
	}
}
