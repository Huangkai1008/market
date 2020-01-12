package schema

import (
	"gopkg.in/go-playground/validator.v9"
	"market/internal/pkg/ecode"
	"market/internal/pkg/utils"
)

type Schema interface {
	Validate(errs validator.ValidationErrors) ecode.MarketError
}

type BaseSchema struct {
	ID        uint           `json:"id"`
	CreatedAt utils.JsonTime `json:"create_time"`
	UpdatedAt utils.JsonTime `json:"update_time"`
}
