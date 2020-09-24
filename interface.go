package shop_sdk_douyin

import (
	"github.com/cgghui/shop_sdk_douyin/aftersale"
	"github.com/cgghui/shop_sdk_douyin/logistics"
	"github.com/cgghui/shop_sdk_douyin/order"
	"github.com/cgghui/shop_sdk_douyin/product"
	"github.com/cgghui/shop_sdk_douyin/product/sku"
	"github.com/cgghui/shop_sdk_douyin/product/spec"
	"github.com/cgghui/shop_sdk_douyin/unit"
)

// 以下接口 shop_sdk_douyin.App 都已经实现

// ProductAPI 商品接口
type ProductAPI interface {
	ProductList(product.ArgList) (product.ResponseList, error)
	ProductDetail(unit.ProductID, ...bool) (product.ResponseDetail, error)
	ProductCategory(...unit.ProductCID) ([]product.ResponseCategory, error)
	ProductCateProperty(cid1, cid2, cid3 unit.ProductCID) ([]product.ResponseCateProperty, error)
	ProductAdd(product.ArgAdd) (product.Product, error)
	ProductEdit(product.ArgAdd) error
	ProductDel(unit.ProductID) error
}

// ProductSpecAPI 规格接口
type ProductSpecAPI interface {
	SpecAdd(spec.ArgAdd) (spec.ResponseAdd, error)
	SpecList() ([]spec.ResponseList, error)
	SpecDetail(id unit.SpecID) (spec.ResponseDetail, error)
	SpecDel(unit.SpecID) error
}

// ProductSkuAPI SKU接口
type ProductSkuAPI interface {
	SkuAdd(sku.ArgAdd) (map[uint64]unit.SkuID, error)
	SkuList(unit.ProductID) ([]sku.ResponseList, error)
	SkuDetail(unit.SkuID) (sku.ResponseDetail, error)
	SkuEditPrice(unit.SkuOperate, float64) error
	SkuSyncStock(unit.SkuOperate, uint16) error
	SkuEditCode(unit.SkuOperate, string) error
}

// OrderAPI 订单接口
type OrderAPI interface {
	OrderList(order.ArgList) (order.ResponseList, error)
	OrderDetail(unit.Order) (order.Detail, error)
	OrderStockUp(unit.Order) error
	OrderCancel(unit.Order, string) error
	OrderServiceList(order.ArgServiceList) (order.ResponseServiceList, error)
	OrderReplyService(unit.ServiceID, string) error
	OrderLogisticsAdd(order.ArgLogisticsAdd) error
	OrderLogisticsEdit(order.ArgLogisticsAdd) error
}

// LogisticsAPI 物流接口
type LogisticsAPI interface {
	OrderLogisticsCompanyList() (logistics.ResponseLogisticsCompanyList, error)
	OrderLogisticsAdd(order.ArgLogisticsAdd) error
	OrderLogisticsEdit(order.ArgLogisticsAdd) error
	AddressProvinceList() ([]logistics.Province, error)
	AddressCityList(uint32) ([]logistics.City, error)
	AddressAreaList(uint32) ([]logistics.Area, error)
}

// AfterSaleAPI 售后接口
type AfterSaleAPI interface {
	RefundOrderList(aftersale.ArgRefundOrderList) (order.ResponseList, error)
	RefundShopRefund(aftersale.ArgRefundShopRefund) error
	AfterSaleOrderList(aftersale.ArgAfterSaleOrderList) (order.ResponseList, error)
	AfterSaleBuyerReturn(aftersale.ArgAfterSaleBuyerReturn) error
	AfterSaleFirmReceive(aftersale.ArgAfterSaleFirmReceive) error
	AfterSaleUploadCompensation(aftersale.ArgAfterSaleUploadCompensation) error
	AfterSaleAddOrderRemark(unit.OrderID, string) error
	AfterSaleRefundProcessDetail(unit.OrderID) (aftersale.ResponseAfterSaleRefundProcessDetail, error)
	AfterSaleBuyerRefund(aftersale.ArgAfterSaleBuyerRefund) error
}
