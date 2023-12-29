package http

import (
	"log"
	"net/http"
)


func MiddlewareStack(next http.Handler) http.Handler {
	return logRequest(next)
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
