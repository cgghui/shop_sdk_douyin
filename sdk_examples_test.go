package shop_sdk_douyin

import (
	"github.com/cgghui/shop_sdk_douyin/aftersale"
	"github.com/cgghui/shop_sdk_douyin/order"
	"github.com/cgghui/shop_sdk_douyin/product"
	"github.com/cgghui/shop_sdk_douyin/product/sku"
	"github.com/cgghui/shop_sdk_douyin/product/spec"
	"github.com/cgghui/shop_sdk_douyin/unit"
	"sync"
	"testing"
	"time"
)

const (
	TestAppKey    = "6870386875393410567"
	TestAppSecret = "25dd8e74-216e-42cb-8012-7d3bba90d3bd"
)

var app = NewBaseApp(TestAppKey, TestAppSecret).NewAccessTokenMust("b64614cd-03a6-48d9-9808-d160959e3f8f")
var TestGoodsID = unit.ProductID("3436456108863134126")

func TestRequest(t *testing.T) {

}

// TestExampleSpecManage 商品规格列表
func TestExampleSpecManage(t *testing.T) {

	// 按商品模块取得方法集
	productSpec := ProductSpecAPI(app)

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

// TestExampleSpecAdd 创建商品的规格选项
func TestExampleSpecAdd(t *testing.T) {

	// 按商品模块取得方法集
	productSpec := ProductSpecAPI(app)

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

// TestExampleSkuAdd 创建商品的SKU
func TestExampleSkuAdd(t *testing.T) {

	t.Logf("AccessToken: %v\n\n", app.AccessToken)

	goods, err := ProductAPI(app).ProductDetail(TestGoodsID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("【商品信息】%+v\n\n", goods)

	// 获取规格列表
	list, err := ProductSpecAPI(app).SpecList()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("【规格列表】%+v\n\n", list)

	// 必须取一组规格 76671573 为规格ID
	specObj, err := ProductSpecAPI(app).SpecDetail(unit.SpecID(76671573))
	if err != nil {
		t.Fatal(err)
	}

	// 构建参数
	arg := sku.NewArgAdd(goods.GetProductID())

	// 以下两个sku共享一个规格id
	// 如果使用其它规格，则须要再实例一个sku.NewArgAddSKU
	sss := sku.NewArgAddSKU(specObj) // sku specs
	// 组合 规格:76671573 100ml+白色 价格16元 库存18件
	sss.NewBox().SetPrice(16).SetStock(18).Push(arg, 716177728, 716177734)
	// 组合 规格:76671573 300ml+黄色 价格16元 库存18件
	sss.NewBox().SetPrice(16).SetStock(18).Push(arg, 716177728, 716177735)

	t.Logf("【规格信息】%+v\n\n", specObj)
	argObj, _ := arg.Build()
	t.Logf("【传递参数】%+v\n\n", argObj)
	t.Logf("【传递参数】%+v\n\n", ToParamMap(argObj))
	ret, err := ProductSkuAPI(app).SkuAdd(argObj)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("【执行结果】%+v\n\n", ret)
}

// TestExampleProductAdd 创建商品
func TestExampleProductAdd(t *testing.T) {

	t.Logf("AccessToken: %v\n\n", app.AccessToken)

	arg := product.NewArgAdd("iPhone Xs Max 128GB")
	var (
		argR product.ArgAddRequired = arg // 只填必传的参数
		pic  *product.Pic
	)

	// 设置价格
	argR.SetPrice(99.98, 120)

	// 设置规格 必须取一组规格 76671573 为规格ID
	SpecObj, err := app.SpecDetail(unit.SpecID(76671573))
	if err != nil {
		t.Fatal(err)
	}
	argR.SetSpecID(SpecObj)

	// 设置图片 最多5张
	pic = product.NewPic()
	pic.Add("https://gd4.alicdn.com/imgextra/i2/634491/O1CN01T9PnWc1j2vIMn6gk2_!!634491.jpg") // 主图
	pic.Add("https://gd3.alicdn.com/imgextra/i3/634491/O1CN01ucqqnb1j2vIMn6gs7_!!634491.jpg")
	pic.Add("https://gd2.alicdn.com/imgextra/i2/634491/O1CN01Fhy0h81j2vIRZYmLe_!!634491.jpg")
	pic.Add("https://gd3.alicdn.com/imgextra/i3/634491/O1CN01QeDQIe1j2vII8fMcJ_!!634491.jpg")
	pic.Add("https://gd1.alicdn.com/imgextra/i1/634491/O1CN01B8TSxS1j2vIPbeu8I_!!634491.jpg")
	if err := argR.SetPic(*pic); err != nil {
		t.Fatal(err)
	}

	// 设置描述 抖音只能用图片
	pic = product.NewPic()
	pic.Add("https://img.alicdn.com/imgextra/i2/634491/O1CN01DNccfG1j2vIROeneG_!!634491.jpg")
	pic.Add("https://img.alicdn.com/imgextra/i3/634491/O1CN01t9uyth1j2vIQ4yyVx_!!634491.jpg")
	pic.Add("https://img.alicdn.com/imgextra/i4/634491/O1CN013jOqBe1j2vIROeX41_!!634491.jpg")
	pic.Add("https://img.alicdn.com/imgextra/i2/634491/O1CN01wgPXHW1j2vIPJFLYo_!!634491.jpg")
	pic.Add("https://img.alicdn.com/imgextra/i1/634491/O1CN01LzQNUl1j2vISIuhp0_!!634491.jpg")
	argR.SetDescription(*pic)

	// 设置分类
	argR.SetCid(17, 2834, 2838) // 宠物生活/宠物主粮/猫粮

	// 设置客服
	argR.SetMobile("17715464009")

	// 设置支付方式
	argR.SetPayType(unit.CashDelivery)

	// 设置重量
	if err := argR.SetWeight(80); err != nil {
		t.Fatal(err)
	}

	// 设置属性
	pr := make(unit.PropertyOPTS, 0)
	pr.Add("年龄", "25-29周岁")
	pr.Add("尺码", "XS S M L XL")
	pr.Add("面料", "羊毛")
	pr.Add("图案", "纯色")
	pr.Add("通勤", "OL风格")
	if err := argR.SetProductFormat(pr); err != nil {
		t.Fatal(err)
	}

	// 结果
	ret, err := ProductAPI(app).ProductAdd(argR.Build())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("uint64 product id: %d create time: %s\n\n", ret.ProductID, ret.CreateTime)
}

// TestExampleProductEdit 编辑商品
func TestExampleProductEdit(t *testing.T) {
	t.Logf("AccessToken: %v\n\n", app.AccessToken)
	arg := product.NewArgEdit("3437558429391149439")
	arg.SetName("iPhone Xs Max 256GB")
	err := ProductAPI(app).ProductEdit(arg.Build())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("edit success")
}

// TestExampleProductDel 删除商品
func TestExampleProductDel(t *testing.T) {

	t.Logf("AccessToken: %v\n\n", app.AccessToken)

	ret := ProductAPI(app).ProductDel("3437555203686143344")
	t.Logf("result: %v\n", ret)
}

// TestExampleOrderList 订单列表
func TestExampleOrderList(t *testing.T) {

	t.Logf("AccessToken: %v\n\n", app.AccessToken)

	st, _ := time.Parse(unit.TimeYmd, "2020-09-20")

	arg := order.ArgList{
		//Status:    order.WaitPay,
		StartTime: st,
		EndTime:   st.Add((7 * 24) * time.Hour),
		OrderBy:   "create_time",
		IsDesc:    unit.TrueInt,
		Page:      0,
		Size:      100,
	}
	ret, err := app.OrderList(arg)
	t.Logf("result: %+v %v\n\n", ret, err)

	id := unit.OrderID("4710483426047603049A")

	detail, err := app.OrderDetail(id)
	t.Logf("detail: %+v %v\n", detail, err)

	err = app.OrderStockUp(id)
	t.Logf("stockup: %v\n", err)
}

// TestExampleOrderServiceList 订单列表
func TestExampleOrderServiceList(t *testing.T) {

	t.Logf("AccessToken: %v\n\n", app.AccessToken)

	startT, _ := time.Parse(unit.TimeYmd, "2020-09-23")

	arg := order.ArgServiceList{
		StartTime: startT,
		EndTime:   startT.Add(24 * time.Hour),
		Status:    unit.StructS{Value: "1"},
		Supply:    unit.StructS{Value: "0"},
		Page:      0,
		Size:      100,
	}
	t.Logf("Param: %v\n\n", ToParamMap(arg))

	list, err := app.OrderServiceList(arg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Result: %+v", list)
}

// TestExampleLogisticsAdd 发货
func TestExampleLogisticsAdd(t *testing.T) {

	t.Logf("AccessToken: %v\n\n", app.AccessToken)

	//
	company, err := LogisticsAPI(app).OrderLogisticsCompanyList()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("快速公司列表：%+v\n\n", company)

	arg := order.ArgLogisticsAdd{
		OrderID:       "4710483426047603049A",
		LogisticsID:   7,
		Company:       "圆通快递",
		LogisticsCode: "YT4800634233119",
	}
	t.Logf("发货结果： %+v\n", OrderAPI(app).OrderLogisticsAdd(arg))
}

// TestExampleAddressProvinceList 获取平台支持的省列表
func TestExampleAddressProvinceList(t *testing.T) {

	t.Logf("AccessToken: %v\n\n", app.AccessToken)

	//
	list, err := app.AddressProvinceList()
	if err != nil {
		t.Fatal(err)
	}

	for _, ls := range list {
		city, _ := LogisticsAPI(app).AddressCityList(ls.ID)
		for _, c := range city {
			area, err := LogisticsAPI(app).AddressAreaList(c.ID)
			t.Logf("%s %s \n%+v %v\n\n", ls.Province, c.City, area, err)
		}
	}
}

// TestExampleRefundOrderList 获取备货中有退款的订单列表
func TestExampleRefundOrderList(t *testing.T) {

	t.Logf("AccessToken: %v\n\n", app.AccessToken)

	st, _ := time.Parse(unit.TimeYmd, "2020-09-20")

	arg := aftersale.ArgRefundOrderList{
		Type:      aftersale.RFD01,
		StartTime: st,
		EndTime:   st.Add((7 * 24) * time.Hour),
		OrderBy:   "create_time",
		IsDesc:    unit.TrueInt,
		Page:      0,
		Size:      100,
	}
	list, err := app.RefundOrderList(arg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v\n\n", list)
}

// TestExampleRefundShopRefund 商家处理备货中退款申请
func TestExampleRefundShopRefund(t *testing.T) {

	t.Logf("AccessToken: %v\n\n", app.AccessToken)

	arg := aftersale.ArgRefundShopRefund{
		OrderID: "4710136396971563369A",
		Type:    aftersale.RSR01,
	}
	err := app.RefundShopRefund(arg)
	t.Logf("%v\n\n", err)
}
