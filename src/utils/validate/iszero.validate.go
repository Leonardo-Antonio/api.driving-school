package validate

import (
	"github.com/Leonardo-Antonio/api.driving-school/src/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsZeroObjectIds(ids ...primitive.ObjectID) error {
	for _, id := range ids {
		if id.IsZero() {
			return utils.ErrIdInvalid
		}
	}
	return nil
}
