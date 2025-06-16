package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	_ "entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"kcers-survey/biz/dal/db/mysql/ent/schema/mixins"
)

// SurveyResponseAnswers 回答
type SurveyResponseAnswers struct {
	ent.Schema
}

func (SurveyResponseAnswers) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("survey_id").Default(0).Comment("survey_id"),
		field.Int64("survey_response_id").Default(0).Comment("survey_response_id"),
		field.Int64("survey_question_id").Default(0).Comment("survey_question_id"),
		field.String("answer_text").Comment("回答文本"),
		field.Int64("answer_value").Default(1).Comment("回答数值"),
	}
}

func (SurveyResponseAnswers) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (SurveyResponseAnswers) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (SurveyResponseAnswers) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "survey_response_answers"},
		entsql.WithComments(true),
	}
}
