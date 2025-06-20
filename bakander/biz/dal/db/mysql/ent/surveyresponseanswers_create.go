// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kcers-survey/biz/dal/db/mysql/ent/surveyresponseanswers"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SurveyResponseAnswersCreate is the builder for creating a SurveyResponseAnswers entity.
type SurveyResponseAnswersCreate struct {
	config
	mutation *SurveyResponseAnswersMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (srac *SurveyResponseAnswersCreate) SetCreatedAt(t time.Time) *SurveyResponseAnswersCreate {
	srac.mutation.SetCreatedAt(t)
	return srac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (srac *SurveyResponseAnswersCreate) SetNillableCreatedAt(t *time.Time) *SurveyResponseAnswersCreate {
	if t != nil {
		srac.SetCreatedAt(*t)
	}
	return srac
}

// SetUpdatedAt sets the "updated_at" field.
func (srac *SurveyResponseAnswersCreate) SetUpdatedAt(t time.Time) *SurveyResponseAnswersCreate {
	srac.mutation.SetUpdatedAt(t)
	return srac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (srac *SurveyResponseAnswersCreate) SetNillableUpdatedAt(t *time.Time) *SurveyResponseAnswersCreate {
	if t != nil {
		srac.SetUpdatedAt(*t)
	}
	return srac
}

// SetDelete sets the "delete" field.
func (srac *SurveyResponseAnswersCreate) SetDelete(i int64) *SurveyResponseAnswersCreate {
	srac.mutation.SetDelete(i)
	return srac
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (srac *SurveyResponseAnswersCreate) SetNillableDelete(i *int64) *SurveyResponseAnswersCreate {
	if i != nil {
		srac.SetDelete(*i)
	}
	return srac
}

// SetCreatedID sets the "created_id" field.
func (srac *SurveyResponseAnswersCreate) SetCreatedID(i int64) *SurveyResponseAnswersCreate {
	srac.mutation.SetCreatedID(i)
	return srac
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (srac *SurveyResponseAnswersCreate) SetNillableCreatedID(i *int64) *SurveyResponseAnswersCreate {
	if i != nil {
		srac.SetCreatedID(*i)
	}
	return srac
}

// SetStatus sets the "status" field.
func (srac *SurveyResponseAnswersCreate) SetStatus(i int64) *SurveyResponseAnswersCreate {
	srac.mutation.SetStatus(i)
	return srac
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (srac *SurveyResponseAnswersCreate) SetNillableStatus(i *int64) *SurveyResponseAnswersCreate {
	if i != nil {
		srac.SetStatus(*i)
	}
	return srac
}

// SetSurveyID sets the "survey_id" field.
func (srac *SurveyResponseAnswersCreate) SetSurveyID(i int64) *SurveyResponseAnswersCreate {
	srac.mutation.SetSurveyID(i)
	return srac
}

// SetNillableSurveyID sets the "survey_id" field if the given value is not nil.
func (srac *SurveyResponseAnswersCreate) SetNillableSurveyID(i *int64) *SurveyResponseAnswersCreate {
	if i != nil {
		srac.SetSurveyID(*i)
	}
	return srac
}

// SetSurveyResponseID sets the "survey_response_id" field.
func (srac *SurveyResponseAnswersCreate) SetSurveyResponseID(i int64) *SurveyResponseAnswersCreate {
	srac.mutation.SetSurveyResponseID(i)
	return srac
}

// SetNillableSurveyResponseID sets the "survey_response_id" field if the given value is not nil.
func (srac *SurveyResponseAnswersCreate) SetNillableSurveyResponseID(i *int64) *SurveyResponseAnswersCreate {
	if i != nil {
		srac.SetSurveyResponseID(*i)
	}
	return srac
}

// SetSurveyQuestionID sets the "survey_question_id" field.
func (srac *SurveyResponseAnswersCreate) SetSurveyQuestionID(i int64) *SurveyResponseAnswersCreate {
	srac.mutation.SetSurveyQuestionID(i)
	return srac
}

// SetNillableSurveyQuestionID sets the "survey_question_id" field if the given value is not nil.
func (srac *SurveyResponseAnswersCreate) SetNillableSurveyQuestionID(i *int64) *SurveyResponseAnswersCreate {
	if i != nil {
		srac.SetSurveyQuestionID(*i)
	}
	return srac
}

// SetAnswerText sets the "answer_text" field.
func (srac *SurveyResponseAnswersCreate) SetAnswerText(s string) *SurveyResponseAnswersCreate {
	srac.mutation.SetAnswerText(s)
	return srac
}

// SetAnswerValue sets the "answer_value" field.
func (srac *SurveyResponseAnswersCreate) SetAnswerValue(i int64) *SurveyResponseAnswersCreate {
	srac.mutation.SetAnswerValue(i)
	return srac
}

// SetNillableAnswerValue sets the "answer_value" field if the given value is not nil.
func (srac *SurveyResponseAnswersCreate) SetNillableAnswerValue(i *int64) *SurveyResponseAnswersCreate {
	if i != nil {
		srac.SetAnswerValue(*i)
	}
	return srac
}

// SetID sets the "id" field.
func (srac *SurveyResponseAnswersCreate) SetID(i int64) *SurveyResponseAnswersCreate {
	srac.mutation.SetID(i)
	return srac
}

// Mutation returns the SurveyResponseAnswersMutation object of the builder.
func (srac *SurveyResponseAnswersCreate) Mutation() *SurveyResponseAnswersMutation {
	return srac.mutation
}

// Save creates the SurveyResponseAnswers in the database.
func (srac *SurveyResponseAnswersCreate) Save(ctx context.Context) (*SurveyResponseAnswers, error) {
	srac.defaults()
	return withHooks(ctx, srac.sqlSave, srac.mutation, srac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (srac *SurveyResponseAnswersCreate) SaveX(ctx context.Context) *SurveyResponseAnswers {
	v, err := srac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (srac *SurveyResponseAnswersCreate) Exec(ctx context.Context) error {
	_, err := srac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (srac *SurveyResponseAnswersCreate) ExecX(ctx context.Context) {
	if err := srac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (srac *SurveyResponseAnswersCreate) defaults() {
	if _, ok := srac.mutation.CreatedAt(); !ok {
		v := surveyresponseanswers.DefaultCreatedAt()
		srac.mutation.SetCreatedAt(v)
	}
	if _, ok := srac.mutation.UpdatedAt(); !ok {
		v := surveyresponseanswers.DefaultUpdatedAt()
		srac.mutation.SetUpdatedAt(v)
	}
	if _, ok := srac.mutation.Delete(); !ok {
		v := surveyresponseanswers.DefaultDelete
		srac.mutation.SetDelete(v)
	}
	if _, ok := srac.mutation.CreatedID(); !ok {
		v := surveyresponseanswers.DefaultCreatedID
		srac.mutation.SetCreatedID(v)
	}
	if _, ok := srac.mutation.Status(); !ok {
		v := surveyresponseanswers.DefaultStatus
		srac.mutation.SetStatus(v)
	}
	if _, ok := srac.mutation.SurveyID(); !ok {
		v := surveyresponseanswers.DefaultSurveyID
		srac.mutation.SetSurveyID(v)
	}
	if _, ok := srac.mutation.SurveyResponseID(); !ok {
		v := surveyresponseanswers.DefaultSurveyResponseID
		srac.mutation.SetSurveyResponseID(v)
	}
	if _, ok := srac.mutation.SurveyQuestionID(); !ok {
		v := surveyresponseanswers.DefaultSurveyQuestionID
		srac.mutation.SetSurveyQuestionID(v)
	}
	if _, ok := srac.mutation.AnswerValue(); !ok {
		v := surveyresponseanswers.DefaultAnswerValue
		srac.mutation.SetAnswerValue(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (srac *SurveyResponseAnswersCreate) check() error {
	if _, ok := srac.mutation.SurveyID(); !ok {
		return &ValidationError{Name: "survey_id", err: errors.New(`ent: missing required field "SurveyResponseAnswers.survey_id"`)}
	}
	if _, ok := srac.mutation.SurveyResponseID(); !ok {
		return &ValidationError{Name: "survey_response_id", err: errors.New(`ent: missing required field "SurveyResponseAnswers.survey_response_id"`)}
	}
	if _, ok := srac.mutation.SurveyQuestionID(); !ok {
		return &ValidationError{Name: "survey_question_id", err: errors.New(`ent: missing required field "SurveyResponseAnswers.survey_question_id"`)}
	}
	if _, ok := srac.mutation.AnswerText(); !ok {
		return &ValidationError{Name: "answer_text", err: errors.New(`ent: missing required field "SurveyResponseAnswers.answer_text"`)}
	}
	if _, ok := srac.mutation.AnswerValue(); !ok {
		return &ValidationError{Name: "answer_value", err: errors.New(`ent: missing required field "SurveyResponseAnswers.answer_value"`)}
	}
	return nil
}

func (srac *SurveyResponseAnswersCreate) sqlSave(ctx context.Context) (*SurveyResponseAnswers, error) {
	if err := srac.check(); err != nil {
		return nil, err
	}
	_node, _spec := srac.createSpec()
	if err := sqlgraph.CreateNode(ctx, srac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	srac.mutation.id = &_node.ID
	srac.mutation.done = true
	return _node, nil
}

func (srac *SurveyResponseAnswersCreate) createSpec() (*SurveyResponseAnswers, *sqlgraph.CreateSpec) {
	var (
		_node = &SurveyResponseAnswers{config: srac.config}
		_spec = sqlgraph.NewCreateSpec(surveyresponseanswers.Table, sqlgraph.NewFieldSpec(surveyresponseanswers.FieldID, field.TypeInt64))
	)
	if id, ok := srac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := srac.mutation.CreatedAt(); ok {
		_spec.SetField(surveyresponseanswers.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := srac.mutation.UpdatedAt(); ok {
		_spec.SetField(surveyresponseanswers.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := srac.mutation.Delete(); ok {
		_spec.SetField(surveyresponseanswers.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
	}
	if value, ok := srac.mutation.CreatedID(); ok {
		_spec.SetField(surveyresponseanswers.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := srac.mutation.Status(); ok {
		_spec.SetField(surveyresponseanswers.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := srac.mutation.SurveyID(); ok {
		_spec.SetField(surveyresponseanswers.FieldSurveyID, field.TypeInt64, value)
		_node.SurveyID = value
	}
	if value, ok := srac.mutation.SurveyResponseID(); ok {
		_spec.SetField(surveyresponseanswers.FieldSurveyResponseID, field.TypeInt64, value)
		_node.SurveyResponseID = value
	}
	if value, ok := srac.mutation.SurveyQuestionID(); ok {
		_spec.SetField(surveyresponseanswers.FieldSurveyQuestionID, field.TypeInt64, value)
		_node.SurveyQuestionID = value
	}
	if value, ok := srac.mutation.AnswerText(); ok {
		_spec.SetField(surveyresponseanswers.FieldAnswerText, field.TypeString, value)
		_node.AnswerText = value
	}
	if value, ok := srac.mutation.AnswerValue(); ok {
		_spec.SetField(surveyresponseanswers.FieldAnswerValue, field.TypeInt64, value)
		_node.AnswerValue = value
	}
	return _node, _spec
}

// SurveyResponseAnswersCreateBulk is the builder for creating many SurveyResponseAnswers entities in bulk.
type SurveyResponseAnswersCreateBulk struct {
	config
	err      error
	builders []*SurveyResponseAnswersCreate
}

// Save creates the SurveyResponseAnswers entities in the database.
func (sracb *SurveyResponseAnswersCreateBulk) Save(ctx context.Context) ([]*SurveyResponseAnswers, error) {
	if sracb.err != nil {
		return nil, sracb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sracb.builders))
	nodes := make([]*SurveyResponseAnswers, len(sracb.builders))
	mutators := make([]Mutator, len(sracb.builders))
	for i := range sracb.builders {
		func(i int, root context.Context) {
			builder := sracb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SurveyResponseAnswersMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sracb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sracb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sracb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sracb *SurveyResponseAnswersCreateBulk) SaveX(ctx context.Context) []*SurveyResponseAnswers {
	v, err := sracb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sracb *SurveyResponseAnswersCreateBulk) Exec(ctx context.Context) error {
	_, err := sracb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sracb *SurveyResponseAnswersCreateBulk) ExecX(ctx context.Context) {
	if err := sracb.Exec(ctx); err != nil {
		panic(err)
	}
}
