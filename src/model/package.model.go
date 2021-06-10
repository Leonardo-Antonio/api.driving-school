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
	pack struct {
		collection *mongo.Collection
	}
	IPackage interface {
		Insert(pack entity.Package) (result *mongo.InsertOneResult, err error)
		FindAll() (packages []entity.Package, err error)
		FindById(Id primitive.ObjectID) (pack entity.Package, err error)
		FindByName(name string) (pack entity.Package, err error)
		Update(pack entity.Package) (result *mongo.UpdateResult, err error)
		RemoveById(Id primitive.ObjectID) (result *mongo.UpdateResult, err error)
	}
)

func NewPackage(db *mongo.Database) *pack {
	return &pack{
		collection: db.Collection(dbutil.CollectionPackages),
	}
}

func (p *pack) Insert(pack entity.Package) (result *mongo.InsertOneResult, err error) {
	pack.ID = primitive.NewObjectID()
	pack.CreatedAt = time.Now()
	pack.Active = true

	result, err = p.collection.InsertOne(context.TODO(), &pack)
	if err != nil {
		return nil, err
	}

	return
}

func (p *pack) FindAll() (packages []entity.Package, err error) {
	filter := bson.M{
		"active": true,
	}
	cursor, err := p.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &packages); err != nil {
		return nil, err
	}

	return
}

func (p *pack) FindById(Id primitive.ObjectID) (pack entity.Package, err error) {
	filter := bson.M{
		"_id":    Id,
		"active": true,
	}
	if err := p.collection.FindOne(context.TODO(), filter).Decode(&pack); err != nil {
		return pack, err
	}
	return
}

func (p *pack) FindByName(name string) (pack entity.Package, err error) {
	filter := bson.M{
		"name":   name,
		"active": true,
	}
	if err := p.collection.FindOne(context.TODO(), filter).Decode(&pack); err != nil {
		return pack, err
	}
	return
}

func (p *pack) Update(pack entity.Package) (result *mongo.UpdateResult, err error) {
	pack.UpdatedAt = time.Now()
	update := bson.M{
		"$set": pack,
	}

	result, err = p.collection.UpdateByID(
		context.TODO(),
		pack.ID,
		update,
	)

	if err != nil {
		return nil, err
	}
	return
}

func (p *pack) RemoveById(Id primitive.ObjectID) (result *mongo.UpdateResult, err error) {
	update := bson.M{
		"$set": bson.M{
			"active": false,
		},
	}

	result, err = p.collection.UpdateByID(
		context.TODO(),
		Id,
		update,
	)

	if err != nil {
		return nil, err
	}
	return
}
