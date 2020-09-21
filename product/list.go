package product

// ArgList ProductList方法的参数
type ArgList struct {
	Page        uint8 // 第几页（第一页为0）
	Size        uint8 // 每页返回条数
	Status      SS    // 指定状态返回商品列表
	CheckStatus SC    `paramName:"check_status"` // 指定审核状态返回商品列表
}

// ResponseList ProductList方法的响应结果
type ResponseList struct {
	All         uint32    `mapstructure:"all"`          // 商品总数
	AllPages    uint32    `mapstructure:"all_pages"`    // 已当前size所得的分页数
	Count       uint32    `mapstructure:"count"`        // 当前条件data返回结果数量
	CurrentPage uint32    `mapstructure:"current_page"` // 当前页
	Data        []Product `mapstructure:"data"`         // 商品列表
	PageSize    uint32    `mapstructure:"page_size"`    // 每页条数
}
