// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kcers-survey/biz/dal/db/mysql/ent/dictionary"
	"kcers-survey/biz/dal/db/mysql/ent/dictionarydetail"
	"kcers-survey/biz/dal/db/mysql/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DictionaryDetailUpdate is the builder for updating DictionaryDetail entities.
type DictionaryDetailUpdate struct {
	config
	hooks     []Hook
	mutation  *DictionaryDetailMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the DictionaryDetailUpdate builder.
func (ddu *DictionaryDetailUpdate) Where(ps ...predicate.DictionaryDetail) *DictionaryDetailUpdate {
	ddu.mutation.Where(ps...)
	return ddu
}

// SetUpdatedAt sets the "updated_at" field.
func (ddu *DictionaryDetailUpdate) SetUpdatedAt(t time.Time) *DictionaryDetailUpdate {
	ddu.mutation.SetUpdatedAt(t)
	return ddu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ddu *DictionaryDetailUpdate) ClearUpdatedAt() *DictionaryDetailUpdate {
	ddu.mutation.ClearUpdatedAt()
	return ddu
}

// SetDelete sets the "delete" field.
func (ddu *DictionaryDetailUpdate) SetDelete(i int64) *DictionaryDetailUpdate {
	ddu.mutation.ResetDelete()
	ddu.mutation.SetDelete(i)
	return ddu
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (ddu *DictionaryDetailUpdate) SetNillableDelete(i *int64) *DictionaryDetailUpdate {
	if i != nil {
		ddu.SetDelete(*i)
	}
	return ddu
}

// AddDelete adds i to the "delete" field.
func (ddu *DictionaryDetailUpdate) AddDelete(i int64) *DictionaryDetailUpdate {
	ddu.mutation.AddDelete(i)
	return ddu
}

// ClearDelete clears the value of the "delete" field.
func (ddu *DictionaryDetailUpdate) ClearDelete() *DictionaryDetailUpdate {
	ddu.mutation.ClearDelete()
	return ddu
}

// SetCreatedID sets the "created_id" field.
func (ddu *DictionaryDetailUpdate) SetCreatedID(i int64) *DictionaryDetailUpdate {
	ddu.mutation.ResetCreatedID()
	ddu.mutation.SetCreatedID(i)
	return ddu
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (ddu *DictionaryDetailUpdate) SetNillableCreatedID(i *int64) *DictionaryDetailUpdate {
	if i != nil {
		ddu.SetCreatedID(*i)
	}
	return ddu
}

// AddCreatedID adds i to the "created_id" field.
func (ddu *DictionaryDetailUpdate) AddCreatedID(i int64) *DictionaryDetailUpdate {
	ddu.mutation.AddCreatedID(i)
	return ddu
}

// ClearCreatedID clears the value of the "created_id" field.
func (ddu *DictionaryDetailUpdate) ClearCreatedID() *DictionaryDetailUpdate {
	ddu.mutation.ClearCreatedID()
	return ddu
}

// SetStatus sets the "status" field.
func (ddu *DictionaryDetailUpdate) SetStatus(i int64) *DictionaryDetailUpdate {
	ddu.mutation.ResetStatus()
	ddu.mutation.SetStatus(i)
	return ddu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ddu *DictionaryDetailUpdate) SetNillableStatus(i *int64) *DictionaryDetailUpdate {
	if i != nil {
		ddu.SetStatus(*i)
	}
	return ddu
}

// AddStatus adds i to the "status" field.
func (ddu *DictionaryDetailUpdate) AddStatus(i int64) *DictionaryDetailUpdate {
	ddu.mutation.AddStatus(i)
	return ddu
}

// ClearStatus clears the value of the "status" field.
func (ddu *DictionaryDetailUpdate) ClearStatus() *DictionaryDetailUpdate {
	ddu.mutation.ClearStatus()
	return ddu
}

// SetTitle sets the "title" field.
func (ddu *DictionaryDetailUpdate) SetTitle(s string) *DictionaryDetailUpdate {
	ddu.mutation.SetTitle(s)
	return ddu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ddu *DictionaryDetailUpdate) SetNillableTitle(s *string) *DictionaryDetailUpdate {
	if s != nil {
		ddu.SetTitle(*s)
	}
	return ddu
}

// SetKey sets the "key" field.
func (ddu *DictionaryDetailUpdate) SetKey(s string) *DictionaryDetailUpdate {
	ddu.mutation.SetKey(s)
	return ddu
}

// SetNillableKey sets the "key" field if the given value is not nil.
func (ddu *DictionaryDetailUpdate) SetNillableKey(s *string) *DictionaryDetailUpdate {
	if s != nil {
		ddu.SetKey(*s)
	}
	return ddu
}

// SetValue sets the "value" field.
func (ddu *DictionaryDetailUpdate) SetValue(s string) *DictionaryDetailUpdate {
	ddu.mutation.SetValue(s)
	return ddu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (ddu *DictionaryDetailUpdate) SetNillableValue(s *string) *DictionaryDetailUpdate {
	if s != nil {
		ddu.SetValue(*s)
	}
	return ddu
}

// SetDictionaryID sets the "dictionary_id" field.
func (ddu *DictionaryDetailUpdate) SetDictionaryID(i int64) *DictionaryDetailUpdate {
	ddu.mutation.SetDictionaryID(i)
	return ddu
}

// SetNillableDictionaryID sets the "dictionary_id" field if the given value is not nil.
func (ddu *DictionaryDetailUpdate) SetNillableDictionaryID(i *int64) *DictionaryDetailUpdate {
	if i != nil {
		ddu.SetDictionaryID(*i)
	}
	return ddu
}

// ClearDictionaryID clears the value of the "dictionary_id" field.
func (ddu *DictionaryDetailUpdate) ClearDictionaryID() *DictionaryDetailUpdate {
	ddu.mutation.ClearDictionaryID()
	return ddu
}

// SetDictionary sets the "dictionary" edge to the Dictionary entity.
func (ddu *DictionaryDetailUpdate) SetDictionary(d *Dictionary) *DictionaryDetailUpdate {
	return ddu.SetDictionaryID(d.ID)
}

// Mutation returns the DictionaryDetailMutation object of the builder.
func (ddu *DictionaryDetailUpdate) Mutation() *DictionaryDetailMutation {
	return ddu.mutation
}

// ClearDictionary clears the "dictionary" edge to the Dictionary entity.
func (ddu *DictionaryDetailUpdate) ClearDictionary() *DictionaryDetailUpdate {
	ddu.mutation.ClearDictionary()
	return ddu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ddu *DictionaryDetailUpdate) Save(ctx context.Context) (int, error) {
	ddu.defaults()
	return withHooks(ctx, ddu.sqlSave, ddu.mutation, ddu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ddu *DictionaryDetailUpdate) SaveX(ctx context.Context) int {
	affected, err := ddu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ddu *DictionaryDetailUpdate) Exec(ctx context.Context) error {
	_, err := ddu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddu *DictionaryDetailUpdate) ExecX(ctx context.Context) {
	if err := ddu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ddu *DictionaryDetailUpdate) defaults() {
	if _, ok := ddu.mutation.UpdatedAt(); !ok && !ddu.mutation.UpdatedAtCleared() {
		v := dictionarydetail.UpdateDefaultUpdatedAt()
		ddu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ddu *DictionaryDetailUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DictionaryDetailUpdate {
	ddu.modifiers = append(ddu.modifiers, modifiers...)
	return ddu
}

func (ddu *DictionaryDetailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(dictionarydetail.Table, dictionarydetail.Columns, sqlgraph.NewFieldSpec(dictionarydetail.FieldID, field.TypeInt64))
	if ps := ddu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ddu.mutation.CreatedAtCleared() {
		_spec.ClearField(dictionarydetail.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := ddu.mutation.UpdatedAt(); ok {
		_spec.SetField(dictionarydetail.FieldUpdatedAt, field.TypeTime, value)
	}
	if ddu.mutation.UpdatedAtCleared() {
		_spec.ClearField(dictionarydetail.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := ddu.mutation.Delete(); ok {
		_spec.SetField(dictionarydetail.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := ddu.mutation.AddedDelete(); ok {
		_spec.AddField(dictionarydetail.FieldDelete, field.TypeInt64, value)
	}
	if ddu.mutation.DeleteCleared() {
		_spec.ClearField(dictionarydetail.FieldDelete, field.TypeInt64)
	}
	if value, ok := ddu.mutation.CreatedID(); ok {
		_spec.SetField(dictionarydetail.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := ddu.mutation.AddedCreatedID(); ok {
		_spec.AddField(dictionarydetail.FieldCreatedID, field.TypeInt64, value)
	}
	if ddu.mutation.CreatedIDCleared() {
		_spec.ClearField(dictionarydetail.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := ddu.mutation.Status(); ok {
		_spec.SetField(dictionarydetail.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := ddu.mutation.AddedStatus(); ok {
		_spec.AddField(dictionarydetail.FieldStatus, field.TypeInt64, value)
	}
	if ddu.mutation.StatusCleared() {
		_spec.ClearField(dictionarydetail.FieldStatus, field.TypeInt64)
	}
	if value, ok := ddu.mutation.Title(); ok {
		_spec.SetField(dictionarydetail.FieldTitle, field.TypeString, value)
	}
	if value, ok := ddu.mutation.Key(); ok {
		_spec.SetField(dictionarydetail.FieldKey, field.TypeString, value)
	}
	if value, ok := ddu.mutation.Value(); ok {
		_spec.SetField(dictionarydetail.FieldValue, field.TypeString, value)
	}
	if ddu.mutation.DictionaryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dictionarydetail.DictionaryTable,
			Columns: []string{dictionarydetail.DictionaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionary.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ddu.mutation.DictionaryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dictionarydetail.DictionaryTable,
			Columns: []string{dictionarydetail.DictionaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionary.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ddu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ddu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dictionarydetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ddu.mutation.done = true
	return n, nil
}

// DictionaryDetailUpdateOne is the builder for updating a single DictionaryDetail entity.
type DictionaryDetailUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *DictionaryDetailMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (dduo *DictionaryDetailUpdateOne) SetUpdatedAt(t time.Time) *DictionaryDetailUpdateOne {
	dduo.mutation.SetUpdatedAt(t)
	return dduo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (dduo *DictionaryDetailUpdateOne) ClearUpdatedAt() *DictionaryDetailUpdateOne {
	dduo.mutation.ClearUpdatedAt()
	return dduo
}

// SetDelete sets the "delete" field.
func (dduo *DictionaryDetailUpdateOne) SetDelete(i int64) *DictionaryDetailUpdateOne {
	dduo.mutation.ResetDelete()
	dduo.mutation.SetDelete(i)
	return dduo
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (dduo *DictionaryDetailUpdateOne) SetNillableDelete(i *int64) *DictionaryDetailUpdateOne {
	if i != nil {
		dduo.SetDelete(*i)
	}
	return dduo
}

// AddDelete adds i to the "delete" field.
func (dduo *DictionaryDetailUpdateOne) AddDelete(i int64) *DictionaryDetailUpdateOne {
	dduo.mutation.AddDelete(i)
	return dduo
}

// ClearDelete clears the value of the "delete" field.
func (dduo *DictionaryDetailUpdateOne) ClearDelete() *DictionaryDetailUpdateOne {
	dduo.mutation.ClearDelete()
	return dduo
}

// SetCreatedID sets the "created_id" field.
func (dduo *DictionaryDetailUpdateOne) SetCreatedID(i int64) *DictionaryDetailUpdateOne {
	dduo.mutation.ResetCreatedID()
	dduo.mutation.SetCreatedID(i)
	return dduo
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (dduo *DictionaryDetailUpdateOne) SetNillableCreatedID(i *int64) *DictionaryDetailUpdateOne {
	if i != nil {
		dduo.SetCreatedID(*i)
	}
	return dduo
}

// AddCreatedID adds i to the "created_id" field.
func (dduo *DictionaryDetailUpdateOne) AddCreatedID(i int64) *DictionaryDetailUpdateOne {
	dduo.mutation.AddCreatedID(i)
	return dduo
}

// ClearCreatedID clears the value of the "created_id" field.
func (dduo *DictionaryDetailUpdateOne) ClearCreatedID() *DictionaryDetailUpdateOne {
	dduo.mutation.ClearCreatedID()
	return dduo
}

// SetStatus sets the "status" field.
func (dduo *DictionaryDetailUpdateOne) SetStatus(i int64) *DictionaryDetailUpdateOne {
	dduo.mutation.ResetStatus()
	dduo.mutation.SetStatus(i)
	return dduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dduo *DictionaryDetailUpdateOne) SetNillableStatus(i *int64) *DictionaryDetailUpdateOne {
	if i != nil {
		dduo.SetStatus(*i)
	}
	return dduo
}

// AddStatus adds i to the "status" field.
func (dduo *DictionaryDetailUpdateOne) AddStatus(i int64) *DictionaryDetailUpdateOne {
	dduo.mutation.AddStatus(i)
	return dduo
}

// ClearStatus clears the value of the "status" field.
func (dduo *DictionaryDetailUpdateOne) ClearStatus() *DictionaryDetailUpdateOne {
	dduo.mutation.ClearStatus()
	return dduo
}

// SetTitle sets the "title" field.
func (dduo *DictionaryDetailUpdateOne) SetTitle(s string) *DictionaryDetailUpdateOne {
	dduo.mutation.SetTitle(s)
	return dduo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (dduo *DictionaryDetailUpdateOne) SetNillableTitle(s *string) *DictionaryDetailUpdateOne {
	if s != nil {
		dduo.SetTitle(*s)
	}
	return dduo
}

// SetKey sets the "key" field.
func (dduo *DictionaryDetailUpdateOne) SetKey(s string) *DictionaryDetailUpdateOne {
	dduo.mutation.SetKey(s)
	return dduo
}

// SetNillableKey sets the "key" field if the given value is not nil.
func (dduo *DictionaryDetailUpdateOne) SetNillableKey(s *string) *DictionaryDetailUpdateOne {
	if s != nil {
		dduo.SetKey(*s)
	}
	return dduo
}

// SetValue sets the "value" field.
func (dduo *DictionaryDetailUpdateOne) SetValue(s string) *DictionaryDetailUpdateOne {
	dduo.mutation.SetValue(s)
	return dduo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (dduo *DictionaryDetailUpdateOne) SetNillableValue(s *string) *DictionaryDetailUpdateOne {
	if s != nil {
		dduo.SetValue(*s)
	}
	return dduo
}

// SetDictionaryID sets the "dictionary_id" field.
func (dduo *DictionaryDetailUpdateOne) SetDictionaryID(i int64) *DictionaryDetailUpdateOne {
	dduo.mutation.SetDictionaryID(i)
	return dduo
}

// SetNillableDictionaryID sets the "dictionary_id" field if the given value is not nil.
func (dduo *DictionaryDetailUpdateOne) SetNillableDictionaryID(i *int64) *DictionaryDetailUpdateOne {
	if i != nil {
		dduo.SetDictionaryID(*i)
	}
	return dduo
}

// ClearDictionaryID clears the value of the "dictionary_id" field.
func (dduo *DictionaryDetailUpdateOne) ClearDictionaryID() *DictionaryDetailUpdateOne {
	dduo.mutation.ClearDictionaryID()
	return dduo
}

// SetDictionary sets the "dictionary" edge to the Dictionary entity.
func (dduo *DictionaryDetailUpdateOne) SetDictionary(d *Dictionary) *DictionaryDetailUpdateOne {
	return dduo.SetDictionaryID(d.ID)
}

// Mutation returns the DictionaryDetailMutation object of the builder.
func (dduo *DictionaryDetailUpdateOne) Mutation() *DictionaryDetailMutation {
	return dduo.mutation
}

// ClearDictionary clears the "dictionary" edge to the Dictionary entity.
func (dduo *DictionaryDetailUpdateOne) ClearDictionary() *DictionaryDetailUpdateOne {
	dduo.mutation.ClearDictionary()
	return dduo
}

// Where appends a list predicates to the DictionaryDetailUpdate builder.
func (dduo *DictionaryDetailUpdateOne) Where(ps ...predicate.DictionaryDetail) *DictionaryDetailUpdateOne {
	dduo.mutation.Where(ps...)
	return dduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dduo *DictionaryDetailUpdateOne) Select(field string, fields ...string) *DictionaryDetailUpdateOne {
	dduo.fields = append([]string{field}, fields...)
	return dduo
}

// Save executes the query and returns the updated DictionaryDetail entity.
func (dduo *DictionaryDetailUpdateOne) Save(ctx context.Context) (*DictionaryDetail, error) {
	dduo.defaults()
	return withHooks(ctx, dduo.sqlSave, dduo.mutation, dduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dduo *DictionaryDetailUpdateOne) SaveX(ctx context.Context) *DictionaryDetail {
	node, err := dduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dduo *DictionaryDetailUpdateOne) Exec(ctx context.Context) error {
	_, err := dduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dduo *DictionaryDetailUpdateOne) ExecX(ctx context.Context) {
	if err := dduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dduo *DictionaryDetailUpdateOne) defaults() {
	if _, ok := dduo.mutation.UpdatedAt(); !ok && !dduo.mutation.UpdatedAtCleared() {
		v := dictionarydetail.UpdateDefaultUpdatedAt()
		dduo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (dduo *DictionaryDetailUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DictionaryDetailUpdateOne {
	dduo.modifiers = append(dduo.modifiers, modifiers...)
	return dduo
}

func (dduo *DictionaryDetailUpdateOne) sqlSave(ctx context.Context) (_node *DictionaryDetail, err error) {
	_spec := sqlgraph.NewUpdateSpec(dictionarydetail.Table, dictionarydetail.Columns, sqlgraph.NewFieldSpec(dictionarydetail.FieldID, field.TypeInt64))
	id, ok := dduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DictionaryDetail.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dictionarydetail.FieldID)
		for _, f := range fields {
			if !dictionarydetail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dictionarydetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if dduo.mutation.CreatedAtCleared() {
		_spec.ClearField(dictionarydetail.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := dduo.mutation.UpdatedAt(); ok {
		_spec.SetField(dictionarydetail.FieldUpdatedAt, field.TypeTime, value)
	}
	if dduo.mutation.UpdatedAtCleared() {
		_spec.ClearField(dictionarydetail.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := dduo.mutation.Delete(); ok {
		_spec.SetField(dictionarydetail.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := dduo.mutation.AddedDelete(); ok {
		_spec.AddField(dictionarydetail.FieldDelete, field.TypeInt64, value)
	}
	if dduo.mutation.DeleteCleared() {
		_spec.ClearField(dictionarydetail.FieldDelete, field.TypeInt64)
	}
	if value, ok := dduo.mutation.CreatedID(); ok {
		_spec.SetField(dictionarydetail.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := dduo.mutation.AddedCreatedID(); ok {
		_spec.AddField(dictionarydetail.FieldCreatedID, field.TypeInt64, value)
	}
	if dduo.mutation.CreatedIDCleared() {
		_spec.ClearField(dictionarydetail.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := dduo.mutation.Status(); ok {
		_spec.SetField(dictionarydetail.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := dduo.mutation.AddedStatus(); ok {
		_spec.AddField(dictionarydetail.FieldStatus, field.TypeInt64, value)
	}
	if dduo.mutation.StatusCleared() {
		_spec.ClearField(dictionarydetail.FieldStatus, field.TypeInt64)
	}
	if value, ok := dduo.mutation.Title(); ok {
		_spec.SetField(dictionarydetail.FieldTitle, field.TypeString, value)
	}
	if value, ok := dduo.mutation.Key(); ok {
		_spec.SetField(dictionarydetail.FieldKey, field.TypeString, value)
	}
	if value, ok := dduo.mutation.Value(); ok {
		_spec.SetField(dictionarydetail.FieldValue, field.TypeString, value)
	}
	if dduo.mutation.DictionaryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dictionarydetail.DictionaryTable,
			Columns: []string{dictionarydetail.DictionaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionary.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dduo.mutation.DictionaryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dictionarydetail.DictionaryTable,
			Columns: []string{dictionarydetail.DictionaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionary.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(dduo.modifiers...)
	_node = &DictionaryDetail{config: dduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dictionarydetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dduo.mutation.done = true
	return _node, nil
}
