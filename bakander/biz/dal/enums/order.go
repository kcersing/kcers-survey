package enums

func ReturnOrderStatusValues(key int64) (values string) {
	switch key {
	case 1:
		values = "待支付"
	case 2:
		values = "订单取消"
	case 3:
		values = "订单失效"
	case 4:
		values = "订单退款完成"
	case 5:
		values = "订单已完成"
	default:
		values = "状态异常"
	}
	return
}
func ReturnOrderDeviceValues(key string) (values string) {
	switch key {
	case "pc":
		values = "电脑端"
	case "wxc":
		values = "微信小程序"

	default:
		values = "状态异常"
	}
	return
}
func ReturnOrderPayWayValues(key string) (values string) {
	switch key {
	case "pc":
		values = "电脑支付"
	case "wxc":
		values = "微信小程序支付"

	default:
		values = "状态异常"
	}
	return
}

type OrderEvent string

const (
	// 订单状态
	OrderStateCreated   int64 = 1 //订单已创建
	OrderStateCancelled       = 2 //订单取消
	OrderStateTimeout         = 3 //订单超时取消
	OrderStatRefunded         = 4 //订单退款完成
	OrderStateCompleted       = 5 //订单已完成

	// 订单事件
	OrderEventPay       OrderEvent = "pay"       //支付
	OrderEventCancelled            = "cancelled" //取消
	OrderEventShipped              = "shipped"   //发货
	OrderEventCompleted            = "completed" //完成
	OrderEventRefund               = "refund"    //退款

)

// 状态行为规则
var StateTransitions = map[int64]map[OrderEvent]bool{
	OrderStateCreated: {
		OrderEventPay:       true,
		OrderEventCancelled: true,
	},
	OrderStateCompleted: {
		OrderEventShipped:   true,
		OrderEventRefund:    true,
		OrderEventCompleted: true,
	},
	OrderStateCancelled: {
		OrderEventCancelled: true,
	},
	OrderStateTimeout: {
		OrderEventCancelled: true,
	},
	OrderStatRefunded: {
		OrderEventRefund: true,
	},
}
