package autorization

import (
	"errors"
	"time"

	"github.com/Leonardo-Antonio/api.driving-school/src/entity"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(data *entity.User) (string, error) {
	claim := entity.ClaimUser{
		ID:        data.ID,
		Email:     data.Email,
		Dni:       data.DNI,
		Names:     data.Names,
		LastNames: data.LastNames,
		Rol:       data.Rol,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 50).Unix(),
			Issuer:    "Leonardo Antonio Nolasco Leyva",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ValidateToken .
func ValidateToken(t string) (entity.ClaimUser, error) {
	token, err := jwt.ParseWithClaims(t, &entity.ClaimUser{}, verifyFunction)
	if err != nil {
		return entity.ClaimUser{}, err
	}
	if !token.Valid {
		return entity.ClaimUser{}, errors.New("Token invalid")
	}

	claim, ok := token.Claims.(*entity.ClaimUser)
	if !ok {
		return entity.ClaimUser{}, errors.New("claims could not be obtained")
	}
	return *claim, nil
}

func verifyFunction(token *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
