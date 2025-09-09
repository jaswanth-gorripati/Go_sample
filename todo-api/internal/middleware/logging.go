package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// fmt.Printf("Request: %s %s\n", r.Method, r.URL)
		// Print input request body
		// body, err := io.ReadAll(r.Body)
		// if err != nil {
		// 	http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		// 	return
		// }
		// fmt.Printf("Request body: %s\n", body)

		next.ServeHTTP(w, r)

		// Response logging

		duration := time.Since(start)
		fmt.Printf("Request: %s %s took %v\n", r.Method, r.URL, duration)
	})
}
