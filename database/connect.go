package database

import (
	"context"
	"fmt"
	_ "github.com/go-kivik/kivik/v4/couchdb" // The CouchDB driver
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Handler struct {
	Url      string `json:"url"`
	Database string `json:"database"`
}

func (dh *Handler) ConnectToDatabase() *mongo.Database {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dh.Url))
	if err != nil {
		panic(err)
	}
	var result bson.M
	if err := client.Database(dh.Database).RunCommand(context.TODO(), bson.D{{"ping", 5}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	database := client.Database(dh.Database)
	print(database)
	return database
}

func (dh *Handler) ConnectToDatabaseWithOptions(dbOptions *options.ClientOptions) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dh.Url), dbOptions)
	if err != nil {
		panic(err)
	}
	return client.Database(dh.Database)
}
