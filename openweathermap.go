package openweathermap

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
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

func New(apiKey string, opts ...Option) *OpenWeatherMap {
	o := OpenWeatherMap{apiKey: apiKey}
	for _, opt := range opts {
		opt(&o)
	}
	return &o
}

func (o *OpenWeatherMap) ReverseGeocode(ctx context.Context, r ReverseGeocodingRequest) (*GeocodingResponse, error) {
	return o.reverseGeocode(ctx, r)
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

func (o *OpenWeatherMap) getCredentialedValues() url.Values {
	return url.Values{"appid": {o.apiKey}}
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

func (o *OpenWeatherMap) reverseGeocode(ctx context.Context, b requestBuilder) (*GeocodingResponse, error) {
	v := o.getCredentialedValues()
	p := o.getUrlAppendingPath("/geo/1.0/reverse")

	var r GeocodingResponse
	if err := o.makeRequest(ctx, b.endpoint(p, v), &r); err != nil {
		return nil, err
	}
	return &r, nil
}
