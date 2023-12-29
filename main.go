package main

import (
    "os"
    "os/signal"
    "syscall"

	"github.com/yugarinn/github-issues-notificator/core"
	"github.com/yugarinn/github-issues-notificator/lib"
	"github.com/yugarinn/github-issues-notificator/http"
	"github.com/yugarinn/github-issues-notificator/worker"
)


func main() {
	lib.LoadEnvFile()
	app := core.BootstrapApplication()

	go worker.InitiNotificationWorker(app)

    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    <-sigs

    os.Exit(0)
}
