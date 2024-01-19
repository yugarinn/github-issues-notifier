package core

import (
	"github.com/yugarinn/github-issues-notifier/config"
	"github.com/yugarinn/github-issues-notifier/lib"
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
