package models

type ProductCategory struct {
	// 商品分类模型
	BaseModel
	ParentId    uint64 `gorm:"index:parent_id;not null;default:0" json:"parent_id"` // 父分类, 0表示一级分类
	CatName     string `gorm:"type:varchar(64);unique" json:"cat_name"`             // 分类名
	CatLevel    uint8  `gorm:"type:tinyint(1);index:cat_level" json:"cat_level"`    // 分类等级, 0->1级；1->2级
	CatKeyWords string `gorm:"type:varchar(255)" json:"cat_key_words"`              // 分类关键词
	CatDesc     string `gorm:"type:text" json:"cat_desc"`                           // 分类描述
}
