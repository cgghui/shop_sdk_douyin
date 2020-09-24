package logistics

//  Province 省
type Province struct {
	ID       uint32 `mapstructure:"province_id"`
	Province string `mapstructure:"province"`
	FatherID uint32 `mapstructure:"father_id"`
}

// City 市
type City struct {
	ID       uint32 `mapstructure:"city_id"`
	City     string `mapstructure:"city"`
	FatherID uint32 `mapstructure:"father_id"`
}

// Area 区
type Area struct {
	ID       uint32 `mapstructure:"area_id"`
	Area     string `mapstructure:"area"`
	FatherID uint32 `mapstructure:"father_id"`
}
