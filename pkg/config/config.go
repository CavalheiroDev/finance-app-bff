package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	APP_ENV_PRODUCTION  string = "production"
	APP_ENV_SANDBOX     string = "sandbox"
	APP_ENV_DEVELOPMENT string = "development"
	APP_ENV_TEST        string = "test"
)

var (
	APP_NAME          string
	APP_PORT          int
	APP_VERSION       string
	APP_SERVER_HEADER string
	APP_ENV           string
	DB_HOST           string
	DB_PORT           int
	DB_USER           string
	DB_PASSWORD       string
	DB_NAME           string
)

func Load() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	APP_NAME = os.Getenv("APP_NAME")
	convertIntEnv(&APP_PORT, "APP_PORT")
	APP_VERSION = os.Getenv("APP_VERSION")
	APP_SERVER_HEADER = os.Getenv("APP_SERVER_HEADER")
	APP_ENV = os.Getenv("APP_ENV")

	DB_HOST = os.Getenv("DB_HOST")
	convertIntEnv(&DB_PORT, "DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")

	return nil
}

func convertIntEnv(env *int, envName string) {
	if envString := os.Getenv(envName); envString != "" {
		var err error
		if *env, err = strconv.Atoi(envString); err != nil {
			log.Fatalf("could not parse %s, permitted int value, got %v: %w", envName, envString, err)
		}
	}
}

func IsProductionEnvironment() bool {
	return APP_ENV == APP_ENV_PRODUCTION
}

func IsSandboxEnvironment() bool {
	return APP_ENV == APP_ENV_SANDBOX
}

func IsDevelopmentEnvironment() bool {
	return APP_ENV == APP_ENV_DEVELOPMENT
}

func IsTestEnvironment() bool {
	return APP_ENV == APP_ENV_TEST
}
