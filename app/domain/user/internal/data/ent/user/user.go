// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUID holds the string denoting the uid field in the database.
	FieldUID = "uid"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldIdentification holds the string denoting the identification field in the database.
	FieldIdentification = "identification"
	// FieldDeleted holds the string denoting the deleted field in the database.
	FieldDeleted = "deleted"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeProfiles holds the string denoting the profiles edge name in mutations.
	EdgeProfiles = "profiles"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ProfilesTable is the table that holds the profiles relation/edge.
	ProfilesTable = "profiles"
	// ProfilesInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	ProfilesInverseTable = "profiles"
	// ProfilesColumn is the table column denoting the profiles relation/edge.
	ProfilesColumn = "user_profiles"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUID,
	FieldPhone,
	FieldEmail,
	FieldPassword,
	FieldIdentification,
	FieldDeleted,
	FieldCreatedAt,
	FieldUpdatedAt,
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

var (
	// UIDValidator is a validator for the "uid" field. It is called by the builders before save.
	UIDValidator func(int64) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int64) error
)
