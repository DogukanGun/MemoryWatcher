package database

import (
	"MemoryWatcher/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gotest.tools/v3/assert"
	"testing"
)

func TestInsertDataWithAutoGenerateID(t *testing.T) {
	db := InitDatabase()
	insert := database.Insert{
		Connection: db,
		Collection: "InsertTest",
		Data: map[string]interface{}{
			"Model": "BMW 120",
		},
		IdAutoGenerate: true,
	}
	id := insert.InsertData()
	assert.Assert(t, id != "")
}

func TestInsertDataWithoutAutoGenerateID(t *testing.T) {
	db := InitDatabase()
	id := primitive.NewObjectID()
	insert := database.Insert{
		Connection: db,
		Collection: "InsertTest",
		Data: map[string]interface{}{
			"Model": "BMW 1",
			"_id":   id,
		},
		IdAutoGenerate: false,
	}
	returnedID := insert.InsertData()
	assert.Assert(t, returnedID == id.Hex())
}
