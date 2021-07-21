package configs

import "os"

const (
	SentryURL = "SENTRY_URL_RETAIL_CLIENT"
)

func SentryGetURL() string {
	return os.Getenv(SentryURL)
}
