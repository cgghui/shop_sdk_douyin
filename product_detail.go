package shop_sdk_douyin

// ResponseProductDetail ProductDetail方法的响应结果
type ResponseProductDetail struct {
	Product       `mapstructure:",squash"`
	Pic           ProductPic         `mapstructure:"pic"`            //
	ProductFormat string             `mapstructure:"product_format"` //
	Specs         []ProductSpecs     `mapstructure:"specs"`          // 商品选项 信息表
	SpecPics      []ProductSpecPic   `mapstructure:"spec_pics"`      // 商品选项 图片表
	SpecPrices    []ProductSpecPrice `mapstructure:"spec_prices"`    // 商品选项 价格表
}
