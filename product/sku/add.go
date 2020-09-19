package sku

import (
	"errors"
	"github.com/cgghui/shop_sdk_douyin/unit"
	"log"
	"strings"
)

type ResponseAdd interface {
}

// ArgAdd ProductAdd方法的参数
type ArgAdd struct {
	Info
	ProductStrID    unit.ProductID `paramName:"product_id"`                // 商品id
	OutSkuID        uint64         `paramName:"out_sku_id,optional"`       // 业务方自己的sku_id，唯一需为数字字符串，max = int64
	SpecID          string         `paramName:"spec_id"`                   // 规格id，依赖/spec/list接口的返回
	SettlementPrice unit.Price     `paramName:"settlement_price,optional"` // 结算价格 (单位 分)
	SpecDetailIDS   string         `paramName:"spec_detail_ids"`           // 子规格id,最多3级,如 100041|150041|160041 （ 女款|白色|XL）
	Multiple        bool           `paramName:"-"`
	_id             []unit.SpecID  `paramName:"-"`
	_ids            []string       `paramName:"-"`
}

func NewArgAdd(price float64, stock uint16, codes ...string) *ArgAdd {
	code := ""
	if len(codes) == 1 {
		code = codes[0]
	}
	r := ArgAdd{
		Info: Info{
			StockNum: stock,
			Price:    unit.PriceToYuan(price),
			Code:     code,
		},
		_ids: make([]string, 0),
	}
	return &r
}

// Build 取出参数集
func (a *ArgAdd) Build() (ArgAdd, error) {
	if len(a._id) == 0 {
		return ArgAdd{}, errors.New("SpecID empty")
	}
	tmp := make([]string, len(a._id))
	for i, val := range a._id {
		tmp[i] = val.ToString()
	}
	a.SpecID = strings.Join(tmp, "|")
	tmp = make([]string, 0)
	for _, spec := range a._ids {
		tmp = append(tmp, spec)
	}
	if len(a._ids) > 1 {
		a.Multiple = true
	} else {
		a.Multiple = false
	}
	a.SpecDetailIDS = strings.Join(tmp, "^")
	a._id = nil
	a._ids = nil
	ret := *a
	*a = ArgAdd{}
	return ret, nil
}

// SetProduct 设置SKU归属商品
func (a *ArgAdd) SetProduct(product unit.Product) *ArgAdd {
	a.ProductStrID = product.GetProductID()
	return a
}

// SetOutSkuID 业务方自己的sku_id，唯一需为数字字符串，max = int64
func (a *ArgAdd) SetOutSkuID(id uint64) *ArgAdd {
	a.OutSkuID = id
	return a
}

// SetSpecID 设置商品主项规格id
func (a *ArgAdd) SetSpecID(spec unit.ProductSpec) *SpecBox {
	specID := spec.GetProductSpecID()
	a._id = append(a._id, specID)
	return &SpecBox{a: a, x: spec, id: specID, ss: make([]string, 0)}
}

type SpecBox struct {
	a  *ArgAdd
	x  unit.ProductSpec
	id unit.SpecID
	ss []string
}

func (s *SpecBox) Add(arg ...unit.SpecID) *SpecBox {
	l1 := s.x.Len()
	l2 := len(arg)
	if s.x.Len() != len(arg) {
		log.Panicf("spec len %d, arg len %d. arg len == spec len", l1, l2)
	}
	tmp := make([]string, l2)
	for i := 0; i < l2; i++ {
		if !s.x.HasSub(i, arg[i]) {
			log.Panicf("arg %d values: %v", i, s.x.GetSub(i))
		}
		tmp[i] = arg[i].ToString()
	}
	result := strings.Join(tmp, "|")
	for _, ss := range s.ss {
		if ss == result {
			log.Panicf("result %s repeat", result)
		}
	}
	s.ss = append(s.ss, result)
	return s
}

func (s *SpecBox) Done() {
	tmp := *s
	tmp.a._ids = append(s.a._ids, tmp.ss...)
	*s = SpecBox{}
}
