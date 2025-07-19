package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
)

var validatorMessages = map[string]string{
	"required": "Поле является обязательным",
	"number":   "Значение поля должно быть числом",
	"min":      "Значение поля меньше минимального",
	"max":      "Значение поля превышает допустимое значение",
	"datetime": "Неправильный формат даты",
	"oneof":    "Недопустимое значение",
}

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator.New(),
	}
}

func (v Validator) Struct(structInput interface{}) *exceptions.ValidationError {

	err := v.validate.Struct(structInput)
	if err != nil {
		return exceptions.NewValidationError(
			v.getErrors(err.(validator.ValidationErrors)),
		)
	}

	return nil
}

func (v Validator) getErrors(errors validator.ValidationErrors) []exceptions.ValidationField {

	var validationErr []exceptions.ValidationField

	for _, err := range errors {

		field := strings.ToLower(err.Field())
		tag := err.Tag()

		msg, ok := validatorMessages[tag]
		if !ok {
			msg = "Недопустимое значение"
		}

		validationErr = append(validationErr, exceptions.ValidationField{
			Name: field,
			Err:  fmt.Errorf("%s", msg),
		})
	}

	return validationErr
}
