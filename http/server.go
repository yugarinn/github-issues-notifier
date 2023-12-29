package http

import (
	"log"
	"net/http"

	"github.com/yugarinn/github-issues-notificator/core"
)


// TODO: move this to .env
const HTTP_PORT = ":8081"

func InitServer(app *core.App) {
	log.Println("starting server...")

	mux := http.NewServeMux()
	mux.HandleFunc("/notifications", func(w http.ResponseWriter, r *http.Request) {
		createNotificationHandler(app, w, r)
	})

	handler := MiddlewareStack(mux)

	log.Fatal(http.ListenAndServe(HTTP_PORT, handler))
}
