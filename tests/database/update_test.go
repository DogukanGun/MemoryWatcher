package database

import (
	"MemoryWatcher/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gotest.tools/v3/assert"
	"testing"
)

func TestUpdateOneDocument(t *testing.T) {
	db := InitDatabase()
	model := "BMW 120"
	insert := database.Insert{
		Connection: db,
		Collection: "UpdateTest",
		Data: map[string]interface{}{
			"Model": "BMW 1",
		},
		IdAutoGenerate: true,
	}
	id := insert.InsertData()
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// Handle error if the string is not a valid ObjectID
		panic(err)
	}
	oidSlice := []interface{}{oid}
	modelSlice := []interface{}{model}
	update := database.Update{
		Connection:         db,
		Collection:         "UpdateTest",
		FilterFields:       []string{"_id"},
		FilterFieldsValue:  oidSlice,
		UpdatedFields:      []string{"Model"},
		UpdatedFieldsValue: modelSlice,
	}
	isUpdated := update.UpdateDocument()
	assert.Assert(t, isUpdated)
	get := database.Select{
		Connection:        db,
		Collection:        "UpdateTest",
		TotalItem:         -1,
		Data:              nil,
		FilterFields:      []string{"_id"},
		FilterFieldsValue: oidSlice,
	}
	models := get.Get()
	assert.Assert(t, models[0]["Model"] == model)
}
