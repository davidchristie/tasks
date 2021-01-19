package main

import (
	"log"

	"github.com/davidchristie/tasks/internal/app/gateway/server"
)

func main() {
	srv := server.NewServer()
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
