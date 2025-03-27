package main

import (
	"time"

	"github.com/ryderwsos/pokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &Config{
		pokeapiClient: client,
	}
	startRepl(cfg)
}
