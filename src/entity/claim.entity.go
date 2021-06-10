package entity

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClaimUser struct {
	ID        primitive.ObjectID `json:"_id,omitempty" xml:"_id,omitempty" bson:"_id,omitempty"`
	Email     string             `json:"email,omitempty" xml:"email,omitempty" bson:"email,omitempty"`
	Dni       string             `json:"dni,omitempty" xml:"dni,omitempty" bson:"dni,omitempty"`
	Names     string             `json:"names,omitempty" xml:"names,omitempty" bson:"names,omitempty"`
	LastNames string             `json:"last_names,omitempty" xml:"last_names,omitempty" bson:"last_names,omitempty"`
	Rol       string             `json:"rol,omitempty" xml:"rol,omitempty" bson:"rol,omitempty"`
	jwt.StandardClaims
}
