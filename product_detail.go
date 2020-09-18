package shop_sdk_douyin

import (
	"github.com/cgghui/shop_sdk_douyin/product_spec/spec"
)

// ResponseProductDetail ProductDetail方法的响应结果
type ResponseProductDetail struct {
	Product       `mapstructure:",squash"`
	Pic           ProductPic              `mapstructure:"pic"`            //
	ProductFormat string                  `mapstructure:"product_format"` //
	Specs         []spec.ProductSpecs     `mapstructure:"specs"`          // 商品选项 信息表
	SpecPics      []spec.ProductSpecPic   `mapstructure:"spec_pics"`      // 商品选项 图片表
	SpecPrices    []spec.ProductSpecPrice `mapstructure:"spec_prices"`    // 商品选项 价格表
}
