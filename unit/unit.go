package unit

import "strconv"

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

// GoodsP 预售类型
type GoodsP uint8

func (t GoodsP) String() string {
	return strconv.FormatUint(uint64(t), 10)
}

const (
	GoodsPOff GoodsP = iota // 非预售
	GoodsPOn                // 全款预售
)

type BoolInt uint8

const (
	FalseInt BoolInt = iota
	TrueInt
)

type BoolStr string

const (
	FalseStr BoolStr = "0"
	TrueStr  BoolStr = "1"
)

const TimeYmd = "2006-01-02"
const TimeYmdHis = "2006-01-02 15:04:05"
