package middleware

import (
	"context"
	"net/http"

	"github.com/amarantec/picpay/internal/utils"
)

type contextKey string

const UserIdKey contextKey = "userId"

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "token is empty", http.StatusUnauthorized)
			return
		}

		userId, err := utils.VerifyToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIdKey, userId)
		next(w, r.WithContext(ctx))
	}
}
