package shop_sdk_douyin

import (
	spec2 "github.com/cgghui/shop_sdk_douyin/product_spec/spec"
	"sync"
	"testing"
)

const (
	TestAppKey    = "6870386875393410567"
	TestAppSecret = "25dd8e74-216e-42cb-8012-7d3bba90d3bd"
)

var app = NewBaseApp(TestAppKey, TestAppSecret).NewAccessTokenMust("3c1691ab-0b23-4656-bfc2-32b65d1d1276")

func TestSpec(t *testing.T) {

	// 按商品模块取得方法集
	spec := GetProductSpec(app)

	// 获取规格列表
	list, err := spec.SpecList()
	if err != nil {
		t.Fatal(err)
	}
	wg := sync.WaitGroup{}
	for i := 0; i < len(list); i++ {
		wg.Add(1)
		go func(s spec2.ResponseSpecList) {
			defer wg.Done()
			detail, err := spec.SpecDetail(s.ID)
			if err != nil {
				t.Logf("spec %d detail error: %v\n", s.ID, err)
			} else {
				t.Logf("spec %d detail : %+v\n", s.ID, detail)
			}
		}(list[i])
	}
	wg.Wait()
	t.Logf("%+v", list)
}

// TestSpecAdd 创建商品的规格选项
func TestSpecAdd(t *testing.T) {

	// 按商品模块取得方法集
	spec := GetProductSpec(app)

	// 构建参数
	arg := spec2.SpecAddArg{Name: "规格参数一"}
	arg.NewSpecBox("容量").Add("100ml").Add("300ml").Add("500ml").Add("1000ml").Done()
	arg.NewSpecBox("颜色").Add("红色").Add("白色").Add("黄色").Add("绿色").Done()
	arg, err := arg.GetArgs()
	if err != nil {
		t.Fatal(err)
	}

	// 发起请求
	ret, err := spec.SpecAdd(arg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("【添加】%+v", ret)

	// 删除规格
	err = spec.SpecDel(ret.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("【删除】spec id %d 删除成功\n", ret.ID)
}
