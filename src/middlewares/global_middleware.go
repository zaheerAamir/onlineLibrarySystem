package middlewares

import (
	"log"
	"net/http"
)

// LoggerMiddleware is a middleware to log the request
func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RawQuery != "" {
			log.Printf("********* Request: %s %s?%s ***********", r.Method, r.URL.Path, r.URL.RawQuery)
		} else {
			log.Printf("********* Request: %s %s ***********", r.Method, r.URL.Path)
		}

		next.ServeHTTP(w, r)
	}
}

func SetContentType(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Content-Type") != "application/json" {
			r.Header.Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}

	})
}
