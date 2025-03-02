package openweathermap

import (
	"fmt"
	"net/url"
)

func requestUrl(path string, v url.Values) string {
	return fmt.Sprintf("%s?%s", path, v.Encode())
}

type requestBuilder interface {
	endpoint(path string, v url.Values) string
}
