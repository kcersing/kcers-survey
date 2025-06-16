package enums

func ReturnMemberConditionValues(key int64) (values string) {
	switch key {
	case 1:
		values = "潜在"
	case 2:
		values = "正式"
	default:
		values = "状态异常"
	}
	return
}
func ReturnMemberGenderValues(key int64) (values string) {
	switch key {
	case 1:
		values = "女"
	case 2:
		values = "男"
	case 3:
		values = "未知"
	default:
		values = "异常"
	}
	return
}
func ReturnMemberGenderKey(values string) (key int64) {
	switch values {
	case "女":
		key = 1
	case "男":
		key = 2
	case "未知":
		key = 3
	default:
		key = 0
	}
	return
}

// MPStatus 会员产品状态
type MPStatus int

const (
	MPStatusUnfinished = iota
	MPStatusNotActivated
	MPStatusActivated
	MPStatusExpire
	MPStatusExhaust
	MPStatusUpgrade
	MPStatusFreeze
)

var MPStatusNames = map[MPStatus]string{
	MPStatusUnfinished:   "未完成",
	MPStatusNotActivated: "未激活",
	MPStatusActivated:    "已激活",
	MPStatusExpire:       "已到期",
	MPStatusExhaust:      "已完结",
	MPStatusUpgrade:      "已升级",
	MPStatusFreeze:       "已冻结",
}

func (s MPStatus) String() string {
	return MPStatusNames[s]
}
