package main

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestGetRate(t *testing.T) {
	s := NewExchangeService("http://localhost:8080")

	_, err := s.GetRate("USD", "KZT")
	if err != nil {
		t.Errorf("Got error %s", err)
	}

	_, err = s.GetRate("KZT", "KZT")
	if err == nil {
		t.Errorf("Expected to see invalid currency pair")
	}

	_, err = s.GetRate("KZT", "")
	if err == nil {
		t.Errorf("Expected to see invalid currency pair")
	}
}

type a struct {
	Msg string `json:"error"`
}

func TestGetRate2(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/convert?from=KZT&to=KZT", nil)
	defer req.Body.Close()

	var result a
	err := json.NewDecoder(req.Body).Decode(&result)
	if err != nil {
		t.Errorf("Failed %s", err)
	}
	fmt.Println(result)
	if result.Msg != "invalid currency pair" {
		t.Errorf("Expected invalid currency pair %s", err)
	}
}
