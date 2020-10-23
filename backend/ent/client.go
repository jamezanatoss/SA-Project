// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/james/app/ent/migrate"

	"github.com/james/app/ent/gender"
	"github.com/james/app/ent/province"
	"github.com/james/app/ent/user"
	"github.com/james/app/ent/usertype"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Gender is the client for interacting with the Gender builders.
	Gender *GenderClient
	// Province is the client for interacting with the Province builders.
	Province *ProvinceClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserType is the client for interacting with the UserType builders.
	UserType *UserTypeClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Gender = NewGenderClient(c.config)
	c.Province = NewProvinceClient(c.config)
	c.User = NewUserClient(c.config)
	c.UserType = NewUserTypeClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Gender:   NewGenderClient(cfg),
		Province: NewProvinceClient(cfg),
		User:     NewUserClient(cfg),
		UserType: NewUserTypeClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:   cfg,
		Gender:   NewGenderClient(cfg),
		Province: NewProvinceClient(cfg),
		User:     NewUserClient(cfg),
		UserType: NewUserTypeClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Gender.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Gender.Use(hooks...)
	c.Province.Use(hooks...)
	c.User.Use(hooks...)
	c.UserType.Use(hooks...)
}

// GenderClient is a client for the Gender schema.
type GenderClient struct {
	config
}

// NewGenderClient returns a client for the Gender from the given config.
func NewGenderClient(c config) *GenderClient {
	return &GenderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `gender.Hooks(f(g(h())))`.
func (c *GenderClient) Use(hooks ...Hook) {
	c.hooks.Gender = append(c.hooks.Gender, hooks...)
}

// Create returns a create builder for Gender.
func (c *GenderClient) Create() *GenderCreate {
	mutation := newGenderMutation(c.config, OpCreate)
	return &GenderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Gender.
func (c *GenderClient) Update() *GenderUpdate {
	mutation := newGenderMutation(c.config, OpUpdate)
	return &GenderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GenderClient) UpdateOne(ge *Gender) *GenderUpdateOne {
	mutation := newGenderMutation(c.config, OpUpdateOne, withGender(ge))
	return &GenderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GenderClient) UpdateOneID(id int) *GenderUpdateOne {
	mutation := newGenderMutation(c.config, OpUpdateOne, withGenderID(id))
	return &GenderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Gender.
func (c *GenderClient) Delete() *GenderDelete {
	mutation := newGenderMutation(c.config, OpDelete)
	return &GenderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GenderClient) DeleteOne(ge *Gender) *GenderDeleteOne {
	return c.DeleteOneID(ge.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GenderClient) DeleteOneID(id int) *GenderDeleteOne {
	builder := c.Delete().Where(gender.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GenderDeleteOne{builder}
}

// Create returns a query builder for Gender.
func (c *GenderClient) Query() *GenderQuery {
	return &GenderQuery{config: c.config}
}

// Get returns a Gender entity by its id.
func (c *GenderClient) Get(ctx context.Context, id int) (*Gender, error) {
	return c.Query().Where(gender.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GenderClient) GetX(ctx context.Context, id int) *Gender {
	ge, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ge
}

// QueryGenderID queries the Gender_ID edge of a Gender.
func (c *GenderClient) QueryGenderID(ge *Gender) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ge.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(gender.Table, gender.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, gender.GenderIDTable, gender.GenderIDColumn),
		)
		fromV = sqlgraph.Neighbors(ge.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GenderClient) Hooks() []Hook {
	return c.hooks.Gender
}

// ProvinceClient is a client for the Province schema.
type ProvinceClient struct {
	config
}

// NewProvinceClient returns a client for the Province from the given config.
func NewProvinceClient(c config) *ProvinceClient {
	return &ProvinceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `province.Hooks(f(g(h())))`.
func (c *ProvinceClient) Use(hooks ...Hook) {
	c.hooks.Province = append(c.hooks.Province, hooks...)
}

// Create returns a create builder for Province.
func (c *ProvinceClient) Create() *ProvinceCreate {
	mutation := newProvinceMutation(c.config, OpCreate)
	return &ProvinceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Province.
func (c *ProvinceClient) Update() *ProvinceUpdate {
	mutation := newProvinceMutation(c.config, OpUpdate)
	return &ProvinceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProvinceClient) UpdateOne(pr *Province) *ProvinceUpdateOne {
	mutation := newProvinceMutation(c.config, OpUpdateOne, withProvince(pr))
	return &ProvinceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProvinceClient) UpdateOneID(id int) *ProvinceUpdateOne {
	mutation := newProvinceMutation(c.config, OpUpdateOne, withProvinceID(id))
	return &ProvinceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Province.
func (c *ProvinceClient) Delete() *ProvinceDelete {
	mutation := newProvinceMutation(c.config, OpDelete)
	return &ProvinceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProvinceClient) DeleteOne(pr *Province) *ProvinceDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProvinceClient) DeleteOneID(id int) *ProvinceDeleteOne {
	builder := c.Delete().Where(province.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProvinceDeleteOne{builder}
}

// Create returns a query builder for Province.
func (c *ProvinceClient) Query() *ProvinceQuery {
	return &ProvinceQuery{config: c.config}
}

// Get returns a Province entity by its id.
func (c *ProvinceClient) Get(ctx context.Context, id int) (*Province, error) {
	return c.Query().Where(province.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProvinceClient) GetX(ctx context.Context, id int) *Province {
	pr, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pr
}

// QueryProvinceID queries the Province_ID edge of a Province.
func (c *ProvinceClient) QueryProvinceID(pr *Province) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(province.Table, province.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, province.ProvinceIDTable, province.ProvinceIDColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProvinceClient) Hooks() []Hook {
	return c.hooks.Province
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryGenderID queries the Gender_ID edge of a User.
func (c *UserClient) QueryGenderID(u *User) *GenderQuery {
	query := &GenderQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(gender.Table, gender.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.GenderIDTable, user.GenderIDColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUserTypeID queries the UserType_ID edge of a User.
func (c *UserClient) QueryUserTypeID(u *User) *UserTypeQuery {
	query := &UserTypeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(usertype.Table, usertype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.UserTypeIDTable, user.UserTypeIDColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProvinceID queries the Province_ID edge of a User.
func (c *UserClient) QueryProvinceID(u *User) *ProvinceQuery {
	query := &ProvinceQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(province.Table, province.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.ProvinceIDTable, user.ProvinceIDColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// UserTypeClient is a client for the UserType schema.
type UserTypeClient struct {
	config
}

// NewUserTypeClient returns a client for the UserType from the given config.
func NewUserTypeClient(c config) *UserTypeClient {
	return &UserTypeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `usertype.Hooks(f(g(h())))`.
func (c *UserTypeClient) Use(hooks ...Hook) {
	c.hooks.UserType = append(c.hooks.UserType, hooks...)
}

// Create returns a create builder for UserType.
func (c *UserTypeClient) Create() *UserTypeCreate {
	mutation := newUserTypeMutation(c.config, OpCreate)
	return &UserTypeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for UserType.
func (c *UserTypeClient) Update() *UserTypeUpdate {
	mutation := newUserTypeMutation(c.config, OpUpdate)
	return &UserTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserTypeClient) UpdateOne(ut *UserType) *UserTypeUpdateOne {
	mutation := newUserTypeMutation(c.config, OpUpdateOne, withUserType(ut))
	return &UserTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserTypeClient) UpdateOneID(id int) *UserTypeUpdateOne {
	mutation := newUserTypeMutation(c.config, OpUpdateOne, withUserTypeID(id))
	return &UserTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserType.
func (c *UserTypeClient) Delete() *UserTypeDelete {
	mutation := newUserTypeMutation(c.config, OpDelete)
	return &UserTypeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserTypeClient) DeleteOne(ut *UserType) *UserTypeDeleteOne {
	return c.DeleteOneID(ut.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserTypeClient) DeleteOneID(id int) *UserTypeDeleteOne {
	builder := c.Delete().Where(usertype.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserTypeDeleteOne{builder}
}

// Create returns a query builder for UserType.
func (c *UserTypeClient) Query() *UserTypeQuery {
	return &UserTypeQuery{config: c.config}
}

// Get returns a UserType entity by its id.
func (c *UserTypeClient) Get(ctx context.Context, id int) (*UserType, error) {
	return c.Query().Where(usertype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserTypeClient) GetX(ctx context.Context, id int) *UserType {
	ut, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ut
}

// QueryUserTypeID queries the UserType_ID edge of a UserType.
func (c *UserTypeClient) QueryUserTypeID(ut *UserType) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ut.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(usertype.Table, usertype.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, usertype.UserTypeIDTable, usertype.UserTypeIDColumn),
		)
		fromV = sqlgraph.Neighbors(ut.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserTypeClient) Hooks() []Hook {
	return c.hooks.UserType
}