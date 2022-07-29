package config

import (
	"github.com/joho/godotenv"
	"os"
)

// LoadEnv Load will read your env file(s) and load them into ENV for this process.
func LoadEnv() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		panic("Failed to load env failed")
	}
}

// GetEnv get key environment variable if exist otherwise return defaultValue
func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

var Environment *EnvironmentData

type EnvironmentData struct {
	ApiCrm  string
	ApiAuth string
}

func LoadEnvironment() {
	appEnv := ""
	if AppConfigDefault.AppEnv != "" {
		appEnv = AppConfigDefault.AppEnv
	} else {
		appEnv = GetEnv("APP_ENV", "local")
	}

	data := EnvironmentData{}
	switch appEnv {
	case "live":
		data = EnvironmentLive()
		break
	case "dev":
		data = EnvironmentDev()
		break
	case "product":
		data = EnvironmentProduct()
		break
	default:
		data = EnvironmentLocal()
	}
	Environment = &data
}

func EnvironmentLocal() EnvironmentData {
	env := EnvironmentData{}
	env.ApiCrm = GetEnv("API_SERVICE_DEVELOP", "")
	env.ApiAuth = GetEnv("API_SERVICE_AUTH", "")
	return env
}

func EnvironmentDev() EnvironmentData {
	env := EnvironmentData{}
	env.ApiCrm = "https://api.dev.monkeyuni.net"
	env.ApiAuth = "https://auth.dev.monkeyuni.com/"
	return env
}

func EnvironmentProduct() EnvironmentData {
	env := EnvironmentData{}
	env.ApiCrm = "https://api.dev.monkeyuni.net"
	env.ApiAuth = "https://auth.dev.monkeyuni.com/"
	return env
}

func EnvironmentLive() EnvironmentData {
	env := EnvironmentData{}
	env.ApiCrm = "https://www.api.monkeyuni.net/"
	env.ApiAuth = "https://auth.monkeyuni.net/"
	return env
}
