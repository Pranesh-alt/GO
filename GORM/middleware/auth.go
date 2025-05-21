package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("supersecretkey")

// Context key type to avoid conflicts
type contextKey string

const userEmailKey contextKey = "user_email"

// AuthMiddleware validates JWT and injects user info into request context
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

			// Add user email to context
			ctx := context.WithValue(r.Context(), userEmailKey, claims.Subject)

			// Optional: implement role-based access control if roles are used
			// Currently skipping role validation logic as role is not in token

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GenerateToken creates a JWT for a given email
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

// GetUserEmailFromContext extracts the email from context
func GetUserEmailFromContext(r *http.Request) (string, bool) {
	email, ok := r.Context().Value(userEmailKey).(string)
	return email, ok
}
