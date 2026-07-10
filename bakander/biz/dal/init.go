package dal

import (
	"context"

	aliyun_sms "kcers-survey/biz/dal/aliyun-sms"
	"kcers-survey/biz/dal/cache"
	"kcers-survey/biz/dal/casbin"
	"kcers-survey/biz/dal/config"
	db "kcers-survey/biz/dal/db"
	"kcers-survey/biz/dal/logger"
	"kcers-survey/biz/infras/cron"
	"kcers-survey/biz/pkg/upload"

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
	config.GlobalUploadService = upload.Init(upload.Config{
		Host:     config.GlobalServerConfig.RabbitMQ.Host,
		Port:     config.GlobalServerConfig.RabbitMQ.Port,
		User:     config.GlobalServerConfig.RabbitMQ.User,
		Password: config.GlobalServerConfig.RabbitMQ.Password,
	})
	hlog.Info("Init upload ok!")
	go func() {
		if err := config.GlobalUploadService.RunImageUpload(context.Background()); err != nil {
			hlog.Fatal("upload image service err", err)
		}
	}()
	//go func() {
	//	if err := config.GlobalUploadService.RunVideoUpload(context.Background()); err != nil {
	//		hlog.Fatal("upload video service err", err)
	//	}
	//}()
	//go func() {
	//	if err := config.GlobalUploadService.RunDocUpload(context.Background()); err != nil {
	//		hlog.Fatal("upload doc service err", err)
	//	}
	//}()
	aliyun_sms.InitAliyunSms()
	hlog.Info("Init aliyun sms ok!")
	//go func() {
	//	wechat.InitWXPaymentApp()
	//	wechat.InitMiniProgramApp()
	//}()
	hlog.Info("Init ok!")
	cron.InitCron()

}
