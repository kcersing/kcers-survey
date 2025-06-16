package system

import (
	"fmt"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"math/rand"
	"saas/biz/dal/sms"
	"testing"
	"time"
)

func TestSms(t *testing.T) {

	verifyCode := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	sms.SendAliyunSms(
		dysmsapi20170525.SendSmsRequest{
			SignName:      tea.String("验证码短信"),
			TemplateCode:  tea.String("SMS_314756319"),
			PhoneNumbers:  tea.String("13937173036"),
			TemplateParam: tea.String(fmt.Sprintf("{\"code\":\"%s\"}", verifyCode)),
		},
	)

}
