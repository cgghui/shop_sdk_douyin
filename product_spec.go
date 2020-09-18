package shop_sdk_douyin

// SpecID 商品选项ID
type SpecID uint32

// ResponseSpecList SpecList方法的响应结果
type ResponseSpecList struct {
	ID   SpecID `mapstructure:"id"`   // 项id
	Name string `mapstructure:"name"` // 项名称
}
