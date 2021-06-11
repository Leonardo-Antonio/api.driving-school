package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sale struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" xml:"_id,omitempty"`
	IdClient  primitive.ObjectID `bson:"id_client,omitempty" json:"id_client,omitempty" xml:"id_client,omitempty"`
	IdPackage primitive.ObjectID `bson:"id_package,omitempty" json:"id_package,omitempty" xml:"id_package,omitempty"`
	Turn      string             `bson:"turn,omitempty" json:"turn,omitempty" xml:"turn,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	Active    bool               `bson:"active,omitempty" json:"active,omitempty"`
}
