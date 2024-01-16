package config

import (
	"os"
)


type Config struct {
	WorkerFrequency      string
	ListenersFilePath string
}

func Get() Config {
	return Config{
		WorkerFrequency: os.Getenv("WORKER_CRON_FREQUENCY"),
		ListenersFilePath: os.Getenv("LISTENERS_FILE_PATH"),
	}
}
