package configs

import (
	"fmt"
	"os"
)

const (
	MindboxEndPointID            = "MINDBOX_END_POINT_ID"
	MindboxOperation             = "MINDBOX_OPERATION_UPDATE_CLIENT"
	MindboxSecretKey             = "MINDBOX_SECRET_KEY"
	MindboxOperationCreateClient = "MINDBOX_OPERATION_CREATE_CLIENT"
)

func MindboxGetUrl(itsNewClient bool) string {
	operation := os.Getenv(MindboxOperation)
	if itsNewClient {
		operation = os.Getenv(MindboxOperationCreateClient)
	}
	return fmt.Sprintf("https://api.mindbox.ru/v3/operations/async?endpointId=%s&operation=%s", os.Getenv(MindboxEndPointID), operation)
}

func MindboxGetSecretKey() string {
	return os.Getenv(MindboxSecretKey)
}
