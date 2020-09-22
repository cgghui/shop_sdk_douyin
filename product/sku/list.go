package sku

import (
	"github.com/cgghui/shop_sdk_douyin/unit"
)

// ResponseList SkuList方法的响应结果
type ResponseList struct {
	ResponseDetail `mapstructure:",squash"`
}

type Info struct {
	StockNum uint16     `mapstructure:"stock_num" paramName:"stock_num"` // 库存余量
	Price    unit.Price `mapstructure:"price"`                           // 价格
	Code     string     `mapstructure:"code" paramName:",optional"`      // 商家自定义的sku代码
}
