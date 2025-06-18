package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// StatusMixin implements the ent.Mixin for sharing
// status fields with package schemas.
type StatusMixin struct {
	mixin.Schema
}

func (StatusMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("status").
			Default(1).
			Optional().
			Comment("状态[0:禁用;1:正常]"),
	}
}
