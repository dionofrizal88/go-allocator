package config

import (
	"os"
	"strconv"
)

// Configuration struct is used to hold environment configuration.
type Configuration struct {
	AppEnv        string
	AppName       string
	AppPort       string
	RedisHost     string
	RedisPassword string
	RedisPort     string
	RedisDB       int

	AgentAllocatorWorkerSleep     int
	AgentAllocatorWorkerMaxAssign int

	QiscusUsername  string
	QiscusPassword  string
	QiscusBaseURL   string
	QiscusAppID     string
	QiscusSecretKey string
}

// getEnvAsInt gets an environment variable and converts it to int.
func getEnvAsInt(key string, defaultVal int) int {
	valStr := os.Getenv(key)
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultVal
}

// GetConfig returns a Configuration struct populated from environment variables.
func GetConfig() Configuration {
	return Configuration{
		AppEnv:        os.Getenv("APP_ENV"),
		AppName:       os.Getenv("APP_NAME"),
		AppPort:       os.Getenv("PORT"),
		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		RedisDB:       getEnvAsInt("REDIS_DB", 0),

		AgentAllocatorWorkerSleep:     getEnvAsInt("AGENT_ALLOCATOR_WORKER_SLEEP", 1),
		AgentAllocatorWorkerMaxAssign: getEnvAsInt("AGENT_ALLOCATOR_WORKER_MAX_ASSIGN", 1),

		QiscusUsername:  os.Getenv("QISCUS_USERNAME"),
		QiscusPassword:  os.Getenv("QISCUS_PASSWORD"),
		QiscusBaseURL:   os.Getenv("QISCUS_BASE_URL"),
		QiscusAppID:     os.Getenv("QISCUS_APP_ID"),
		QiscusSecretKey: os.Getenv("QISCUS_SECRET_KEY"),
	}
}
