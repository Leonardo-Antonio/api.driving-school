package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Package struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" xml:"_id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Src         string             `bson:"src,omitempty" json:"src,omitempty" xml:"src,omitempty"`
	Content     []string           `bson:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	DeletedAt   time.Time          `bson:"deleted_at,omitempty" json:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
	Active      bool               `bson:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
	CreatedBy   primitive.ObjectID `bson:"created_by,omitempty" json:"created_by,omitempty" xml:"created_by,omitempty"`
	UpdatedBy   primitive.ObjectID `bson:"updated_by,omitempty" json:"updated_by,omitempty" xml:"updated_by,omitempty"`
	DeletedBy   primitive.ObjectID `bson:"deleted_by,omitempty" json:"deleted_by,omitempty" xml:"deleted_by,omitempty"`
}
