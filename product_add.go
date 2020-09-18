package shop_sdk_douyin

// ProductListArg ProductAdd方法的参数
type ProductAddArg struct {
	OutProductID  uint64       `paramName:"out_product_id"` // 外部商品id，接入方的商品id需为数字字符串，max = int64
	Name          string       // 商品名称，最多30个字符，不能含emoj表情
	Pic           ProductPic   // 商品轮播图，多张图片用 "|" 分开，第一张图为主图，最多5张，至少600x600，大小不超过1M
	Description   ProductPic   // 商品描述，目前只支持图片。多张图片用 "|" 分开。不能用其他网站的文本粘贴，这样会出现css样式不兼容
	MarketPrice   ProductPrice `paramName:"market_price"`   // 划线价，单位分
	DiscountPrice ProductPrice `paramName:"discount_price"` // 售价，单位分
	FirstCid      PCid         `paramName:"first_cid"`      // 一级分类id（三个分类级别请确保从属正确）
	SecondCid     PCid         `paramName:"second_cid"`     // 二级分类id
	ThirdCid      PCid         `paramName:"third_cid"`      // 三级分类id
}

// SetMarketPrice 设置划线价
func (p *ProductAddArg) SetMarketPrice(price float64) *ProductAddArg {
	p.MarketPrice = ProductPriceYuan(price)
	return p
}

// SetMarketPrice 设置售价
func (p *ProductAddArg) SetDiscountPrice(price float64) *ProductAddArg {
	p.DiscountPrice = ProductPriceYuan(price)
	return p
}

// SetCid1 设置商品一级分类
func (p *ProductAddArg) SetCid1(c *ResponseProductCategory) *ProductAddArg {
	p.FirstCid = c.ID
	return p
}

// SetCid2 设置商品二级分类
func (p *ProductAddArg) SetCid2(c *ResponseProductCategory) *ProductAddArg {
	p.SecondCid = c.ID
	return p
}

// SetCid3 设置商品三级分类
func (p *ProductAddArg) SetCid3(c *ResponseProductCategory) *ProductAddArg {
	p.ThirdCid = c.ID
	return p
}
