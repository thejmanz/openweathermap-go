package openweathermap

import "strings"

type OneCallWeatherType string

type OneCallDataSet []OneCallWeatherType

const (
	OneCallCurrentWeather   OneCallWeatherType = "current"
	OneCallForecastMinutely OneCallWeatherType = "minutely"
	OneCallForecastHourly   OneCallWeatherType = "hourly"
	OneCallForecastDaily    OneCallWeatherType = "daily"
	OneCallWeatherAlerts    OneCallWeatherType = "alerts"
)

func (o OneCallDataSet) Excluding() string {
	e := make([]string, len(o))
	for i, s := range o {
		e[i] = string(s)
	}
	return strings.Join(e, ",")
}
