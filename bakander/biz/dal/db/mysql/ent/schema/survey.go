package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	_ "entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"kcers-survey/biz/dal/db/mysql/ent/schema/mixins"
)

type Survey struct {
	ent.Schema
}

func (Survey) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("title"),
		field.String("pic").Comment("pic"),
		field.Text("desc").Comment("desc"),
		field.Time("start_at").Default(nil).Comment("开始时间"),
		field.Time("end_at").Default(nil).Comment("结束时间"),
	}
}

func (Survey) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Survey) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("question", SurveyQuestion.Type),
	}
}

func (Survey) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "survey"},
		entsql.WithComments(true),
	}
}
