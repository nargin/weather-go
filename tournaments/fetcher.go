package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Intentionally left blank
	fmt.Println("Fetcher!")

	parseEnv()

	url := "https://api.pandascore.co/lol/tournaments?sort=&search[region]=WEU&page=1&per_page=10"
	PandaScoreAPIKey := os.Getenv("PANDASCORE_API_KEY")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer "+PandaScoreAPIKey)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	body, _ := io.ReadAll(resp.Body)

	var jsonFormatted []any
	err = json.Unmarshal(body, &jsonFormatted)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	jsonPretty, err := json.MarshalIndent(jsonFormatted, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	os.WriteFile("tournaments.json", jsonPretty, 0644)
}
