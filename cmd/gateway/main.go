package main

import (
	"log"

	"github.com/davidchristie/tasks/internal/app/gateway/server"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
