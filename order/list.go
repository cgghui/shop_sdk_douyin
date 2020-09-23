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
	Count uint32   `mapstructure:"count"`
	Total uint32   `mapstructure:"total"`
	List  []Detail `mapstructure:"list"`
}
