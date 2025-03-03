package openweathermap

import "net/url"

type GeocodingResult struct {
	Name       string            `json:"name"`
	LocalNames map[string]string `json:"local_names"`
	Lat        float64           `json:"lat"`
	Lon        float64           `json:"lon"`
	Country    string            `json:"country"`
	State      string            `json:"state,omitempty"` // US only
}

type GeocodingResponse []GeocodingResult

func (g GeocodingResponse) Empty() bool {
	return len(g) == 0
}

type ReverseGeocodingRequest struct {
	Lat   float64
	Lon   float64
	Limit int
}

func (r ReverseGeocodingRequest) endpoint(path string, v url.Values) string {
	addFloat64UrlValue("lat", r.Lat, v)
	addFloat64UrlValue("lon", r.Lon, v)
	addIntUrlValue("limit", r.Limit, v)
	return requestUrl(path, v)
}
