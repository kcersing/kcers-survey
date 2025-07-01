package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"kcers-survey/biz/dal/db/mysql/ent/schema/mixins"
)

type Area struct {
	ent.Schema
}

func (Area) Fields() []ent.Field {
	return []ent.Field{

		field.Int64("parent_id").Optional().Comment("上级区域ID"),
		field.Int64("level").Optional().Comment("行政区域等级 1-省 2-市 3-区县 4-街道镇"),

		field.String("name").Optional().Comment("名称"),
		field.String("whole_name").Optional().Unique().Comment("完整名称"),
		field.String("lon").Optional().Comment("本区域经度"),
		field.String("lat").Optional().Comment("本区域维度"),

		field.String("city_code").Optional().Comment("电话区号"),
		field.String("zip_code").Optional().Comment("邮政编码"),
		field.String("area_code").Optional().Comment("行政区划代码"),
		field.String("pin_yin").Optional().Comment("名称全拼"),
		field.String("simple_py").Optional().Comment("首字母简拼"),
		field.String("per_pin_yin").Optional().Comment("区域名称拼音的第一个字母"),
	}
}

func (Area) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Area) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Area) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_area"},
		entsql.WithComments(true),
	}
}
