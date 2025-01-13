package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationResponse struct {
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location/")
	if err != nil {
		return fmt.Errorf("error fetching Pokedex locations: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}
	defer res.Body.Close()

	
	var locations LocationResponse
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return fmt.Errorf("error unmarshaling response: %w", err)
	}
	
	//display all locations 
	for _, location := range locations.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}