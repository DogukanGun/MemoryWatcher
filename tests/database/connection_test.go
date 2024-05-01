package database

import (
	"MemoryWatcher/database"
	"testing"
)

func TestConnection(t *testing.T) {
	handler := database.Handler{
		Url:      "mongodb://localhost:27017",
		Database: "test",
	}
	handler.ConnectToDatabase()
}
