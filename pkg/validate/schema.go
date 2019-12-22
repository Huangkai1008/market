package validate

import (
	"market/pkg/e"

	"gopkg.in/go-playground/validator.v9"
)

type Schema interface {
	Validate(errs validator.ValidationErrors) e.MarketError
}
