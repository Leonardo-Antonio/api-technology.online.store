package dbutil

import (
	"context"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Connection() (db *mongo.Database) {
	clientOptions := options.Client().ApplyURI(utils.Config().MongoUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.TODO(), nil) ;err != nil {
		log.Fatal(err)
	}

	return client.Database(utils.Config().DBName)
}
