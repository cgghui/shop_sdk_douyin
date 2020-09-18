package product

import (
	"strings"
)

/////////////////////////////////////////////////////////////////////

// Pic 商品图片集合
type Pic []string

// Add 添加图片
func (p *Pic) Add(imgURL string) *Pic {
	*p = append(*p, imgURL)
	return p
}

// Distinct 去除重复的图片
func (p Pic) Distinct() Pic {
	x := map[string]string{}
	for _, imgURL := range p {
		x[imgURL] = imgURL
	}
	r := make(Pic, len(x))
	i := 0
	for imgURL, _ := range x {
		r[i] = imgURL
		i++
	}
	return r
}

// JoinString 将图片拼合抖音规定的格式
func (p Pic) JoinString() string {
	return strings.Join(p.Distinct(), "|")
}

/////////////////////////////////////////////////////////////////////

// SS 商品上下架状态
type SS uint8

const (
	StatusOn  SS = iota // 上架
	StatusOff           // 下架
)

/////////////////////////////////////////////////////////////////////

// SC 商品审核状态
type SC uint8

const (
	CheckNot    SC = iota + 1 // 未提审
	CheckIng                  // 审核中
	CheckPass                 // 审核通过
	CheckReject               // 审核驳回
	CheckForbid               // 封禁
)

/////////////////////////////////////////////////////////////////////

// PT 付款方式
type PT uint8

const (
	CashDelivery  PT = iota // 货到付款
	OnlinePayment           // 在线支付
	Casual                  // 让客户选择
)
