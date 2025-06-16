package enums

func ReturnMemberProductStatusValues(key int64) (values string) {
	switch key {
	case 1:
		values = "未激活"
	case 2:
		values = "正常"
	case 3:
		values = "过期"
	case 4:
		values = "退费"
	default:
		values = "状态异常"
	}
	return
}
