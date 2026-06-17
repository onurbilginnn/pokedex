package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/onurbilginnn/pokecache"
)

func TestLocationCache(t *testing.T) {
	cacheInterval := 5 * time.Second
	const initialLocationURL = "https://pokeapi.co/api/v2/location-area/?limit=20"
	state := state{
		Next: initialLocationURL,
	}
	cache := pokecache.NewCache(cacheInterval)
	locations, err := getFromURL[pokemonLocationResponse](state.Next, cache)
	if err != nil {
		t.Fatalf("Failed to fetch locations: %v", err)
	}
	if len(locations.Results) == 0 {
		t.Fatal("Expected to fetch some locations, but got none")
	}
	fmt.Printf("Fetched %d locations\n", len(locations.Results))

	// Wait for the cache to expire
	time.Sleep(cacheInterval + 1*time.Second)

	// Try fetching again, it should fetch from the API again
	locations, err = getFromURL[pokemonLocationResponse](state.Next, cache)
	if err != nil {
		t.Fatalf("Failed to fetch locations after cache expiration: %v", err)
	}
	if len(locations.Results) == 0 {
		t.Fatal("Expected to fetch some locations after cache expiration, but got none")
	}
	fmt.Printf("Fetched %d locations after cache expiration\n", len(locations.Results))
}
