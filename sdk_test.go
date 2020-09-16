package shop_sdk_douyin

import (
	"testing"
)

const (
	TestAppKey    = "6870386875393410567"
	TestAppSecret = "25dd8e74-216e-42cb-8012-7d3bba90d3bd"
)

var app = NewApp(TestAppKey, TestAppSecret, "3c1691ab-0b23-4656-bfc2-32b65d1d1276")

func TestShopBrandList(t *testing.T) {
	app.ProductList(ProductListArg{Page: 0, Size: 10, Status: PStatusOn, CheckStatus: PCheckPass})
}
