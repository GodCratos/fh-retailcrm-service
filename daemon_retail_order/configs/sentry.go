package configs

import "os"

const (
	SentryURL = "SENTRY_URL_RETAIL_ORDER"
)

func SentryGetURL() string {
	return os.Getenv(SentryURL)
}
