package commands

import "github.com/mateus-f-torres/boot_pokedex/internal/pokeapi"

type Command struct {
	order       int
	name        string
	description string
	Callback    func(api *pokeapi.PokeAPI, args []string) error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"map": {
			order:       0,
			name:        "map",
			description: "print the names of the next 20 locations in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			order:       1,
			name:        "mapb",
			description: "print the names of the previous 20 locations in the Pokemon world",
			Callback:    commandMapB,
		},
		"explore": {
			order:       2,
			name:        "explore <location>",
			description: "explore a location and print list of available Pokémon",
			Callback:    commandExplore,
		},
		"catch": {
			order:       3,
			name:        "catch <pokemon>",
			description: "try to catch a given Pokémon",
			Callback:    commandCatch,
		},
		"help": {
			order:       4,
			name:        "help",
			description: "displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			order:       5,
			name:        "exit",
			description: "exit the Pokedex",
			Callback:    commandExit,
		},
	}

}
