package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/weisskaiser/pokedexcli/commands"
)

func cleanInput(text string) []string {
	var t []string

	for _, v := range strings.Fields(text) {
		trimmed := strings.TrimSpace(v)
		t = append(t, strings.ToLower(trimmed))
	}

	return t
}

func main() {
	cliCommands := commands.InitCommands()
	scanner := bufio.NewScanner(os.Stdin)
	config := commands.NewConfig()
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		command, ok := cliCommands[words[0]]
		args := words[1:]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.Callback(config, args)
		if err != nil {
			fmt.Println(err)
		}
	}
}
