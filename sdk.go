package shop_sdk_douyin

import (
	"github.com/cgghui/shop_sdk_douyin/product"
	"github.com/cgghui/shop_sdk_douyin/product/sku"
	"github.com/cgghui/shop_sdk_douyin/product/spec"
	"github.com/cgghui/shop_sdk_douyin/unit"
)

type App struct {
	base         *BaseApp
	AccessToken  string `mapstructure:"access_token"`
	ExpiresIn    uint32 `mapstructure:"expires_in"`
	RefreshToken string `mapstructure:"refresh_token"`
	Scope        string `mapstructure:"scope"`
	ShopID       uint64 `mapstructure:"shop_id"`
	ShopName     string `mapstructure:"shop_name"`
	Error        error  `mapstructure:"-"`
}

// ShopBrandList 获权店铺列表
// https://op.jinritemai.com/docs/api-docs/13/54
func (a *App) ShopBrandList() {
}

// ProductList 获取商品列表
// https://op.jinritemai.com/docs/api-docs/14/57
func (a *App) ProductList(arg product.ArgList) (product.ResponseList, error) {
	if arg.Size == 0 {
		arg.Size = 10
	}
	if arg.CheckStatus == 0 {
		arg.CheckStatus = product.CheckPass
	}
	var body product.ResponseList
	if err := a.base.NewRequest("product.list", arg, &body); err != nil {
		return body, err
	}
	return body, nil
}

// ProductDetail 获取商品详细信息
// strID 字符串格式的商品id
// draft 是否从草稿加载 true从草稿加载 false不从草稿加载 默认false
// https://op.jinritemai.com/docs/api-docs/14/56
func (a *App) ProductDetail(ProductStrID unit.ProductID, draft ...bool) (product.ResponseDetail, error) {
	dt := "false"
	if len(draft) == 1 {
		if draft[0] {
			dt = "true"
		}
	}
	var body product.ResponseDetail
	err := a.base.NewRequest("product.detail", ParamMap{"product_id": ProductStrID, "show_draft": dt}, &body)
	if err != nil {
		return body, err
	}
	return body, nil
}

// ProductCategory 获取商品分类列表
// id 分类id，如果不指则获取最顶级
// https://op.jinritemai.com/docs/api-docs/14/58
func (a *App) ProductCategory(id ...unit.ProductCID) ([]product.ResponseCategory, error) {
	cid := unit.CidTOP
	if len(id) == 1 {
		cid = id[0]
	}
	var body []product.ResponseCategory
	if err := a.base.NewRequest("product.getGoodsCategory", ParamMap{"cid": cid}, &body); err != nil {
		return body, err
	}
	return body, nil
}

// ProductAdd 添加商品
// https://op.jinritemai.com/docs/api-docs/14/59
func (a *App) ProductAdd(arg product.ArgAdd) {

}

// SpecAdd 添加选项规格
// https://op.jinritemai.com/docs/api-docs/14/64
func (a *App) SpecAdd(arg spec.ArgAdd) (spec.ResponseAdd, error) {
	var body spec.ResponseAdd
	if err := a.base.NewRequest("spec.add", arg, &body); err != nil {
		return body, err
	}
	return body, nil
}

// SpecList 获取选项规格列表
// https://op.jinritemai.com/docs/api-docs/14/64
func (a *App) SpecList() ([]spec.ResponseList, error) {
	var body []spec.ResponseList
	if err := a.base.NewRequest("spec.list", nil, &body); err != nil {
		return body, err
	}
	return body, nil
}

// SpecDetail 获取选项规格详情
// https://op.jinritemai.com/docs/api-docs/14/63
func (a *App) SpecDetail(id unit.SpecID) (spec.ResponseDetail, error) {
	var body spec.ResponseDetail
	if err := a.base.NewRequest("spec.specDetail", ParamMap{"id": id}, &body); err != nil {
		return body, err
	}
	return body, nil
}

// SpecDel 删除选项规格
// https://op.jinritemai.com/docs/api-docs/14/65
func (a *App) SpecDel(id unit.SpecID) error {
	return a.base.NewRequest("spec.del", ParamMap{"id": id}, nil)
}

// SkuAdd 添加SKU
// https://op.jinritemai.com/docs/api-docs/14/64
func (a *App) SkuAdd(arg sku.ArgAdd) (sku.ResponseAdd, error) {
	var body interface{}
	if err := a.base.NewRequest("sku.addAll", arg, &body); err != nil {
		return sku.ResponseAdd{}, err
	}
	return sku.ResponseAdd{R: body}, nil
}

// SkuList 获取商品sku列表
// id 分类id，如果不指则获取最顶级
// https://op.jinritemai.com/docs/api-docs/14/82
func (a *App) SkuList(ProductStrID unit.ProductID) ([]sku.ResponseList, error) {
	var body []sku.ResponseList
	if err := a.base.NewRequest("sku.list", ParamMap{"product_id": ProductStrID}, &body); err != nil {
		return body, err
	}
	return body, nil
}
