package configs

import (
	"fmt"
	"os"
)

const (
	MindboxEndPointID           = "MINDBOX_END_POINT_ID"
	MindboxOperation            = "MINDBOX_OPERATION_UPDATE_ORDER"
	MindboxSecretKey            = "MINDBOX_SECRET_KEY"
	MindboxOperationCreateOrder = "MINDBOX_OPERATION_CREATE_ORDER"
)

func MindboxGetUrl(itsNewOrder bool) string {
	operation := os.Getenv(MindboxOperation)
	if itsNewOrder {
		operation = os.Getenv(MindboxOperationCreateOrder)
	}
	return fmt.Sprintf("https://api.mindbox.ru/v3/operations/async?endpointId=%s&operation=%s", os.Getenv(MindboxEndPointID), operation)
}

func MindboxGetSecretKey() string {
	return os.Getenv(MindboxSecretKey)
}
