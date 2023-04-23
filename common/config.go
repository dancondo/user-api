package common

import (
	"os"
	"strconv"
)

var configuration = map[string]string{
	"APP_ENV":           "dev",
	"APP_SERVICE_TOKEN": "xpto",
	"APP_NAME":          "users-api",

	// LOG LEVEL
	"LOG_LEVEL": "debug",

	// HTTP CONFIGURATION
	"HTTP_PORT":                           "4040",
	"HTTP_SERVER_READ_TIMEOUT_SECONDS":    "30",
	"HTTP_SERVER_WRITE_TIMEOUT_SECONDS":   "30",
	"HTTP_SERVER_MAX_IDLE_CONNS":          "3",
	"HTTP_SERVER_MAX_IDLE_CONNS_PER_HOST": "2",
}

// GetEnv
// Gets an env variable
func GetEnv(configKey string) string {
	return configuration[configKey]
}

// GetEnvInt
// Gets env variable as int
func GetEnvInt(configKey string) int {
	i := 0
	i, _ = strconv.Atoi(GetEnv(configKey))
	return i
}

// GetEnvFloat64
// Gets env variable as float64
func GetEnvFloat64(configKey string) float64 {
	var i float64 = 0.0 //nolint:ineffassign
	i, _ = strconv.ParseFloat(GetEnv(configKey), 64)
	return i
}

// GetEnvFloat32
// Gets env variable as float32
func GetEnvFloat32(configKey string) float32 {
	var i float32 = 0.0
	val, err := strconv.ParseFloat(GetEnv(configKey), 32)

	if err != nil {
		return i
	}

	return float32(val)
}

func GetEnvBool(configKey string) bool {
	val, err := strconv.ParseBool(GetEnv(configKey))
	if err != nil {
		return false
	}
	return val
}

// SetEnv
// Sets an env variable
func SetEnv(configKey string, value string) {
	configuration[configKey] = value
}

// LoadEnv
// Set default value from config map, else load all setted environment variables.
func LoadEnv() {
	for k := range configuration {
		v := os.Getenv(k)
		if v != "" {
			configuration[k] = v
		}
	}
}
