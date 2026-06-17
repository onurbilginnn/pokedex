package main

import (
	"fmt"
	"os"

	"math/rand/v2"

	"github.com/onurbilginnn/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(state *state, cache *pokecache.Cache, location string) error
}

type pokedex struct {
	pokemon map[string]Pokemon
}

type state struct {
	Next     string
	Previous string
	Pokedex  pokedex
}

func commandExit(state *state, cache *pokecache.Cache, location string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(state *state, cache *pokecache.Cache, location string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("map: Show 20 location areas in Pokemon world")
	fmt.Println("mapb: Show previous 20 location areas in Pokemon world")
	fmt.Println("explore: Explore a location area in Pokemon world")
	fmt.Println("catch: Catch a Pokemon")
	fmt.Println("inspect: Inspect a caught Pokemon")
	fmt.Println("pokedex: List all caught Pokemon")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandMap(state *state, cache *pokecache.Cache, location string) error {
	fmt.Println("Showing 20 location areas in Pokemon world...")
	locations, err := getFromURL[pokemonLocationResponse](state.Next, cache)
	if err != nil {
		return err
	}
	state.Next = locations.Next
	state.Previous = locations.Previous
	for _, location := range locations.Results {
		fmt.Printf("- %s\n", location.Name)
	}
	return nil
}

func commandMapBack(state *state, cache *pokecache.Cache, location string) error {
	fmt.Println("Showing previous 20 location areas in Pokemon world...")
	locations, err := getFromURL[pokemonLocationResponse](state.Previous, cache)
	if err != nil {
		return err
	}
	state.Next = locations.Next
	state.Previous = locations.Previous
	for _, location := range locations.Results {
		fmt.Printf("- %s\n", location.Name)
	}
	return nil
}

func commandExplore(state *state, cache *pokecache.Cache, location string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + location
	fmt.Printf("Exploring %s...\n", location)
	locationData, err := getFromURL[locationInfo](url, cache)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	for _, encounter := range locationData.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(state *state, cache *pokecache.Cache, pokemon string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	pokemonData, err := getFromURL[Pokemon](url, cache)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	catchChance := rand.Float64() * 100 / float64(pokemonData.BaseExperience)
	fmt.Printf("%.2f\n", catchChance)
	if catchChance < 0.5 {
		fmt.Printf("%s escaped!\n", pokemon)
		return nil
	}
	state.Pokedex.pokemon[pokemon] = pokemonData
	fmt.Printf("%s was caught!\n", pokemon)
	return nil
}

func commandInspect(state *state, cache *pokecache.Cache, pokemon string) error {
	if _, exists := state.Pokedex.pokemon[pokemon]; !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	pokemonData := state.Pokedex.pokemon[pokemon]
	fmt.Printf("Name: %s\n", pokemonData.Name)
	fmt.Printf("Base Experience: %d\n", pokemonData.BaseExperience)
	fmt.Printf("Height: %d\n", pokemonData.Height)
	fmt.Printf("Weight: %d\n", pokemonData.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonData.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemonData.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}
	return nil
}

func commandPokedex(state *state, cache *pokecache.Cache, pokemon string) error {
	if len(state.Pokedex.pokemon) == 0 {
		fmt.Println("Your Pokedex is empty. Catch some Pokemon first!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for name := range state.Pokedex.pokemon {
		fmt.Printf("- %s\n", name)
	}
	return nil
}
