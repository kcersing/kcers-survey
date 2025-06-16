namespace go captcha

include "../base/base.thrift"

service CaptchaAdminService {
    //获取验证码
    base.NilResponse Captcha(1: base.Empty req) (api.post = "/service/captcha")
    base.NilResponse SmsCaptcha(1: CaptchaReq req) (api.post = "/service/sms-captcha")
    base.NilResponse ImgCaptcha(1: base.Empty req) (api.post = "/service/img-captcha")
}
struct CaptchaReq {
    1:  optional string mobile="" (api.raw = "mobile")
    2:  optional i64 type=1 (api.raw = "type")
}