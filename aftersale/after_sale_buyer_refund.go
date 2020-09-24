package aftersale

import "github.com/cgghui/shop_sdk_douyin/unit"

type ArgAfterSaleBuyerRefund struct {
	OrderID  unit.OrderID `paramName:"order_id"`
	Type     RSR          `paramName:"type"`
	Comment  CommX        `paramName:"comment,optional"`  // type = 2 时需要选择拒绝原因
	Evidence string       `paramName:"evidence,optional"` // type = 2 时需要上传图片凭证
}
