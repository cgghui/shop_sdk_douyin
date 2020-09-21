package sku

import (
	"errors"
	"github.com/cgghui/shop_sdk_douyin/unit"
	"log"
	"strconv"
	"strings"
)

type ResponseAdd struct {
	R interface{}
}

func (r ResponseAdd) Result() (ret map[uint64]unit.SkuID, has bool) {
	switch val := r.R.(type) {
	case []interface{}:
		{
			l := len(val)
			ret = make(map[uint64]unit.SkuID, l)
			for i := 0; i < l; i++ {
				ret[uint64(i)] = unit.SkuID(val[i].(float64))
			}
			has =
				true
		}
	case map[string]interface{}:
		{
			l := len(val)
			ret = make(map[uint64]unit.SkuID, l)
			for k, v := range val {
				id, _ := strconv.ParseUint(k, 10, 64)
				ret[id] = unit.SkuID(v.(float64))
			}
			has = true
		}
	}
	return
}

// ArgAdd ProductAdd方法的参数
type ArgAdd struct {
	ArgAddSKU
	ProductStrID unit.ProductID `paramName:"product_id"`                         // 商品id
	StockNum     string         `mapstructure:"stock_num" paramName:"stock_num"` // 库存余量
	Price        string         `mapstructure:"price"`                           // 价格
	Code         string         `mapstructure:"code" paramName:",optional"`      // 商家自定义的sku代码
	params       []ArgAddSKU    `paramName:"-"`
}

type ArgAddInterface interface {
	Build() (ArgAdd, error)
	addSku(s *ArgAddSKU)
}

func NewArgAdd(product unit.Product) ArgAddInterface {
	r := ArgAdd{
		ProductStrID: product.GetProductID(),
		params:       make([]ArgAddSKU, 0),
	}
	return &r
}

func (a *ArgAdd) addSku(s *ArgAddSKU) {
	ss := *s
	a.params = append(a.params, ss)
}

// Build 取出参数集
func (a *ArgAdd) Build() (ArgAdd, error) {
	if len(a.params) == 0 {
		return ArgAdd{}, errors.New("empty SKU")
	}
	l := len(a.params)
	tmp := [7][]string{}
	for i := 0; i < 7; i++ {
		tmp[i] = make([]string, l)
	}
	for i, param := range a.params {
		tmp[0][i] = param.SpecID
		tmp[1][i] = param.SpecDetailIDS
		tmp[2][i] = param.OutSkuID
		tmp[3][i] = param.Code
		tmp[4][i] = strconv.FormatUint(uint64(param.StockNum), 10)
		tmp[5][i] = strconv.FormatUint(uint64(param.Price), 10)
		tmp[6][i] = param.SettlementPrice
	}
	ret := *a
	*a = ArgAdd{}
	ret.ArgAddSKU.SpecID = strings.Join(tmp[0], unit.SPE1)
	ret.ArgAddSKU.SpecDetailIDS = strings.Join(tmp[1], unit.SPE2)
	ret.ArgAddSKU.OutSkuID = strings.Join(tmp[2], unit.SPE1)
	if ret.ArgAddSKU.OutSkuID == unit.SPE1 {
		ret.ArgAddSKU.OutSkuID = ""
	}
	ret.Code = strings.Join(tmp[3], unit.SPE1)
	if ret.Code == unit.SPE1 {
		ret.Code = ""
	}
	ret.StockNum = strings.Join(tmp[4], unit.SPE1)
	ret.Price = strings.Join(tmp[5], unit.SPE1)
	ret.ArgAddSKU.SettlementPrice = strings.Join(tmp[6], unit.SPE1)
	if ret.ArgAddSKU.SettlementPrice == unit.SPE1 {
		ret.ArgAddSKU.SettlementPrice = ""
	}
	return ret, nil
}

/////////////////////////////////

type ArgAddBuild struct {
	ArgAddSKU
}

type ArgAddBuildInterface interface {
	Box() ArgAddSKUInterface
}

func NewArgAddSKU(spec unit.ProductSpec) ArgAddBuildInterface {
	return &ArgAddBuild{
		ArgAddSKU: ArgAddSKU{
			spec:   spec,
			SpecID: spec.GetProductSpecID().ToString(),
		},
	}
}

func (s *ArgAddBuild) Box() ArgAddSKUInterface {
	ss := *s
	return &ss.ArgAddSKU
}

///////////////

type ArgAddSKU struct {
	Info            `paramName:"-"`
	OutSkuID        string           `paramName:"out_sku_id,optional"`       // 业务方自己的sku_id，唯一需为数字字符串，max = int64
	SpecID          string           `paramName:"spec_id"`                   // 规格id，依赖/spec/list接口的返回
	SpecDetailIDS   string           `paramName:"spec_detail_ids"`           // 子规格id,最多3级,如 100041|150041|160041（ 女款|白色|XL）
	SettlementPrice string           `paramName:"settlement_price,optional"` // 结算价格 (单位 分)
	spec            unit.ProductSpec `paramName:"-"`
}

type ArgAddSKUInterface interface {
	SetStock(uint16) *ArgAddSKU
	SetPrice(float64) *ArgAddSKU
	SetOutSkuID(uint64) *ArgAddSKU
	SetCode(string) *ArgAddSKU
	Push(ArgAddInterface, ...unit.SpecID)
}

func (s *ArgAddSKU) SetStock(n uint16) *ArgAddSKU {
	s.StockNum = n
	return s
}

func (s *ArgAddSKU) SetPrice(n float64) *ArgAddSKU {
	s.Price = unit.PriceToYuan(n)
	return s
}

func (s *ArgAddSKU) SetOutSkuID(id uint64) *ArgAddSKU {
	s.OutSkuID = strconv.FormatUint(id, 10)
	return s
}

func (s *ArgAddSKU) SetCode(c string) *ArgAddSKU {
	s.Code = c
	return s
}

func (s *ArgAddSKU) Push(box ArgAddInterface, arg ...unit.SpecID) {
	l1 := s.spec.Len()
	l2 := len(arg)
	if s.spec.Len() != len(arg) {
		log.Panicf("spec len %d, arg len %d. arg len == spec len", l1, l2)
	}
	tmp := make([]string, l2)
	for i := 0; i < l2; i++ {
		if !s.spec.HasSub(i, arg[i]) {
			log.Panicf("arg %d values: %v", i, s.spec.GetSub(i))
		}
		tmp[i] = arg[i].ToString()
	}
	s.SpecDetailIDS = strings.Join(tmp, unit.SPE1)
	box.addSku(s)
}
