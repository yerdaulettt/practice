package middleware

import (
	"log"
	"net/http"
)

func AuthAndLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-KEY") != "123" {
			log.Println(r.Method, r.URL.String(), "server is using", r.Proto, "!!!Unautorized!!!")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error":"unauthorized"}`))
			return
		}

		log.Println(r.Method, r.URL.String(), " server is using", r.Proto)
		next.ServeHTTP(w, r)
	})
}
