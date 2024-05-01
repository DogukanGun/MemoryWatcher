package database

import (
	"MemoryWatcher/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Insert struct {
	Connection     *mongo.Database
	Collection     string
	Data           map[string]interface{}
	IdAutoGenerate bool
}

func (dh *Insert) InsertData() string {
	if dh.IdAutoGenerate {
		dh.Data["_id"] = primitive.NewObjectID()
	}
	putRev, err := dh.Connection.Collection(dh.Collection).InsertOne(context.Background(), dh.Data)
	if err != nil {
		logger.LogError(logger.LogErrorStruct{Message: err.Error()})
		return ""
	}
	id := putRev.InsertedID.(primitive.ObjectID).Hex()
	logger.LogInfo(logger.LogInfoStruct{Message: id})
	return id
}
