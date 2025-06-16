package common

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"kcers-survey/biz/dal/cache"
	db "kcers-survey/biz/dal/db/mysql"
	"kcers-survey/biz/dal/db/mysql/ent"
	"kcers-survey/biz/infras/do"
)

type Common struct {
	ctx   context.Context
	c     *app.RequestContext
	db    *ent.Client
	cache *ristretto.Cache
}

func NewCommon(ctx context.Context, c *app.RequestContext) do.Common {
	return &Common{
		ctx:   ctx,
		c:     c,
		db:    db.DB,
		cache: cache.Cache,
	}
}
