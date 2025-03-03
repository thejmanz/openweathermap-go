package openweathermap

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestOpenWeatherMap_ReverseGeocode(t *testing.T) {
	t.Run("reverse_geocode_request", func(t *testing.T) {
		expected, err := jsonDataFromFile("data/geocoding/reverse_geocoding_response.json")
		if err != nil {
			t.Error(err.Error())
		}

		owm := mockClient(expected, http.StatusOK)
		rs, err := owm.ReverseGeocode(context.TODO(), ReverseGeocodingRequest{})
		if err != nil {
			t.Error(err.Error())
		}

		got, err := json.Marshal(rs)
		if err != nil {
			t.Error(err.Error())
		}

		err = compareJson(expected, got)
		if err != nil {
			t.Error(err.Error())
		}
	})
}

func TestOpenWeatherMap_DirectGeocode(t *testing.T) {
	t.Run("direct_geocode_request", func(t *testing.T) {
		expected, err := jsonDataFromFile("data/geocoding/direct_geocoding_response.json")
		if err != nil {
			t.Error(err.Error())
		}

		owm := mockClient(expected, http.StatusOK)
		rs, err := owm.DirectGeocode(context.TODO(), DirectGeocodingRequest{})
		if err != nil {
			t.Error(err.Error())
		}

		got, err := json.Marshal(rs)
		if err != nil {
			t.Error(err.Error())
		}

		err = compareJson(expected, got)
		if err != nil {
			t.Error(err.Error())
		}
	})
}

func TestAPIError_Error(t *testing.T) {
	t.Run("api_error_response", func(t *testing.T) {
		expected, err := jsonDataFromFile("data/error/api_error_response.json")
		if err != nil {
			t.Error(err.Error())
		}

		owm := mockClient(expected, http.StatusBadRequest)
		_, err = owm.DirectGeocode(context.TODO(), DirectGeocodingRequest{})

		got, err := json.Marshal(err)
		if err != nil {
			t.Error(err.Error())
		}

		err = compareJson(expected, got)
		if err != nil {
			t.Error(err.Error())
		}
	})
}

func mockClient(expected []byte, status int) *OpenWeatherMap {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		_, _ = w.Write(expected)
	}))

	return New("YOURAPIKEY", withBaseUrl(srv.URL))
}

func jsonDataFromFile(fileName string) ([]byte, error) {
	by, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return by, nil
}

func compareJson(expected, got []byte) error {
	var e, g interface{}
	if err := json.Unmarshal(expected, &e); err != nil {
		return err
	}

	if err := json.Unmarshal(got, &g); err != nil {
		return err
	}

	if !reflect.DeepEqual(e, g) {
		return fmt.Errorf("\nexpected = %v\ngot = %v", e, g)
	}

	return nil
}
