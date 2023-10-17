package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ravibhagw/xbs_adv_stats_job/dal"
	"github.com/ravibhagw/xbs_adv_stats_job/models"
)

var teamIdMap = map[string]int{
	"xbs team 10": 3301,
	"xbs team 4":  1428,
	"xbs team 3":  5091,
	"xbs team1":   7560, // Someone can't type lol
	"xbs team 1":  7560,
	"xbs team 5":  8516,
	"xbs team 8":  9529,
	"xbs team 6":  11090,
	"xbs team 7":  14824,
}

func main() {

	matches := make([]models.Match, 0)
	for _, clubID := range teamIdMap {
		var foundMatches, err = fetchClubMatches(clubID)
		if err != nil {
			panic(err)
		}
		matches = append(matches, foundMatches...)
	}

	connStr := "lel"
	client, err := dal.NewClient(connStr)
	if err != nil {
		panic(err)
	}

	client.SaveMatches(matches)

	fmt.Printf("done")
}

func fetchClubMatches(clubID int) ([]models.Match, error) {
	url := "https://proclubs.ea.com/api/nhl/clubs/matches?matchType=gameType5&platform=common-gen5&clubIds=" + fmt.Sprint(clubID)

	// Create an HTTP client with a timeout of 10 seconds
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create an HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set the User-Agent header to an empty string
	req.Header.Set("User-Agent", "")

	// Perform the HTTP GET request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the response status code is OK (200)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response into a slice of Root structs
	var data []models.Match
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
