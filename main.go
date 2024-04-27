package main

import (
	"MemoryWatcher/logger"
	"MemoryWatcher/watcher"
	"fmt"
	"os"
)

/*func main() {
	logger.Init()
	database.Install()
	cmd.Execute()
}*/

func main() {
	err := os.Setenv(watcher.DOCKER_KEY, "true")
	if err != nil {
		logger.LogError(logger.LogErrorStruct{Message: fmt.Sprintf("Error happened in docker run: %v", err.Error())})
		return
	}
	watcher.WatchDockerContainers()
}
