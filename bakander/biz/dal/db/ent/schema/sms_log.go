package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"kcers-survey/biz/dal/db/mysql/ent/schema/mixins"
)

type SmsLog struct {
	ent.Schema
}

func (SmsLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("mobile").Comment("手机号"),
		field.String("biz_id").Comment("BizId"),
		field.String("code").Comment("验证码"),
		field.String("content").Default("").Comment("内容"),
		field.Int64("notify_type").Comment("通知类型[1会员;2员工]").Optional(),
		field.String("template").Comment("短信模板"),
	}
}

func (SmsLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (SmsLog) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (SmsLog) Indexes() []ent.Index {
	return []ent.Index{}
}

func (SmsLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_sms_log"},
		entsql.WithComments(true),
	}
}
