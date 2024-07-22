package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Backend struct {
		Port     int
		BasePath string
	}
	Frontend struct {
		BasePath string
	}
}

func LoadEnvs() *Config {
	var config Config

	loadServerPort(&config)
	loadBackendAPIBasePath(&config)
	loadFrontendBasePath(&config)

	flag.Parse()

	return &config
}

func loadServerPort(config *Config) {
	portDefaultValue := 3000
	if port := getEnvInt("LOCAL_TO_DO_MANAGER_BACKEND_SERVER_PORT"); port != nil {
		portDefaultValue = *port
	}
	flag.IntVar(&config.Backend.Port, "back-port", portDefaultValue, "Backend API sever Port")
}

func loadBackendAPIBasePath(config *Config) {
	basePathDefaultValue := "/api"
	if basePath := getEnvStr("LOCAL_TO_DO_MANAGER_BACKEND_SERVER_API_BASE_PATH"); basePath != nil {
		basePathDefaultValue = *basePath
	}
	flag.StringVar(&config.Backend.BasePath, "back-basePath", basePathDefaultValue, "Backend API sever base Path")
}

func loadFrontendBasePath(config *Config) {
	basePathDefaultValue := "/app"
	if basePath := getEnvStr("LOCAL_TO_DO_MANAGER_FRONTEND_BASE_PATH"); basePath != nil {
		basePathDefaultValue = *basePath
	}
	flag.StringVar(&config.Frontend.BasePath, "front-basePath", basePathDefaultValue, "Frontend base Path")
}

func getEnvStr(key string) *string {
	strValue := os.Getenv(key)
	if strValue == "" {
		return nil
	}
	return &strValue
}

func getEnvInt(key string) *int {
	strValue := os.Getenv(key)
	if strValue == "" {
		return nil
	}
	converted, err := strconv.Atoi(strValue)
	if err != nil {
		panic(fmt.Sprintf("could not convert env value. Env Name: %s - Env Value: %s", key, strValue))
	}
	return &converted
}
