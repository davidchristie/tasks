package seed

import (
	"fmt"
	"net/http"

	"github.com/davidchristie/tasks/internal/app/gateway/database"
)

func resetHandler(db database.Database) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		err := db.DeleteAllTasks()
    if err != nil {
			fmt.Println("error:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = db.DeleteAllUsers()
		if err != nil {
			fmt.Println("error:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
