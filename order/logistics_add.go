package order

import "github.com/cgghui/shop_sdk_douyin/unit"

type ArgLogisticsAdd struct {
	OrderID       unit.OrderID   `paramName:"order_id"`         // 父订单ID，由orderList接口返回
	LogisticsID   unit.CompanyID `paramName:"logistics_id"`     //
	Company       string         `paramName:"company,optional"` // 物流公司名称
	LogisticsCode string         `paramName:"logistics_code"`   // 运单号
}
