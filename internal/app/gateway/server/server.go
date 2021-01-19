package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/davidchristie/tasks/internal/app/gateway/auth"
	"github.com/davidchristie/tasks/internal/app/gateway/github"
	"github.com/davidchristie/tasks/internal/app/gateway/graphql"
	"github.com/davidchristie/tasks/internal/app/gateway/graphql/generated"
	"github.com/davidchristie/tasks/internal/app/gateway/logging"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const queryEndpoint = "/query"

func getRequiredEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal(key + " is not defined")
	}
	return value
}

func NewServer() http.Server {
	domain := getRequiredEnv("DOMAIN")
	githubClientID := getRequiredEnv("GITHUB_CLIENT_ID")
	githubClientSecret := getRequiredEnv("GITHUB_CLIENT_SECRET")
	port := getRequiredEnv("PORT")
	webApp := getRequiredEnv("WEB_APP")

	logger := log.New(os.Stdout, "", log.LstdFlags)
	router := mux.NewRouter()

	router.Use(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
	))
	router.Use(logging.Middleware(logger))
	router.Use(auth.Middleware())

	router.HandleFunc("/login/github", github.LoginHandler(logger, githubClientID, domain)).Methods("GET")
	router.HandleFunc("/login/github/callback", github.CallbackHandler(logger, githubClientID, githubClientSecret, webApp)).Methods("GET")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))
	router.Handle("/", playground.Handler("Tasks", queryEndpoint))
	router.Handle(queryEndpoint, srv)

	fmt.Println("Port: " + port)
	return http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
}
