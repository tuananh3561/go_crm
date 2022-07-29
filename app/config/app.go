package config

import (
	"strconv"
)

type AppConfig struct {
	AppName     string
	AppEnv      string
	AppKey      string
	AppPort     string
	AppDebug    bool
	AppTimeZone string
}

var AppConfigDefault AppConfig

func SetupConfigApp() AppConfig {
	appConfig := AppConfig{}
	appConfig.AppName = GetEnv("APP_NAME", "go_crm")
	appConfig.AppEnv = GetEnv("APP_ENV", "live")
	appConfig.AppKey = GetEnv("APP_KEY", "")
	appConfig.AppPort = GetEnv("APP_PORT", "8085")
	parseBool, _ := strconv.ParseBool(GetEnv("APP_DEBUG", "true"))
	appConfig.AppDebug = parseBool
	appConfig.AppTimeZone = GetEnv("APP_TIMEZONE", "Asia/Ho_Chi_Minh")

	AppConfigDefault = appConfig
	return appConfig
}
