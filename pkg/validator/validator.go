package validator

import (
	form "github.com/go-playground/form/v4"
	playValidator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Validator struct {
	validator   *playValidator.Validate
	formDecoder *form.Decoder
}

var instance *Validator

func Initialize() {
	instance = &Validator{
		validator:   playValidator.New(),
		formDecoder: form.NewDecoder(),
	}

	registerUUIDCustomType()
	RegisterCustomValidation("cpf", cpfValidation)
}

func registerUUIDCustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (any, error) {
		return uuid.Parse(vals[0])
	}, uuid.UUID{})
}

func Struct(object any) error {
	return instance.validator.Struct(object)
}

func FormDecode(object any, values map[string][]string) error {
	return instance.formDecoder.Decode(object, values)
}

func RegisterCustomValidation(tag string, fn playValidator.Func) {
	instance.validator.RegisterValidation(tag, fn)
}
