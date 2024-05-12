package commands

func GetCommands() map[string]pokedexCommand {
	commands := map[string]pokedexCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View details about a caught Pokemon",
			Callback:    commandInspect,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Searches a location for Pokemon",
			Callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Displays the next page of locations",
			Callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of locations",
			Callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See all the pokemon you've caught",
			Callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			Callback:    commandExit,
		},
	}
	return commands
}
