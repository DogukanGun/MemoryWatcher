package database

import (
	"MemoryWatcher/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Select struct {
	Connection        *mongo.Database
	Collection        string
	TotalItem         int
	Data              interface{}
	FilterFields      []string
	FilterFieldsValue []interface{}
}

func (sl *Select) Get() []map[string]interface{} {
	filter := bson.M{}
	for i, field := range sl.FilterFields {
		filter[field] = sl.FilterFieldsValue[i]
	}
	findOptions := options.Find()
	if sl.TotalItem != -1 {
		findOptions.SetLimit(int64(sl.TotalItem))
	}
	cursor, err := sl.Connection.Collection(sl.Collection).Find(context.Background(), filter, findOptions)
	if err != nil {
		logger.LogError(logger.LogErrorStruct{Message: err.Error()})
	}
	// Loop through the cursor and decode each document into a map
	var resultArr []map[string]interface{}
	for cursor.Next(context.Background()) {
		var result bson.M // This will be a map[string]interface{}
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		resultArr = append(resultArr, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return resultArr
}
