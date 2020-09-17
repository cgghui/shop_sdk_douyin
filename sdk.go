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
	//resp, err := a.base.NewRequest("shop.brandList", nil)
	//fmt.Printf("%+v %v", resp, err)
}

// ProductList 获取商品列表
// https://op.jinritemai.com/docs/api-docs/14/57
func (a *App) ProductList(arg ProductListArg) (ProductListResponse, error) {
	if arg.Size == 0 {
		arg.Size = 10
	}
	if arg.CheckStatus == 0 {
		arg.CheckStatus = PCheckPass
	}
	var body ProductListResponse
	err := a.base.NewRequest("product.list", arg, &body)
	if err != nil {
		return body, err
	}
	return body, nil
}

// ProductList 获取商品列表
// https://op.jinritemai.com/docs/api-docs/14/56
func (a *App) ProductDetail(arg ProductDetailArg) (ProductDetailResponse, error) {
	var body ProductDetailResponse
	err := a.base.NewRequest("product.detail", arg, &body)
	if err != nil {
		return body, err
	}
	return body, nil
}
