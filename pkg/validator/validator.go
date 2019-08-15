package validator

import (
	"github.com/gin-gonic/gin/binding"
)

func init() {
	binding.Validator = new(defaultValidator)
}
