package system

import (
	"context"
	"github.com/bytedance/sonic"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"kcers-survey/biz/dal/cache"
	"kcers-survey/biz/dal/config"
	db "kcers-survey/biz/dal/db/mysql"
	"kcers-survey/biz/dal/db/mysql/ent"
	"kcers-survey/biz/dal/db/mysql/ent/api"
	"kcers-survey/biz/dal/db/mysql/ent/predicate"
	"kcers-survey/biz/infras/do"
	"kcers-survey/idl_gen/model/base"
	"kcers-survey/idl_gen/model/menu"
	"time"
)

type Api struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (a *Api) ApiTree(req *menu.ListApiReq) (resp []*base.Tree, total int, err error) {

	inter, exist := a.cache.Get("ApiTree")
	if exist {
		if v, ok := inter.([]*base.Tree); ok {
			return v, len(v), nil
		}
	}

	var predicates []predicate.API

	apis, err := a.db.API.Query().All(a.ctx)
	if err != nil {
		err = errors.Wrap(err, "get api list failed")
		return resp, total, err
	}
	apiGroups, err := a.db.API.Query().GroupBy(api.FieldAPIGroup).Strings(a.ctx)

	for _, apiGroup := range apiGroups {
		g := &base.Tree{
			Title: apiGroup,
		}
		for _, v := range apis {
			value, _ := sonic.Marshal(map[string]string{"path": v.Path, "method": v.Method})
			if v.APIGroup == g.Title {
				g.Children = append(g.Children, &base.Tree{
					Title: v.Title,
					Value: string(value),
					Key:   strconv.FormatInt(v.ID, 10), // map[string]string{"path": v.Path, "method": v.Method},
				})
			}
		}

		resp = append(resp, g)
	}

	total, _ = a.db.API.Query().Where(predicates...).Count(a.ctx)

	a.cache.SetWithTTL("ApiTree", &resp, 1, 30*time.Hour)
	return
}

func (a *Api) Create(req *menu.ApiInfo) error {
	_, err := a.db.API.Create().
		SetPath(req.Path).
		SetDescription(req.Description).
		SetAPIGroup(req.Group).
		SetMethod(req.Method).
		Save(a.ctx)
	if err != nil {
		err = errors.Wrap(err, "create Api failed")
		return err
	}
	return nil
}

func (a *Api) Update(req *menu.ApiInfo) error {
	_, err := a.db.API.UpdateOneID(req.ID).
		SetPath(req.Path).
		SetDescription(req.Description).
		SetAPIGroup(req.Group).
		SetMethod(req.Method).
		Save(a.ctx)
	if err != nil {
		err = errors.Wrap(err, "update Api failed")
		return err
	}
	return nil
}

func (a *Api) Delete(id int64) error {
	err := a.db.API.DeleteOneID(id).Exec(a.ctx)
	return err
}

func (a *Api) List(req *menu.ListApiReq) (resp []*menu.ApiInfo, total int, err error) {
	var predicates []predicate.API
	//if req.Path != "" {
	//	predicates = append(predicates, api.PathContains(req.Path))
	//}
	//if req.Description != "" {
	//	predicates = append(predicates, api.DescriptionContains(req.Description))
	//}
	//if req.Method != "" {
	//	predicates = append(predicates, api.MethodContains(req.Method))
	//}
	//if req.Group != "" {
	//	predicates = append(predicates, api.APIGroupContains(req.Group))
	//}

	apis, err := a.db.API.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(a.ctx)
	if err != nil {
		err = errors.Wrap(err, "get api list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &apis)
	if err != nil {
		err = errors.Wrap(err, "copy Api failed")
		return resp, 0, err
	}
	for i, v := range apis {
		resp[i].CreatedAt = v.CreatedAt.Format(time.DateTime)
		resp[i].UpdatedAt = v.UpdatedAt.Format(time.DateTime)

	}

	total, _ = a.db.API.Query().Where(predicates...).Count(a.ctx)
	return resp, total, nil
}

func NewApi(ctx context.Context, c *app.RequestContext) do.Api {
	return &Api{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
