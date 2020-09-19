package sku

import (
	"github.com/cgghui/shop_sdk_douyin/unit"
)

// ResponseList SkuList方法的响应结果
type ResponseList struct {
	Info            `mapstructure:",squash"`
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

type Info struct {
	StockNum uint16     `mapstructure:"stock_num" paramName:"stock_num"` // 库存余量
	Price    unit.Price `mapstructure:"price"`                           // 价格
	Code     string     `mapstructure:"code" paramName:",optional"`      // 商家自定义的sku代码
}
