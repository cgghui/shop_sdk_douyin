package product

import (
	"errors"
	"fmt"
	"github.com/cgghui/shop_sdk_douyin/product/spec"
	"github.com/cgghui/shop_sdk_douyin/unit"
	"strconv"
	"strings"
	"time"
)

const (
	DelayDay02 uint8 = 2
	DelayDay03 uint8 = 3
	DelayDay05 uint8 = 5
	DelayDay07 uint8 = 7
	DelayDay10 uint8 = 10
	DelayDay15 uint8 = 15
)

// ArgAdd ProductAdd方法的参数
// 不建议直接使用这个结构体，因为体内的字段全部是string，无法预知其格式
// 请通过NewArgAdd方法来进行添加参数
type ArgAdd struct {
	ProductID        string `paramName:"product_id,optional"`
	OutProductID     string `paramName:"out_product_id,optional"`     // 外部商品id，接入方的商品id需为数字字符串 int64
	Name             string `paramName:"name"`                        // 商品名称，最多30个字符，不能含emoj表情
	Pic              string `paramName:"pic"`                         // 商品轮播图
	Description      string `paramName:"description"`                 // 商品描述
	MarketPrice      string `paramName:"market_price"`                // 划线价，单位分
	DiscountPrice    string `paramName:"discount_price"`              // 售价，单位分
	FirstCid         string `paramName:"first_cid"`                   // 一级分类id（三个分类级别请确保从属正确）
	SecondCid        string `paramName:"second_cid"`                  // 二级分类id
	ThirdCid         string `paramName:"third_cid"`                   // 三级分类id
	SpecID           string `paramName:"spec_id"`                     // 规格id, 要先创建商品通用规格, 如颜色-尺寸
	SpecPic          string `paramName:"spec_pic,optional"`           // 主规格id, 如颜色-尺寸 id|图片url
	Mobile           string `paramName:"mobile"`                      // 客服号码
	Weight           string `paramName:"weight"`                      // 商品重量 (单位:克)。范围: 10克 - 9999990克
	ProductFormat    string `paramName:"product_format"`              // 属性名称|属性值之间用|分隔, 多组之间用^分开
	PayType          string `paramName:"pay_type"`                    // 支付方式
	RecommendRemark  string `paramName:"recommend_remark,optional"`   // 商家推荐语，不能含emoj表情
	BrandID          string `paramName:"brand_id,optional"`           // 品牌id
	PresellType      string `paramName:"presell_type,optional"`       // 预售类型，1-全款预售，0-非预售，默认0
	PresellDelay     string `paramName:"presell_delay,optional"`      // 预售结束后，几天发货，可以选择2-30
	PresellEndTime   string `paramName:"presell_end_time,optional"`   // 预售结束时间，格式2020-02-21 18:54:27，最大30天
	DeliveryDelayDay string `paramName:"delivery_delay_day,optional"` // 承诺发货时间，单位是天 2|3|5|7|10|15
	QualityReport    string `paramName:"quality_report,optional"`     // 质检报告链接,多个图片以逗号分隔
	ClassQuality     string `paramName:"class_quality,optional"`      // 品类资质链接,多个图片以逗号分隔
	Commit           string `paramName:"commit,optional"`             // 编辑时使用
}

// ArgAddInterface 所有参数方法
type ArgAddInterface interface {
	ArgAddRequired
	ArgAddOptional
}

// ArgAddRequired 必填参数方法
type ArgAddRequired interface {
	SetPic(Pic) error
	SetDescription(Pic)
	SetPrice(p1, p2 float64)
	SetCid(c1, c2, c3 unit.ProductCID)
	SetSpecID(unit.ProductSpec)
	SetMobile(string)
	SetWeight(uint32) error
	SetProductFormat(unit.PropertyOPTS) error
	SetPayType(unit.PayT)
	ArgBuild
}

// ArgAddOptional 选填参数方法
type ArgAddOptional interface {
	SetOutProductID(uint64)
	SetSpecPic([]spec.Pic) error
	SetRecommendRemark(string)
	SetBrandID(unit.BrandID)
	SetPresellType(unit.GoodsP)
	SetPresellDelay(uint8) error
	SetPresellEndTime(time.Duration) error
	SetDeliveryDelayDay(uint8)
	SetQualityReport(Pic)
	SetClassQuality(Pic)
	ArgBuild
}

// ArgEdit 编辑参数方法
type ArgEdit interface {
	SetName(string)
	SetPic(Pic) error
	SetDescription(Pic)
	SetCid(c1, c2, c3 unit.ProductCID)
	SetSpecID(unit.ProductSpec)
	SetMobile(string)
	SetProductFormat(unit.PropertyOPTS) error
	SetPresellType(unit.GoodsP)
	SetPresellDelay(uint8) error
	SetPresellEndTime(time.Duration) error
	SetCommit(bool)
	ArgBuild
}

// Build 构造参数列表
type ArgBuild interface {
	Build() ArgAdd
}

// NewArgAdd 实例商品添加参数
func NewArgAdd(name string) ArgAddInterface {
	return &ArgAdd{Name: name}
}

// NewArgEdit 实例商品编辑参数
func NewArgEdit(p unit.ProductID) ArgEdit {
	return &ArgAdd{ProductID: string(p.GetProductID())}
}

// HookSkipCheck
// 这个方法是shop_sdk_douyin.ToMapData方法调用的
func (p ArgAdd) HookSkipCheck(_, _, v string) bool {
	// p.Commit == "" 即新增 新增时不可跳过 强制带值
	// p.Commit != "" 即修改 修改时可以跳过
	if p.Commit != "" {
		return v == ""
	}
	return false
}

// SetName 设置商品标题
func (p *ArgAdd) SetName(n string) {
	p.Name = n
}

// SetOutProductID 设置商家算成的商品id
func (p *ArgAdd) SetOutProductID(id uint64) {
	p.OutProductID = strconv.FormatUint(id, 10)
}

// SetOutProductID 商品轮播图
// 多张图片用 "|" 分开，第一张图为主图，最多5张，至少600x600，大小不超过1M
func (p *ArgAdd) SetPic(img Pic) error {
	if len(img) > 5 {
		return errors.New("pic max 5 Zhang")
	}
	p.Pic = img.JoinString()
	return nil
}

// SetDescription 商品描述
// 目前只支持图片。多张图片用 "|" 分开。不能用其他网站的文本粘贴，这样会出现css样式不兼容
func (p *ArgAdd) SetDescription(img Pic) {
	p.Description = img.JoinString()
}

// SetPrice 设置价格 p1售价 p2划线价 单位元
func (p *ArgAdd) SetPrice(p1, p2 float64) {
	p.DiscountPrice = unit.PriceToYuan(p1).String()
	p.MarketPrice = unit.PriceToYuan(p2).String()
}

// SetCid 设置商品分类 c1一级 c2二级 c3三级
func (p *ArgAdd) SetCid(c1, c2, c3 unit.ProductCID) {
	p.FirstCid = c1.String()
	p.SecondCid = c2.String()
	p.ThirdCid = c3.String()
}

// SetSpecID 规格id
//要先创建商品通用规格, 如颜色-尺寸
func (p *ArgAdd) SetSpecID(s unit.ProductSpec) {
	p.SpecID = s.GetProductSpecID().String()
}

// SetSpecPic 主规格id
// 如颜色-尺寸, 颜色就是主规格, 颜色如黑,白,黄,它们的id|图片url 可含图片，亦可不含图片
func (p *ArgAdd) SetSpecPic(s []spec.Pic) error {
	l := len(s)
	r := make([]string, l)
	for i := 0; i < l; i++ {
		if s[i].ID == 0 {
			return fmt.Errorf("index %d field id cannot be 0", i)
		}
		if s[i].Pic == "" {
			r[i] = s[i].ID.String()
		} else {
			r[i] = s[i].ID.String() + unit.SPE1 + s[i].Pic
		}
	}
	p.SpecPic = strings.Join(r, unit.SPE2)
	return nil
}

// SetMobile 客服号码
func (p *ArgAdd) SetMobile(m string) {
	p.Mobile = m
}

// SetWeight 商品重量 (单位:克)。范围: 10克 - 9999990克
func (p *ArgAdd) SetWeight(w uint32) error {
	if w < 10 {
		return errors.New("weight min 10g")
	}
	if w > 9999990 {
		return errors.New("weight max 9999990g")
	}
	p.Weight = strconv.FormatUint(uint64(w), 10)
	return nil
}

// SetProductFormat 属性名称|属性值 之间用|分隔, 多组之间用^分开
func (p *ArgAdd) SetProductFormat(o unit.PropertyOPTS) error {
	l := len(o)
	r := make([]string, l)
	for i := 0; i < l; i++ {
		if o[i].Name == "" {
			return fmt.Errorf("index %d field name empty", i)
		}
		r[i] = o[i].Name + unit.SPE1 + o[i].Value
	}
	p.ProductFormat = strings.Join(r, unit.SPE2)
	return nil
}

// SetPayType 支付方式，必填，0-货到付款，1-在线支付，2-二者都支持
func (p *ArgAdd) SetPayType(t unit.PayT) {
	p.PayType = t.String()
}

// SetRecommendRemark 商家推荐语，不能含emoj表情
func (p *ArgAdd) SetRecommendRemark(r string) {
	p.RecommendRemark = r
}

// SetBrandID 品牌id
func (p *ArgAdd) SetBrandID(b unit.BrandID) {
	p.BrandID = b.String()
}

// SetPresellType 预售类型，1-全款预售，0-非预售，默认0
func (p *ArgAdd) SetPresellType(t unit.GoodsP) {
	p.PresellType = t.String()
}

//  SetPresellDelay 预售结束后，几天发货，可以选择2-30
func (p *ArgAdd) SetPresellDelay(d uint8) error {
	if d < 2 {
		return errors.New("presell delay min 2 day")
	}
	if d > 30 {
		return errors.New("presell delay max 30 day")
	}
	p.PresellDelay = strconv.FormatUint(uint64(d), 10)
	return nil
}

// SetPresellEndTime 预售结束时间，格式2020-02-21 18:54:27，最多支持设置距离当前30天
func (p *ArgAdd) SetPresellEndTime(t time.Duration) error {
	if t < 1 {
		return errors.New("presell end time min 1 day")
	}
	if t > 30 {
		return errors.New("presell end time max 30 day")
	}
	p.PresellEndTime = time.Now().Add((time.Hour * 24) * t).Format(unit.TimeYmdHis)
	return nil
}

// SetDeliveryDelayDay 承诺发货时间，单位是天，可选值为: 2、3、5、7、10、15
func (p *ArgAdd) SetDeliveryDelayDay(d uint8) {
	p.DeliveryDelayDay = strconv.FormatUint(uint64(d), 10)
}

// SetQualityReport 商品创建和编辑操作支持设置质检报告链接,多个图片以逗号分隔
func (p *ArgAdd) SetQualityReport(img Pic) {
	p.QualityReport = strings.Join(img, unit.SPE3)
}

// SetClassQuality 商品创建和编辑操作支持设置品类资质链接,多个图片以逗号分隔
func (p *ArgAdd) SetClassQuality(img Pic) {
	p.ClassQuality = strings.Join(img, unit.SPE3)
}

// SetCommit
// true：编辑后立即提交审核；false：编辑后仅保存，不提审
func (p *ArgAdd) SetCommit(is bool) {
	if is {
		p.Commit = "1"
	} else {
		p.Commit = "2"
	}
}

// Build 获取参数
func (p *ArgAdd) Build() ArgAdd {
	if p.ProductID != "" && p.Commit == "" {
		p.Commit = "2"
	} else {
		p.Commit = "1"
	}
	return *p
}
