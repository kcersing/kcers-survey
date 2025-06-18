package system

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/mojocn/base64Captcha"
	"kcers-survey/biz/dal/config"
	"kcers-survey/biz/infras/do"
)

type Captcha struct {
}

func (c *Captcha) GetCaptcha() (id, b64s, answer string, err error) {
	captchaDriver := base64Captcha.NewDriverDigit(
		config.GlobalServerConfig.Captcha.ImgHeight,
		config.GlobalServerConfig.Captcha.ImgWidth,
		config.GlobalServerConfig.Captcha.KeyLong,
		0.7,
		80,
	)
	captchaGen := base64Captcha.NewCaptcha(captchaDriver, base64Captcha.DefaultMemStore)

	id, b64s, answer, err = captchaGen.Generate()
	return

}

func NewCaptcha(ctx context.Context, c *app.RequestContext) do.Captcha {
	return &Captcha{}
}
