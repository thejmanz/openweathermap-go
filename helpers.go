package openweathermap

import (
	"fmt"
	"net/url"
)

func addStringValueWithDefault(v url.Values, key, val, defaultVal string) {
	if val == "" {
		v.Add(key, defaultVal)
	} else {
		v.Add(key, val)
	}
}

func addFloat64Value(v url.Values, key string, val float64) {
	v.Add(key, fmt.Sprintf("%g", val))
}

func addIntValue(v url.Values, key string, val int) {
	v.Add(key, fmt.Sprintf("%d", val))
}

func addInt64Value(v url.Values, key string, val int64) {
	v.Add(key, fmt.Sprintf("%d", val))
}
