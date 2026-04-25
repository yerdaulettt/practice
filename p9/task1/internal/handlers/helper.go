package handlers

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"
)

func IsRetryable(resp *http.Response, err error) bool {
	switch resp.StatusCode {
	case 429, 500, 502, 503, 504:
		return true
	case 401, 404:
		return false
	}

	return false
}

func CalculateBackoff(attempt int, cfg *retryConfig) time.Duration {
	backoff := cfg.baseDelay * time.Duration(math.Pow(2, float64(attempt)))
	if backoff > cfg.maxDelay {
		backoff = cfg.baseDelay
	}
	jitter := time.Duration(rand.Int63n(int64(backoff)))

	return jitter
}

func executePayment(ctx context.Context, cfg *retryConfig) error {
	var err error
	for attempt := 0; attempt < cfg.maxRetries; attempt++ {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		resp, err := http.Get(testURL)
		if !IsRetryable(resp, err) {
			return err
		}
		defer resp.Body.Close()

		// if err == nil {
		// 	break
		// }

		if attempt == cfg.maxRetries-1 {
			return err
		}

		backoff := CalculateBackoff(attempt, cfg)

		fmt.Printf("Attempt %d failed, waiting ~%v...\n", attempt+1, backoff)
		time.Sleep(backoff)
	}

	return err
}
