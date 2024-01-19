package main

import (
	"github.com/yugarinn/github-issues-notifier/core"
	"github.com/yugarinn/github-issues-notifier/lib"
	"github.com/yugarinn/github-issues-notifier/worker"
)


func main() {
	lib.LoadEnvFile()

	app := core.BootstrapApplication()
	worker.InitiNotificationWorker(app)
}
