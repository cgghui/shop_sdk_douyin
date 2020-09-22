package product

import (
	"github.com/cgghui/shop_sdk_douyin/unit"
	"strings"
)

/////////////////////////////////////////////////////////////////////

// Pic 商品图片集合
type Pic []string

func NewPic() *Pic {
	p := make(Pic, 0)
	return &p
}

// Add 添加图片
func (p *Pic) Add(imgURL string) *Pic {
	*p = append(*p, imgURL)
	return p
}

// JoinString 将图片拼合抖音规定的格式
func (p Pic) JoinString() string {
	return strings.Join(p, unit.SPE1)
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
