package shop_sdk_douyin

import (
	"github.com/cgghui/shop_sdk_douyin/product"
	"github.com/cgghui/shop_sdk_douyin/product/sku"
	"github.com/cgghui/shop_sdk_douyin/product/spec"
	"github.com/cgghui/shop_sdk_douyin/unit"
)

type ProductAPI interface {
	ProductList(product.ArgList) (product.ResponseList, error)
	ProductDetail(unit.ProductID, ...bool) (product.ResponseDetail, error)
	ProductCategory(...unit.ProductCID) ([]product.ResponseCategory, error)
	SkuList(unit.ProductID) ([]sku.ResponseList, error)
	ProductSpecAPI
}

type ProductSpecAPI interface {
	SpecAdd(spec.ArgAdd) (spec.ResponseAdd, error)
	SpecList() ([]spec.ResponseList, error)
	SpecDetail(id unit.SpecID) (spec.ResponseDetail, error)
	SpecDel(unit.SpecID) error
}

// GetProduct 从App独立出商品操作管理方法
func GetProduct(app *App) ProductAPI {
	return app
}

// GetProductSpec 从App独立出商品规格操作管理方法
func GetProductSpec(app *App) ProductSpecAPI {
	return app
}
