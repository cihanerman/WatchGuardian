package main

import (
	"bufio"
	"github.com/cihanerman/WatchGuardian/utils"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	defer log.Println("WatchGuardian is closing...")
	reader := bufio.NewReader(os.Stdin)

	// read file path from command line
	log.Println("Enter file path: ")

	filePath, err := reader.ReadString('\n')
	utils.CheckError(err)
	filePath = strings.TrimSuffix(filePath, "\n")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatal("File does not exist")
	}

	// read post change data url from command line
	log.Println("Enter post url: ")
	postUrl, err := reader.ReadString('\n')
	utils.CheckError(err)
	postUrl = strings.TrimSuffix(postUrl, "\n")

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
	token = strings.TrimSuffix(token, "\n")
	token = strings.Trim(token, " ")

	if token == "" {
		log.Fatal("Token is empty")
	}

	// read header variable from command line
	log.Println("Enter header variable(defaul Authorization): ")
	headerVariable, err := reader.ReadString('\n')
	utils.CheckError(err)
	headerVariable = strings.TrimSuffix(headerVariable, "\n")
	headerVariable = strings.Trim(headerVariable, " ")

	if headerVariable == "" {
		headerVariable = "Authorization"
	}

	// TODO watch the file for changes

	// TODO send post request to url with data from file
}
