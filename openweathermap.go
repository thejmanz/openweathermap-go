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

func (o *OpenWeatherMap) OneCall(ctx context.Context, r OneCallRequest, exclude OneCallDataSet) (*OneCallWeatherResponse, error) {
	return o.oneCall(ctx, r, exclude)
}

func (o *OpenWeatherMap) OneCallTimeMachine(ctx context.Context, r OneCallRequest, timestamp int64) (*OneCallTimeMachineResponse, error) {
	return o.oneCallTimeMachine(ctx, r, timestamp)
}

func (o *OpenWeatherMap) ReverseGeocode(ctx context.Context, r ReverseGeocodingRequest) (*GeocodingResponse, error) {
	return o.geocode(ctx, r, "/geo/1.0/reverse")
}

func (o *OpenWeatherMap) DirectGeocode(ctx context.Context, r DirectGeocodingRequest) (*GeocodingResponse, error) {
	return o.geocode(ctx, r, "/geo/1.0/direct")
}

func (o *OpenWeatherMap) ZipGeocode(ctx context.Context, r ZipGeocodingRequest) (*ZipGeocodingResponse, error) {
	return o.zipGeocode(ctx, r, "/geo/1.0/zip")
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
	rq, err := http.NewRequestWithContext(ctx, "GET", url, nil)
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

func (o *OpenWeatherMap) geocode(ctx context.Context, b requestBuilder, path string) (*GeocodingResponse, error) {
	v := o.getCredentialedValues()
	p := o.getUrlAppendingPath(path)

	var r GeocodingResponse
	if err := o.makeRequest(ctx, b.endpoint(p, v), &r.Locations); err != nil {
		return nil, err
	}
	return &r, nil
}

func (o *OpenWeatherMap) zipGeocode(ctx context.Context, b requestBuilder, path string) (*ZipGeocodingResponse, error) {
	v := o.getCredentialedValues()
	p := o.getUrlAppendingPath(path)

	var r ZipGeocodingResponse
	if err := o.makeRequest(ctx, b.endpoint(p, v), &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (o *OpenWeatherMap) oneCall(ctx context.Context, b requestBuilder, exclude OneCallDataSet) (*OneCallWeatherResponse, error) {
	v := o.getCredentialedValues()
	p := o.getUrlAppendingPath("/data/3.0/onecall")
	if len(exclude) > 0 {
		v.Add("exclude", exclude.Excluding())
	}

	var r OneCallWeatherResponse
	if err := o.makeRequest(ctx, b.endpoint(p, v), &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (o *OpenWeatherMap) oneCallTimeMachine(ctx context.Context, b requestBuilder, timestamp int64) (*OneCallTimeMachineResponse, error) {
	v := o.getCredentialedValues()
	p := o.getUrlAppendingPath("/data/3.0/onecall/timemachine")
	addInt64Value(v, "dt", timestamp)

	var r OneCallTimeMachineResponse
	if err := o.makeRequest(ctx, b.endpoint(p, v), &r); err != nil {
		return nil, err
	}
	return &r, nil
}
