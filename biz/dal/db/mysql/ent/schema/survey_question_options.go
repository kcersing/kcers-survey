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

type SurveyQuestionOptions struct {
	ent.Schema
}

func (SurveyQuestionOptions) Fields() []ent.Field {
	return []ent.Field{

		field.Int64("survey_question_id").Optional().Default(0).Comment("survey_question_id"),
		field.String("serial").Optional().Default("").Comment("serial"),
		field.Text("content").Optional().Default("").Comment("content"),
	}
}

func (SurveyQuestionOptions) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (SurveyQuestionOptions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question", SurveyQuestion.Type).Ref("option").Field("survey_question_id").Unique(),
	}
}

func (SurveyQuestionOptions) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "survey_question_options"},
		entsql.WithComments(true),
	}
}
