package seed

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davidchristie/tasks/internal/app/gateway/database"
	"github.com/davidchristie/tasks/internal/app/gateway/entity"
)

type tasksRequestBody struct {
	Tasks []entity.Task `json:"tasks"`
}

func tasksHandler(db database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := tasksRequestBody{}

		err := json.NewDecoder(r.Body).Decode(&body)
    if err != nil {
			fmt.Println("error:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		for _, task := range body.Tasks {
			err = db.InsertTask(r.Context(), &task)
			if err != nil {
				fmt.Println("error:", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
			}
		}
	}
}
