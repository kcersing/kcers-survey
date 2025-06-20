// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kcers-survey/biz/dal/db/mysql/ent/predicate"
	"kcers-survey/biz/dal/db/mysql/ent/surveyresponseanswers"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SurveyResponseAnswersDelete is the builder for deleting a SurveyResponseAnswers entity.
type SurveyResponseAnswersDelete struct {
	config
	hooks    []Hook
	mutation *SurveyResponseAnswersMutation
}

// Where appends a list predicates to the SurveyResponseAnswersDelete builder.
func (srad *SurveyResponseAnswersDelete) Where(ps ...predicate.SurveyResponseAnswers) *SurveyResponseAnswersDelete {
	srad.mutation.Where(ps...)
	return srad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (srad *SurveyResponseAnswersDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, srad.sqlExec, srad.mutation, srad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (srad *SurveyResponseAnswersDelete) ExecX(ctx context.Context) int {
	n, err := srad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (srad *SurveyResponseAnswersDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(surveyresponseanswers.Table, sqlgraph.NewFieldSpec(surveyresponseanswers.FieldID, field.TypeInt64))
	if ps := srad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, srad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	srad.mutation.done = true
	return affected, err
}

// SurveyResponseAnswersDeleteOne is the builder for deleting a single SurveyResponseAnswers entity.
type SurveyResponseAnswersDeleteOne struct {
	srad *SurveyResponseAnswersDelete
}

// Where appends a list predicates to the SurveyResponseAnswersDelete builder.
func (srado *SurveyResponseAnswersDeleteOne) Where(ps ...predicate.SurveyResponseAnswers) *SurveyResponseAnswersDeleteOne {
	srado.srad.mutation.Where(ps...)
	return srado
}

// Exec executes the deletion query.
func (srado *SurveyResponseAnswersDeleteOne) Exec(ctx context.Context) error {
	n, err := srado.srad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{surveyresponseanswers.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (srado *SurveyResponseAnswersDeleteOne) ExecX(ctx context.Context) {
	if err := srado.Exec(ctx); err != nil {
		panic(err)
	}
}
