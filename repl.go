package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	Next	string
	Previous 	string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}


func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := &config{
		Next:     "",
		Previous: "",
	}

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

		command, exists := getCommands()[firstWord]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

		//fmt.Printf("Your command was: %s\n", firstWord)
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


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name: "map",
			description: "Display the next 20 locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Display the previous 20 locations",
			callback: commandMapb,
		},
	}
}
