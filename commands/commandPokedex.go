package commands

import "fmt"

func commandPokedex(c *config, args []string) error {
	if len(c.PokemonsCaught) == 0 {
		fmt.Println("Your pokedex is empty")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, p := range c.PokemonsCaught {
		fmt.Printf("- %s\n", p.Name)
	}
	return nil
}
