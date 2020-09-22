package unit

import "strconv"

// Price 商品价格 以分为单位
type Price uint32

// PriceToYuan 以元为参数 返回 Price
func PriceToYuan(price float64) Price {
	return Price(price * 100)
}

// Yuan 将 Price 输出为元
func (p Price) Yuan() float64 {
	return float64(p) / 100
}

// String 将 Price 输出为字符串
func (p Price) String() string {
	return strconv.FormatUint(uint64(p), 10)
}
