package database

import (
	"MemoryWatcher/logger"
	"bytes"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os/exec"
	"runtime"
	"time"
)

func Install() {
	// Check if CouchDB is installed
	if !isMongoDBInstalled() {
		err := installMongoDB()
		if err != nil {
			logger.LogError(logger.LogErrorStruct{Message: fmt.Sprintf("Error installing MongoDB: %v", err)})
		} else {
			fmt.Println("Mongo DB installed successfully.")
		}
	} else {
		fmt.Println("Mongo DB is already installed.")
		startMongoDB()
	}
}

func isMongoDBInstalled() bool {
	// Check if the command exists
	cmd := exec.Command("command", "-v", "mongod")
	err := cmd.Run()
	return err == nil
}

func installMongoDB() error {
	// Install CouchDB (commands may vary depending on the OS)
	switch runtime.GOOS {
	case "linux":
		return installForLinux()
	case "darwin":
		return installForMac()
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

func startMongoDB() {
	// Start CouchDB (commands may vary depending on the OS)
	switch runtime.GOOS {
	case "linux":
		startDbForLinux()
	case "darwin":
		startDbForMac()
	default:
		fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

func installForMac() error {
	cmd := exec.Command("brew", "install", "mongodb")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		logger.LogError(logger.LogErrorStruct{Message: fmt.Sprint(err) + ": " + stderr.String()})
		return err
	}
	fmt.Println("Result: " + out.String())
	return nil
}

func installForLinux() error {
	cmd := exec.Command("sudo", "apt-get", "update")
	err := cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("sudo", "apt-get", "install", "-y", "mongodb")
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// starts couch db for mac
func startDbForMac() {
	if !isMongoDBWorking() {
		// Start Mongo DB using brew services
		cmdStart := exec.Command("brew", "services", "start", "mongodb-community")
		if err := cmdStart.Run(); err != nil {
			fmt.Println("Error starting Mongo DB:", err)
			return
		}
		fmt.Println("Mongo DB started successfully.")
	}

}

// starts couch db for linux
func startDbForLinux() {
	if !isMongoDBWorking() {
		// Start Mongo DB using brew services
		cmdStart := exec.Command("sudo", "systemctl", "start", "mongodb")
		if err := cmdStart.Run(); err != nil {
			fmt.Println("Error starting Mongo DB:", err)
			return
		}
		fmt.Println("Mongo DB started successfully.")
	}

}

func isMongoDBWorking() bool {
	mongoDBURL := "mongodb://localhost:27017"

	// Create a MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDBURL))
	if err != nil {
		fmt.Printf("Error creating MongoDB client: %v\n", err)
		return false
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v\n", err)
		return false
	}
	defer client.Disconnect(ctx)

	// Ping MongoDB server
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Printf("Error pinging MongoDB server: %v\n", err)
		return false
	}

	fmt.Println("MongoDB is working!")
	return true
}
