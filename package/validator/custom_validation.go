package validator

import (
	"fmt"
	"reflect"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

func validateRequiredIf(fl validator.FieldLevel) bool {
	f := strings.Split(fl.Param(), ":")
	field := f[0]
	value := f[1]
	divefield := strings.Split(field, ".")

	q := reflect.Indirect(fl.Top())

	for _, field := range divefield {
		q = q.FieldByName(field)
	}

	if fmt.Sprintf("%v", q.Interface()) == reflect.ValueOf(value).Interface() {
		validate := validator.New()
		if err := validate.Var(fl.Field().Interface(), "required"); err != nil {
			return false
		}
		return true
	}
	return true
}
