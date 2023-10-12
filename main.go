package main

import (
	"bufio"
	"github.com/cihanerman/WatchGuardian/utils"
	"github.com/cihanerman/WatchGuardian/watchers"
	"log"
	"net/url"
	"os"
)

func main() {
	defer log.Println("WatchGuardian is closing...")
	reader := bufio.NewReader(os.Stdin)

	// read file path from command line
	log.Println("Enter file path: ")

	filePath, err := reader.ReadString('\n')
	utils.CheckError(err)
	filePath = utils.TrimInput(filePath)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatal("File does not exist")
	}

	// read post change data url from command line
	log.Println("Enter post url: ")
	postUrl, err := reader.ReadString('\n')
	utils.CheckError(err)
	postUrl = utils.TrimInput(postUrl)

	// parse url
	parsedURL, err := url.Parse(postUrl)
	utils.CheckError(err)

	// check if url is valid
	if parsedURL.Scheme == "" && parsedURL.Host == "" {
		log.Fatal("Invalid url")
	}

	//  read token for url from command line
	log.Println("Enter token: ")
	token, err := reader.ReadString('\n')
	utils.CheckError(err)
	token = utils.TrimInput(token)

	// read header variable from command line
	log.Println("Enter header variable(default Authorization): ")
	headerVariable, err := reader.ReadString('\n')
	utils.CheckError(err)
	headerVariable = utils.TrimInput(headerVariable)

	if headerVariable == "" {
		headerVariable = "Authorization"
	}

	// watch the file for changes
	// open file
	file, err := os.Open(filePath)
	utils.CheckError(err)
	defer file.Close()

	fileReader := bufio.NewReader(file)

	// Create new watchers.
	watcher := watchers.CreateWatcher()
	defer watcher.Close()

	closeWatcher := make(chan struct{})

	// Start listening for events.
	log.Println("File monitoring started...")
	go watchers.StartWatcher(watcher, closeWatcher, fileReader, postUrl, headerVariable, token)

	// Add a path.
	watchers.AddFileToWatcher(watcher, filePath)

	// Block main goroutine until error.
	<-closeWatcher

}
