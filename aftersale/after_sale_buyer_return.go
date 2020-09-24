package aftersale

import "github.com/cgghui/shop_sdk_douyin/unit"

type ArgAfterSaleBuyerReturn struct {
	OrderID   unit.OrderID `paramName:"order_id"`
	Type      RSR          `paramName:"type"`
	SmsID     unit.BoolStr `paramName:"sms_id"`
	Comment   Comm         `paramName:"comment,optional"`    // type = 2 时需要选择拒绝原因
	Evidence  string       `paramName:"evidence,optional"`   // type = 2 时需要上传图片凭证
	AddressID string       `paramName:"address_id,optional"` // 商家退货物流收货地址id,不传则使用默认售后收货地址
}
