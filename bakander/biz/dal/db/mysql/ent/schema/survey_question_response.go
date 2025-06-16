package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	_ "entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"kcers-survey/biz/dal/db/mysql/ent/schema/mixins"
)

type SurveyQuestionResponse struct {
	ent.Schema
}

func (SurveyQuestionResponse) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("title"),
		field.Int64("question_id").Default(0).Comment("question_id"),
		field.String("type").Comment("type"),
		field.Int64("required").Default(1).Comment("required"),
	}
}

func (SurveyQuestionResponse) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (SurveyQuestionResponse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question", SurveyQuestion.Type).Ref("response").Field("question_id"),
	}
}

func (SurveyQuestionResponse) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "survey_question_response"},
		entsql.WithComments(true),
	}
}
