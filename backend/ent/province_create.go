// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/james/app/ent/province"
	"github.com/james/app/ent/user"
)

// ProvinceCreate is the builder for creating a Province entity.
type ProvinceCreate struct {
	config
	mutation *ProvinceMutation
	hooks    []Hook
}

// SetProvinceName sets the Province_Name field.
func (pc *ProvinceCreate) SetProvinceName(s string) *ProvinceCreate {
	pc.mutation.SetProvinceName(s)
	return pc
}

// AddProvinceIDIDs adds the Province_ID edge to User by ids.
func (pc *ProvinceCreate) AddProvinceIDIDs(ids ...int) *ProvinceCreate {
	pc.mutation.AddProvinceIDIDs(ids...)
	return pc
}

// AddProvinceID adds the Province_ID edges to User.
func (pc *ProvinceCreate) AddProvinceID(u ...*User) *ProvinceCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddProvinceIDIDs(ids...)
}

// Mutation returns the ProvinceMutation object of the builder.
func (pc *ProvinceCreate) Mutation() *ProvinceMutation {
	return pc.mutation
}

// Save creates the Province in the database.
func (pc *ProvinceCreate) Save(ctx context.Context) (*Province, error) {
	if _, ok := pc.mutation.ProvinceName(); !ok {
		return nil, &ValidationError{Name: "Province_Name", err: errors.New("ent: missing required field \"Province_Name\"")}
	}
	if v, ok := pc.mutation.ProvinceName(); ok {
		if err := province.ProvinceNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Province_Name", err: fmt.Errorf("ent: validator failed for field \"Province_Name\": %w", err)}
		}
	}
	var (
		err  error
		node *Province
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProvinceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProvinceCreate) SaveX(ctx context.Context) *Province {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *ProvinceCreate) sqlSave(ctx context.Context) (*Province, error) {
	pr, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pr.ID = int(id)
	return pr, nil
}

func (pc *ProvinceCreate) createSpec() (*Province, *sqlgraph.CreateSpec) {
	var (
		pr    = &Province{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: province.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: province.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.ProvinceName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: province.FieldProvinceName,
		})
		pr.ProvinceName = value
	}
	if nodes := pc.mutation.ProvinceIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   province.ProvinceIDTable,
			Columns: []string{province.ProvinceIDColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pr, _spec
}
