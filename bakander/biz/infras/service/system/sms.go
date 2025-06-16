package system

// This file is auto-generated, don't edit it. Thanks.

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"kcers-survey/biz/dal/db/mysql/ent"
	"kcers-survey/biz/dal/db/mysql/ent/predicate"
	sms2 "kcers-survey/biz/dal/db/mysql/ent/sms"
	"kcers-survey/biz/dal/db/mysql/ent/smslog"
	"kcers-survey/idl_gen/model/sms"
	"sync"
	"time"
)

type Sms struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
	mu    sync.Mutex
}

func (s *Sms) Info(id int64) (resp *sms.SmsInfo, err error) {

	first, err := s.db.Sms.Query().Where(sms2.IDEQ(id)).First(s.ctx)
	if first != nil {
		resp = &sms.SmsInfo{
			NoticeCount: first.NoticeCount,
			UsedNotice:  first.UsedNotice,
		}
	}
	return
}

func (s *Sms) SendList(req sms.SmsSendListReq) (resp []*sms.SmsSend, total int, err error) {
	var predicates []predicate.SmsLog

	if req.Mobile != "" {
		predicates = append(predicates, smslog.MobileEQ(req.Mobile))
	}

	all, err := s.db.SmsLog.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(smslog.FieldID)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Venue Sms log list failed")
		return resp, total, err
	}

	for _, l := range all {
		log := sms.SmsSend{
			CreatedAt:  l.CreatedAt.Format(time.DateTime),
			Status:     l.Status,
			Mobile:     l.Mobile,
			Code:       l.Code,
			BizId:      l.BizID,
			NotifyType: l.NotifyType,
			Content:    l.Content,
			Templates:  l.Template,
		}
		resp = append(resp, &log)
	}

	total, _ = s.db.SmsLog.Query().Where(predicates...).Count(s.ctx)
	return
}
