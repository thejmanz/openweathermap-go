package main

import (
	"context"
	"fmt"
	"github.com/thejmanz/openweathermap-go"
	"log"
	"os"
	"time"
)

func main() {
	token := os.Getenv("OPENWEATHERMAP_API_KEY")

	client := openweathermap.New(token)

	request := openweathermap.OneCallRequest{
		Lat:   34.4645,
		Lon:   -118.6500,
		Units: "imperial", // optional, the default is 'metric' units
		Lang:  "en",       // optional, the default is 'en' language
	}

	// Use a OneCallDataSet to exclude One Call data types from the resulting response.
	exclude := openweathermap.OneCallDataSet{
		openweathermap.OneCallForecastMinutely,
		openweathermap.OneCallForecastHourly,
	}

	w, err := client.OneCall(context.Background(), request, exclude)
	if err != nil {
		log.Fatalln(fmt.Errorf("request failed, reason: %w", err))
	}

	s := fmt.Sprintf("Current Temp: %gF @ %v in %s", w.Current.Temp, time.Unix(w.Current.Dt, 0), w.Timezone)

	log.Println(s)
}
