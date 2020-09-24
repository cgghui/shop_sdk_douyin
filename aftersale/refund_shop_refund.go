package aftersale

import (
	"github.com/cgghui/shop_sdk_douyin/unit"
)

// ArgRefundShopRefund RefundShopRefund方法的参数
type ArgRefundShopRefund struct {
	OrderID       unit.OrderID   `paramName:"order_id"`                // 父订单ID，须带字母A
	Type          RSR            `paramName:"type"`                    // 1：同意退款 2：不同意退款
	LogisticsID   unit.CompanyID `paramName:"logistics_id,optional"`   // type = 2 时必须填写物流公司ID
	LogisticsCode string         `paramName:"logistics_code,optional"` // type = 2 时必须填写物流单号
}
