# openweathermap-go


OpenWeatherMap-Go is a wrapper library for the [OpenWeatherMap API](https://api.openweathermap.org). Primary use case
is to provide an easy-to-use interface for interacting with common [OpenWeatherMap API](https://api.openweathermap.org)
services such as their [One Call 3.0 API](https://openweathermap.org/api/one-call-3)
and [Geocoding API](https://openweathermap.org/api/geocoding-api).


## Installation

    $ go get github.com/thejmanz/openweathermap-go

## Usage

A reverse geocoding request example.
```go
import github.com/thejmanz/openweathermap

func main() {
    token := os.Getenv("OPENWEATHERMAP_API_KEY")
	
    client := openweathermap.New(token)

    request := openweathermap.ReverseGeocodingRequest{
        Lat:   34.3233,
        Lon:   -118.4643,
        Limit: 0, // Optional value, default value is 0 resulting in a full list.
    }

    geo, err := client.ReverseGeocode(context.Background(), request)
    if err != nil {
        //Handle any errors
    }
	
    if !geo.Empty() {
        first := geo.Locations[0]
        log.Printf(first)
    }
}
```
#### Usage Disclaimer
OpenWeatherMap-Go is currently in beta and under ongoing development, therefore it does not provide a full range of
[OpenWeatherMap API](https://api.openweathermap.org) functionality.
