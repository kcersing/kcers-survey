package system

// This file is auto-generated, don't edit it. Thanks.

import (
	"context"
	"errors"
	"fmt"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/cloudwego/hertz/pkg/app"
	"kcers-survey/biz/dal/cache"
	"kcers-survey/biz/dal/config"
	db "kcers-survey/biz/dal/db/mysql"
	"kcers-survey/biz/dal/sms"
	"kcers-survey/biz/infras/do"
	"math/rand"
	"time"
)

func (s *Sms) SendVerificationCode(mobile string, types int64) (err error) {

	_, exist := s.cache.Get("VerificationCode_" + mobile)
	if exist {
		err = errors.New("请勿重复发送验证码")
		return
	}
	verifyCode := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	bizId, err := sms.SendAliyunSms(
		dysmsapi20170525.SendSmsRequest{
			SignName:      tea.String(config.GlobalServerConfig.Aliyun.Sms.Captcha.SignName),
			TemplateCode:  tea.String(config.GlobalServerConfig.Aliyun.Sms.Captcha.TemplateCode),
			PhoneNumbers:  tea.String(mobile),
			TemplateParam: tea.String(fmt.Sprintf("{\"code\":\"%s\"}", verifyCode)),
		},
	)

	if err != nil {
		return err
	}
	s.cache.SetWithTTL("VerificationCode_"+mobile, verifyCode, 1, 5*time.Minute)

	go func() {
		if bizId != "" {
			s.db.SmsLog.Create().
				SetMobile(mobile).
				SetCode(verifyCode).
				SetBizID(bizId).
				SetTemplate(config.GlobalServerConfig.Aliyun.Sms.Captcha.TemplateCode).
				SetNotifyType(types).
				Save(s.ctx)
		}
	}()

	return nil
}

func (s *Sms) CheckVerificationCode(mobile, verificationCode string, types int64) (err error) {
	code, exist := s.cache.Get("VerificationCode_" + mobile)
	if !exist {
		err = errors.New("验证码已失效")
		return
	}
	if vc, ok := code.(string); ok {
		if vc != verificationCode {
			err = errors.New("验证码输入错误")
			return
		}
	} else {
		err = errors.New("验证码解析错误")
		return
	}

	return nil
}

func NewAliyunSms(ctx context.Context, c *app.RequestContext) do.Sms {
	return &Sms{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
