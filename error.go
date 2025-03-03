package openweathermap

import (
	"encoding/json"
)

type APIError struct {
	Cod        json.Number `json:"cod"`
	Message    string      `json:"message"`
	Parameters []string    `json:"parameters,omitempty"`
}

func (a *APIError) Error() string {
	return a.Message
}
