// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeReads holds the string denoting the reads edge name in mutations.
	EdgeReads = "reads"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ReadsTable is the table that holds the reads relation/edge.
	ReadsTable = "reads"
	// ReadsInverseTable is the table name for the Read entity.
	// It exists in this package in order to avoid circular dependency with the "read" package.
	ReadsInverseTable = "reads"
	// ReadsColumn is the table column denoting the reads relation/edge.
	ReadsColumn = "user_reads"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPassword,
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