package aftersale

import (
	"github.com/cgghui/shop_sdk_douyin/order"
	"github.com/cgghui/shop_sdk_douyin/unit"
)

type ResponseAfterSaleRefundProcessDetail struct {
	OrderInfo         order.Detail      `mapstructure:"order_info"`
	Process           ProcessInfo       `mapstructure:"process_info"`
	Logs              []Log             `mapstructure:"logs"`
	RefundTotalAmount unit.Price        `mapstructure:"refund_total_amount"`
	RefundPostAmount  unit.Price        `mapstructure:"refund_post_amount"`
	ProcessList       []ProcessInfoList `mapstructure:"process_info_list"`
}

type ProcessInfoList struct {
	Process ProcessInfo `mapstructure:"process_info"`
	Logs    []Log       `mapstructure:"process_info"`
}

type ProcessInfo struct {
	ApplyTime       string   `mapstructure:"apply_time"`       // demo 2020-07-11 00:29:20
	Reason          string   `mapstructure:"reason"`           // demo 其他
	TypeDesc        string   `mapstructure:"type_desc"`        // demo 退货退款
	ApplyRemark     string   `mapstructure:"apply_remark"`     // demo 到货后试了衣服，我太高了，所以衣服不合身
	Evidence        []string `mapstructure:"evidence"`         // 凭证图片列表
	LogisticsTime   string   `mapstructure:"logistics_time"`   // 买家填写退货物流时间
	LogisticsCode   string   `mapstructure:"logistics_code"`   // 退货物流单号
	LogisticsName   string   `mapstructure:"logistics_name"`   // 退货快递公司
	AfterSaleID     uint64   `mapstructure:"aftersale_id"`     // 售后单ID
	AfterSaleStatus ASS      `mapstructure:"aftersale_status"` // 售后单状态
}

type Log struct {
	CreateTime string   `mapstructure:"create_time"` // demo 2020-07-11 00:29:20
	Action     string   `mapstructure:"action"`      // demo 申请退货
	Desc       string   `mapstructure:"desc"`        // demo 退货理由：其他
	Operator   string   `mapstructure:"operator"`    // demo 67939615271
	Evidence   []string `mapstructure:"evidence"`    // 凭证图片列表
}
