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

type SurveyQuestion struct {
	ent.Schema
}

func (SurveyQuestion) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("title"),
		field.Int64("survey_id").Default(0).Comment("survey_id"),
		field.String("type").Comment("type"),
		field.Int64("required").Default(1).Comment("required"),

		field.Int64("pid").Default(0).Comment("pid"),
	}
}

func (SurveyQuestion) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (SurveyQuestion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("survey", Survey.Type).Ref("question").Field("survey_id"),
		edge.To("response", SurveyQuestionResponse.Type),
	}
}

func (SurveyQuestion) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "survey_question"},
		entsql.WithComments(true),
	}
}
