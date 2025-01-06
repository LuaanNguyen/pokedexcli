package main

import (
	"fmt"
	"strings"
)

func main () {
	fmt.Println("Hello, World!")
}


//This function splits the users' input into "words" based on white space.
//It should also lowercase the input and trim any leading or trailing whitespace 
func cleanInput(text string) []string {
	//Lower case the input string 
	loweredText := strings.ToLower(text)

	//Trim all white splace and split the input into words
	return strings.Fields(loweredText)
}