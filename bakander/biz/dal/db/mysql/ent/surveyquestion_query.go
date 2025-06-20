// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"kcers-survey/biz/dal/db/mysql/ent/predicate"
	"kcers-survey/biz/dal/db/mysql/ent/surveyquestion"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SurveyQuestionQuery is the builder for querying SurveyQuestion entities.
type SurveyQuestionQuery struct {
	config
	ctx        *QueryContext
	order      []surveyquestion.OrderOption
	inters     []Interceptor
	predicates []predicate.SurveyQuestion
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SurveyQuestionQuery builder.
func (sqq *SurveyQuestionQuery) Where(ps ...predicate.SurveyQuestion) *SurveyQuestionQuery {
	sqq.predicates = append(sqq.predicates, ps...)
	return sqq
}

// Limit the number of records to be returned by this query.
func (sqq *SurveyQuestionQuery) Limit(limit int) *SurveyQuestionQuery {
	sqq.ctx.Limit = &limit
	return sqq
}

// Offset to start from.
func (sqq *SurveyQuestionQuery) Offset(offset int) *SurveyQuestionQuery {
	sqq.ctx.Offset = &offset
	return sqq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sqq *SurveyQuestionQuery) Unique(unique bool) *SurveyQuestionQuery {
	sqq.ctx.Unique = &unique
	return sqq
}

// Order specifies how the records should be ordered.
func (sqq *SurveyQuestionQuery) Order(o ...surveyquestion.OrderOption) *SurveyQuestionQuery {
	sqq.order = append(sqq.order, o...)
	return sqq
}

// First returns the first SurveyQuestion entity from the query.
// Returns a *NotFoundError when no SurveyQuestion was found.
func (sqq *SurveyQuestionQuery) First(ctx context.Context) (*SurveyQuestion, error) {
	nodes, err := sqq.Limit(1).All(setContextOp(ctx, sqq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{surveyquestion.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sqq *SurveyQuestionQuery) FirstX(ctx context.Context) *SurveyQuestion {
	node, err := sqq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SurveyQuestion ID from the query.
// Returns a *NotFoundError when no SurveyQuestion ID was found.
func (sqq *SurveyQuestionQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sqq.Limit(1).IDs(setContextOp(ctx, sqq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{surveyquestion.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sqq *SurveyQuestionQuery) FirstIDX(ctx context.Context) int64 {
	id, err := sqq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SurveyQuestion entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SurveyQuestion entity is found.
// Returns a *NotFoundError when no SurveyQuestion entities are found.
func (sqq *SurveyQuestionQuery) Only(ctx context.Context) (*SurveyQuestion, error) {
	nodes, err := sqq.Limit(2).All(setContextOp(ctx, sqq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{surveyquestion.Label}
	default:
		return nil, &NotSingularError{surveyquestion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sqq *SurveyQuestionQuery) OnlyX(ctx context.Context) *SurveyQuestion {
	node, err := sqq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SurveyQuestion ID in the query.
// Returns a *NotSingularError when more than one SurveyQuestion ID is found.
// Returns a *NotFoundError when no entities are found.
func (sqq *SurveyQuestionQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sqq.Limit(2).IDs(setContextOp(ctx, sqq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{surveyquestion.Label}
	default:
		err = &NotSingularError{surveyquestion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sqq *SurveyQuestionQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := sqq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SurveyQuestions.
func (sqq *SurveyQuestionQuery) All(ctx context.Context) ([]*SurveyQuestion, error) {
	ctx = setContextOp(ctx, sqq.ctx, ent.OpQueryAll)
	if err := sqq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SurveyQuestion, *SurveyQuestionQuery]()
	return withInterceptors[[]*SurveyQuestion](ctx, sqq, qr, sqq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sqq *SurveyQuestionQuery) AllX(ctx context.Context) []*SurveyQuestion {
	nodes, err := sqq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SurveyQuestion IDs.
func (sqq *SurveyQuestionQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if sqq.ctx.Unique == nil && sqq.path != nil {
		sqq.Unique(true)
	}
	ctx = setContextOp(ctx, sqq.ctx, ent.OpQueryIDs)
	if err = sqq.Select(surveyquestion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sqq *SurveyQuestionQuery) IDsX(ctx context.Context) []int64 {
	ids, err := sqq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sqq *SurveyQuestionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sqq.ctx, ent.OpQueryCount)
	if err := sqq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sqq, querierCount[*SurveyQuestionQuery](), sqq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sqq *SurveyQuestionQuery) CountX(ctx context.Context) int {
	count, err := sqq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sqq *SurveyQuestionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sqq.ctx, ent.OpQueryExist)
	switch _, err := sqq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sqq *SurveyQuestionQuery) ExistX(ctx context.Context) bool {
	exist, err := sqq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SurveyQuestionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sqq *SurveyQuestionQuery) Clone() *SurveyQuestionQuery {
	if sqq == nil {
		return nil
	}
	return &SurveyQuestionQuery{
		config:     sqq.config,
		ctx:        sqq.ctx.Clone(),
		order:      append([]surveyquestion.OrderOption{}, sqq.order...),
		inters:     append([]Interceptor{}, sqq.inters...),
		predicates: append([]predicate.SurveyQuestion{}, sqq.predicates...),
		// clone intermediate query.
		sql:       sqq.sql.Clone(),
		path:      sqq.path,
		modifiers: append([]func(*sql.Selector){}, sqq.modifiers...),
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SurveyQuestion.Query().
//		GroupBy(surveyquestion.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sqq *SurveyQuestionQuery) GroupBy(field string, fields ...string) *SurveyQuestionGroupBy {
	sqq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SurveyQuestionGroupBy{build: sqq}
	grbuild.flds = &sqq.ctx.Fields
	grbuild.label = surveyquestion.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.SurveyQuestion.Query().
//		Select(surveyquestion.FieldCreatedAt).
//		Scan(ctx, &v)
func (sqq *SurveyQuestionQuery) Select(fields ...string) *SurveyQuestionSelect {
	sqq.ctx.Fields = append(sqq.ctx.Fields, fields...)
	sbuild := &SurveyQuestionSelect{SurveyQuestionQuery: sqq}
	sbuild.label = surveyquestion.Label
	sbuild.flds, sbuild.scan = &sqq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SurveyQuestionSelect configured with the given aggregations.
func (sqq *SurveyQuestionQuery) Aggregate(fns ...AggregateFunc) *SurveyQuestionSelect {
	return sqq.Select().Aggregate(fns...)
}

func (sqq *SurveyQuestionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sqq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sqq); err != nil {
				return err
			}
		}
	}
	for _, f := range sqq.ctx.Fields {
		if !surveyquestion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sqq.path != nil {
		prev, err := sqq.path(ctx)
		if err != nil {
			return err
		}
		sqq.sql = prev
	}
	return nil
}

func (sqq *SurveyQuestionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SurveyQuestion, error) {
	var (
		nodes = []*SurveyQuestion{}
		_spec = sqq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SurveyQuestion).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SurveyQuestion{config: sqq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(sqq.modifiers) > 0 {
		_spec.Modifiers = sqq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sqq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sqq *SurveyQuestionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sqq.querySpec()
	if len(sqq.modifiers) > 0 {
		_spec.Modifiers = sqq.modifiers
	}
	_spec.Node.Columns = sqq.ctx.Fields
	if len(sqq.ctx.Fields) > 0 {
		_spec.Unique = sqq.ctx.Unique != nil && *sqq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sqq.driver, _spec)
}

func (sqq *SurveyQuestionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(surveyquestion.Table, surveyquestion.Columns, sqlgraph.NewFieldSpec(surveyquestion.FieldID, field.TypeInt64))
	_spec.From = sqq.sql
	if unique := sqq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sqq.path != nil {
		_spec.Unique = true
	}
	if fields := sqq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, surveyquestion.FieldID)
		for i := range fields {
			if fields[i] != surveyquestion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sqq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sqq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sqq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sqq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sqq *SurveyQuestionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sqq.driver.Dialect())
	t1 := builder.Table(surveyquestion.Table)
	columns := sqq.ctx.Fields
	if len(columns) == 0 {
		columns = surveyquestion.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sqq.sql != nil {
		selector = sqq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sqq.ctx.Unique != nil && *sqq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range sqq.modifiers {
		m(selector)
	}
	for _, p := range sqq.predicates {
		p(selector)
	}
	for _, p := range sqq.order {
		p(selector)
	}
	if offset := sqq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sqq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sqq *SurveyQuestionQuery) Modify(modifiers ...func(s *sql.Selector)) *SurveyQuestionSelect {
	sqq.modifiers = append(sqq.modifiers, modifiers...)
	return sqq.Select()
}

// SurveyQuestionGroupBy is the group-by builder for SurveyQuestion entities.
type SurveyQuestionGroupBy struct {
	selector
	build *SurveyQuestionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sqgb *SurveyQuestionGroupBy) Aggregate(fns ...AggregateFunc) *SurveyQuestionGroupBy {
	sqgb.fns = append(sqgb.fns, fns...)
	return sqgb
}

// Scan applies the selector query and scans the result into the given value.
func (sqgb *SurveyQuestionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sqgb.build.ctx, ent.OpQueryGroupBy)
	if err := sqgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SurveyQuestionQuery, *SurveyQuestionGroupBy](ctx, sqgb.build, sqgb, sqgb.build.inters, v)
}

func (sqgb *SurveyQuestionGroupBy) sqlScan(ctx context.Context, root *SurveyQuestionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sqgb.fns))
	for _, fn := range sqgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sqgb.flds)+len(sqgb.fns))
		for _, f := range *sqgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sqgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sqgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SurveyQuestionSelect is the builder for selecting fields of SurveyQuestion entities.
type SurveyQuestionSelect struct {
	*SurveyQuestionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sqs *SurveyQuestionSelect) Aggregate(fns ...AggregateFunc) *SurveyQuestionSelect {
	sqs.fns = append(sqs.fns, fns...)
	return sqs
}

// Scan applies the selector query and scans the result into the given value.
func (sqs *SurveyQuestionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sqs.ctx, ent.OpQuerySelect)
	if err := sqs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SurveyQuestionQuery, *SurveyQuestionSelect](ctx, sqs.SurveyQuestionQuery, sqs, sqs.inters, v)
}

func (sqs *SurveyQuestionSelect) sqlScan(ctx context.Context, root *SurveyQuestionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sqs.fns))
	for _, fn := range sqs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sqs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sqs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sqs *SurveyQuestionSelect) Modify(modifiers ...func(s *sql.Selector)) *SurveyQuestionSelect {
	sqs.modifiers = append(sqs.modifiers, modifiers...)
	return sqs
}
