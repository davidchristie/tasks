package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

type accessTokenFile struct {
	AccessToken string `json:"access_token"`
}

type config struct {
	AccessTokenFile string `default:"./data/access_token.json" split_words:"true"`
	Port string `required:"true"`
	UserFile string `default:"./data/user.json" split_words:"true"`
}

var conf = config{}

var logger = log.New(os.Stdout, "", log.LstdFlags)

func init() {
	envconfig.MustProcess("", &conf)
}

func main() {
	router := mux.NewRouter()

	router.Use(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
	))
	router.Use(func(next http.Handler) http.Handler {
		return handlers.LoggingHandler(logger.Writer(), next)
	})

	router.HandleFunc("/login/oauth/access_token", accessTokenHandler()).Methods("GET", "POST")
	router.HandleFunc("/login/oauth/authorize", authorizeHandler()).Methods("GET")
	router.HandleFunc("/user", userHandler()).Methods("GET")

	fmt.Printf("Port: %v\n", conf.Port)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", conf.Port),
		Handler:      router,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func accessTokenHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dat, err := ioutil.ReadFile(conf.AccessTokenFile)
		if (err != nil) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(dat)
	}
}

func authorizeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		redirectURI := r.URL.Query().Get("redirect_uri")
		code := "CODE"
		url := fmt.Sprintf("%s?code=%s", redirectURI, code)
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func userHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileToken := readAccessTokenFromFile()
		headerToken := getAccessTokenFromHeader(r)

		if (fileToken != headerToken) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{}"))
		}

		dat, err := ioutil.ReadFile(conf.UserFile)
		if (err != nil) {
			panic(err)
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(dat)
	}
}

func readAccessTokenFromFile() string {
	file, err := os.Open(conf.AccessTokenFile)
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	data := accessTokenFile{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		panic(err)
	}
	return data.AccessToken
}

func getAccessTokenFromHeader(r *http.Request) (string) {
	const prefix = "token "
	auth := r.Header.Get("Authorization")
	if auth == "" || !strings.HasPrefix(auth, prefix) {
		return ""
	}
	return strings.TrimPrefix(auth, prefix)
}
