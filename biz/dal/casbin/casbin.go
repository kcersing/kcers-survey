package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	entAdapter "github.com/casbin/ent-adapter"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"kcers-survey/biz/dal/config"
)

var casbinEnforcer *casbin.Enforcer

func InitCasbin() {
	var err error
	casbinEnforcer, err = newCasbin()
	if err != nil {
		hlog.Fatal(err)
	}

}

func CasbinEnforcer() *casbin.Enforcer {
	return casbinEnforcer
}

func newCasbin() (enforcer *casbin.Enforcer, err error) {
	adapter, err := entAdapter.NewAdapter("mysql", config.GlobalServerConfig.MySQLInfo.Host)

	//adapter, err := entAdapter.NewAdapter("pgx", config.GlobalServerConfig.PostgreSQLInfo.Host)

	if err != nil {
		hlog.Error(err)
		return
	}

	var text string
	if config.GlobalServerConfig.Casbin.ModelText == "" {
		text = `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
	} else {
		text = config.GlobalServerConfig.Casbin.ModelText
	}

	m, err := model.NewModelFromString(text)
	if err != nil {
		hlog.Error(err)
		return
	}

	enforcer, err = casbin.NewEnforcer(m, adapter)
	if err != nil {
		hlog.Error(err)
		return
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		hlog.Error(err)
		return
	}

	return
}
