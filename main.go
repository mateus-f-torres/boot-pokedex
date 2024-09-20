package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/mateus-f-torres/boot_pokedex/internal/cache"
	"github.com/mateus-f-torres/boot_pokedex/internal/commands"
)

func main() {
	prompt := "pokedex > "
	buf := bufio.NewScanner(os.Stdin)
	cache := cache.NewCache(1 * time.Minute)
	cmds := commands.GetCommands()
	cnf := commands.GetConfig(cache)
	for {
		fmt.Print(prompt)
		buf.Scan()
		input := buf.Text()
		c, ok := cmds[input]
		if ok {
			err := c.Callback(&cnf)
			if err != nil {
				fmt.Printf("an error occurred: %v\n", err)
			}
		} else {
			fmt.Printf("unknown command: %v\n", input)
			fmt.Println("type 'help' to see list of available commands")
		}
	}
}
