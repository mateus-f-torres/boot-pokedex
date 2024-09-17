package commands

type Command struct {
	name        string
	description string
	Callback    func(*config) error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"map": {
			name:        "map",
			description: "Print the names of the next 20 locations in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Print the names of the previous 20 locations in the Pokemon world",
			Callback:    commandMapB,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}

}
