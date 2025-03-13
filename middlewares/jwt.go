package middleware

import (
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("notes321")

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Tokens Missing", http.StatusUnauthorized)
			return
		}
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}
		username := claims["username"].(string)
		log.Printf("User Authenticated: %v", username)
		next.ServeHTTP(w, r)
	})
}
