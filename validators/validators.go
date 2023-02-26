package validators

import (
	"strings"

	"github.com/go-playground/validator"
)

func ValidateCoolTitle(field validator.FieldLevel) bool {
	fl := field.Field().String()

	if !strings.Contains(fl,"Cool"){
		return false
	}
	
	return true
}
