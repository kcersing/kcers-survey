// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kcers-survey/biz/dal/db/mysql/ent/predicate"
	"kcers-survey/biz/dal/db/mysql/ent/token"
	"kcers-survey/biz/dal/db/mysql/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TokenUpdate is the builder for updating Token entities.
type TokenUpdate struct {
	config
	hooks     []Hook
	mutation  *TokenMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TokenUpdate builder.
func (tu *TokenUpdate) Where(ps ...predicate.Token) *TokenUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TokenUpdate) SetUpdatedAt(t time.Time) *TokenUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tu *TokenUpdate) ClearUpdatedAt() *TokenUpdate {
	tu.mutation.ClearUpdatedAt()
	return tu
}

// SetDelete sets the "delete" field.
func (tu *TokenUpdate) SetDelete(i int64) *TokenUpdate {
	tu.mutation.ResetDelete()
	tu.mutation.SetDelete(i)
	return tu
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableDelete(i *int64) *TokenUpdate {
	if i != nil {
		tu.SetDelete(*i)
	}
	return tu
}

// AddDelete adds i to the "delete" field.
func (tu *TokenUpdate) AddDelete(i int64) *TokenUpdate {
	tu.mutation.AddDelete(i)
	return tu
}

// ClearDelete clears the value of the "delete" field.
func (tu *TokenUpdate) ClearDelete() *TokenUpdate {
	tu.mutation.ClearDelete()
	return tu
}

// SetCreatedID sets the "created_id" field.
func (tu *TokenUpdate) SetCreatedID(i int64) *TokenUpdate {
	tu.mutation.ResetCreatedID()
	tu.mutation.SetCreatedID(i)
	return tu
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableCreatedID(i *int64) *TokenUpdate {
	if i != nil {
		tu.SetCreatedID(*i)
	}
	return tu
}

// AddCreatedID adds i to the "created_id" field.
func (tu *TokenUpdate) AddCreatedID(i int64) *TokenUpdate {
	tu.mutation.AddCreatedID(i)
	return tu
}

// ClearCreatedID clears the value of the "created_id" field.
func (tu *TokenUpdate) ClearCreatedID() *TokenUpdate {
	tu.mutation.ClearCreatedID()
	return tu
}

// SetUserID sets the "user_id" field.
func (tu *TokenUpdate) SetUserID(i int64) *TokenUpdate {
	tu.mutation.ResetUserID()
	tu.mutation.SetUserID(i)
	return tu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableUserID(i *int64) *TokenUpdate {
	if i != nil {
		tu.SetUserID(*i)
	}
	return tu
}

// AddUserID adds i to the "user_id" field.
func (tu *TokenUpdate) AddUserID(i int64) *TokenUpdate {
	tu.mutation.AddUserID(i)
	return tu
}

// SetToken sets the "token" field.
func (tu *TokenUpdate) SetToken(s string) *TokenUpdate {
	tu.mutation.SetToken(s)
	return tu
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableToken(s *string) *TokenUpdate {
	if s != nil {
		tu.SetToken(*s)
	}
	return tu
}

// SetType sets the "type" field.
func (tu *TokenUpdate) SetType(i int64) *TokenUpdate {
	tu.mutation.ResetType()
	tu.mutation.SetType(i)
	return tu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableType(i *int64) *TokenUpdate {
	if i != nil {
		tu.SetType(*i)
	}
	return tu
}

// AddType adds i to the "type" field.
func (tu *TokenUpdate) AddType(i int64) *TokenUpdate {
	tu.mutation.AddType(i)
	return tu
}

// ClearType clears the value of the "type" field.
func (tu *TokenUpdate) ClearType() *TokenUpdate {
	tu.mutation.ClearType()
	return tu
}

// SetSource sets the "source" field.
func (tu *TokenUpdate) SetSource(s string) *TokenUpdate {
	tu.mutation.SetSource(s)
	return tu
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableSource(s *string) *TokenUpdate {
	if s != nil {
		tu.SetSource(*s)
	}
	return tu
}

// SetExpiredAt sets the "expired_at" field.
func (tu *TokenUpdate) SetExpiredAt(t time.Time) *TokenUpdate {
	tu.mutation.SetExpiredAt(t)
	return tu
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableExpiredAt(t *time.Time) *TokenUpdate {
	if t != nil {
		tu.SetExpiredAt(*t)
	}
	return tu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (tu *TokenUpdate) SetOwnerID(id int64) *TokenUpdate {
	tu.mutation.SetOwnerID(id)
	return tu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (tu *TokenUpdate) SetNillableOwnerID(id *int64) *TokenUpdate {
	if id != nil {
		tu = tu.SetOwnerID(*id)
	}
	return tu
}

// SetOwner sets the "owner" edge to the User entity.
func (tu *TokenUpdate) SetOwner(u *User) *TokenUpdate {
	return tu.SetOwnerID(u.ID)
}

// Mutation returns the TokenMutation object of the builder.
func (tu *TokenUpdate) Mutation() *TokenMutation {
	return tu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (tu *TokenUpdate) ClearOwner() *TokenUpdate {
	tu.mutation.ClearOwner()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TokenUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TokenUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TokenUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TokenUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TokenUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok && !tu.mutation.UpdatedAtCleared() {
		v := token.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tu *TokenUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TokenUpdate {
	tu.modifiers = append(tu.modifiers, modifiers...)
	return tu
}

func (tu *TokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(token.Table, token.Columns, sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt64))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tu.mutation.CreatedAtCleared() {
		_spec.ClearField(token.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(token.FieldUpdatedAt, field.TypeTime, value)
	}
	if tu.mutation.UpdatedAtCleared() {
		_spec.ClearField(token.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.Delete(); ok {
		_spec.SetField(token.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedDelete(); ok {
		_spec.AddField(token.FieldDelete, field.TypeInt64, value)
	}
	if tu.mutation.DeleteCleared() {
		_spec.ClearField(token.FieldDelete, field.TypeInt64)
	}
	if value, ok := tu.mutation.CreatedID(); ok {
		_spec.SetField(token.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedCreatedID(); ok {
		_spec.AddField(token.FieldCreatedID, field.TypeInt64, value)
	}
	if tu.mutation.CreatedIDCleared() {
		_spec.ClearField(token.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := tu.mutation.UserID(); ok {
		_spec.SetField(token.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedUserID(); ok {
		_spec.AddField(token.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.Token(); ok {
		_spec.SetField(token.FieldToken, field.TypeString, value)
	}
	if value, ok := tu.mutation.GetType(); ok {
		_spec.SetField(token.FieldType, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedType(); ok {
		_spec.AddField(token.FieldType, field.TypeInt64, value)
	}
	if tu.mutation.TypeCleared() {
		_spec.ClearField(token.FieldType, field.TypeInt64)
	}
	if value, ok := tu.mutation.Source(); ok {
		_spec.SetField(token.FieldSource, field.TypeString, value)
	}
	if value, ok := tu.mutation.ExpiredAt(); ok {
		_spec.SetField(token.FieldExpiredAt, field.TypeTime, value)
	}
	if tu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   token.OwnerTable,
			Columns: []string{token.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   token.OwnerTable,
			Columns: []string{token.OwnerColumn},
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
	_spec.AddModifiers(tu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TokenUpdateOne is the builder for updating a single Token entity.
type TokenUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TokenMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TokenUpdateOne) SetUpdatedAt(t time.Time) *TokenUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tuo *TokenUpdateOne) ClearUpdatedAt() *TokenUpdateOne {
	tuo.mutation.ClearUpdatedAt()
	return tuo
}

// SetDelete sets the "delete" field.
func (tuo *TokenUpdateOne) SetDelete(i int64) *TokenUpdateOne {
	tuo.mutation.ResetDelete()
	tuo.mutation.SetDelete(i)
	return tuo
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableDelete(i *int64) *TokenUpdateOne {
	if i != nil {
		tuo.SetDelete(*i)
	}
	return tuo
}

// AddDelete adds i to the "delete" field.
func (tuo *TokenUpdateOne) AddDelete(i int64) *TokenUpdateOne {
	tuo.mutation.AddDelete(i)
	return tuo
}

// ClearDelete clears the value of the "delete" field.
func (tuo *TokenUpdateOne) ClearDelete() *TokenUpdateOne {
	tuo.mutation.ClearDelete()
	return tuo
}

// SetCreatedID sets the "created_id" field.
func (tuo *TokenUpdateOne) SetCreatedID(i int64) *TokenUpdateOne {
	tuo.mutation.ResetCreatedID()
	tuo.mutation.SetCreatedID(i)
	return tuo
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableCreatedID(i *int64) *TokenUpdateOne {
	if i != nil {
		tuo.SetCreatedID(*i)
	}
	return tuo
}

// AddCreatedID adds i to the "created_id" field.
func (tuo *TokenUpdateOne) AddCreatedID(i int64) *TokenUpdateOne {
	tuo.mutation.AddCreatedID(i)
	return tuo
}

// ClearCreatedID clears the value of the "created_id" field.
func (tuo *TokenUpdateOne) ClearCreatedID() *TokenUpdateOne {
	tuo.mutation.ClearCreatedID()
	return tuo
}

// SetUserID sets the "user_id" field.
func (tuo *TokenUpdateOne) SetUserID(i int64) *TokenUpdateOne {
	tuo.mutation.ResetUserID()
	tuo.mutation.SetUserID(i)
	return tuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableUserID(i *int64) *TokenUpdateOne {
	if i != nil {
		tuo.SetUserID(*i)
	}
	return tuo
}

// AddUserID adds i to the "user_id" field.
func (tuo *TokenUpdateOne) AddUserID(i int64) *TokenUpdateOne {
	tuo.mutation.AddUserID(i)
	return tuo
}

// SetToken sets the "token" field.
func (tuo *TokenUpdateOne) SetToken(s string) *TokenUpdateOne {
	tuo.mutation.SetToken(s)
	return tuo
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableToken(s *string) *TokenUpdateOne {
	if s != nil {
		tuo.SetToken(*s)
	}
	return tuo
}

// SetType sets the "type" field.
func (tuo *TokenUpdateOne) SetType(i int64) *TokenUpdateOne {
	tuo.mutation.ResetType()
	tuo.mutation.SetType(i)
	return tuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableType(i *int64) *TokenUpdateOne {
	if i != nil {
		tuo.SetType(*i)
	}
	return tuo
}

// AddType adds i to the "type" field.
func (tuo *TokenUpdateOne) AddType(i int64) *TokenUpdateOne {
	tuo.mutation.AddType(i)
	return tuo
}

// ClearType clears the value of the "type" field.
func (tuo *TokenUpdateOne) ClearType() *TokenUpdateOne {
	tuo.mutation.ClearType()
	return tuo
}

// SetSource sets the "source" field.
func (tuo *TokenUpdateOne) SetSource(s string) *TokenUpdateOne {
	tuo.mutation.SetSource(s)
	return tuo
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableSource(s *string) *TokenUpdateOne {
	if s != nil {
		tuo.SetSource(*s)
	}
	return tuo
}

// SetExpiredAt sets the "expired_at" field.
func (tuo *TokenUpdateOne) SetExpiredAt(t time.Time) *TokenUpdateOne {
	tuo.mutation.SetExpiredAt(t)
	return tuo
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableExpiredAt(t *time.Time) *TokenUpdateOne {
	if t != nil {
		tuo.SetExpiredAt(*t)
	}
	return tuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (tuo *TokenUpdateOne) SetOwnerID(id int64) *TokenUpdateOne {
	tuo.mutation.SetOwnerID(id)
	return tuo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableOwnerID(id *int64) *TokenUpdateOne {
	if id != nil {
		tuo = tuo.SetOwnerID(*id)
	}
	return tuo
}

// SetOwner sets the "owner" edge to the User entity.
func (tuo *TokenUpdateOne) SetOwner(u *User) *TokenUpdateOne {
	return tuo.SetOwnerID(u.ID)
}

// Mutation returns the TokenMutation object of the builder.
func (tuo *TokenUpdateOne) Mutation() *TokenMutation {
	return tuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (tuo *TokenUpdateOne) ClearOwner() *TokenUpdateOne {
	tuo.mutation.ClearOwner()
	return tuo
}

// Where appends a list predicates to the TokenUpdate builder.
func (tuo *TokenUpdateOne) Where(ps ...predicate.Token) *TokenUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TokenUpdateOne) Select(field string, fields ...string) *TokenUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Token entity.
func (tuo *TokenUpdateOne) Save(ctx context.Context) (*Token, error) {
	tuo.defaults()
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TokenUpdateOne) SaveX(ctx context.Context) *Token {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TokenUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TokenUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TokenUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok && !tuo.mutation.UpdatedAtCleared() {
		v := token.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuo *TokenUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TokenUpdateOne {
	tuo.modifiers = append(tuo.modifiers, modifiers...)
	return tuo
}

func (tuo *TokenUpdateOne) sqlSave(ctx context.Context) (_node *Token, err error) {
	_spec := sqlgraph.NewUpdateSpec(token.Table, token.Columns, sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt64))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Token.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, token.FieldID)
		for _, f := range fields {
			if !token.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != token.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tuo.mutation.CreatedAtCleared() {
		_spec.ClearField(token.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(token.FieldUpdatedAt, field.TypeTime, value)
	}
	if tuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(token.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.Delete(); ok {
		_spec.SetField(token.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedDelete(); ok {
		_spec.AddField(token.FieldDelete, field.TypeInt64, value)
	}
	if tuo.mutation.DeleteCleared() {
		_spec.ClearField(token.FieldDelete, field.TypeInt64)
	}
	if value, ok := tuo.mutation.CreatedID(); ok {
		_spec.SetField(token.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedCreatedID(); ok {
		_spec.AddField(token.FieldCreatedID, field.TypeInt64, value)
	}
	if tuo.mutation.CreatedIDCleared() {
		_spec.ClearField(token.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := tuo.mutation.UserID(); ok {
		_spec.SetField(token.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedUserID(); ok {
		_spec.AddField(token.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.Token(); ok {
		_spec.SetField(token.FieldToken, field.TypeString, value)
	}
	if value, ok := tuo.mutation.GetType(); ok {
		_spec.SetField(token.FieldType, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedType(); ok {
		_spec.AddField(token.FieldType, field.TypeInt64, value)
	}
	if tuo.mutation.TypeCleared() {
		_spec.ClearField(token.FieldType, field.TypeInt64)
	}
	if value, ok := tuo.mutation.Source(); ok {
		_spec.SetField(token.FieldSource, field.TypeString, value)
	}
	if value, ok := tuo.mutation.ExpiredAt(); ok {
		_spec.SetField(token.FieldExpiredAt, field.TypeTime, value)
	}
	if tuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   token.OwnerTable,
			Columns: []string{token.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   token.OwnerTable,
			Columns: []string{token.OwnerColumn},
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
	_spec.AddModifiers(tuo.modifiers...)
	_node = &Token{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
