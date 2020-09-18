package product

import (
	"github.com/cgghui/shop_sdk_douyin/product/spec"
)

// ResponseDetail ProductDetail方法的响应结果
type ResponseDetail struct {
	Product       `mapstructure:",squash"`
	Pic           Pic          `mapstructure:"pic"`            //
	ProductFormat string       `mapstructure:"product_format"` //
	Specs         []spec.Specs `mapstructure:"specs"`          // 商品选项 信息表
	SpecPics      []spec.Pic   `mapstructure:"spec_pics"`      // 商品选项 图片表
	SpecPrices    []spec.Price `mapstructure:"spec_prices"`    // 商品选项 价格表
}
