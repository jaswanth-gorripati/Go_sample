package middleware

import (
	"context"
	"net/http"
	"regexp"

	"github.com/example/todo-api/internal/auth"
)

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if regexp.MustCompile(`^/swagger.*`).MatchString(r.URL.Path) || r.URL.Path == "/users/register" || r.URL.Path == "/users/login" {
			next.ServeHTTP(w, r)
			return
		}
		token := auth.ExtractToken(r)
		if token == "" {
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}
		userID, err := auth.ValidateJWT(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		// Set the user ID in the context
		ctx := context.WithValue(r.Context(), "userID", userID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
