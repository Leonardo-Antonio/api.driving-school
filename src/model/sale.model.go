package model

import (
	"context"
	"time"

	"github.com/Leonardo-Antonio/api.driving-school/src/dbutil"
	"github.com/Leonardo-Antonio/api.driving-school/src/entity"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	sale struct {
		collection     *mongo.Collection
		collectionUser *mongo.Collection
	}
	ISale interface {
		changeStateSaleUser(id primitive.ObjectID, turn string, state bool) error
		Buy(sale entity.Sale) (*mongo.InsertOneResult, error)
	}
)

func NewSale(db *mongo.Database) *sale {
	return &sale{
		collection:     db.Collection(dbutil.CollectionSales),
		collectionUser: db.Collection(dbutil.CollectionUsers),
	}
}

func (s *sale) changeStateSaleUser(id primitive.ObjectID, turn string, state bool) error {
	update := bson.M{
		"$set": bson.M{
			"_id":         id,
			"active":      true,
			"turn":        turn,
			"status_sale": state,
		},
	}
	result, err := s.collectionUser.UpdateByID(context.TODO(), id, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return utils.ErrNotUpdated
	}
	return nil
}

func (s *sale) Buy(sale entity.Sale) (*mongo.InsertOneResult, error) {
	sale.ID = primitive.NewObjectID()
	sale.CreatedAt = time.Now()
	sale.Active = true
	result, err := s.collection.InsertOne(context.TODO(), &sale)
	if err != nil {
		return nil, err
	}

	if err := s.changeStateSaleUser(
		sale.IdClient,
		sale.Turn,
		true,
	); err != nil {
		return nil, err
	}
	return result, err
}
