package shop_sdk_douyin

// PStatus 商品上下架状态
type PStatus uint8

const (
	PStatusOn  PStatus = iota // 上架
	PStatusOff                // 下架
)

// PCheck 商品审核状态
type PCheck uint8

const (
	PCheckNot    PCheck = iota + 1 // 未提审
	PCheckIng                      // 审核中
	PCheckPass                     // 审核通过
	PCheckReject                   // 审核驳回
	PCheckForbid                   // 封禁
)

// PayT 付款方式
type PayT uint8

const (
	CashDelivery  PayT = iota // 货到付款
	OnlinePayment             // 在线支付
	Casual                    // 让客户选择
)

// PCid 商品分类id
type PCid uint16

const PCidTOP PCid = 0

// ProductListArg ProductList方法的参数
type ProductListArg struct {
	Page        uint8   // 第几页（第一页为0）
	Size        uint8   // 每页返回条数
	Status      PStatus // 指定状态返回商品列表
	CheckStatus PCheck  `paramName:"check_status"` // 指定审核状态返回商品列表
}

// ResponseProductList ProductList方法的响应结果
type ResponseProductList struct {
	All         uint32    `mapstructure:"all"`          // 商品总数
	AllPages    uint32    `mapstructure:"all_pages"`    // 已当前size所得的分页数
	Count       uint32    `mapstructure:"count"`        // 当前条件data返回结果数量
	CurrentPage uint32    `mapstructure:"current_page"` // 当前页
	Data        []Product `mapstructure:"data"`         // 商品列表
	PageSize    uint32    `mapstructure:"page_size"`    // 每页条数
}

// Product 商品基础信息
type Product struct {
	ProductStrID    string  `mapstructure:"product_id_str"`   // 商品id的字符串版本
	ProductID       uint64  `mapstructure:"product_id"`       // 商品id todo 好像不准，也不知道这是为什么
	OpenUserID      uint64  `mapstructure:"open_user_id"`     //
	Name            string  `mapstructure:"name"`             // 商品名称，最多30个字符，不能含emoj表情
	Description     string  `mapstructure:"description"`      // 商品描述，目前只支持图片。多张图片用 "|" 分开
	Img             string  `mapstructure:"img"`              //
	MarketPrice     uint32  `mapstructure:"market_price"`     // 划线价，单位分
	DiscountPrice   uint32  `mapstructure:"discount_price"`   // 售价，单位分
	SettlementPrice uint32  `mapstructure:"settlement_price"` //
	Status          PStatus `mapstructure:"status"`           //
	SpecID          uint64  `mapstructure:"spec_id"`          // 规格id, 要先创建商品通用规格, 如颜色-尺寸
	CheckStatus     PCheck  `mapstructure:"check_status"`     //
	Mobile          string  `mapstructure:"mobile"`           // 客服号码
	FirstCid        PCid    `mapstructure:"first_cid"`        // 一级分类id（三个分类级别请确保从属正确）
	SecondCid       PCid    `mapstructure:"second_cid"`       // 二级分类id
	ThirdCid        PCid    `mapstructure:"third_cid"`        // 三级分类id
	PayType         PayT    `mapstructure:"pay_type"`         // 支付方式，必填，0-货到付款，1-在线支付，2-二者都支持
	RecommendRemark string  `mapstructure:"recommend_remark"` // 商家推荐语，不能含emoj表情
	IsCreate        uint8   `mapstructure:"is_create"`
	CreateTime      string  `mapstructure:"create_time"`
	UpdateTime      string  `mapstructure:"update_time"`
}
