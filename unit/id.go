package unit

import "strconv"

type (
	BrandID    uint64 // 已授权品牌ID
	ProductID  string // 商品ID
	ProductCID uint16 // 商品分类ID
	SpecID     uint64 // 规格选项ID
	SkuID      uint64 // SKU ID
	OrderID    string // 订单ID
	ServiceID  uint64 // 服务ID
	CompanyID  uint16 // 物流公司ID
)

const CidTOP ProductCID = 0 // 商品的最顶级分类

func (p ProductID) GetProductID() ProductID {
	if p == "" {
		panic("product id empty")
	}
	return p
}

func (o OrderID) GetParentID() OrderID {
	if o == "" {
		panic("order id empty")
	}
	return o
}

func (b BrandID) String() string {
	return strconv.FormatUint(uint64(b), 10)
}

func (p ProductCID) String() string {
	return strconv.FormatUint(uint64(p), 10)
}

func (s SpecID) String() string {
	return strconv.FormatUint(uint64(s), 10)
}

const SPE1 = "|"
const SPE2 = "^"
const SPE3 = ","
