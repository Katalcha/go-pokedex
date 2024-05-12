package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Katalcha/go-pokedex/internal/commands"
)

func cleanUserInput(userInput string) []string {
	lowered := strings.ToLower(userInput)
	loweredWords := strings.Fields(lowered)
	return loweredWords
}

func loadPokedex(cfg *commands.Config) {
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

		command, exists := commands.GetCommands()[commandName]
		if exists {
			err := command.Callback(cfg, args...)
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
