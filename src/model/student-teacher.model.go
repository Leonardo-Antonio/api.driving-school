package model

import (
	"context"
	"time"

	"github.com/Leonardo-Antonio/api.driving-school/src/dbutil"
	"github.com/Leonardo-Antonio/api.driving-school/src/entity"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils/const/roles"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	studentTeacher struct {
		collection     *mongo.Collection
		collectionUser *mongo.Collection
	}

	IStudentTeacher interface {
		findByTurnStudent(turn string) (student []entity.User, err error)
		findByTurnTeacher(turn string) (teacher []entity.User, err error)
		FindByTurn(turn string) (studentTeacher entity.StudentTeacher, err error)
		AssingStudentToTeacher(assingStudentToTeacher entity.AssignStudentTeacher) (result *mongo.InsertOneResult, err error)
	}
)

func NewStudentTeacher(db *mongo.Database) *studentTeacher {
	return &studentTeacher{
		collection:     db.Collection(dbutil.CollectionStudentTeacher),
		collectionUser: db.Collection(dbutil.CollectionUsers),
	}
}

func (st *studentTeacher) findByTurnStudent(turn string) (student []entity.User, err error) {
	filterStudent := bson.M{
		"turn":        turn,
		"rol":         roles.CLIENT,
		"active":      true,
		"status_sale": true,
	}
	cursor, err := st.collectionUser.Find(context.TODO(), filterStudent)
	if err != nil {
		return
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &student); err != nil {
		return student, err
	}
	return
}

func (st *studentTeacher) findByTurnTeacher(turn string) (teacher []entity.User, err error) {
	filterStudent := bson.M{
		"turn":   turn,
		"rol":    roles.INSTRUCTOR,
		"active": true,
	}
	cursor, err := st.collectionUser.Find(context.TODO(), filterStudent)
	if err != nil {
		return
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &teacher); err != nil {
		return teacher, err
	}
	return
}

func (st *studentTeacher) FindByTurn(turn string) (studentTeacher entity.StudentTeacher, err error) {
	student, err := st.findByTurnStudent(turn)
	if err != nil {
		return
	}

	teacher, err := st.findByTurnTeacher(turn)
	if err != nil {
		return
	}

	studentTeacher.Client = student
	studentTeacher.Teacher = teacher

	return
}

func (st *studentTeacher) AssingStudentToTeacher(
	assingStudentToTeacher entity.AssignStudentTeacher,
) (result *mongo.InsertOneResult, err error) {
	assingStudentToTeacher.ID = primitive.NewObjectID()
	assingStudentToTeacher.CreatedAt = time.Now()
	assingStudentToTeacher.Active = true

	result, err = st.collection.InsertOne(context.TODO(), &assingStudentToTeacher)
	if err != nil {
		return nil, err
	}

	return
}

/*
func (st *studentTeacher) UpdateStudentStateAssignTeacher(Id user) (result *mongo.InsertOneResult, err error) {
	assingStudentToTeacher.ID = primitive.NewObjectID()
	assingStudentToTeacher.CreatedAt = time.Now()
	assingStudentToTeacher.Active = true

	result, err = st.collection.InsertOne(context.TODO(), &assingStudentToTeacher)
	if err != nil {
		return nil, err
	}

	return
} */
