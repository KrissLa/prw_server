package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prw_server/app/pkg/api"
)

const apiPath = "api?format=json"

type JokeClient struct {
	url string
}

func (j *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := j.url + apiPath

	response, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request error: %v", err)
	}

	var data api.JokeResponse

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

