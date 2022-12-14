// Code generated by ent, DO NOT EDIT.

package attendanceday

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/vmkevv/rigelapi/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Day applies equality check predicate on the "day" field. It's identical to DayEQ.
func Day(v time.Time) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDay), v))
	})
}

// DayEQ applies the EQ predicate on the "day" field.
func DayEQ(v time.Time) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDay), v))
	})
}

// DayNEQ applies the NEQ predicate on the "day" field.
func DayNEQ(v time.Time) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDay), v))
	})
}

// DayIn applies the In predicate on the "day" field.
func DayIn(vs ...time.Time) predicate.AttendanceDay {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttendanceDay(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDay), v...))
	})
}

// DayNotIn applies the NotIn predicate on the "day" field.
func DayNotIn(vs ...time.Time) predicate.AttendanceDay {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttendanceDay(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDay), v...))
	})
}

// DayGT applies the GT predicate on the "day" field.
func DayGT(v time.Time) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDay), v))
	})
}

// DayGTE applies the GTE predicate on the "day" field.
func DayGTE(v time.Time) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDay), v))
	})
}

// DayLT applies the LT predicate on the "day" field.
func DayLT(v time.Time) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDay), v))
	})
}

// DayLTE applies the LTE predicate on the "day" field.
func DayLTE(v time.Time) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDay), v))
	})
}

// HasAttendances applies the HasEdge predicate on the "attendances" edge.
func HasAttendances() predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttendancesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttendancesTable, AttendancesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttendancesWith applies the HasEdge predicate on the "attendances" edge with a given conditions (other predicates).
func HasAttendancesWith(preds ...predicate.Attendance) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttendancesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttendancesTable, AttendancesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClassPeriod applies the HasEdge predicate on the "classPeriod" edge.
func HasClassPeriod() predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassPeriodTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClassPeriodTable, ClassPeriodColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassPeriodWith applies the HasEdge predicate on the "classPeriod" edge with a given conditions (other predicates).
func HasClassPeriodWith(preds ...predicate.ClassPeriod) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassPeriodInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClassPeriodTable, ClassPeriodColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AttendanceDay) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AttendanceDay) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
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
func Not(p predicate.AttendanceDay) predicate.AttendanceDay {
	return predicate.AttendanceDay(func(s *sql.Selector) {
		p(s.Not())
	})
}
