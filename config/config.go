package config

import (
	"os"
)


type Config struct {
	WorkerFrequency           string
	ListenersFilePath         string
	ListenersDatabaseFilePath string
}

func Get() Config {
    return Config{
        WorkerFrequency:           getEnvOrDefault("WORKER_CRON_FREQUENCY", "*/30 * * * *"),
        ListenersFilePath:         getEnvOrDefault("LISTENERS_FILE_PATH", "./listeners.yml"),
        ListenersDatabaseFilePath: getEnvOrDefault("LISTENERS_DATABASE_PATH", "./listeners.db"),
    }
}

func getEnvOrDefault(envVar, defaultValue string) string {
    value := os.Getenv(envVar)

    if len(value) == 0 {
        return defaultValue
    }

    return value
}
