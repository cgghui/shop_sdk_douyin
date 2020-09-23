package order

import (
	"github.com/cgghui/shop_sdk_douyin/unit"
	"reflect"
	"time"
)

// ArgServiceList ServiceList方法的参数
type ArgServiceList struct {
	StartTime time.Time    `paramName:"start_time"`
	EndTime   time.Time    `paramName:"end_time"`
	Status    unit.StructS `paramName:"status"`
	Supply    unit.StructS `paramName:"supply"`        //是否获取分销商服务申请，0获取本店铺的服务申请，1获取分销商的服务申请
	Page      uint8        `paramName:"page,optional"` // 页数（默认值为0，第一页从0开始）
	Size      uint8        `paramName:"size,optional"` // 每页订单数（默认为10，最大100）
}

func (a ArgServiceList) HookConvertValue(f reflect.StructField, v reflect.Value) string {
	switch f.Type.String() {
	case "time.Time":
		return v.Interface().(time.Time).Format(unit.TimeYmdHis)
	case "unit.StructS":
		val := v.Interface().(unit.StructS).Value
		if val == nil {
			return "-" // 这个值将在HookSkipCheck中忽略
		}
		return val.(string)
	}
	return ""
}

func (a ArgServiceList) HookSkipCheck(t, _, v string) bool {
	if t == "unit.StructS" {
		if v == "-" {
			return true
		}
	}
	return false
}

// ResponseServiceList ServiceList方法的响应结果
type ResponseServiceList struct {
	Total uint32    `mapstructure:"total"`
	List  []Service `mapstructure:"list"`
}

// Service 服务请求
type Service struct {
	ID                unit.ServiceID `mapstructure:"id"`
	OrderID           unit.OrderID   `mapstructure:"order_id"`
	Reply             string         `mapstructure:"reply"`
	Detail            string         `mapstructure:"detail"`
	CreateTime        string         `mapstructure:"create_time"`
	OperateStatus     uint8          `mapstructure:"operate_status"`
	OperateStatusDesc string         `mapstructure:"operate_status_desc"`
	OperatorID        uint64         `mapstructure:"operator_id"`
	ReplyTime         string         `mapstructure:"reply_time"`
	ShopID            uint64         `mapstructure:"shop_id"`
	SupplyID          uint64         `mapstructure:"supply_id"`
	ReplyOpID         uint64         `mapstructure:"reply_op_id"`
}
