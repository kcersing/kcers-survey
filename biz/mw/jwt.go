package mw

import (
	"context"
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
	"kcers-survey/biz/dal/config"
	"kcers-survey/biz/infras/service/common"
	userService "kcers-survey/biz/infras/service/user"
	"kcers-survey/biz/pkg/errno"
	"kcers-survey/biz/pkg/utils"
	"kcers-survey/idl_gen/model/user"
	user3 "kcers-survey/idl_gen/model/user"
	"strconv"
	"time"
)

type jwtLogin struct {
	Username  string `form:"username,required" json:"username,required"`   //lint:ignore SA5008 ignoreCheck
	Password  string `form:"password,required" json:"password,required"`   //lint:ignore SA5008 ignoreCheck
	Captcha   string `form:"captcha,required" json:"captcha,required"`     //lint:ignore SA5008 ignoreCheck
	CaptchaID string `form:"captchaId,required" json:"captchaId,required"` //lint:ignore SA5008 ignoreCheck
}

// jwt identityKey
var (
	identityKey   = "jwt-id"
	jwtMiddleware = new(jwt.HertzJWTMiddleware)
)

func GetJWTMw(e *casbin.Enforcer) *jwt.HertzJWTMiddleware {
	jwtMiddleware, err := newJWT(e)
	if err != nil {
		hlog.Fatal(err, "JWT Init Error")
	}
	return jwtMiddleware
}

func newJWT(enforcer *casbin.Enforcer) (jwtMiddleware *jwt.HertzJWTMiddleware, err error) {

	jwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "saas",
		Key:         []byte(config.GlobalServerConfig.Auth.AccessSecret),
		Timeout:     time.Duration(config.GlobalServerConfig.Auth.AccessExpire) * time.Second,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		// PayloadFunc is used to define a custom token payload.
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(map[string]interface{}); ok {
				return jwt.MapClaims{
					identityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		// IdentityHandler is used to define a custom identity handler.
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			payloadMap, ok := claims[identityKey].(map[string]interface{})
			if !ok {
				hlog.Error("get payloadMap error:", " claims data:", claims[identityKey])
				return nil
			}
			c.Set("userId", payloadMap["userId"])
			c.Set("userRole", payloadMap["userRole"])
			c.Set("userRoleIds", payloadMap["userRoleIds"])
			c.Set("userType", payloadMap["userType"])

			return payloadMap
		},
		// Authenticator is used to validate the login data.
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			res := new(user.LoginResp)

			var loginVal jwtLogin
			if err := c.BindAndValidate(&loginVal); err != nil {
				return "", err
			}
			// 验证码

			username := loginVal.Username
			password := loginVal.Password
			res, err = userService.NewUser(ctx, c).Login(&user3.LoginReq{
				Username: username,
				Password: password,
			})
			if err != nil {
				return nil, errors.New("账号或密码错误")
			}

			payLoadMap := make(map[string]interface{})
			payLoadMap["userId"] = strconv.Itoa(int(res.UserId))
			payLoadMap["userRole"] = res.UserRole
			payLoadMap["userRoleIds"] = res.UserRoleIds
			payLoadMap["userType"] = strconv.Itoa(int(res.UserId))
			return payLoadMap, nil
		},
		// Authorizator is used to validate the authentication of the current request.
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			obj := string(c.URI().Path())
			act := string(c.Method())

			hlog.Info(obj, act)
			payloadMap, ok := data.(map[string]interface{})

			if !ok {
				hlog.Error("get payloadMap error:", " claims data:", data)
				return false
			}
			var userRoleIds []int64
			hlog.Info(payloadMap)
			userType, ok := payloadMap["userType"].(string)
			if !ok {
				hlog.Error("userType 解析错误", err)
				return false
			}
			userId, ok := payloadMap["userId"].(string)
			if !ok {
				hlog.Error("userId 解析错误")
				return false
			}

			hlog.Info(userId)
			if userType == "10" {

			}

			if userType == "1" {
				roleIds, ok := payloadMap["userRoleIds"].([]interface{})
				if !ok {
					hlog.Error("payloadMap:", payloadMap)
					return false
				}
				for _, v := range roleIds {
					i := v.(float64)
					userRoleIds = append(userRoleIds, int64(i))
				}
			}

			//
			//existToken := service.NewToken(ctx, c).IsExistByUserId(int64(userIdInt))
			//if !existToken {
			//	return false
			//}
			// check the role status
			//roleInfo, err := service.NewRole(ctx, c).RoleInfoByID(cast.ToInt64(roleId))
			//// if the role is not exist or the role is not active, return false
			//if err != nil {
			//	hlog.Error(err, "role is not exist")
			//	return false
			//}

			//if roleInfo.Status != 1 {
			//	hlog.Error("role cache is not a valid *ent.Role or the role is not active")
			//	return false
			//}

			//sub := roleId
			//check the permission
			//pass, err := enforcer.Enforce(sub, obj, act)
			//if err != nil {
			//	hlog.Error("casbin err,  role id: ", roleId, " path: ", obj, " method: ", act, " pass: ", pass, " err: ", err.Error())
			//	return false
			//}
			//if !pass {
			//	hlog.Info("casbin forbid role id: ", roleId, " path: ", obj, " method: ", act, " pass: ", pass)
			//}
			//hlog.Info("casbin allow role id: ", roleId, " path: ", obj, " method: ", act, " pass: ", pass)
			//return pass

			return true
		},
		LogoutResponse: func(ctx context.Context, c *app.RequestContext, code int) {
			id := common.GetTokenUserID(c)
			hlog.Info(id)
			err = common.NewToken(ctx, c).Delete(id)
			if err != nil {
				utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
			}
			utils.SendResponse(c, errno.Success, nil, 0, "")
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
		RefreshResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {

			utils.SendResponse(c, errno.Success,
				map[string]interface{}{
					"token":  token,
					"expire": expire.Format(time.RFC3339),
				}, 0, "")
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
	})

	return

}
