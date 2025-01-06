package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		//Use fmt.Print to print the prompt Pokedex > without a newline character
		fmt.Print("Pokedex > ")
		scanner.Scan()

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

//This function splits the users' input into "words" based on white space.
//It should also lowercase the input and trim any leading or trailing whitespace
func cleanInput(text string) []string {
	//Lower case the input string 
	loweredText := strings.ToLower(text)

	//Trim all white splace and split the input into words
	return strings.Fields(loweredText)
}