package unit

import "strconv"

/////////////////////////////////////////////////////////////////////////////////
// PayType 付款方式
type PayT uint8

func (t PayT) String() string {
	return strconv.FormatUint(uint64(t), 10)
}

const (
	CashDelivery  PayT = iota // 货到付款
	OnlinePayment             // 在线支付
	Casual                    // 让客户选择
)

/////////////////////////////////////////////////////////////////////////////////
// GoodsP 预售类型
type GoodsP uint8

func (t GoodsP) String() string {
	return strconv.FormatUint(uint64(t), 10)
}

const (
	GoodsPOff GoodsP = iota // 非预售
	GoodsPOn                // 全款预售
)

/////////////////////////////////////////////////////////////////////////////////
// BoolInt 布尔int型
type BoolInt uint8

const (
	FalseInt BoolInt = iota
	TrueInt
)

/////////////////////////////////////////////////////////////////////////////////
// BoolStr 布尔string型
type BoolStr string

const (
	FalseStr BoolStr = "0"
	TrueStr  BoolStr = "1"
)

/////////////////////////////////////////////////////////////////////////////////
// StructS 通用状态（struct转换为map）
// unit.StructS{} 传入该值时，字段应当被忽略
// unit.StructS{Value:"1"}
type StructS struct {
	Value interface{}
}

const TimeYmd = "2006-01-02"
const TimeYmdHis = "2006-01-02 15:04:05"
