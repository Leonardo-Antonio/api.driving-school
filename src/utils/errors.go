package utils

import "errors"

var (
	ErrNotRecordsFoundUser = errors.New("no user data found")
	ErrDniNotFound         = errors.New("the dni was not found")
	ErrDniInvalid          = errors.New("the dni entered is not correct")
)
