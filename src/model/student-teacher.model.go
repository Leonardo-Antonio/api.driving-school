package model

import (
	"context"
	"time"

	"github.com/Leonardo-Antonio/api.driving-school/src/dbutil"
	"github.com/Leonardo-Antonio/api.driving-school/src/entity"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils"
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
		findById(ID primitive.ObjectID) (user entity.User, err error)
		FindByTurn(turn string) (studentTeacher entity.StudentTeacher, err error)
		AssingStudentToTeacher(assingStudentToTeacher entity.AssignStudentTeacher) (result *mongo.InsertOneResult, err error)
		StudentsByTeacher(id primitive.ObjectID) (teacherStudents entity.TeacherStudents, err error)
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

func (st *studentTeacher) findById(ID primitive.ObjectID) (user entity.User, err error) {
	filter := bson.M{
		"_id":    ID,
		"active": true, /*agrefar el estado*/
	}
	// ademas creo q se debe tener un campoc en la collection users como userAsignetd
	if st.collectionUser.FindOne(context.TODO(), filter).Decode(&user) != nil {
		return
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

	student, err := st.findById(assingStudentToTeacher.IdClient)
	teacher, err := st.findById(assingStudentToTeacher.IdTeacher)
	if err != nil {
		return nil, err
	}

	if student.Turn != teacher.Turn {
		return nil, utils.ErrAssignTurn
	}

	assingStudentToTeacher.ID = primitive.NewObjectID()
	assingStudentToTeacher.CreatedAt = time.Now()
	assingStudentToTeacher.Active = true

	result, err = st.collection.InsertOne(context.TODO(), &assingStudentToTeacher)
	if err != nil {
		return nil, err
	}

	return
}

func (st *studentTeacher) StudentsByTeacher(id primitive.ObjectID) (
	teacherStudents entity.TeacherStudents,
	err error,
) {
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"id_teacher": id,
				"active":     true,
			},
		},
		{
			"$lookup": bson.M{
				"from":         dbutil.CollectionUsers,
				"localField":   "id_client",
				"foreignField": "_id",
				"as":           "students",
			},
		},
	}

	teacher, err := st.findById(id)
	if err != nil {
		return teacherStudents, err
	}

	cursor, err := st.collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return
	}

	var response []entity.TeacherStudents
	if err := cursor.All(context.TODO(), &response); err != nil {
		return teacherStudents, err
	}
	for _, value := range response {
		for _, student := range value.Students {
			teacherStudents.Students = append(teacherStudents.Students, student)
		}
	}
	teacherStudents.Teacher = teacher
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

/* pipeline := []bson.M{
	{
		"$match": bson.M{
			"id_teacher": id,
			"active":     true,
		},
	},
	{
		"lookup": bson.M{
			"from":         dbutil.CollectionUsers,
			"localField":   "id_client",
			"foreignField": "_id",
			"as":           "students",
		},
	},
} */
/*
pipeline := mongo.Pipeline{
	{
		bson.M{
			primitive.E{
				Key: "$match",
				Value: bson.D{
					primitive.E{
						Key:   "id_teacher",
						Value: id,
					},
					primitive.E{
						Key:   "active",
						Value: true,
					},
				},
			},
		},

		primitive.E{
			Key: "$lookup",
			Value: bson.D{
				primitive.E{
					Key:   "from",
					Value: dbutil.CollectionUsers,
				},
				primitive.E{
					Key:   "localField",
					Value: "id_client",
				},
				primitive.E{
					Key:   "foreignField",
					Value: "_id",
				},
				primitive.E{
					Key:   "as",
					Value: "students",
				},
			},
		},
	},
} */
