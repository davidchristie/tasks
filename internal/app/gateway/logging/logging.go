package logging

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func Middleware(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return handlers.LoggingHandler(logger.Writer(), next)
	}
}
