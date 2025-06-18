package mixins

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// BaseMixin implements the ent.Mixin for sharing
// base fields with package schemas.
type BaseMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Comment("primary key"),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Comment("created time").
			Optional(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("last update time").
			Optional(),
		field.Int64("delete").
			Default(0).
			Comment("last delete  1:已删除").
			Optional(),
		field.Int64("created_id").Default(0).Comment("created").
			Optional(),
	}
}
