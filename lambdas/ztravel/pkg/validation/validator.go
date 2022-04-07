package validation

import "github.com/go-playground/validator"

func New() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("nonblank", NonBlankVal)
	return validate
}
