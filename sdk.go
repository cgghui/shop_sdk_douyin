package shop_sdk_douyin

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"net/url"
)

type App struct {
	BaseApp
	AccessToken  string `mapstructure:"access_token"`
	ExpiresIn    uint32 `mapstructure:"expires_in"`
	RefreshToken string `mapstructure:"refresh_token"`
	Scope        string `mapstructure:"scope"`
	ShopID       uint64 `mapstructure:"shop_id"`
	ShopName     string `mapstructure:"shop_name"`
	Error        error  `mapstructure:"-"`
}

func NewApp(k, s, t string) *App {
	return &App{
		BaseApp: BaseApp{Key: k, Secret: s, AccessToken: t},
	}
}

// NewAccessToken 获权AccessToken
// NewApp和NewAccessToken不是同一个对象的实例 该方法将创建新的app
// https://op.jinritemai.com/docs/guide-docs/9/21
func (a *App) NewAccessToken() (*App, error) {
	body := url.Values{}
	body.Add("app_id", a.Key)
	body.Add("app_secret", a.Secret)
	body.Add("grant_type", "authorization_self")
	resp, err := http.Get(GatewayURL + "/oauth2/access_token?" + body.Encode())
	if err != nil {
		return nil, err
	}
	var ret BaseResp
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}
	var app App
	if err := mapstructure.Decode(ret.Data, &app); err != nil {
		return nil, err
	}
	app.BaseApp = a.BaseApp
	app.BaseApp.AccessToken = app.AccessToken
	return &app, nil
}

// NewAccessTokenMust 获权AccessToken
// 同NewAccessToken，只不过error信息存储至对象内的Error属性
// https://op.jinritemai.com/docs/guide-docs/9/21
func (a *App) NewAccessTokenMust() *App {
	app, err := a.NewAccessToken()
	if err != nil {
		a.Error = err
	}
	return app
}

// ShopBrandList 获权店铺列表
// https://op.jinritemai.com/docs/api-docs/13/54
func (a *App) ShopBrandList() {
	resp, err := NewRequest(a.BaseApp, "shop.brandList", nil)
	fmt.Printf("%+v %v", resp, err)
}

type (
	PStatus        uint8
	PCheck         uint8
	ProductListArg struct {
		Page        uint8
		Size        uint8
		Status      PStatus
		CheckStatus PCheck
	}
)

const (
	PStatusOn PStatus = iota
	PStatusOff
)

const (
	PCheckNot PCheck = iota + 1 // 未提审
	PCheckIng
	PCheckPass
	PCheckReject
	PCheckForbid
)

// ProductList 获取商品列表
// https://op.jinritemai.com/docs/api-docs/14/57
// page 第几页（第一页为0）
func (a *App) ProductList(arg ProductListArg) {
	if arg.Size == 0 {
		arg.Size = 10
	}
	if arg.CheckStatus == 0 {
		arg.CheckStatus = PCheckPass
	}
	resp, err := NewRequest(a.BaseApp, "product.list", ParamMap{
		"page":         arg.Page,
		"size":         arg.Size,
		"status":       arg.Status,
		"check_status": arg.CheckStatus,
	})
	fmt.Printf("%+v %v", resp, err)
}
