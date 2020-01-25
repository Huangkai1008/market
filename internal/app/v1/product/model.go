package product

import (
	"market/internal/pkg/database/model"
	"market/internal/pkg/utils"
)

// CategorySchema 商品分类模型
type Category struct {
	ID          uint           `gorm:"primary_key" json:"id"`
	ParentId    uint           `gorm:"index;not null;default:0;comment:'父分类, 0表示一级分类'" json:"parent_id"`           // 父分类, 0表示一级分类
	CatName     string         `gorm:"type:varchar(64);not null;unique;comment:'分类名'" json:"cat_name"`             // 分类名
	CatLevel    uint8          `gorm:"type:tinyint(1);not null;index;comment:'分类等级,0->1级,1->2级'" json:"cat_level"` // 分类等级, 0->1级,1->2级
	CatKeywords string         `gorm:"type:varchar(255);comment:'分类关键词'" json:"cat_keywords"`                      // 分类关键词
	CatIcon     string         `gorm:"type:varchar(255);comment:'分类图标'" json:"cat_icon"`                           // 分类图标
	CatDesc     string         `gorm:"type:text;comment:'分类描述'" json:"cat_desc"`                                   // 分类描述
	CreatedAt   utils.JsonTime `gorm:"type:datetime;column:create_time;comment:'创建时间'" json:"create_time"`
	UpdatedAt   utils.JsonTime `gorm:"type:datetime;column:update_time;comment:'更新时间'" json:"update_time"`
}

type Categories []*Category

func (Category) TableName() string {
	return "product_category"
}

func (category *Category) ToCategorySchema() (categorySchema *CategorySchema) {
	categorySchema = &CategorySchema{
		ID:          category.ID,
		ParentId:    category.ParentId,
		CatName:     category.CatName,
		CatLevel:    category.CatLevel,
		CatKeywords: category.CatKeywords,
		CatIcon:     category.CatIcon,
		CatDesc:     category.CatDesc,
	}
	return
}

func (categories Categories) ToCategorySchema() []*CategorySchema {
	categorySchemas := make([]*CategorySchema, len(categories))
	for index, category := range categories {
		categorySchemas[index] = category.ToCategorySchema()
	}
	return categorySchemas
}

// CategorySpec  商品分类规格 用于确定商品的规格模板
type CategorySpec struct {
	model.BaseModel
	SpecName   string `gorm:"type:varchar(64);not null;unique_index:uq_cat_id_spec;comment:'分类规格名称'" json:"spec_name"` // 分类规格名称, 颜色 ...
	JoinSelect *bool  `gorm:"type:tinyint(1);index;not null;comment:'是否可以筛选'" json:"join_select"`                      // 是否可以筛选
	SpecType   uint   `gorm:"type:tinyint(1);index;not null;comment:'规格类型  1 销售规格属性 2 展示属性'" json:"spec_type"`         // 规格类型  1 销售规格属性 2 展示属性
	CatId      uint   `gorm:"index;not null;unique_index:uq_cat_id_spec;comment:'商品分类id'" json:"cat_id"`               // 商品分类id
}

type CategorySpecs []*CategorySpec

func (CategorySpec) TableName() string {
	return "product_category_spec"
}

func (spec *CategorySpec) ToCategorySpecSchema() (categorySpecSchema *CategorySpecSchema) {
	categorySpecSchema = &CategorySpecSchema{
		ID:         spec.ID,
		SpecName:   spec.SpecName,
		JoinSelect: spec.JoinSelect,
		SpecType:   spec.SpecType,
		CatId:      spec.CatId,
	}
	return
}

func (specs CategorySpecs) ToCategorySpecSchemas() []*CategorySpecSchema {
	specSchemas := make([]*CategorySpecSchema, len(specs))
	for index, spec := range specs {
		specSchemas[index] = spec.ToCategorySpecSchema()
	}
	return specSchemas
}

// Product 商品SPU模型
type Product struct {
	model.BaseModel
	ProductName string `gorm:"type:varchar(64);index;comment:'商品名称'" json:"product_name"`      // 商品名称
	ProductSn   string `gorm:"type:varchar(24);unique;comment:'商品货号'" json:"product_sn"`       // 商品货号
	SubTitle    string `gorm:"type:varchar(128);comment:'副标题'" json:"sub_title"`               // 副标题
	CatId       uint   `gorm:"type:bigint(10);index;not null;comment:'商品分类id'" json:"cat_id"`  // 商品分类id
	BrandId     uint   `gorm:"type:bigint(10);index;not null;comment:'品牌id'" json:"brand_id"`  // 品牌id
	StoreId     uint   `gorm:"type:bigint(10);index;not null;comment:'商铺id'" json:"store_id"`  // 商铺id
	Unit        uint   `gorm:"type:varchar(32);comment:'单位(件/台...)'" json:"unit"`              // 单位(件/台...)
	Published   *bool  `gorm:"type:tinyint(1);index;not null;comment:'上架状态'" json:"published"` // 上架状态
}
