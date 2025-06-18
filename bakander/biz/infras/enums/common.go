package enums

func ReturnMemberGenderValues(key int64) (values string) {
	switch key {
	case 1:
		values = "女"
	case 2:
		values = "男"
	case 3:
		values = "未知"
	default:
		values = "其他"
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
