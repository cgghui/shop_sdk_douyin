package spec

import (
	"github.com/cgghui/shop_sdk_douyin/product/sku"
	"github.com/cgghui/shop_sdk_douyin/unit"
)

// 规格关系
type PLeaf uint8

const (
	PLeafN PLeaf = iota // 父级
	PLeafY              // 子级
)

// ResponseDetail SpecDetail方法的响应结果
type ResponseDetail struct {
	ID    unit.SpecID
	Name  string  `mapstructure:"name"`  //
	Specs []Specs `mapstructure:"specs"` // 商品选项 信息表
}

// Specs 商品选项<信息表> 用pid进行关系绑定
type Specs struct {
	Spec   `mapstructure:",squash"` // 主选项规格
	Values []Spec                   `mapstructure:"values"` // 子选项规格
}

// Spec 商品选项信息
type Spec struct {
	ID     uint64 `mapstructure:"id"`      // 规格创建时生成的id 父子各不相同
	SpecID uint64 `mapstructure:"spec_id"` // 项id 这个是抖音系统自带的  有：颜色、尺码、长度等 父子同值
	Name   string `mapstructure:"name"`    // 名称
	PID    uint64 `mapstructure:"pid"`     // 父级id 如果本身就是父级，则为0
	IsLeaf PLeaf  `mapstructure:"is_leaf"` // 是否为父级
	Status uint8  `mapstructure:"status"`  // todo 目前还不知道这字段是什么意思
}

// Pic 商品选项<图片表>
type Pic struct {
	SpecDetailID uint64 `mapstructure:"spec_detail_id"` // 规格id 与 ProductSpec.ID 对应
	Pic          string `mapstructure:"pic"`            // 图片路径
}

// Price 商品选项<价格表>
type Price struct {
	sku.Info        `mapstructure:",squash"`
	SkuID           uint64     `mapstructure:"sku_id"`           // todo 目前还不知道这字段是什么意思
	SpecDetailIDS   []uint64   `mapstructure:"spec_detail_ids"`  // 规格id 与 ProductSpec.ID 对应
	SettlementPrice unit.Price `mapstructure:"settlement_price"` // 结算价格
}
