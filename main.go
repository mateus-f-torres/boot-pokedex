package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		inputs := strings.Split(buf.Text(), " ")
		if len(inputs) > 0 {
			c, ok := cmds[inputs[0]]
			if ok {
				err := c.Callback(&cnf, inputs[1:])
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
