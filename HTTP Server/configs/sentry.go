package configs

import "os"

const (
	SentryURL = "SENTRY_URL_HTTP_SERVER"
)

func SentryGetURL() string {
	return os.Getenv(SentryURL)
}
