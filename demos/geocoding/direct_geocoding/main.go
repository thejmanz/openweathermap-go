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

	request := openweathermap.DirectGeocodingRequest{
		Query: "Seattle, WA, US", // City, State(US only), Country Code
		Limit: 0,                 // Optional value, default value is 0 resulting in a full list.
	}

	geo, err := client.DirectGeocode(context.Background(), request)
	if err != nil {
		log.Println(fmt.Errorf("request failed, reson: %w", err))
	}

	if !geo.Empty() {
		first := geo.Locations[0]
		log.Printf("Name: %s, Country: %s, (Lat: %g, Lon: %g)", first.Name, first.Country, first.Lat, first.Lon)
	} else {
		log.Printf("no results where found for %s", request.Query)
	}
}
