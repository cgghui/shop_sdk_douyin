package shop_sdk_douyin

// PCid 商品分类id
type PCid uint16

const PCidTOP PCid = 0

// ResponseProductCategory ProductCategory方法的响应结果
type ResponseProductCategory struct {
	ID   PCid   `mapstructure:"id"`
	Name string `mapstructure:"name"`
}
