package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func cpfValidation(fl validator.FieldLevel) bool {
	const cpfRegex = `^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`

	r, _ := regexp.Compile(cpfRegex)

	return r.MatchString(fl.Field().String())
}
