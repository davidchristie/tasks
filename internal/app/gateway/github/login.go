package github

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davidchristie/tasks/internal/app/gateway/config"
	"github.com/davidchristie/tasks/internal/app/gateway/database"
	"github.com/davidchristie/tasks/internal/app/gateway/entity"
	"github.com/google/uuid"
)

func CallbackHandler(conf *config.Config, db database.Database, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		code := r.FormValue("code")

		token, err := requestAccessToken(conf, logger, code)
		if err != nil {
			logger.Printf("error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		githubUser, err := FetchUser(conf, token)
		if err != nil {
			logger.Printf("error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		dbUserExists, err := db.HasUserWithGithubID(githubUser.ID)
		if err != nil {
			logger.Printf("error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !dbUserExists {
			err = db.InsertUser(r.Context(), &entity.User{
				AvatarURL: githubUser.AvatarURL,
				GithubID:  githubUser.ID,
				Email:     githubUser.Email,
				ID:        uuid.New(),
				Name:      githubUser.Name,
			})
			if err != nil {
				logger.Printf("error: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		url := fmt.Sprintf("%s?token=%s", conf.WebApp, token)
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func LoginHandler(conf *config.Config, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		redirectURI := fmt.Sprintf("%s/login/github/callback", conf.Domain)
		url := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s", conf.GithubAuthorizeURL, conf.GithubClientID, redirectURI)
		http.Redirect(w, r, url, http.StatusFound)
	}
}
