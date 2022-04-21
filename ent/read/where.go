// Code generated by entc, DO NOT EDIT.

package read

import (
	"pingpong2/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Result applies equality check predicate on the "result" field. It's identical to ResultEQ.
func Result(v int) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResult), v))
	})
}

// ResultEQ applies the EQ predicate on the "result" field.
func ResultEQ(v int) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResult), v))
	})
}

// ResultNEQ applies the NEQ predicate on the "result" field.
func ResultNEQ(v int) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResult), v))
	})
}

// ResultIn applies the In predicate on the "result" field.
func ResultIn(vs ...int) predicate.Read {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Read(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldResult), v...))
	})
}

// ResultNotIn applies the NotIn predicate on the "result" field.
func ResultNotIn(vs ...int) predicate.Read {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Read(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldResult), v...))
	})
}

// ResultGT applies the GT predicate on the "result" field.
func ResultGT(v int) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldResult), v))
	})
}

// ResultGTE applies the GTE predicate on the "result" field.
func ResultGTE(v int) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldResult), v))
	})
}

// ResultLT applies the LT predicate on the "result" field.
func ResultLT(v int) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldResult), v))
	})
}

// ResultLTE applies the LTE predicate on the "result" field.
func ResultLTE(v int) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldResult), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSentence applies the HasEdge predicate on the "sentence" edge.
func HasSentence() predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SentenceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SentenceTable, SentenceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSentenceWith applies the HasEdge predicate on the "sentence" edge with a given conditions (other predicates).
func HasSentenceWith(preds ...predicate.Sentense) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SentenceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SentenceTable, SentenceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Read) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Read) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
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
func Not(p predicate.Read) predicate.Read {
	return predicate.Read(func(s *sql.Selector) {
		p(s.Not())
	})
}