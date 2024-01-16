package worker

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"

	"github.com/yugarinn/github-issues-notificator/core"
	"github.com/yugarinn/github-issues-notificator/internal"
)


func InitiNotificationWorker(app *core.App) {
	c := cron.New()
	cronFrequency := app.Config.WorkerFrequency

	cronID, cronError := c.AddFunc(cronFrequency, func() {
		log.Println("checking for new issues...")

		listeners, err := internal.LoadListeners(app)
		if err != nil {
			log.Println("error during listeners file parsing, aborting execution:", err)
		}

		internal.SendNotifications(app, listeners)

		log.Println("checking for new issues... done")
	})

	if cronError != nil {
		log.Println("error during worker boot, aborting execution:", cronError)
		os.Exit(1)
	}

	c.Start()
	log.Println("started worker on entry", cronID)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-ch)

	c.Stop()
	log.Println("killing woker... done")
}
