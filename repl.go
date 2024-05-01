package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pokedexCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]pokedexCommand {
	commands := map[string]pokedexCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}
	return commands
}

func cleanUserInput(userInput string) []string {
	lowered := strings.ToLower(userInput)
	loweredWords := strings.Fields(lowered)
	return loweredWords
}

func loadPokedex() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanUserInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
