package middleware

import (
	"context"
	"net/http"
)

type contextKey string

const userIdKey contextKey = "userId"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		userId, err := verifyTokenAndGetUserId(token)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userIdKey, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserIdFromContext(r *http.Request) (string, bool) {
	userId, ok := r.Context().Value(userIdKey).(string)
	return userId, ok
}

func verifyTokenAndGetUserId(token string) (string, error) {
	//TODO: change with real authentication
	return "admin", nil
}
