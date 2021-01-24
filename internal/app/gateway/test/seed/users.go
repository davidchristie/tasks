package seed

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davidchristie/tasks/internal/app/gateway/database"
	"github.com/davidchristie/tasks/internal/app/gateway/entity"
)

type usersRequestBody struct {
	Users []entity.User `json:"users"`
}

func usersHandler(db database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := usersRequestBody{}

		err := json.NewDecoder(r.Body).Decode(&body)
    if err != nil {
			fmt.Println("error:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		for _, user := range body.Users {
			err = db.InsertUser(r.Context(), &user)
			if err != nil {
				fmt.Println("error:", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
			}
		}
	}
}
