package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mateus-f-torres/boot_pokedex/internal/commands"
	"github.com/mateus-f-torres/boot_pokedex/internal/pokeapi"
)

func main() {
	prompt := "pokedex > "
	buffer := bufio.NewScanner(os.Stdin)
	cmds := commands.GetCommands()
	pokeapi := pokeapi.GetPokeAPI(1 * time.Minute)
	for {
		fmt.Print(prompt)
		buffer.Scan()
		inputs := strings.Split(buffer.Text(), " ")
		if len(inputs) > 0 {
			c, ok := cmds[inputs[0]]
			if ok {
				err := c.Callback(pokeapi, inputs[1:])
				if err != nil {
					fmt.Printf("an error occurred: %v\n", err)
				}
			} else {
				fmt.Printf("unknown command: %v\n", inputs[0])
				fmt.Println("type 'help' to see list of available commands")
			}
		}
	}
}
