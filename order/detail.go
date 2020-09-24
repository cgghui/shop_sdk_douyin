package order

import (
	"github.com/cgghui/shop_sdk_douyin/unit"
)

// Detail 订单信息
type Detail struct {
	OrderID          unit.OrderID `mapstructure:"order_id"`           // 订单ID
	ShopID           uint64       `mapstructure:"shop_id"`            // 店铺ID
	OpenID           interface{}  `mapstructure:"open_id"`            // 在抖音小程序下单，买家的抖音小程序ID TODO 不知道是什么类型
	PostAddr         Address      `mapstructure:"post_addr"`          // 收件人地址
	PostCode         string       `mapstructure:"post_code"`          // 邮政编码
	PostReceiver     string       `mapstructure:"post_receiver"`      // 收件人姓名
	PostTel          string       `mapstructure:"post_tel"`           // 收件人电话
	BuyerWords       string       `mapstructure:"buyer_words"`        // 买家备注
	SellerWords      string       `mapstructure:"seller_words"`       // 卖家备注
	LogisticsID      uint64       `mapstructure:"logistics_id"`       // 物流公司ID
	LogisticsCode    string       `mapstructure:"logistics_code"`     // 物流单号
	LogisticsTime    string       `mapstructure:"logistics_time"`     // 发货时间 string型unix时间戳
	ReceiptTime      string       `mapstructure:"receipt_time"`       // 收货时间 string型unix时间戳
	OrderStatus      SS           `mapstructure:"order_status"`       // 订单状态
	CreateTime       string       `mapstructure:"create_time"`        // 订单创建时间 string型unix时间戳
	UpdateTime       uint64       `mapstructure:"update_time"`        // 订单更新时间 unix时间戳
	OrderType        OT           `mapstructure:"order_type"`         // 订单类型 (0:普通订单，2:虚拟订单，4:电子券)
	ExpShipTime      uint64       `mapstructure:"exp_ship_time"`      // 订单最晚发货时间 unix时间戳
	CancelReason     string       `mapstructure:"cancel_reason"`      // 订单取消原因
	PayType          PT           `mapstructure:"pay_type"`           // 支付类型 (0：货到付款，1：微信，2：支付宝）
	PayTime          string       `mapstructure:"pay_time"`           // 支付时间 string型unix时间戳
	PostAmount       unit.Price   `mapstructure:"post_amount"`        // 邮费金额 (单位: 分)
	CouponAmount     unit.Price   `mapstructure:"coupon_amount"`      // 平台优惠券金额 (单位: 分)
	CouponInfo       []Coupon     `mapstructure:"coupon_info"`        // 优惠券详情
	ShopCouponAmount unit.Price   `mapstructure:"shop_coupon_amount"` // 商家优惠券金额 (单位: 分)
	ShopFullCampaign interface{}  `mapstructure:"shop_full_campaign"` // TODO 不知道干麻用的 未知的一个字段
	OrderTotalAmount unit.Price   `mapstructure:"order_total_amount"` // 订单实付金额（不包含运费）
	IsComment        unit.BoolStr `mapstructure:"is_comment"`         // 是否评价 (1:已评价)
	UrgeCnt          uint8        `mapstructure:"urge_cnt"`           // 催单次数
	BType            BT           `mapstructure:"b_type"`             // 订单APP渠道
	SubBType         SBT          `mapstructure:"sub_b_type"`         // 订单来源类型 0:未知 1:app 2:小程序 3:h5
	CBiz             CB           `mapstructure:"c_biz"`              // 订单业务类型
	CType            interface{}  `mapstructure:"c_type"`             // TODO 不知道干麻用的 未知的一个字段
	Child            []Child      `mapstructure:"child"`              // 子订单列表
}

func (d Detail) GetParentID() unit.OrderID {
	return d.OrderID.GetParentID()
}

// Child 子订单信息
type Child struct {
	Detail            `mapstructure:",squash"`
	PID               unit.OrderID      `mapstructure:"pid"`                // 父订单ID
	OutProductID      uint64            `mapstructure:"out_product_id"`     // 该子订单购买的商品外部id
	OutSkuID          uint64            `mapstructure:"out_sku_id"`         // 该子订单购买的商品外部 sku_id
	ProductName       string            `mapstructure:"product_name"`       // 商品名称
	ProductPic        string            `mapstructure:"product_pic"`        // 商品图片
	ComboID           unit.SkuID        `mapstructure:"combo_id"`           // 该子订单购买的商品 sku_id
	ComboAmount       unit.Price        `mapstructure:"combo_amount"`       // 该子订单所购买的sku的售价
	ComboNum          uint16            `mapstructure:"combo_num"`          // 该子订单所购买的sku的数量
	Code              string            `mapstructure:"code"`               // 该子订单购买的商品的编码 code
	SpecDesc          unit.PropertyOPTS `mapstructure:"spec_desc"`          // 该子订单所属商品规格描述
	FinalStatus       SS                `mapstructure:"final_status"`       // 子订单状态
	PreSaleType       interface{}       `mapstructure:"pre_sale_type"`      // 订单预售类型 (1:全款预售订单)
	CouponMetaID      string            `mapstructure:"coupon_meta_id"`     // 优惠券id
	CampaignID        string            `mapstructure:"campaign_id"`        // 活动id
	CampaignInfo      []Campaign        `mapstructure:"campaign_info"`      // 活动细则 (title为活动标题)
	WarehouseID       interface{}       `mapstructure:"warehouse_id"`       // 仓库ID
	OutWarehouseID    interface{}       `mapstructure:"out_warehouse_id"`   // 仓库外部ID
	WarehouseSupplier interface{}       `mapstructure:"warehouse_supplier"` // 供应商ID
}

func (c Child) GetParentID() unit.OrderID {
	return c.PID.GetParentID()
}

// Address 收货地址
type Address struct {
	City     unit.Relation `mapstructure:"city"`
	Detail   string        `mapstructure:"detail"`
	Province unit.Relation `mapstructure:"province"`
	Town     unit.Relation `mapstructure:"town"`
}

// Coupon 优惠券
type Coupon struct {
	ID          uint64     `mapstructure:"id"`
	Name        string     `mapstructure:"name"`
	Description string     `mapstructure:"description"`
	Credit      unit.Price `mapstructure:"credit"` // 优惠金额, 单位分
	Type        CT         `mapstructure:"type"`
	Discount    float64    `mapstructure:"discount"`
}

// Campaign 活动细则
type Campaign struct {
	ID    string `mapstructure:"campaign_id"`
	Title string `mapstructure:"title"`
}
