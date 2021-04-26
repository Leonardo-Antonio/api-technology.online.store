package utils

import "errors"

var (
	ErrDeleteOne = errors.New("could not delete user")
	ErrUpdateOne = errors.New("more than one record is updated")
)

