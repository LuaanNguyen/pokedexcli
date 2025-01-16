package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(cfg *config) error {
	url := cfg.Previous 
	
	if url == "" {
		fmt.Println("You are on the first page.")
		return nil 
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching previous Podedex locations: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response's body: %w", err)
	}
	defer res.Body.Close() 

	var locations LocationResponse
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return fmt.Errorf("error unmarshaling response: %w", err)
	}

	if len(locations.Results) > 0 {
		for _, location := range locations.Results {
			fmt.Printf("%s\n", location.Name)
		}	
	}

	// Update the config for pagination
    cfg.Next = ""
    if locations.Next != nil {
        cfg.Next = *locations.Next // Set the Next URL if it exists
    }
    cfg.Previous = ""
    if locations.Previous != nil {
        cfg.Previous = *locations.Previous // Set the Previous URL if it exists
    }
	
	return nil
}