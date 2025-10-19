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
	fmt.Println("Fetching Worlds 2025 matches from PandaScore API...")

	parseEnv()

	PandaScoreAPIKey := os.Getenv("PANDASCORE_API_KEY")
	if PandaScoreAPIKey == "" {
		log.Fatalf("PANDASCORE_API_KEY not found in environment")
	}

	// Fetch all LoL matches in 2025 and filter for Worlds
	allMatchesURL := "https://api.pandascore.co/lol/matches?range[begin_at]=2025-01-01,2025-12-31&page=1&per_page=100&sort=-begin_at"
	fmt.Println("\nFetching LoL matches in 2025...")

	body := fetchRawData(allMatchesURL, PandaScoreAPIKey)

	var matches []Match
	err := json.Unmarshal(body, &matches)
	if err != nil {
		log.Fatalf("Error unmarshalling matches: %v", err)
	}

	// Filter for Worlds matches
	var worldsMatches []Match
	for _, match := range matches {
		if match.League.Name == "Worlds" ||
			match.Serie.Slug == "league-of-legends-world-championship-2025" ||
			match.Serie.Slug == "league-of-legends-world-championship-2025-playoffs" {
			worldsMatches = append(worldsMatches, match)
		}
	}

	fmt.Printf("\n\n=== WORLDS 2025 MATCHES ===\n")
	fmt.Printf("Found %d Worlds 2025 matches!\n\n", len(worldsMatches))

	for i, match := range worldsMatches {
		fmt.Printf("%d. %s\n", i+1, match.Name)
		fmt.Printf("   Date: %s\n", match.BeginAt)
		fmt.Printf("   Status: %s\n", match.Status)
		fmt.Printf("   Tournament: %s\n", match.Tournament.Name)
		if len(match.Opponents) >= 2 {
			fmt.Printf("   Match: %s vs %s\n",
				match.Opponents[0].Opponent.Acronym,
				match.Opponents[1].Opponent.Acronym)
		}
		fmt.Println()
	}

	// Save all Worlds matches to file
	jsonPretty, err := json.MarshalIndent(worldsMatches, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	os.WriteFile("worlds_2025_matches.json", jsonPretty, 0644)
	fmt.Println("Data saved to worlds_2025_matches.json!")
}

func fetchRawData(url string, apiKey string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer "+apiKey)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	body, _ := io.ReadAll(resp.Body)

	// Try to detect if response is an error
	if resp.StatusCode >= 400 {
		fmt.Printf("API Error: %s\n", string(body))
		return []byte{}
	}

	return body
}
