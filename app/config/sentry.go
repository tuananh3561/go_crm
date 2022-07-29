package config

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type SentryConfig struct {
	Dsn         string
	Release     string
	Environment string
	Debug       bool
}

func getConfigSentry() SentryConfig {
	var sentryConfig = SentryConfig{}

	sentryConfig.Dsn = GetEnv("SENTRY_DSN", "")
	sentryConfig.Release = GetEnv("APP_NAME", "edu_cms")
	sentryConfig.Environment = GetEnv("APP_ENV", "live")
	sentryConfig.Debug = true
	debug, err := strconv.ParseBool(GetEnv("APP_DEBUG", "true"))
	if err == nil {
		sentryConfig.Debug = debug
	}

	return sentryConfig
}

func SentryInit() SentryConfig {

	var sentryConfig = getConfigSentry()

	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: sentryConfig.Dsn,
		// Either set environment and release here or set the SENTRY_ENVIRONMENT
		// and SENTRY_RELEASE environment variables.
		Release:     sentryConfig.Release,
		Environment: sentryConfig.Environment,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if hint.Context != nil {
				if req, ok := hint.Context.Value(sentry.RequestContextKey).(*http.Request); ok {
					// You have access to the original Request
					fmt.Println(req)
				}
			}
			fmt.Println(event)
			return event
		},
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug:            sentryConfig.Debug,
		AttachStacktrace: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	return sentryConfig
}

func SetUseSentry(service *gin.Engine) {
	service.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	service.Use(func(ctx *gin.Context) {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
		}
		ctx.Next()
	})
}
