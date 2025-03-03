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

type OneCallCurrentWeatherData struct {
	Dt         int64         `json:"dt"`
	Sunrise    int64         `json:"sunrise"`
	Sunset     int64         `json:"sunset"`
	Temp       float64       `json:"temp"`
	FeelsLike  float64       `json:"feels_like"`
	Pressure   int64         `json:"pressure"`
	Humidity   int64         `json:"humidity"`
	DewPoint   float64       `json:"dew_point"`
	Uvi        float64       `json:"uvi"`
	Clouds     int64         `json:"clouds"`
	Visibility int64         `json:"visibility"`
	WindSpeed  float64       `json:"wind_speed"`
	WindDeg    int64         `json:"wind_deg"`
	WindGust   float64       `json:"wind_gust,omitempty"`
	Weather    []WeatherData `json:"weather"`
	Rain       *struct {
		OneHour float64 `json:"1h"`
	} `json:"rain,omitempty"`
	Snow *struct {
		OneHour float64 `json:"1h"`
	} `json:"snow,omitempty"`
}

type OneCallForecastMinutelyData struct {
	Dt            int64   `json:"dt"`
	Precipitation float64 `json:"precipitation"`
}

type OneCallForecastHourlyData struct {
	Dt         int64         `json:"dt"`
	Temp       float64       `json:"temp"`
	FeelsLike  float64       `json:"feels_like"`
	Pressure   int64         `json:"pressure"`
	Humidity   int64         `json:"humidity"`
	DewPoint   float64       `json:"dew_point"`
	Uvi        float64       `json:"uvi"`
	Clouds     int64         `json:"clouds"`
	Visibility int64         `json:"visibility"`
	WindSpeed  float64       `json:"wind_speed"`
	WindDeg    int64         `json:"wind_deg"`
	WindGust   float64       `json:"wind_gust,omitempty"`
	Weather    []WeatherData `json:"weather"`
	Pop        float64       `json:"pop"`
	Rain       *struct {
		OneHour float64 `json:"1h"`
	} `json:"rain,omitempty"`
	Snow *struct {
		OneHour float64 `json:"1h"`
	} `json:"snow,omitempty"`
}

type OneCallForecastDailyData struct {
	Dt        int64   `json:"dt"`
	Sunrise   int64   `json:"sunrise"`
	Sunset    int64   `json:"sunset"`
	Moonrise  int64   `json:"moonrise"`
	Moonset   int64   `json:"moonset"`
	MoonPhase float64 `json:"moon_phase"`
	Summary   string  `json:"summary"`
	Temp      struct {
		Day   float64 `json:"day"`
		Min   float64 `json:"min"`
		Max   float64 `json:"max"`
		Night float64 `json:"night"`
		Eve   float64 `json:"eve"`
		Morn  float64 `json:"morn"`
	} `json:"temp"`
	FeelsLike struct {
		Day   float64 `json:"day"`
		Night float64 `json:"night"`
		Eve   float64 `json:"eve"`
		Morn  float64 `json:"morn"`
	} `json:"feels_like"`
	Pressure  int64         `json:"pressure"`
	Humidity  int64         `json:"humidity"`
	DewPoint  float64       `json:"dew_point"`
	WindSpeed float64       `json:"wind_speed"`
	WindDeg   float64       `json:"wind_deg"`
	WindGust  float64       `json:"wind_gust"`
	Weather   []WeatherData `json:"weather"`
	Clouds    int64         `json:"clouds"`
	Pop       float64       `json:"pop"`
	Rain      float64       `json:"rain,omitempty"`
	Snow      float64       `json:"snow,omitempty"`
	Uvi       float64       `json:"uvi"`
}

type OneCallWeatherAlertData struct {
	SenderName  string   `json:"sender_name"`
	Event       string   `json:"event"`
	Start       int64    `json:"start"`
	End         int64    `json:"end"`
	Description string   `json:"description"`
	Tags        []string `json:"tags,omitempty"`
}

type OneCallWeatherResponse struct {
	Lat            float64                       `json:"lat"`
	Lon            float64                       `json:"lon"`
	Timezone       string                        `json:"timezone"`
	TimezoneOffset int64                         `json:"timezone_offset"`
	Current        *OneCallCurrentWeatherData    `json:"current,omitempty"`
	Minutely       []OneCallForecastMinutelyData `json:"minutely,omitempty"`
	Hourly         []OneCallForecastHourlyData   `json:"hourly,omitempty"`
	Daily          []OneCallForecastDailyData    `json:"daily,omitempty"`
	Alerts         []OneCallWeatherAlertData     `json:"alerts,omitempty"`
}
