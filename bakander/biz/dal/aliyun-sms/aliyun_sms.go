package aliyun_sms

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dypnsapi20170525 "github.com/alibabacloud-go/dypnsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"kcers-survey/biz/dal/config"
	"sync"
)

func CreateClient() (_result *dypnsapi20170525.Client, _err error) {
	AccessKeyId := config.GlobalServerConfig.Aliyun.Access.AccessKeyId
	AccessKeySecret := config.GlobalServerConfig.Aliyun.Access.AccessKeySecret

	configSms := &openapi.Config{
		AccessKeyId:     &AccessKeyId,
		AccessKeySecret: &AccessKeySecret,
		Endpoint:        tea.String("dypnsapi.aliyuncs.com"),
	}
	_result = &dypnsapi20170525.Client{}
	_result, _err = dypnsapi20170525.NewClient(configSms)
	return _result, _err
}

var onceAliyunSms sync.Once

var AliyunSms *dypnsapi20170525.Client

func InitAliyunSms() {
	onceAliyunSms.Do(func() {
		AliyunSms, _ = CreateClient()
	})
}
