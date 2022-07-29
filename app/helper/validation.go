package helper

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func ValidateBase(data interface{}) (bool, string) {
	var validate *validator.Validate
	Message := ""

	validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}
		reflected := reflect.ValueOf(data)

		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflected.Type().FieldByName(err.StructField())
			var name string

			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				Message = "The " + name + " is required"
				break
			case "email":
				Message = "The " + name + " should be a valid email"
				break
			default:
				Message = "The " + name + " is invalid"
				break
			}

			return false, Message
		}
	}
	return true, ""
}
