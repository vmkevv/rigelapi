// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/vmkevv/rigelapi/ent/classperiodsync"
	"github.com/vmkevv/rigelapi/ent/teacher"
)

// ClassPeriodSync is the model entity for the ClassPeriodSync schema.
type ClassPeriodSync struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// LastSyncID holds the value of the "last_sync_id" field.
	LastSyncID string `json:"last_sync_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClassPeriodSyncQuery when eager-loading is set.
	Edges                      ClassPeriodSyncEdges `json:"edges"`
	teacher_class_period_syncs *string
}

// ClassPeriodSyncEdges holds the relations/edges for other nodes in the graph.
type ClassPeriodSyncEdges struct {
	// Teacher holds the value of the teacher edge.
	Teacher *Teacher `json:"teacher,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TeacherOrErr returns the Teacher value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClassPeriodSyncEdges) TeacherOrErr() (*Teacher, error) {
	if e.loadedTypes[0] {
		if e.Teacher == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: teacher.Label}
		}
		return e.Teacher, nil
	}
	return nil, &NotLoadedError{edge: "teacher"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ClassPeriodSync) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case classperiodsync.FieldID, classperiodsync.FieldLastSyncID:
			values[i] = new(sql.NullString)
		case classperiodsync.ForeignKeys[0]: // teacher_class_period_syncs
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ClassPeriodSync", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ClassPeriodSync fields.
func (cps *ClassPeriodSync) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case classperiodsync.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				cps.ID = value.String
			}
		case classperiodsync.FieldLastSyncID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_sync_id", values[i])
			} else if value.Valid {
				cps.LastSyncID = value.String
			}
		case classperiodsync.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field teacher_class_period_syncs", values[i])
			} else if value.Valid {
				cps.teacher_class_period_syncs = new(string)
				*cps.teacher_class_period_syncs = value.String
			}
		}
	}
	return nil
}

// QueryTeacher queries the "teacher" edge of the ClassPeriodSync entity.
func (cps *ClassPeriodSync) QueryTeacher() *TeacherQuery {
	return (&ClassPeriodSyncClient{config: cps.config}).QueryTeacher(cps)
}

// Update returns a builder for updating this ClassPeriodSync.
// Note that you need to call ClassPeriodSync.Unwrap() before calling this method if this ClassPeriodSync
// was returned from a transaction, and the transaction was committed or rolled back.
func (cps *ClassPeriodSync) Update() *ClassPeriodSyncUpdateOne {
	return (&ClassPeriodSyncClient{config: cps.config}).UpdateOne(cps)
}

// Unwrap unwraps the ClassPeriodSync entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cps *ClassPeriodSync) Unwrap() *ClassPeriodSync {
	_tx, ok := cps.config.driver.(*txDriver)
	if !ok {
		panic("ent: ClassPeriodSync is not a transactional entity")
	}
	cps.config.driver = _tx.drv
	return cps
}

// String implements the fmt.Stringer.
func (cps *ClassPeriodSync) String() string {
	var builder strings.Builder
	builder.WriteString("ClassPeriodSync(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cps.ID))
	builder.WriteString("last_sync_id=")
	builder.WriteString(cps.LastSyncID)
	builder.WriteByte(')')
	return builder.String()
}

// ClassPeriodSyncs is a parsable slice of ClassPeriodSync.
type ClassPeriodSyncs []*ClassPeriodSync

func (cps ClassPeriodSyncs) config(cfg config) {
	for _i := range cps {
		cps[_i].config = cfg
	}
}
