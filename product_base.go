package shop_sdk_douyin

import "strings"

type ProductPic []string

func (p *ProductPic) Add(imgURL string) *ProductPic {
	*p = append(*p, imgURL)
	return p
}

func (p ProductPic) Distinct() ProductPic {
	x := map[string]string{}
	for _, imgURL := range p {
		x[imgURL] = imgURL
	}
	r := make(ProductPic, len(x))
	i := 0
	for imgURL, _ := range x {
		r[i] = imgURL
		i++
	}
	return r
}

func (p ProductPic) JoinString() string {
	return strings.Join(p.Distinct(), "|")
}

// ProductPrice 商品价格 以分为单位
type ProductPrice uint32

// ProductPriceYuan 以元为参数 返回ProductPrice
func ProductPriceYuan(price float64) ProductPrice {
	return ProductPrice(price * 100)
}

// Yuan 将ProductPrice输出为元
func (p ProductPrice) Yuan() float64 {
	return float64(p) / 100
}
