package validate

import (
	"errors"
	"strings"

	"github.com/Leonardo-Antonio/api.driving-school/src/utils/const/turn"
)

var errTurnRequired = errors.New("the shift field is required")

func Turn(turnUser string) error {
	if len(turnUser) == 0 {
		return errTurnRequired
	}
	if strings.EqualFold(turn.AFTERNOON, turnUser) ||
		strings.EqualFold(turn.MORNING, turnUser) ||
		strings.EqualFold(turn.NIGHT, turnUser) {
		return nil
	}

	return errTurnInvalid
}
