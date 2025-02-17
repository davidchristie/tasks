package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/davidchristie/tasks/internal/app/gateway/config"
	"github.com/davidchristie/tasks/internal/app/gateway/database"
	"github.com/davidchristie/tasks/internal/app/gateway/entity"
	"github.com/davidchristie/tasks/internal/app/gateway/github"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// Middleware decodes the bearer token and packs the user into context.
func Middleware(conf *config.Config, db database.Database) func(http.Handler) http.Handler {
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

			token := strings.TrimPrefix(a, tokenPrefix)

			githubUser, err := github.FetchUser(conf, token)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			user, err := db.FindUserByGithubID(githubUser.ID)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, user)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. Requires Middleware to have run.
func ForContext(ctx context.Context) *entity.User {
	raw, _ := ctx.Value(userCtxKey).(*entity.User)
	return raw
}
