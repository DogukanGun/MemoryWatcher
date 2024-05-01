package database

import (
	"MemoryWatcher/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitDatabase() *mongo.Database {
	handler := database.Handler{
		Url:      "mongodb://localhost:27017",
		Database: "test",
	}
	return handler.ConnectToDatabase()
}
