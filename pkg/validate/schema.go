package validate

import (
	"gopkg.in/go-playground/validator.v9"
	"market/pkg/e"
)

type Schema interface {
	Validate(err validator.ValidationErrors) e.MarketError
}
