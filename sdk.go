package shop_sdk_douyin

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
func (a *App) ProductList(arg ProductListArg) (ResponseProductList, error) {
	if arg.Size == 0 {
		arg.Size = 10
	}
	if arg.CheckStatus == 0 {
		arg.CheckStatus = PCheckPass
	}
	var body ResponseProductList
	if err := a.base.NewRequest("product.list", arg, &body); err != nil {
		return body, err
	}
	return body, nil
}

// ProductDetail 获取商品详细信息
// strID 字符串格式的商品id
// draft 是否从草稿加载 true从草稿加载 false不从草稿加载 默认false
// https://op.jinritemai.com/docs/api-docs/14/56
func (a *App) ProductDetail(ProductStrID string, draft ...bool) (ResponseProductDetail, error) {
	dt := "false"
	if len(draft) == 1 {
		if draft[0] {
			dt = "true"
		}
	}
	var body ResponseProductDetail
	err := a.base.NewRequest("product.detail", ParamMap{"product_id": ProductStrID, "show_draft": dt}, &body)
	if err != nil {
		return body, err
	}
	return body, nil
}

// ProductCategory 获取商品分类列表
// id 分类id，如果不指则获取最顶级
// https://op.jinritemai.com/docs/api-docs/14/58
func (a *App) ProductCategory(id ...PCid) ([]ResponseProductCategory, error) {
	cid := PCidTOP
	if len(id) == 1 {
		cid = id[0]
	}
	var body []ResponseProductCategory
	if err := a.base.NewRequest("product.getGoodsCategory", ParamMap{"cid": cid}, &body); err != nil {
		return body, err
	}
	return body, nil
}

// SkuList 获取商品sku列表
// id 分类id，如果不指则获取最顶级
// https://op.jinritemai.com/docs/api-docs/14/82
func (a *App) SkuList(ProductStrID string) ([]ResponseSkuList, error) {
	var body []ResponseSkuList
	if err := a.base.NewRequest("sku.list", ParamMap{"product_id": ProductStrID}, &body); err != nil {
		return body, err
	}
	return body, nil
}
