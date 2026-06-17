package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/onurbilginnn/pokecache"
)

const initialLocationURL = "https://pokeapi.co/api/v2/location-area/?limit=20"

func main() {
	cacheInterval := 5 * time.Minute
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Pokedex help",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show 20 location areas in Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapback",
			description: "Show previous 20 location areas in Pokemon world",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area in Pokemon world",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokemon",
			callback:    commandPokedex,
		},
	}
	state := state{
		Next:     initialLocationURL,
		Previous: initialLocationURL,
		Pokedex:  pokedex{pokemon: make(map[string]Pokemon)},
	}
	scanner := bufio.NewScanner(os.Stdin)
	pokeCache := pokecache.NewCache(cacheInterval)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputText := scanner.Text()
		command := strings.ToLower(cleanInput(inputText)[0])
		location := ""
		if len(cleanInput(inputText)) > 1 {
			location = cleanInput(inputText)[1]
		}
		if _, isCommandExists := commands[command]; !isCommandExists {
			fmt.Println("Unknown command")
		} else {
			err := commands[command].callback(&state, pokeCache, location)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Command error: %v\n", err)
			}
		}
		inputText = ""
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Scanner error: %v\n", err)
			os.Exit(1)
		}
	}
}
