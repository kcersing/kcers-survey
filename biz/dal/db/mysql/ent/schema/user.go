package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"kcers-survey/biz/dal/db/mysql/ent/schema/mixins"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("mobile").Unique().Comment("mobile number | 手机号"),
		field.String("name").Optional().Comment("姓名"),
		field.Int64("gender").Default(3).Comment("性别 | [0:女性;1:男性;3:保密]").Optional(),

		field.String("username").Unique().Comment("user's login name | 登录名"),
		field.String("password").Comment("password | 密码"),

		field.JSON("functions", []string{}).Comment("functions | 职能"),
		field.Int64("job_time").Default(1).Comment("job time | [1:全职;2:兼职;]").Optional(),

		field.String("detail").Comment("详情").Optional(),

		field.String("side_mode").Optional().Default("dark").Comment("template mode | 布局方式"),
		field.String("base_color").Optional().Default("#fff").Comment("base color of template | 后台页面色调"),
		field.String("active_color").Optional().Default("#1890ff").Comment("active color of template | 当前激活的颜色设定"),

		field.String("email").Optional().Comment("email | 邮箱号"),
		field.String("wecom").Optional().Comment("wecom | 微信号"),

		field.String("organization").Optional().Comment("部门"),

		field.Int64("default_venue_id").Optional().Comment("登陆后默认场馆ID"),

		field.String("avatar").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Comment("avatar | 头像路径"),

		field.Time("birthday").Comment("出生日期").Optional(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{

		edge.To("token", Token.Type).Unique(),

		edge.To("roles", Role.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("username"),
		index.Fields("mobile"),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_users"},
		entsql.WithComments(true),
	}
}
