// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"redbook/app/domain/mall/internal/data/ent/order"
	"redbook/app/domain/mall/internal/data/ent/predicate"
	"redbook/app/domain/mall/internal/data/ent/schema"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetUID sets the "uid" field.
func (ou *OrderUpdate) SetUID(i int64) *OrderUpdate {
	ou.mutation.ResetUID()
	ou.mutation.SetUID(i)
	return ou
}

// AddUID adds i to the "uid" field.
func (ou *OrderUpdate) AddUID(i int64) *OrderUpdate {
	ou.mutation.AddUID(i)
	return ou
}

// SetOrderItems sets the "order_items" field.
func (ou *OrderUpdate) SetOrderItems(si []schema.OrderItem) *OrderUpdate {
	ou.mutation.SetOrderItems(si)
	return ou
}

// AppendOrderItems appends si to the "order_items" field.
func (ou *OrderUpdate) AppendOrderItems(si []schema.OrderItem) *OrderUpdate {
	ou.mutation.AppendOrderItems(si)
	return ou
}

// ClearOrderItems clears the value of the "order_items" field.
func (ou *OrderUpdate) ClearOrderItems() *OrderUpdate {
	ou.mutation.ClearOrderItems()
	return ou
}

// SetState sets the "state" field.
func (ou *OrderUpdate) SetState(i int32) *OrderUpdate {
	ou.mutation.ResetState()
	ou.mutation.SetState(i)
	return ou
}

// AddState adds i to the "state" field.
func (ou *OrderUpdate) AddState(i int32) *OrderUpdate {
	ou.mutation.AddState(i)
	return ou
}

// SetTotalAmount sets the "total_amount" field.
func (ou *OrderUpdate) SetTotalAmount(i int64) *OrderUpdate {
	ou.mutation.ResetTotalAmount()
	ou.mutation.SetTotalAmount(i)
	return ou
}

// SetNillableTotalAmount sets the "total_amount" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableTotalAmount(i *int64) *OrderUpdate {
	if i != nil {
		ou.SetTotalAmount(*i)
	}
	return ou
}

// AddTotalAmount adds i to the "total_amount" field.
func (ou *OrderUpdate) AddTotalAmount(i int64) *OrderUpdate {
	ou.mutation.AddTotalAmount(i)
	return ou
}

// ClearTotalAmount clears the value of the "total_amount" field.
func (ou *OrderUpdate) ClearTotalAmount() *OrderUpdate {
	ou.mutation.ClearTotalAmount()
	return ou
}

// SetTotalQuantity sets the "total_quantity" field.
func (ou *OrderUpdate) SetTotalQuantity(i int64) *OrderUpdate {
	ou.mutation.ResetTotalQuantity()
	ou.mutation.SetTotalQuantity(i)
	return ou
}

// SetNillableTotalQuantity sets the "total_quantity" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableTotalQuantity(i *int64) *OrderUpdate {
	if i != nil {
		ou.SetTotalQuantity(*i)
	}
	return ou
}

// AddTotalQuantity adds i to the "total_quantity" field.
func (ou *OrderUpdate) AddTotalQuantity(i int64) *OrderUpdate {
	ou.mutation.AddTotalQuantity(i)
	return ou
}

// ClearTotalQuantity clears the value of the "total_quantity" field.
func (ou *OrderUpdate) ClearTotalQuantity() *OrderUpdate {
	ou.mutation.ClearTotalQuantity()
	return ou
}

// SetUpdatedAt sets the "updated_at" field.
func (ou *OrderUpdate) SetUpdatedAt(t time.Time) *OrderUpdate {
	ou.mutation.SetUpdatedAt(t)
	return ou
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	ou.defaults()
	return withHooks[int, OrderMutation](ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OrderUpdate) defaults() {
	if _, ok := ou.mutation.UpdatedAt(); !ok {
		v := order.UpdateDefaultUpdatedAt()
		ou.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrderUpdate) check() error {
	if v, ok := ou.mutation.UID(); ok {
		if err := order.UIDValidator(v); err != nil {
			return &ValidationError{Name: "uid", err: fmt.Errorf(`ent: validator failed for field "Order.uid": %w`, err)}
		}
	}
	return nil
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.UID(); ok {
		_spec.SetField(order.FieldUID, field.TypeInt64, value)
	}
	if value, ok := ou.mutation.AddedUID(); ok {
		_spec.AddField(order.FieldUID, field.TypeInt64, value)
	}
	if value, ok := ou.mutation.OrderItems(); ok {
		_spec.SetField(order.FieldOrderItems, field.TypeJSON, value)
	}
	if value, ok := ou.mutation.AppendedOrderItems(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, order.FieldOrderItems, value)
		})
	}
	if ou.mutation.OrderItemsCleared() {
		_spec.ClearField(order.FieldOrderItems, field.TypeJSON)
	}
	if value, ok := ou.mutation.State(); ok {
		_spec.SetField(order.FieldState, field.TypeInt32, value)
	}
	if value, ok := ou.mutation.AddedState(); ok {
		_spec.AddField(order.FieldState, field.TypeInt32, value)
	}
	if value, ok := ou.mutation.TotalAmount(); ok {
		_spec.SetField(order.FieldTotalAmount, field.TypeInt64, value)
	}
	if value, ok := ou.mutation.AddedTotalAmount(); ok {
		_spec.AddField(order.FieldTotalAmount, field.TypeInt64, value)
	}
	if ou.mutation.TotalAmountCleared() {
		_spec.ClearField(order.FieldTotalAmount, field.TypeInt64)
	}
	if value, ok := ou.mutation.TotalQuantity(); ok {
		_spec.SetField(order.FieldTotalQuantity, field.TypeInt64, value)
	}
	if value, ok := ou.mutation.AddedTotalQuantity(); ok {
		_spec.AddField(order.FieldTotalQuantity, field.TypeInt64, value)
	}
	if ou.mutation.TotalQuantityCleared() {
		_spec.ClearField(order.FieldTotalQuantity, field.TypeInt64)
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetUID sets the "uid" field.
func (ouo *OrderUpdateOne) SetUID(i int64) *OrderUpdateOne {
	ouo.mutation.ResetUID()
	ouo.mutation.SetUID(i)
	return ouo
}

// AddUID adds i to the "uid" field.
func (ouo *OrderUpdateOne) AddUID(i int64) *OrderUpdateOne {
	ouo.mutation.AddUID(i)
	return ouo
}

// SetOrderItems sets the "order_items" field.
func (ouo *OrderUpdateOne) SetOrderItems(si []schema.OrderItem) *OrderUpdateOne {
	ouo.mutation.SetOrderItems(si)
	return ouo
}

// AppendOrderItems appends si to the "order_items" field.
func (ouo *OrderUpdateOne) AppendOrderItems(si []schema.OrderItem) *OrderUpdateOne {
	ouo.mutation.AppendOrderItems(si)
	return ouo
}

// ClearOrderItems clears the value of the "order_items" field.
func (ouo *OrderUpdateOne) ClearOrderItems() *OrderUpdateOne {
	ouo.mutation.ClearOrderItems()
	return ouo
}

// SetState sets the "state" field.
func (ouo *OrderUpdateOne) SetState(i int32) *OrderUpdateOne {
	ouo.mutation.ResetState()
	ouo.mutation.SetState(i)
	return ouo
}

// AddState adds i to the "state" field.
func (ouo *OrderUpdateOne) AddState(i int32) *OrderUpdateOne {
	ouo.mutation.AddState(i)
	return ouo
}

// SetTotalAmount sets the "total_amount" field.
func (ouo *OrderUpdateOne) SetTotalAmount(i int64) *OrderUpdateOne {
	ouo.mutation.ResetTotalAmount()
	ouo.mutation.SetTotalAmount(i)
	return ouo
}

// SetNillableTotalAmount sets the "total_amount" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableTotalAmount(i *int64) *OrderUpdateOne {
	if i != nil {
		ouo.SetTotalAmount(*i)
	}
	return ouo
}

// AddTotalAmount adds i to the "total_amount" field.
func (ouo *OrderUpdateOne) AddTotalAmount(i int64) *OrderUpdateOne {
	ouo.mutation.AddTotalAmount(i)
	return ouo
}

// ClearTotalAmount clears the value of the "total_amount" field.
func (ouo *OrderUpdateOne) ClearTotalAmount() *OrderUpdateOne {
	ouo.mutation.ClearTotalAmount()
	return ouo
}

// SetTotalQuantity sets the "total_quantity" field.
func (ouo *OrderUpdateOne) SetTotalQuantity(i int64) *OrderUpdateOne {
	ouo.mutation.ResetTotalQuantity()
	ouo.mutation.SetTotalQuantity(i)
	return ouo
}

// SetNillableTotalQuantity sets the "total_quantity" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableTotalQuantity(i *int64) *OrderUpdateOne {
	if i != nil {
		ouo.SetTotalQuantity(*i)
	}
	return ouo
}

// AddTotalQuantity adds i to the "total_quantity" field.
func (ouo *OrderUpdateOne) AddTotalQuantity(i int64) *OrderUpdateOne {
	ouo.mutation.AddTotalQuantity(i)
	return ouo
}

// ClearTotalQuantity clears the value of the "total_quantity" field.
func (ouo *OrderUpdateOne) ClearTotalQuantity() *OrderUpdateOne {
	ouo.mutation.ClearTotalQuantity()
	return ouo
}

// SetUpdatedAt sets the "updated_at" field.
func (ouo *OrderUpdateOne) SetUpdatedAt(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetUpdatedAt(t)
	return ouo
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ouo *OrderUpdateOne) Where(ps ...predicate.Order) *OrderUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	ouo.defaults()
	return withHooks[*Order, OrderMutation](ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OrderUpdateOne) defaults() {
	if _, ok := ouo.mutation.UpdatedAt(); !ok {
		v := order.UpdateDefaultUpdatedAt()
		ouo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrderUpdateOne) check() error {
	if v, ok := ouo.mutation.UID(); ok {
		if err := order.UIDValidator(v); err != nil {
			return &ValidationError{Name: "uid", err: fmt.Errorf(`ent: validator failed for field "Order.uid": %w`, err)}
		}
	}
	return nil
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.UID(); ok {
		_spec.SetField(order.FieldUID, field.TypeInt64, value)
	}
	if value, ok := ouo.mutation.AddedUID(); ok {
		_spec.AddField(order.FieldUID, field.TypeInt64, value)
	}
	if value, ok := ouo.mutation.OrderItems(); ok {
		_spec.SetField(order.FieldOrderItems, field.TypeJSON, value)
	}
	if value, ok := ouo.mutation.AppendedOrderItems(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, order.FieldOrderItems, value)
		})
	}
	if ouo.mutation.OrderItemsCleared() {
		_spec.ClearField(order.FieldOrderItems, field.TypeJSON)
	}
	if value, ok := ouo.mutation.State(); ok {
		_spec.SetField(order.FieldState, field.TypeInt32, value)
	}
	if value, ok := ouo.mutation.AddedState(); ok {
		_spec.AddField(order.FieldState, field.TypeInt32, value)
	}
	if value, ok := ouo.mutation.TotalAmount(); ok {
		_spec.SetField(order.FieldTotalAmount, field.TypeInt64, value)
	}
	if value, ok := ouo.mutation.AddedTotalAmount(); ok {
		_spec.AddField(order.FieldTotalAmount, field.TypeInt64, value)
	}
	if ouo.mutation.TotalAmountCleared() {
		_spec.ClearField(order.FieldTotalAmount, field.TypeInt64)
	}
	if value, ok := ouo.mutation.TotalQuantity(); ok {
		_spec.SetField(order.FieldTotalQuantity, field.TypeInt64, value)
	}
	if value, ok := ouo.mutation.AddedTotalQuantity(); ok {
		_spec.AddField(order.FieldTotalQuantity, field.TypeInt64, value)
	}
	if ouo.mutation.TotalQuantityCleared() {
		_spec.ClearField(order.FieldTotalQuantity, field.TypeInt64)
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}