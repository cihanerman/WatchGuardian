package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func CheckError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func TrimInput(input string) string {
	input = strings.TrimSuffix(input, "\n")
	input = strings.Trim(input, " ")
	return input
}

func SendUpdate(msg, file, op, url, authKey, toke string) {
	// Prepare request body
	event := Event{
		Message:   msg,
		File:      file,
		Operation: op,
	}
	data, err := json.Marshal(event)
	CheckError(err)

	// Make HTTP POST request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	CheckError(err)
	request.Header.Set("Content-Type", "application/json")
	if toke != "" {
		request.Header.Set(authKey, toke)
	}

	// Send request
	client := &http.Client{}
	response, err := client.Do(request)
	CheckError(err)
	defer response.Body.Close()

	// Read response body
	body, err := io.ReadAll(response.Body)
	CheckError(err)

	log.Println("response status:", response.Status, "response body:", string(body), "url:", url)
}
