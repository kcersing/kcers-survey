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
	"kcers-survey/biz/dal/db/mysql/ent/area"
	"kcers-survey/biz/dal/db/mysql/ent/predicate"
	"kcers-survey/idl_gen/model/base"
	"strconv"

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

func (s *Sys) Area(req *sys.SysListReq) (list []*base.Tree, total int64, err error) {
	all, err := s.db.Area.
		Query().
		Where(area.ParentID(0)).
		All(s.ctx)
	if err != nil {
		return nil, 0, err
	}
	for _, v := range all {
		list = append(list, &base.Tree{
			Title: v.Name,
			Value: strconv.FormatInt(v.ID, 10),
		})
	}
	total = int64(len(all))
	return list, total, nil
}

func (s *Sys) City(req *sys.SysListReq) (list []*base.Tree, total int64, err error) {
	all, err := s.db.Area.
		Query().
		Where(area.ParentID(req.ID)).
		All(s.ctx)
	if err != nil {
		return nil, 0, err
	}
	for _, v := range all {
		list = append(list, &base.Tree{
			Title: v.Name,
			Value: strconv.FormatInt(v.ID, 10),
		})
	}
	total = int64(len(all))
	return list, total, nil
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
