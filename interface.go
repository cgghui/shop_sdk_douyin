package shop_sdk_douyin

type ProductInterface interface {
	ProductList(ProductListArg) (ResponseProductList, error)
	ProductDetail(string, ...bool) (ResponseProductDetail, error)
	ProductCategory(...PCid) ([]ResponseProductCategory, error)
	SpecAdd(SpecAddArg) (ResponseSpecAdd, error)
	SpecList() ([]ResponseSpecList, error)
	SpecDetail(SpecID) (ResponseSpecDetail, error)
	SkuList(string) ([]ResponseSkuList, error)
}

// GetProduct 从App独立出商品操作管理方法
func GetProduct(app *App) ProductInterface {
	return app
}
