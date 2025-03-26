package main

import (
	"errors"
	"fmt"
)

func commandMap(config *Config) error {
	locationAreaList, err := config.pokeapiClient.GetAreaLocations(config.Next)
	if err != nil {
		return err
	}
	config.Next = locationAreaList.Next
	config.Previous = locationAreaList.Previous

	for _, i := range locationAreaList.Results {
		fmt.Println(i.Name)
	}
	return nil
}

func commandMapb(config *Config) error {
	if config.Previous == nil {
		return errors.New("no previous map, you are on the first page")
	}
	locationAreaList, err := config.pokeapiClient.GetAreaLocations(config.Previous)
	if err != nil {
		return err
	}
	config.Next = locationAreaList.Next
	config.Previous = locationAreaList.Previous

	for _, i := range locationAreaList.Results {
		fmt.Println(i.Name)
	}
	return nil
}
