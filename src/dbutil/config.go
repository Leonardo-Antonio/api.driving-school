package dbutil

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createIndexUser(db *mongo.Database) {
	_, err := db.Collection(CollectionUsers).
		Indexes().
		CreateOne(
			context.TODO(),
			mongo.IndexModel{
				Keys: bson.D{
					{Key: "email", Value: 1},
					{Key: "dni", Value: 1},
				},
				Options: options.Index().SetUnique(true),
			},
		)
	if err != nil {
		log.Fatalln(err)
	}
}
