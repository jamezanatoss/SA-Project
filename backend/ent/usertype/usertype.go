// Code generated by entc, DO NOT EDIT.

package usertype

const (
	// Label holds the string label denoting the usertype type in the database.
	Label = "user_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserTypeName holds the string denoting the usertype_name field in the database.
	FieldUserTypeName = "user_type_name"

	// EdgeUserTypeID holds the string denoting the usertype_id edge name in mutations.
	EdgeUserTypeID = "UserType_ID"

	// Table holds the table name of the usertype in the database.
	Table = "user_types"
	// UserTypeIDTable is the table the holds the UserType_ID relation/edge.
	UserTypeIDTable = "users"
	// UserTypeIDInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserTypeIDInverseTable = "users"
	// UserTypeIDColumn is the table column denoting the UserType_ID relation/edge.
	UserTypeIDColumn = "user_type_user_type_id"
)

// Columns holds all SQL columns for usertype fields.
var Columns = []string{
	FieldID,
	FieldUserTypeName,
}

var (
	// UserTypeNameValidator is a validator for the "UserType_name" field. It is called by the builders before save.
	UserTypeNameValidator func(string) error
)