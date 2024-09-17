package commands

type config struct {
	Next     string
	Previous string
}

func GetConfig() config {
	return config{
		Next:     "https://pokeapi.co/api/v2/location",
		Previous: "",
	}
}
