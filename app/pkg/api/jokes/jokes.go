package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prw_server/app/pkg/api"
)

const apiPath = "api?format=json"

func NewJokeClient(baseUrl string) *JokeClient {
	return &JokeClient{
		url: baseUrl,
	}
}

type JokeClient struct {
	url string
}

func (j *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := j.url + apiPath

	response, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request status: %s", http.StatusText(response.StatusCode))
	}

	var data api.JokeResponse

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
