package sku

import (
	"github.com/cgghui/shop_sdk_douyin/unit"
)

// ResponseList SkuList方法的响应结果
type ResponseList struct {
	unit.SpecSkuInfo `mapstructure:",squash"`
	ID               uint64 `mapstructure:"id"`
	OpenUserID       uint64 `mapstructure:"open_user_id"`
	OutSkuID         uint64 `mapstructure:"out_sku_id"`
	SpecDetailID1    uint64 `mapstructure:"spec_detail_id1"`
	SpecDetailID2    uint64 `mapstructure:"spec_detail_id2"`
	SpecDetailID3    uint64 `mapstructure:"spec_detail_id3"`
	SpecDetailName1  string `mapstructure:"spec_detail_name1"`
	SpecDetailName2  string `mapstructure:"spec_detail_name2"`
	SpecDetailName3  string `mapstructure:"spec_detail_name3"`
}
