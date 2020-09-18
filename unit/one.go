package unit

// SpecSkuInfo 规格和SKU共用的价格、库存
type SpecSkuInfo struct {
	StockNum uint16 `mapstructure:"stock_num" paramName:"stock_num"` // 库存余量
	Price    Price  `mapstructure:"price"`                           // 价格
	Code     string `mapstructure:"code"`                            // 商家自定义的sku代码
}
