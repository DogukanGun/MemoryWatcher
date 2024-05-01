package database

import (
	"MemoryWatcher/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Update struct {
	Connection         *mongo.Database
	Collection         string
	FilterFields       []string
	FilterFieldsValue  []interface{}
	UpdatedFields      []string
	UpdatedFieldsValue []interface{}
}

func (ud *Update) UpdateDocument() bool {
	filter := bson.M{}
	for i, field := range ud.FilterFields {
		filter[field] = ud.FilterFieldsValue[i]
	}
	updateBody := bson.M{}
	for i, field := range ud.UpdatedFields {
		updateBody[field] = ud.UpdatedFieldsValue[i]
	}
	update := bson.D{{"$set",
		updateBody,
	}}
	updateResult, err := ud.Connection.Collection(ud.Collection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.LogError(logger.LogErrorStruct{Message: err.Error()})
	}
	return updateResult.ModifiedCount == 1
}

func (ud *Update) UpdateDocuments() int64 {
	filter := bson.M{}
	for i, field := range ud.FilterFields {
		filter[field] = ud.FilterFieldsValue[i]
	}
	updateBody := bson.M{}
	for i, field := range ud.UpdatedFields {
		updateBody[field] = ud.UpdatedFieldsValue[i]
	}
	update := bson.D{{"$set",
		updateBody,
	}}
	updateResult, err := ud.Connection.Collection(ud.Collection).UpdateMany(context.Background(), filter, update)
	if err != nil {
		logger.LogError(logger.LogErrorStruct{Message: err.Error()})
	}
	return updateResult.ModifiedCount
}
