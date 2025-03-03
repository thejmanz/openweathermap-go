package openweathermap

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

const baseUrl = "https://api.openweathermap.org"

type OpenWeatherMap struct {
	apiKey  string
	baseUrl string
}

type Option func(o *OpenWeatherMap)

// withBaseUrl Internal function used for testing purposes
func withBaseUrl(url string) Option {
	return func(o *OpenWeatherMap) {
		o.baseUrl = url
	}
}

func New(apiKey string) *OpenWeatherMap {
	o := OpenWeatherMap{apiKey: apiKey}
	return &o
}

func (o *OpenWeatherMap) getUrlAppendingPath(path string) string {
	var u string
	if o.baseUrl == "" {
		u = baseUrl
	} else {
		u = o.baseUrl
	}
	return u + path
}

func (o *OpenWeatherMap) makeRequest(ctx context.Context, url string, destination interface{}) error {
	rq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	rs, err := http.DefaultClient.Do(rq)
	if err != nil {
		return err
	}
	defer rs.Body.Close()

	by, err := io.ReadAll(rs.Body)
	if err != nil {
		return err
	}

	if rs.StatusCode != http.StatusOK {
		var apiError APIError
		if err = json.Unmarshal(by, &apiError); err != nil {
			return err
		}
		return &apiError
	}

	return json.Unmarshal(by, &destination)
}
