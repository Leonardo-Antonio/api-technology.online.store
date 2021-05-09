package utils

import "errors"

func errRequiredField(field string) error {
	return errors.New("the " + field + " field is required")
}
