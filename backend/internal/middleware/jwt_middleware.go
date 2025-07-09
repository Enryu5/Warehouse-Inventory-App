package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey []byte

func SetupJWTKey() {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		log.Fatal("JWT_SECRET_KEY is not set in environment variables")
	}
	jwtKey = []byte(secret)
}

type contextKey string

const AdminIDKey contextKey = "admin_id"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Accept both float64 and int admin_id (robust type assertion)
		var adminID int
		switch v := claims["admin_id"].(type) {
		case float64:
			adminID = int(v)
		case int:
			adminID = v
		default:
			http.Error(w, "Invalid admin_id in token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), AdminIDKey, adminID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
