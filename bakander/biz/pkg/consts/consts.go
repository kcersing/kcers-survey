package consts

import "time"

const (
	FreeCar    = "kcers"
	Issuer     = "kcers"
	Admin      = "Admin"
	User       = "User"
	ThirtyDays = time.Hour * 24 * 30
	AccountID  = "accountId"
	ID         = "id"

	ApiConfigPath = "./config/config.yaml"

	HlogFilePath = "./tmp/hlog/logs/"

	ExportFilePath = "./tmp/export/"

	WechatFilePath = "./tmp/wechat/"
	MySqlDSN       = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	RabbitMqURI    = "amqp://%s:%s@%s:%d/"

	TCP                  = "tcp"
	FreePortAddress      = ""
	RedisProfileClientDB = 0
	RedisBlobClientDB    = 0
	RedisCarClientDB     = 0

	OCRUrl = "https://aip.baidubce.com/rest/2.0/ocr/v1/driving_license"

	GPTUrl = "https://api.openai.com/v1/chat/completions"
)
