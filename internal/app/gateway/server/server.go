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
	"github.com/davidchristie/tasks/internal/app/gateway/config"
	"github.com/davidchristie/tasks/internal/app/gateway/database"
	"github.com/davidchristie/tasks/internal/app/gateway/github"
	"github.com/davidchristie/tasks/internal/app/gateway/graphql"
	"github.com/davidchristie/tasks/internal/app/gateway/graphql/generated"
	"github.com/davidchristie/tasks/internal/app/gateway/logging"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const queryEndpoint = "/query"

func NewServer() (*http.Server, error) {
	conf := config.Read()

	logger := log.New(os.Stdout, "", log.LstdFlags)
	router := mux.NewRouter()

	db, err := database.Connect(conf.DatabaseURL)
	if err != nil {
		return nil, err
	}
	db.Migrate(conf.DatabaseMigrations)

	router.Use(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
	))
	router.Use(logging.Middleware(logger))
	router.Use(auth.Middleware(conf, db))

	router.HandleFunc("/login/github", github.LoginHandler(conf, logger)).Methods("GET")
	router.HandleFunc("/login/github/callback", github.CallbackHandler(conf, db, logger)).Methods("GET")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graphql.Resolver{
			Database: db,
		},
	}))
	router.Handle("/", playground.Handler("Tasks", queryEndpoint))
	router.Handle(queryEndpoint, srv)

	fmt.Printf("Port: %v\n", conf.Port)
	return &http.Server{
		Addr:         fmt.Sprintf(":%v", conf.Port),
		Handler:      router,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}, nil
}
