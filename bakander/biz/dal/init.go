package dal

import (
	"context"

	aliyun_sms "kcers-survey/biz/dal/aliyun-sms"
	"kcers-survey/biz/dal/cache"
	"kcers-survey/biz/dal/casbin"
	"kcers-survey/biz/dal/config"
	db "kcers-survey/biz/dal/db"
	"kcers-survey/biz/dal/db/ent/role"
	"kcers-survey/biz/dal/logger"
	"kcers-survey/biz/infras/cron"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func Init() {
	config.InitConfig()
	hlog.Info("Init config ok!")
	logger.InitLogger()
	hlog.Info("Init logger ok!")
	cache.InitCache()
	hlog.Info("Init cache ok!")
	db.InitDB()
	hlog.Info("Init db ok!")
	casbin.InitCasbin()
	hlog.Info("Init casbin ok!")
	aliyun_sms.InitAliyunSms()
	hlog.Info("Init aliyun sms ok!")
	//go func() {
	//	wechat.InitWXPaymentApp()
	//	wechat.InitMiniProgramApp()
	//}()
	hlog.Info("Init ok!")
	cron.InitCron()

	seedRoles()
}

func seedRoles() {
	ctx := context.Background()
	exist, _ := db.DB.Role.Query().Where(role.ValueEQ("interviewer")).Exist(ctx)
	if !exist {
		_, err := db.DB.Role.Create().
			SetName("调查员").
			SetValue("interviewer").
			SetDefaultRouter("dashboard").
			SetRemark("调查员角色，可查看和修改自己调研的问卷").
			SetOrderNo(10).
			SetStatus(1).
			Save(ctx)
		if err != nil {
			hlog.Error("create interviewer role failed:", err)
		} else {
			hlog.Info("seed interviewer role ok!")
		}
	}
}
