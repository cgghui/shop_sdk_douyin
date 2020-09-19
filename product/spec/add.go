package spec

import (
	"errors"
	"github.com/cgghui/shop_sdk_douyin/unit"
	"strings"
)

// ResponseAdd SpecAdd方法的响应结果
type ResponseAdd struct {
	ResponseDetail `mapstructure:",squash"` // 和详情返回的结果一样
}

// ArgAdd SpecAdd方法的参数
type ArgAdd struct {
	Name   string
	Specs  string
	_specs []CreateBox `paramName:"-"` // 规格盒子阵列
}

func NewArgAdd(name string) ArgAdd {
	return ArgAdd{Name: name}
}

// Build 取出参数集
func (s *ArgAdd) Build() (ArgAdd, error) {
	if len(s._specs) == 0 {
		return ArgAdd{}, errors.New("specs empty")
	}
	var tmp []string
	for _, c := range s._specs {
		tmp = append(tmp, c.join())
	}
	s.Specs = strings.Join(tmp, unit.SPE2)
	s._specs = nil
	x := *s
	*s = ArgAdd{}
	return x, nil
}

// NewBox 做一个选项盒子 name 盒子的名称
func (s *ArgAdd) NewBox(name string) *CreateBox {
	return &CreateBox{s: s, name: name, sub: make([]string, 0)}
}

// CreateBox 规格盒子
type CreateBox struct {
	s    *ArgAdd
	name string
	sub  []string
}

// Add 往盒子里添加规格 name规格的名称
func (p *CreateBox) Add(name string) *CreateBox {
	if p.name == "" {
		panic("name empty")
	}
	p.sub = append(p.sub, name)
	return p
}

// Done 完成 规格添加好了后，必须调用该方法 完成操作
func (p *CreateBox) Done() {
	tmp := *p
	p.s._specs = append(p.s._specs, tmp)
	*p = CreateBox{}
}

// join 将盒子里的数据拼合成字符串
func (p CreateBox) join() string {
	return p.name + unit.SPE1 + strings.Join(p.sub, unit.SPE3)
}
