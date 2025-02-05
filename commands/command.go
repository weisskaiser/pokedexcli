package commands

import (
	"time"

	"github.com/weisskaiser/pokedexcli/internal"
)

type NameValue struct {
	Name  string
	Value int
}

type Pokemon struct {
	Name   string
	Weight int
	Height int
	Stats  []NameValue
	Types  []string
}

type config struct {
	Next             string
	Previous         string
	Cache            *internal.Cache
	PokemonApiClient *internal.PokemonApiClient
	PokemonsCaught   map[string]Pokemon
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*config, []string) error
}

var cliCommands map[string]CliCommand

func NewConfig() *config {
	cache := internal.NewCache(time.Second * 5)
	return &config{
		Next:             "https://pokeapi.co/api/v2/location-area?limit=20",
		Cache:            cache,
		PokemonApiClient: internal.NewPokemonApiClient(cache, "https://pokeapi.co"),
		PokemonsCaught:   make(map[string]Pokemon),
	}
}

func InitCommands() map[string]CliCommand {
	cliCommands = map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map displays the next 20 locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to mapb displays the previous 20 locations",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore <area-name>",
			Description: "Displays the names of the pokemons of the given area",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch <pokemon-name>",
			Description: "Tries to capture the given pokemon",
			Callback:    commandCapture,
		},
		"inspect": {
			Name:        "inspect <pokemon-name>",
			Description: "Displays information about a caught pokemon",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Displays all caught pokemons",
			Callback:    commandPokedex,
		},
	}
	return cliCommands
}
