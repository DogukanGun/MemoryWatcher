package database

import (
	"MemoryWatcher/database"
	"gotest.tools/v3/assert"
	"testing"
)

func TestSelectData(t *testing.T) {
	db := InitDatabase()
	model := "BMW 1"
	insert := database.Insert{
		Connection: db,
		Collection: "GetTest",
		Data: map[string]interface{}{
			"Model": model,
		},
		IdAutoGenerate: true,
	}
	insert.InsertData()
	get := database.Select{
		Connection: db,
		Collection: "GetTest",
		Data:       nil,
	}
	data := get.Get()
	assert.Assert(t, data[0]["Model"].(string) == model)
}
