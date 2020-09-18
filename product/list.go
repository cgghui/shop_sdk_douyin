package product

import "github.com/cgghui/shop_sdk_douyin/unit"

// ArgList ProductList方法的参数
type ArgList struct {
	Page        uint8 // 第几页（第一页为0）
	Size        uint8 // 每页返回条数
	Status      SS    // 指定状态返回商品列表
	CheckStatus SC    `paramName:"check_status"` // 指定审核状态返回商品列表
}

// ResponseList ProductList方法的响应结果
type ResponseList struct {
	All         uint32    `mapstructure:"all"`          // 商品总数
	AllPages    uint32    `mapstructure:"all_pages"`    // 已当前size所得的分页数
	Count       uint32    `mapstructure:"count"`        // 当前条件data返回结果数量
	CurrentPage uint32    `mapstructure:"current_page"` // 当前页
	Data        []Product `mapstructure:"data"`         // 商品列表
	PageSize    uint32    `mapstructure:"page_size"`    // 每页条数
}

// Product 商品基础信息
type Product struct {
	ProductStrID    unit.ProductID  `mapstructure:"product_id_str"`   // 商品id的字符串版本
	ProductID       uint64          `mapstructure:"product_id"`       // 商品id todo 好像不准，也不知道这是为什么
	OpenUserID      uint64          `mapstructure:"open_user_id"`     //
	Name            string          `mapstructure:"name"`             // 商品名称，最多30个字符，不能含emoj表情
	Description     string          `mapstructure:"description"`      // 商品描述，目前只支持图片。多张图片用 "|" 分开
	Img             string          `mapstructure:"img"`              //
	MarketPrice     unit.Price      `mapstructure:"market_price"`     // 划线价，单位分
	DiscountPrice   unit.Price      `mapstructure:"discount_price"`   // 售价，单位分
	SettlementPrice unit.Price      `mapstructure:"settlement_price"` //
	Status          SS              `mapstructure:"status"`           //
	SpecID          uint64          `mapstructure:"spec_id"`          // 规格id, 要先创建商品通用规格, 如颜色-尺寸
	CheckStatus     SC              `mapstructure:"check_status"`     //
	Mobile          string          `mapstructure:"mobile"`           // 客服号码
	FirstCid        unit.ProductCID `mapstructure:"first_cid"`        // 一级分类id（三个分类级别请确保从属正确）
	SecondCid       unit.ProductCID `mapstructure:"second_cid"`       // 二级分类id
	ThirdCid        unit.ProductCID `mapstructure:"third_cid"`        // 三级分类id
	PayType         PT              `mapstructure:"pay_type"`         // 支付方式，必填，0-货到付款，1-在线支付，2-二者都支持
	RecommendRemark string          `mapstructure:"recommend_remark"` // 商家推荐语，不能含emoj表情
	IsCreate        uint8           `mapstructure:"is_create"`
	CreateTime      string          `mapstructure:"create_time"`
	UpdateTime      string          `mapstructure:"update_time"`
}
