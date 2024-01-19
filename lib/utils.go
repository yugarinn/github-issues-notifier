package lib

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "github-issues-notifier"

func LoadEnvFile() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
    currentWorkDirectory, _ := os.Getwd()
    rootPath := projectName.Find([]byte(currentWorkDirectory))

	var err error

	err = godotenv.Load(string(rootPath) + `/.env`)

    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }
}
