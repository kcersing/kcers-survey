package system

import (
	"context"
	"encoding/json"
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"kcers-survey/biz/dal/cache"
	casbin2 "kcers-survey/biz/dal/casbin"
	"kcers-survey/biz/dal/config"
	db "kcers-survey/biz/dal/db/mysql"
	"kcers-survey/biz/dal/db/mysql/ent"
	"kcers-survey/biz/dal/db/mysql/ent/api"
	"kcers-survey/biz/dal/db/mysql/ent/role"
	"kcers-survey/biz/infras/do"
	"kcers-survey/idl_gen/model/auth"
	"strconv"
)

type Auth struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	Cbs   *casbin.Enforcer
	db    *ent.Client
	cache *ristretto.Cache
}

func (a *Auth) QueryApiAll(id []int64) (resp []*auth.ApiAuthInfo, err error) {

	//ApiAuthInterface, exist := a.cache.Get("apiAll")
	//if exist {
	//	if u, ok := ApiAuthInterface.([]*do.ApiAuthInfo); ok {
	//		return u, nil
	//	}
	//}
	all, err := a.db.API.Query().Where(api.IDIn(id...)).All(a.ctx)
	if err != nil {
		return resp, err
	}

	err = copier.Copy(&resp, &all)
	if err != nil {
		err = errors.Wrap(err, "copy api all failed")
		return resp, err
	}

	//a.cache.SetWithTTL("apiAll", &resp, 1, 30*time.Hour)
	return

}

func (a *Auth) UpdateApiAuth(roleIDStr string, apis []int64) error {
	// clear old policies
	var oldPolicies [][]string
	oldPolicies, _ = a.Cbs.GetFilteredPolicy(0, roleIDStr)
	if len(oldPolicies) != 0 {
		removeResult, err := a.Cbs.RemoveFilteredPolicy(0, roleIDStr)
		if err != nil {
			return err
		}
		if !removeResult {
			return errors.New("casbin policies remove failed")
		}
	}
	infos, _ := a.QueryApiAll(apis)
	// add new policies
	var policies [][]string
	for _, v := range infos {
		policies = append(policies, []string{roleIDStr, v.Path, v.Method})
	}
	addResult, err := a.Cbs.AddPolicies(policies)
	if err != nil {
		return err
	}
	if !addResult {
		return errors.New("casbin policies add failed")
	}
	roleId, _ := strconv.ParseUint(roleIDStr, 10, 64)

	jsonBytes, _ := json.Marshal(apis)
	var intSlice []int
	err = json.Unmarshal(jsonBytes, &intSlice)
	if err != nil {
		return err
	}
	a.db.Role.Update().Where(role.ID(int64(roleId))).SetApis(intSlice).Save(a.ctx)
	return nil
}

func (a *Auth) ApiAuth(roleIDStr string) (infos []*auth.ApiAuthInfo, err error) {

	policies, _ := a.Cbs.GetFilteredPolicy(0, roleIDStr)
	for _, v := range policies {
		infos = append(infos, &auth.ApiAuthInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return
}

func (a *Auth) UpdateMenuAuth(roleID int64, menuIDs []int64) error {
	tx, err := a.db.Tx(a.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction err")
	}
	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				hlog.Error("UpdateMenuAuthority err:", err, "rollback err:", rollbackErr)
			}
		}
	}()

	//tx.Role.UpdateOneID(roleID).ClearMenus().Exec(a.ctx)
	err = tx.Role.UpdateOneID(roleID).ClearMenus().Exec(a.ctx)
	if err != nil {
		return errors.Wrap(err, "delete role's menu failed, error")
	}

	err = tx.Role.UpdateOneID(roleID).AddMenuIDs(menuIDs...).Exec(a.ctx)
	if err != nil {
		return errors.Wrap(err, "add role's menu failed, error")
	}

	return tx.Commit()
}

func (a *Auth) MenuAuth(roleID int64) (menuIDs []int64, err error) {
	menus, err := a.db.Role.Query().Where(role.IDEQ(roleID)).QueryMenus().All(a.ctx)
	for _, v := range menus {
		if v.ID != 1 {
			menuIDs = append(menuIDs, v.ID)
		}
	}
	return
}

func NewAuth(ctx context.Context, c *app.RequestContext) do.Auth {
	return &Auth{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
		Cbs:   casbin2.CasbinEnforcer(),
	}
}
