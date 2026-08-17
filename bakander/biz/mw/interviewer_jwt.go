package mw

import (
	"context"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"

	"kcers-survey/biz/dal/config"
	interviewerService "kcers-survey/biz/infras/service/interviewer"
	"kcers-survey/biz/pkg/errno"
	"kcers-survey/biz/pkg/utils"
)

type interviewerLogin struct {
	Mobile   string `form:"mobile,required" json:"mobile,required"`
	Password string `form:"password,required" json:"password,required"`
}

var interviewerJwtMiddleware *jwt.HertzJWTMiddleware

func GetInterviewerJWTMw() *jwt.HertzJWTMiddleware {
	if interviewerJwtMiddleware != nil {
		return interviewerJwtMiddleware
	}
	mw, err := newInterviewerJWT()
	if err != nil {
		hlog.Fatal(err, "Interviewer JWT Init Error")
	}
	interviewerJwtMiddleware = mw
	return interviewerJwtMiddleware
}

func newInterviewerJWT() (*jwt.HertzJWTMiddleware, error) {
	return jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "saas-interviewer",
		Key:         []byte(config.GlobalServerConfig.Auth.AccessSecret),
		Timeout:     time.Duration(config.GlobalServerConfig.Auth.AccessExpire) * time.Second,
		MaxRefresh:  time.Hour,
		IdentityKey: "interviewer-id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(map[string]interface{}); ok {
				return jwt.MapClaims{
					"interviewer-id": v,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			payloadMap, ok := claims["interviewer-id"].(map[string]interface{})
			if !ok {
				hlog.Error("get interviewer payloadMap error:", " claims data:", claims["interviewer-id"])
				return nil
			}
			c.Set("interviewerUserId", payloadMap["userId"])
			c.Set("interviewerMobile", payloadMap["mobile"])
			return payloadMap
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVal interviewerLogin
			if err := c.BindAndValidate(&loginVal); err != nil {
				return nil, err
			}

			s := interviewerService.NewInterviewer(ctx, c)
			resp, err := s.Login(loginVal.Mobile, loginVal.Password)
			if err != nil {
				return nil, err
			}

			roleIDs, _ := s.GetRoleIDs(resp.UserId)
			roleValues, _ := s.GetRoleValues(resp.UserId)

			payLoadMap := make(map[string]interface{})
			payLoadMap["userId"] = strconv.Itoa(int(resp.UserId))
			payLoadMap["mobile"] = resp.Mobile
			payLoadMap["roleIds"] = roleIDs
			payLoadMap["roleValues"] = roleValues
			return payLoadMap, nil
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			payloadMap, ok := data.(map[string]interface{})
			if !ok {
				return false
			}
			roleValuesRaw, ok := payloadMap["roleValues"].([]interface{})
			if !ok {
				return false
			}
			for _, v := range roleValuesRaw {
				if vStr, ok := v.(string); ok && vStr == interviewerService.InterviewerRoleValue {
					return true
				}
			}
			return false
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			utils.SendResponse(c, errno.NewErrNo(10002, "您没有访问此资源的权限"), message, 0, "")
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			utils.SendResponse(c, errno.Success,
				map[string]interface{}{
					"token":  token,
					"expire": expire.Format(time.RFC3339),
				}, 0, "")
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "interviewer jwt biz err = %+v", e.Error())
			return e.Error()
		},
	})
}
