// Code generated by ent, DO NOT EDIT.

package area

const (
	// Label holds the string label denoting the area type in the database.
	Label = "area"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPoints holds the string denoting the points field in the database.
	FieldPoints = "points"
	// EdgeActivities holds the string denoting the activities edge name in mutations.
	EdgeActivities = "activities"
	// EdgeYear holds the string denoting the year edge name in mutations.
	EdgeYear = "year"
	// Table holds the table name of the area in the database.
	Table = "areas"
	// ActivitiesTable is the table that holds the activities relation/edge.
	ActivitiesTable = "activities"
	// ActivitiesInverseTable is the table name for the Activity entity.
	// It exists in this package in order to avoid circular dependency with the "activity" package.
	ActivitiesInverseTable = "activities"
	// ActivitiesColumn is the table column denoting the activities relation/edge.
	ActivitiesColumn = "area_activities"
	// YearTable is the table that holds the year relation/edge.
	YearTable = "areas"
	// YearInverseTable is the table name for the Year entity.
	// It exists in this package in order to avoid circular dependency with the "year" package.
	YearInverseTable = "years"
	// YearColumn is the table column denoting the year relation/edge.
	YearColumn = "year_areas"
)

// Columns holds all SQL columns for area fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPoints,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "areas"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"year_areas",
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
