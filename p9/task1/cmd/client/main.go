package main

import (
	"net/http"
	"net/http/httptest"

	"p9task1/internal/app"
)

func main() {
	counter := 1

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if counter <= 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			counter++
			return
		}

		counter = 1
		w.Write([]byte(`{"status": "success"}`))
	}))
	defer testServer.Close()

	app.Run(testServer.URL)
}
