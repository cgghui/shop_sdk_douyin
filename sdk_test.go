package shop_sdk_douyin

import (
	"fmt"
	"testing"
	"time"
)

const (
	TestAppKey    = "6870386875393410567"
	TestAppSecret = "25dd8e74-216e-42cb-8012-7d3bba90d3bd"
)

var app = NewBaseApp(TestAppKey, TestAppSecret)

func TestShopBrandList(t *testing.T) {
	time.Now()
	appx := app.NewAccessTokenMust("3c1691ab-0b23-4656-bfc2-32b65d1d1276")
	product := GetProduct(appx)
	//l, e := appx.ProductList(ProductListArg{
	//	Page:        0,
	//	Size:        10,
	//	Status:      PStatusOn,
	//	CheckStatus: PCheckPass,
	//})
	//fmt.Printf("%+v %v\n\n", l, e)
	//d, e := appx.ProductDetail(ProductDetailArg{ProductStrID: "3436456108863134126", ShowDraft: true})
	//fmt.Printf("%+v %v\n\n", d, e)
	//fmt.Printf("%+v\n\n", d.SpecPics)
	//fmt.Printf("%+v\n\n", d.SpecPrices)
	//fmt.Printf("%+v\n\n", d.Specs[0])
	//c, e := product.ProductCategory(PCid(28))
	//fmt.Printf("%+v %v\n\n", c, e)
	s, _ := product.SkuList("3436456108863134126")

	fmt.Printf("%+v\n\n", s)
}
