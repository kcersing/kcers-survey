package enums

func ReturnScheduleStatusValues(key int64) (values string) {
	switch key {
	case 1:
		values = "未发布"
	case 2:
		values = "待上课"
	case 3:
		values = "已上课"
	case 4:
		values = "已下课"
	case 5:
		values = "取消"
	case 6:
		values = "爽约"
	default:
		values = "状态异常"
	}
	return
}
func ReturnScheduleMemberStatusName(key int64) (values string) {
	switch key {
	case 1:
		values = "待上课"
	case 2:
		values = "上课中"
	case 3:
		values = "已下课"
	case 4:
		values = "爽约"
	case 5:
		values = "取消"
	case 6:
		values = "爽约"
	default:
		values = "状态异常"
	}
	return
}
