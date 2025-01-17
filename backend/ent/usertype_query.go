// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/james/app/ent/predicate"
	"github.com/james/app/ent/user"
	"github.com/james/app/ent/usertype"
)

// UserTypeQuery is the builder for querying UserType entities.
type UserTypeQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.UserType
	// eager-loading edges.
	withUserTypeID *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (utq *UserTypeQuery) Where(ps ...predicate.UserType) *UserTypeQuery {
	utq.predicates = append(utq.predicates, ps...)
	return utq
}

// Limit adds a limit step to the query.
func (utq *UserTypeQuery) Limit(limit int) *UserTypeQuery {
	utq.limit = &limit
	return utq
}

// Offset adds an offset step to the query.
func (utq *UserTypeQuery) Offset(offset int) *UserTypeQuery {
	utq.offset = &offset
	return utq
}

// Order adds an order step to the query.
func (utq *UserTypeQuery) Order(o ...OrderFunc) *UserTypeQuery {
	utq.order = append(utq.order, o...)
	return utq
}

// QueryUserTypeID chains the current query on the UserType_ID edge.
func (utq *UserTypeQuery) QueryUserTypeID() *UserQuery {
	query := &UserQuery{config: utq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := utq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usertype.Table, usertype.FieldID, utq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, usertype.UserTypeIDTable, usertype.UserTypeIDColumn),
		)
		fromU = sqlgraph.SetNeighbors(utq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserType entity in the query. Returns *NotFoundError when no usertype was found.
func (utq *UserTypeQuery) First(ctx context.Context) (*UserType, error) {
	uts, err := utq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(uts) == 0 {
		return nil, &NotFoundError{usertype.Label}
	}
	return uts[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (utq *UserTypeQuery) FirstX(ctx context.Context) *UserType {
	ut, err := utq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return ut
}

// FirstID returns the first UserType id in the query. Returns *NotFoundError when no id was found.
func (utq *UserTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = utq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usertype.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (utq *UserTypeQuery) FirstXID(ctx context.Context) int {
	id, err := utq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only UserType entity in the query, returns an error if not exactly one entity was returned.
func (utq *UserTypeQuery) Only(ctx context.Context) (*UserType, error) {
	uts, err := utq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(uts) {
	case 1:
		return uts[0], nil
	case 0:
		return nil, &NotFoundError{usertype.Label}
	default:
		return nil, &NotSingularError{usertype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (utq *UserTypeQuery) OnlyX(ctx context.Context) *UserType {
	ut, err := utq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return ut
}

// OnlyID returns the only UserType id in the query, returns an error if not exactly one id was returned.
func (utq *UserTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = utq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usertype.Label}
	default:
		err = &NotSingularError{usertype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (utq *UserTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := utq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserTypes.
func (utq *UserTypeQuery) All(ctx context.Context) ([]*UserType, error) {
	if err := utq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return utq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (utq *UserTypeQuery) AllX(ctx context.Context) []*UserType {
	uts, err := utq.All(ctx)
	if err != nil {
		panic(err)
	}
	return uts
}

// IDs executes the query and returns a list of UserType ids.
func (utq *UserTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := utq.Select(usertype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (utq *UserTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := utq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (utq *UserTypeQuery) Count(ctx context.Context) (int, error) {
	if err := utq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return utq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (utq *UserTypeQuery) CountX(ctx context.Context) int {
	count, err := utq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (utq *UserTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := utq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return utq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (utq *UserTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := utq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (utq *UserTypeQuery) Clone() *UserTypeQuery {
	return &UserTypeQuery{
		config:     utq.config,
		limit:      utq.limit,
		offset:     utq.offset,
		order:      append([]OrderFunc{}, utq.order...),
		unique:     append([]string{}, utq.unique...),
		predicates: append([]predicate.UserType{}, utq.predicates...),
		// clone intermediate query.
		sql:  utq.sql.Clone(),
		path: utq.path,
	}
}

//  WithUserTypeID tells the query-builder to eager-loads the nodes that are connected to
// the "UserType_ID" edge. The optional arguments used to configure the query builder of the edge.
func (utq *UserTypeQuery) WithUserTypeID(opts ...func(*UserQuery)) *UserTypeQuery {
	query := &UserQuery{config: utq.config}
	for _, opt := range opts {
		opt(query)
	}
	utq.withUserTypeID = query
	return utq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserTypeName string `json:"UserType_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserType.Query().
//		GroupBy(usertype.FieldUserTypeName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (utq *UserTypeQuery) GroupBy(field string, fields ...string) *UserTypeGroupBy {
	group := &UserTypeGroupBy{config: utq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := utq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return utq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		UserTypeName string `json:"UserType_name,omitempty"`
//	}
//
//	client.UserType.Query().
//		Select(usertype.FieldUserTypeName).
//		Scan(ctx, &v)
//
func (utq *UserTypeQuery) Select(field string, fields ...string) *UserTypeSelect {
	selector := &UserTypeSelect{config: utq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := utq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return utq.sqlQuery(), nil
	}
	return selector
}

func (utq *UserTypeQuery) prepareQuery(ctx context.Context) error {
	if utq.path != nil {
		prev, err := utq.path(ctx)
		if err != nil {
			return err
		}
		utq.sql = prev
	}
	return nil
}

func (utq *UserTypeQuery) sqlAll(ctx context.Context) ([]*UserType, error) {
	var (
		nodes       = []*UserType{}
		_spec       = utq.querySpec()
		loadedTypes = [1]bool{
			utq.withUserTypeID != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &UserType{config: utq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, utq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := utq.withUserTypeID; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*UserType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.User(func(s *sql.Selector) {
			s.Where(sql.InValues(usertype.UserTypeIDColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.user_type_user_type_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "user_type_user_type_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_type_user_type_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.UserTypeID = append(node.Edges.UserTypeID, n)
		}
	}

	return nodes, nil
}

func (utq *UserTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := utq.querySpec()
	return sqlgraph.CountNodes(ctx, utq.driver, _spec)
}

func (utq *UserTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := utq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (utq *UserTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usertype.Table,
			Columns: usertype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usertype.FieldID,
			},
		},
		From:   utq.sql,
		Unique: true,
	}
	if ps := utq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := utq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := utq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := utq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (utq *UserTypeQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(utq.driver.Dialect())
	t1 := builder.Table(usertype.Table)
	selector := builder.Select(t1.Columns(usertype.Columns...)...).From(t1)
	if utq.sql != nil {
		selector = utq.sql
		selector.Select(selector.Columns(usertype.Columns...)...)
	}
	for _, p := range utq.predicates {
		p(selector)
	}
	for _, p := range utq.order {
		p(selector)
	}
	if offset := utq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := utq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserTypeGroupBy is the builder for group-by UserType entities.
type UserTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (utgb *UserTypeGroupBy) Aggregate(fns ...AggregateFunc) *UserTypeGroupBy {
	utgb.fns = append(utgb.fns, fns...)
	return utgb
}

// Scan applies the group-by query and scan the result into the given value.
func (utgb *UserTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := utgb.path(ctx)
	if err != nil {
		return err
	}
	utgb.sql = query
	return utgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (utgb *UserTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := utgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (utgb *UserTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(utgb.fields) > 1 {
		return nil, errors.New("ent: UserTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := utgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (utgb *UserTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := utgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (utgb *UserTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = utgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usertype.Label}
	default:
		err = fmt.Errorf("ent: UserTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (utgb *UserTypeGroupBy) StringX(ctx context.Context) string {
	v, err := utgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (utgb *UserTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(utgb.fields) > 1 {
		return nil, errors.New("ent: UserTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := utgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (utgb *UserTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := utgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (utgb *UserTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = utgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usertype.Label}
	default:
		err = fmt.Errorf("ent: UserTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (utgb *UserTypeGroupBy) IntX(ctx context.Context) int {
	v, err := utgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (utgb *UserTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(utgb.fields) > 1 {
		return nil, errors.New("ent: UserTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := utgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (utgb *UserTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := utgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (utgb *UserTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = utgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usertype.Label}
	default:
		err = fmt.Errorf("ent: UserTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (utgb *UserTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := utgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (utgb *UserTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(utgb.fields) > 1 {
		return nil, errors.New("ent: UserTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := utgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (utgb *UserTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := utgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (utgb *UserTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = utgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usertype.Label}
	default:
		err = fmt.Errorf("ent: UserTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (utgb *UserTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := utgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (utgb *UserTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := utgb.sqlQuery().Query()
	if err := utgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (utgb *UserTypeGroupBy) sqlQuery() *sql.Selector {
	selector := utgb.sql
	columns := make([]string, 0, len(utgb.fields)+len(utgb.fns))
	columns = append(columns, utgb.fields...)
	for _, fn := range utgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(utgb.fields...)
}

// UserTypeSelect is the builder for select fields of UserType entities.
type UserTypeSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (uts *UserTypeSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := uts.path(ctx)
	if err != nil {
		return err
	}
	uts.sql = query
	return uts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uts *UserTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := uts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (uts *UserTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(uts.fields) > 1 {
		return nil, errors.New("ent: UserTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := uts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uts *UserTypeSelect) StringsX(ctx context.Context) []string {
	v, err := uts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (uts *UserTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usertype.Label}
	default:
		err = fmt.Errorf("ent: UserTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uts *UserTypeSelect) StringX(ctx context.Context) string {
	v, err := uts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (uts *UserTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(uts.fields) > 1 {
		return nil, errors.New("ent: UserTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := uts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uts *UserTypeSelect) IntsX(ctx context.Context) []int {
	v, err := uts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (uts *UserTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usertype.Label}
	default:
		err = fmt.Errorf("ent: UserTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uts *UserTypeSelect) IntX(ctx context.Context) int {
	v, err := uts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (uts *UserTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(uts.fields) > 1 {
		return nil, errors.New("ent: UserTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := uts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uts *UserTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := uts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (uts *UserTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usertype.Label}
	default:
		err = fmt.Errorf("ent: UserTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uts *UserTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := uts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (uts *UserTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(uts.fields) > 1 {
		return nil, errors.New("ent: UserTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := uts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uts *UserTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := uts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (uts *UserTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usertype.Label}
	default:
		err = fmt.Errorf("ent: UserTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uts *UserTypeSelect) BoolX(ctx context.Context) bool {
	v, err := uts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uts *UserTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uts.sqlQuery().Query()
	if err := uts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uts *UserTypeSelect) sqlQuery() sql.Querier {
	selector := uts.sql
	selector.Select(selector.Columns(uts.fields...)...)
	return selector
}
