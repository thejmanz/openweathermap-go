package openweathermap

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"time"
)

type GeocodingRequestType = int

const (
	ReverseGeocode GeocodingRequestType = iota
	DirectGeocode  GeocodingRequestType = 1
)

type OneCallRequestType = int

const (
	OneCall     OneCallRequestType = iota
	TimeMachine OneCallRequestType = 1
)

func TestOpenWeatherMap_OneCall(t *testing.T) {
	t.Run("one_call_request", func(t *testing.T) {
		oneCall(t, "data/onecall/one_call_response.json", OneCall)
	})

	t.Run("one_call_current_request", func(t *testing.T) {
		oneCall(t, "data/onecall/one_call_current_response.json", OneCall)
	})
}

func TestOpenWeatherMap_OneCallTimeMachine(t *testing.T) {
	oneCall(t, "data/onecall/one_call_timemachine_response.json", TimeMachine)
}

func TestOpenWeatherMap_DirectGeocode(t *testing.T) {
	geocode(t, "data/geocoding/direct_geocoding_response.json", DirectGeocode)
}

func TestOpenWeatherMap_ReverseGeocode(t *testing.T) {
	geocode(t, "data/geocoding/reverse_geocoding_response.json", ReverseGeocode)
}

func TestAPIError_Error(t *testing.T) {
	owm, expected, err := getTestClient("data/error/api_error_response.json", http.StatusUnauthorized)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = owm.OneCall(context.TODO(), OneCallRequest{}, nil)
	if err == nil {
		t.Errorf("Expected request to produce an error")
	}

	var apiError *APIError
	if errors.As(err, &apiError) {
		got, marshalErr := json.Marshal(apiError)
		if marshalErr != nil {
			t.Error(marshalErr.Error())
		}

		if err = compareJson(expected, got); err != nil {
			t.Error(err.Error())
		}
	} else {
		t.Errorf("Invalid error type")
	}
}

func oneCall(t *testing.T, filename string, requestType OneCallRequestType) {
	owm, expected, err := getTestClient(filename, http.StatusOK)
	if err != nil {
		t.Fatal(err.Error())
	}

	var rs interface{}
	switch requestType {
	case OneCall:
		rs, err = owm.OneCall(context.TODO(), OneCallRequest{}, nil)
	case TimeMachine:
		rs, err = owm.OneCallTimeMachine(context.TODO(), OneCallRequest{}, time.Now().Unix())
	}

	if err != nil {
		t.Error(err.Error())
	}

	got, err := json.Marshal(rs)
	if err != nil {
		t.Error(err.Error())
	}

	if err = compareJson(expected, got); err != nil {
		t.Error(err.Error())
	}
}

func geocode(t *testing.T, filename string, requestType GeocodingRequestType) {
	owm, expected, err := getTestClient(filename, http.StatusOK)
	if err != nil {
		t.Fatal(err.Error())
	}

	var rs interface{}
	switch requestType {
	case ReverseGeocode:
		rs, err = owm.ReverseGeocode(context.TODO(), ReverseGeocodingRequest{})
	case DirectGeocode:
		rs, err = owm.DirectGeocode(context.TODO(), DirectGeocodingRequest{})
	}

	if err != nil {
		t.Error(err.Error())
	}

	got, err := json.Marshal(rs)
	if err != nil {
		t.Error(err.Error())
	}

	if err = compareJson(expected, got); err != nil {
		t.Error(err.Error())
	}
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
		return fmt.Errorf("\nexpected=%s\ngot=%s", e, g)
	}

	return nil
}

func getTestClient(filename string, status int) (*OpenWeatherMap, []byte, error) {
	expected, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		_, _ = w.Write(expected)
	}))

	cl := New("YOURAPIKEY", withBaseUrl(srv.URL))

	return cl, expected, nil
}
