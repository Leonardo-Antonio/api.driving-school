package utils

import "errors"

var (
	ErrNotRecordsFoundUser = errors.New("no user data found")
	ErrDniNotFound         = errors.New("the dni was not found")
	ErrDniInvalid          = errors.New("the dni entered is not correct")
	ErrIdInvalid           = errors.New("the id you have entered is invalid")
	ErrNamePackageNotExist = errors.New("package name does not exist")
)
