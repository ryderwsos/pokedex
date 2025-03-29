package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := client.cache.Get(url); ok {
		encountersData := Pokemon{}
		err := json.Unmarshal(val, &encountersData)
		if err != nil {
			return Pokemon{}, err
		}
		return encountersData, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonData := Pokemon{}
	err = json.Unmarshal(data, &pokemonData)
	if err != nil {
		return Pokemon{}, err
	}

	client.cache.Add(url, data)

	return pokemonData, nil
}
