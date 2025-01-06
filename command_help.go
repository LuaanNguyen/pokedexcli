package main

import "fmt"

//Call back func for help command
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	// List available commands
	for command, properties := range getCommands(){
		fmt.Printf("%s: %s\n", command, properties.description)
	}

	return nil
}