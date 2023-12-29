package core

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/yugarinn/github-issues-notificator/database"
	"github.com/yugarinn/github-issues-notificator/lib"
)


type App struct {
	Database *mongo.Database
	GithubClient *lib.GithubClient
	EmailClient *lib.EmailClient
}

func BootstrapApplication() *App {
    app := App{
		Database: database.Database(),
		GithubClient: &lib.GithubClient{},
		EmailClient: &lib.EmailClient{},
    }

	return &app
}
