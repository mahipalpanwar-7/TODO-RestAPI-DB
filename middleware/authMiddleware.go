package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)

			json.NewEncoder(w).Encode("Missing Token")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer")

		token, err := jwt.Parse(
			tokenString,
			func(t *jwt.Token) (interface{}, error) {
				return utils.JwtKey, nil
			},
		)

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)

			json.NewEncoder(w).Encode("Invalid Token")
			return
		}

		next.ServeHTTP(w, r)

	})
}
