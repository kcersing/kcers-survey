package do

import (
	"kcers-survey/idl_gen/model/sms"
)

type Sms interface {
	SendVerificationCode(mobile string, types int64) (err error)
	CheckVerificationCode(mobile, verificationCode string, types int64) (err error)

	SendList(req sms.SmsSendListReq) (resp []*sms.SmsSend, total int, err error)
	Info(id int64) (resp *sms.SmsInfo, err error)
}
