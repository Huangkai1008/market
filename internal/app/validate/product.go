package validate

import (
	"gopkg.in/go-playground/validator.v9"
	"market/pkg/e"
)

/************************************/
/**********   商品模块验证    ********/
/************************************/
type CategoryQuery struct {
	ParentId string `form:"parent_id" validate:"required,gte=0"`
}

func (c *CategoryQuery) Validate(errs validator.ValidationErrors) e.MarketError {
	var marketError e.MarketError
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

type CategorySpecUri struct {
	CatId uint `uri:"cat_id" validate:"required,gt=0"`
}

func (c CategorySpecUri) Validate(errs validator.ValidationErrors) e.MarketError {
	var marketError e.MarketError
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
