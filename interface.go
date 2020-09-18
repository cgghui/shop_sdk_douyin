package shop_sdk_douyin

import (
	"github.com/cgghui/shop_sdk_douyin/product_spec/sku"
	"github.com/cgghui/shop_sdk_douyin/product_spec/spec"
)

type ProductInterface interface {
	ProductList(ProductListArg) (ResponseProductList, error)
	ProductDetail(string, ...bool) (ResponseProductDetail, error)
	ProductCategory(...PCid) ([]ResponseProductCategory, error)
	SkuList(string) ([]sku.ResponseSkuList, error)
	ProductSpecInterface
}

type ProductSpecInterface interface {
	SpecAdd(spec.SpecAddArg) (spec.ResponseSpecAdd, error)
	SpecList() ([]spec.ResponseSpecList, error)
	SpecDetail(spec.SpecID) (spec.ResponseSpecDetail, error)
	SpecDel(spec.SpecID) error
}

// GetProduct 从App独立出商品操作管理方法
func GetProduct(app *App) ProductInterface {
	return app
}

// GetProductSpec 从App独立出商品规格操作管理方法
func GetProductSpec(app *App) ProductSpecInterface {
	return app
}
