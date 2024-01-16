package main

import (
	"github.com/yugarinn/github-issues-notificator/core"
	"github.com/yugarinn/github-issues-notificator/lib"
	"github.com/yugarinn/github-issues-notificator/worker"
)


func main() {
	lib.LoadEnvFile()

	app := core.BootstrapApplication()
	worker.InitiNotificationWorker(app)
}
