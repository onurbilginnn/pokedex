package main

type pokemonLocation struct {
	Name string `json:"name"`
}

type pokemonLocationResponse struct {
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []pokemonLocation `json:"results"`
}

type pokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
	} `json:"pokemon"`
}

type locationInfo struct {
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

type PokemonStat struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type PokemonType struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}
