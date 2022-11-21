package mongodb

import (
	"context"
	"fmt"
	"food_mlem/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Mongodb MongoInstance

// Source: https://github.com/mongodb/mongo-go-driver
func init() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(conf.Config.MONGODB_CONFIG.ENDPOINT).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(conf.Config.MONGODB_CONFIG.DATABASE)
	Mongodb = MongoInstance{
		Client: client,
		Db:     db,
	}

	fmt.Println("CONNECT MONGODB SUCCESS!")
}
