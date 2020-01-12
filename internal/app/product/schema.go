package product

import (
	"gopkg.in/go-playground/validator.v9"

	"market/internal/pkg/ecode"
)

type CategoryQuerySchema struct {
	ParentId string `form:"parent_id" validate:"required,gte=0"`
}

func (c *CategoryQuerySchema) Validate(errs validator.ValidationErrors) ecode.MarketError {
	var marketError ecode.MarketError
	err := errs[0]

	if err.Field() == "ParentId" {
		switch err.Tag() {
		case "required":
			marketError.Message = "分类父级id不能为空"
		case "gte":
			marketError.Message = "分类父级id不能小于0"
		}
	}
	return marketError
}

type CategorySpecURISchema struct {
	CatId uint `uri:"cat_id" validate:"required,gt=0"`
}

func (c CategorySpecURISchema) Validate(errs validator.ValidationErrors) ecode.MarketError {
	var marketError ecode.MarketError
	err := errs[0]

	if err.Field() == "CatId" {
		switch err.Tag() {
		case "required":
			marketError.Message = "分类id不能为空"
		case "gt":
			marketError.Message = "分类id不能小于0"
		}
	}
	return marketError
}

type CategorySchema struct {
	ID          uint   `json:"id"`
	ParentId    uint   `json:"parent_id"`     // 父分类, 0表示一级分类
	CatName     string `json:"cat_name"`      // 分类名
	CatLevel    uint8  `json:"cat_level"`     // 分类等级, 0->1级; 1->2级
	CatKeyWords string `json:"cat_key_words"` // 分类关键词
	CatIcon     string `json:"cat_icon"`      // 分类图标
	CatDesc     string `json:"cat_desc"`      // 分类描述
}

type CategoryListSchema struct {
	categories []*CategorySchema
	total      int
}

type CategorySpecSchema struct {
	ID         uint   `json:"id"`
	SpecName   string `json:"spec_name"`   // 分类规格名称, 颜色 ...
	JoinSelect *bool  `json:"join_select"` // 是否可以筛选
	SpecType   uint   `json:"spec_type"`   // 规格类型  1 销售规格属性 2 展示属性
	CatId      uint   `json:"cat_id"`      // 商品分类id
}
