// Code generated by entc, DO NOT EDIT.

package sentense

import (
	"pingpong2/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Chinese applies equality check predicate on the "chinese" field. It's identical to ChineseEQ.
func Chinese(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChinese), v))
	})
}

// Pinyin applies equality check predicate on the "pinyin" field. It's identical to PinyinEQ.
func Pinyin(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPinyin), v))
	})
}

// English applies equality check predicate on the "english" field. It's identical to EnglishEQ.
func English(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnglish), v))
	})
}

// ChineseEQ applies the EQ predicate on the "chinese" field.
func ChineseEQ(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChinese), v))
	})
}

// ChineseNEQ applies the NEQ predicate on the "chinese" field.
func ChineseNEQ(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChinese), v))
	})
}

// ChineseIn applies the In predicate on the "chinese" field.
func ChineseIn(vs ...string) predicate.Sentense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sentense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldChinese), v...))
	})
}

// ChineseNotIn applies the NotIn predicate on the "chinese" field.
func ChineseNotIn(vs ...string) predicate.Sentense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sentense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldChinese), v...))
	})
}

// ChineseGT applies the GT predicate on the "chinese" field.
func ChineseGT(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChinese), v))
	})
}

// ChineseGTE applies the GTE predicate on the "chinese" field.
func ChineseGTE(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChinese), v))
	})
}

// ChineseLT applies the LT predicate on the "chinese" field.
func ChineseLT(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChinese), v))
	})
}

// ChineseLTE applies the LTE predicate on the "chinese" field.
func ChineseLTE(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChinese), v))
	})
}

// ChineseContains applies the Contains predicate on the "chinese" field.
func ChineseContains(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldChinese), v))
	})
}

// ChineseHasPrefix applies the HasPrefix predicate on the "chinese" field.
func ChineseHasPrefix(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldChinese), v))
	})
}

// ChineseHasSuffix applies the HasSuffix predicate on the "chinese" field.
func ChineseHasSuffix(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldChinese), v))
	})
}

// ChineseEqualFold applies the EqualFold predicate on the "chinese" field.
func ChineseEqualFold(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldChinese), v))
	})
}

// ChineseContainsFold applies the ContainsFold predicate on the "chinese" field.
func ChineseContainsFold(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldChinese), v))
	})
}

// PinyinEQ applies the EQ predicate on the "pinyin" field.
func PinyinEQ(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPinyin), v))
	})
}

// PinyinNEQ applies the NEQ predicate on the "pinyin" field.
func PinyinNEQ(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPinyin), v))
	})
}

// PinyinIn applies the In predicate on the "pinyin" field.
func PinyinIn(vs ...string) predicate.Sentense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sentense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPinyin), v...))
	})
}

// PinyinNotIn applies the NotIn predicate on the "pinyin" field.
func PinyinNotIn(vs ...string) predicate.Sentense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sentense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPinyin), v...))
	})
}

// PinyinGT applies the GT predicate on the "pinyin" field.
func PinyinGT(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPinyin), v))
	})
}

// PinyinGTE applies the GTE predicate on the "pinyin" field.
func PinyinGTE(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPinyin), v))
	})
}

// PinyinLT applies the LT predicate on the "pinyin" field.
func PinyinLT(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPinyin), v))
	})
}

// PinyinLTE applies the LTE predicate on the "pinyin" field.
func PinyinLTE(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPinyin), v))
	})
}

// PinyinContains applies the Contains predicate on the "pinyin" field.
func PinyinContains(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPinyin), v))
	})
}

// PinyinHasPrefix applies the HasPrefix predicate on the "pinyin" field.
func PinyinHasPrefix(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPinyin), v))
	})
}

// PinyinHasSuffix applies the HasSuffix predicate on the "pinyin" field.
func PinyinHasSuffix(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPinyin), v))
	})
}

// PinyinEqualFold applies the EqualFold predicate on the "pinyin" field.
func PinyinEqualFold(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPinyin), v))
	})
}

// PinyinContainsFold applies the ContainsFold predicate on the "pinyin" field.
func PinyinContainsFold(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPinyin), v))
	})
}

// EnglishEQ applies the EQ predicate on the "english" field.
func EnglishEQ(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnglish), v))
	})
}

// EnglishNEQ applies the NEQ predicate on the "english" field.
func EnglishNEQ(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnglish), v))
	})
}

// EnglishIn applies the In predicate on the "english" field.
func EnglishIn(vs ...string) predicate.Sentense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sentense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEnglish), v...))
	})
}

// EnglishNotIn applies the NotIn predicate on the "english" field.
func EnglishNotIn(vs ...string) predicate.Sentense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sentense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEnglish), v...))
	})
}

// EnglishGT applies the GT predicate on the "english" field.
func EnglishGT(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEnglish), v))
	})
}

// EnglishGTE applies the GTE predicate on the "english" field.
func EnglishGTE(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEnglish), v))
	})
}

// EnglishLT applies the LT predicate on the "english" field.
func EnglishLT(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEnglish), v))
	})
}

// EnglishLTE applies the LTE predicate on the "english" field.
func EnglishLTE(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEnglish), v))
	})
}

// EnglishContains applies the Contains predicate on the "english" field.
func EnglishContains(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEnglish), v))
	})
}

// EnglishHasPrefix applies the HasPrefix predicate on the "english" field.
func EnglishHasPrefix(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEnglish), v))
	})
}

// EnglishHasSuffix applies the HasSuffix predicate on the "english" field.
func EnglishHasSuffix(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEnglish), v))
	})
}

// EnglishEqualFold applies the EqualFold predicate on the "english" field.
func EnglishEqualFold(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEnglish), v))
	})
}

// EnglishContainsFold applies the ContainsFold predicate on the "english" field.
func EnglishContainsFold(v string) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEnglish), v))
	})
}

// HasReads applies the HasEdge predicate on the "reads" edge.
func HasReads() predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReadsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReadsTable, ReadsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReadsWith applies the HasEdge predicate on the "reads" edge with a given conditions (other predicates).
func HasReadsWith(preds ...predicate.Read) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReadsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReadsTable, ReadsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Sentense) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Sentense) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
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
func Not(p predicate.Sentense) predicate.Sentense {
	return predicate.Sentense(func(s *sql.Selector) {
		p(s.Not())
	})
}
