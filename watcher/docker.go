package watcher

import (
	"MemoryWatcher/logger"
	"MemoryWatcher/utils"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"os/user"
	"time"
)

const (
	DOCKER_KEY = "DOCKER_WATCH"
)

func WatchDockerContainers() {
	ctx := context.Background()

	// Set up file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	defer logger.LogInfo(logger.LogInfoStruct{Message: "Docker watch has been ended."})

	// Set up Docker client
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	// Watch for file events
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				// Get current time
				currentTime := time.Now().Format("2006-01-02 15:04:05")
				// Get user who made the operation
				currentUser, _ := user.Current()

				var operationSummary string
				if event.Op&fsnotify.Create == fsnotify.Create {
					operationSummary = fmt.Sprintf("File created: %s", event.Name)
					fi, err := os.Stat(event.Name)
					if err != nil {
						log.Println("Error getting file info:", err)
					} else {
						log.Println("File modified:", event.Name)
						log.Println("File size:", fi.Size(), "bytes")
					}
				} else if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					operationSummary = fmt.Sprintf("File permission changed: %s", event.Name)
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					fi, err := os.Stat(event.Name)
					if err != nil {
						log.Println("Error getting file info:", err)
					} else {
						log.Println("File modified:", event.Name)
						log.Println("File size:", fi.Size(), "bytes")
					}
					operationSummary = fmt.Sprintf("File removed: %s", event.Name)
				} else if event.Op&fsnotify.Write == fsnotify.Write {
					operationSummary = fmt.Sprintf("File changed: %s", event.Name)
					fi, err := os.Stat(event.Name)
					if err != nil {
						log.Println("Error getting file info:", err)
					} else {
						log.Println("File modified:", event.Name)
						log.Println("File size:", fi.Size(), "bytes")
					}
				}
				logger.LogInfo(logger.LogInfoStruct{Message: "Time: " + currentTime + "\nUser: " + currentUser.Username + "\nOperation " + operationSummary})
				if err != nil {
					log.Println("Error sending email:", err)
				}
			case err := <-watcher.Errors:
				log.Println("Error watching files:", err)
			}
		}
	}()

	// Watch current directory for file events

	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}

	// Watch for Docker events
	go func() {
		dockerEvents, dockerErr := dockerClient.Events(ctx, types.EventsOptions{})
		for {
			select {
			case event := <-dockerEvents:
				if event.Action == "create" && event.Type == "container" {
					// Container creation event occurred
					log.Println("Container created:", event.ID)
					// Get current time
					currentTime := time.Now().Format("2006-01-02 15:04:05")
					// Get user who made the operation
					currentUser, _ := user.Current()
					// Send email notification
					logger.LogInfo(logger.LogInfoStruct{Message: "Time: " + currentTime + "\nUser: " + currentUser.Username + "\nA new Docker container was created: " + event.ID})
					/*err := sendEmail("Container Created", "Time: "+currentTime+"\nUser: "+currentUser.Username+"\nA new Docker container was created: "+event.ID)
					if err != nil {
						log.Println("Error sending email:", err)
					}*/
					utils.Send("Container Created" + " \nTime: " + currentTime + "\nUser: " + currentUser.Username + "\nA new Docker container was created: " + event.ID)
				}
			case eventErr := <-dockerErr:
				fmt.Printf("error from docker client: %s", eventErr)
			}
		}
	}()

	for os.Getenv(DOCKER_KEY) == "true" {
	}
}
