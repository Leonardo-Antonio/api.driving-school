package model

import (
	"context"
	"time"

	"github.com/Leonardo-Antonio/api.driving-school/src/dbutil"
	"github.com/Leonardo-Antonio/api.driving-school/src/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	user struct {
		collection *mongo.Collection
	}

	IUser interface {
		Insert(entity entity.User) (result *mongo.InsertOneResult, err error)
		FindAll() (users []entity.User, err error)
		FindByDNI(logIn entity.User) (user entity.User, err error)
		FindByEmail(logIn entity.User) (user entity.User, err error)
		Update(user entity.User) (result *mongo.UpdateResult, err error)
		Remove(ID primitive.ObjectID) (result *mongo.UpdateResult, err error)
	}
)

func NewUser(db *mongo.Database) *user {
	return &user{
		collection: db.Collection(dbutil.CollectionUsers),
	}
}

func (u *user) Insert(entity entity.User) (result *mongo.InsertOneResult, err error) {
	entity.ID = primitive.NewObjectID()
	entity.Active = true
	entity.CreatedAt = time.Now()
	entity.StatusSale = false // cambiarlo porq no debe crear como true

	result, err = u.collection.InsertOne(context.TODO(), &entity)
	if err != nil {
		return nil, err
	}
	return
}

func (u *user) FindAll() (users []entity.User, err error) {
	filter := bson.M{
		"active": true,
	}
	cursor, err := u.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return
}

func (u *user) FindByDNI(logIn entity.User) (user entity.User, err error) {
	filter := bson.M{
		"dni":      logIn.DNI,
		"password": logIn.Password,
		"active":   true,
	}

	if u.collection.FindOne(context.TODO(), filter).Decode(&user) != nil {
		return
	}

	return
}

func (u *user) FindByEmail(logIn entity.User) (user entity.User, err error) {
	filter := bson.M{
		"email":    logIn.Email,
		"password": logIn.Password,
	}

	if u.collection.FindOne(context.TODO(), filter).Decode(&user) != nil {
		return
	}

	return
}

func (u *user) FindById(ID primitive.ObjectID) (user entity.User, err error) {
	filter := bson.M{
		"_id":    ID,
		"active": true,
	}

	if u.collection.FindOne(context.TODO(), filter).Decode(&user) != nil {
		return
	}

	return
}

func (u *user) Update(user entity.User) (result *mongo.UpdateResult, err error) {
	user.UpdatedAt = time.Now()
	udpate := bson.M{
		"$set": user,
	}

	result, err = u.collection.UpdateByID(
		context.TODO(),
		user.ID,
		udpate,
	)
	if err != nil {
		return nil, err
	}

	return
}

func (u *user) Remove(ID primitive.ObjectID) (result *mongo.UpdateResult, err error) {
	delete := bson.M{
		"$set": bson.M{
			"active": false,
		},
	}
	result, err = u.collection.UpdateByID(context.TODO(), ID, delete)
	if err != nil {
		return nil, err
	}

	return
}
