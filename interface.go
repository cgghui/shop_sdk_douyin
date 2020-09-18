package shop_sdk_douyin

import (
	"github.com/cgghui/shop_sdk_douyin/product"
	"github.com/cgghui/shop_sdk_douyin/product/sku"
	"github.com/cgghui/shop_sdk_douyin/product/spec"
	"github.com/cgghui/shop_sdk_douyin/unit"
)

type ProductInterface interface {
	ProductList(product.ArgList) (product.ResponseList, error)
	ProductDetail(unit.ProductID, ...bool) (product.ResponseDetail, error)
	ProductCategory(...unit.ProductCID) ([]product.ResponseCategory, error)
	SkuList(unit.ProductID) ([]sku.ResponseList, error)
	ProductSpecInterface
}

type ProductSpecInterface interface {
	SpecAdd(spec.ArgAdd) (spec.ResponseAdd, error)
	SpecList() ([]spec.ResponseList, error)
	SpecDetail(id unit.SpecID) (spec.ResponseDetail, error)
	SpecDel(unit.SpecID) error
}

// GetProduct 从App独立出商品操作管理方法
func GetProduct(app *App) ProductInterface {
	return app
}

// GetProductSpec 从App独立出商品规格操作管理方法
func GetProductSpec(app *App) ProductSpecInterface {
	return app
}
