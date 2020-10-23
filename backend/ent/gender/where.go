// Code generated by entc, DO NOT EDIT.

package gender

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/james/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// GenderName applies equality check predicate on the "Gender_Name" field. It's identical to GenderNameEQ.
func GenderName(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGenderName), v))
	})
}

// GenderNameEQ applies the EQ predicate on the "Gender_Name" field.
func GenderNameEQ(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGenderName), v))
	})
}

// GenderNameNEQ applies the NEQ predicate on the "Gender_Name" field.
func GenderNameNEQ(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGenderName), v))
	})
}

// GenderNameIn applies the In predicate on the "Gender_Name" field.
func GenderNameIn(vs ...string) predicate.Gender {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gender(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGenderName), v...))
	})
}

// GenderNameNotIn applies the NotIn predicate on the "Gender_Name" field.
func GenderNameNotIn(vs ...string) predicate.Gender {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gender(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGenderName), v...))
	})
}

// GenderNameGT applies the GT predicate on the "Gender_Name" field.
func GenderNameGT(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGenderName), v))
	})
}

// GenderNameGTE applies the GTE predicate on the "Gender_Name" field.
func GenderNameGTE(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGenderName), v))
	})
}

// GenderNameLT applies the LT predicate on the "Gender_Name" field.
func GenderNameLT(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGenderName), v))
	})
}

// GenderNameLTE applies the LTE predicate on the "Gender_Name" field.
func GenderNameLTE(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGenderName), v))
	})
}

// GenderNameContains applies the Contains predicate on the "Gender_Name" field.
func GenderNameContains(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGenderName), v))
	})
}

// GenderNameHasPrefix applies the HasPrefix predicate on the "Gender_Name" field.
func GenderNameHasPrefix(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGenderName), v))
	})
}

// GenderNameHasSuffix applies the HasSuffix predicate on the "Gender_Name" field.
func GenderNameHasSuffix(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGenderName), v))
	})
}

// GenderNameEqualFold applies the EqualFold predicate on the "Gender_Name" field.
func GenderNameEqualFold(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGenderName), v))
	})
}

// GenderNameContainsFold applies the ContainsFold predicate on the "Gender_Name" field.
func GenderNameContainsFold(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGenderName), v))
	})
}

// HasGenderID applies the HasEdge predicate on the "Gender_ID" edge.
func HasGenderID() predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenderIDTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GenderIDTable, GenderIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGenderIDWith applies the HasEdge predicate on the "Gender_ID" edge with a given conditions (other predicates).
func HasGenderIDWith(preds ...predicate.User) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenderIDInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GenderIDTable, GenderIDColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Gender) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Gender) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Gender) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		p(s.Not())
	})
}