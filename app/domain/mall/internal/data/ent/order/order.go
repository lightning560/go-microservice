// Code generated by ent, DO NOT EDIT.

package order

import (
	"time"
)

const (
	// Label holds the string label denoting the order type in the database.
	Label = "order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUID holds the string denoting the uid field in the database.
	FieldUID = "uid"
	// FieldOrderItems holds the string denoting the order_items field in the database.
	FieldOrderItems = "order_items"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldTotalAmount holds the string denoting the total_amount field in the database.
	FieldTotalAmount = "total_amount"
	// FieldTotalQuantity holds the string denoting the total_quantity field in the database.
	FieldTotalQuantity = "total_quantity"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the order in the database.
	Table = "orders"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
	FieldUID,
	FieldOrderItems,
	FieldState,
	FieldTotalAmount,
	FieldTotalQuantity,
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
