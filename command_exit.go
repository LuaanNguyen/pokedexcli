package main

import (
	"fmt"
	"os"
)

//Call back func for exit command
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
