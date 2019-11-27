package models

import (
	"market/schema"
)

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

type ProductCategories []*ProductCategory

func (category *ProductCategory) ToSchemaCategory() (schemaCategory *schema.Category) {
	schemaCategory = &schema.Category{
		ID:          category.ID,
		ParentId:    category.ParentId,
		CatName:     category.CatName,
		CatLevel:    category.CatLevel,
		CatKeyWords: category.CatKeyWords,
		CatIcon:     category.CatIcon,
		CatDesc:     category.CatDesc,
	}
	return
}

func (categories ProductCategories) ToSchemaCategories() []*schema.Category {
	schemaCategories := make([]*schema.Category, len(categories))
	for index, category := range categories {
		schemaCategories[index] = category.ToSchemaCategory()
	}
	return schemaCategories
}

type ProductCategorySpec struct {
	// 商品分类规格 用于确定商品的规格模板
	BaseModel
	SpecName   string `gorm:"type:varchar(64);not null;unique_index:uq_cat_id_spec" json:"spec_name"` // 分类规格名称, 颜色 ...
	JoinSelect *bool  `gorm:"type:tinyint(1);index;not null" json:"join_select"`                      // 是否可以筛选
	SpecType   uint   `gorm:"type:tinyint(1);index;not null" json:"spec_type"`                        // 规格类型  1 销售规格属性 2 展示属性
	CatId      uint   `gorm:"index;not null;unique_index:uq_cat_id_spec" json:"cat_id"`               // 商品分类id
}

type ProductCategorySpecs []*ProductCategorySpec

func (spec *ProductCategorySpec) ToSchemaCategorySpec() (schemaCategorySpec *schema.CategorySpec) {
	schemaCategorySpec = &schema.CategorySpec{
		ID:         spec.ID,
		SpecName:   spec.SpecName,
		JoinSelect: spec.JoinSelect,
		SpecType:   spec.SpecType,
		CatId:      spec.CatId,
	}
	return
}

func (specs ProductCategorySpecs) ToSchemaCategorySpecs() []*schema.CategorySpec {
	schemaSpecs := make([]*schema.CategorySpec, len(specs))
	for index, spec := range specs {
		schemaSpecs[index] = spec.ToSchemaCategorySpec()
	}
	return schemaSpecs
}

//Product 商品 SPU
type Product struct {
	BaseModel
	ProductName string `gorm:"type:varchar(64);index:product_name" json:"product_name"` // 商品名称
	ProductSn   string `gorm:"type:varchar(24);unique" json:"product_sn"`               //商品货号
	SubTitle    string `gorm:"type:varchar(128)" json:"sub_title"`                      // 副标题
	CatId       uint   `gorm:"index;not null" json:"cat_id"`                            //商品分类id
	BrandId     uint   `gorm:"index;not null" json:"brand_id"`                          // 品牌id
	StoreId     uint   `gorm:"index;not null" json:"store_id"`                          // 商铺id
	Unit        uint   `gorm:"type:varchar(32)" json:"unit"`                            // 单位(件/台...)
	Published   *bool  `gorm:"type:tinyint(1);index;not null" json:"published"`         // 上架状态
}

//	获取商品分类
func GetCategories(condition interface{}) (categories ProductCategories, err error) {
	err = db.Where(condition).Find(&categories).Error
	return
}

// 获取商品分类总数
func GetCategoryCount(condition interface{}) (count int, err error) {
	err = db.Model(ProductCategory{}).Where(condition).Count(&count).Error
	return
}

// 获取商品分类规格信息
func GetCategorySpecs(condition interface{}) (specs ProductCategorySpecs, err error) {
	err = db.Where(condition).Find(&specs).Error
	return
}
