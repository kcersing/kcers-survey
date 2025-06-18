package sms

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func CreateClient() (_result *dysmsapi20170525.Client, _err error) {

	AccessKeyId := ""
	AccessKeySecret := ""

	configSms := &openapi.Config{
		AccessKeyId:     &AccessKeyId,
		AccessKeySecret: &AccessKeySecret,
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
	}
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(configSms)
	return _result, _err
}
func SendAliyunSms(req dysmsapi20170525.SendSmsRequest) (bizId string, _err error) {
	client, _ := CreateClient()

	// 1.发送短信
	sendReq := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  req.PhoneNumbers,
		SignName:      req.SignName,
		TemplateCode:  req.TemplateCode,
		TemplateParam: req.TemplateParam,
	}
	sendResp, _err := client.SendSms(sendReq)
	if _err != nil {
		return "", _err
	}

	hlog.Info(sendResp)
	code := sendResp.Body.Code
	if !tea.BoolValue(util.EqualString(code, tea.String("OK"))) {
		hlog.Info(tea.String("错误信息: " + tea.StringValue(sendResp.Body.Message)))
		return "", _err
	}

	bizId = *sendResp.Body.BizId
	return bizId, nil
}
