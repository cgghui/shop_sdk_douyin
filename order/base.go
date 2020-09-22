package order

// SS 订单状态
type SS uint8

const (
	WaitPay   SS = iota + 1 // 在线支付订单待支付；货到付款订单待确认
	StockUp                 // 备货中（只有此状态下，才可发货）
	Delivered               // 已发货
	Cancelled               // 已取消：1.用户未支付时取消订单；2.用户超时未支付，系统自动取消订单；3.货到付款订单，用户拒收
	Completed               // 已完成：1.在线支付订单，商家发货后，用户收货、拒收或者15天无物流更新；2.货到付款订单，用户确认收货
)

// OT 订单类型
type OT uint8

const (
	GeneralOrder      OT = 0 // 普通订单
	VirtualOrder      OT = 2 // 虚拟订单
	ElectronicVoucher OT = 4 // 电子券
)

// PT 支付方式
type PT uint8

const (
	CashDelivery PT = iota // 货到付款
	WeChat                 // 微信支付
	Alipay                 // 支付宝支付
)

// CT 优惠券类型
type CT uint8

const (
	PlatformDC CT = iota + 1  // 平台折扣券 (平台承担)
	PlatformDT                // 平台直减券 (平台承担)
	PlatformFD                // 平台满减券 (平台承担)
	CategoryDC CT = iota + 8  // 品类折扣券 (暂未开放)
	CategoryDT                // 品类直减券 (暂未开放)
	CategoryFD                // 品类满减券 (暂未开放)
	ShopDC     CT = iota + 15 // 店铺折扣券 (店铺承担)
	ShopDT                    // 店铺直减券 (店铺承担)
	ShopFD                    // 店铺满减券 (店铺承担)
	ChannelDC  CT = iota + 22 // 渠道折扣券 (平台承担)
	ChannelDT                 // 渠道直减券 (平台承担)
	ChannelFD                 // 渠道满减券 (平台承担)
	ProductDC  CT = iota + 29 // 单品折扣券 (店铺承担)
	ProductDT                 // 单品直减券 (店铺承担)
	ProductFD                 // 单品满减券 (店铺承担)
)
