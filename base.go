package shop_sdk_douyin

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ParamMap 用于包装请求数据
// 值虽然是interface，但最终在请求时都被转成了string
type ParamMap map[string]interface{}

// GatewayURL 抖音小店网关地址
const GatewayURL = "https://openapi-fxg.jinritemai.com"

// SortKeyList 公共参数排序后的字段列表，签名时用到
var SortKeyList = [5]string{
	"app_key",
	"method",
	"param_json",
	"timestamp",
	"v",
}

// BaseApp 应用的基础配置
type BaseApp struct {
	Key         string
	Secret      string
	accessToken *string
	gatewayURL  string
}

// NewBaseApp 实例化基础应用
func NewBaseApp(k, s string) *BaseApp {
	return &BaseApp{Key: k, Secret: s}
}

// SetGatewayURL 重置抖音小店网关地址
func (b *BaseApp) SetGatewayURL(u string) *BaseApp {
	b.gatewayURL = u
	return b
}

// NewAccessToken 获权AccessToken
// NewApp和NewAccessToken不是同一个对象的实例 该方法将创建新的app
// https://op.jinritemai.com/docs/guide-docs/9/21
func (b *BaseApp) NewAccessToken(t ...string) (*App, error) {
	app := App{}
	if len(t) == 0 {
		body := url.Values{}
		body.Add("app_id", b.Key)
		body.Add("app_secret", b.Secret)
		body.Add("grant_type", "authorization_self")
		resp, err := http.Get(GatewayURL + "/oauth2/access_token?" + body.Encode())
		if err != nil {
			return nil, err
		}
		var ret BaseResp
		if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
			return nil, err
		}
		if err := mapstructure.Decode(ret.Data, &app); err != nil {
			return nil, err
		}
		b.accessToken = &app.AccessToken
	} else {
		app.AccessToken = t[0]
		b.accessToken = &app.AccessToken
	}
	app.base = b
	return &app, nil
}

// NewAccessTokenMust 获权AccessToken
// 同NewAccessToken，只不过error信息存储至对象内的Error属性
// https://op.jinritemai.com/docs/guide-docs/9/21
func (b *BaseApp) NewAccessTokenMust(t ...string) *App {
	app, err := b.NewAccessToken(t...)
	if err != nil {
		return &App{Error: err}
	}
	return app
}

type BaseResp struct {
	Data    interface{} `json:"data"`
	ErrNo   int         `json:"err_no"`
	Message string      `json:"message"`
}

func toParamMap(data interface{}) ParamMap {
	r := ParamMap{}
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		x := v.Field(i)
		val := ""
		switch x.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val = strconv.FormatInt(x.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val = strconv.FormatUint(x.Uint(), 10)
		case reflect.String:
			val = x.String()
		case reflect.Bool:
			if x.Bool() {
				val = "true"
			} else {
				val = "false"
			}
		}
		if n := f.Tag.Get("paramName"); n == "" || n == "-" {
			r[strings.ToLower(f.Name)] = val
		} else {
			r[n] = val
		}
	}
	return r
}

// NewRequest 执行请求
func (b *BaseApp) NewRequest(method string, postData interface{}, d interface{}) error {
	var dat = ParamMap{}
	if postData != nil {
		if values, ok := postData.(ParamMap); ok {
			if len(values) > 0 {
				for k, v := range values {
					dat[k] = fmt.Sprint(v)
				}
			}
		} else {
			dat = toParamMap(postData)
		}
	}
	params := ParamMap{
		"method":       method,
		"app_key":      b.Key,
		"access_token": *b.accessToken,
		"param_json":   dat,
		"timestamp":    time.Now().Format("2006-01-02 15:04:05"),
		"v":            "2",
		"sign":         "",
	}
	params["sign"] = Sign(params, b.Secret)

	query := url.Values{}
	for k, v := range params {
		if s, ok := v.(string); ok {
			query.Add(k, s)
		}
	}
	body := strings.NewReader(query.Encode())
	req, err := http.NewRequest("POST", GatewayURL+"/"+strings.ReplaceAll(method, ".", "/"), body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	var ret BaseResp
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return err
	}
	if ret.ErrNo != 0 || ret.Message != "success" {
		return fmt.Errorf("response error %d %s", ret.ErrNo, ret.Message)
	}
	return mapstructure.Decode(ret.Data, d)
}

// Sign 参数签名
// 该方法会将param_json转换为json
func Sign(param ParamMap, secret string) string {
	paramJSON := param["param_json"].(ParamMap)
	if len(paramJSON) == 0 {
		param["param_json"] = "{}"
	} else {
		var ks []string
		for k := range paramJSON {
			ks = append(ks, k)
		}
		sort.Strings(ks)
		for i, k := range ks {
			ks[i] = fmt.Sprintf(`"%v":"%v"`, k, paramJSON[k])
		}
		param["param_json"] = "{" + strings.Join(ks, ",") + "}"
	}
	signStr := ""
	for _, k := range SortKeyList {
		if len(param[k].(string)) == 0 {
			continue
		}
		signStr += fmt.Sprintf("%v%v", k, param[k])
	}
	signStr = ReplaceSpecial(secret + signStr + secret)
	h := md5.New()
	h.Write([]byte(signStr))
	return hex.EncodeToString(h.Sum(nil))
}

func ReplaceSpecial(param string) string {
	param = strings.ReplaceAll(param, "&", "\\u0026")
	param = strings.ReplaceAll(param, "<", "\\u003c")
	param = strings.ReplaceAll(param, ">", "\\u00ce")
	return param
}
