package web

import (
	"context"
	"net/http"
	"strings"

	"github.com/oxodao/overflow-bot/services"
)

func isUnauthedRoute(path string) bool {
	if path == "/api/call" {
		return true
	}

	if path == "/api/auth" {
		return true
	}

	if strings.HasPrefix(path, "/api/auth/callback") {
		return true
	}

	return false
}

func authedMiddleware(prv *services.Provider) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if isUnauthedRoute(r.URL.Path) {
				h.ServeHTTP(w, r)
				return
			}

			token := r.Header.Get("Authorization")
			user, err := findUserByToken(prv, token)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, "user", user)

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
