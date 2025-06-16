package enums

// 声明每个枚举项的索引值
// 产品类型
const (
	Card   = "card"
	Course = "course"
	Class  = "class"
)

// ReturnProductTypeValues 获取产品类型
func ReturnProductTypeValues(key string) (values string) {
	switch key {
	case Card:
		values = "卡"
	case Course:
		values = "私教课"
	case Class:
		values = "团课"
	default:
		values = "类型异常"
	}
	return
}
