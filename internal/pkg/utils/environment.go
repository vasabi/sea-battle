package utils

import (
	"github.com/spf13/viper"
)

func bindEnv(env string) func() error {
	return func() error {
		return viper.BindEnv(env)
	}
}

// InitEnvironmentSettings binds environment variables.
func InitEnvironmentSettings() (err error) {
	bindings := []func() error{
		bindEnv("DEBUG"),
		bindEnv("HTTP_PORT"),
		bindEnv("SWAGGER_INFO_HOST"),
	}

	for _, bind := range bindings {
		if err = bind(); err != nil {
			return
		}
	}

	return
}

// GetDebugModeON returns true if the DEBUG environment variable is set to true. Default false.
func GetDebugModeON() bool {
	return viper.GetBool("DEBUG")
}

// GetHTTPPort returns the HTTP_PORT environment variable. Default 8080.
func GetHTTPPort() string {
	return getFromEnvOrDefault("HTTP_PORT", "8080")
}

// GetSwaggerInfoHost returns the SWAGGER_INFO_HOST environment variable. Default localhost:8080.
func GetSwaggerInfoHost() string {
	return getFromEnvOrDefault("SWAGGER_INFO_HOST", "localhost:8080")
}

func getFromEnvOrDefault(env string, defaultVal string) string {
	val := viper.GetString(env)

	if len(val) == 0 {
		return defaultVal
	}

	return val
}
