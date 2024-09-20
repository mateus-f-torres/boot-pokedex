package pokeapi

import (
	"time"

	"github.com/mateus-f-torres/boot_pokedex/internal/cache"
)

const location_api string = "https://pokeapi.co/api/v2/location"

type PokeAPI struct {
	Next  string
	Prev  string
	cache cache.Cache
}

func GetPokeAPI(cacheInterval time.Duration) *PokeAPI {
	return &PokeAPI{
		Next:  location_api,
		Prev:  "",
		cache: cache.NewCache(cacheInterval),
	}
}
