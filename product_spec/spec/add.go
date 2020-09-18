package spec

import (
	"errors"
	"strings"
)

// ResponseSpecAdd SpecAdd方法的响应结果
type ResponseSpecAdd struct {
	ResponseSpecDetail `mapstructure:",squash"` // 和详情返回的结果一样
}

// SpecAddArg SpecAdd方法的参数
type SpecAddArg struct {
	Name   string
	Specs  string
	_specs []SpecCreateBox // 规格盒子阵列
}

// GetArgs 取出参数集
func (s *SpecAddArg) GetArgs() (SpecAddArg, error) {
	if len(s._specs) == 0 {
		return SpecAddArg{}, errors.New("specs empty")
	}
	var tmp []string
	for _, c := range s._specs {
		tmp = append(tmp, c.join())
	}
	s.Specs = strings.Join(tmp, "^")
	s._specs = nil
	x := *s
	*s = SpecAddArg{}
	return x, nil
}

// NewSpecBox 做一个选项盒子 name 盒子的名称
func (s *SpecAddArg) NewSpecBox(name string) *SpecCreateBox {
	return &SpecCreateBox{s: s, name: name, sub: make([]string, 0)}
}

// SpecCreateBox 规格盒子
type SpecCreateBox struct {
	s    *SpecAddArg
	name string
	sub  []string
}

// Add 往盒子里添加规格 name规格的名称
func (p *SpecCreateBox) Add(name string) *SpecCreateBox {
	if p.name == "" {
		panic("name empty")
	}
	p.sub = append(p.sub, name)
	return p
}

// Done 完成 规格添加好了后，必须调用该方法 完成操作
func (p *SpecCreateBox) Done() {
	tmp := *p
	p.s._specs = append(p.s._specs, tmp)
	*p = SpecCreateBox{}
}

// join 将盒子里的数据拼合成字符串
func (p SpecCreateBox) join() string {
	return p.name + "|" + strings.Join(p.sub, ",")
}
