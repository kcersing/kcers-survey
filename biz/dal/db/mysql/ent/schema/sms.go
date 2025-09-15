package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"kcers-survey/biz/dal/db/mysql/ent/schema/mixins"
)

type Sms struct {
	ent.Schema
}

func (Sms) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("notice_count").Default(0).Comment("通知短信数量"),
		field.Int64("used_notice").Default(0).Comment("已用通知"),
	}
}

func (Sms) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Sms) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Sms) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Sms) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_sms"},
		entsql.WithComments(true),
	}
}
