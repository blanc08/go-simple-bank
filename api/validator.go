package api

import (
	"github.com/blanc08/go-simple-bank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// Check currency support
		return util.IsSupportedCurrency(currency)
	}
	return false
}
