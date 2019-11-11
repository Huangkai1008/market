package models

type ProductCategory struct {
	// 商品分类
	BaseModel
	ParentId    uint   `gorm:"index:parent_id;not null;default:0" json:"parent_id"` // 父分类, 0表示一级分类
	CatName     string `gorm:"type:varchar(64);unique" json:"cat_name"`             // 分类名
	CatLevel    uint8  `gorm:"type:tinyint(1);index:cat_level" json:"cat_level"`    // 分类等级, 0->1级; 1->2级
	CatKeyWords string `gorm:"type:varchar(255)" json:"cat_key_words"`              // 分类关键词
	CatIcon     string `gorm:"type:varchar(255)" json:"cat_icon"`                   // 分类图标
	CatDesc     string `gorm:"type:text" json:"cat_desc"`                           // 分类描述
}

type ProductCategorySpec struct {
	// 商品分类规格 用于确定商品的规格模板
	BaseModel
	SpecName   string `gorm:"type:varchar(64);not null;unique_index:uq_cat_id_spec" json:"spec_name"` // 分类规格名称, 颜色 ...
	JoinSelect *bool  `gorm:"type:tinyint(1);index;not null" json:"join_select"`                      // 是否可以筛选
	SpecType   uint   `gorm:"type:tinyint(1);index;not null" json:"spec_type"`                        // 规格类型  1 销售规格属性 2 展示属性
	CatId      uint   `gorm:"index;not null;unique_index:uq_cat_id_spec" json:"cat_id"`               // 商品分类id
}

func GetCategories(condition interface{}) (categories []ProductCategory, err error) {
	//	获取商品分类
	err = db.Where(condition).Find(&categories).Error
	return
}

func GetCategoryCount(condition interface{}) (count int, err error) {
	// 获取商品分类总数
	err = db.Model(ProductCategory{}).Where(condition).Count(&count).Error
	return
}
