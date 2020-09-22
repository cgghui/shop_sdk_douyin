package sku

import "github.com/cgghui/shop_sdk_douyin/unit"

// ResponseDetail SkuDetail的响应结果
type ResponseDetail struct {
	Detail
}

type Detail struct {
	Info            `mapstructure:",squash"`
	ID              unit.SkuID     `mapstructure:"id"`
	OutSkuID        uint64         `mapstructure:"out_sku_id"`
	OpenUserID      uint64         `mapstructure:"open_user_id"`
	ProductID       uint64         `mapstructure:"product_id"`
	ProductStrID    unit.ProductID `mapstructure:"product_id_str"`
	SpecID          unit.SpecID    `mapstructure:"spec_id"`
	SpecDetailID1   unit.SpecID    `mapstructure:"spec_detail_id1"`
	SpecDetailID2   unit.SpecID    `mapstructure:"spec_detail_id2"`
	SpecDetailID3   unit.SpecID    `mapstructure:"spec_detail_id3"`
	SpecDetailName1 string         `mapstructure:"spec_detail_name1"`
	SpecDetailName2 string         `mapstructure:"spec_detail_name2"`
	SpecDetailName3 string         `mapstructure:"spec_detail_name3"`
	SettlementPrice unit.Price     `mapstructure:"settlement_price"`
	CreateTime      uint64         `mapstructure:"create_time"`
	SkuType         uint8          `mapstructure:"sku_type"`
}

func (d Detail) GetProductID() unit.ProductID {
	return d.ProductStrID.GetProductID()
}

func (d Detail) GetSkuID() unit.SkuID {
	return d.ID
}
