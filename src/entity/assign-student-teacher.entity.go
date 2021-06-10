package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssignStudentTeacher struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" xml:"_id,omitempty"`
	IdClient  primitive.ObjectID `bson:"id_client,omitempty" json:"id_client,omitempty" xml:"id_client,omitempty"`
	IdTeacher primitive.ObjectID `bson:"id_teacher,omitempty" json:"id_teacher,omitempty" xml:"id_teacher,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	Active    bool               `bson:"active,omitempty" json:"active,omitempty" xml:"active"`
	CreatedBy primitive.ObjectID `bson:"created_by,omitempty" json:"created_by,omitempty" xml:"created_by,omitempty"`
	UpdatedBy primitive.ObjectID `bson:"updated_by,omitempty" json:"updated_by,omitempty" xml:"updated_by,omitempty"`
	DeletedBy primitive.ObjectID `bson:"deleted_by,omitempty" json:"deleted_by,omitempty" xml:"deleted_by,omitempty"`
}
