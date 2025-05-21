package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

var JwtSecret = []byte("supersecretkey")

type contextKey string

const userEmailKey contextKey = "user_email"

func AuthMiddleware(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Authorization header missing or invalid", http.StatusUnauthorized)
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			claims := &jwt.RegisteredClaims{}

			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return JwtSecret, nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userEmailKey, claims.Subject)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GenerateToken(email string) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour)
	claims := jwt.RegisteredClaims{
		Subject:   email,
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

func GetUserEmailFromContext(r *http.Request) (string, bool) {
	email, ok := r.Context().Value(userEmailKey).(string)
	return email, ok
}
