package utils

import "errors"

var (
	ErrDeleteOne         = errors.New("could not delete user")
	ErrUpdateOne         = errors.New("more than one record is updated")
	ErrInsertOne         = errors.New("more than one record is inserted")
	ErrIDRequired        = errors.New("the id field is required")
	ErrUserNotExists     = errors.New("username does not exist")
	ErrCategoryNotExists = errors.New("category does not exist")
)
