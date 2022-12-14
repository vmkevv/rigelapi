// Code generated by ent, DO NOT EDIT.

package dpto

const (
	// Label holds the string label denoting the dpto type in the database.
	Label = "dpto"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeProvincias holds the string denoting the provincias edge name in mutations.
	EdgeProvincias = "provincias"
	// Table holds the table name of the dpto in the database.
	Table = "dptos"
	// ProvinciasTable is the table that holds the provincias relation/edge.
	ProvinciasTable = "provincia"
	// ProvinciasInverseTable is the table name for the Provincia entity.
	// It exists in this package in order to avoid circular dependency with the "provincia" package.
	ProvinciasInverseTable = "provincia"
	// ProvinciasColumn is the table column denoting the provincias relation/edge.
	ProvinciasColumn = "dpto_provincias"
)

// Columns holds all SQL columns for dpto fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
