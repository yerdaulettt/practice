package handlers

import (
	"context"
	"net/http"
	"time"
)

var testURL string

func SetURL(urlSting string) {
	testURL = urlSting
}

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	config := &retryConfig{maxRetries: 5, baseDelay: 500 * time.Millisecond, maxDelay: 5000 * time.Millisecond}

	err := executePayment(ctx, config)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	w.Write([]byte(`{"message": "ok"}`))
}
