package utils

import (
	"errors"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/entity"
	"strings"
)

func errRequiredField(field string) error {
	return errors.New("the " + field + " field is required")
}

func UserValidate(user entity.User) error {
	if len(user.Name) == 0 { return errRequiredField("name") }
	if len(user.LastName) == 0 { return errRequiredField("last name") }
	if len(user.Nick) == 0 { return errRequiredField("nick name") }
	if len(user.Email) == 0 { return errRequiredField("email") }
	if !strings.ContainsAny(user.Email, "@") { return errors.New("enter a valid email address") }
	if len(user.Password) < 8 { return errors.New("enter a strong password") }
	if len(user.Rol) == 0 { return errRequiredField("rol") }
	switch strings.ToUpper(user.Rol) {
	case "USER":
		return nil
	case "ADMIN":
		return nil
	case "MANAGER":
		return nil
	default:
		return errors.New("does not have a correctly assigned role")
	}
	return nil
}
