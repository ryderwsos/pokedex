package main

import (
	"time"

	"github.com/ryderwsos/pokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cfg := &Config{
		pokeapiClient: client,
	}
	startRepl(cfg)
}
