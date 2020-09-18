package shop_sdk_douyin

import (
	"testing"
)

const (
	TestAppKey    = "6870386875393410567"
	TestAppSecret = "25dd8e74-216e-42cb-8012-7d3bba90d3bd"
)

var app = NewBaseApp(TestAppKey, TestAppSecret).NewAccessTokenMust("3c1691ab-0b23-4656-bfc2-32b65d1d1276")

func TestShopBrandList(t *testing.T) {

}

// TestSpecAdd 创建商品的规格选项
func TestSpecAdd(t *testing.T) {

	// 按商品模块取得方法集
	product := GetProduct(app)

	// 构建参数
	arg := SpecAddArg{Name: "规格参数一"}
	arg.NewSpecBox("容量").Add("100ml").Add("300ml").Add("500ml").Add("1000ml").Done()
	arg.NewSpecBox("颜色").Add("红色").Add("白色").Add("黄色").Add("绿色").Done()
	arg, err := arg.GetArgs()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("请求参数：%+v\n\n", arg)

	// 发起请求
	ret, err := product.SpecAdd(arg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ret)
}
