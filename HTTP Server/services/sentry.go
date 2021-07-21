package services

import (
	"log"
	"time"

	"github.com/GodCratos/mindbox/configs"
	"github.com/getsentry/sentry-go"
)

func SentryPushMessage(msg string) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: configs.SentryGetURL(),
	})
	if err != nil {
		log.Println("[SERVICES:SENTRY] Sentry initialization error. Error description : ", err)
	}
	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage(msg)
}
