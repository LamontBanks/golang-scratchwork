package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Issue struct {
	Title    string `json:"title"`
	Estimate int    `json:"estimate"`
}

func main() {
	fmt.Println(getBootDevIssues("https://api.boot.dev/v1/courses_rest_api/learn-http/issues?limit=2"))
}

func getBootDevIssues(url string) ([]Issue, error) {
	// Make the request
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	// Decode the JSON into Issue structs
	var issues []Issue
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&issues); err != nil {
		return issues, fmt.Errorf("error decoding response body: %w", err)
	}

	// Return values
	return issues, nil

}
