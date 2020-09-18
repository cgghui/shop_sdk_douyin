package shop_sdk_douyin

// ResponseSkuList SkuList方法的响应结果
type ResponseSkuList struct {
	BaseSkuInfo     `mapstructure:",squash"`
	ID              uint64 `mapstructure:"id"`
	OpenUserID      uint64 `mapstructure:"open_user_id"`
	OutSkuID        uint64 `mapstructure:"out_sku_id"`
	SpecDetailID1   uint64 `mapstructure:"spec_detail_id1"`
	SpecDetailID2   uint64 `mapstructure:"spec_detail_id2"`
	SpecDetailID3   uint64 `mapstructure:"spec_detail_id3"`
	SpecDetailName1 string `mapstructure:"spec_detail_name1"`
	SpecDetailName2 string `mapstructure:"spec_detail_name2"`
	SpecDetailName3 string `mapstructure:"spec_detail_name3"`
}

type BaseSkuInfo struct {
	StockNum uint32 `mapstructure:"stock_num"` // 库存余量
	Price    uint32 `mapstructure:"price"`     // 价格
	Code     string `mapstructure:"code"`      // 商家自定义的sku代码
}
