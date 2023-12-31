// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"redbook/app/domain/user/internal/data/ent/predicate"
	"redbook/app/domain/user/internal/data/ent/profile"
	"redbook/app/domain/user/internal/data/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProfileUpdate is the builder for updating Profile entities.
type ProfileUpdate struct {
	config
	hooks    []Hook
	mutation *ProfileMutation
}

// Where appends a list predicates to the ProfileUpdate builder.
func (pu *ProfileUpdate) Where(ps ...predicate.Profile) *ProfileUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUID sets the "uid" field.
func (pu *ProfileUpdate) SetUID(i int64) *ProfileUpdate {
	pu.mutation.ResetUID()
	pu.mutation.SetUID(i)
	return pu
}

// AddUID adds i to the "uid" field.
func (pu *ProfileUpdate) AddUID(i int64) *ProfileUpdate {
	pu.mutation.AddUID(i)
	return pu
}

// SetName sets the "name" field.
func (pu *ProfileUpdate) SetName(s string) *ProfileUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetSex sets the "sex" field.
func (pu *ProfileUpdate) SetSex(i int32) *ProfileUpdate {
	pu.mutation.ResetSex()
	pu.mutation.SetSex(i)
	return pu
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableSex(i *int32) *ProfileUpdate {
	if i != nil {
		pu.SetSex(*i)
	}
	return pu
}

// AddSex adds i to the "sex" field.
func (pu *ProfileUpdate) AddSex(i int32) *ProfileUpdate {
	pu.mutation.AddSex(i)
	return pu
}

// SetAvatar sets the "avatar" field.
func (pu *ProfileUpdate) SetAvatar(s string) *ProfileUpdate {
	pu.mutation.SetAvatar(s)
	return pu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableAvatar(s *string) *ProfileUpdate {
	if s != nil {
		pu.SetAvatar(*s)
	}
	return pu
}

// SetSign sets the "sign" field.
func (pu *ProfileUpdate) SetSign(s string) *ProfileUpdate {
	pu.mutation.SetSign(s)
	return pu
}

// SetNillableSign sets the "sign" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableSign(s *string) *ProfileUpdate {
	if s != nil {
		pu.SetSign(*s)
	}
	return pu
}

// ClearSign clears the value of the "sign" field.
func (pu *ProfileUpdate) ClearSign() *ProfileUpdate {
	pu.mutation.ClearSign()
	return pu
}

// SetBirthday sets the "birthday" field.
func (pu *ProfileUpdate) SetBirthday(t time.Time) *ProfileUpdate {
	pu.mutation.SetBirthday(t)
	return pu
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableBirthday(t *time.Time) *ProfileUpdate {
	if t != nil {
		pu.SetBirthday(*t)
	}
	return pu
}

// ClearBirthday clears the value of the "birthday" field.
func (pu *ProfileUpdate) ClearBirthday() *ProfileUpdate {
	pu.mutation.ClearBirthday()
	return pu
}

// SetLevel sets the "level" field.
func (pu *ProfileUpdate) SetLevel(i int32) *ProfileUpdate {
	pu.mutation.ResetLevel()
	pu.mutation.SetLevel(i)
	return pu
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableLevel(i *int32) *ProfileUpdate {
	if i != nil {
		pu.SetLevel(*i)
	}
	return pu
}

// AddLevel adds i to the "level" field.
func (pu *ProfileUpdate) AddLevel(i int32) *ProfileUpdate {
	pu.mutation.AddLevel(i)
	return pu
}

// SetVerifyType sets the "verify_type" field.
func (pu *ProfileUpdate) SetVerifyType(i int8) *ProfileUpdate {
	pu.mutation.ResetVerifyType()
	pu.mutation.SetVerifyType(i)
	return pu
}

// SetNillableVerifyType sets the "verify_type" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableVerifyType(i *int8) *ProfileUpdate {
	if i != nil {
		pu.SetVerifyType(*i)
	}
	return pu
}

// AddVerifyType adds i to the "verify_type" field.
func (pu *ProfileUpdate) AddVerifyType(i int8) *ProfileUpdate {
	pu.mutation.AddVerifyType(i)
	return pu
}

// SetAttr sets the "attr" field.
func (pu *ProfileUpdate) SetAttr(i int32) *ProfileUpdate {
	pu.mutation.ResetAttr()
	pu.mutation.SetAttr(i)
	return pu
}

// SetNillableAttr sets the "attr" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableAttr(i *int32) *ProfileUpdate {
	if i != nil {
		pu.SetAttr(*i)
	}
	return pu
}

// AddAttr adds i to the "attr" field.
func (pu *ProfileUpdate) AddAttr(i int32) *ProfileUpdate {
	pu.mutation.AddAttr(i)
	return pu
}

// SetState sets the "state" field.
func (pu *ProfileUpdate) SetState(i int32) *ProfileUpdate {
	pu.mutation.ResetState()
	pu.mutation.SetState(i)
	return pu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableState(i *int32) *ProfileUpdate {
	if i != nil {
		pu.SetState(*i)
	}
	return pu
}

// AddState adds i to the "state" field.
func (pu *ProfileUpdate) AddState(i int32) *ProfileUpdate {
	pu.mutation.AddState(i)
	return pu
}

// SetDeleted sets the "deleted" field.
func (pu *ProfileUpdate) SetDeleted(b bool) *ProfileUpdate {
	pu.mutation.SetDeleted(b)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *ProfileUpdate) SetUpdatedAt(t time.Time) *ProfileUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pu *ProfileUpdate) SetUserID(id int64) *ProfileUpdate {
	pu.mutation.SetUserID(id)
	return pu
}

// SetUser sets the "user" edge to the User entity.
func (pu *ProfileUpdate) SetUser(u *User) *ProfileUpdate {
	return pu.SetUserID(u.ID)
}

// Mutation returns the ProfileMutation object of the builder.
func (pu *ProfileUpdate) Mutation() *ProfileMutation {
	return pu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pu *ProfileUpdate) ClearUser() *ProfileUpdate {
	pu.mutation.ClearUser()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProfileUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks[int, ProfileMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProfileUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProfileUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProfileUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := profile.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProfileUpdate) check() error {
	if _, ok := pu.mutation.UserID(); pu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Profile.user"`)
	}
	return nil
}

func (pu *ProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(profile.Table, profile.Columns, sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt64))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UID(); ok {
		_spec.SetField(profile.FieldUID, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedUID(); ok {
		_spec.AddField(profile.FieldUID, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(profile.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Sex(); ok {
		_spec.SetField(profile.FieldSex, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.AddedSex(); ok {
		_spec.AddField(profile.FieldSex, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.Avatar(); ok {
		_spec.SetField(profile.FieldAvatar, field.TypeString, value)
	}
	if value, ok := pu.mutation.Sign(); ok {
		_spec.SetField(profile.FieldSign, field.TypeString, value)
	}
	if pu.mutation.SignCleared() {
		_spec.ClearField(profile.FieldSign, field.TypeString)
	}
	if value, ok := pu.mutation.Birthday(); ok {
		_spec.SetField(profile.FieldBirthday, field.TypeTime, value)
	}
	if pu.mutation.BirthdayCleared() {
		_spec.ClearField(profile.FieldBirthday, field.TypeTime)
	}
	if value, ok := pu.mutation.Level(); ok {
		_spec.SetField(profile.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.AddedLevel(); ok {
		_spec.AddField(profile.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.VerifyType(); ok {
		_spec.SetField(profile.FieldVerifyType, field.TypeInt8, value)
	}
	if value, ok := pu.mutation.AddedVerifyType(); ok {
		_spec.AddField(profile.FieldVerifyType, field.TypeInt8, value)
	}
	if value, ok := pu.mutation.Attr(); ok {
		_spec.SetField(profile.FieldAttr, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.AddedAttr(); ok {
		_spec.AddField(profile.FieldAttr, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.State(); ok {
		_spec.SetField(profile.FieldState, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.AddedState(); ok {
		_spec.AddField(profile.FieldState, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.Deleted(); ok {
		_spec.SetField(profile.FieldDeleted, field.TypeBool, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(profile.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profile.UserTable,
			Columns: []string{profile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profile.UserTable,
			Columns: []string{profile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProfileUpdateOne is the builder for updating a single Profile entity.
type ProfileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProfileMutation
}

// SetUID sets the "uid" field.
func (puo *ProfileUpdateOne) SetUID(i int64) *ProfileUpdateOne {
	puo.mutation.ResetUID()
	puo.mutation.SetUID(i)
	return puo
}

// AddUID adds i to the "uid" field.
func (puo *ProfileUpdateOne) AddUID(i int64) *ProfileUpdateOne {
	puo.mutation.AddUID(i)
	return puo
}

// SetName sets the "name" field.
func (puo *ProfileUpdateOne) SetName(s string) *ProfileUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetSex sets the "sex" field.
func (puo *ProfileUpdateOne) SetSex(i int32) *ProfileUpdateOne {
	puo.mutation.ResetSex()
	puo.mutation.SetSex(i)
	return puo
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableSex(i *int32) *ProfileUpdateOne {
	if i != nil {
		puo.SetSex(*i)
	}
	return puo
}

// AddSex adds i to the "sex" field.
func (puo *ProfileUpdateOne) AddSex(i int32) *ProfileUpdateOne {
	puo.mutation.AddSex(i)
	return puo
}

// SetAvatar sets the "avatar" field.
func (puo *ProfileUpdateOne) SetAvatar(s string) *ProfileUpdateOne {
	puo.mutation.SetAvatar(s)
	return puo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableAvatar(s *string) *ProfileUpdateOne {
	if s != nil {
		puo.SetAvatar(*s)
	}
	return puo
}

// SetSign sets the "sign" field.
func (puo *ProfileUpdateOne) SetSign(s string) *ProfileUpdateOne {
	puo.mutation.SetSign(s)
	return puo
}

// SetNillableSign sets the "sign" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableSign(s *string) *ProfileUpdateOne {
	if s != nil {
		puo.SetSign(*s)
	}
	return puo
}

// ClearSign clears the value of the "sign" field.
func (puo *ProfileUpdateOne) ClearSign() *ProfileUpdateOne {
	puo.mutation.ClearSign()
	return puo
}

// SetBirthday sets the "birthday" field.
func (puo *ProfileUpdateOne) SetBirthday(t time.Time) *ProfileUpdateOne {
	puo.mutation.SetBirthday(t)
	return puo
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableBirthday(t *time.Time) *ProfileUpdateOne {
	if t != nil {
		puo.SetBirthday(*t)
	}
	return puo
}

// ClearBirthday clears the value of the "birthday" field.
func (puo *ProfileUpdateOne) ClearBirthday() *ProfileUpdateOne {
	puo.mutation.ClearBirthday()
	return puo
}

// SetLevel sets the "level" field.
func (puo *ProfileUpdateOne) SetLevel(i int32) *ProfileUpdateOne {
	puo.mutation.ResetLevel()
	puo.mutation.SetLevel(i)
	return puo
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableLevel(i *int32) *ProfileUpdateOne {
	if i != nil {
		puo.SetLevel(*i)
	}
	return puo
}

// AddLevel adds i to the "level" field.
func (puo *ProfileUpdateOne) AddLevel(i int32) *ProfileUpdateOne {
	puo.mutation.AddLevel(i)
	return puo
}

// SetVerifyType sets the "verify_type" field.
func (puo *ProfileUpdateOne) SetVerifyType(i int8) *ProfileUpdateOne {
	puo.mutation.ResetVerifyType()
	puo.mutation.SetVerifyType(i)
	return puo
}

// SetNillableVerifyType sets the "verify_type" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableVerifyType(i *int8) *ProfileUpdateOne {
	if i != nil {
		puo.SetVerifyType(*i)
	}
	return puo
}

// AddVerifyType adds i to the "verify_type" field.
func (puo *ProfileUpdateOne) AddVerifyType(i int8) *ProfileUpdateOne {
	puo.mutation.AddVerifyType(i)
	return puo
}

// SetAttr sets the "attr" field.
func (puo *ProfileUpdateOne) SetAttr(i int32) *ProfileUpdateOne {
	puo.mutation.ResetAttr()
	puo.mutation.SetAttr(i)
	return puo
}

// SetNillableAttr sets the "attr" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableAttr(i *int32) *ProfileUpdateOne {
	if i != nil {
		puo.SetAttr(*i)
	}
	return puo
}

// AddAttr adds i to the "attr" field.
func (puo *ProfileUpdateOne) AddAttr(i int32) *ProfileUpdateOne {
	puo.mutation.AddAttr(i)
	return puo
}

// SetState sets the "state" field.
func (puo *ProfileUpdateOne) SetState(i int32) *ProfileUpdateOne {
	puo.mutation.ResetState()
	puo.mutation.SetState(i)
	return puo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableState(i *int32) *ProfileUpdateOne {
	if i != nil {
		puo.SetState(*i)
	}
	return puo
}

// AddState adds i to the "state" field.
func (puo *ProfileUpdateOne) AddState(i int32) *ProfileUpdateOne {
	puo.mutation.AddState(i)
	return puo
}

// SetDeleted sets the "deleted" field.
func (puo *ProfileUpdateOne) SetDeleted(b bool) *ProfileUpdateOne {
	puo.mutation.SetDeleted(b)
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *ProfileUpdateOne) SetUpdatedAt(t time.Time) *ProfileUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (puo *ProfileUpdateOne) SetUserID(id int64) *ProfileUpdateOne {
	puo.mutation.SetUserID(id)
	return puo
}

// SetUser sets the "user" edge to the User entity.
func (puo *ProfileUpdateOne) SetUser(u *User) *ProfileUpdateOne {
	return puo.SetUserID(u.ID)
}

// Mutation returns the ProfileMutation object of the builder.
func (puo *ProfileUpdateOne) Mutation() *ProfileMutation {
	return puo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (puo *ProfileUpdateOne) ClearUser() *ProfileUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// Where appends a list predicates to the ProfileUpdate builder.
func (puo *ProfileUpdateOne) Where(ps ...predicate.Profile) *ProfileUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProfileUpdateOne) Select(field string, fields ...string) *ProfileUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Profile entity.
func (puo *ProfileUpdateOne) Save(ctx context.Context) (*Profile, error) {
	puo.defaults()
	return withHooks[*Profile, ProfileMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProfileUpdateOne) SaveX(ctx context.Context) *Profile {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProfileUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProfileUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := profile.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProfileUpdateOne) check() error {
	if _, ok := puo.mutation.UserID(); puo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Profile.user"`)
	}
	return nil
}

func (puo *ProfileUpdateOne) sqlSave(ctx context.Context) (_node *Profile, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(profile.Table, profile.Columns, sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt64))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Profile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, profile.FieldID)
		for _, f := range fields {
			if !profile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != profile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UID(); ok {
		_spec.SetField(profile.FieldUID, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedUID(); ok {
		_spec.AddField(profile.FieldUID, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(profile.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Sex(); ok {
		_spec.SetField(profile.FieldSex, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.AddedSex(); ok {
		_spec.AddField(profile.FieldSex, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.Avatar(); ok {
		_spec.SetField(profile.FieldAvatar, field.TypeString, value)
	}
	if value, ok := puo.mutation.Sign(); ok {
		_spec.SetField(profile.FieldSign, field.TypeString, value)
	}
	if puo.mutation.SignCleared() {
		_spec.ClearField(profile.FieldSign, field.TypeString)
	}
	if value, ok := puo.mutation.Birthday(); ok {
		_spec.SetField(profile.FieldBirthday, field.TypeTime, value)
	}
	if puo.mutation.BirthdayCleared() {
		_spec.ClearField(profile.FieldBirthday, field.TypeTime)
	}
	if value, ok := puo.mutation.Level(); ok {
		_spec.SetField(profile.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.AddedLevel(); ok {
		_spec.AddField(profile.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.VerifyType(); ok {
		_spec.SetField(profile.FieldVerifyType, field.TypeInt8, value)
	}
	if value, ok := puo.mutation.AddedVerifyType(); ok {
		_spec.AddField(profile.FieldVerifyType, field.TypeInt8, value)
	}
	if value, ok := puo.mutation.Attr(); ok {
		_spec.SetField(profile.FieldAttr, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.AddedAttr(); ok {
		_spec.AddField(profile.FieldAttr, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.State(); ok {
		_spec.SetField(profile.FieldState, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.AddedState(); ok {
		_spec.AddField(profile.FieldState, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.Deleted(); ok {
		_spec.SetField(profile.FieldDeleted, field.TypeBool, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(profile.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profile.UserTable,
			Columns: []string{profile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profile.UserTable,
			Columns: []string{profile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Profile{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
