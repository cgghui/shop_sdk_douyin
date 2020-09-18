package product

import "github.com/cgghui/shop_sdk_douyin/unit"

// ResponseCategory ProductCategory方法的响应结果
type ResponseCategory struct {
	ID   unit.ProductCID `mapstructure:"id"`
	Name string          `mapstructure:"name"`
}
