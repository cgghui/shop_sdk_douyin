package shop_sdk_douyin

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type ParamMap map[string]interface{}

const GatewayURL = "https://openapi-fxg.jinritemai.com"

var SortKeyList = [5]string{"app_key", "method", "param_json", "timestamp", "v"}

type BaseApp struct {
	Key         string
	Secret      string
	AccessToken string
}

type BaseResp struct {
	Data    interface{} `json:"data"`
	ErrNo   int         `json:"err_no"`
	Message string      `json:"message"`
}

// NewRequest 执行请求
func NewRequest(app BaseApp, method string, data ParamMap) (BaseResp, error) {
	var ret BaseResp
	if data == nil {
		data = ParamMap{}
	} else {
		if len(data) > 0 {
			for k, val := range data {
				data[k] = fmt.Sprint(val)
			}
		}
	}
	params := ParamMap{
		"method":       method,
		"app_key":      app.Key,
		"access_token": app.AccessToken,
		"param_json":   data,
		"timestamp":    time.Now().Format("2006-01-02 15:04:05"),
		"v":            "2",
		"sign":         "",
	}
	params["sign"] = Sign(params, app.Secret)

	query := url.Values{}
	for k, v := range params {
		if s, ok := v.(string); ok {
			query.Add(k, s)
		}
	}
	body := strings.NewReader(query.Encode())
	req, err := http.NewRequest("POST", GatewayURL+"/"+strings.ReplaceAll(method, ".", "/"), body)
	if err != nil {
		return ret, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ret, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return ret, err
	}
	return ret, nil
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
