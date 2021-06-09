package dbutil

import (
	"context"
	"log"

	"github.com/Leonardo-Antonio/api.driving-school/src/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection() (db *mongo.Database) {
	clientOptions := options.Client().ApplyURI(utils.Config().MongoUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln(err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalln(err)
	}
	db = client.Database(utils.Config().NameDataBase)
	createIndexUser(db)
	return
}
