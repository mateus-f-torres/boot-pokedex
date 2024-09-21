package pokeapi

import (
	"time"

	"github.com/mateus-f-torres/boot_pokedex/internal/cache"
)

type PokeAPI struct {
	Next    string
	Prev    string
	Pokedex map[string]Pokemon

	cache cache.Cache

	// TODO: maybe game objects that track where the player currently is
	// so we can have the list of "visible pokemons" in the area?
}

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Stats  map[string]int
	Types  []string
}

func GetPokeAPI(cacheInterval time.Duration) *PokeAPI {
	return &PokeAPI{
		Next:    "https://pokeapi.co/api/v2/location",
		Prev:    "",
		Pokedex: map[string]Pokemon{},

		cache: cache.NewCache(cacheInterval),
	}
}
