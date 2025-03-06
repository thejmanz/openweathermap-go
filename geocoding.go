package openweathermap

import "net/url"

type GeocodingLocationData struct {
	Name       string            `json:"name"`
	LocalNames map[string]string `json:"local_names,omitempty"`
	Lat        float64           `json:"lat"`
	Lon        float64           `json:"lon"`
	Country    string            `json:"country"`
	State      string            `json:"state,omitempty"` // US only
}

type GeocodingResponse struct {
	Locations []GeocodingLocationData
}

func (g GeocodingResponse) Empty() bool {
	return len(g.Locations) == 0
}

type ReverseGeocodingRequest struct {
	Lat   float64
	Lon   float64
	Limit int
}

func (r ReverseGeocodingRequest) endpoint(path string, v url.Values) string {
	addFloat64Value(v, "lat", r.Lat)
	addFloat64Value(v, "lon", r.Lon)
	addIntValue(v, "limit", r.Limit)
	return requestUrl(path, v)
}

type DirectGeocodingRequest struct {
	Query string
	Limit int
}

func (d DirectGeocodingRequest) endpoint(path string, v url.Values) string {
	v.Add("q", d.Query)
	addIntValue(v, "limit", d.Limit)
	return requestUrl(path, v)
}

type ZipGeocodingResponse struct {
	Zip     string  `json:"zip"`
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
}

type ZipGeocodingRequest struct {
	Query string
}

func (z ZipGeocodingRequest) endpoint(path string, v url.Values) string {
	v.Add("zip", z.Query)
	return requestUrl(path, v)
}
