package unit

import "strconv"

type (
	ProductID  string // 商品ID
	ProductCID uint16 // 商品分类id
	SpecID     uint64 // 规格选项ID
	SkuID      uint64 // SKU ID
)

const CidTOP ProductCID = 0 // 商品的最顶级分类

func (p ProductID) GetProductID() ProductID {
	return p
}

func (s SpecID) ToString() string {
	return strconv.FormatUint(uint64(s), 10)
}

const SPE1 = "|"
const SPE2 = "^"
const SPE3 = ","
