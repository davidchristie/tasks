package github

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func CallbackHandler(logger *log.Logger, clientID, clientSecret, webApp string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		code := r.FormValue("code")

		token, err := requestAccessToken(logger, clientID, clientSecret, code)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
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
