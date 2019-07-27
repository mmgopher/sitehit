package config

import (
	"os"
	"strconv"
	"sync"
)


type Configuration struct {
	HttpAddr string
	WindowTime int
	StorageFile string
}

var instance *Configuration
var once sync.Once

func GetConfiguration() *Configuration {
	once.Do(func() {
		instance = &Configuration{
			HttpAddr: getEnvDefault("HTTP_ADDR", ":8081"),
			WindowTime:getEnvIntDefault("WINDOW_TIME", 60),
			StorageFile: getEnvDefault("STORAGE_FILE", "storage.gob"),
		}

	})

	return instance
}


func getEnvDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvIntDefault(key string, fallback int) int {

	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	intValue, err := strconv.Atoi(value)
	if err != nil  {
		intValue = fallback
	}
	return intValue
}
