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

	return locationAreaData, nil
}
