package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) GetAreaLocations(urlExt *string) (locationArea, error) {
	url := baseURL + "/location-area"

	if urlExt != nil {
		url = *urlExt
	}

	//check if current url already exist in cache
	if val, ok := client.cache.Get(url); ok {
		locationAreaData := locationArea{}
		err := json.Unmarshal(val, &locationAreaData)
		if err != nil {
			return locationArea{}, err
		}
		return locationAreaData, nil
	}

	//move on with rest of request

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationArea{}, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return locationArea{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationArea{}, err
	}

	locationAreaData := locationArea{}
	err = json.Unmarshal(data, &locationAreaData)
	if err != nil {
		return locationArea{}, err
	}

	//add url to cache since it didnt exist already
	client.cache.Add(url, data)

	return locationAreaData, nil
}
