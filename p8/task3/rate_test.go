package main

import (
	"encoding/json"
	"errors"
	"math/rand/v2"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var errInvalidCurrencyPair error = errors.New("api error: invalid currency pair")
var errNetwork error = errors.New("network error")
var errServerPanic error = errors.New("api error: internal server error")
var errJson error = errors.New("decode error")
var errEmpty error = errors.New("decode error: EOF")

type testCase struct {
	name     string
	from, to string
	want     string
}

func TestGateRate3(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")

		if from == "" || to == "" || from == to {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "invalid currency pair"}`))
			return
		}

		if from == "json" {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if from == "slow" {
			time.Sleep(6 * time.Second)
		}

		if from == "panic" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "internal server error"}`))
			return
		}

		if from == "empty" {
			w.Write([]byte{})
			return
		}

		randomRate := rand.IntN(500)
		rate := &RateResponse{Base: from, Target: to, Rate: float64(randomRate)}

		err := json.NewEncoder(w).Encode(rate)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		}
	}))
	defer testServer.Close()

	client := NewExchangeService(testServer.URL)

	tests := []testCase{
		{"Correct response", "USD", "KZT", ""},
		{"Correct response 2", "KZT", "USD", ""},
		{"Missing from and to", "", "", errInvalidCurrencyPair.Error()},
		{"Missing from", "", "KZT", errInvalidCurrencyPair.Error()},
		{"Missing to", "KZT", "", errInvalidCurrencyPair.Error()},
		{"Same currency", "KZT", "KZT", errInvalidCurrencyPair.Error()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.GetRate(tt.from, tt.to)

			if err != nil && err.Error() != tt.want {
				t.Errorf("Expected (%s), but got (%s)", tt.want, err)
			}
		})
	}

	tests2 := []testCase{
		{"Malformed JSON", "json", "...", errJson.Error()},
		{"Server panic", "panic", "...", errServerPanic.Error()},
		{"Empty", "empty", "...", errEmpty.Error()},
		{"Slow response", "slow", "...", errNetwork.Error()},
	}

	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.GetRate(tt.from, tt.to)

			if err != nil && !strings.Contains(err.Error(), tt.want) {
				t.Errorf("Expected (%s), but got (%s)", tt.want, err)
			}
		})
	}
}
