package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// UserType holds the schema definition for UserType entity.
type UserType struct {
	ent.Schema
}

// Fields of the UserType.
func (UserType) Fields() []ent.Field {
	return []ent.Field{
		
		field.String("UserType_name").NotEmpty(),
	}
}

// Edges of the UserType.
func (UserType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("UserType_ID", User.Type),
	
	}
}
