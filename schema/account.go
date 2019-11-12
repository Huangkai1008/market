package schema

type Address struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`      // 用户id    用户1 --> 地址N
	Consignee   string `json:"consignee"`    // 收货人姓名
	Mobile      string `json:"mobile"`       // 手机号码
	Region      string `json:"region"`       // 所在地区
	FullAddress string `json:"full_address"` // 详细地址
	Tag         string `json:"tag"`          // 标签
	IsDefault   *bool  `json:"is_default"`   // 是否默认地址
}
