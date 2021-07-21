package configs

import (
	"fmt"
	"os"
)

const (
	RetailURL    = "RETAIL_URL_GET_ORDER"
	RetailDomain = "CRM_DOMAIN"
	RetailApiKey = "CRM_API_KEY"
)

func RetailGetURL(clientID string) string {
	return fmt.Sprintf(os.Getenv(RetailURL), os.Getenv(RetailDomain), os.Getenv(RetailApiKey), clientID)
}
