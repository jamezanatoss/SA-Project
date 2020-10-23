package schema
import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)
// User holds the schema definition for User entity.
type User struct {
	ent.Schema
}
// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		
		field.Int("Identity_card").Positive(),
		field.String("Password").NotEmpty(),
		field.String("Confirm_password").NotEmpty(),
		field.String("First_name").NotEmpty(),
		field.String("Last_name").NotEmpty(),
		field.String("Email").NotEmpty(),
		field.Int("Phone").Positive(),
		field.Int("Date_of_birth").Positive(),
	}
}
// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Gender_ID", Gender.Type).Ref("Gender_ID").Unique(),
		edge.From("UserType_ID", UserType.Type).Ref("UserType_ID").Unique(),
		edge.From("Province_ID", Province.Type).Ref("Province_ID").Unique(),
	}
}
