package github

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davidchristie/tasks/internal/app/gateway/database"
	"github.com/google/uuid"
)

func CallbackHandler(db database.Database, logger *log.Logger, clientID, clientSecret, webApp string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		code := r.FormValue("code")

		token, err := requestAccessToken(logger, clientID, clientSecret, code)
		if err != nil {
			logger.Printf("error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		githubUser, err := FetchUser(token)
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
			err = db.InsertUser(r.Context(), &database.User{
				GithubID: githubUser.ID,
				Email:    githubUser.Email,
				ID:       uuid.New(),
				Name:     githubUser.Name,
			})
			if err != nil {
				logger.Printf("error: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		http.Redirect(w, r, webApp+"?token="+token, http.StatusFound)
	}
}

func LoginHandler(logger *log.Logger, clientID, domain string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		redirectURI := fmt.Sprintf("%s/login/github/callback", domain)
		url := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s", clientID, redirectURI)
		http.Redirect(w, r, url, http.StatusFound)
	}
}
