package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/davidchristie/tasks/internal/app/gateway/github"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// Middleware decodes the bearer token and packs the user into context.
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			const tokenPrefix = "token "

			a := r.Header.Get("Authorization")

			if a == "" {
				next.ServeHTTP(w, r)
				return
			}

			if !strings.HasPrefix(a, tokenPrefix) {
				next.ServeHTTP(w, r)
				return
			}
			// if !strings.HasPrefix(a, tokenPrefix) {
			// 	fmt.Println("doesn't have token prefix")
			// 	http.Error(w, "Invalid access token", http.StatusForbidden)
			// 	return
			// }

			token := strings.TrimPrefix(a, tokenPrefix)

			user, err := github.FetchUser(token)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			// if err != nil {
			// 	fmt.Printf("github error: %s\n", err)
			// 	http.Error(w, "Invalid access token", http.StatusForbidden)
			// 	return
			// }

			ctx := context.WithValue(r.Context(), userCtxKey, user)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *github.User {
	raw, _ := ctx.Value(userCtxKey).(*github.User)
	return raw
}
