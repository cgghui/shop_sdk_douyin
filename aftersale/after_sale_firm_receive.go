package aftersale

import "github.com/cgghui/shop_sdk_douyin/unit"

type ArgAfterSaleFirmReceive struct {
	OrderID  unit.OrderID `paramName:"order_id"`          // 子订单ID
	Type     RSR          `paramName:"type"`              // 处理方式 1：确认收货并退款 2：拒绝
	Comment  Comm         `paramName:"comment,optional"`  // type = 2 时需要选择拒绝原因
	Evidence string       `paramName:"evidence,optional"` // type = 2 凭证图片（货到付款订单，必填）
	Register string       `paramName:"register"`
	Package  string       `paramName:"package"`
	Facade   string       `paramName:"facade"`
	Function string       `paramName:"function"`
}
