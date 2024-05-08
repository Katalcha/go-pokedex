package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Katalcha/go-pokedex/internal/pokeApi"
)

type config struct {
	pokeApiClient        pokeApi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

type pokedexCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]pokedexCommand {
	commands := map[string]pokedexCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"explore": {
			name:        "explore",
			description: "Searches <location_name> for Pokemon",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Displays the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of locations",
			callback:    commandMapb,
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

func loadPokedex(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanUserInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
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
