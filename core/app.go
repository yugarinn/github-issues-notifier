package core

import (
	"github.com/yugarinn/github-issues-notifier/config"
)


type App struct {
	Config config.Config
}

func BootstrapApplication() *App {
    app := App{
		Config: config.Get(),
    }

	return &app
}
