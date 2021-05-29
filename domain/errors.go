package domain

import "errors"

var (
	ErrInternalServer = errors.New("Internal Server Error")
	ErrNotFounds      = errors.New("Your Requested User is not found")
	ErrConflict       = errors.New("Your User Already exist!")
	ErrBadParamInput  = errors.New("Given Param is not valid")
)
