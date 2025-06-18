package system

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"kcers-survey/biz/dal/cache"
	"kcers-survey/biz/dal/config"
	db "kcers-survey/biz/dal/db/mysql"
	"kcers-survey/biz/dal/db/mysql/ent"
	logs2 "kcers-survey/biz/dal/db/mysql/ent/logs"
	"kcers-survey/biz/dal/db/mysql/ent/predicate"
	"kcers-survey/biz/infras/do"
	"kcers-survey/idl_gen/model/logs"
	"time"
)

type Logs struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func NewLogs(ctx context.Context, c *app.RequestContext) do.Logs {
	return &Logs{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}

func (l *Logs) Create(logsReq *logs.LogsInfo) error {
	err := l.db.Logs.Create().
		SetType(logsReq.Type).
		SetMethod(logsReq.Method).
		SetAPI(logsReq.API).
		SetSuccess(logsReq.Success).
		SetReqContent(logsReq.ReqContent).
		SetRespContent(logsReq.RespContent).
		SetIP(logsReq.IP).
		SetUserAgent(logsReq.UserAgent).
		SetOperatorsr(logsReq.Operatorsr).
		SetTime(logsReq.Time).
		SetIdentity(logsReq.Identity).
		Exec(l.ctx)
	if err != nil {
		err = errors.Wrap(err, "create logs failed")
		return err
	}
	return nil
}

func (l *Logs) List(req *logs.LogsListReq) (list []*logs.LogsInfo, total int, err error) {
	var predicates []predicate.Logs
	if req.Type != "" {
		predicates = append(predicates, logs2.TypeEQ(req.Type))
	}
	if req.Method != "" {
		predicates = append(predicates, logs2.MethodEQ(req.Method))
	}
	if req.API != "" {
		predicates = append(predicates, logs2.APIContains(req.API))
	}
	if req.Operatorsr != "" {
		predicates = append(predicates, logs2.OperatorsrContains(req.Operatorsr))
	}
	if req.Success != true {
		predicates = append(predicates, logs2.SuccessEQ(req.Success))
	}
	logsData, err := l.db.Logs.Query().Where(predicates...).
		Offset(int((req.Page - 1) * req.PageSize)).
		Limit(int(req.PageSize)).
		Order(ent.Desc(logs2.FieldCreatedAt)).All(l.ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "query logsData list failed")
	}
	for _, v := range logsData {
		list = append(list, &logs.LogsInfo{
			Type:        v.Type,
			Method:      v.Method,
			API:         v.API,
			Success:     v.Success,
			ReqContent:  v.ReqContent,
			RespContent: v.RespContent,
			IP:          v.IP,
			UserAgent:   v.UserAgent,
			Operatorsr:  v.Operatorsr,
			Identity:    v.Identity,
			Time:        v.Time,
			CreatedAt:   v.CreatedAt.Format(time.DateTime),
			UpdatedAt:   v.UpdatedAt.Format(time.DateTime),
		})
	}
	total, err = l.db.Logs.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "query logsData count failed")
	}
	return
}

func (l *Logs) DeleteAll() error {
	_, err := l.db.Logs.Delete().Exec(l.ctx)
	if err != nil {
		return errors.Wrap(err, "delete logsData failed")
	}
	return nil
}
