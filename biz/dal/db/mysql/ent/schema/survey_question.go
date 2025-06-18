package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	_ "entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"kcers-survey/biz/dal/db/mysql/ent/schema/mixins"
	"entgo.io/ent/schema/edge"
)

type SurveyQuestion struct {
	ent.Schema
}

func (SurveyQuestion) Fields() []ent.Field {
	return []ent.Field{

		field.Int64("survey_id").Optional().Default(0).Comment("survey_id"),
		field.Int64("parent_id").Optional().Default(0).Comment("parent_id"),
		field.Text("content").Optional().Default("").Comment("content"),
		field.String("type").Optional().Default("").Comment("type"),

		field.Int64("sort").Optional().Default(0).Comment("sort"),

		field.Int64("required").Optional().Default(1).Comment("是否必填 1必填 2选填"),
		field.JSON("options", map[string]string{}).Optional().Comment("存储选项"),
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
		edge.From("survey", Survey.Type).Ref("question").Field("survey_id").Unique(),

	}
}

func (SurveyQuestion) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "survey_question"},
		entsql.WithComments(true),
	}
}
