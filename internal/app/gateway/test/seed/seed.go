package seed

import (
	"github.com/davidchristie/tasks/internal/app/gateway/database"
	"github.com/gorilla/mux"
)

func Populate(router *mux.Router, db database.Database) {
	router.Handle("/reset", resetHandler(db)).Methods("POST")
	router.Handle("/tasks", tasksHandler(db)).Methods("POST")
	router.Handle("/users", usersHandler(db)).Methods("POST")
}
