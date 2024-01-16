package core

import (
	"github.com/yugarinn/github-issues-notificator/config"
	"github.com/yugarinn/github-issues-notificator/lib"
)


type App struct {
	Config config.Config
	GithubClient *lib.GithubClient
	EmailClient *lib.EmailClient
}

func BootstrapApplication() *App {
    app := App{
		Config: config.Get(),
		GithubClient: &lib.GithubClient{},
		EmailClient: &lib.EmailClient{},
    }

	return &app
}
