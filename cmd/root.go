package cmd

import (
	"MemoryWatcher/logger"
	"MemoryWatcher/watcher"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "info",
	Short: "Program is written by DogukanGun",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
}

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Start docker container watcher",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.Setenv(watcher.DOCKER_KEY, "true")
		if err != nil {
			logger.LogError(logger.LogErrorStruct{Message: fmt.Sprintf("Error happened in docker run: %v", err.Error())})
			return
		}
		watcher.WatchDockerContainers()
	},
}

func Execute() {
	rootCmd.AddCommand(dockerCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
