package utils

import "github.com/Leonardo-Antonio/api-technology.online.store/src/entity"

func CategoryValidation(category entity.Category) error {
	if len(category.Name) == 0 {
		return errRequiredField("name")
	}
	return nil
}
