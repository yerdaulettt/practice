package handlers

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
)

type rateResponse struct {
	Base     string  `json:"base"`
	Target   string  `json:"target"`
	Rate     float64 `json:"rate"`
	ErrorMsg string  `json:"error,omitempty"`
}

func errorJsonResponse(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(`{"error": "` + msg + `"}`))
}

func RateConvert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	if from == "" || to == "" || from == to {
		errorJsonResponse(w, http.StatusBadRequest, "invalid currency pair")
		return
	}

	randomRate := rand.IntN(500)
	rate := &rateResponse{Base: from, Target: to, Rate: float64(randomRate)}

	err := json.NewEncoder(w).Encode(rate)
	if err != nil {
		errorJsonResponse(w, http.StatusInternalServerError, err.Error())
	}
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "ok"}`))
}
