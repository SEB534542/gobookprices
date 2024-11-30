package gobookprices

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// getWorks takes an ISBN, looks up the edition on openlibrary and returns the work.
func getWorks(isbn string) (string, error) {
	url := fmt.Sprintf("https://openlibrary.org/isbn/%s.json", isbn)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Works []struct {
			Key string `json:"key"`
		} `json:"works"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.Works[0].Key, nil
	// TODO: only first work is returned, check if necessary to return a list of all works
}
