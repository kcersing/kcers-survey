package config

type ServerConfig struct {
	Name           string           `mapstructure:"Name" json:"Name"`
	Host           string           `mapstructure:"Host" json:"Host"`
	Port           int              `mapstructure:"Port" json:"Port"`
	Timeout        int              `mapstructure:"Timeout" json:"Timeout"`
	IsProd         bool             `mapstructure:"IsProd" json:"IsProd"`
	Domain         string           `mapstructure:"Domain" json:"Domain"`
	MySQLInfo      MySQLConfig      `mapstructure:"MySQL" json:"MySQL"`
	PostgreSQLInfo PostgreSQLConfig `mapstructure:"PostgreSQL" json:"PostgreSQL"`
	Captcha        Captcha          `mapstructure:"Captcha" json:"Captcha"`
	Auth           Auth             `mapstructure:"Auth" json:"Auth"`
	Redis          Redis            `mapstructure:"Redis" json:"Redis"`
	Casbin         CasbinConf       `mapstructure:"Casbin" json:"Casbin"`
	Minio          Minio            `mapstructure:"Minio" json:"Minio"`
	Swagger        Swagger          `mapstructure:"Swagger" json:"Swagger"`
	Aliyun         Aliyun           `mapstructure:"Aliyun" json:"Aliyun"`
	Wechat         Wechat           `mapstructure:"Wechat" json:"Wechat"`
	Alipay         AliPay           `mapstructure:"Alipay" yaml:"Alipay"`
	PicHost        string           `mapstructure:"PicHost" json:"PicHost"`
}

type MySQLConfig struct {
	Host string `mapstructure:"Host" json:"Host"`
	Salt string `mapstructure:"Salt" json:"Salt"`
}
type PostgreSQLConfig struct {
	Host string `mapstructure:"Host" json:"Host"`
	Salt string `mapstructure:"Salt" json:"Salt"`
}
type Captcha struct {
	KeyLong   int `mapstructure:"KeyLong" json:"KeyLong"`
	ImgWidth  int `mapstructure:"ImgWidth" json:"ImgWidth"`
	ImgHeight int `mapstructure:"ImgHeight" json:"ImgHeight"`
}
type Auth struct {
	OAuthKey     string `mapstructure:"OAuthKey" json:"OAuthKey"`
	AccessSecret string `mapstructure:"AccessSecret" json:"AccessSecret"`
	AccessExpire int    `mapstructure:"AccessExpire" json:"AccessExpire"`
}

type Redis struct {
	Host     string `mapstructure:"Host" json:"Host"`
	Password string `mapstructure:"Password" json:"Password"`
}

type CasbinConf struct {
	ModelText string `mapstructure:"ModelText" json:"ModelText"`
}

type Wechat struct {
	Appid              string `mapstructure:"appid" yaml:"appid"`
	AppSecret          string `mapstructure:"app_secret" yaml:"app_secret"`
	MchId              string `mapstructure:"mch_id" yaml:"mch_id"`
	ApiKey             string `mapstructure:"api_key" yaml:"api_key"`
	ApiV3Key           string `mapstructure:"api_v3_key" yaml:"api_v3_key"`
	CertFileContent    string `mapstructure:"cert_file_content" yaml:"cert_file_content"`
	KeyFileContent     string `mapstructure:"key_file_content" yaml:"key_file_content"`
	Pkcs12FileContent  string `mapstructure:"pkcs12_file_content" yaml:"pkcs12_file_content"`
	SerialNo           string `mapstructure:"serial_no" yaml:"serial_no"`
	NotifyUrl          string `mapstructure:"notify_url" yaml:"notify_url"`
	RefundNotifyUrl    string `mapstructure:"refund_notify_url" yaml:"refund_notify_url"`
	RSAPublicKeyPath   string `mapstructure:"rsa_public_key_path" yaml:"rsa_public_key_path"`
	WechatPaySerialNo  string `mapstructure:"wechat_pay_serial_no" yaml:"wechat_pay_serial_no"`
	CertificateKeyPath string `mapstructure:"certificate_key_path" yaml:"certificate_key_path"`
}

type AliPay struct {
	Appid                   string `mapstructure:"appid" yaml:"appid"`
	PrivateKey              string `mapstructure:"private_key" yaml:"private_key"`
	AppPublicCertContent    string `mapstructure:"app_public_cert_content" yaml:"app_public_cert_content"`
	AlipayRootCertContent   string `mapstructure:"alipay_root_cert_content" yaml:"alipay_root_cert_content"`
	AlipayPublicCertContent string `mapstructure:"alipay_public_cert_content" yaml:"alipay_public_cert_content"`
	NotifyUrl               string `mapstructure:"notify_url" yaml:"notify_url"`
}

type Minio struct {
	EndPoint        string `mapstructure:"EndPoint" yaml:"EndPoint"`
	AccessKeyID     string `mapstructure:"AccessKeyID" yaml:"AccessKeyID"`
	SecretAccessKey string `mapstructure:"SecretAccessKey" yaml:"SecretAccessKey"`
	UseSSL          bool   `mapstructure:"UseSSL" yaml:"UseSSL"`

	VideoBucketName string `mapstructure:"VideoBucketName" yaml:"VideoBucketName"`
	ImgBucketName   string `mapstructure:"ImgBucketName" yaml:"ImgBucketName"`

	Url string `mapstructure:"Url" yaml:"Url"`
}
type Swagger struct {
	Url string `mapstructure:"url" yaml:"url"`
}

type MiniProgram struct {
	AppID  string `mapstructure:"appid" yaml:"appid"`
	Secret string `mapstructure:"secret" yaml:"secret"`
	Token  string `mapstructure:"token" yaml:"token"`
	AESKey string `mapstructure:"aes_key" yaml:"aes_key"`

	AppKey  string `mapstructure:"app_key" yaml:"app_key"`
	OfferID string `mapstructure:"offer_id" yaml:"offer_id"`
}

type Aliyun struct {
	Access Access `mapstructure:"Access" yaml:"Access"`
	Sms    Sms    `mapstructure:"Sms" yaml:"Sms"`
}
type Access struct {
	AccessKeyId     string `mapstructure:"AccessKeyId" yaml:"AccessKeyId"`
	AccessKeySecret string `mapstructure:"AccessKeySecret" yaml:"AccessKeySecret"`
}
type Sms struct {
	Captcha SmsTemplate `mapstructure:"Captcha" yaml:"Captcha"`
}
type SmsTemplate struct {
	SignName     string `mapstructure:"SignName" yaml:"SignName"`
	TemplateCode string `mapstructure:"TemplateCode" yaml:"TemplateCode"`
}
