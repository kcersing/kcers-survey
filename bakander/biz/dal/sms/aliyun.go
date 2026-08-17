package sms

import (
	dypnsapi20170525 "github.com/alibabacloud-go/dypnsapi-20170525/v3/client"
	aliyun_sms "kcers-survey/biz/dal/aliyun-sms"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func SendVerifyCode(phoneNumber, signName, templateCode, code string) (bizId string, err error) {
	client := aliyun_sms.AliyunSms
	if client == nil {
		hlog.Error("AliyunSms client is nil, sending SMS may fail")
		client, err = aliyun_sms.CreateClient()
		if err != nil {
			return "", err
		}
	}

	templateParam := `{"code":"` + code + `","min":"5"}`
	req := &dypnsapi20170525.SendSmsVerifyCodeRequest{
		PhoneNumber:   &phoneNumber,
		SignName:      &signName,
		TemplateCode:  &templateCode,
		TemplateParam: &templateParam,
	}

	resp, err := client.SendSmsVerifyCode(req)
	if err != nil {
		return "", err
	}

	hlog.Infof("SendSmsVerifyCode resp: %+v", resp)
	if resp.Body.Success != nil && *resp.Body.Success {
		if resp.Body.Model != nil && resp.Body.Model.BizId != nil {
			return *resp.Body.Model.BizId, nil
		}
		return "", nil
	}
	return "", nil
}
