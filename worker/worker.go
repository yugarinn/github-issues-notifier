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
	"github.com/yugarinn/github-issues-notificator/lib"
)


func InitiNotificationWorker(app *core.App) {
	lib.LoadEnvFile()

	c := cron.New()
	cronFrequency := os.Getenv("WORKER_CRON_FREQUENCY")

	cronID, cronError := c.AddFunc(cronFrequency, func() {
		log.Println("checking for new issues...")
		internal.SendNotifications(app)
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
	log.Println("starting woker... done")
}
