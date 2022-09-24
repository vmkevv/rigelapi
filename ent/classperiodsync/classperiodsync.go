// Code generated by ent, DO NOT EDIT.

package classperiodsync

const (
	// Label holds the string label denoting the classperiodsync type in the database.
	Label = "class_period_sync"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLastSyncID holds the string denoting the last_sync_id field in the database.
	FieldLastSyncID = "last_sync_id"
	// EdgeTeacher holds the string denoting the teacher edge name in mutations.
	EdgeTeacher = "teacher"
	// Table holds the table name of the classperiodsync in the database.
	Table = "class_period_syncs"
	// TeacherTable is the table that holds the teacher relation/edge.
	TeacherTable = "class_period_syncs"
	// TeacherInverseTable is the table name for the Teacher entity.
	// It exists in this package in order to avoid circular dependency with the "teacher" package.
	TeacherInverseTable = "teachers"
	// TeacherColumn is the table column denoting the teacher relation/edge.
	TeacherColumn = "teacher_class_period_syncs"
)

// Columns holds all SQL columns for classperiodsync fields.
var Columns = []string{
	FieldID,
	FieldLastSyncID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "class_period_syncs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"teacher_class_period_syncs",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
