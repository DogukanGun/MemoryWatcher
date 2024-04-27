package database

import (
	"MemoryWatcher/logger"
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Install() {
	// Check if CouchDB is installed
	if !isCouchDBInstalled() {
		err := installCouchDB()
		if err != nil {
			logger.LogError(logger.LogErrorStruct{Message: fmt.Sprintf("Error installing CouchDB: %v", err)})
		} else {
			fmt.Println("CouchDB installed successfully.")
		}
	} else {
		fmt.Println("CouchDB is already installed.")
		startCouchDB()
	}
}

func isCouchDBInstalled() bool {
	// Check if the command exists
	cmd := exec.Command("command", "-v", "couchdb")
	err := cmd.Run()
	return err == nil
}

func installCouchDB() error {
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

func startCouchDB() {
	// Start CouchDB (commands may vary depending on the OS)
	switch runtime.GOOS {
	case "linux":
		installForLinux()
	case "darwin":
		startDbForMac()
	default:
		fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

func installForMac() error {
	cmd := exec.Command("brew", "install", "couchdb")
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
	cmd = exec.Command("sudo", "apt-get", "install", "-y", "couchdb")
	err = cmd.Run()
	if err != nil {
		return err
	}
	// Start CouchDB service
	cmd = exec.Command("sudo", "systemctl", "start", "couchdb")
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// starts couch db for mac
func startDbForMac() {
	if !isDbHasPasswordForMac() {
		cmdSetPassword := exec.Command("bash", "-c", "echo '[admins]\nadmin = memorywatcher' >> /opt/homebrew/etc/local.ini")
		if err := cmdSetPassword.Run(); err != nil {
			fmt.Println("Error setting admin password:", err)
			return
		}
	}
	if !isDbWorking() {
		// Start CouchDB using brew services
		cmdStart := exec.Command("brew", "services", "start", "couchdb")
		if err := cmdStart.Run(); err != nil {
			fmt.Println("Error starting CouchDB:", err)
			return
		}
		fmt.Println("CouchDB started successfully.")
	}

}

func isDbWorking() bool {
	couchDBURL := "http://localhost:5984"

	// Send a GET request to the CouchDB server
	resp, err := http.Get(couchDBURL)
	if err != nil {
		fmt.Printf("Error connecting to CouchDB: %v\n", err)
		return false
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode == http.StatusOK {
		fmt.Println("CouchDB is working!")
		return true
	} else {
		fmt.Printf("CouchDB is not working, status code: %d\n", resp.StatusCode)
		return false
	}
}

// checks if the couch db has credentials
func isDbHasPasswordForMac() (foundAdminLine bool) {
	// Open the file
	file, err := os.Open("/opt/homebrew/etc/local.ini")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Variables to track if we are in the [admins] section and if we found the admin line
	inAdminsSection := false
	foundAdminLine = false

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line contains [admins]
		if strings.Contains(line, "[admins]") {
			inAdminsSection = true
			fmt.Println(line)
		}

		// Check if we are in the [admins] section and if the line contains "admin = memorywatcher"
		if inAdminsSection && strings.Contains(line, "admin = memorywatcher") {
			foundAdminLine = true
			fmt.Println(line)
		}

		// Break the loop if we found the admin line
		if foundAdminLine {
			break
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	return
}
