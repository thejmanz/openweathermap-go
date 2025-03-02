package openweathermap

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
