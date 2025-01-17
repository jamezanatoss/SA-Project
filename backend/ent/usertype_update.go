// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/james/app/ent/predicate"
	"github.com/james/app/ent/user"
	"github.com/james/app/ent/usertype"
)

// UserTypeUpdate is the builder for updating UserType entities.
type UserTypeUpdate struct {
	config
	hooks      []Hook
	mutation   *UserTypeMutation
	predicates []predicate.UserType
}

// Where adds a new predicate for the builder.
func (utu *UserTypeUpdate) Where(ps ...predicate.UserType) *UserTypeUpdate {
	utu.predicates = append(utu.predicates, ps...)
	return utu
}

// SetUserTypeName sets the UserType_name field.
func (utu *UserTypeUpdate) SetUserTypeName(s string) *UserTypeUpdate {
	utu.mutation.SetUserTypeName(s)
	return utu
}

// AddUserTypeIDIDs adds the UserType_ID edge to User by ids.
func (utu *UserTypeUpdate) AddUserTypeIDIDs(ids ...int) *UserTypeUpdate {
	utu.mutation.AddUserTypeIDIDs(ids...)
	return utu
}

// AddUserTypeID adds the UserType_ID edges to User.
func (utu *UserTypeUpdate) AddUserTypeID(u ...*User) *UserTypeUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return utu.AddUserTypeIDIDs(ids...)
}

// Mutation returns the UserTypeMutation object of the builder.
func (utu *UserTypeUpdate) Mutation() *UserTypeMutation {
	return utu.mutation
}

// RemoveUserTypeIDIDs removes the UserType_ID edge to User by ids.
func (utu *UserTypeUpdate) RemoveUserTypeIDIDs(ids ...int) *UserTypeUpdate {
	utu.mutation.RemoveUserTypeIDIDs(ids...)
	return utu
}

// RemoveUserTypeID removes UserType_ID edges to User.
func (utu *UserTypeUpdate) RemoveUserTypeID(u ...*User) *UserTypeUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return utu.RemoveUserTypeIDIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (utu *UserTypeUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := utu.mutation.UserTypeName(); ok {
		if err := usertype.UserTypeNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "UserType_name", err: fmt.Errorf("ent: validator failed for field \"UserType_name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(utu.hooks) == 0 {
		affected, err = utu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			utu.mutation = mutation
			affected, err = utu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(utu.hooks) - 1; i >= 0; i-- {
			mut = utu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, utu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (utu *UserTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := utu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (utu *UserTypeUpdate) Exec(ctx context.Context) error {
	_, err := utu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utu *UserTypeUpdate) ExecX(ctx context.Context) {
	if err := utu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (utu *UserTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usertype.Table,
			Columns: usertype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usertype.FieldID,
			},
		},
	}
	if ps := utu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := utu.mutation.UserTypeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usertype.FieldUserTypeName,
		})
	}
	if nodes := utu.mutation.RemovedUserTypeIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usertype.UserTypeIDTable,
			Columns: []string{usertype.UserTypeIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := utu.mutation.UserTypeIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usertype.UserTypeIDTable,
			Columns: []string{usertype.UserTypeIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, utu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usertype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserTypeUpdateOne is the builder for updating a single UserType entity.
type UserTypeUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserTypeMutation
}

// SetUserTypeName sets the UserType_name field.
func (utuo *UserTypeUpdateOne) SetUserTypeName(s string) *UserTypeUpdateOne {
	utuo.mutation.SetUserTypeName(s)
	return utuo
}

// AddUserTypeIDIDs adds the UserType_ID edge to User by ids.
func (utuo *UserTypeUpdateOne) AddUserTypeIDIDs(ids ...int) *UserTypeUpdateOne {
	utuo.mutation.AddUserTypeIDIDs(ids...)
	return utuo
}

// AddUserTypeID adds the UserType_ID edges to User.
func (utuo *UserTypeUpdateOne) AddUserTypeID(u ...*User) *UserTypeUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return utuo.AddUserTypeIDIDs(ids...)
}

// Mutation returns the UserTypeMutation object of the builder.
func (utuo *UserTypeUpdateOne) Mutation() *UserTypeMutation {
	return utuo.mutation
}

// RemoveUserTypeIDIDs removes the UserType_ID edge to User by ids.
func (utuo *UserTypeUpdateOne) RemoveUserTypeIDIDs(ids ...int) *UserTypeUpdateOne {
	utuo.mutation.RemoveUserTypeIDIDs(ids...)
	return utuo
}

// RemoveUserTypeID removes UserType_ID edges to User.
func (utuo *UserTypeUpdateOne) RemoveUserTypeID(u ...*User) *UserTypeUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return utuo.RemoveUserTypeIDIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (utuo *UserTypeUpdateOne) Save(ctx context.Context) (*UserType, error) {
	if v, ok := utuo.mutation.UserTypeName(); ok {
		if err := usertype.UserTypeNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "UserType_name", err: fmt.Errorf("ent: validator failed for field \"UserType_name\": %w", err)}
		}
	}

	var (
		err  error
		node *UserType
	)
	if len(utuo.hooks) == 0 {
		node, err = utuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			utuo.mutation = mutation
			node, err = utuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(utuo.hooks) - 1; i >= 0; i-- {
			mut = utuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, utuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (utuo *UserTypeUpdateOne) SaveX(ctx context.Context) *UserType {
	ut, err := utuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return ut
}

// Exec executes the query on the entity.
func (utuo *UserTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := utuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utuo *UserTypeUpdateOne) ExecX(ctx context.Context) {
	if err := utuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (utuo *UserTypeUpdateOne) sqlSave(ctx context.Context) (ut *UserType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usertype.Table,
			Columns: usertype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usertype.FieldID,
			},
		},
	}
	id, ok := utuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UserType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := utuo.mutation.UserTypeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usertype.FieldUserTypeName,
		})
	}
	if nodes := utuo.mutation.RemovedUserTypeIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usertype.UserTypeIDTable,
			Columns: []string{usertype.UserTypeIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := utuo.mutation.UserTypeIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usertype.UserTypeIDTable,
			Columns: []string{usertype.UserTypeIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	ut = &UserType{config: utuo.config}
	_spec.Assign = ut.assignValues
	_spec.ScanValues = ut.scanValues()
	if err = sqlgraph.UpdateNode(ctx, utuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usertype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return ut, nil
}
