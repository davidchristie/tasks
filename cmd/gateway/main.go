package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davidchristie/tasks/internal/app/gateway/github"
	"github.com/davidchristie/tasks/internal/app/gateway/middleware"
	"github.com/gorilla/mux"
)

func getRequiredEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal(key + " is not defined")
	}
	return value
}

func main() {
	githubClientID := getRequiredEnv("GITHUB_CLIENT_ID")
	githubClientSecret := getRequiredEnv("GITHUB_CLIENT_SECRET")
	port := getRequiredEnv("PORT")
	webApp := getRequiredEnv("WEB_APP")

	logger := log.New(os.Stdout, "", log.LstdFlags)
	router := mux.NewRouter()

	router.Use(middleware.Logging(logger))

	router.HandleFunc("/login/github", github.LoginHandler(logger, port, githubClientID)).Methods("GET")
	router.HandleFunc("/login/github/callback", github.CallbackHandler(logger, githubClientID, githubClientSecret, webApp)).Methods("GET")

	fmt.Println("Listening on port", port)
	server := http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
