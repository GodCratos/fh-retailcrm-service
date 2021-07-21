package configs

import (
	"fmt"
	"os"
)

const (
	HTTPPort = "HTTP_PORT"
)

//HTTPGetPort is func to get port http
func HTTPGetPort() string {
	return fmt.Sprintf(":%v", os.Getenv(HTTPPort))
}
