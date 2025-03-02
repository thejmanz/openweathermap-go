package openweathermap

import (
	"encoding/json"
	"fmt"
)

type APIError struct {
	Cod        json.Number `json:"cod"`
	Message    string      `json:"message"`
	Parameters []string    `json:"parameters"`
}

func (a *APIError) Error() string {
	return fmt.Sprintf("http: %s, %s", a.Cod, a.Error())
}
