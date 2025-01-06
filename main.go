package main

import (
	"bufio"
	"fmt"
	"os"
)

func main () {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		//Use fmt.Print to print the prompt Pokedex > without a newline character
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		//Use the scanner's .Scan and .Text methods to get the user's input as a string
		input := scanner.Text()

		//Clean the users input by trimming any leading or trailing whitespace, and converting it to lowercase.
		stringSlice := cleanInput(input)

		// If there are no words, just continue to the next iteration
		if len(stringSlice) == 0 {
			continue
		}
		
		//Capture the first "word" of the input and use it to print
		firstWord := stringSlice[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}
}


