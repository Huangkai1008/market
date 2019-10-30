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

func GetCategories(maps interface{}) (categories []ProductCategory, err error) {
	//	获取商品分类
	err = db.Where(maps).Find(&categories).Error
	return
}

func GetCategoryCount(maps interface{}) (count int, err error) {
	// 获取商品分类总数
	err = db.Model(ProductCategory{}).Where(maps).Count(&count).Error
	return
}
