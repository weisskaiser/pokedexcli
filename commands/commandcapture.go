package commands

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCapture(c *config, args []string) error {
	if len(args) == 0 {
		return errors.New("required argument pokemon name is missing")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	root, err := c.PokemonApiClient.GetPokemonInformation(args[0])
	if err != nil {
		return err
	}

	chance := 0.9 - (root.BaseExperience/1000)*2.

	if chance < 0.01 {
		chance = 0.01
	}

	got := rand.Float64()

	fmt.Printf("chance %.2f got %.2f\n", chance, got)

	if got > chance {
		fmt.Printf("%s escaped!\n", root.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", root.Name)

	stats := []NameValue{}
	types := []string{}
	for _, s := range root.Stats {
		stats = append(stats, NameValue{s.Stat.Name, s.BaseStat})
	}
	for _, t := range root.Types {
		types = append(types, t.Type.Name)
	}

	c.PokemonsCaught[root.Name] = Pokemon{
		Name:   root.Name,
		Weight: root.Weight,
		Height: root.Height,
		Stats:  stats,
		Types:  types,
	}

	return nil
}
