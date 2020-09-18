package sku

import (
	"github.com/cgghui/shop_sdk_douyin/unit"
)

// ArgAdd ProductAdd方法的参数
type ArgAdd struct {
	unit.SpecSkuInfo
	ProductStrID    unit.ProductID  `paramName:"product_id"`       // 商品id
	outSkuID        uint64          `paramName:"out_sku_id"`       // 业务方自己的sku_id，唯一需为数字字符串，max = int64
	SpecID          unit.SpecID     `paramName:"spec_id"`          // 规格id，依赖/spec/list接口的返回
	SpecDetailIDS   [][]unit.SpecID `paramName:"spec_detail_ids	"` // 子规格id,最多3级,如 100041|150041|160041 （ 女款|白色|XL）
	SettlementPrice unit.Price      `paramName:"settlement_price"` // 结算价格 (单位 分)
}

func NewArgAdd() {

}
