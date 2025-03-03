package openweathermap

import (
	"net/url"
	"strings"
)

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

type OneCallRequest struct {
	Lat   float64
	Lon   float64
	Units string
	Lang  string
}

func (o OneCallRequest) endpoint(path string, v url.Values) string {
	addFloat64UrlValue("lat", o.Lat, v)
	addFloat64UrlValue("lon", o.Lon, v)
	addStringUrlValue("units", o.Units, "metric", v)
	addStringUrlValue("lang", o.Lang, "en", v)
	return requestUrl(path, v)
}
