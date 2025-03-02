package openweathermap

import "testing"

func TestOneCallDataSet_Excluding(t *testing.T) {
	ts := []struct {
		Name     string
		Set      OneCallDataSet
		Expected string
	}{
		{
			"one_call_data_excluding_none",
			OneCallDataSet{},
			"",
		}, {
			"one_call_data_set_excluding_all",
			OneCallDataSet{
				OneCallCurrentWeather,
				OneCallForecastMinutely,
				OneCallForecastHourly,
				OneCallForecastDaily,
				OneCallWeatherAlerts,
			},
			"current,minutely,hourly,daily,alerts",
		},
	}
	for _, tt := range ts {
		t.Run(tt.Name, func(t *testing.T) {
			if got := tt.Set.Excluding(); got != tt.Expected {
				t.Errorf("\nExcluding() = %s\n expected = %s", got, tt.Expected)
			}
		})
	}
}
