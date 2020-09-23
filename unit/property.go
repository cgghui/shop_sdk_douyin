package unit

// Property 属性项
type Property struct {
	Name  string `mapstructure:"name"`
	Value string `mapstructure:"value"`
}

// PropertyOPTS 多项属性
type PropertyOPTS []Property

// Add 新增属性
func (p *PropertyOPTS) Add(n, v string) {
	*p = append(*p, Property{Name: n, Value: v})
}

// Relation 通用关系结构
type Relation struct {
	ID   string `mapstructure:"id"`
	Name string `mapstructure:"name"`
}
