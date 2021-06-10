package validate

import (
	"errors"

	"github.com/Leonardo-Antonio/api.driving-school/src/entity"
)

var (
	errFieldNameRequired = errors.New("the name field is required")
	errDescriptionLength = errors.New("the package description is too little")
	errItems             = errors.New("must have at least 2 items")
)

func Package(pack entity.Package) error {
	if len(pack.Name) == 0 {
		return errFieldNameRequired
	}

	if len(pack.Description) < 10 {
		return errDescriptionLength
	}

	if len(pack.Content) < 2 {
		return errItems
	}

	return nil
}
