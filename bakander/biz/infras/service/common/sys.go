package common

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"kcers-survey/biz/dal/cache"
	"kcers-survey/biz/dal/config"
	db "kcers-survey/biz/dal/db/mysql"
	"kcers-survey/biz/dal/db/mysql/ent"
	"kcers-survey/biz/dal/db/mysql/ent/predicate"

	"kcers-survey/biz/infras/do"

	"kcers-survey/idl_gen/model/sys"
)

type Sys struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (s *Sys) RoleList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error) {
	var predicates []predicate.Role

	lists, err := s.db.Role.Query().Where(predicates...).All(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Role list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, &sys.SysList{
			ID:   v.ID,
			Name: v.Remark,
		})
	}
	total = int64(len(list))

	return
}

func NewSys(ctx context.Context, c *app.RequestContext) do.Sys {
	return &Sys{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
