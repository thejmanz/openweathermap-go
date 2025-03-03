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
	expected, err := jsonDataFromFile("data/geocoding/reverse_geocoding_response.json")
	if err != nil {
		t.Error(err.Error())
	}

	owm := mockClient(expected)
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
}

func mockClient(expected []byte) *OpenWeatherMap {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
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
		return fmt.Errorf("\nexpected = %v\ngot = %v", expected, got)
	}

	return nil
}
