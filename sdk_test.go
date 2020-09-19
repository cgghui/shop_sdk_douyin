package shop_sdk_douyin

import (
	"github.com/cgghui/shop_sdk_douyin/product/sku"
	"github.com/cgghui/shop_sdk_douyin/product/spec"
	"github.com/cgghui/shop_sdk_douyin/unit"
	"sync"
	"testing"
)

const (
	TestAppKey    = "6870386875393410567"
	TestAppSecret = "25dd8e74-216e-42cb-8012-7d3bba90d3bd"
)

var app = NewBaseApp(TestAppKey, TestAppSecret).NewAccessTokenMust("b64614cd-03a6-48d9-9808-d160959e3f8f")

func TestSpecManage(t *testing.T) {

	// 按商品模块取得方法集
	productSpec := GetProductSpec(app)

	// 获取规格列表
	list, err := productSpec.SpecList()
	if err != nil {
		t.Fatal(err)
	}
	wg := sync.WaitGroup{}
	for i := 0; i < len(list); i++ {
		wg.Add(1)
		go func(s spec.ResponseList) {
			defer wg.Done()
			detail, err := productSpec.SpecDetail(s.ID)
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
	productSpec := GetProductSpec(app)

	// 构建参数
	arg := spec.NewArgAdd("规格参数一")
	arg.NewBox("颜色").Add("红色").Add("白色").Add("黄色").Add("绿色").Done()
	arg.NewBox("容量").Add("100ml").Add("300ml").Add("500ml").Add("1000ml").Done()
	arg, err := arg.Build()
	if err != nil {
		t.Fatal(err)
	}

	// 发起请求
	ret, err := productSpec.SpecAdd(arg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("【添加】%+v", ret)

	// 删除规格
	err = productSpec.SpecDel(ret.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("【删除】spec id %d 删除成功\n", ret.ID)
}

// TestSkuAdd 创建商品的SKU
func TestSkuAdd(t *testing.T) {

	t.Logf("AccessToken: %v\n\n", app.AccessToken)

	goods, err := app.ProductDetail("3436456108863134126")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("【商品信息】%+v\n\n", goods)

	// 获取规格列表
	list, err := app.SpecList()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("【规格列表】%+v\n\n", list)

	// 必须取一组规格 76671573 为规格ID
	spec, err := app.SpecDetail(unit.SpecID(76671573))
	if err != nil {
		t.Fatal(err)
	}

	// 构建参数
	arg := sku.NewArgAdd(16, 888)
	arg.SetProduct(goods.GetProductID())
	sku := arg.SetSpecID(spec)
	sku.Add(716177728, 716177734) // 组合 100ml+白色
	//sku.Add(716177728, 716177735) // 组合 300ml+黄色
	// 调用Done后sku这个对象将被重置或销毁，之后如果试图再次调用sku.Add将会引起panic
	// 一切Add工作，请在Done之前完成
	sku.Done()
	t.Logf("【规格信息】%+v\n\n", spec)
	argObj, _ := arg.Build()
	t.Logf("【传递参数】%+v\n\n", argObj)
	t.Logf("【传递参数】%+v\n\n", ToParamMap(argObj))
	ret, err := app.SkuAdd(argObj)
	t.Logf("【执行结果】%+v %v\n\n", ret, err)
}
