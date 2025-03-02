package openweathermap

import (
	"fmt"
	"net/url"
)

func addStringUrlValue(key, val, defaultVal string, v url.Values) {
	if val == "" {
		v.Add(key, defaultVal)
	} else {
		v.Add(key, val)
	}
}

func addFloat64UrlValue(key string, val float64, v url.Values) {
	v.Add(key, fmt.Sprintf("%g", val))
}

func addIntUrlValue(key string, val int, v url.Values) {
	if val > 0 {
		v.Add(key, fmt.Sprintf("%d", val))
	}
}
