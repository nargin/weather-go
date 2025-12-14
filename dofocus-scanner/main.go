package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Item struct {
	ID              int               `json:"id"`
	MongoID         string            `json:"_id"`
	Name            map[string]string `json:"name"`
	Level           int               `json:"level"`
	ImageURL        string            `json:"imageUrl"`
	Supertype       Supertype         `json:"supertype"`
	Characteristics []Characteristic  `json:"characteristics"`
}

type Supertype struct {
	ID int `json:"id"`
}

type Characteristic struct {
	ID int `json:"id"`
}

type ItemSimple struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Level             int    `json:"level"`
	ImageURL          string `json:"imageUrl"`
	SupertypeID       int    `json:"supertypeId"`
	CharacteristicIDs []int  `json:"characteristicIds"`
}

func main() {
	// Fetch API
	resp, err := http.Get("https://dofocus.fr/api/items")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading: %v\n", err)
		return
	}

	// Parse JSON
	var items []Item
	err = json.Unmarshal(body, &items)
	if err != nil {
		fmt.Printf("Error parsing: %v\n", err)
		return
	}

	fmt.Printf("Found %d items\n", len(items))

	// Transform to simple format with only French names
	var simpleItems []ItemSimple
	for _, item := range items {
		charIDs := make([]int, len(item.Characteristics))
		for i, char := range item.Characteristics {
			charIDs[i] = char.ID
		}

		simpleItems = append(simpleItems, ItemSimple{
			ID:                item.ID,
			Name:              item.Name["fr"], // Just French name
			Level:             item.Level,
			ImageURL:          item.ImageURL,
			SupertypeID:       item.Supertype.ID,
			CharacteristicIDs: charIDs,
		})
	}

	// Save to JSON file
	file, err := os.Create("items.json")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Marshal with indentation
	jsonData, err := json.MarshalIndent(simpleItems, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		return
	}

	file.Write(jsonData)

	fmt.Printf("Saved %d items to items.json\n", len(simpleItems))
}
