package validator

import (
	"github.com/google/uuid"
	validatorV9 "gopkg.in/go-playground/validator.v9"
	"nbhd/env/validator"
	"regexp"
)

type Validator struct {
	v *validatorV9.Validate
}

func (validator Validator) Process(i interface{}) error {
	return validator.v.Struct(i)
}

func NewValidator() validator.Validator {
	v := validatorV9.New()
	v.RegisterValidation("uuid", isUUID)
	v.RegisterValidation("phone", IsPhone)
	return Validator{v}
}

func isUUID(f validatorV9.FieldLevel) bool {
	v := f.Field().String()
	_, err := uuid.Parse(v)
	return err == nil
}

func IsPhone(f validatorV9.FieldLevel) bool {
	v := f.Field().String()
	p, _ := regexp.Compile("^[0-9]{10}$")
	return p.MatchString(v)
}
