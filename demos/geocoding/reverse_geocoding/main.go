package main

import (
	"context"
	"fmt"
	"github.com/thejmanz/openweathermap-go"
	"log"
	"os"
)

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
		log.Println(fmt.Errorf("request failed, reason: %w", err))
	}

	if !geo.Empty() {
		first := geo.Locations[0]
		log.Printf("Name: %s, Country: %s, (Lat: %g, Lon: %g)", first.Name, first.Country, first.Lat, first.Lon)
	} else {
		log.Printf("no results where found for (Lat: %g, Lon: %g)", request.Lat, request.Lon)
	}
}
