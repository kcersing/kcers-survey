package wechat

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"kcers-survey/biz/dal/config"
	"kcers-survey/biz/pkg/consts"
	"os"
	"sync"
)

var PaymentWechatApp *payment.Payment
var oncePaymentWechat sync.Once

func InitWXPaymentApp() {
	oncePaymentWechat.Do(func() {
		var err error
		PaymentWechatApp, err = NewWXPaymentApp()
		if err != nil || PaymentWechatApp == nil {
			hlog.Error("NewWXPaymentApp err: %s", err)
		}

	})
}

const TRANSACTION_SUCCESS = "TRANSACTION.SUCCESS"
const TRANSACTION_FAILED = "TRANSACTION.FAILED"

func NewWXPaymentApp() (*payment.Payment, error) {

	conf := config.GlobalServerConfig.Wechat
	var cache kernel.CacheInterface
	if config.GlobalServerConfig.Redis.Host != "" {
		cache = kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{config.GlobalServerConfig.Redis.Host},
			Password: config.GlobalServerConfig.Redis.Password,
			DB:       7,
		})
	}
	wechatFilePath := consts.WechatFilePath
	if err := os.MkdirAll(wechatFilePath, 0o777); err != nil {
		panic(err)
	}
	Payment, err := payment.NewPayment(&payment.UserConfig{
		AppID:              conf.Appid,    // 小程序、公众号或者企业微信的appid
		MchID:              conf.MchId,    // 商户号 appID
		MchApiV3Key:        conf.ApiV3Key, //
		Key:                conf.ApiKey,
		CertPath:           conf.CertFileContent,
		KeyPath:            conf.KeyFileContent,
		SerialNo:           conf.SerialNo,
		CertificateKeyPath: conf.CertificateKeyPath,
		WechatPaySerial:    conf.WechatPaySerialNo,
		//RSAPublicKeyPath:   conf.RSAPublicKeyPath,
		NotifyURL: conf.NotifyUrl,
		//SubMchID:           conf.MchId,
		//SubAppID:           conf.Appid,
		ResponseType: response.TYPE_MAP,
		Cache:        cache,
		Log: payment.Log{
			Level: "debug",
			File:  wechatFilePath + "wechatpay.log",
		},
		Http: payment.Http{
			Timeout: 30.0,
			//BaseURI: "http://127.0.0.1:8888",
			BaseURI: "https://api.mch.weixin.qq.com",
		},

		HttpDebug: false,
		Debug:     false,
		//Debug:     true,
	})

	return Payment, err
}
