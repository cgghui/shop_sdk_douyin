package shop_sdk_douyin

import (
	"errors"
	"github.com/cgghui/shop_sdk_douyin/logistics"
	"github.com/cgghui/shop_sdk_douyin/order"
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

// ShopBrandList 获取店铺的已授权品牌列表
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

// ProductCateProperty 根据商品分类获取对应的属性列表
// id 分类id，如果不指则获取最顶级
// https://op.jinritemai.com/docs/api-docs/14/58
func (a *App) ProductCateProperty(cid1, cid2, cid3 unit.ProductCID) ([]product.ResponseCateProperty, error) {
	dat := ParamMap{"first_cid": cid1, "second_cid": cid2, "third_cid": cid3}
	var body []product.ResponseCateProperty
	if err := a.base.NewRequest("product.getCateProperty", dat, &body); err != nil {
		return body, err
	}
	return body, nil
}

// ProductAdd 添加商品
// https://op.jinritemai.com/docs/api-docs/14/59
func (a *App) ProductAdd(arg product.ArgAdd) (product.Product, error) {
	var body product.ResponseDetail
	if err := a.base.NewRequest("product.add", arg, &body); err != nil {
		return product.Product{}, err
	}
	return body.Product, nil
}

// ProductEdit 编辑商品
// 编辑商品的参数虽与ProductAdd共用，但须要使用product.NewArgEdit方法进行实例
// https://op.jinritemai.com/docs/api-docs/14/60
func (a *App) ProductEdit(arg product.ArgAdd) error {
	var body interface{}
	if err := a.base.NewRequest("product.edit", arg, &body); err != nil {
		return err
	}
	if ret, ok := body.(bool); ok && ret {
		return nil
	}
	return errors.New("edit fail")
}

// ProductDel 删除商品
// https://op.jinritemai.com/docs/api-docs/14/61
func (a *App) ProductDel(id unit.ProductID) error {
	if err := a.base.NewRequest("product.del", ParamMap{"product_id": id}, nil); err != nil {
		return err
	}
	return nil
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
func (a *App) SkuAdd(arg sku.ArgAdd) (map[uint64]unit.SkuID, error) {
	var body interface{}
	if err := a.base.NewRequest("sku.addAll", arg, &body); err != nil {
		return nil, err
	}
	ret, ok := sku.ResponseAdd{R: body}.Result()
	if ok {
		return ret, nil
	}
	return nil, errors.New("response data unable to parse")
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

// SkuDetail 获取商品sku详情
// https://op.jinritemai.com/docs/api-docs/14/104
// todo 官方的这个方法有问题，404。可以通过SkuList方法暂时代替
func (a *App) SkuDetail(sid unit.SkuID) (sku.ResponseDetail, error) {
	var body sku.ResponseDetail
	if err := a.base.NewRequest("sku.detail", ParamMap{"sku_id": sid}, &body); err != nil {
		return body, err
	}
	return body, nil
}

// SkuEditPrice 修改商品sku的价格
// op参数的实现 sku.Detail 可调用 App.SkuDetail、App.SkuList 方法
// https://op.jinritemai.com/docs/api-docs/14/84
func (a *App) SkuEditPrice(op unit.SkuOperate, p float64) error {
	arg := ParamMap{"product_id": op.GetProductID(), "sku_id": op.GetSkuID(), "price": unit.PriceToYuan(p)}
	return a.base.NewRequest("sku.editPrice", arg, nil)
}

// SkuSyncStock 修改商品sku的库存
// op参数的实现 sku.Detail 可调用 App.SkuDetail、App.SkuList 方法
// https://op.jinritemai.com/docs/api-docs/14/85
func (a *App) SkuSyncStock(op unit.SkuOperate, n uint16) error {
	arg := ParamMap{"product_id": op.GetProductID(), "sku_id": op.GetSkuID(), "stock_num": n}
	return a.base.NewRequest("sku.syncStock", arg, nil)
}

// SkuEditCode 修改商品sku的编码
// op参数的实现 sku.Detail 可调用 App.SkuDetail、App.SkuList 方法
// https://op.jinritemai.com/docs/api-docs/14/86
func (a *App) SkuEditCode(op unit.SkuOperate, c string) error {
	arg := ParamMap{"product_id": op.GetProductID(), "sku_id": op.GetSkuID(), "code": c}
	return a.base.NewRequest("sku.editCode", arg, nil)
}

// OrderList 订单列表
// https://op.jinritemai.com/docs/api-docs/15/55
func (a *App) OrderList(arg order.ArgList) (order.ResponseList, error) {
	var body order.ResponseList
	if err := a.base.NewRequest("order.list", arg, &body); err != nil {
		return order.ResponseList{}, err
	}
	return body, nil
}

// OrderDetail 订单详情
// https://op.jinritemai.com/docs/api-docs/15/68
func (a *App) OrderDetail(o unit.Order) (order.Detail, error) {
	var body order.ResponseList
	if err := a.base.NewRequest("order.detail", ParamMap{"order_id": o.GetParentID()}, &body); err != nil {
		return order.Detail{}, err
	}
	if body.Total != 1 {
		return order.Detail{}, errors.New("order total not is 1")
	}
	return body.List[0], nil
}

// OrderStockUp 确认货到付款订单
// https://op.jinritemai.com/docs/api-docs/15/69
func (a *App) OrderStockUp(o unit.Order) error {
	return a.base.NewRequest("order.stockUp", ParamMap{"order_id": o.GetParentID()}, nil)
}

// OrderCancel 取消货到付款订单
// https://op.jinritemai.com/docs/api-docs/15/72
func (a *App) OrderCancel(o unit.Order, reason string) error {
	return a.base.NewRequest("order.cancel", ParamMap{"order_id": o.GetParentID(), "reason": reason}, nil)
}

// OrderServiceList 获取客服向店铺发起的服务请求列表
// https://op.jinritemai.com/docs/api-docs/15/74
func (a *App) OrderServiceList(arg order.ArgServiceList) (order.ResponseServiceList, error) {
	var body order.ResponseServiceList
	if err := a.base.NewRequest("order.serviceList", arg, &body); err != nil {
		return order.ResponseServiceList{}, err
	}
	return body, nil
}

// OrderReplyService 回复服务请求
// https://op.jinritemai.com/docs/api-docs/15/75
func (a *App) OrderReplyService(id unit.ServiceID, reply string) error {
	return a.base.NewRequest("order.replyService", ParamMap{"id": id, "reply": reply}, nil)
}

// OrderLogisticsCompanyList 回复服务请求
// https://op.jinritemai.com/docs/api-docs/16/76
func (a *App) OrderLogisticsCompanyList() (logistics.ResponseLogisticsCompanyList, error) {
	var body logistics.ResponseLogisticsCompanyList
	if err := a.base.NewRequest("order.logisticsCompanyList", nil, &body); err != nil {
		return body, err
	}
	return body, nil
}

// OrderLogisticsAdd 订单发货
// https://op.jinritemai.com/docs/api-docs/16/77
func (a *App) OrderLogisticsAdd(arg order.ArgLogisticsAdd) error {
	return a.base.NewRequest("order.logisticsAdd", arg, nil)
}

// OrderLogisticsEdit 修改发货物流
// https://op.jinritemai.com/docs/api-docs/16/79
func (a *App) OrderLogisticsEdit(arg order.ArgLogisticsAdd) error {
	return a.base.NewRequest("order.logisticsEdit", arg, nil)
}

// AddressProvinceList 获取省列表
// https://op.jinritemai.com/docs/api-docs/16/101
func (a *App) AddressProvinceList() ([]logistics.Province, error) {
	var body []logistics.Province
	if err := a.base.NewRequest("address.provinceList", nil, &body); err != nil {
		return nil, err
	}
	return body, nil
}

// AddressCityList 获取市列表
// https://op.jinritemai.com/docs/api-docs/16/102
func (a *App) AddressCityList(provID uint32) ([]logistics.City, error) {
	var body []logistics.City
	if err := a.base.NewRequest("address.cityList", ParamMap{"province_id": provID}, &body); err != nil {
		return nil, err
	}
	return body, nil
}

// AddressCityList 获取市列表
// https://op.jinritemai.com/docs/api-docs/16/103
func (a *App) AddressAreaList(cityID uint32) ([]logistics.Area, error) {
	var body []logistics.Area
	if err := a.base.NewRequest("address.areaList", ParamMap{"city_id": cityID}, &body); err != nil {
		return nil, err
	}
	return body, nil
}
