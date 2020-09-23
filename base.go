package shop_sdk_douyin

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cgghui/shop_sdk_douyin/unit"
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
	return &BaseApp{Key: k, Secret: s, gatewayURL: GatewayURL}
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

// ToParamMap 将任意struct转换为成ParamMap
// paramName "-" 这个字段将被忽略，须要注意的是如果字段是bool，那么将被转换成字符串
// 类型的"true"和"false"
func ToParamMap(data interface{}, ret ...*ParamMap) ParamMap {
	var (
		r ParamMap // 最终结果
		t reflect.Type
		v reflect.Value
	)
	if len(ret) == 0 {
		r = ParamMap{} // 非递归
	} else {
		r = *ret[0] // 递归时 将以指针式进行赋值
	}
	if val, ok := data.(reflect.Value); ok {
		// 递归
		t = val.Type()
		v = val
	} else {
		// 非递归
		t = reflect.TypeOf(data)
		v = reflect.ValueOf(data)
	}
	// HookConvertValue
	var Func1has bool
	var Func1 reflect.Value
	if _, has := t.MethodByName("HookConvertValue"); has {
		Func1has = has
		Func1 = v.MethodByName("HookConvertValue")
	}
	// HookSkipCheck
	var Func2has bool
	var Func2 reflect.Value
	if _, has := t.MethodByName("HookSkipCheck"); has {
		Func2has = has
		Func2 = v.MethodByName("HookSkipCheck")
	}
	// 遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)       // 字段
		x := v.Field(i)       // 值
		fs := f.Type.String() // 字段类型 字符串
		ff := reflect.ValueOf(f)
		// tag 结构体后的标记 n标记名称 o标记参数
		tag := f.Tag.Get("paramName")
		if tag == "-" {
			continue
		}
		n := ""
		o := ""
		if strings.Index(tag, unit.SPE3) == -1 {
			n = tag
		} else {
			xx := strings.Split(tag, unit.SPE3)
			n = xx[0]
			o = xx[1]
		}
		if n == "" {
			n = strings.ToLower(f.Name)
		}
		val := ""
		switch x.Kind() {
		// int
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val = strconv.FormatInt(x.Int(), 10)
		// uint
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val = strconv.FormatUint(x.Uint(), 10)
		// float
		case reflect.Float32, reflect.Float64:
			val = strconv.FormatFloat(x.Float(), 'f', -1, 64)
		// string
		case reflect.String:
			val = x.String()
		// struct
		case reflect.Struct:
			if f.Name == f.Type.Name() {
				ToParamMap(x, &r)
			} else {
				if Func1has {
					val = Func1.Call([]reflect.Value{ff, reflect.ValueOf(x)})[0].String()
				}
			}
		// bool
		case reflect.Bool:
			if x.Bool() {
				val = "true"
			} else {
				val = "false"
			}
		}
		if x.Kind() != reflect.Struct {
			if o == "optional" {
				if val != "" && val != "0" {
					r[n] = val
				}
			} else {
				if Func2has {
					arg := []reflect.Value{reflect.ValueOf(fs), reflect.ValueOf(n), reflect.ValueOf(val)}
					if ret := Func2.Call(arg); !ret[0].Bool() {
						r[n] = val
					}
				} else {
					r[n] = val
				}
			}
		} else {
			if Func2has {
				arg := []reflect.Value{reflect.ValueOf(fs), reflect.ValueOf(n), reflect.ValueOf(val)}
				if ret := Func2.Call(arg); !ret[0].Bool() {
					r[n] = val
				}
			}
		}
	}
	// 递归返回
	if len(ret) != 0 {
		*ret[0] = r
		return nil
	}
	// 最终返回
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
			dat = ToParamMap(postData)
		}
	}
	params := ParamMap{
		"method":       method,
		"app_key":      b.Key,
		"access_token": *b.accessToken,
		"param_json":   dat,
		"timestamp":    time.Now().Format(unit.TimeYmdHis),
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
	req, err := http.NewRequest("POST", b.gatewayURL+"/"+strings.ReplaceAll(method, ".", "/"), body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	var ret BaseResp
	//r, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("%s\n\n", r)
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return err
	}
	if ret.ErrNo != 0 || ret.Message != "success" {
		return fmt.Errorf("response error %d %s", ret.ErrNo, ret.Message)
	}
	if d == nil {
		return nil
	}
	if ret.Data == nil {
		return errors.New("response error data is nil")
	}
	if reflect.TypeOf(d).Elem().Kind() == reflect.Interface {
		rd := reflect.ValueOf(ret.Data)
		reflect.ValueOf(d).Elem().Set(rd)
		return nil
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
		param["param_json"] = "{" + strings.Join(ks, unit.SPE3) + "}"
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
