package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	_ "entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"kcers-survey/biz/dal/db/mysql/ent/schema/mixins"
)

// SurveyResponse 回答
type SurveyResponse struct {
	ent.Schema
}

func (SurveyResponse) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("survey_id").Default(0).Comment("survey_id"),
		field.String("ip").Default("").Comment("用户IP地址"),
		field.String("map").Default("").Comment("用户地图坐标"),
		field.String("device").Default("").Comment("设备信息"),
		field.String("audio").Default("").Comment("音频"),

		field.Time("started_at").Default(nil).Comment("开始时间"),
		field.Time("completed_at").Default(nil).Comment("完成时间"),
	}
}

func (SurveyResponse) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (SurveyResponse) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("survey", Survey.Type).Ref("question").Field("survey_id"),
	}

}

func (SurveyResponse) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "survey_response"},
		entsql.WithComments(true),
	}
}
