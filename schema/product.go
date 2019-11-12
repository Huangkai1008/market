package schema

type Category struct {
	ID          uint   `json:"id"`
	ParentId    uint   `json:"parent_id"`     // 父分类, 0表示一级分类
	CatName     string `json:"cat_name"`      // 分类名
	CatLevel    uint8  `json:"cat_level"`     // 分类等级, 0->1级; 1->2级
	CatKeyWords string `json:"cat_key_words"` // 分类关键词
	CatIcon     string `json:"cat_icon"`      // 分类图标
	CatDesc     string `json:"cat_desc"`      // 分类描述
}

type CategoryList struct {
	categories []*Category
	total      int
}

type CategorySpec struct {
	ID         uint   `json:"id"`
	SpecName   string `json:"spec_name"`   // 分类规格名称, 颜色 ...
	JoinSelect *bool  `json:"join_select"` // 是否可以筛选
	SpecType   uint   `json:"spec_type"`   // 规格类型  1 销售规格属性 2 展示属性
	CatId      uint   `json:"cat_id"`      // 商品分类id
}
