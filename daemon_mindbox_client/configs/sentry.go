package configs

import "os"

const (
	SentryURL = "SENTRY_URL_MINDBOX_CLIENT"
)

func SentryGetURL() string {
	return os.Getenv(SentryURL)
}
