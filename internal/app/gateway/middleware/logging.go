package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func Logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return handlers.LoggingHandler(logger.Writer(), next)
	}
}
