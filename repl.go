package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/firstsano/pokedex/internal/pokeapi"
)

type configuration struct {
	pokeapiClient pokeapi.Client
	nextLocation  *string
	prevLocation  *string
}

func startREPL(cfg *configuration) {
	scanner := bufio.NewScanner(os.Stdin)
	registry := getCommands()

	for {
		fmt.Print("Pokedex > ")
		if scanned := scanner.Scan(); !scanned {
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, ok := registry[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)

	return strings.Fields(text)
}

type cliCommand struct {
	name        string
	description string
	callback    func(*configuration) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
