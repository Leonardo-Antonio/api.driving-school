package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	DNI        string             `bson:"dni,omitempty" json:"dni,omitempty"`
	Email      string             `bson:"email,omitempty" json:"email,omitempty"`
	Names      string             `bson:"names,omitempty" json:"names,omitempty"`
	LastNames  string             `bson:"last_names,omitempty" json:"last_names,omitempty"`
	Password   string             `bson:"password,omitempty" json:"password,omitempty"`
	Src        string             `bson:"src,omitempty" json:"src,omitempty"`
	Rol        string             `bson:"rol,omitempty" json:"rol,omitempty"`
	StatusSale bool               `bson:"status_sale,omitempty" json:"status_sale,omitempty"`
	Turn       string             `bson:"turn,omitempty" json:"turn,omitempty"`
	CreatedAt  time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt  time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	Active     bool               `bson:"active,omitempty" json:"active,omitempty"`
}
