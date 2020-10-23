// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/james/app/ent/gender"
	"github.com/james/app/ent/province"
	"github.com/james/app/ent/user"
	"github.com/james/app/ent/usertype"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// IdentityCard holds the value of the "Identity_card" field.
	IdentityCard int `json:"Identity_card,omitempty"`
	// Password holds the value of the "Password" field.
	Password string `json:"Password,omitempty"`
	// ConfirmPassword holds the value of the "Confirm_password" field.
	ConfirmPassword string `json:"Confirm_password,omitempty"`
	// FirstName holds the value of the "First_name" field.
	FirstName string `json:"First_name,omitempty"`
	// LastName holds the value of the "Last_name" field.
	LastName string `json:"Last_name,omitempty"`
	// Email holds the value of the "Email" field.
	Email string `json:"Email,omitempty"`
	// Phone holds the value of the "Phone" field.
	Phone int `json:"Phone,omitempty"`
	// DateOfBirth holds the value of the "Date_of_birth" field.
	DateOfBirth int `json:"Date_of_birth,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges                  UserEdges `json:"edges"`
	gender_gender_id       *int
	province_province_id   *int
	user_type_user_type_id *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// GenderID holds the value of the Gender_ID edge.
	GenderID *Gender
	// UserTypeID holds the value of the UserType_ID edge.
	UserTypeID *UserType
	// ProvinceID holds the value of the Province_ID edge.
	ProvinceID *Province
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// GenderIDOrErr returns the GenderID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) GenderIDOrErr() (*Gender, error) {
	if e.loadedTypes[0] {
		if e.GenderID == nil {
			// The edge Gender_ID was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: gender.Label}
		}
		return e.GenderID, nil
	}
	return nil, &NotLoadedError{edge: "Gender_ID"}
}

// UserTypeIDOrErr returns the UserTypeID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) UserTypeIDOrErr() (*UserType, error) {
	if e.loadedTypes[1] {
		if e.UserTypeID == nil {
			// The edge UserType_ID was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: usertype.Label}
		}
		return e.UserTypeID, nil
	}
	return nil, &NotLoadedError{edge: "UserType_ID"}
}

// ProvinceIDOrErr returns the ProvinceID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ProvinceIDOrErr() (*Province, error) {
	if e.loadedTypes[2] {
		if e.ProvinceID == nil {
			// The edge Province_ID was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: province.Label}
		}
		return e.ProvinceID, nil
	}
	return nil, &NotLoadedError{edge: "Province_ID"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // Identity_card
		&sql.NullString{}, // Password
		&sql.NullString{}, // Confirm_password
		&sql.NullString{}, // First_name
		&sql.NullString{}, // Last_name
		&sql.NullString{}, // Email
		&sql.NullInt64{},  // Phone
		&sql.NullInt64{},  // Date_of_birth
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*User) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // gender_gender_id
		&sql.NullInt64{}, // province_province_id
		&sql.NullInt64{}, // user_type_user_type_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field Identity_card", values[0])
	} else if value.Valid {
		u.IdentityCard = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Password", values[1])
	} else if value.Valid {
		u.Password = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Confirm_password", values[2])
	} else if value.Valid {
		u.ConfirmPassword = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field First_name", values[3])
	} else if value.Valid {
		u.FirstName = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Last_name", values[4])
	} else if value.Valid {
		u.LastName = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Email", values[5])
	} else if value.Valid {
		u.Email = value.String
	}
	if value, ok := values[6].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field Phone", values[6])
	} else if value.Valid {
		u.Phone = int(value.Int64)
	}
	if value, ok := values[7].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field Date_of_birth", values[7])
	} else if value.Valid {
		u.DateOfBirth = int(value.Int64)
	}
	values = values[8:]
	if len(values) == len(user.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field gender_gender_id", value)
		} else if value.Valid {
			u.gender_gender_id = new(int)
			*u.gender_gender_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field province_province_id", value)
		} else if value.Valid {
			u.province_province_id = new(int)
			*u.province_province_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_type_user_type_id", value)
		} else if value.Valid {
			u.user_type_user_type_id = new(int)
			*u.user_type_user_type_id = int(value.Int64)
		}
	}
	return nil
}

// QueryGenderID queries the Gender_ID edge of the User.
func (u *User) QueryGenderID() *GenderQuery {
	return (&UserClient{config: u.config}).QueryGenderID(u)
}

// QueryUserTypeID queries the UserType_ID edge of the User.
func (u *User) QueryUserTypeID() *UserTypeQuery {
	return (&UserClient{config: u.config}).QueryUserTypeID(u)
}

// QueryProvinceID queries the Province_ID edge of the User.
func (u *User) QueryProvinceID() *ProvinceQuery {
	return (&UserClient{config: u.config}).QueryProvinceID(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", Identity_card=")
	builder.WriteString(fmt.Sprintf("%v", u.IdentityCard))
	builder.WriteString(", Password=")
	builder.WriteString(u.Password)
	builder.WriteString(", Confirm_password=")
	builder.WriteString(u.ConfirmPassword)
	builder.WriteString(", First_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", Last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", Email=")
	builder.WriteString(u.Email)
	builder.WriteString(", Phone=")
	builder.WriteString(fmt.Sprintf("%v", u.Phone))
	builder.WriteString(", Date_of_birth=")
	builder.WriteString(fmt.Sprintf("%v", u.DateOfBirth))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
