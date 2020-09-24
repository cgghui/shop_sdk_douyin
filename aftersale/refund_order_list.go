package aftersale

import (
	"github.com/cgghui/shop_sdk_douyin/unit"
	"reflect"
	"time"
)

// ArgRefundOrderList RefundOrderList方法的参数
type ArgRefundOrderList struct {
	Type      RFD          `paramName:"type,optional"`
	StartTime time.Time    `paramName:"start_time,optional"`
	EndTime   time.Time    `paramName:"end_time,optional"`
	OrderBy   string       `paramName:"order_by"`
	IsDesc    unit.BoolInt `paramName:"is_desc,optional"`
	Page      uint8        `paramName:"page,optional"`
	Size      uint8        `paramName:"size,optional"`
}

func (a ArgRefundOrderList) HookConvertValue(f reflect.StructField, v reflect.Value) string {
	if f.Type.String() == "time.Time" {
		return v.Interface().(time.Time).Format(unit.TimeYmdHis)
	}
	return ""
}
