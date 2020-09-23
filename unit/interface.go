package unit

type Product interface {
	GetProductID() ProductID
}

type ProductSpec interface {
	GetProductSpecID() SpecID
	Len() int
	HasSub(int, SpecID) bool
	GetSub(int) []SpecID
}

type SkuOperate interface {
	Product
	GetSkuID() SkuID
}

type Order interface {
	GetParentID() OrderID
}
