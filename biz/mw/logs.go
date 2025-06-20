package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"kcers-survey/biz/infras/service/system"
	"kcers-survey/biz/infras/service/user"
	"kcers-survey/idl_gen/model/logs"
	"strconv"
	"time"
)

func LogMw() app.HandlerFunc {

	return func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()
		c.Next(ctx)
		var log logs.LogsInfo
		log.Type = "Interface"
		log.Method = string(c.Request.Method())
		log.API = string(c.Request.Path())
		log.UserAgent = string(c.Request.Header.UserAgent())
		log.IP = c.ClientIP()

		reqBodyStr := string(c.Request.Body())
		if len(reqBodyStr) > 200 {
			reqBodyStr = reqBodyStr[:200]
		}
		log.ReqContent = reqBodyStr

		respBodyStr := string(c.Request.Body())
		if len(respBodyStr) > 200 {
			respBodyStr = respBodyStr[:200]
		}

		if c.Response.Header.StatusCode() == 200 {
			log.Success = true
		}

		costTime := time.Since(start).Milliseconds()
		log.Time = costTime

		var username = "Anonymous"

		userIn, exist := c.Get("user_id")
		if !(exist || userIn == nil) {
			userId := toInt(userIn)
			userInfo, _ := user.NewUser(ctx, c).Info(userId)
			if userInfo != nil {
				username = userInfo.Name
			}
			log.Operatorsr = username
			log.Identity = 2
		}

		err := system.NewLogs(ctx, c).Create(&log)
		if err != nil {
			hlog.Error(err)
		}

	}
}
func toInt(idIn interface{}) int64 {
	var idStr string
	var ok bool
	idStr, ok = idIn.(string)
	if !ok {
		idStr = "0"
	}
	id, _ := strconv.Atoi(idStr)
	return int64(id)
}
