package order

import (
	"github.com/cgghui/shop_sdk_douyin/unit"
	"reflect"
	"time"
)

// ArgList OrderList方法的参数
type ArgList struct {
	Status    SS           `paramName:"order_status,optional"`
	StartTime time.Time    `paramName:"start_time"`
	EndTime   time.Time    `paramName:"end_time"`
	OrderBy   string       `paramName:"order_by"`
	IsDesc    unit.BoolInt `paramName:"is_desc,optional"`
	Page      uint8        `paramName:"page,optional"`
	Size      uint8        `paramName:"size,optional"`
}

func (a ArgList) HookConvertValue(f reflect.StructField, v reflect.Value) string {
	if f.Type.String() == "time.Time" {
		return v.Interface().(time.Time).Format(unit.TimeYmdHis)
	}
	return ""
}

// ResponseList OrderList方法的响应结果
type ResponseList struct {
	Count uint32  `mapstructure:"count"`
	Total uint32  `mapstructure:"total"`
	List  []Order `mapstructure:"list"`
}

// Order 订单信息
type Order struct {
	OrderID          unit.OrderID `mapstructure:"order_id"`           // 订单ID
	ShopID           uint64       `mapstructure:"shop_id"`            // 店铺ID
	OpenID           interface{}  `mapstructure:"open_id"`            // 在抖音小程序下单时，买家的抖音小程序ID TODO 不知道是什么类型
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
	CouponInfo       Coupon       `mapstructure:"coupon_info"`        // 优惠券详情
	ShopCouponAmount unit.Price   `mapstructure:"shop_coupon_amount"` // 商家优惠券金额 (单位: 分)
}

// Child 子订单信息
type Child struct {
	Order
	PID          unit.OrderID      `mapstructure:"pid"`            // 父订单ID
	ProductName  string            `mapstructure:"product_name"`   // 商品名称
	ProductPic   string            `mapstructure:"product_pic"`    // 商品图片
	ComboID      unit.SkuID        `mapstructure:"combo_id"`       // 该子订单购买的商品 sku_id
	ComboAmount  unit.Price        `mapstructure:"combo_amount"`   // 该子订单所购买的sku的售价
	ComboNum     uint16            `mapstructure:"combo_num"`      // 该子订单所购买的sku的数量
	Code         string            `mapstructure:"code"`           // 该子订单购买的商品的编码 code
	SpecDesc     unit.PropertyOPTS `mapstructure:"spec_desc"`      // 该子订单所属商品规格描述
	CouponMetaID string            `mapstructure:"coupon_meta_id"` // 优惠券id

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
