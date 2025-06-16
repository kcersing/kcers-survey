package do

type Captcha interface {
	GetCaptcha() (id, b64s, answer string, err error)
}
