package internal

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ComparePassword(passwordConfirmation string) validation.RuleFunc {
	return func(value interface{}) error {
		if value.(string) == passwordConfirmation {
			return nil
		}
		return errors.New("password_confirmation_failed")
	}
}
