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

	request := openweathermap.ZipGeocodingRequest{
		Query: "90022,US", // Zip/Postal Code, Country Code
	}

	geo, err := client.ZipGeocode(context.TODO(), request)
	if err != nil {
		log.Fatal(fmt.Errorf("request failed, reason: %w", err))
	}

	log.Printf("Name: %s, Zip: %s, Lat: %g, Lon: %g, Country: %s", geo.Name, geo.Zip, geo.Lat, geo.Lon, geo.Country)
}
