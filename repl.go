package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func startREPL() {
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

		err := command.callback()
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
