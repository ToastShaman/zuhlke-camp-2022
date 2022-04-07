package validation

import (
	"strings"

	"github.com/go-playground/validator"
)

func NonBlankVal(fl validator.FieldLevel) bool {
	return len(strings.TrimSpace(fl.Field().String())) > 0
}
